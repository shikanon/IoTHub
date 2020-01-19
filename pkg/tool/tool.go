package tool

import (
	"fmt"
	"github.com/oklog/ulid"
	"github.com/rs/xid"
	"github.com/satori/go.uuid"
	"github.com/segmentio/ksuid"
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
