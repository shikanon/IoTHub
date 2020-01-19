package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shikanon/IoTOrbHub/config"
	"github.com/shikanon/IoTOrbHub/pkg/database"
	"github.com/shikanon/IoTOrbHub/pkg/util"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
)

// *******************************************************************************************************
func Home(c *gin.Context) {
	//db_client := database.MongoDbClient()
	//
	//ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//collection := db_client.Collection("base_model_concise")
	//cur, _ := collection.Find(ctx, bson.D{})
	//defer cur.Close(ctx)
	//for cur.Next(ctx) {
	//	var result bson.M
	//	err := cur.Decode(&result)
	//	if err != nil { log.Fatal(err) }
	//	fmt.Println(result)
	//	fmt.Printf("%T\n", result)
	//}
	//if err := cur.Err(); err != nil {
	//	log.Fatal(err)
	//}
	//
	//db := database.DbConn()
	//defer db.Close()
	//
	//data := Model{
	//	MongodbID:   2,
	//	Name:        "路灯照明",
	//	Scene:       "公共服务",
	//	TerritoryID: 1,
	//}
	//db.Create(&data)

	//
	//var devices []Device
	//db.Where("activation_time = ?", "0000-00-00 00:00:00").Find(&devices)
	//fmt.Println(len(devices))

	//
	//var device Device
	//db.First(&device, 1)
	////time := TimeDeal(device.CreateTime)
	//time := TimeDeal(device.ActivationTime)
	//fmt.Println(time)
	//fmt.Printf("%T\n", time)

	//db.AutoMigrate(&Product{}, &Device{}, &Model{}, &ModelTerritory{}, &NodeType{},
	//	&NetworkWay{}, &DataFormat{}, &AuthMethod{}, &DeviceStatus{}, &CustomTopic{},
	//	&TopicPermission{}, &MongodbModel{})
	//var device Device
	//d
	//db.First(&device, 1)

	//data := GetIntactModel("01DYPKYMSGD4CSNK86Z8H09WBC") // 5e201036c39081ab3d77e230
	//data := DeviceNameToDevice("01DYPJKERMSARPYM1ZM8ZJEV7A", "demo_01")

	db := config.MongodbConfig.BaseModelConcise
	filter := bson.M{"id": 1}
	data := database.MongoDbGetFilterData(db, filter)
	fmt.Println(data["_id"])
	fmt.Printf("%T", data["_id"])


	//data := Model{
	//	MongodbID:   1,
	//	Name:        "车辆定位卡",
	//	Scene:       "公共服务",
	//	TerritoryID: 1,
	//}
	//db.Create(&data)

	//
	//data := &TopicPermission{
	//	Name: "发布和订阅",
	//}
	//db.Create(data)

	//db := config.MongodbConfig.Db
	//port := config.MongodbConfig.Port
	//host := config.MongodbConfig.Host
	//base1 := config.MongodbConfig.BaseModelConcise
	//base2 := config.MongodbConfig.BaseModelIntact
	//product := config.MongodbConfig.ProductModel
	//fmt.Println(db, port, host, base1, base2, product)

	resp := gin.H{
		"status":  "Y",
		"message": "数据格式查询成功",
		//"data":    ".............",
		"data":    data,
	}
	c.JSON(200, resp)

}

// *********************************************************************************************************

// TODO  权限；参数校验；捕捉错误;标签字段的拆解

func GetProductModels(c *gin.Context) {
	db := database.DbConn()
	defer db.Close()

	var models []database.Model
	db.Preload("Territory").Find(&models)

	resp := gin.H{
		"status":  "Y",
		"message": "所有模型信息查询成功",
		"data":    models,
	}
	c.JSON(200, resp)
}

func GetProductNodeTypes(c *gin.Context) {
	db := database.DbConn()
	defer db.Close()

	var nodetypes []database.NodeType
	db.Find(&nodetypes)

	resp := gin.H{
		"status":  "Y",
		"message": "节点类型查询成功",
		"data":    nodetypes,
	}
	c.JSON(200, resp)
}

func GetProductNetworkWays(c *gin.Context) {
	db := database.DbConn()
	defer db.Close()

	var networks []database.NetworkWay
	db.Find(&networks)

	resp := gin.H{
		"status":  "Y",
		"message": "联网方式查询成功",
		"data":    networks,
	}
	c.JSON(200, resp)
}

func GetProductDataFormats(c *gin.Context) {
	db := database.DbConn()
	defer db.Close()

	var dataformats []database.DataFormat
	db.Find(&dataformats)

	resp := gin.H{
		"status":  "Y",
		"message": "数据格式查询成功",
		"data":    dataformats,
	}
	c.JSON(200, resp)
}

func GetProductAuthMethods(c *gin.Context) {
	db := database.DbConn()
	defer db.Close()

	var authmethos []database.AuthMethod
	db.Find(&authmethos)

	resp := gin.H{
		"status":  "Y",
		"message": "数据格式查询成功",
		"data":    authmethos,
	}
	c.JSON(200, resp)
}

func GetProducts(c *gin.Context) {
	// 获取参数
	page_str := c.Query("page")
	page, _ := strconv.Atoi(page_str)
	item_str := c.Query("item")
	item, _ := strconv.Atoi(item_str)

	// 查询数据
	db := database.DbConn()
	defer db.Close()
	var products []database.Product
	db.Limit(item).Offset((page - 1) * item).Order("id desc").Preload("NodeType").Find(&products)
	// 构建响应数据结构
	type info struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		ProductKey string `json:"product_key"`
		CreateTime string `json:"create_time"`
		NodeType   string `json:"node_type"`
	}
	result := []info{}
	for _, val := range products {
		data := info{
			ID:         val.ID,
			Name:       val.Name,
			ProductKey: val.ProductKey,
			CreateTime: database.TimeDeal(val.CreateTime),
			NodeType:   val.NodeType.Name,
		}
		result = append(result, data)
	}

	var total = 0
	db.Model(&database.Product{}).Count(&total)

	type RespData struct {
		NumResults int    `json:"num_results"`
		DataList   []info `json:"data_list"`
	}

	var resp_data = RespData{
		NumResults: total,
		DataList:   result,
	}

	//  响应
	resp := gin.H{
		"status":  "Y",
		"message": "产品首页信息查询成功",
		"data":    resp_data,
	}
	c.JSON(200, resp)
}

func AddProduct(c *gin.Context) {
	type Data struct {
		Name         string `form:"name" json:"name" binding:"required"`
		Category     string `form:"category" json:"category" binding:"required"`
		ModelID      int    `form:"model_id" json:"model_id" binding:"required"`
		NodeTypeID   int    `form:"node_type_id" json:"node_type_id" binding:"required"`
		NetworkID    int    `form:"network_id" json:"network_id" binding:"required"`
		DataFormatID int    `form:"data_format_id" json:"data_format_id" binding:"required"`
		AuthMethodID int    `form:"auth_method_id" json:"auth_method_id" binding:"required"`
		Describe     string `form:"desc" json:"desc" binding:"required"`
	}

	data := Data{}
	if err := c.ShouldBind(&data); err != nil {
		fmt.Println(err)
	}

	name := data.Name
	category := data.Category
	model_id := data.ModelID
	node_type_id := data.NodeTypeID
	network_id := data.NetworkID
	data_format_id := data.DataFormatID
	auth_method_id := data.AuthMethodID
	desc := data.Describe

	// 构建实例
	product := &database.Product{
		Name:          name,
		Category:      category,
		ObjectModelID: model_id,
		NodeTypeID:    node_type_id,
		NetworkWayID:  network_id,
		DataFormatID:  data_format_id,
		AuthMethodID:  auth_method_id,
		Describe:      desc,
	}

	// mysql持久化存储。存储topic
	id := product.SaveProduct()
	database.ProductSaveCustomTopic(id)

	// 构建，返回响应
	resp := gin.H{
		"status":  "Y",
		"message": "产品创建成功",
		"data":    map[string]int{"product_id": id},
		//"data": id,
	}
	c.JSON(200, resp)
}

func GetProduct(c *gin.Context) {
	product_id := c.Query("pid")
	db := database.DbConn()
	defer db.Close()

	type Response struct {
		ID              int    `json:"id"`
		Name            string `json:"name"`
		NodeType        string `json:"node_type"`
		NodeTypeID      int    `json:"node_type_id"`
		CreateTime      string `json:"create_time"`
		ObjectModelName string `json:"object_model_name"`
		DataFormat      string `json:"data_format"`
		DataFormatID    int    `json:"data_format_id"`
		AuthMethod      string `json:"auth_method"`
		AuthMethodID    int    `json:"auth_method_id"`
		Status          string `json:"status"`
		NetworkWay      string `json:"network_way"`
		NetworkWayID    int    `json:"network_way_id"`
		Describe        string `json:"desc"`
		ConciseModelID  string `json:"concise_model_id"`
		IntactModelID   string `json:"intact_model_id"`
		ProductKey      string `json:"product_key"`
		ProductSecret   string `json:"product_secret"`
		DeviceCount     int    `json:"device_count"` // TODO
	}

	var product database.Product
	db.First(&product, product_id)
	db.Model(&product).Related(&product.ObjectModel, "ObjectModel")
	db.Model(&product).Related(&product.MongodbModel, "MongodbModel")
	db.Model(&product).Related(&product.NodeType, "NodeType")
	db.Model(&product).Related(&product.NetworkWay, "NetworkWay")
	db.Model(&product).Related(&product.DataFormat, "DataFormat")
	db.Model(&product).Related(&product.AuthMethod, "AuthMethod")
	db.Model(&product).Related(&product.Device, "Device")

	response := Response{
		ID:              product.ID,
		Name:            product.Name,
		NodeType:        product.NodeType.Name,
		NodeTypeID:      product.NodeTypeID,
		CreateTime:      database.TimeDeal(product.CreateTime),
		ObjectModelName: product.ObjectModel.Name,
		DataFormat:      product.DataFormat.Name,
		DataFormatID:    product.DataFormatID,
		AuthMethod:      product.AuthMethod.Name,
		AuthMethodID:    product.AuthMethodID,
		Status:          product.Status,
		NetworkWay:      product.NetworkWay.Name,
		NetworkWayID:    product.NetworkWayID,
		Describe:        product.Describe,
		ConciseModelID:  product.MongodbModel.ConciseModelID,
		IntactModelID:   product.MongodbModel.IntactModelID,
		ProductKey:      product.ProductKey,
		ProductSecret:   product.ProductSecret,
		DeviceCount:     len(product.Device),
	}

	resp := gin.H{
		"status":  "Y",
		"message": "产品信息查询成功",
		"data":    response,
	}
	c.JSON(200, resp)
}

func UpdateProduct(c *gin.Context) {
	name := c.PostForm("name")
	desc := c.PostForm("desc")
	product_id := c.PostForm("pid")

	db := database.DbConn()
	defer db.Close()

	var product database.Product
	db.First(&product, product_id)
	product.Name = name
	product.Describe = desc
	db.Save(&product)

	resp := gin.H{
		"status":  "Y",
		"message": "产品信息更新成功",
		"data":    nil,
	}
	c.JSON(200, resp)
}

func GetProductTopic(c *gin.Context) {
	product_str := c.Query("pid")
	product_id, _ := strconv.Atoi(product_str)
	device_id := 0
	data := database.GetTopics(product_id, device_id)

	resp := gin.H{
		"status":  "Y",
		"message": "Topic查询成功",
		"data":    data,
	}
	c.JSON(200, resp)
}

func AddProductTopic(c *gin.Context) {
	product_str := c.PostForm("pid")
	product_id, _ := strconv.Atoi(product_str)
	permission_str := c.PostForm("permission_id")
	permission_id, _ := strconv.Atoi(permission_str)
	topic := c.PostForm("topic")
	desc := c.PostForm("desc")

	topic_detail := "/user/%s"
	detail := "/%s/%s" + fmt.Sprintf(topic_detail, topic)

	data := database.CustomTopic{
		ProductID:    product_id,
		PermissionID: permission_id,
		Detail:       detail,
		Describe:     desc,
	}

	id := database.MysqlInsertOneData(&data)
	resp := gin.H{
		"status":  "Y",
		"message": "自定义topic创建成功",
		"data":    id,
	}
	c.JSON(200, resp)
}

func UpdateProductTopic(c *gin.Context) {
	permission_str := c.PostForm("permission_id")
	permission_id, _ := strconv.Atoi(permission_str)
	topic := c.PostForm("topic")
	desc := c.DefaultPostForm("desc", "")

	topic_str := c.PostForm("pid")
	topic_id, _ := strconv.Atoi(topic_str)

	topic_detail := "/user/%s"
	detail := "/%s/%s" + fmt.Sprintf(topic_detail, topic)

	db := database.DbConn()
	defer db.Close()
	var topic_model database.CustomTopic
	db.Find(&topic_model, topic_id)
	topic_model.PermissionID = permission_id
	topic_model.Detail = detail
	topic_model.Describe = desc
	db.Save(&topic_model)

	resp := gin.H{
		"status":  "Y",
		"message": "数据更新成功",
		"data":    nil,
	}
	c.JSON(200, resp)
}

//// 产品-查看-topic类，自定义，删除topic类
//func DeleteProductTopic(c *gin.Context) {
//	topidid := 1
//	accesskeyid := "aaa"
//	signature := "bbb"
//}

//// 产品-查看-功能定义查看
//// functions?pid=1&accesskeyid=aaaa&signature=bbb
//func GetProductFunction(c *gin.Context) {}
//
//// 产品-查看-导入物模型 TODO
//// 产品-查看-查看物模型
//// model?pid=1&accesskeyid=aaaa&signature=bbb
//func DownloadModelJS(c *gin.Context) {
//	// 返回两个模型。完整模型和简单模型
//}
//
//// 产品-查看-生成设备端代码
//// code?pid=1&accesskeyid=aaaa&signature=bbb
//func DownloadEquipmentCode(c *gin.Context) {}
//
//// 产品-查看-标准功能，添加 TODO
//// 产品-查看-标准功能，编辑 TODO
//// 产品-查看-自定义功能，添加 TODO
//// 产品-查看-自定义功能，编辑 TODO
//// 产品-查看-自定义功能，删除 TODO

func DeleteProduct(c *gin.Context) {
	type Data struct {
		ProductID int `json:"pid"`
	}

	data := Data{}
	if err := c.ShouldBind(&data); err != nil {
		fmt.Println(err)
	}

	product_id := data.ProductID

	fmt.Println(product_id)

	db := database.DbConn()
	defer db.Close()

	var devices []database.Device
	db.Where("product_id = ?", product_id).Find(&devices)

	if (len(devices) > 0 ){
		resp := gin.H{
			"status":  "N",
			"message": "产品下还存在设备",
			"data":    nil,
		}
		c.JSON(200, resp)
	} else {
		var product database.Product
		db.Where("id = ?", product_id).First(&product)
		mongodb_model_id := product.MongodbModelID

		database.ProductDeleteMongodbModel(product_id)
		db.Where("id = ?", product_id).Delete(&database.Product{})
		db.Where("id = ?", mongodb_model_id).Delete(&database.MongodbModel{})
		db.Where("product_id = ?", product_id).Delete(&database.CustomTopic{})

		resp := gin.H{
			"status":  "Y",
			"message": "设备删除成功",
			"data":    nil,
		}
		c.JSON(200, resp)
	}
}

func GetDevices(c *gin.Context) {
	// product为0，设备首页
	// product不为0，产品-管理设备 / 产品-查看-前往管理
	product_str := c.Query("pid")
	product_id, _ := strconv.Atoi(product_str)
	page_str := c.Query("page")
	page, _ := strconv.Atoi(page_str)
	item_str := c.Query("item")
	item, _ := strconv.Atoi(item_str)

	var total = 0
	var activate_num = 0
	var online_num = 0
	db := database.DbConn()
	defer db.Close()
	var devices []database.Device
	if product_id == 0 {
		db.Where("activation_time != ?", "0000-00-00 00:00:00").Find(&devices)
		activate_num = len(devices)
		db.Where("last_on_line_time != ?", "0000-00-00 00:00:00").Find(&devices)
		online_num = len(devices)
		db.Limit(item).Offset((page - 1) * item).Order("id desc").Preload("Status").Find(&devices)
		db.Model(&database.Device{}).Count(&total)
	} else {
		db.Where("product_id = ? AND activation_time != ?", product_id, "0000-00-00 00:00:00").Find(&devices)
		activate_num = len(devices)
		db.Where("product_id = ? AND last_on_line_time != ?", product_id, "0000-00-00 00:00:00").Find(&devices)
		online_num = len(devices)
		db.Where("product_id = ?", product_id).Find(&devices)
		total = len(devices)
		db.Where("product_id = ?", product_id).Limit(item).Offset((page - 1) * item).Order("id desc").Preload("Status").Find(&devices)
	}

	type Response struct {
		ID               int    `json:"id"`                 // 设备id
		Name             string `json:"name"`               // 备注名称
		TheirProductName string `json:"their_product_name"` // 所属产品
		NodeType         string `json:"node_type"`          // 节点类型
		NodeTypeID       int    `json:"node_type_id"`       // 节点类型id
		StatusID         int    `json:"status_id"`          // 状态id
		Status           string `json:"status"`             // 状态
		LastOnLineTime   string `json:"last_on_line_time"`  // 最后上线时间
	}
	var responses []Response
	for _, device := range devices {
		var data Response
		data.ID = device.ID
		data.Name = device.Name
		var product database.Product
		db.First(&product, device.ProductID).Related(&product.NodeType, "NodeType")
		data.TheirProductName = product.Name
		data.NodeType = product.NodeType.Name
		data.NodeTypeID = product.NodeTypeID
		data.Status = device.Status.Name
		data.StatusID = device.StatusID
		data.LastOnLineTime = database.TimeDeal(device.LastOnLineTime)
		responses = append(responses, data)
	}

	type RespData struct {
		ActivateCount int        `json:"device_active_count"`
		OnlineCount   int        `json:"device_online_count"`
		NumResults    int        `json:"num_results"`
		DataList      []Response `json:"data_list"`
	}

	var resp_data = RespData{
		ActivateCount: activate_num,
		OnlineCount:   online_num,
		NumResults:    total,
		DataList:      responses,
	}

	resp := gin.H{
		"status":  "Y",
		"message": "所有设备查询成功",
		"data":    resp_data,
	}
	c.JSON(200, resp)
}

func GetSimpleProducts(c *gin.Context) {
	db := database.DbConn()
	defer db.Close()

	var products []database.Product
	db.Find(&products)

	type info struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	result := []info{}
	for _, val := range products {
		data := info{
			ID:   val.ID,
			Name: val.Name,
		}
		result = append(result, data)
	}

	resp := gin.H{
		"status":  "Y",
		"message": "所有产品名称查询成功",
		"data":    result,
	}
	c.JSON(200, resp)
}

func AddDevice(c *gin.Context) {
	db := database.DbConn()
	defer db.Close()

	type Data struct {
		ProductID int    `form:"pid" json:"pid" binding:"required"`
		Name      string `form:"name" json:"name" binding:"required"`
		Remark    string `form:"remark" json:"remark" binding:"required"`
	}

	data := Data{}
	if err := c.ShouldBind(&data); err != nil {
		fmt.Println(err)
	}

	product_id := data.ProductID
	name := data.Name
	remark := data.Remark

	device := database.Device{
		ProductID:   product_id,
		StatusID:    1,
		Name:        name,
		Remark:      remark,
		BatchCreate: false,
	}

	id := device.SaveDevice()

	var product database.Product
	db.First(&product, product_id)
	var device_save database.Device
	db.First(&device_save, id)

	type RespData struct {
		ProductKey   string `json:"product_key"`
		DeviceName   string `json:"device_name"`
		DeviceSecret string `json:"device_secret"`
	}
	var resp_data RespData
	resp_data.ProductKey = product.ProductKey
	resp_data.DeviceName = device_save.Name
	resp_data.DeviceSecret = device_save.DeviceSecret

	resp := gin.H{
		"status":  "Y",
		"message": "产品创建成功",
		"data":    resp_data,
	}
	c.JSON(200, resp)
}

//// 设备-批量添加，自动生成
//func BatchAutomaticAddProduct(c *gin.Context) {
//	productid := 1
//	number := 10
//	accesskeyid := "aaa"
//	signature := "bbb"
//}
//
//// 设备-批量添加，批量上传，上传文件 TODO
//// 设备-批量添加，批量上传 TODO
//
//// 设备-批次管理
//// batchdevices?page=1&item=10&accesskeyid=aaaa&signature=bbb
//func GetBatchDevices(c *gin.Context) {}

// 设备-批次管理-详情
// batchdevice?pid=1&accesskeyid=aaaa&signature=bbb
//func GetBatchDevice(c *gin.Context) {}
//
//// 设备-批次管理-下载CSV
//func DownloadBatchDeviceCSV(c *gin.Context) {
//	productid := 1
//	accesskeyid := "aaa"
//	signature := "bbb"
//	// 过滤出批量生成的设备
//}
//
//// 设备-修改信息
//func UpdateDevice(c *gin.Context) {
//	statusid := 0
//	remark := "备注名称"
//	productid := 1
//	accesskeyid := "aaa"
//	signature := "bbb"
//}

func GetDevice(c *gin.Context) {
	device_str := c.Query("did")
	device_id, _ := strconv.Atoi(device_str)
	db := database.DbConn()
	defer db.Close()
	device := database.Device{}
	db.First(&device, device_id)
	type Response struct {
		ID             int    `json:"id"`
		ProductID      int    `json:"product_id"`
		ProductKey     string `json:"product_key"`
		ProductName    string `json:"product_name"`
		NoteTypeID     int    `json:"node_type_id"`
		NodeType       string `json:"node_type"`
		StatusID       int    `json:"status_id"`
		Name           string `json:"name"`
		Remark         string `json:"remark"`
		DeviceSecret   string `json:"device_secret"`
		IP             string `json:"ip"`
		CreateTime     string `json:"create_time"`
		ActivationTime string `json:"activate_time"`
		LastOnLineTime string `json:"last_online_time"`
		IotID          string `json:"iot_id"`
		Label          string `json:"label"`
		BatchCreate    bool   `json:"batch_create"`
	}
	var response Response
	response.ID = device.ID
	response.ProductID = device.ProductID
	response.StatusID = device.StatusID
	response.Name = device.Name
	response.Remark = device.Remark
	response.DeviceSecret = device.DeviceSecret
	response.IP = device.IP
	response.CreateTime = database.TimeDeal(device.CreateTime)
	response.ActivationTime = database.TimeDeal(device.ActivationTime)
	response.LastOnLineTime = database.TimeDeal(device.LastOnLineTime)
	response.IotID = device.IotID
	response.Label = device.Label
	response.BatchCreate = device.BatchCreate

	var product database.Product
	db.First(&product, response.ProductID)
	db.Model(&product).Related(&product.NodeType, "NodeType")
	response.ProductKey = product.ProductKey
	response.ProductName = product.Name
	response.NoteTypeID = product.NetworkWayID
	response.NodeType = product.NodeType.Name

	resp := gin.H{
		"status":  "Y",
		"message": "设备信息查询成功",
		"data":    response,
	}
	c.JSON(200, resp)
}

func GetDeviceTopic(c *gin.Context) {
	device_str := c.Query("did")
	device_id, _ := strconv.Atoi(device_str)

	db := database.DbConn()
	defer db.Close()

	device := database.Device{}
	db.First(&device, device_id)
	product_id := device.ProductID

	data := database.GetTopics(product_id, device_id)

	resp := gin.H{
		"status":  "Y",
		"message": "Topic查询成功",
		"data":    data,
	}
	c.JSON(200, resp)
}

//// 设备-查看-topic类，自定义，发布消息(限订阅)
//func CustomTopicSendMessage(c *gin.Context) {
//	topicid := 1
//	message := "hello world"
//	accesskeyid := "aaa"
//	signature := "bbb"
//}
//
//// 设备-查看-运行状态，图
//// mapstatus?did=1&accesskeyid=aaaa&signature=bbb
//func GetRunningStatusMap(c *gin.Context) {}
//
//// 设备-查看-运行状态，表
//// tablestatus??did=1&accesskeyid=aaaa&signature=bbb
//func GetRunningStatusTable(c *gin.Context) {}
//
//// 设备-查看-运行状态，实时刷新 TODO 待商讨
//// 设备-查看-运新状态，查看数据 TODO
//// 设备-查看-事件管理 TODO 暂缓
//// 设备-查看-服务调用 TODO 暂缓
//
//// 设备-查看-影子设备，查看
//// shadowdevice?did=1&accesskeyid=aaaa&signature=bbb
//func GetShadowDeviceInfo(c *gin.Context) {}
//
//// 设备-查看-影子设备，更新影子
//func UpdateShadowDevice(c *gin.Context) {
//	device := 1
//	message := `{
// "state": {
//   "reported": {},
//   "desired": {}
// },
// "metadata": {
//   "reported": {},
//   "desired": {}
// },
// "timestamp": 0,
// "version": 0
//}`
//	accesskeyid := "aaa"
//	signature := "bbb"
//}
//
//// 设备-查看-文件管理 TODO 暂缓
//// 设备-查看-日志服务 TODO 暂缓

func DeleteDevice(c *gin.Context) {
	type Data struct {
		DeviceIDList []int `json:"dids"`
	}

	data := Data{}
	if err := c.ShouldBind(&data); err != nil {
		fmt.Println(err)
	}

	device_id_list := data.DeviceIDList

	db := database.DbConn()
	defer db.Close()

	for index, _ := range device_id_list{
		db.Where("id = ?", device_id_list[index]).Delete(&database.Device{})
	}

	resp := gin.H{
		"status":  "Y",
		"message": "设备删除成功",
		"data":    nil,
	}
	c.JSON(200, resp)
}

func GetDeviceDesireStatus(c *gin.Context) {
	device_str := c.Query("did")
	device_id, _ := strconv.Atoi(device_str)

	db := database.DbConn()
	defer db.Close()

	var device database.Device
	db.First(&device, device_id)
	product_id := device.ProductID
	device_iot := device.IotID
	var product database.Product
	db.First(&product, product_id)
	product_key := product.ProductKey

	data := util.GetDeviceDesiredPropertyInfo(product_key, device_iot)

	resp := gin.H{
		"status":  "Y",
		"message": "设备期望状态查询成功",
		"data":    data,
	}
	c.JSON(200, resp)
}

func GetDevicePropertyStatus(c *gin.Context) {
	device_str := c.Query("did")
	device_id, _ := strconv.Atoi(device_str)

	db := database.DbConn()
	defer db.Close()

	var device database.Device
	db.First(&device, device_id)
	product_id := device.ProductID
	device_iot := device.IotID
	var product database.Product
	db.First(&product, product_id)
	product_key := product.ProductKey

	data := util.GetDevicePropertyStatusInfo(product_key, device_iot)

	resp := gin.H{
		"status":  "Y",
		"message": "设备期望状态查询成功",
		"data":    data,
	}
	c.JSON(200, resp)
}