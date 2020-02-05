package influxdb

import (
	"encoding/json"
	"errors"
	"fmt"
	client "github.com/influxdata/influxdb1-client/v2"
	"log"
	"strconv"
	"time"
)

const (
	addr            = "http://139.199.89.93:8086"
	MyDB            = "orbbeciot"
	username        = "orbbec"
	password        = "orbbec"
	measurementName = "property_reported"
	query           = "device_id"
)

var (
	InfluxClient client.Client
)

func main() {
	// Create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     addr,
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	//field := map[string]interface{}{"a":1,"b":1}
	//msgId := GetDeviceMsgIdFromPropertyDesired(c, "1")
	//AddPointToPropertyDesired(c, "1", msgId, field)
	//field := map[string]interface{}{"a":1}
	//msgId := GetDeviceMsgIdFromPropertyDesired(c, "1")
	//AddPointToPropertyDesired(c, "1", msgId, field)
	//field2 := map[string]interface{}{"isReply":"true"}
	//AddPointToPropertyDesired(c, "1", 4, field2)
	//now := time.Now().Unix()
	currentTime := time.Now()
	m, _ := time.ParseDuration("-1h")
	result := strconv.FormatInt(currentTime.Add(m).UnixNano(), 10)
	fmt.Println(result)
	q := client.NewQuery(fmt.Sprintf("SELECT * FROM %s where time > %s order by time desc tz('Asia/Shanghai')", "test", result), MyDB, "")
	if response, err := c.Query(q); err == nil && response.Error() == nil && len(response.Results[0].Series) != 0 {
		fmt.Println(response.Results[0].Series[0].Values)
	}
	q = client.NewQuery(fmt.Sprintf("SELECT * FROM %s where time > now() - 1h order by time desc tz('Asia/Shanghai')", "test"), MyDB, "")
	if response, err := c.Query(q); err == nil && response.Error() == nil && len(response.Results[0].Series) != 0 {
		fmt.Println(response.Results[0].Series[0].Values)
	}
}

func NewInfluxdbClient() client.Client {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     addr,
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func AddPointToPropertyReported(deviceId string, fields map[string]interface{}) {
	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}
	// Create a point and add to batch
	tags := map[string]string{"device_id": deviceId}
	pt, err := client.NewPoint("property_reported", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	// Write the batch
	if err := InfluxClient.Write(bp); err != nil {
		log.Fatal(err)
	}
}

func AddPointToPropertyDesired(deviceId string, fields map[string]interface{}) {
	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}
	// Create a point and add to batch
	tags := map[string]string{"device_id": deviceId}
	pt, err := client.NewPoint("property_desired", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	// Write the batch
	if err := InfluxClient.Write(bp); err != nil {
		log.Fatal(err)
	}
}

func AddPointToService(deviceId string, fields map[string]interface{}) {
	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}
	// Create a point and add to batch
	tags := map[string]string{"device_id": deviceId}
	pt, err := client.NewPoint("service", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	// Write the batch
	if err := InfluxClient.Write(bp); err != nil {
		log.Fatal(err)
	}
}

func AddPointToEvent(deviceId string, fields map[string]interface{}) {
	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}
	// Create a point and add to batch
	tags := map[string]string{"device_id": deviceId}
	pt, err := client.NewPoint("event", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	// Write the batch
	if err := InfluxClient.Write(bp); err != nil {
		log.Fatal(err)
	}
}

func GetDeviceMsgIdFromService() (msgId int64) {
	q := client.NewQuery(fmt.Sprintf("SELECT last(msg_id) FROM %s", "service"), MyDB, "")
	if response, err := InfluxClient.Query(q); err == nil && response.Error() == nil && len(response.Results[0].Series) != 0 {
		msgId, err = response.Results[0].Series[0].Values[0][1].(json.Number).Int64()
		if err != nil {
			fmt.Println(err)
			return 1
		} else {
			msgId += 1
			return
		}
	} else {
		return 1
	}
}

func GetPropertyVersionFromPropertyDesired(deviceId, property string) (version int64) {
	q := client.NewQuery(fmt.Sprintf("SELECT last(%s) FROM %s where %s='%s'", property, "property_desired", "device_id", deviceId), MyDB, "")
	if response, err := InfluxClient.Query(q); err == nil && response.Error() == nil && len(response.Results[0].Series) != 0 {
		ts := response.Results[0].Series[0].Values[0][0]
		q = client.NewQuery(fmt.Sprintf("SELECT last(version) FROM %s where %s='%s' and %s='%s'", "property_desired", "device_id", deviceId, "time", ts), MyDB, "")
		if response, err := InfluxClient.Query(q); err == nil && response.Error() == nil && len(response.Results[0].Series) != 0 {
			version, err = response.Results[0].Series[0].Values[0][1].(json.Number).Int64()
			if err != nil {
				fmt.Println(err)
				return 1
			} else {
				version += 1
				return
			}
		} else {
			return 1
		}
	} else {
		return 1
	}
}

func GetDevicePropertyFromPropertyReported(deviceId, property string) (ts, value interface{}, err error) {
	q := client.NewQuery(fmt.Sprintf("SELECT last(%s) FROM %s where %s='%s' tz('Asia/Shanghai')", property, "property_reported", "device_id", deviceId), MyDB, "")
	if response, err := InfluxClient.Query(q); err == nil && response.Error() == nil && len(response.Results[0].Series) != 0 {
		ts = response.Results[0].Series[0].Values[0][0]
		value = response.Results[0].Series[0].Values[0][1]
		return ts, value, nil
	} else {
		err = errors.New(fmt.Sprintf("Property %s does not exist or is nil", property))
		return nil, nil, err
	}
}

func GetDevicePropertyFromPropertyDesired(deviceId, property string) (ts, value, version interface{}, err error) {
	q := client.NewQuery(fmt.Sprintf("SELECT last(%s) FROM %s where %s='%s'", property, "property_desired", "device_id", deviceId), MyDB, "")
	if response, err := InfluxClient.Query(q); err == nil && response.Error() == nil && len(response.Results[0].Series) != 0 {
		t := response.Results[0].Series[0].Values[0][0]
		value = response.Results[0].Series[0].Values[0][1]
		q = client.NewQuery(fmt.Sprintf("SELECT last(version) FROM %s where %s='%s' and %s='%s' tz('Asia/Shanghai')", "property_desired", "device_id", deviceId, "time", t), MyDB, "")
		if response, err := InfluxClient.Query(q); err == nil && response.Error() == nil && len(response.Results[0].Series) != 0 {
			version = response.Results[0].Series[0].Values[0][1]
			ts = response.Results[0].Series[0].Values[0][0]
			return ts, value, version, nil
		} else {
			fmt.Print(err, "*", response.Error())
			err = errors.New(fmt.Sprintf("Property %s does not exist or is nil", property))
			return nil, nil, nil, err
		}
	} else {
		err = errors.New(fmt.Sprintf("Property %s does not exist or is nil", property))
		return nil, nil, nil, err
	}
}

func GetDeviceServiceInfoFromService(deviceId string) (serviceInfo [][]interface{}) {
	q := client.NewQuery(fmt.Sprintf("SELECT * FROM %s where %s='%s' and time > now() - 1h order by time desc tz('Asia/Shanghai')", "service", "device_id", deviceId), MyDB, "")
	if response, err := InfluxClient.Query(q); err == nil && response.Error() == nil && len(response.Results[0].Series) != 0 {
		return response.Results[0].Series[0].Values
	} else {
		return nil
	}
}

func GetDeviceEventInfoFromEvent(deviceId, event_type string, start, end int64, page int) (eventInfo [][]interface{}) {
	t1 := time.Unix(start, 0).Format(time.RFC3339)
	t2 := time.Unix(end, 0).Format(time.RFC3339)
	offset := (page - 1) * 9
	if event_type == "all" {
		q := client.NewQuery(fmt.Sprintf("SELECT * FROM %s where %s='%s' and time >= '%s' and time <= '%s' order by time desc LIMIT 9 OFFSET %s tz('Asia/Shanghai')", "event", "device_id", deviceId, t1, t2, strconv.Itoa(offset)), MyDB, "")
		fmt.Print(q)
		if response, err := InfluxClient.Query(q); err == nil && response.Error() == nil && len(response.Results[0].Series) != 0 {
			return response.Results[0].Series[0].Values
		} else {
			return nil
		}
	} else {
		q := client.NewQuery(fmt.Sprintf("SELECT * FROM %s where %s='%s' and %s='%s' and time >= '%s' and time <= '%s' order by time desc LIMIT 9 OFFSET %s tz('Asia/Shanghai')", "event", "device_id", deviceId, "event_type", event_type, t1, t2, strconv.Itoa(offset)), MyDB, "")
		fmt.Println(q)
		if response, err := InfluxClient.Query(q); err == nil && response.Error() == nil && len(response.Results[0].Series) != 0 {
			return response.Results[0].Series[0].Values
		} else {
			return nil
		}
	}

}

func GetDevicePropertyHistoryFromPropertyReported(deviceId, property string, Hour int) (propertyHistory [][]interface{}) {
	q := client.NewQuery(fmt.Sprintf("SELECT %s FROM %s where %s='%s' and time > now() - %sh order by time desc tz('Asia/Shanghai')", property, "property_reported", "device_id", deviceId, strconv.Itoa(Hour)), MyDB, "")
	if response, err := InfluxClient.Query(q); err == nil && response.Error() == nil && len(response.Results[0].Series) != 0 {
		return response.Results[0].Series[0].Values
	} else {
		return nil
	}
}
