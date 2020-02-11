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

