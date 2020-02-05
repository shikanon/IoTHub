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
	deviceId := "1XKCqOEZn90hsUtWFafPWhfqu1A"
	property := "LightAdjustLevel"
	value := 88
	params := map[string]interface{}{property: value}
	paramsStr, _ := json.Marshal(params)
	service := map[string]interface{}{"identifier": "set", "service_name": "set", "params": string(paramsStr), "msg_id": 1}
	AddPointToService(deviceId, service)
	serviceInfo := GetDeviceServiceInfoFromService(deviceId)
	if serviceInfo == nil { //try a unit test on function
		t.Error("时许数据库写入设备服务测试不通过") // 如果不是如预期的那么就报错
	} else {
		t.Log("时许数据库写入设备服务测试通过:") //记录一些你期望记录的信息
	}
}

func Test_AddEvent(t *testing.T) {
	deviceId := "1XKCqOEZn90hsUtWFafPWhfqu1A"
	property := "alarmType"
	value := 1
	params := map[string]interface{}{property: value}
	paramsStr, _ := json.Marshal(params)
	service := map[string]interface{}{
		"identifier": "alarmEvent",
		"event_name": "报警事件",
		"event_type": "alert",
		"outputData": string(paramsStr),
		"msg_id":     1,
	}
	AddPointToEvent(deviceId, service)
	end := time.Now().Unix()
	m, _ := time.ParseDuration("-10m")
	start := time.Now().Add(m).Unix()
	eventInfo := GetDeviceEventInfoFromEvent(deviceId, "alert", start, end, 1)
	if eventInfo == nil { //try a unit test on function
		t.Error("时许数据库写入设备事件测试不通过", eventInfo) // 如果不是如预期的那么就报错
	} else {
		t.Log("时许数据库写入设备事件测试通过:", eventInfo) //记录一些你期望记录的信息
	}
}
