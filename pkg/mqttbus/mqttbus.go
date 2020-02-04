package mqttbus

/*
 * Copyright (c) 2013 IBM Corp.
 *
 * All rights reserved. This program and the accompanying materials
 * are made available under the terms of the Eclipse Public License v1.0
 * which accompanies this distribution, and is available at
 * http://www.eclipse.org/legal/epl-v10.html
 *
 * Contributors:
 *    Seth Hoenig
 *    Allan Stockdill-Mander
 *    Mike Robertson
 */
import (
	"encoding/json"
	"errors"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/shikanon/IoTOrbHub/pkg/constants"
	"github.com/shikanon/IoTOrbHub/pkg/database"
	"github.com/shikanon/IoTOrbHub/pkg/influxdb"
	"github.com/shikanon/IoTOrbHub/pkg/tool"
	"github.com/shikanon/IoTOrbHub/pkg/util"
	logs "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	MQTTHub *Client
	// TODO 用户认证，允许订阅SYS主题消息
	MQTTUrl = "10.10.6.53:1883"
	// TODO: 从数据库查询所有产品对应topic
	SubTopics = []string{
		// 设备上下线
		"$SYS/broker/connection/clients/+",
		// 设备上报属性
		"/sys/+/+/thing/event/property/post",
		// 设置设备属性
		"/sys/+/+/thing/service/property/set_reply",
		// 设备事件上报
		"/sys/+/+/thing/event/+/post",
		// 设备服务调用（异步调用）
		"/sys/+/+/thing/service/+",
		// 网关批量上报数据
		"/sys/+/+/thing/event/property/pack/post",
		// 物模型历史数据上报
		"/sys/+/+/thing/event/property/history/post",
		// 设备获取期望属性值
		"/sys/+/+/thing/property/desired/get",
		// 设备清空期望属性值
		"/sys/+/+/thing/property/desired/delete",
		// 设备标签信息上报
		"/sys/+/+/thing/deviceinfo/update",
		// 设备删除标签消息
		"/sys/+/+/thing/deviceinfo/delete",
		// 更新设备影子
		"/shadow/update/+/+",
	}
)

// Client struct
type Client struct {
	MQTTUrl string
	PubCli  MQTT.Client
	SubCli  MQTT.Client
}

type ConnectionMsg struct {
	ClientID  string `json:"clientID"`
	Online    bool   `json:"online"`
	TimeStamp string `json:"timestamp"`
}

type DeviceMsg struct {
	Id      string      `json:"id"`
	Version string      `json:"version"`
	Params  interface{} `json:"params"`
	Method  string      `json:"method"`
}

type ReplyMsg struct {
	Id      string                 `json:"id"`
	Code    int                    `json:"code"`
	Data    map[string]interface{} `json:"data"`
	Message string                 `json:"message"`
	Event   string                 `json:"event"`
}

type ParamsParse struct {
	Version   string `json:"version"`
	AttrKey   string `json:"attrKey"`
	AttrValue string `json:"attrValue"`
	Value     string `json:"value"`
}

type ShadowMsg struct {
	Method    string      `json:"method"`
	Version   int64       `json:"version"`
	Payload   interface{} `json:"payload"`
	Timestamp int64       `json:"timestamp"`
}

func onPubConnectionLost(client MQTT.Client, err error) {
	logs.Errorf("onPubConnectionLost with error: %v", err)
	go MQTTHub.InitPubClient()
}

func onSubConnectionLost(client MQTT.Client, err error) {
	logs.Errorf("onSubConnectionLost with error: %v", err)
	go MQTTHub.InitSubClient()
}

func onSubConnect(client MQTT.Client) {
	for _, t := range SubTopics {
		token := client.Subscribe(t, 1, OnSubMessageReceived)
		if rs, err := util.CheckClientToken(token); !rs {
			logs.Errorf("edge-hub-cli subscribe topic: %s, %v", t, err)
			return
		}
		logs.Infof("edge-hub-cli subscribe topic to %s", t)
	}
}

// OnSubMessageReceived msg received callback
func OnSubMessageReceived(client MQTT.Client, message MQTT.Message) {
	logs.Infof("OnSubMessageReceived receive msg from topic: %s,msg: %s", message.Topic(), message.Payload())
	if reg, err := regexp.MatchString("\\$SYS/broker/connection/clients/(?P<clientID>.+)", message.Topic()); reg == true && err == nil {
		// 设备上下线
		var connectionMsg ConnectionMsg
		err := json.Unmarshal(message.Payload(), &connectionMsg)
		if err != nil {
			logs.Error(err)
		}
		if strings.HasPrefix(connectionMsg.ClientID, "hub-client") {
			return
		}
		productKey, deviceName, err := util.GetDeviceDetailFromClientID(connectionMsg.ClientID)
		if err != nil {
			logs.Error(err)
		}
		if connectionMsg.Online == true {
			User := make(map[string]interface{})
			logs.Info(fmt.Sprintf("设备上线:%s&%s", deviceName, productKey))
			ts := time.Now()
			device := database.DeviceNameToDevice(productKey, deviceName)
			if tool.TimeDeal(device.ActivationTime) == "-" {
				User["ActivationTime"] = ts
			}
			User["LastOnLineTime"] = ts
			db := database.DbConn()
			db.Model(&device).Updates(User)
			defer db.Close()
		} else {
			logs.Info(fmt.Sprintf("设备下线:%s", connectionMsg.ClientID))
		}
		// 消息传递给controller
		//event := deviceController.Event{Operate: 1, Topic:message.Topic(), Message:message.Payload()}
		//deviceController.DeviceController.Manager.SendMessage(&event)
	} else if reg, err := regexp.MatchString("/sys/(?P<productKey>.+)/(?P<deviceName>.+)/thing/event/property/post", message.Topic()); reg == true && err == nil {
		// 设备上报属性
		logs.Info("设备上报属性处理逻辑")
		var replyMsg ReplyMsg
		data := make(map[string]interface{})
		fields := make(map[string]interface{})
		replyMsg.Code = constants.Success
		replyMsg.Message = constants.SuccessMsg
		productKey, deviceName, err := util.GetDeviceDetailFromTopic(message.Topic())
		if err != nil {
			logs.Error(err)
			replyMsg.Code = constants.RequestParameterError
			replyMsg.Message = constants.RequestParameterErrorMsg
		}
		fmt.Println(productKey, deviceName)
		id := gjson.GetBytes(message.Payload(), "id")
		if !id.Exists() || len(id.String()) == 0 {
			logs.Error("缺少参数: id")
			replyMsg.Code = constants.RequestParameterError
			replyMsg.Message = constants.RequestParameterErrorMsg
		}
		replyMsg.Id = id.String()
		// 数据校验
		device := database.DeviceNameToDevice(productKey, deviceName)
		deviceId := device.IotID
		model := database.GetIntactModel(productKey)
		modelJson, err := json.Marshal(model)
		if err != nil {
			fmt.Println(err)
		}
		params := gjson.GetBytes(message.Payload(), "params")
		for k, v := range params.Map() {
			result := gjson.Get(string(modelJson), fmt.Sprintf("properties.#(identifier==%s)", k))
			if !result.Exists() {
				err := errors.New("tsl parse: params not exist")
				data[k] = err.Error()
				continue
			}
			dataType := result.Get("dataType.type")
			if value, err := util.DataVerification(k, dataType.String(), v, result); err != nil {
				data[k] = err.Error()
			} else {
				fields[k] = value
			}
		}
		if len(fields) != 0 {
			fields["msg_id"] = id.Int()
			influxdb.AddPointToPropertyReported(deviceId, fields)
		}
		replyMsg.Data = data
		res, err := json.Marshal(replyMsg)
		if err != nil {
			logs.Error(err)
			return
		}
		reply := message.Topic() + "_reply"
		MQTTHub.Publish(reply, res)
		//event := deviceController.Event{Operate: 2, Topic:message.Topic(), Message:message.Payload()}
		//deviceController.DeviceController.Manager.SendMessage(&event)
	} else if reg, err := regexp.MatchString("/sys/(?P<productKey>.+)/(?P<deviceName>.+)/thing/service/property/set_reply", message.Topic()); reg == true && err == nil {
		// 设置设备属性reply
		logs.Info("设置设备属性reply处理逻辑")
		productKey, deviceName, err := util.GetDeviceDetailFromTopic(message.Topic())
		if err != nil {
			logs.Error(err)
		}
		fmt.Println(productKey, deviceName)
		code := gjson.GetBytes(message.Payload(), "code")
		if !code.Exists() || len(code.String()) == 0 {
			logs.Error("缺少参数: code")
		}
		//deviceId := "1" //从数据库查
		//msgId := gjson.GetBytes(message.Payload(), "id")
		//if !code.Exists() || len(code.String()) == 0 {
		//	logs.Error("缺少参数: id")
		//}
		// TODO: 如果code不为200，将信息同步至对应服务
		if code.Int() == 200 {
			fmt.Println("sent set property Msg ok")

		}
		//event := deviceController.Event{Operate: 3, Topic:message.Topic(), Message:message.Payload()}
		//deviceController.DeviceController.Manager.SendMessage(&event)
	} else if reg, err := regexp.MatchString("/sys/(?P<productKey>.+)/(?P<deviceName>.+)/thing/event/(?P<event>.+)/post", message.Topic()); reg == true && err == nil {
		// 设备事件上报
		logs.Info("设备事件上报处理逻辑")
		var replyMsg ReplyMsg
		replyMsg.Code = constants.Success
		replyMsg.Message = constants.SuccessMsg
		replyMsg.Data = make(map[string]interface{})
		productKey, deviceName, err := util.GetDeviceDetailFromTopic(message.Topic())
		if err != nil {
			logs.Error(err)
			replyMsg.Code = constants.RequestParameterError
			replyMsg.Message = constants.RequestParameterErrorMsg
		}
		event, err := util.GetEventlFromTopic(message.Topic())
		if err != nil {
			logs.Error(err)
			replyMsg.Code = constants.RequestParameterError
			replyMsg.Message = constants.RequestParameterErrorMsg
		}
		replyMsg.Event = event
		fmt.Println(productKey, deviceName, event)
		id := gjson.GetBytes(message.Payload(), "id")
		if !id.Exists() || len(id.String()) == 0 {
			logs.Error("缺少参数: id")
			replyMsg.Code = constants.RequestParameterError
			replyMsg.Message = constants.RequestParameterErrorMsg
		}
		replyMsg.Id = id.String()
		// 数据校验
		device := database.DeviceNameToDevice(productKey, deviceName)
		deviceId := device.IotID
		model := database.GetIntactModel(productKey)
		modelJson, err := json.Marshal(model)
		if err != nil {
			fmt.Println(err)
		}
		method := gjson.GetBytes(message.Payload(), "method")
		rs := gjson.Get(string(modelJson), fmt.Sprintf("events.#(method==%s)", method.String()))
		if !rs.Exists() {
			replyMsg.Message = "event not exist"
		} else {
			outputData := make(map[string]interface{})
			params := gjson.GetBytes(message.Payload(), "params.value")
			for k, v := range params.Map() {
				result := rs.Get(fmt.Sprintf("outputData.#(identifier==%s)", k))
				if !result.Exists() {
					replyMsg.Message = fmt.Sprintf("tsl parse: params not exist -> %s", k)
					break
				}
				dataType := result.Get("dataType.type")
				if value, err := util.DataVerification(k, dataType.String(), v, result); err != nil {
					replyMsg.Message = fmt.Sprintf("tsl parse: %s specs error -> %s", dataType.String(), k)
					break
				} else {
					outputData[k] = value
				}
			}
			if len(outputData) != 0 {
				dataStr, err := json.Marshal(outputData)
				if err == nil {
					service := map[string]interface{}{
						"identifier": event,
						"event_name": rs.Get("name").String(),
						"event_type": rs.Get("type").String(),
						"outputData": string(dataStr),
						"msg_id":     id.Int(),
					}
					influxdb.AddPointToEvent(deviceId, service)
				} else {
					logs.Error(err)
				}
			}
		}
		res, err := json.Marshal(replyMsg)
		if err != nil {
			logs.Error(err)
			return
		}
		reply := message.Topic() + "_reply"
		MQTTHub.Publish(reply, res)
		//event := deviceController.Event{Operate: 4, Topic:message.Topic(), Message:message.Payload()}
		//deviceController.DeviceController.Manager.SendMessage(&event)
	} else if reg, err := regexp.MatchString("/sys/(?P<productKey>.+)/(?P<deviceName>.+)/thing/service/(?P<identifier>.+_reply)", message.Topic()); reg == true && err == nil {
		// 设备服务调用（异步调用）reply
		logs.Info("设备服务调用（异步调用）reply处理逻辑")
		productKey, deviceName, err := util.GetDeviceDetailFromTopic(message.Topic())
		if err != nil {
			logs.Error(err)
		}
		service, err := util.GetServicelFromTopic(message.Topic()) // service可根据消息id查询数据库获得
		if err != nil {
			logs.Error(err)
		}
		fmt.Println(productKey, deviceName, service)
		code := gjson.GetBytes(message.Payload(), "code")
		if !code.Exists() || len(code.String()) == 0 {
			logs.Error("缺少参数: code")
		}
		// TODO: 设备服务调用（异步调用）reply
		if code.Int() == 200 {
			fmt.Println("sent call service Msg ok")
		}
		//event := deviceController.Event{Operate: 5, Topic:message.Topic(), Message:message.Payload()}
		//deviceController.DeviceController.Manager.SendMessage(&event)
	} else if reg, err := regexp.MatchString("/sys/(?P<productKey>.+)/(?P<deviceName>.+)/thing/event/property/pack/post", message.Topic()); reg == true && err == nil {
		// 网关批量上报数据
		logs.Info("网关批量上报数据处理逻辑")
	} else if reg, err := regexp.MatchString("/sys/(?P<productKey>.+)/(?P<deviceName>.+)/thing/event/property/history/post", message.Topic()); reg == true && err == nil {
		// 物模型历史数据上报
		logs.Info("物模型历史数据上报处理逻辑")
	} else if reg, err := regexp.MatchString("/sys/(?P<productKey>.+)/(?P<deviceName>.+)/thing/property/desired/get", message.Topic()); reg == true && err == nil {
		// 设备获取期望属性值
		logs.Info("设备获取期望属性值处理逻辑")
		var replyMsg ReplyMsg
		replyMsg.Code = constants.Success
		replyMsg.Message = constants.SuccessMsg
		productKey, deviceName, err := util.GetDeviceDetailFromTopic(message.Topic())
		if err != nil {
			logs.Error(err)
			replyMsg.Code = constants.RequestParameterError
			replyMsg.Message = constants.RequestParameterErrorMsg
		}
		fmt.Println(productKey, deviceName)
		id := gjson.GetBytes(message.Payload(), "id")
		if !id.Exists() || len(id.String()) == 0 {
			logs.Error("缺少参数: id")
			replyMsg.Code = constants.RequestParameterError
			replyMsg.Message = constants.RequestParameterErrorMsg
		}
		replyMsg.Id = id.String()
		params := gjson.GetBytes(message.Payload(), "params")
		if !params.Exists() || len(params.Array()) == 0 {
			logs.Error("缺少参数: params")
			replyMsg.Code = constants.RequestParameterError
			replyMsg.Message = constants.RequestParameterErrorMsg
		}
		dataMap := make(map[string]interface{})
		var paramsParse ParamsParse
		for _, v := range params.Array() {
			fmt.Println(v)
			// TODO: 查询prams期望属性值，构造响应数据
			// 		data:= {
			//		 "power": {
			//		   "value": "on",
			//		   "version": 2
			//		 }
			//		}
			paramsParse.Value = ""
			paramsParse.Version = ""
			dataMap[v.String()] = paramsParse
		}
		replyMsg.Data = dataMap
		res, err := json.Marshal(replyMsg)
		if err != nil {
			logs.Error(err)
		}
		reply := message.Topic() + "_reply"
		MQTTHub.Publish(reply, res)

		//event := deviceController.Event{Operate: 8, Topic:message.Topic(), Message:message.Payload()}
		//deviceController.DeviceController.Manager.SendMessage(&event)
	} else if reg, err := regexp.MatchString("/sys/(?P<productKey>.+)/(?P<deviceName>.+)/thing/property/desired/delete", message.Topic()); reg == true && err == nil {
		// 设备清空期望属性值
		logs.Info("设备清空期望属性值处理逻辑")
		var replyMsg ReplyMsg
		replyMsg.Code = constants.Success
		replyMsg.Message = constants.SuccessMsg
		replyMsg.Data = make(map[string]interface{})
		productKey, deviceName, err := util.GetDeviceDetailFromTopic(message.Topic())
		if err != nil {
			logs.Error(err)
		}
		fmt.Println(productKey, deviceName)
		id := gjson.GetBytes(message.Payload(), "id")
		if !id.Exists() || len(id.String()) == 0 {
			logs.Error("缺少参数: id")
			replyMsg.Code = constants.RequestParameterError
			replyMsg.Message = constants.RequestParameterErrorMsg
		}
		replyMsg.Id = id.String()
		params := gjson.GetBytes(message.Payload(), "params")
		if !params.Exists() || len(params.Map()) == 0 {
			logs.Error("缺少参数: params")
			replyMsg.Code = constants.RequestParameterError
			replyMsg.Message = constants.RequestParameterErrorMsg
		}
		//  	params = {
		//  		"power": {
		//		   		"version": 1
		//		 		},
		//		 	"temperature": { }
		//		}
		for k, v := range params.Map() {
			version := v.Get("version")
			fmt.Println(k, version)
			// TODO: if version == null or version == versionLast 则清空属性k
		}
		res, err := json.Marshal(replyMsg)
		if err != nil {
			logs.Error(err)
		}
		reply := message.Topic() + "_reply"
		MQTTHub.Publish(reply, res)
		//event := deviceController.Event{Operate: 9, Topic:message.Topic(), Message:message.Payload()}
		//deviceController.DeviceController.Manager.SendMessage(&event)
	} else if reg, err := regexp.MatchString("/sys/(?P<productKey>.+)/(?P<deviceName>.+)/thing/deviceinfo/update", message.Topic()); reg == true && err == nil {
		// 设备标签信息上报
		logs.Info("设备标签信息上报处理逻辑")

		var replyMsg ReplyMsg
		replyMsg.Code = constants.Success
		replyMsg.Message = constants.SuccessMsg
		replyMsg.Data = make(map[string]interface{})
		productKey, deviceName, err := util.GetDeviceDetailFromTopic(message.Topic())
		if err != nil {
			logs.Error(err)
		}
		fmt.Println(productKey, deviceName)
		id := gjson.GetBytes(message.Payload(), "id")
		if !id.Exists() || len(id.String()) == 0 {
			logs.Error("缺少参数: id")
			replyMsg.Code = constants.RequestParameterError
			replyMsg.Message = constants.RequestParameterErrorMsg
		}
		replyMsg.Id = id.String()
		params := gjson.GetBytes(message.Payload(), "params")
		if !params.Exists() || len(params.Array()) == 0 {
			logs.Error("缺少参数: params")
			replyMsg.Code = constants.RequestParameterError
			replyMsg.Message = constants.RequestParameterErrorMsg
		}
		for _, v := range params.Array() {
			attrKey := v.Get("attrKey")
			attrValue := v.Get("attrValue")
			// TODO:保存标签到数据库
			fmt.Println(attrKey, attrValue)
		}
		res, err := json.Marshal(replyMsg)
		if err != nil {
			logs.Error(err)
		}
		reply := message.Topic() + "_reply"
		MQTTHub.Publish(reply, res)
		//event := deviceController.Event{Operate: 10, Topic:message.Topic(), Message:message.Payload()}
		//deviceController.DeviceController.Manager.SendMessage(&event)
	} else if reg, err := regexp.MatchString("/sys/(?P<productKey>.+)/(?P<deviceName>.+)/thing/deviceinfo/delete", message.Topic()); reg == true && err == nil {
		// 设备删除标签消息
		logs.Info("设备删除标签消息处理逻辑")
		var replyMsg ReplyMsg
		replyMsg.Code = constants.Success
		replyMsg.Message = constants.SuccessMsg
		replyMsg.Data = make(map[string]interface{})
		productKey, deviceName, err := util.GetDeviceDetailFromTopic(message.Topic())
		if err != nil {
			logs.Error(err)
		}
		fmt.Println(productKey, deviceName)
		id := gjson.GetBytes(message.Payload(), "id")
		if !id.Exists() || len(id.String()) == 0 {
			logs.Error("缺少参数: id")
			replyMsg.Code = constants.RequestParameterError
			replyMsg.Message = constants.RequestParameterErrorMsg
		}
		replyMsg.Id = id.String()
		params := gjson.GetBytes(message.Payload(), "params")
		if !params.Exists() || len(params.Array()) == 0 {
			logs.Error("缺少参数: params")
			replyMsg.Code = constants.RequestParameterError
			replyMsg.Message = constants.RequestParameterErrorMsg
		}
		for _, v := range params.Array() {
			attrKey := v.Get("attrKey")
			// TODO:删除标签
			fmt.Println(attrKey)
		}
		res, err := json.Marshal(replyMsg)
		if err != nil {
			logs.Error(err)
		}
		reply := message.Topic() + "_reply"
		MQTTHub.Publish(reply, res)
		//event := deviceController.Event{Operate: 11, Topic:message.Topic(), Message:message.Payload()}
		//deviceController.DeviceController.Manager.SendMessage(&event)
	} else if reg, err := regexp.MatchString("/shadow/update/(?P<productKey>.+)/(?P<deviceName>.+)", message.Topic()); reg == true && err == nil {
		// 设备更新影子设备状态
		logs.Info("设备更新影子设备状态处理逻辑")
		productKey, deviceName, err := util.GetDeviceDetailFromShadowTopic(message.Topic())
		if err != nil {
			logs.Error(err)
		}
		fmt.Println(productKey, deviceName)
		var shadowMsg ShadowMsg
		results := gjson.GetManyBytes(message.Payload(), "method", "version", "state")
		method := results[0]
		version := results[1]
		state := results[2]
		versionLast := 1 // 从数据库查
		timestamp := time.Now().Unix()
		shadowMsg.Timestamp = timestamp
		if version.Int() == -1 {
			// TODO: 清空设备影子数据,影子版本置为0
		} else if version.Int() > int64(versionLast) {
			if method.String() == "update" {
				for k, v := range state.Map() {
					switch k {
					case "reported":
						shadowMsg.Method = "reply"
						for k, v := range v.Map() {
							// TODO: 数据上报
							fmt.Println(k, v.String())
						}
						payLoadMap := make(map[string]interface{})
						payLoadMap["status"] = "success"
						payLoadMap["version"] = version.Int()
						shadowMsg.Payload = payLoadMap
					case "desired":
						shadowMsg.Method = "control"
						shadowMsg.Version = version.Int()
						for k, v := range v.Map() {
							// TODO: 期望值
							fmt.Println(k, v.String(), timestamp)
						}
					}
				}
			} else if method.String() == "get" {
				shadowMsg.Method = "reply"
				shadowMsg.Version = int64(versionLast)
				// TODO: 构造payload
			} else if method.String() == "delete" {
				reported := state.Get("reported")
				if !reported.Exists() || len(reported.String()) == 0 {
					logs.Error("缺少参数: reported")
				}
				if reported.String() == "null" {
					// TODO:删除影子全部属性
				} else {
					for k, v := range reported.Map() {
						if v.String() == "null" {
							// TODO:删除影子属性
							fmt.Println("删除影子属性:", k)
						}
					}
				}
			}
		} else {
			shadowMsg.Method = "reply"
			payLoadMap := make(map[string]interface{})
			payLoadMap["status"] = "error"
			errMap := make(map[string]interface{})
			errMap["errorcode"] = constants.VersionConflictError
			errMap["errormessage"] = constants.VersionConflictErrorMsg
			payLoadMap["content"] = errMap
			shadowMsg.Payload = payLoadMap
		}
		res, err := json.Marshal(shadowMsg)
		if err != nil {
			logs.Error(err)
		}
		reply := fmt.Sprintf(constants.ShadowTopic, productKey, deviceName)
		MQTTHub.Publish(reply, res)
		//event := deviceController.Event{Operate: 11, Topic: message.Topic(), Message: message.Payload()}
		//deviceController.DeviceController.Manager.SendMessage(&event)
	}
}

// InitSubClient init sub client
func (mq *Client) InitSubClient() {
	timeStr := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	right := len(timeStr)
	if right > 10 {
		right = 10
	}
	subID := fmt.Sprintf("hub-client-sub-%s", timeStr[0:right])
	subOpts := util.HubClientInit(mq.MQTTUrl, subID, "", "")
	subOpts.OnConnect = onSubConnect
	subOpts.AutoReconnect = false
	subOpts.OnConnectionLost = onSubConnectionLost
	mq.SubCli = MQTT.NewClient(subOpts)
	util.LoopConnect(subID, mq.SubCli)
	logs.Info("finish hub-client sub")
}

// InitPubClient init pub client
func (mq *Client) InitPubClient() {
	timeStr := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	right := len(timeStr)
	if right > 10 {
		right = 10
	}
	pubID := fmt.Sprintf("hub-client-pub-%s", timeStr[0:right])
	pubOpts := util.HubClientInit(mq.MQTTUrl, pubID, "", "")
	pubOpts.OnConnectionLost = onPubConnectionLost
	pubOpts.AutoReconnect = false
	mq.PubCli = MQTT.NewClient(pubOpts)
	util.LoopConnect(pubID, mq.PubCli)
	logs.Info("finish hub-client pub")
}

func (mq *Client) Publish(topic string, payload []byte) {
	token := mq.PubCli.Publish(topic, 1, false, payload)
	if token.WaitTimeout(util.TokenWaitTime) && token.Error() != nil {
		logs.Errorf("Error in pubMQTT with topic: %s, %v", topic, token.Error())
	} else {
		logs.Infof("Success in pubMQTT with topic: %s", topic)
	}
}

func (mq *Client) Subscribe(topic string) {
	// subscribe topic to external mqtt broker.
	token := mq.SubCli.Subscribe(topic, 1, OnSubMessageReceived)
	if rs, err := util.CheckClientToken(token); !rs {
		logs.Errorf("Edge-hub-cli subscribe topic: %s, %v", topic, err)
	}
}

func (mq *Client) SetProperty(productKey, deviceName string, args string) (errData map[string]interface{}) {
	// 设置设备属性
	// TODO：设备是否在线
	// 数据校验
	errData = make(map[string]interface{})
	fields := make(map[string]interface{})
	model := database.GetIntactModel(productKey)
	modelJson, err := json.Marshal(model)
	if err != nil {
		fmt.Println(err)
	}
	res := gjson.Parse(args)
	for k, v := range res.Map() {
		result := gjson.Get(string(modelJson), fmt.Sprintf("services.#(identifier==set).inputData.#(identifier==%s)", k))
		if !result.Exists() {
			err := errors.New("tsl parse: params not exist")
			errData[k] = err.Error()
			continue
		}
		dataType := result.Get("dataType.type")
		if value, err := util.DataVerification(k, dataType.String(), v, result); err != nil {
			errData[k] = err.Error()
		} else {
			fields[k] = value
		}
	}
	if len(fields) != 0 {
		topic := fmt.Sprintf("/sys/%s/%s/thing/service/property/set", productKey, deviceName)
		device := database.DeviceNameToDevice(productKey, deviceName)
		deviceId := device.IotID
		var deviceMsg DeviceMsg
		deviceMsgId := influxdb.GetDeviceMsgIdFromService() //消息ID号，由物联网平台生成。
		deviceMsg.Id = strconv.FormatInt(deviceMsgId, 10)
		deviceMsg.Version = constants.Version
		deviceMsg.Method = constants.SetProperty
		deviceMsg.Params = fields
		deviceMsgStr, err := json.Marshal(deviceMsg)
		if err != nil {
			logs.Error(err)
		}
		payLoad := deviceMsgStr
		mq.Publish(topic, payLoad)
		paramsStr, err := json.Marshal(fields)
		if err == nil {
			service := map[string]interface{}{"identifier": "set", "service_name": "set", "params": string(paramsStr), "msg_id": deviceMsgId}
			influxdb.AddPointToService(deviceId, service)
		} else {
			fmt.Println(err)
		}
		for k, v := range fields {
			version := influxdb.GetPropertyVersionFromPropertyDesired(deviceId, k)
			param := map[string]interface{}{k: v, "version": version, "msg_id": deviceMsgId}
			influxdb.AddPointToPropertyDesired(deviceId, param)
		}
	}
	return errData
}

func (mq *Client) CallService(productKey, deviceName, service string, args string) (errData map[string]interface{}) {
	// 设备服务调用
	// TODO：设备是否在线
	// 数据校验
	errData = make(map[string]interface{})
	fields := make(map[string]interface{})
	model := database.GetIntactModel(productKey)
	modelJson, err := json.Marshal(model)
	if err != nil {
		fmt.Println(err)
	}
	svc := gjson.Get(string(modelJson), fmt.Sprintf("services.#(identifier==%s)", service))
	if !svc.Exists() {
		err := errors.New("tsl parse: service not exist")
		errData[service] = err.Error()
		return
	}
	res := gjson.Parse(args)
	for k, v := range res.Map() {
		result := svc.Get(fmt.Sprintf("inputData.#(identifier==%s)", k))
		if !result.Exists() {
			err := errors.New("tsl parse: params not exist")
			errData[k] = err.Error()
			return
		}
		dataType := result.Get("dataType.type")
		if value, err := util.DataVerification(k, dataType.String(), v, result); err != nil {
			errData[k] = err.Error()
		} else {
			fields[k] = value
		}
	}
	topic := fmt.Sprintf("/sys/%s/%s/thing/service/%s", productKey, deviceName, service)
	device := database.DeviceNameToDevice(productKey, deviceName)
	deviceId := device.IotID
	var deviceMsg DeviceMsg
	deviceMsgId := influxdb.GetDeviceMsgIdFromService() //消息ID号，由物联网平台生成。
	deviceMsg.Id = strconv.FormatInt(deviceMsgId, 10)
	deviceMsg.Version = constants.Version
	deviceMsg.Method = fmt.Sprintf(constants.CallService, service)
	deviceMsg.Params = fields
	deviceMsgStr, err := json.Marshal(deviceMsg)
	if err != nil {
		logs.Error(err)
	}
	payLoad := deviceMsgStr
	mq.Publish(topic, payLoad)
	paramsStr, err := json.Marshal(fields)
	if err == nil {
		service := map[string]interface{}{"identifier": service, "service_name": service, "params": string(paramsStr), "msg_id": deviceMsgId}
		influxdb.AddPointToService(deviceId, service)
	} else {
		fmt.Println(err)
	}
	return
}

func (mq *Client) GetProperty(productKey, deviceName string, args string) {
	//TODO:获取设备属性逻辑
	// 如果设备在线,获取设备真实属性，如果设备不在线,返回时序数据库数据
}