package api

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shikanon/IoTOrbHub/pkg/database"
	"github.com/shikanon/IoTOrbHub/pkg/tool"
	"github.com/shikanon/IoTOrbHub/pkg/util"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var log = logrus.New()

func GetProductModels(c *gin.Context) {
	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
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
	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
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
	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
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
	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
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
	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
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
	//获取参数
	page := tool.StringNumberToInTNumber(c.Query("page"))
	item := tool.StringNumberToInTNumber(c.Query("item"))
	name := c.Query("name")
	key := c.Query("key")
	value := c.Query("value")

	if len(key) != 0 {
		if key_res, msg := CheckProductLabelKeyQualify(key); key_res != true {
			ErrResponse(msg, c)
			return
		}
	}
	if len(value) != 0 {
		if value_res, msg := CheckProductLabelValueQualify(value); value_res != true {
			ErrResponse(msg, c)
			return
		}
	}

	label_filter := database.DeatLabelQueryFilter(key, value)

	var total = 0

	// 查询数据
	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
	defer db.Close()
	var products []database.Product
	if len(name) == 0 && len(label_filter) == 0 {
		db.Model(&database.Product{}).Count(&total)
		db.Limit(item).Offset((page - 1) * item).Order("id desc").Find(&products)
	} else if len(name) == 0 && len(label_filter) != 0 {
		db.Model(&database.Product{}).Where("label LIKE ?", label_filter).Count(&total)
		db.Where("label LIKE ?", label_filter).Limit(item).Offset((page - 1) * item).Order("id desc").Find(&products)
	} else if len(name) != 0 && len(label_filter) == 0 {
		db.Model(&database.Product{}).Where("name = ?", name).Count(&total)
		db.Where("name = ?", name).Limit(item).Offset((page - 1) * item).Order("id desc").Find(&products)
	} else if len(name) != 0 && len(label_filter) != 0 {
		db.Model(&database.Product{}).Where("name = ? AND label LIKE ?", name, label_filter).Count(&total)
		db.Where("name = ? AND label LIKE ?", name, label_filter).Limit(item).Offset((page - 1) * item).Order("id desc").Find(&products)
	}

	// 构建响应数据结构
	type info struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		ProductKey string `json:"product_key"`
		CreateTime string `json:"create_time"`
		NodeTypeID int    `json:"node_type_id"`
	}
	result := []info{}
	for _, val := range products {
		data := info{
			ID:         val.ID,
			Name:       val.Name,
			ProductKey: val.ProductKey,
			CreateTime: tool.TimeDeal(val.CreateTime),
			NodeTypeID: val.NodeTypeID,
		}
		result = append(result, data)
	}

	resp_data := map[string]interface{}{
		"num_results": total,
		"data_list": result,
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
		Name         string `form:"name" json:"name"`
		Category     string `form:"category" json:"category" `
		ModelID      int    `form:"model_id" json:"model_id" `
		NodeTypeID   int    `form:"node_type_id" json:"node_type_id" `
		NetworkID    int    `form:"network_id" json:"network_id" `
		DataFormatID int    `form:"data_format_id" json:"data_format_id" `
		AuthMethodID int    `form:"auth_method_id" json:"auth_method_id" `
		Describe     string `form:"desc" json:"desc"`
	}

	data := Data{}
	if err := c.ShouldBind(&data); err != nil {
		AnalyticParameterErrResponse(c)
		return
	}

	name := data.Name
	category := data.Category
	model_id := data.ModelID
	node_type_id := data.NodeTypeID
	network_id := data.NetworkID
	data_format_id := data.DataFormatID
	auth_method_id := data.AuthMethodID
	desc := data.Describe

	if nameRes, msg := CheckProductNameQualify(name); nameRes != true {
		ErrResponse(msg, c)
		return
	}
	if categoryRes, msg := CheckProductCategoryQualify(category); categoryRes != true {
		ErrResponse(msg, c)
		return
	}
	if modelRes, msg := CheckModelIDQualify(model_id); modelRes != true {
		ErrResponse(msg, c)
		return
	}
	if nodeTypeRes, msg := CheckNodeTypeIDQualify(node_type_id); nodeTypeRes != true {
		ErrResponse(msg, c)
		return
	}
	if netWorkRes, msg := CheckNetworkIDQualify(network_id); netWorkRes != true {
		ErrResponse(msg, c)
		return
	}
	if dataFormatRes, msg := CheckDataFormatIDQualify(data_format_id); dataFormatRes != true {
		ErrResponse(msg, c)
		return
	}
	if authMethodRes, msg := CheckAuthMethodIDQualify(auth_method_id); authMethodRes != true {
		ErrResponse(msg, c)
		return
	}
	if descRes, msg := CheckProductDescQualify(desc); descRes != true {
		ErrResponse(msg, c)
		return
	}

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
	id, msg := product.SaveProduct()
	if id == 0 {
		DbErrorResponse(msg, c)
		return
	}
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
	product_id := tool.StringNumberToInTNumber(c.Query("pid"))
	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
	defer db.Close()

	if productRes, msg := CheckProductIDQualify(product_id); productRes != true {
		ErrResponse(msg, c)
		return
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

	label_map := tool.JsonStrToMap(product.Label)
	label_spl := []map[string]string{}

	for key, value := range label_map {
		data := map[string]string{
			"key":   key,
			"value": value,
		}
		label_spl = append(label_spl, data)
	}

	response := map[string]interface{}{
		"id":                product.ID,
		"name":              product.Name,
		"node_type":         product.NodeType.Name,
		"node_type_id":      product.NodeTypeID,
		"create_time":       tool.TimeDeal(product.CreateTime),
		"object_model_name": product.ObjectModel.Name,
		"data_format":       product.DataFormat.Name,
		"data_format_id":    product.DataFormatID,
		"auth_method":       product.AuthMethod.Name,
		"auth_method_id":    product.AuthMethodID,
		"status":            product.Status,
		"network_way":       product.NetworkWay.Name,
		"network_way_id":    product.NetworkWayID,
		"desc":              product.Describe,
		"concise_model_id":  product.MongodbModel.ConciseModelID,
		"intact_model_id":   product.MongodbModel.IntactModelID,
		"product_key":       product.ProductKey,
		"product_secret":    product.ProductSecret,
		"device_count":      len(product.Device),
		"label":             label_spl,
	}

	resp := gin.H{
		"status":  "Y",
		"message": "产品信息查询成功",
		"data":    response,
	}
	c.JSON(200, resp)
}

func UpdateProduct(c *gin.Context) {
	type Require struct {
		Name      string `json:"name"`
		Describe  string `json:"desc"`
		ProductID int    `json:"pid"`
	}

	var require Require
	if err := c.ShouldBind(&require); err != nil {
		AnalyticParameterErrResponse(c)
		return
	}

	name := require.Name
	desc := require.Describe
	product_id := require.ProductID

	if nameRes, msg := CheckProductNameQualify(name); nameRes != true {
		ErrResponse(msg, c)
		return
	}
	if productIDRes, msg := CheckProductIDQualify(product_id); productIDRes != true {
		ErrResponse(msg, c)
		return
	}
	if descRes, msg := CheckProductDescQualify(desc); descRes != true {
		ErrResponse(msg, c)
		return
	}

	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
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
	product_id := tool.StringNumberToInTNumber(c.Query("pid"))
	device_id := 0
	data := database.GetTopics(product_id, device_id)

	if productRes, msg := CheckProductIDQualify(product_id); productRes != true {
		ErrResponse(msg, c)
		return
	}

	resp := gin.H{
		"status":  "Y",
		"message": "Topic查询成功",
		"data":    data,
	}
	c.JSON(200, resp)
}

func AddProductTopic(c *gin.Context) {
	type Require struct {
		ProductID int    `json:"pid"`
		Name      string `json:"name"`
		Operation int    `json:"operation"`
		Describe  string `json:"desc"`
	}

	var require Require
	if err := c.ShouldBind(&require); err != nil {
		AnalyticParameterErrResponse(c)
		return
	}

	product_id := require.ProductID
	name := require.Name
	topic_name := "/%s/%s/user/" + name
	operation := require.Operation
	describe := require.Describe

	if productIDRes, msg := CheckProductIDQualify(product_id); productIDRes != true {
		ErrResponse(msg, c)
		return
	}
	if nameRes, msg := CheckProductTopicNameQualify(name); nameRes != true {
		ErrResponse(msg, c)
		return
	}
	if operationRes, msg := CheckProductTopicOperationQualify(operation); operationRes != true {
		ErrResponse(msg, c)
		return
	}
	if topDescRes, msg := CheckProductTopicDescQualify(describe); topDescRes != true {
		ErrResponse(msg, c)
		return
	}

	data := database.CustomTopic{
		ProductID:    product_id,
		PermissionID: operation,
		Detail:       topic_name,
		Describe:     describe,
	}
	id, msg := database.MysqlInsertOneData(&data)
	if id == 0 {
		DbErrorResponse(msg, c)
		return
	}

	resp := gin.H{
		"status":  "Y",
		"message": "自定义topic创建成功",
		"data":    id,
	}
	c.JSON(200, resp)
}

func UpdateProductTopic(c *gin.Context) {
	type Require struct {
		TopicID   int    `json:"tid"`
		Name      string `json:"name"`
		Operation int    `json:"operation"`
		Describe  string `json:"desc"`
	}

	var require Require
	if err := c.ShouldBind(&require); err != nil {
		AnalyticParameterErrResponse(c)
		return
	}

	topic_id := require.TopicID
	name := require.Name
	operation := require.Operation
	describe := require.Describe

	if topicNameRes, msg := CheckProductTopicNameQualify(name); topicNameRes != true {
		ErrResponse(msg, c)
		return
	}
	if operationRes, msg := CheckProductTopicOperationQualify(operation); operationRes != true {
		ErrResponse(msg, c)
		return
	}
	if descRes, msg := CheckProductTopicDescQualify(describe); descRes != true {
		ErrResponse(msg, c)
		return
	}
	if topidRes, msg := CheckProductTopicIDQualify(topic_id); topidRes != true {
		ErrResponse(msg, c)
		return
	}

	var topic database.CustomTopic
	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
	defer db.Close()

	db.First(&topic, topic_id)
	topic.PermissionID = operation
	topic.Describe = describe
	topic.Detail = "/%s/%s/user/" + name
	db.Save(&topic)

	resp := gin.H{
		"status":  "Y",
		"message": "Topic更新成功",
		"data":    nil,
	}
	c.JSON(200, resp)
}

func DeleteProductTopic(c *gin.Context) {
	type Require struct {
		TopicID int `json:"tid"`
	}

	var require Require
	if err := c.ShouldBind(&require); err != nil {
		AnalyticParameterErrResponse(c)
		return
	}
	topic_id := require.TopicID

	if topidRes, msg := CheckProductTopicIDQualify(topic_id); topidRes != true {
		ErrResponse(msg, c)
		return
	}

	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
	defer db.Close()

	db.Where("id = ?", topic_id).Delete(&database.CustomTopic{})

	resp := gin.H{
		"status":  "Y",
		"message": "Topic删除成功",
		"data":    nil,
	}
	c.JSON(200, resp)
}

func DeleteProduct(c *gin.Context) {
	type Data struct {
		ProductID int `json:"pid"`
	}

	data := Data{}
	if err := c.ShouldBind(&data); err != nil {
		AnalyticParameterErrResponse(c)
		return
	}

	product_id := data.ProductID

	if productIDRes, msg := CheckProductIDQualify(product_id); productIDRes != true {
		ErrResponse(msg, c)
		return
	}

	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
	defer db.Close()

	var devices []database.Device
	db.Where("product_id = ?", product_id).Find(&devices)

	if len(devices) > 0 {
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

	product_id := tool.StringNumberToInTNumber(c.Query("pid"))
	page := tool.StringNumberToInTNumber(c.Query("page"))
	item := tool.StringNumberToInTNumber(c.Query("item"))
	name := c.Query("name")
	remark := c.Query("remark")
	key := c.Query("key")
	value := c.Query("value")

	if product_id != 0 {
		if productIDRes, msg := CheckProductIDQualify(product_id); productIDRes != true {
			ErrResponse(msg, c)
			return
		}
	}
	if name != "" {
		if namrRes, msg := CheckDeviceNameQualify(name); namrRes != true {
			ErrResponse(msg, c)
			return
		}
	}
	if key != "" {
		if keyRes, msg := CheckDeviceLabelKeyQualify(key); keyRes != true {
			ErrResponse(msg, c)
			return
		}
	}
	if value != "" {
		if valueRes, msg := CheckDeviceLabelValueQualify(value); valueRes != true {
			ErrResponse(msg, c)
			return
		}
	}
	if remark != "" {
		if remarkRes, msg := CheckDeviceRemarkQualify(remark); remarkRes != true {
			ErrResponse(msg, c)
			return
		}
	}

	label_filter := database.DeatLabelQueryFilter(key, value)

	var total = 0
	var activate_num = 0
	var online_num = 0

	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
	defer db.Close()
	var devices []database.Device

	if product_id == 0 {
		if len(name) == 0 && len(remark) == 0 {
			if len(label_filter) == 0 {
				db.Where("activation_time != ?", "0000-00-00 00:00:00").Find(&devices)
				activate_num = len(devices)
				db.Where("online = ?", 1).Find(&devices)
				online_num = len(devices)
				db.Limit(item).Offset((page - 1) * item).Order("id desc").Preload("Status").Find(&devices)
				db.Model(&database.Device{}).Count(&total)
			} else {
				db.Where("activation_time != ? AND label LIKE ?", "0000-00-00 00:00:00", label_filter).Find(&devices)
				activate_num = len(devices)
				db.Where("online = ? AND label LIKE ?", 1, label_filter).Find(&devices)
				online_num = len(devices)
				db.Model(&database.Device{}).Where("label LIKE ?", label_filter).Count(&total)
				db.Where("label LIKE ?", label_filter).Limit(item).Offset((page - 1) * item).Order("id desc").Preload("Status").Find(&devices)
			}
		} else {
			if len(name) != 0 {
				if len(label_filter) == 0 {
					db.Where("activation_time != ? AND name = ?", "0000-00-00 00:00:00", name).Find(&devices)
					activate_num = len(devices)
					db.Where("online = ? AND name = ?", 1, name).Find(&devices)
					online_num = len(devices)
					db.Model(&database.Device{}).Where("name = ?", name).Count(&total)
					db.Where("name = ?", name).Limit(item).Offset((page - 1) * item).Order("id desc").Preload("Status").Find(&devices)
				} else {
					db.Where("activation_time != ? AND label LIKE ? AND name = ?", "0000-00-00 00:00:00", label_filter, name).Find(&devices)
					activate_num = len(devices)
					db.Where("online = ? AND label LIKE ? AND name = ?", 1, label_filter, name).Find(&devices)
					online_num = len(devices)
					db.Where("label LIKE ? AND name = ?", label_filter, name).Limit(item).Offset((page - 1) * item).Order("id desc").Preload("Status").Find(&devices)
					db.Model(&database.Device{}).Where("label LIKE ? AND name = ?", label_filter, name).Count(&total)
				}
			} else {
				if len(label_filter) == 0 {
					db.Where("activation_time != ? AND remark = ?", "0000-00-00 00:00:00", remark).Find(&devices)
					activate_num = len(devices)
					db.Where("online = ? AND remark = ?", 1, remark).Find(&devices)
					online_num = len(devices)
					db.Where("remark = ?", remark).Limit(item).Offset((page - 1) * item).Order("id desc").Preload("Status").Find(&devices)
					db.Model(&database.Device{}).Where("remark = ?", remark).Count(&total)
				} else {
					db.Where("activation_time != ? AND label LIKE ? AND remark = ?", "0000-00-00 00:00:00", label_filter, remark).Find(&devices)
					activate_num = len(devices)
					db.Where("online = ? AND label LIKE ? AND remark = ?", 1, label_filter, remark).Find(&devices)
					online_num = len(devices)
					db.Where("label LIKE ? AND remark = ?", label_filter, remark).Limit(item).Offset((page - 1) * item).Order("id desc").Preload("Status").Find(&devices)
					db.Model(&database.Device{}).Where("label LIKE ? AND remark = ?", label_filter, remark).Count(&total)
				}
			}
		}
	} else {
		if len(name) == 0 && len(remark) == 0 {
			if len(label_filter) == 0 {
				db.Where("activation_time != ? AND product_id = ?", "0000-00-00 00:00:00", product_id).Find(&devices)
				activate_num = len(devices)
				db.Where("online = ? AND product_id = ?", 1, product_id).Find(&devices)
				online_num = len(devices)
				db.Where("product_id = ?", product_id).Limit(item).Offset((page - 1) * item).Order("id desc").Preload("Status").Find(&devices)
				db.Model(&database.Device{}).Where("product_id = ?", product_id).Count(&total)
			} else {
				db.Where("activation_time != ? AND label LIKE ? AND product_id = ?", "0000-00-00 00:00:00", label_filter, product_id).Find(&devices)
				activate_num = len(devices)
				db.Where("online = ? AND label LIKE ? AND product_id = ?", 1, label_filter, product_id).Find(&devices)
				online_num = len(devices)
				db.Where("label LIKE ? AND product_id = ?", label_filter, product_id).Limit(item).Offset((page - 1) * item).Order("id desc").Preload("Status").Find(&devices)
				db.Model(&database.Device{}).Where("label LIKE ? AND product_id = ?", label_filter, product_id).Count(&total)
			}
		} else {
			if len(name) != 0 {
				if len(label_filter) == 0 {
					db.Where("activation_time != ? AND name = ? AND product_id = ?", "0000-00-00 00:00:00", name, product_id).Find(&devices)
					activate_num = len(devices)
					db.Where("online = ? AND name = ? AND product_id = ?", 1, name, product_id).Find(&devices)
					online_num = len(devices)
					db.Where("name = ? AND product_id = ?", name, product_id).Limit(item).Offset((page - 1) * item).Order("id desc").Preload("Status").Find(&devices)
					db.Model(&database.Device{}).Where("name = ? AND product_id = ?", name, product_id).Count(&total)
				} else {
					db.Where("activation_time != ? AND label LIKE ? AND name = ? AND product_id = ?", "0000-00-00 00:00:00", label_filter, name, product_id).Find(&devices)
					activate_num = len(devices)
					db.Where("online = ? AND label LIKE ? AND name = ? AND product_id = ?", 1, label_filter, name, product_id).Find(&devices)
					online_num = len(devices)
					db.Where("label LIKE ? AND name = ? AND product_id = ?", label_filter, name, product_id).Limit(item).Offset((page - 1) * item).Order("id desc").Preload("Status").Find(&devices)
					db.Model(&database.Device{}).Where("label LIKE ? AND name = ? AND product_id = ?", label_filter, name, product_id).Count(&total)
				}
			} else {
				if len(label_filter) == 0 {
					db.Where("activation_time != ? AND remark = ? AND product_id = ?", "0000-00-00 00:00:00", remark, product_id).Find(&devices)
					activate_num = len(devices)
					db.Where("online = ? AND remark = ? AND product_id = ?", 1, remark, product_id).Find(&devices)
					online_num = len(devices)
					db.Where("remark = ? AND product_id = ?", remark, product_id).Limit(item).Offset((page - 1) * item).Order("id desc").Preload("Status").Find(&devices)
					db.Model(&database.Device{}).Where("remark = ? AND product_id = ?", remark, product_id).Count(&total)
				} else {
					db.Where("activation_time != ? AND label LIKE ? AND remark = ? AND product_id = ?", "0000-00-00 00:00:00", label_filter, remark, product_id).Find(&devices)
					activate_num = len(devices)
					db.Where("online = ? AND label LIKE ? AND remark = ? AND product_id = ?", 1, label_filter, remark, product_id).Find(&devices)
					online_num = len(devices)
					db.Where("label LIKE ? AND remark = ? AND product_id = ?", label_filter, remark, product_id).Limit(item).Offset((page - 1) * item).Order("id desc").Preload("Status").Find(&devices)
					db.Model(&database.Device{}).Where("label LIKE ? AND remark = ? AND product_id = ?", label_filter, remark, product_id).Count(&total)
				}
			}
		}
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
		//db.First(&product, device.ProductID).Related(&product.NodeType, "NodeType")
		db.First(&product, device.ProductID)
		data.TheirProductName = product.Name
		//data.NodeType = product.NodeType.Name
		data.NodeTypeID = product.NodeTypeID
		//data.Status = device.Status.Name
		data.StatusID = device.StatusID
		data.LastOnLineTime = tool.TimeDeal(device.LastOnLineTime)
		responses = append(responses, data)
	}

	resp_data := map[string]interface{}{
		"device_active_count": activate_num,
		"device_online_count": online_num,
		"num_results":         total,
		"data_list":           responses,
	}

	resp := gin.H{
		"status":  "Y",
		"message": "所有设备查询成功",
		"data":    resp_data,
	}
	c.JSON(200, resp)
}

func GetSimpleProducts(c *gin.Context) {
	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
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
	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
	defer db.Close()

	type Data struct {
		ProductID int    `form:"pid" json:"pid"`
		Name      string `form:"name" json:"name"`
		Remark    string `form:"remark" json:"remark"`
	}

	data := Data{}
	if err := c.ShouldBind(&data); err != nil {
		AnalyticParameterErrResponse(c)
		return
	}

	product_id := data.ProductID
	name := data.Name
	remark := data.Remark

	if productIDRes, msg := CheckProductIDQualify(product_id); productIDRes != true {
		ErrResponse(msg, c)
		return
	}
	if nameRes, msg := CheckDeviceNameQualify(name); nameRes != true {
		ErrResponse(msg, c)
		return
	}
	if remarkRes, msg := CheckDeviceRemarkQualify(remark); remarkRes != true {
		ErrResponse(msg, c)
		return
	}

	device := database.Device{
		ProductID:   product_id,
		StatusID:    1,
		Name:        name,
		Remark:      remark,
		BatchCreate: false,
		CreateTime:  time.Now(),
		Online:      false,
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

func GetDevice(c *gin.Context) {
	device_id := tool.StringNumberToInTNumber(c.Query("did"))
	if deviceIDRes, msg := CheckDeviceIDQualify(device_id); deviceIDRes != true {
		ErrResponse(msg, c)
		return
	}

	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
	defer db.Close()
	device := database.Device{}
	db.First(&device, device_id)

	label_map := tool.JsonStrToMap(device.Label)
	label_spl := []map[string]string{}

	for key, value := range label_map {
		data := map[string]string{
			"key":   key,
			"value": value,
		}
		label_spl = append(label_spl, data)
	}

	var product database.Product
	db.First(&product, device.ProductID)
	db.Model(&product).Related(&product.NodeType, "NodeType")

	response := map[string]interface{}{
		"id":               device.ID,
		"product_id":       device.ProductID,
		"product_key":      product.ProductKey,
		"product_name":     product.Name,
		"node_type_id":     product.NetworkWayID,
		"node_type":        product.NodeType.Name,
		"status_id":        device.StatusID,
		"name":             device.Name,
		"remark":           device.Remark,
		"device_secret":    device.DeviceSecret,
		"ip":               device.IP,
		"create_time":      tool.TimeDeal(device.CreateTime),
		"activate_time":    tool.TimeDeal(device.ActivationTime),
		"last_online_time": tool.TimeDeal(device.LastOnLineTime),
		"iot_id":           device.IotID,
		"label":            label_spl,
		"batch_create":     device.BatchCreate,
	}

	resp := gin.H{
		"status":  "Y",
		"message": "设备信息查询成功",
		"data":    response,
	}
	c.JSON(200, resp)
}

func GetDeviceTopic(c *gin.Context) {
	device_id := tool.StringNumberToInTNumber(c.Query("did"))

	if deviceIDRes, msg := CheckDeviceIDQualify(device_id); deviceIDRes != true {
		ErrResponse(msg, c)
		return
	}

	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
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

func DeleteDevice(c *gin.Context) {
	type Data struct {
		DeviceIDList []int `json:"dids"`
	}

	data := Data{}
	if err := c.ShouldBind(&data); err != nil {
		AnalyticParameterErrResponse(c)
		return
	}

	device_id_list := data.DeviceIDList

	if didListRes, msg := CheckDeviceIDListQualify(device_id_list); didListRes != true {
		ErrResponse(msg, c)
		return
	}

	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
	defer db.Close()

	for index, _ := range device_id_list {
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
	device_id := tool.StringNumberToInTNumber(c.Query("did"))
	if deviceIDRes, msg := CheckDeviceIDQualify(device_id); deviceIDRes != true {
		ErrResponse(msg, c)
		return
	}

	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
	defer db.Close()

	var device database.Device
	db.First(&device, device_id)
	product_id := device.ProductID
	device_iot := device.IotID
	var product database.Product
	db.First(&product, product_id)
	product_key := product.ProductKey

	data := util.GetDeviceDesiredPropertyInfo(product_key, device_iot)
	//result := tool.DealSequentialDatabaseData(data)

	resp := gin.H{
		"status":  "Y",
		"message": "设备期望状态查询成功",
		"data":    data,
	}
	c.JSON(200, resp)
}

func GetDevicePropertyStatus(c *gin.Context) {
	device_id := tool.StringNumberToInTNumber(c.Query("did"))
	if deviceIDRes, msg := CheckDeviceIDQualify(device_id); deviceIDRes != true {
		ErrResponse(msg, c)
		return
	}

	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
	defer db.Close()

	var device database.Device
	db.First(&device, device_id)
	product_id := device.ProductID
	device_iot := device.IotID
	var product database.Product
	db.First(&product, product_id)
	product_key := product.ProductKey

	data := util.GetDevicePropertyStatusInfo(product_key, device_iot)
	//result := tool.DealSequentialDatabaseData(data)

	resp := gin.H{
		"status":  "Y",
		"message": "设备实时状态查询成功",
		"data":    data,
	}
	c.JSON(200, resp)
}

func GetDeviceHistoryStatus(c *gin.Context) {
	device_id := tool.StringNumberToInTNumber(c.Query("did"))
	property := c.Query("identifier")
	start := int64(tool.StringNumberToInTNumber(c.Query("start")))
	end := int64(tool.StringNumberToInTNumber(c.Query("end")))

	if deviceIDRes, msg := CheckDeviceIDQualify(device_id); deviceIDRes != true {
		ErrResponse(msg, c)
		return
	}
	if startRes, msg := CheckTimeStampQualify(start); startRes != true {
		ErrResponse(msg, c)
		return
	}
	if endRes, msg := CheckTimeStampQualify(end); endRes != true {
		ErrResponse(msg, c)
		return
	}

	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
	defer db.Close()

	var device database.Device
	db.First(&device, device_id)
	device_iot := device.IotID

	data := util.GetPropertyHistory(device_iot, property, start, end)

	next_time := util.GetNextPropertyHistory(device_iot, property, start, end)
	if next_time == -1 {
		next_time = 0
	}

	response := map[string]interface{}{
		"next_time": next_time,
		"data_list": data,
	}

	resp := gin.H{
		"status":  "Y",
		"message": "设备运行状态单个属性历史记录信息查询成功",
		"data":    response,
	}
	c.JSON(200, resp)
}

func GetModelFunctions(c *gin.Context) {
	model_id := tool.StringNumberToInTNumber(c.Query("id"))

	if modelIDRes, msg := CheckModelIDQualify(model_id); modelIDRes != true {
		ErrResponse(msg, c)
		return
	}

	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
	defer db.Close()

	var data []database.ModelFunction
	db.Where("model_id = ?", model_id).Find(&data)

	for index, value := range data {
		if value.DataType == "" {
			data[index].DataType = "-"
		}
	}

	resp := gin.H{
		"status":  "Y",
		"message": "标准功能定义态查询成功",
		"data":    data,
	}
	c.JSON(200, resp)
}

func GetDeviceEvent(c *gin.Context) {
	device_id := tool.StringNumberToInTNumber(c.Query("did"))
	start := int64(tool.StringNumberToInTNumber(c.Query("start")))
	end := int64(tool.StringNumberToInTNumber(c.Query("end")))
	event_type := c.Query("type")
	identifier := c.Query("identifier")

	if deviceIDRes, msg := CheckDeviceIDQualify(device_id); deviceIDRes != true {
		ErrResponse(msg, c)
		return
	}
	if startRes, msg := CheckTimeStampQualify(start); startRes != true {
		ErrResponse(msg, c)
		return
	}
	if endRes, msg := CheckTimeStampQualify(end); endRes != true {
		ErrResponse(msg, c)
		return
	}
	if event_type != "" {
		if eventTypeRes, msg := CheckEventTypeQualify(event_type); eventTypeRes != true {
			ErrResponse(msg, c)
			return
		}
	}

	if event_type == "" {
		event_type = "all"
	}

	if identifier == "" {
		identifier = "all"
	}

	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
	defer db.Close()

	var device database.Device
	db.First(&device, device_id)
	device_iot := device.IotID

	data := util.GetDeviceEventInfo(device_iot, event_type, identifier, start, end)
	next_time := util.GetNextDeviceEventInfo(device_iot, event_type, identifier, start, end)

	if next_time == -1 {
		next_time = 0
	}

	response := map[string]interface{}{
		"next_time": next_time,
		"data_list": data,
	}

	resp := gin.H{
		"status":  "Y",
		"message": "设备事件管理查询成功",
		"data":    response,
	}
	c.JSON(200, resp)
}

func GetDeviceServer(c *gin.Context) {
	device_id := tool.StringNumberToInTNumber(c.Query("did"))
	start := int64(tool.StringNumberToInTNumber(c.Query("start")))
	end := int64(tool.StringNumberToInTNumber(c.Query("end")))
	identifier := c.Query("identifier")

	if deviceIDRes, msg := CheckDeviceIDQualify(device_id); deviceIDRes != true {
		ErrResponse(msg, c)
		return
	}
	if startRes, msg := CheckTimeStampQualify(start); startRes != true {
		ErrResponse(msg, c)
		return
	}
	if endRes, msg := CheckTimeStampQualify(end); endRes != true {
		ErrResponse(msg, c)
		return
	}

	if identifier == "" {
		identifier = "all"
	}

	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
	defer db.Close()

	var device database.Device
	db.First(&device, device_id)
	device_iot := device.IotID
	data := util.GetDeviceServiceInfo(device_iot, identifier, start, end)
	next_time := util.GetNextDeviceServiceInfo(device_iot, identifier, start, end)
	if next_time == -1 {
		next_time = 0
	}

	response := map[string]interface{}{
		"next_time": next_time,
		"data_list": data,
	}

	resp := gin.H{
		"status":  "Y",
		"message": "设备服务调用查询成功",
		"data":    response,
	}
	c.JSON(200, resp)
}

func GetModelTSL(c *gin.Context) {
	product_id := tool.StringNumberToInTNumber(c.Query("pid"))

	if productIDRes, msg := CheckProductIDQualify(product_id); productIDRes != true {
		ErrResponse(msg, c)
		return
	}

	intact_mode, concise_model, msg := database.GetProductModelInfo(product_id)
	if intact_mode == nil && concise_model == nil {
		DbErrorResponse(msg, c)
		return
	}

	data := map[string]primitive.M{
		"intact_model":  intact_mode,
		"concise_model": concise_model,
	}

	resp := gin.H{
		"status":  "Y",
		"message": "物模型TSL查询成功",
		"data":    data,
	}
	c.JSON(200, resp)
}

func AutoAddDevice(c *gin.Context) {
	type Data struct {
		ProductID int `form:"pid" json:"pid"`
		Number    int `form:"num" json:"num"`
	}

	data := Data{}
	if err := c.ShouldBind(&data); err != nil {
		AnalyticParameterErrResponse(c)
		return
	}

	product_id := data.ProductID
	number := data.Number

	if productIDRes, msg := CheckProductIDQualify(product_id); productIDRes != true {
		ErrResponse(msg, c)
		return
	}
	if numberRes, msg := CheckAutoAddDeviceNumberQualify(number); numberRes != true {
		ErrResponse(msg, c)
		return
	}

	time := time.Now()
	for i := 0; i < number; i++ {
		device_name := tool.GenerateAutoDeviceName()
		device := database.Device{
			ProductID:   product_id,
			StatusID:    1,
			Name:        device_name,
			Remark:      "",
			BatchCreate: true,
			CreateTime:  time,
			Online:      false,
		}
		device.SaveDevice()
	}

	resp := gin.H{
		"status":  "Y",
		"message": "批量创建成功",
		"data":    nil,
	}
	c.JSON(200, resp)
}

func GetBatchDevices(c *gin.Context) {
	page := tool.StringNumberToInTNumber(c.Query("page"))
	item := tool.StringNumberToInTNumber(c.Query("item"))

	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
	defer db.Close()

	type Result struct {
		ProductID  int       `json:"product_id"`
		CreateTime time.Time `json:"create_time"`
		Total      int       `json:"total"`
	}

	type Response struct {
		ProductID  int    `json:"product_id"`
		Name       string `json:"name"`
		ProductKey string `json:"product_key"`
		CreateTime string `json:"create_time"`
		Total      int    `json:"total"`
	}

	var response []Response
	var result []Result

	sql := fmt.Sprintf("select product_id, create_time, COUNT(product_id) as total from device where batch_create = TRUE group by `create_time`, `product_id` ORDER BY `create_time` DESC LIMIT %d, %d;", (page-1)*item, item)
	db.Raw(sql).Scan(&result)

	for _, value := range result {
		var data Response
		data.CreateTime = tool.TimeDeal(value.CreateTime)
		data.Total = value.Total
		product_id := value.ProductID
		data.ProductID = product_id

		var product database.Product
		db.First(&product, product_id)

		data.Name = product.Name
		data.ProductKey = product.ProductKey

		response = append(response, data)
	}

	type RespData struct {
		NumResults int        `json:"num_results"`
		DataList   []Response `json:"data_list"`
	}

	total := len(result)

	var resp_data = RespData{
		NumResults: total,
		DataList:   response,
	}

	resp := gin.H{
		"status":  "Y",
		"message": "批次管理设备查询成功",
		"data":    resp_data,
	}
	c.JSON(200, resp)
}

func GetBatchDevice(c *gin.Context) {
	product_id := tool.StringNumberToInTNumber(c.Query("pid"))
	create_time := tool.StringNumberToInTNumber(c.Query("time"))
	if create_time == 0 {
		ErrResponse("时间不能为0", c)
		return
	}
	time := time.Unix(int64(create_time), 0)
	page := tool.StringNumberToInTNumber(c.Query("page"))
	item := tool.StringNumberToInTNumber(c.Query("item"))

	if productIDRes, msg := CheckProductIDQualify(product_id); productIDRes != true {
		ErrResponse(msg, c)
		return
	}
	if createTimeRes, msg := CheckTimeStampQualify(int64(create_time)); createTimeRes != true {
		ErrResponse(msg, c)
		return
	}

	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
	defer db.Close()

	var product database.Product
	db.First(&product, product_id)
	prodyct_key := product.ProductKey

	var result []database.Device
	db.Where("product_id = ? AND create_time  = ?", product_id, time).Find(&result)
	total := len(result)
	if page != 0 && item != 0 {
		db.Limit(item).Offset((page-1)*item).Order("id desc").Where("product_id = ? AND create_time  = ?", product_id, time).Find(&result)
	}

	type Response struct {
		ProductKey     string `json:"product_key"`
		DeviceName     string `json:"device_name"`
		DeviceSecret   string `json:"device_secret"`
		StatusID       int    `json:"status_id"`
		ActivationTime string `json:"activation_time"`
	}

	var response []Response
	for _, value := range result {
		var data Response
		data.DeviceName = value.Name
		data.DeviceSecret = value.DeviceSecret
		data.StatusID = value.StatusID
		data.ActivationTime = tool.TimeDeal(value.ActivationTime)
		data.ProductKey = prodyct_key
		response = append(response, data)
	}

	type RespData struct {
		NumResults int        `json:"num_results"`
		DataList   []Response `json:"data_list"`
	}

	var resp_data = RespData{
		NumResults: total,
		DataList:   response,
	}

	resp := gin.H{
		"status":  "Y",
		"message": "批次管理设备查询成功",
		"data":    resp_data,
	}
	c.JSON(200, resp)
}

func UpdateDevice(c *gin.Context) {
	type Args struct {
		Remark   string `json:"remark"`
		DeviceId int    `json:"did"`
	}

	var args Args
	if err := c.ShouldBind(&args); err != nil {
		AnalyticParameterErrResponse(c)
		return
	}

	remark := args.Remark
	device_id := args.DeviceId

	if remarkRes, msg := CheckDeviceRemarkQualify(remark); remarkRes != true {
		ErrResponse(msg, c)
		return
	}
	if deviceIDRes, msg := CheckProductIDQualify(device_id); deviceIDRes != true {
		ErrResponse(msg, c)
		return
	}

	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
	defer db.Close()

	var device database.Device
	db.First(&device, device_id)
	device.Remark = remark
	db.Save(&device)

	resp := gin.H{
		"status":  "Y",
		"message": "设备信息更新成功",
		"data":    nil,
	}
	c.JSON(200, resp)

}

func AddProductLabel(c *gin.Context) {
	type Args struct {
		ProductID int                 `json:"pid"`
		Label     []map[string]string `json:"label"`
	}

	var args Args
	if err := c.ShouldBind(&args); err != nil {
		AnalyticParameterErrResponse(c)
		return
	}

	product_id := args.ProductID
	label := args.Label
	data := database.DealLabelArgs(label)

	if productIDRes, msg := CheckProductIDQualify(product_id); productIDRes != true {
		ErrResponse(msg, c)
		return
	}
	if labelRes, msg := CheckProductLabelQualify(label); labelRes != true {
		ErrResponse(msg, c)
		return
	}

	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
	defer db.Close()

	var product database.Product
	db.First(&product, product_id)

	product.Label = data
	db.Save(&product)

	resp := gin.H{
		"status":  "Y",
		"message": "产品添加标签成功",
		"data":    nil,
	}
	c.JSON(200, resp)
}

func AddDeviceLabel(c *gin.Context) {
	type Args struct {
		DeviceID int                 `json:"did"`
		Label    []map[string]string `json:"label"`
	}

	var args Args
	if err := c.ShouldBind(&args); err != nil {
		AnalyticParameterErrResponse(c)
		return
	}

	device_id := args.DeviceID
	label := args.Label
	data := database.DealLabelArgs(label)

	if deviceIDRes, msg := CheckDeviceIDQualify(device_id); deviceIDRes != true {
		ErrResponse(msg, c)
		return
	}
	if labelRes, msg := CheckDeviceLabelQualify(label); labelRes != true {
		ErrResponse(msg, c)
		return
	}

	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
	defer db.Close()

	var device database.Device
	db.First(&device, device_id)

	device.Label = data
	db.Save(&device)

	resp := gin.H{
		"status":  "Y",
		"message": "设备添加标签成功",
		"data":    nil,
	}
	c.JSON(200, resp)
}

func AnalysisUploadCSVFile(c *gin.Context) {
	rFile, err := c.FormFile("file")
	if err != nil {
		AnalyticParameterErrResponse(c)
		return
	}

	file, err := rFile.Open()
	if err != nil {
		ErrResponse("文件错误", c)
		return
	}
	defer file.Close()
	reader := csv.NewReader(bufio.NewReader(file))

	data, _ := reader.ReadAll()

	response_data := map[string]interface{}{
		"number_count": len(data[1:]),
		"file_name":    rFile.Filename,
		"file_size":    fmt.Sprintf("%.2f", float64(rFile.Size)/float64(1024)),
	}

	resp := gin.H{
		"status":  "Y",
		"message": "文件识别成功",
		"data":    response_data,
	}
	c.JSON(200, resp)
}

func FileAddDevice(c *gin.Context) {
	rFile, _ := c.FormFile("file")
	product_id := tool.StringNumberToInTNumber(c.PostForm("pid"))

	if productIDRes, msg := CheckProductIDQualify(product_id); productIDRes != true {
		ErrResponse(msg, c)
		return
	}

	file, err := rFile.Open()
	if err != nil {
		ErrResponse("文件格式错误", c)
		return
	}
	defer file.Close()
	reader := csv.NewReader(bufio.NewReader(file))

	data, _ := reader.ReadAll()
	device_list := data[1:]

	time := time.Now()
	for _, value := range device_list {
		device_name := value[0]
		device := database.Device{
			ProductID:   product_id,
			StatusID:    1,
			Name:        device_name,
			Remark:      "",
			BatchCreate: true,
			CreateTime:  time,
			Online:      false,
		}
		device.SaveDevice()
	}

	resp := gin.H{
		"status":  "Y",
		"message": "设备批量添加成功",
		"data":    nil,
	}
	c.JSON(200, resp)
}

func GetProductFunction(c *gin.Context) {
	product_id := tool.StringNumberToInTNumber(c.Query("pid"))

	if productIDRes, msg := CheckProductIDQualify(product_id); productIDRes != true {
		ErrResponse(msg, c)
		return
	}

	property, msg := database.ProductGetPropertyFunction(product_id)
	if property == nil {
		DbErrorResponse(msg, c)
		return
	}
	services := []map[string]interface{}{}
	events := []map[string]interface{}{}

	response := map[string]interface{}{
		"property": property,
		"services": services,
		"events":   events,
	}

	resp := gin.H{
		"status":  "Y",
		"message": "产品功能定义查询成功",
		"data":    response,
	}
	c.JSON(200, resp)
}
