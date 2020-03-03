package api

import (
	"github.com/gin-gonic/gin"
	"github.com/shikanon/IoTOrbHub/pkg/tool"
	"github.com/shikanon/IoTOrbHub/v-cloud/database"
)

func Home(c *gin.Context) {
	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
	defer db.Close()
	//db.AutoMigrate(&database.CameraManagement{}, &database.AccessType{}, &database.WorkSpace{})

	//var access database.AccessType
	//access.Name = "HTTP-FLV"
	//id, msg := database.MysqlInsertOneData(&access)
	//fmt.Println(id)
	//fmt.Println(msg)

	resp := gin.H{
		"status":  "Y",
		"message": "asasasas",
		"data":    nil,
	}
	c.JSON(200, resp)
}

// *********************************************************************************************************************
func AddWorkSpace(c *gin.Context) {
	type Data struct {
		Name         string `form:"name" json:"name"`
		AccessTypeID int    `form:"access_type_id" json:"access_type_id"`
	}

	data := Data{}
	if err := c.ShouldBind(&data); err != nil {
		AnalyticParameterErrResponse(c)
		return
	}

	name := data.Name
	access_type_id := data.AccessTypeID

	ws := database.WorkSpace{
		Name:         name,
		StatusID:     0,
		AccessTypeID: access_type_id,
		DeviceCount:  0,
	}

	id, msg := ws.SaveWorkSpace()
	if id == 0 {
		DbErrorResponse(msg, c)
		return
	}

	resp := gin.H{
		"status":  "Y",
		"message": "工作空间创建成功",
		"data":    id,
	}
	c.JSON(200, resp)
}

func UpdateWorkSpace(c *gin.Context) {
	type Data struct {
		Name         string `form:"name" json:"name"`
		AccessTypeID int    `form:"access_type_id" json:"access_type_id"`
		ID           int    `form:"id" json:"id"`
	}

	data := Data{}
	if err := c.ShouldBind(&data); err != nil {
		AnalyticParameterErrResponse(c)
		return
	}

	id := data.ID
	name := data.Name
	access_type_id := data.AccessTypeID

	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
	defer db.Close()

	var ws database.WorkSpace
	db.First(&ws, id)
	ws.Name = name
	ws.AccessTypeID = access_type_id
	db.Save(&ws)

	resp := gin.H{
		"status":  "Y",
		"message": "工作空间修改成功",
		"data":    nil,
	}
	c.JSON(200, resp)
}

func GetWorkSpace(c *gin.Context) {
	id := tool.StringNumberToInTNumber(c.Query("id"))

	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
	defer db.Close()

	var ws database.WorkSpace
	db.First(&ws, id)

	response := map[string]interface{}{
		"name":           ws.Name,
		"access_type_id": ws.AccessTypeID,
		"access_address": ws.AccessAddress,
	}

	resp := gin.H{
		"status":  "Y",
		"message": "工作空间查询成功",
		"data":    response,
	}
	c.JSON(200, resp)
}

func GetAllAccessTypes(c *gin.Context) {
	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
	defer db.Close()

	var access_type []database.AccessType
	db.Find(&access_type)

	resp := gin.H{
		"status":  "Y",
		"message": "所有推流类型查询成功",
		"data":    access_type,
	}
	c.JSON(200, resp)
}

func GetAllWorkSpaces(c *gin.Context) {
	page := tool.StringNumberToInTNumber(c.Query("page"))
	item := tool.StringNumberToInTNumber(c.Query("item"))
	name := c.Query("name")

	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
	defer db.Close()

	var total = 0
	var wss []database.WorkSpace
	if len(name) != 0 {
		db.Model(&database.WorkSpace{}).Where("name = ?", name).Count(&total)
		db.Where("name = ?", name).Limit(item).Offset((page - 1) * item).Order("id desc").Find(&wss)
	} else {
		db.Model(&database.WorkSpace{}).Count(&total)
		db.Limit(item).Offset((page - 1) * item).Order("id desc").Find(&wss)
	}

	type info struct {
		ID            int    `json:"work_space_id"`
		Name          string `json:"name"`
		StatusID      int    `json:"status_id"`
		AccessTypeID  int    `json:"access_type_id"`
		AccessAddress string `json:"access_address"`
		DeviceCount   int    `json:"device_count"`
	}

	result := []info{}
	for _, val := range wss {
		data := info{
			ID:            val.ID,
			Name:          val.Name,
			StatusID:      val.StatusID,
			AccessTypeID:  val.AccessTypeID,
			AccessAddress: val.AccessAddress,
			DeviceCount:   val.DeviceCount,
		}
		result = append(result, data)
	}

	resp_data := map[string]interface{}{
		"count":     total,
		"data_list": result,
	}

	resp := gin.H{
		"status":  "Y",
		"message": "产品首页信息查询成功",
		"data":    resp_data,
	}
	c.JSON(200, resp)
}

func UpdateWorkSpaceStatus(c *gin.Context) {
	type Data struct {
		ID int `form:"id" json:"id"`
	}

	data := Data{}
	if err := c.ShouldBind(&data); err != nil {
		AnalyticParameterErrResponse(c)
		return
	}

	id := data.ID

	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
	defer db.Close()

	var ws database.WorkSpace
	db.First(&ws, id)

	if id == 0 {
		ws.StatusID = 1
	} else {
		ws.StatusID = 0
	}

	db.Save(&ws)

	resp := gin.H{
		"status":  "Y",
		"message": "工作空间状态修改成功",
		"data":    nil,
	}
	c.JSON(200, resp)
}

func DeleteWorkSpace(c *gin.Context) {
	type Data struct {
		ID int `form:"id" json:"id"`
	}

	data := Data{}
	if err := c.ShouldBind(&data); err != nil {
		AnalyticParameterErrResponse(c)
		return
	}

	id := data.ID

	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
	defer db.Close()

	db.Where("id = ?", id).Delete(&database.WorkSpace{})

	resp := gin.H{
		"status":  "Y",
		"message": "工作空间删除成功",
		"data":    nil,
	}
	c.JSON(200, resp)
}

func AddCamera(c *gin.Context) {
	type Data struct {
		WS  int    `form:"ws" json:"ws"`
		Sin string `form:"sin" json:"sin"`
	}

	data := Data{}
	if err := c.ShouldBind(&data); err != nil {
		AnalyticParameterErrResponse(c)
		return
	}

	ws := data.WS
	sin := data.Sin

	camera := database.CameraManagement{
		Sin:         sin,
		WorkSpaceID: ws,
	}

	id, msg := camera.SaveCamera()
	if id == 0 {
		DbErrorResponse(msg, c)
		return
	}
	resp := gin.H{
		"status":  "Y",
		"message": "设备添加成功",
		"data":    id,
	}
	c.JSON(200, resp)
}

func UpdateCameraStatus(c *gin.Context) {
	type Data struct {
		ID int `form:"id" json:"id"`
	}

	data := Data{}
	if err := c.ShouldBind(&data); err != nil {
		AnalyticParameterErrResponse(c)
		return
	}

	id := data.ID

	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
	defer db.Close()

	var ca database.CameraManagement
	db.First(&ca, id)

	if id == 0 {
		ca.StatusID = 1
	} else {
		ca.StatusID = 0
	}

	db.Save(&ca)

	resp := gin.H{
		"status":  "Y",
		"message": "设备状态状态修改成功",
		"data":    nil,
	}
	c.JSON(200, resp)
}

func DeleteCamera(c *gin.Context) {
	type Data struct {
		ID int `form:"id" json:"id"`
	}

	data := Data{}
	if err := c.ShouldBind(&data); err != nil {
		AnalyticParameterErrResponse(c)
		return
	}

	id := data.ID

	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
	defer db.Close()

	db.Where("id = ?", id).Delete(&database.CameraManagement{})

	resp := gin.H{
		"status":  "Y",
		"message": "摄像头删除成功",
		"data":    nil,
	}
	c.JSON(200, resp)
}

func GetAllCameras(c *gin.Context) {
	page := tool.StringNumberToInTNumber(c.Query("page"))
	item := tool.StringNumberToInTNumber(c.Query("item"))
	ws := tool.StringNumberToInTNumber(c.Query("ws"))
	sin := c.Query("sin")

	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
	defer db.Close()

	var total = 0
	var cas []database.CameraManagement

	if ws == 0 {
		if len(sin) == 0 {
			db.Model(&database.WorkSpace{}).Count(&total)
			db.Limit(item).Offset((page - 1) * item).Order("id desc").Find(&cas)
		} else {
			db.Model(&database.WorkSpace{}).Where("sin = ?", sin).Count(&total)
			db.Where("sin = ?", sin).Limit(item).Offset((page - 1) * item).Order("id desc").Find(&cas)
		}
	} else {
		if len(sin) == 0 {
			db.Model(&database.WorkSpace{}).Where("id = ?", ws).Count(&total)
			db.Where("id = ?", ws).Limit(item).Offset((page - 1) * item).Order("id desc").Find(&cas)
		} else {
			db.Model(&database.WorkSpace{}).Where("id = ? AND sin = ?", ws, sin).Count(&total)
			db.Where("id = ? AND sin = ?", ws, sin).Limit(item).Offset((page - 1) * item).Order("id desc").Find(&cas)
		}
	}

	type info struct {
		Sin             string `json:"sin"`
		SpaceName       string `json:"space_name"`
		AccessAddress   string `json:"access_address"`
		StatusID        int    `json:"status_id"`
		LastConnectTime string `json:"last_connect_time"`
		CameraID        int    `json:"camera_id"`
	}

	result := []info{}
	for _, val := range cas {
		var ws database.WorkSpace
		ws_id := val.WorkSpaceID
		db.First(&ws, ws_id)
		ws_name := ws.Name
		data := info{
			Sin:             val.Sin,
			SpaceName:       ws_name,
			AccessAddress:   val.AccessAddress,
			StatusID:        val.StatusID,
			LastConnectTime: tool.TimeDeal(val.LastConnectTime),
			CameraID:        val.ID,
		}
		result = append(result, data)
	}

	resp_data := map[string]interface{}{
		"count":     total,
		"data_list": result,
	}

	resp := gin.H{
		"status":  "Y",
		"message": "产品首页信息查询成功",
		"data":    resp_data,
	}
	c.JSON(200, resp)
}
