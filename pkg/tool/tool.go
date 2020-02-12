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
	"strconv"
	"time"
)

func GenerateProductKey() (secret string) {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return id.String()
}

func GenerateProductSecret() (secret string) {
	return xid.New().String()
}

func GenerateDeviceSecret() (secret string) {
	data := fmt.Sprintf("%s", uuid.NewV4())
	filter, _ := regexp.Compile("-")
	result := filter.ReplaceAllString(data, "")
	return result
}

func GenerateIotId() (IotId string) {
	return ksuid.New().String()
}

func GenerateAutoDeviceName() (name string) {
	decret := GenerateDeviceSecret()
	iotid := GenerateIotId()

	return decret[0:10] + iotid[0:10]
}

func StringNumberToInTNumber(num_str string) (num_int int) {
	result, _ := strconv.Atoi(num_str)
	return result
}

func DealSequentialDatabaseData(arg []map[string]interface{}) (result []map[string]interface{}) {
	for index, value := range arg {
		var dat map[string]interface{}
		var json_str string
		json_str = value["Value"].(string)
		if json_str != "" {
			json.Unmarshal([]byte(json_str), &dat)
			arg[index]["Value"] = dat
		}
	}
	return arg
}

func TimeDeal(time time.Time) (result string) {
	time_str := time.Format(config.GeneralConfig.TimeFormat)
	//if (time == ){}else {}
	if time_str == "0001/01/01 00:00:00" {
		return "-"
	} else {
		return time_str
	}
}

func JsonStrToMap(data_str string) (data_map map[string]string) {
	map_data := make(map[string]string)
	json.Unmarshal([]byte(data_str), &map_data)
	return map_data
}

func MapToJsonStr(map_data map[string]string) (data_str string) {
	data, _ := json.Marshal(map_data)
	result := string(data)
	return result
}

func GetStringSpecialCharCount(data, char string) (number int) {
	re, _ := regexp.Compile(char)
	b := re.FindAllString(data, -1)
	return len(b)
}

func GetStringChinaCharCount(data string) (number int) {
	re, _ := regexp.Compile("[\u4e00-\u9fa5]")
	b := re.FindAllString(data, -1)
	return len(b)
}

func GetStringEnglishCharAndNumberCount(data string) (number int) {
	re, _ := regexp.Compile("[0-9a-zA-Z]")
	b := re.FindAllString(data, -1)
	return len(b)
}

func GetPointInStringCount(data string) (number int) {
	re, _ := regexp.Compile(`\.`)
	b := re.FindAll([]byte(data), -1)
	return len(b)
}

func GetParenthesesInStringCount(data string) (number int) {
	re, _ := regexp.Compile(`\(`)
	b := re.FindAll([]byte(data), -1)
	return len(b)
}

func GetTheParenthesesInStringCount(data string) (number int) {
	re, _ := regexp.Compile(`\)`)
	b := re.FindAll([]byte(data), -1)
	return len(b)
}
