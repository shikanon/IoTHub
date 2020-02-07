package apiserver

import (
	"github.com/gin-gonic/gin"
	"github.com/shikanon/IoTOrbHub/pkg/apiserver/api"
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
	router.Use(Cors())
	router.Use(gin.Recovery())
	v1 := router.Group("/iot/api/v1")
	{
		v1.GET("/models", api.GetProductModels)           // 产品-获取物模型
		v1.GET("/nodetypes", api.GetProductNodeTypes)     // 产品-获取节点类型
		v1.GET("/networkways", api.GetProductNetworkWays) // 产品-获取联网方式
		v1.GET("/dataformats", api.GetProductDataFormats) // 产品-获取数据格式
		v1.GET("/authmethods", api.GetProductAuthMethods) // 产品-获取认证方式
		v1.GET("/simpleproducts", api.GetSimpleProducts)  // 获取所有产品的id和名称
		v1.GET("/products", api.GetProducts)              // 产品首页
		v1.POST("/product", api.AddProduct)               // 产品-创建产品
		v1.GET("/product", api.GetProduct)                // 产品-查看
		v1.PUT("/product", api.UpdateProduct)             // 产品-查看-编辑(名称、描述、标签)
		v1.GET("/ptopics", api.GetProductTopic)           // 产品-查看-topic类
		v1.POST("/ptopic", api.AddProductTopic)           // 产品-查看-topic类，自定义，定义topic类
		v1.PUT("/ptopic", api.UpdateProductTopic)         // 产品-查看-topic类，自定义，编辑topic类
		v1.DELETE("/ptopic", api.DeleteProductTopic)      // 产品-查看-topic类，自定义，删除topic类
		v1.GET("/tsl", api.GetModelTSL)                   // 产品-查看-功能定义-物模型TSL
		v1.DELETE("/product", api.DeleteProduct)          // 产品-删除
		v1.GET("/devices", api.GetDevices)                // 列出所有设备(设备首页 / 产品-管理设备 / 产品-查看-前往管理)
		v1.POST("/device", api.AddDevice)                 // 设备-创建设备
		v1.POST("/adevice", api.AutoAddDevice)            // 批量添加设备-自动生成
		v1.GET("/batchdevices", api.GetBatchDevices)      // 设备-批次管理
		v1.GET("/batchdevice", api.GetBatchDevice)        // 设备-批次管理-详情
		v1.GET("/device", api.GetDevice)                  // 设备-查看
		v1.GET("/dtopics", api.GetDeviceTopic)            // 设备-查看-查看topic类
		v1.DELETE("/device", api.DeleteDevice)            // 设备-删除
		v1.GET("/desstatus", api.GetDeviceDesireStatus)   // 获取设备期望状态
		v1.GET("/prostatus", api.GetDevicePropertyStatus) // 获取设备实时状态
		v1.GET("/hisstatus", api.GetDeviceHistoryStatus)  // 查看设备运行状态单个属性历史记录信息  TODO
		v1.GET("/event", api.GetDeviceEvent)              // 查看设备事件管理 TODO
		v1.GET("/server", api.GetDeviceServer)            // 查看设备服务调用 TODO
		v1.GET("/modelfuncs", api.GetModelFunctions)      // 获取物模型标准功能定义
		v1.PUT("/device", api.UpdateDevice)               // 设备-编辑 TODO
		v1.POST("/plabel", api.AddProductLabel)           // 产品-添加标签 TODO
		v1.POST("/dlabel", api.AddDeviceLabel)            // 设备-添加标签 TODO

		v1.GET("/", Cors(), api.Home)
	}
	router.Run("0.0.0.0:9898")
}
