package util

import (
	"encoding/json"
	"github.com/shikanon/IoTOrbHub/pkg/database"
	"github.com/shikanon/IoTOrbHub/pkg/influxdb"
	"strconv"

	//"crypto/tls"
	"errors"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/shikanon/IoTOrbHub/pkg/constants"
	logs "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"os"
	"strings"
	"time"
)

var (
	// TokenWaitTime to wait
	TokenWaitTime = 120 * time.Second
)

type parseStruct struct {
}

// CheckKeyExist check dis info format
func CheckKeyExist(keys []string, disinfo map[string]interface{}) error {
	for _, v := range keys {
		_, ok := disinfo[v]
		if !ok {
			logs.Info("key: %s not found", v)
			return errors.New("key not found")
		}
	}
	return nil
}

// CheckClientToken checks token is right
func CheckClientToken(token MQTT.Token) (bool, error) {
	if token.Wait() && token.Error() != nil {
		return false, token.Error()
	}
	return true, nil
}

// PathExist check file exists or not
func PathExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// HubClientInit create mqtt client config
func HubClientInit(server, clientID, username, password string) *MQTT.ClientOptions {
	opts := MQTT.NewClientOptions().AddBroker(server).SetClientID(clientID).SetCleanSession(true)
	if username != "" {
		opts.SetUsername(username)
		if password != "" {
			opts.SetPassword(password)
		}
	}
	//tlsConfig := &tls.Config{InsecureSkipVerify: true, ClientAuth: tls.NoClientCert}
	//opts.SetTLSConfig(tlsConfig)
	return opts
}

// LoopConnect connect to mqtt server
func LoopConnect(clientID string, client MQTT.Client) {
	for {
		logs.Info("start connect to mqtt server with client id: %s", clientID)
		token := client.Connect()
		logs.Info("client %s isconnected: %v", clientID, client.IsConnected())
		if rs, err := CheckClientToken(token); !rs {
			logs.Errorf("connect error: %v", err)
		} else {
			return
		}
		time.Sleep(5 * time.Second)
	}
}

func GetDeviceDetailFromTopic(topic string) (string, string, error) {
	res := strings.Split(topic, "/")
	if len(res) >= constants.DeviceNameIndex && res[constants.DeviceTopicPrefixIndex] == constants.DeviceTopicPrefix {
		return res[constants.ProductKeyIndex], res[constants.DeviceNameIndex], nil
	}
	return "", "", errors.New("failed to get product key and device name from topic")
}

func GetDeviceDetailFromShadowTopic(topic string) (string, string, error) {
	res := strings.Split(topic, "/")
	if len(res) >= constants.ShadowDeviceNameIndex && res[constants.ShadowTopicPrefixIndex] == constants.ShadowTopicPrefix {
		return res[constants.ShadowProductKeyIndex], res[constants.ShadowDeviceNameIndex], nil
	}
	return "", "", errors.New("failed to get product key and device name from shadow topic")
}

func GetDeviceDetailFromClientID(clientID string) (string, string, error) {
	res := strings.Split(strings.Split(clientID, "|")[0], "&")
	if len(res) == 2 {
		return res[0], res[1], nil
	}
	return "", "", errors.New("failed to get product key and device name from clientID")
}

func GetEventlFromTopic(topic string) (string, error) {
	res := strings.Split(topic, "/")
	if len(res) >= constants.EventIndex && res[constants.DeviceTopicPrefixIndex] == constants.DeviceTopicPrefix {
		return res[constants.EventIndex], nil
	}
	return "", errors.New("failed to get event from topic")
}

func GetServicelFromTopic(topic string) (string, error) {
	res := strings.Split(topic, "/")
	if len(res) >= constants.ServiceIndex && res[constants.DeviceTopicPrefixIndex] == constants.DeviceTopicPrefix {
		servicesArr := strings.Split(res[constants.ServiceIndex], "_")
		if len(servicesArr) == 2 {
			return servicesArr[0], nil
		}
	}
	return "", errors.New("failed to get event from topic")
}

func DataVerification(key, dataType string, value, result gjson.Result) (interface{}, error) {
	switch dataType {
	case "int":
		return CheckInt(key, value, result)
	case "float":
		return CheckFloat(key, value, result)
	case "double":
		return CheckDouble(key, value, result)
	case "enum":
		return CheckEnum(key, value, result)
	case "bool":
		return CheckBool(key, value)
	case "text":
		return CheckText(key, value, result)
	case "date":
		return CheckDate(key, value)
	case "struct":
		return CheckStruct(key, value, result)
	case "array":
		return CheckArray(key, value, result)
	default:
		return nil, errors.New(fmt.Sprintf("6328:tsl parse: date type is not exit -> %s", key))
	}
}

func CheckInt(key string, value, result gjson.Result) (interface{}, error) {
	var x int32
	fmt.Println(value.Raw)
	err := json.Unmarshal([]byte(value.Raw), &x)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New(fmt.Sprintf("6307:tsl parse: data type is not int -> %s", key))
	}
	if min := result.Get("dataType.specs.min"); min.Exists() && value.Int() < min.Int() {
		return nil, errors.New(fmt.Sprintf("6306:tsl parse: int value is smaller than min %s -> %s", min.String(), key))
	}
	if max := result.Get("dataType.specs.max"); max.Exists() && value.Int() > max.Int() {
		return nil, errors.New(fmt.Sprintf("6306:tsl parse: int value is bigger than max %s -> %s", max.String(), key))
	}
	return x, nil
}

func CheckFloat(key string, value, result gjson.Result) (interface{}, error) {
	var x float32
	fmt.Println(value.Raw)
	err := json.Unmarshal([]byte(value.Raw), &x)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New(fmt.Sprintf("6307:tsl parse: data type is not float -> %s", key))
	}
	if min := result.Get("dataType.specs.min"); min.Exists() && value.Float() < min.Float() {
		return nil, errors.New(fmt.Sprintf("6306:tsl parse: float value is smaller than min %s -> %s", min.String(), key))
	}
	if max := result.Get("dataType.specs.max"); max.Exists() && value.Float() > max.Float() {
		return nil, errors.New(fmt.Sprintf("6306:tsl parse: float value is bigger than max %s -> %s", max.String(), key))
	}
	return x, nil
}

func CheckDouble(key string, value, result gjson.Result) (interface{}, error) {
	var x float64
	fmt.Println(value.Raw)
	err := json.Unmarshal([]byte(value.Raw), &x)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New(fmt.Sprintf("6307:tsl parse: data type is not double -> %s", key))
	}
	if min := result.Get("dataType.specs.min"); min.Exists() && value.Float() < min.Float() {
		return nil, errors.New(fmt.Sprintf("6306:tsl parse: double value is smaller than min %s -> %s", min.String(), key))
	}
	if max := result.Get("dataType.specs.max"); max.Exists() && value.Float() > max.Float() {
		return nil, errors.New(fmt.Sprintf("6306:tsl parse: double value is bigger than max %s -> %s", max.String(), key))
	}
	return x, nil
}

func CheckEnum(key string, value, result gjson.Result) (interface{}, error) {
	var x int
	err := json.Unmarshal([]byte(value.Raw), &x)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New(fmt.Sprintf("6309:tsl parse: value of enum type must be int -> %s", key))
	}
	enumArr := result.Get("dataType.specs")
	if !func(enumArr, value gjson.Result) bool {
		for k, _ := range enumArr.Map() {
			if k == value.String() {
				return true
			}
		}
		return false
	}(enumArr, value) {
		return nil, errors.New(fmt.Sprintf("6309:tsl parse: enum specs error -> %s", key))
	}
	return x, nil
}

func CheckBool(key string, value gjson.Result) (interface{}, error) {
	var x int
	err := json.Unmarshal([]byte(value.Raw), &x)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New(fmt.Sprintf("6308:tsl parse: value of bool type must be int -> %s", key))
	} else if x != 0 && x != 1 {
		return nil, errors.New(fmt.Sprintf("6308:tsl parse: bool specs error -> %s", key))
	}
	return x, nil
}

func CheckText(key string, value, result gjson.Result) (interface{}, error) {
	var x string
	fmt.Println(value.Raw)
	err := json.Unmarshal([]byte(value.Raw), &x)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New(fmt.Sprintf("6310:tsl parse: data type is not string -> %s", key))
	}
	if length := result.Get("dataType.specs.length"); length.Exists() && int64(len(value.String())) > length.Int() {
		return nil, errors.New(fmt.Sprintf("6310:tsl parse: string length bigger than max length %s -> %s", length.String(), key))
	}
	return x, nil
}

func CheckDate(key string, value gjson.Result) (interface{}, error) {
	var str string // 单位：毫秒
	fmt.Println(value.Raw)
	err1 := json.Unmarshal([]byte(value.Raw), &str)
	x,err2 := strconv.Atoi(str)
	if err1 != nil || err2 != nil {
		fmt.Println(err1,err2)
		return nil, errors.New(fmt.Sprintf("6311:tsl parse: date type must be a string of long(UTC ms) -> %s", key))
	}

	timeTemplate := "2006-01-02 15:04:05" //常规类型
	date := time.Unix(int64(x/1e3), 0).Format(timeTemplate)
	return date, nil
}

func CheckStruct(key string, value, result gjson.Result) (interface{}, error) {
	var x map[string]interface{}
	fmt.Println(value.Raw)
	err := json.Unmarshal([]byte(value.Raw), &x)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New(fmt.Sprintf("6312:tsl parse: struct type must be json -> %s", key))
	}
	paramsNum := len(result.Get("dataType.specs").Array())
	if len(x) != paramsNum {
		return nil, errors.New(fmt.Sprintf("6312:tsl parse: struct param size error -> %s", key))
	}
	for k, _ := range x {
		subresult := result.Get(fmt.Sprintf("dataType.specs.#(identifier==%s)", k))
		if !subresult.Exists() {
			return nil, errors.New(fmt.Sprintf("6304:tsl parse: field %s not exist in struct %s", k, key))
		}
		subdataType := subresult.Get("dataType.type")
		v := value.Get(k)
		switch subdataType.String() {
		case "int":
			v, err := CheckInt(k, v, subresult)
			if err != nil {
				return nil, err
			} else {
				x[k] = v
			}
		case "float":
			v, err := CheckFloat(k, v, subresult)
			if err != nil {
				return nil, err
			} else {
				x[k] = v
			}
		case "double":
			v, err := CheckDouble(k, v, subresult)
			if err != nil {
				return nil, err
			} else {
				x[k] = v
			}
		case "enum":
			v, err := CheckEnum(k, v, subresult)
			if err != nil {
				return nil, err
			} else {
				x[k] = v
			}
		case "bool":
			v, err := CheckBool(k, v)
			if err != nil {
				return nil, err
			} else {
				x[k] = v
			}
		case "text":
			v, err := CheckText(k, v, subresult)
			if err != nil {
				return nil, err
			} else {
				x[k] = v
			}
		case "date":
			v, err := CheckDate(k, v)
			if err != nil {
				return nil, err
			} else {
				x[k] = v
			}
		default:
			return nil, errors.New(fmt.Sprintf("6328:tsl parse: date type is not exit -> %s", key))
		}
	}
	str,err := json.Marshal(x)
	if err != nil {
		logs.Error(err)
	}
	return string(str), nil
}

func CheckArray(key string, value, result gjson.Result) (interface{}, error) {
	var x []interface{}
	fmt.Println(value.Raw)
	err := json.Unmarshal([]byte(value.Raw), &x)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New(fmt.Sprintf("6328:tsl parse: date type is not array -> %s", key))
	}
	maxSize := result.Get("dataType.specs.size")
	if int64(len(x)) > maxSize.Int() {
		return nil, errors.New(fmt.Sprintf("6324:tsl parse: array size is bigger than max size %s -> %s", maxSize, key))
	}
	arrayType := result.Get("dataType.specs.item.type")
	var array []interface{}
	switch arrayType.String() {
	case "int":
		for _, v := range value.Array() {
			v, err := CheckInt(key, v, result)
			if err != nil {
				return nil, err
			} else {
				array = append(array, v)
			}
		}
	case "float":
		for _, v := range value.Array() {
			v, err := CheckFloat(key, v, result)
			if err != nil {
				return nil, err
			} else {
				array = append(array, v)
			}
		}
	case "double":
		for _, v := range value.Array() {
			v, err := CheckDouble(key, v, result)
			if err != nil {
				return nil, err
			} else {
				array = append(array, v)
			}
		}
	case "text":
		for _, v := range value.Array() {
			v, err := CheckText(key, v, result)
			if err != nil {
				return nil, err
			} else {
				array = append(array, v)
			}
		}
	case "struct":
		for _, v := range value.Array() {
			v, err := CheckStruct(key, v, result)
			if err != nil {
				return nil, err
			} else {
				array = append(array, v)
			}
		}
	default:
		return nil, errors.New(fmt.Sprintf("6328:tsl parse: date type is not exit -> %s", key))
	}
	return array, nil
}

func GetDevicePropertyStatusInfo(productKey,deviceId string) (properties []map[string]interface{}) {
	model := database.GetIntactModel(productKey)
	modelJson,err := json.Marshal(model)
	if err != nil {
		fmt.Println(err)
	}
	result:= gjson.Get(string(modelJson), "properties")
	for _,v := range result.Array() {
		property := make(map[string]interface{})
		property["Identifier"] = v.Map()["identifier"].String()
		property["Name"] = v.Map()["name"].String()
		property["Unit"] = v.Map()["dataType"].Map()["specs"].Map()["unit"].String()
		property["DataType"] = v.Map()["dataType"].Map()["type"].String()
		ts, value, err := influxdb.GetDevicePropertyFromPropertyReported(influxdb.InfluxClient, deviceId, property["Identifier"].(string))
		if err == nil {
			t, err := time.Parse(time.RFC3339, ts.(string))
			if err != nil {
				fmt.Println(err)
			}
			property["Time"] = t.Format("2006/01/02 15:04:05")
			property["Value"] = value
		}
		properties = append(properties, property)
	}
	return properties
}

func GetDeviceDesiredPropertyInfo(productKey,deviceId string) (properties []map[string]interface{}) {
	model := database.GetIntactModel(productKey)
	modelJson,err := json.Marshal(model)
	if err != nil {
		fmt.Println(err)
	}
	result:= gjson.Get(string(modelJson), "properties")
	for _,v := range result.Array() {
		property := make(map[string]interface{})
		property["Identifier"] = v.Map()["identifier"].String()
		property["Name"] = v.Map()["name"].String()
		property["Unit"] = v.Map()["dataType"].Map()["specs"].Map()["unit"].String()
		property["DataType"] = v.Map()["dataType"].Map()["type"].String()
		property["Version"] = 0
		ts, value, version, err := influxdb.GetDevicePropertyFromPropertyDesired(influxdb.InfluxClient, deviceId, property["Identifier"].(string))
		if err == nil {
			t, err := time.Parse(time.RFC3339, ts.(string))
			if err != nil {
				fmt.Println(err)
			}
			property["Time"] = t.Format("2006/01/02 15:04:05")
			property["Value"] = value
			property["Version"] = version
		}
		properties = append(properties, property)
	}
	return properties
}


func GetDeviceServiceInfo(deviceId string) (serviceInfo []map[string]interface{}) {
	res := influxdb.GetDeviceServiceInfoFromService(influxdb.InfluxClient, deviceId)
	if res != nil {
		for _,v := range res {
			t, err := time.Parse(time.RFC3339, v[0].(string))
			if err != nil {
				fmt.Println(err)
			}
			service := map[string]interface{}{
				"Time": t.Format("2006/01/02 15:04:05"),
				"Identifier": v[2],
				"Name": v[5],
				"InputData": v[4],
				"OutputData":fmt.Sprintf(`{"code":200,"data":{},"id":"%s","message":"success","version":"1.0"}`, v[3]),
			}
			serviceInfo = append(serviceInfo, service)
		}
	}
	return
}

func GetDeviceEventInfo(deviceId string) (eventInfo []map[string]interface{}) {
	res := influxdb.GetDeviceEventInfoFromEvent(influxdb.InfluxClient, deviceId)
	if res != nil {
		for _,v := range res {
			t, err := time.Parse(time.RFC3339, v[0].(string))
			if err != nil {
				fmt.Println(err)
			}
			event := map[string]interface{}{
				"Time": t.Format("2006/01/02 15:04:05"),
				"Identifier": v[4],
				"Name": v[2],
				"EventType": v[3],
				"OutputData":v[6],
			}
			eventInfo = append(eventInfo, event)
		}
	}
	return
}