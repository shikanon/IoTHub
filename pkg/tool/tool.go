package tool

import (
	"encoding/json"
	"fmt"
	"github.com/oklog/ulid"
	"github.com/rs/xid"
	"github.com/satori/go.uuid"
	"github.com/segmentio/ksuid"
	"github.com/shikanon/IoTOrbHub/config"
	"math/rand"
	"regexp"
	"time"
)

// 生成productKey (ulid)
func GenerateProductKey() (secret string) {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return id.String()
}

// 生成productSecret (xid)
func GenerateProductSecret() (secret string) {
	return xid.New().String()
}

// 生成deviceSecret  (uuid)
func GenerateDeviceSecret() (secret string) {
	data := fmt.Sprintf("%s", uuid.NewV4())
	filter, _ := regexp.Compile("-")
	result := filter.ReplaceAllString(data, "")
	return result
}

// 生成设备唯一id  (ksuid)
func GenerateIotId() (IotId string) {
	return ksuid.New().String()
}

// TODO 生成随机唯一的产品名称-批量生成

// 处理数据中的json字符串
func DealJsonStr(arg []map[string]interface{})(result []map[string]interface{}){
	for index, value := range arg{
		var dat map[string]interface{}
		var json_str string
		json_str = value["Value"].(string)
		json.Unmarshal([]byte(json_str), &dat)
		arg[index]["Value"] = dat
	}
	return arg
}

// 处理转化时间
func TimeDeal(time time.Time) (result string) {
	time_str := time.Format(config.GeneralConfig.TimeFormat)
	//if (time == ){}else {}
	if time_str == "0001/01/01 00:00:00" {
		return "-"
	} else {
		return time_str
	}
}