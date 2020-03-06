package apiserver

import (
	"github.com/gin-gonic/gin"
	iotApi "github.com/shikanon/IoTOrbHub/pkg/apiserver/api"
	vCloudApi "github.com/shikanon/IoTOrbHub/v-cloud/api"
	"net/http"
)

// 跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func ApiRegister() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	//router.Use(log.LoggerToFile())
	router.Use(Cors())
	router.Use(gin.Recovery())

	iot := router.Group("/iot/api/v1")
	{
		iot.GET("/products", iotApi.GetProducts)              // 产品首页
		iot.GET("/models", iotApi.GetProductModels)           // 产品-获取物模型
		iot.GET("/nodetypes", iotApi.GetProductNodeTypes)     // 产品-获取节点类型
		iot.GET("/networkways", iotApi.GetProductNetworkWays) // 产品-获取联网方式
		iot.GET("/dataformats", iotApi.GetProductDataFormats) // 产品-获取数据格式
		iot.GET("/authmethods", iotApi.GetProductAuthMethods) // 产品-获取认证方式
		iot.GET("/simpleproducts", iotApi.GetSimpleProducts)  // 获取所有产品的id和名称
		iot.POST("/product", iotApi.AddProduct)               // 产品-创建产品
		iot.GET("/product", iotApi.GetProduct)                // 产品-查看
		iot.PUT("/product", iotApi.UpdateProduct)             // 产品-查看-编辑(名称)
		iot.POST("/plabel", iotApi.AddProductLabel)           // 产品-添加标签
		iot.GET("/ptopics", iotApi.GetProductTopic)           // 产品-查看-topic类
		iot.POST("/ptopic", iotApi.AddProductTopic)           // 产品-查看-topic类，自定义，定义topic类
		iot.PUT("/ptopic", iotApi.UpdateProductTopic)         // 产品-查看-topic类，自定义，编辑topic类
		iot.DELETE("/ptopic", iotApi.DeleteProductTopic)      // 产品-查看-topic类，自定义，删除topic类
		iot.GET("/functions", iotApi.GetProductFunction)      // 产品-产看-功能定义
		iot.GET("/tsl", iotApi.GetModelTSL)                   // 产品-查看-功能定义-物模型TSL
		iot.DELETE("/product", iotApi.DeleteProduct)          // 产品-删除
		iot.GET("/devices", iotApi.GetDevices)                // 列出所有设备(设备首页 / 产品-管理设备 / 产品-查看-前往管理)
		iot.POST("/device", iotApi.AddDevice)                 // 设备-创建设备
		iot.POST("/adevice", iotApi.AutoAddDevice)            // 设备-批量添加-自动生成
		iot.POST("/upload", iotApi.AnalysisUploadCSVFile)     // 设备-批量创建-上传文件
		iot.POST("/fdevice", iotApi.FileAddDevice)            // 设备-批量添加-文件确认
		iot.GET("/batchdevices", iotApi.GetBatchDevices)      // 设备-批次管理
		iot.GET("/batchdevice", iotApi.GetBatchDevice)        // 设备-批次管理-详情
		iot.GET("/device", iotApi.GetDevice)                  // 设备-查看
		iot.PUT("/device", iotApi.UpdateDevice)               // 设备-编辑
		iot.POST("/dlabel", iotApi.AddDeviceLabel)            // 设备-添加标签
		iot.GET("/dtopics", iotApi.GetDeviceTopic)            // 设备-查看-查看topic类
		iot.GET("/desstatus", iotApi.GetDeviceDesireStatus)   // 设备-获取设备期望状态
		iot.GET("/prostatus", iotApi.GetDevicePropertyStatus) // 设备-获取设备实时状态
		iot.GET("/hisstatus", iotApi.GetDeviceHistoryStatus)  // 设备-查看设备运行状态单个属性历史记录信息
		iot.GET("/event", iotApi.GetDeviceEvent)              // 设备-事件管理
		iot.GET("/server", iotApi.GetDeviceServer)            // 设备-服务调用
		iot.GET("/modelfuncs", iotApi.GetModelFunctions)      // 产品-获取物模型标准功能定义
		iot.DELETE("/device", iotApi.DeleteDevice)            // 设备-删除
	}

	vCloud := router.Group("/v-cloud/api/v1")
	{
		vCloud.GET("/", vCloudApi.Home)

		vCloud.POST("/workspace", vCloudApi.AddWorkSpace)        // 创建工作空间
		vCloud.PUT("/workspace", vCloudApi.UpdateWorkSpace)      // 修改工作空间
		vCloud.GET("/workspace", vCloudApi.GetWorkSpace)         // 获取工作空间
		vCloud.GET("/ast", vCloudApi.GetAllAccessTypes)          // 获取所有推流类型
		vCloud.GET("/workspaces", vCloudApi.GetAllWorkSpaces)    // 工作空间首页
		vCloud.PUT("/wsstatus", vCloudApi.UpdateWorkSpaceStatus) // 改变工作空间状态
		vCloud.DELETE("/workspace", vCloudApi.DeleteWorkSpace)   // 删除工作空间
		vCloud.POST("/camera", vCloudApi.AddCamera)              // 添加摄像头
		vCloud.DELETE("/camera", vCloudApi.DeleteCamera)         // 删除摄像头
		vCloud.PUT("/cstatus", vCloudApi.UpdateCameraStatus)     // 改变设备状态
		vCloud.GET("/cameras", vCloudApi.GetAllCameras)          // 摄像头管理首页
		vCloud.GET("/wss", vCloudApi.GetAllWorkSpaceIdAndName)   // 获取所有工作空间的id和name
		vCloud.GET("/cas", vCloudApi.GetAllCameraIdAndName)      // 获取所有摄像头的id、sin、状态
	}

	router.Run("0.0.0.0:9898")
}
