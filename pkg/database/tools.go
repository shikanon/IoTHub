package database

import (
	"encoding/json"
	"fmt"
	"github.com/shikanon/IoTOrbHub/config"
	"github.com/shikanon/IoTOrbHub/pkg/tool"
	"github.com/tidwall/gjson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type TopicModel struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Desc           string `json:"desc"`
	TopicShortName string `json:"topic_name"`
	Operation      int    `json:"operation"`
}

type Topics struct {
	Base    []TopicModel `json:"base"`
	Physics []TopicModel `json:"physics"`
	Custom  []TopicModel `json:"custom"`
}

func (p *Product) SaveProduct() (id int) {
	product_key := tool.GenerateProductKey()
	product_secret := tool.GenerateProductSecret()
	p.ProductKey = product_key
	p.ProductSecret = product_secret

	mongodb_model_id := ProductSaveModel(p.ObjectModelID, p.ProductKey)
	p.MongodbModelID = mongodb_model_id

	id = MysqlInsertOneData(p)
	return id
}

func (d *Device) SaveDevice() (id int) {
	device_secret := tool.GenerateDeviceSecret()
	iot_id := tool.GenerateIotId()
	d.DeviceSecret = device_secret
	d.IotID = iot_id
	data_id := MysqlInsertOneData(d)
	return data_id
}

func ProductSaveModel(base_model_id int, product_key string) (mongodb_model_id int) {
	intact_product_tab := config.MongodbConfig.IntactProductModel
	concise_product_tab := config.MongodbConfig.ConciseProductModel
	concise_base_tab := config.MongodbConfig.BaseModelConcise
	intact_base_tab := config.MongodbConfig.BaseModelIntact

	filter := bson.M{"id": base_model_id}

	concise_data := MongoDbGetFilterData(concise_base_tab, filter)
	delete(concise_data, "_id")
	delete(concise_data, "id")

	intact_data := MongoDbGetFilterData(intact_base_tab, filter)
	delete(intact_data, "_id")
	delete(intact_data, "id")
	intact_data["profile"] = map[string]string{
		"productKey": product_key,
	}

	concise_id_str := MongoDbInsertOneData(concise_product_tab, concise_data)
	intact_id_str := MongoDbInsertOneData(intact_product_tab, intact_data)

	mongo_model := &MongodbModel{
		ConciseModelID: concise_id_str,
		IntactModelID:  intact_id_str,
	}
	save_id := MysqlInsertOneData(mongo_model)
	return save_id
}

func ProductDeleteMongodbModel(pid int) {
	db := DbConn()
	defer db.Close()

	var product Product
	db.First(&product, pid)
	db.Model(&product).Related(&product.MongodbModel, "MongodbModel")
	concise_model_id := product.MongodbModel.ConciseModelID
	concise_collection_name := config.MongodbConfig.ConciseProductModel
	MongodbDeleteOneData(concise_collection_name, concise_model_id)

	intact_collection_name := config.MongodbConfig.IntactProductModel
	intact_model_id := product.MongodbModel.IntactModelID
	MongodbDeleteOneData(intact_collection_name, intact_model_id)
}

func ProductSaveCustomTopic(pid int) {
	base_cuntom_topic := []CustomTopic{
		{
			ProductID:    pid,
			PermissionID: 2,
			Detail:       "/%s/%s/user/update",
			Describe:     "",
		},
		{
			ProductID:    pid,
			PermissionID: 2,
			Detail:       "/%s/%s/user/update/error",
			Describe:     "",
		},
		{
			ProductID:    pid,
			PermissionID: 1,
			Detail:       "/%s/%s/user/get",
			Describe:     "",
		},
	}

	for _, value := range base_cuntom_topic {
		data := CustomTopic{}
		data.Describe = value.Describe
		data.ProductID = pid
		data.Detail = value.Detail
		data.PermissionID = value.PermissionID
		_ = MysqlInsertOneData(&data)
	}
}

func GetCustomTopic(pid int) (data []TopicModel) {
	db := DbConn()
	defer db.Close()

	data_list := []CustomTopic{}
	result := []TopicModel{}
	db.Where("product_id = ?", pid).Find(&data_list)
	for _, value := range data_list {
		var a TopicModel
		a.ID = value.ID
		a.Name = ""
		a.Desc = value.Describe
		a.Operation = value.PermissionID
		a.TopicShortName = value.Detail
		result = append(result, a)
	}
	return result
}

func GetTopicModels(id int) (result Topics) {
	var data Topics
	base_data := []TopicModel{
		{
			Name:           "固件升级",
			Desc:           "设备上报固件升级信息",
			TopicShortName: "/ota/device/inform/%s/%s",
			Operation:      2,
		},
		{
			Name:           "固件升级",
			Desc:           "固件升级信息下行",
			TopicShortName: "/ota/device/upgrade/%s/%s",
			Operation:      1,
		},
		{
			Name:           "固件升级",
			Desc:           "设备上报固件升级进度",
			TopicShortName: "/ota/device/progress/%s/%s",
			Operation:      2,
		},
		{
			Name:           "固件升级",
			Desc:           "设备主动拉取固件升级信息",
			TopicShortName: "/ota/device/request/%s/%s",
			Operation:      2,
		},
		{
			Name:           "设备标签",
			Desc:           "设备上报标签数据",
			TopicShortName: "/sys/%s/%s/thing/deviceinfo/update",
			Operation:      2,
		},
		{
			Name:           "设备标签",
			Desc:           "云端响应标签上报",
			TopicShortName: "/sys/%s/%s/thing/deviceinfo/update_reply",
			Operation:      1,
		},
		{
			Name:           "设备标签",
			Desc:           "设备删除标签信息",
			TopicShortName: "/sys/%s/%s/thing/deviceinfo/delete",
			Operation:      1,
		},
		{
			Name:           "设备标签",
			Desc:           "云端响应标签删除",
			TopicShortName: "/sys/%s/%s/thing/deviceinfo/delete_reply",
			Operation:      2,
		},
		{
			Name:           "时钟同步",
			Desc:           "ntp时钟同步请求",
			TopicShortName: "/ext/ntp/%s/%s/request",
			Operation:      2,
		},
		{
			Name:           "时钟同步",
			Desc:           "ntp时钟同步响应",
			TopicShortName: "/ext/ntp/%s/%s/response",
			Operation:      1,
		},
		{
			Name:           "设备影子",
			Desc:           "设备影子发布",
			TopicShortName: "/shadow/update/%s/%s",
			Operation:      2,
		},
		{
			Name:           "设备影子",
			Desc:           "设备接收影子变更",
			TopicShortName: "/shadow/get/%s/%s",
			Operation:      1,
		},
		{
			Name:           "配置更新",
			Desc:           "云端主动下推配置信息",
			TopicShortName: "/sys/%s/%s/thing/config/push",
			Operation:      1,
		},
		{
			Name:           "配置更新",
			Desc:           "设备端查询配置信息",
			TopicShortName: "/sys/%s/%s/thing/config/get",
			Operation:      2,
		},
		{
			Name:           "配置更新",
			Desc:           "云端响应配置信息",
			TopicShortName: "/sys/%s/%s/thing/config/get_reply",
			Operation:      1,
		},
		{
			Name:           "广播",
			Desc:           "广播topic，identifer用户自定义字符串",
			TopicShortName: "/broadcast/%s/%s",
			Operation:      1,
		},
	}
	data.Base = base_data

	physics_data := []TopicModel{
		{
			Name:           "属性上报",
			Desc:           "设备属性上报",
			TopicShortName: "/sys/%s/%s/thing/event/property/post",
			Operation:      2,
		},
		{
			Name:           "属性上报",
			Desc:           "云端响应属性上报",
			TopicShortName: "/sys/%s/%s/thing/event/property/post_reply",
			Operation:      1,
		},
		{
			Name:           "属性设置",
			Desc:           "设备属性设置",
			TopicShortName: "/sys/%s/%s/thing/service/property/set",
			Operation:      1,
		},
		{
			Name:           "事件上报",
			Desc:           "设备事件上报",
			TopicShortName: "/sys/%s/%s/thing/event/%s/post",
			Operation:      2,
		},
		{
			Name:           "事件上报",
			Desc:           "云端响应事件上报",
			TopicShortName: "/sys/%s/%s/thing/event/%s/post_reply",
			Operation:      1,
		},
		{
			Name:           "服务调用",
			Desc:           "设备服务调用",
			TopicShortName: "/sys/%s/%s/thing/service/%s",
			Operation:      1,
		},
		{
			Name:           "服务调用",
			Desc:           "设备端响应服务调用",
			TopicShortName: "/sys/%s/%s/thing/service/%s_reply",
			Operation:      2,
		},
	}
	data.Physics = physics_data

	custom_data := GetCustomTopic(id)
	data.Custom = custom_data
	return data
}

func GetTopics(pid, did int) (topic Topics) {
	db := DbConn()
	defer db.Close()

	datas := GetTopicModels(pid)
	var product Product
	db.First(&product, pid)
	product_key := product.ProductKey

	device_name := ""
	if did == 0 {
		device_name = "${deviceName}"
	} else {
		var device Device
		db.First(&device, did)
		device_name = device.Name
	}

	for index, value := range datas.Base {
		if value.Name == "广播" {
			value.TopicShortName = fmt.Sprintf(value.TopicShortName, product_key, "${identifer}")
			datas.Base[index] = value
		} else {
			value.TopicShortName = fmt.Sprintf(value.TopicShortName, product_key, device_name)
			datas.Base[index] = value
		}
	}

	for index, value := range datas.Physics {
		switch value.Name {
		case "事件上报":
			value.TopicShortName = fmt.Sprintf(value.TopicShortName, product_key, device_name, "event/${tsl.event.identifer}")
			datas.Physics[index] = value
		case "服务调用":
			value.TopicShortName = fmt.Sprintf(value.TopicShortName, product_key, device_name, "${tsl.service.identifer}")
			datas.Physics[index] = value
		default:
			value.TopicShortName = fmt.Sprintf(value.TopicShortName, product_key, device_name)
			datas.Physics[index] = value
		}
	}

	for index, value := range datas.Custom {
		value.TopicShortName = fmt.Sprintf(value.TopicShortName, product_key, device_name)
		datas.Custom[index] = value
	}

	return datas
}

func DeviceNameToDevice(product_key, device_name string) (device Device) {
	db := DbConn()
	defer db.Close()

	var device_model Device
	var product Product
	db.Where("product_key = ?", product_key).First(&product)
	product_id := product.ID
	db.Where("name = ? AND product_id= ?", device_name, product_id).First(&device_model)

	return device_model
}

func GetIntactModel(producy_key string) (result primitive.M) {
	db := DbConn()
	defer db.Close()

	var product Product
	db.Where("product_key = ?", producy_key).First(&product)
	db.Model(&product).Related(&product.MongodbModel, "MongodbModel")
	intact_model_id := product.MongodbModel.IntactModelID
	intact_collection_name := config.MongodbConfig.IntactProductModel

	model_id, _ := primitive.ObjectIDFromHex(intact_model_id)
	filter := bson.M{"_id": model_id}
	data := MongoDbGetFilterData(intact_collection_name, filter)
	delete(data, "_id")

	return data
}

func GetProductModelInfo(pid int) (intact, concise primitive.M) {
	db := DbConn()
	defer db.Close()

	var product Product
	db.First(&product, pid)
	db.Model(&product).Related(&product.MongodbModel, "MongodbModel")

	product_key := product.ProductKey

	intact_model_id_str := product.MongodbModel.IntactModelID
	intact_model_id, _ := primitive.ObjectIDFromHex(intact_model_id_str)
	intact_filter := bson.M{"_id": intact_model_id}

	concise_model_id_str := product.MongodbModel.ConciseModelID
	concise_model_id, _ := primitive.ObjectIDFromHex(concise_model_id_str)
	concise_filter := bson.M{"_id": concise_model_id}

	intact_collection_name := config.MongodbConfig.IntactProductModel
	intact_data := MongoDbGetFilterData(intact_collection_name, intact_filter)
	delete(intact_data, "_id")
	profile := map[string]string{"productKey": product_key}
	intact_data["profile"] = profile

	concise_collection_name := config.MongodbConfig.ConciseProductModel
	concise_data := MongoDbGetFilterData(concise_collection_name, concise_filter)
	delete(concise_data, "_id")

	return intact_data, concise_data
}

func DeatLabelQueryFilter(key, value string) (filter string) {

	key_filter_model := "\"%s\":"
	value_filter_model := ":\"%s\""
	filter_model := "\"%s\":\"%s\""

	label_filter := ""

	if len(key) == 0 && len(value) == 0 {
		label_filter = ""
	} else if len(key) != 0 && len(value) == 0 {
		label_filter = "%" + fmt.Sprintf(key_filter_model, key) + "%"
	} else if len(key) == 0 && len(value) != 0 {
		label_filter = "%" + fmt.Sprintf(value_filter_model, value) + "%"
	} else if len(key) != 0 && len(value) != 0 {
		label_filter = "%" + fmt.Sprintf(filter_model, key, value) + "%"
	}

	return label_filter
}

func DealLabelArgs(args []map[string]string) (label_str string) {
	result := make(map[string]string)
	for _, value := range args {
		result[value["key"]] = value["value"]
	}
	result_str := tool.MapToJsonStr(result)
	return result_str
}

func DeviceOnline(iot_id string) {
	db := DbConn()
	defer db.Close()

	var device Device
	db.Where("iot_id = ?", iot_id).First(&device)
	device.Online = true
	db.Save(&device)
}

func DeviceOutline(iot_id string) {
	db := DbConn()
	defer db.Close()

	var device Device
	db.Where("iot_id = ?", iot_id).First(&device)
	device.Online = false
	device.LastOnLineTime = time.Now()
	db.Save(&device)
}

func ProductGetPropertyFunction(productID int) (properties []map[string]interface{}) {

	db := DbConn()
	defer db.Close()

	var product Product
	db.First(&product, productID)
	db.Model(&product).Related(&product.MongodbModel, "MongodbModel")
	intact_model_id := product.MongodbModel.IntactModelID
	intact_collection_name := config.MongodbConfig.IntactProductModel

	model_id, _ := primitive.ObjectIDFromHex(intact_model_id)
	filter := bson.M{"_id": model_id}
	data := MongoDbGetFilterData(intact_collection_name, filter)
	delete(data, "_id")

	modelJson, _ := json.Marshal(data)
	result := gjson.Get(string(modelJson), "properties")

	for _, v := range result.Array() {
		property := make(map[string]interface{})
		property["identifier"] = v.Map()["identifier"].String()
		property["name"] = v.Map()["name"].String()
		property["accessMode"] = v.Map()["accessMode"].String()
		property["desc"] = v.Map()["desc"].String()
		property["required"] = v.Map()["required"].String()
		property["data_type"] = v.Map()["dataType"].Map()["type"].String()

		condition := v.Map()["dataType"].Map()["specs"].Array()
		if condition == nil {
			property["data_condition"] = tool.JsonStrToMap(v.Map()["dataType"].Map()["specs"].Raw)
		} else {
			var need_data []map[string]interface{}
			for _, value := range condition {
				resolve_data := make(map[string]interface{})
				resolve_data["identifier"] = value.Map()["identifier"].String()
				resolve_data["name"] = value.Map()["name"].String()
				resolve_data["data_type"] = value.Map()["dataType"].Map()["type"].String()
				resolve_data["data_condition"] = tool.JsonStrToMap(value.Map()["dataType"].Map()["specs"].Raw)
				need_data = append(need_data, resolve_data)
			}
			property["data_condition"] = need_data
		}
		properties = append(properties, property)
	}
	return properties
}

func GetAllProductID() (result []int) {
	db := DbConn()
	defer db.Close()

	var ids []int
	db.Model(Product{}).Pluck("id", &ids)

	return ids
}

func GetProductLabel(productKey string) (label map[string]string) {
	db := DbConn()
	defer db.Close()

	var product Product
	db.Where("product_key = ?", productKey).First(&product)

	productLabel := product.Label
	return tool.JsonStrToMap(productLabel)
}

func UpdateProductLabel(productKey string, newLabel map[string]string) {
	db := DbConn()
	defer db.Close()

	var product Product
	db.Where("product_key = ?", productKey).First(&product)
	label := tool.MapToJsonStr(newLabel)
	product.Label = label
	db.Save(&product)
}

func GetDeviceLabel(iotID string)(label map[string]string){
	db := DbConn()
	defer db.Close()

	var device Device
	db.Where("iot_id = ?", iotID).First(&device)

	deviceLabel := device.Label
	return tool.JsonStrToMap(deviceLabel)
}

func UpdateDeviceLabel(iotID string, newLabel map[string]string) {
	db := DbConn()
	defer db.Close()

	var device Device
	db.Where("iot_id = ?", iotID).First(&device)
	label := tool.MapToJsonStr(newLabel)
	device.Label = label
	db.Save(&device)
}
