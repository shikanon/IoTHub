package api

import (
	"fmt"
	"github.com/shikanon/IoTOrbHub/config"
	"github.com/shikanon/IoTOrbHub/pkg/database"
	"github.com/shikanon/IoTOrbHub/pkg/hub"
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

// 1. 生成productKey, productSecret
// 2. 绑定model，存储mongodb
// 3. 存储topic信息到mysql数据库
// 4. 保存product信息到mysql数据库
func (p *Product) SaveProduct() (id int) {
	product_key := hub.GenerateProductKey()
	product_secret := hub.GenerateProductSecret()
	p.ProductKey = product_key
	p.ProductSecret = product_secret

	mongodb_model_id := ProductSaveModel(p.ObjectModelID, p.ProductKey)
	p.MongodbModelID = mongodb_model_id

	id = database.MysqlInsertOneData(p)
	return id
}

// 绑定deviceKey,iotid
func (d *Device) SaveDevice() (id int) {
	device_secret := hub.GenerateDeviceSecret()
	iot_id := hub.GenerateIotId()
	d.DeviceSecret = device_secret
	d.IotID = iot_id
	data_id := database.MysqlInsertOneData(d)
	return data_id
}

// 提产品保存物模型
func ProductSaveModel(id int, key string) (mongodb_model_id int) {
	intact_product_tab := config.MongodbConfig.IntactProductModel
	concise_product_tab := config.MongodbConfig.ConciseProductModel
	concise_base_tab := config.MongodbConfig.BaseModelConcise
	intact_base_tab := config.MongodbConfig.BaseModelIntact

	filter := bson.M{"id": id}

	concise_data := database.MongoDbGetFilterData(concise_base_tab, filter)
	delete(concise_data, "_id")
	delete(concise_data, "id")

	intact_data := database.MongoDbGetFilterData(intact_base_tab, filter)
	delete(intact_data, "_id")
	delete(intact_data, "id")
	intact_data["profile"] = map[string]string{
		"productKey": key,
	}

	concise_id_str := database.MongoDbInsertOneData(concise_product_tab, concise_data)
	intact_id_str := database.MongoDbInsertOneData(intact_product_tab, intact_data)

	mongo_model := &MongodbModel{
		ConciseModelID: concise_id_str,
		IntactModelID:  intact_id_str,
	}
	id = database.MysqlInsertOneData(mongo_model)
	return id
}

// 提供key，生成topic，并存储数据库
func ProductSaveCustomTopic(id int) {
	base_cuntom_topic := []CustomTopic{
		{
			ProductID:    id,
			PermissionID: 2,
			Detail:       "/%s/%s/user/update",
			Describe:     "",
		},
		{
			ProductID:    id,
			PermissionID: 2,
			Detail:       "/%s/%s/user/update/error",
			Describe:     "",
		},
		{
			ProductID:    id,
			PermissionID: 1,
			Detail:       "/%s/%s/user/get",
			Describe:     "",
		},
	}

	for _, value := range base_cuntom_topic {
		data := CustomTopic{}
		data.Describe = value.Describe
		data.ProductID = id
		data.Detail = value.Detail
		data.PermissionID = value.PermissionID
		_ = database.MysqlInsertOneData(&data)
	}
}

// 获取用户自定义Topic模板
func GetCustomTopic(pid int) (data []TopicModel) {
	db := database.DbConn()
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

// 获取topic模板
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

// 获取Topic类
func GetTopics(pid, did int) (topic Topics) {
	db := database.DbConn()
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

// 设备名称获取设备
func DeviceNameToDevice(device_name string) (device Device) {
	db := database.DbConn()
	defer db.Close()

	var device_model Device
	db.Where("name = ?", device_name).First(&device_model)
	return device_model
}

// 获取完整物模型 TODO
func GetIntactModel(producy_key string) (result primitive.M) {
	db := database.DbConn()
	defer db.Close()

	var product Product
	db.Where("product_key = ?", producy_key).First(&product)
	db.Model(&product).Related(&product.MongodbModel, "MongodbModel")
	intact_model_id := product.MongodbModel.IntactModelID
	intact_collection_name := config.MongodbConfig.IntactProductModel
	model_id, _ := primitive.ObjectIDFromHex(intact_model_id)
	filter := bson.M{"_id": model_id}
	data := database.MongoDbGetFilterData(intact_collection_name, filter)
	fmt.Printf("%T\n", data)
	return data
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
