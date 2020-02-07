package influxdb

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func Test_AddPropertyReported(t *testing.T) {
	// 初始化配置
	c := NewInfluxdbClient()
	InfluxClient = c
	deviceId := "1XKCqOEZn90hsUtWFafPWhfqu1A"
	property := "LightAdjustLevel"
	value := int64(77)
	field := map[string]interface{}{property: value}
	AddPointToPropertyReported(deviceId, field)
	_, v, err := GetDevicePropertyFromPropertyReported(deviceId, property)
	if v, _ := v.(json.Number).Int64(); err != nil || v != value { //try a unit test on function
		t.Error(fmt.Sprintf("时许数据库写入设备上报属性测试不通过:%s,", err)) // 如果不是如预期的那么就报错
	} else {
		t.Log("时许数据库写入设备上报属性测试通过") //记录一些你期望记录的信息
	}
	defer InfluxClient.Close()
}

func Test_AddPropertyDesired(t *testing.T) {
	// 初始化配置
	c := NewInfluxdbClient()
	InfluxClient = c
	deviceId := "1XKCqOEZn90hsUtWFafPWhfqu1A"
	property := "LightAdjustLevel"
	value := int64(88)
	param := map[string]interface{}{property: value, "version": 1, "msg_id": 1}
	AddPointToPropertyDesired(deviceId, param)
	_, v, _, err := GetDevicePropertyFromPropertyDesired(deviceId, property)
	if v, _ := v.(json.Number).Int64(); err != nil || v != value { //try a unit test on function
		t.Error(fmt.Sprintf("时许数据库写入设备期望属性测试不通过:%s,", err)) // 如果不是如预期的那么就报错
	} else {
		t.Log("时许数据库写入设备期望属性测试通过") //记录一些你期望记录的信息
	}
	defer InfluxClient.Close()
}

func Test_AddService(t *testing.T) {
	c := NewInfluxdbClient()
	InfluxClient = c
	deviceId := "1XKCqOEZn90hsUtWFafPWhfqu1A"
	property1 := "LightAdjustLevel"
	value1 := 66
	params1 := map[string]interface{}{property1: value1}
	paramsStr1, _ := json.Marshal(params1)
	service1 := map[string]interface{}{
		"identifier": "set",
		"service_name": "set",
		"params": string(paramsStr1),
		"msg_id": 2,
	}
	AddPointToService(deviceId, service1)
	time.Sleep(1*time.Second)
	property2 := "TimeReset"
	value2 := "test"
	params2 := map[string]interface{}{property2: value2}
	paramsStr2, _ := json.Marshal(params2)
	service2 := map[string]interface{}{
		"identifier": "TimeReset",
		"service_name": "TimeReset",
		"params": string(paramsStr2),
		"msg_id": 2,
	}
	AddPointToService(deviceId, service2)
	end := time.Now().Unix()
	m, _ := time.ParseDuration("-10m")
	start := time.Now().Add(m).Unix()
	serviceInfo := GetDeviceServiceInfoFromService(deviceId, "all", start, end, false)
	if serviceInfo == nil { //try a unit test on function
		t.Error("时许数据库写入设备服务测试不通过") // 如果不是如预期的那么就报错
	} else {
		t.Log("时许数据库写入设备服务测试通过", len(serviceInfo), serviceInfo) //记录一些你期望记录的信息
	}
	defer InfluxClient.Close()
}

func Test_AddEvent(t *testing.T) {
	c := NewInfluxdbClient()
	InfluxClient = c
	deviceId := "1XKCqOEZn90hsUtWFafPWhfqu1A"
	property1 := "ErrorCode"
	value1 := 0
	params1 := map[string]interface{}{property1: value1}
	paramsStr1, _ := json.Marshal(params1)
	event1 := map[string]interface{}{
		"identifier": "Error",
		"event_name": "故障上报",
		"event_type": "info",
		"outputData": string(paramsStr1),
		"msg_id":     1,
	}
	AddPointToEvent(deviceId, event1)
	time.Sleep(1*time.Second)
	property2 := "alarmType"
	value2 := 1
	params2 := map[string]interface{}{property2: value2}
	paramsStr2, _ := json.Marshal(params2)
	event2 := map[string]interface{}{
		"identifier": "alarmEvent",
		"event_name": "报警事件",
		"event_type": "alert",
		"outputData": string(paramsStr2),
		"msg_id":     1,
	}
	AddPointToEvent(deviceId, event2)
	time.Sleep(1*time.Second)
	property3 := "Error"
	value3 := 1
	params3 := map[string]interface{}{property3: value3}
	paramsStr3, _ := json.Marshal(params3)
	event3 := map[string]interface{}{
		"identifier": "CommunicationFault",
		"event_name": "通讯故障",
		"event_type": "error",
		"outputData": string(paramsStr3),
		"msg_id":     1,
	}
	AddPointToEvent(deviceId, event3)
	end := time.Now().Unix()
	m, _ := time.ParseDuration("-10m")
	start := time.Now().Add(m).Unix()
	eventInfo := GetDeviceEventInfoFromEvent(deviceId, "all", "all", start, end, false)
	if eventInfo == nil { //try a unit test on function
		t.Error("时许数据库写入设备事件测试不通过") // 如果不是如预期的那么就报错
	} else {
		t.Log("时许数据库写入设备事件测试通过", len(eventInfo), eventInfo) //记录一些你期望记录的信息
	}
	defer InfluxClient.Close()
}

func TestGetDevicePropertyHistoryFromPropertyReported(t *testing.T) {
	c := NewInfluxdbClient()
	InfluxClient = c
	deviceId := "1XKCqOEZn90hsUtWFafPWhfqu1A"
	property := "LightAdjustLevel"
	end := time.Now().Unix()
	m, _ := time.ParseDuration("-10m")
	start := time.Now().Add(m).Unix()
	propertyHistory := GetDevicePropertyHistoryFromPropertyReported(deviceId, property, start, end, false)
	if propertyHistory == nil { //try a unit test on function
		t.Error("时许数据库获取设备属性历史数据测试不通过") // 如果不是如预期的那么就报错
	} else {
		t.Log("时许数据库获取设备属性历史数据测试通过", len(propertyHistory), propertyHistory) //记录一些你期望记录的信息
	}
	defer InfluxClient.Close()
}

func TestGetNextDeviceServiceInfoFromService(t *testing.T) {
	c := NewInfluxdbClient()
	InfluxClient = c
	deviceId := "1XKCqOEZn90hsUtWFafPWhfqu1A"
	end := time.Now().Unix()
	m, _ := time.ParseDuration("-10m")
	start := time.Now().Add(m).Unix()
	serviceInfo := GetDeviceServiceInfoFromService(deviceId, "set", start, end, true)
	t.Log("时许数据库服务数据getNext测试通过")
	if len(serviceInfo) != 0 {
		for _,v := range serviceInfo {
			t, err := time.Parse(time.RFC3339, v[0].(string))
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(t.Unix())
		}
	}
	defer InfluxClient.Close()
}

func TestGetNextDeviceEventInfoFromEvent(t *testing.T) {
	c := NewInfluxdbClient()
	InfluxClient = c
	deviceId := "1XKCqOEZn90hsUtWFafPWhfqu1A"
	end := time.Now().Unix()
	m, _ := time.ParseDuration("-10m")
	start := time.Now().Add(m).Unix()
	eventInfo := GetDeviceEventInfoFromEvent(deviceId, "all", "all", start, end, true)
	t.Log("时许数据库设备事件数据getNext测试通过")
	if len(eventInfo) != 0 {
		for _,v := range eventInfo {
			t, err := time.Parse(time.RFC3339, v[0].(string))
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(t.Unix())
		}
	}
	defer InfluxClient.Close()
}

func TestGetNextDevicePropertyHistoryFromPropertyReported(t *testing.T) {
	c := NewInfluxdbClient()
	InfluxClient = c
	deviceId := "1XKCqOEZn90hsUtWFafPWhfqu1A"
	property := "LightAdjustLevel"
	end := time.Now().Unix()
	m, _ := time.ParseDuration("-10m")
	start := time.Now().Add(m).Unix()
	propertyHistory := GetDevicePropertyHistoryFromPropertyReported(deviceId, property, start, end, true)
	t.Log("时许数据库设备属性历史数据getNext测试通过") //记录一些你期望记录的信息
	if len(propertyHistory) != 0 {
		for _,v := range propertyHistory {
			t, err := time.Parse(time.RFC3339, v[0].(string))
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(t.Unix())
		}
	}
	defer InfluxClient.Close()
}
