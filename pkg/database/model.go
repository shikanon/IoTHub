package database

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 产品
type Product struct {
	ID             int          `gorm:"primary_key" json:"id"`
	Name           string       `json:"name"`                                                                                         // 产品名称
	Category       string       `gorm:"type:enum('standard_category','custom_category');default:'standard_category'" json:"category"` // 所属品类
	ObjectModel    Model        `gorm:"foreignKey:ObjectModelID" json:"object_model"`                                                 //
	ObjectModelID  int          `json:"object_model_id"`                                                                              // 物模型id
	MongodbModel   MongodbModel `gorm:"foreignKey:MongodbModelID" json:"mongodb_model"`                                               //
	MongodbModelID int          `json:"mongodb_model_id"`                                                                             // 在mongodb数据库中model的id
	NodeType       NodeType     `gorm:"foreignKey:NodeTypeID" json:"node_type"`                                                       //
	NodeTypeID     int          `json:"node_type_id"`                                                                                 // 节点类型id
	NetworkWay     NetworkWay   `gorm:"foreignKey:NetworkWayID" json:"network_way"`                                                   //
	NetworkWayID   int          `json:"network_way_id"`                                                                               // 联网方式id
	DataFormat     DataFormat   `gorm:"foreignKey:DataFormatID" json:"data_format"`                                                   //
	DataFormatID   int          `json:"data_format_id"`                                                                               // 数据格式id
	AuthMethod     AuthMethod   `gorm:"foreignKey:AuthMethodID" json:"auth_method"`                                                   //
	AuthMethodID   int          `json:"auth_method_id"`                                                                               // 认证方式id
	Describe       string       `gorm:"type:varchar(255)" json:"desc"`                                                                // 描述
	ProductKey     string       `gorm:"unique;not null" json:"product_key"`                                                           // 产品key
	ProductSecret  string       `gorm:"unique;not null" json:"product_secret"`                                                        // 产品密钥
	Status         string       `gorm:"default:'开发中';type:varchar(255)" json:"status"`                                                // 产品状态
	Device         []Device     `gorm:"foreignKey:ProductID"`                                                                         // 	//
	Label          string       `gorm:"type:varchar(255);default:''" json:"label"`                                                    // 产品标签。键值对
	CreateTime     time.Time    `json:"create_time"`                                                                                  // 创建时间
}

func (p *Product) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("create_time", time.Now())
	return nil
}

// 设备
type Device struct {
	ID             int          `gorm:"primary_key" json:"id"`
	ProductID      int          `json:"product_id"`                           // 对应的产品的ID
	Status         DeviceStatus `gorm:"foreignKey:StatusID" json:"status"`    //
	StatusID       int          `json:"status_id"`                            // 状态id
	Name           string       `json:"name"`                                 // 设备名称
	Remark         string       `json:"remark"`                               // 备注
	DeviceSecret   string       `gorm:"unique;not null" json:"device_secret"` // 设备密钥
	IP             string       `json:"ip"`                                   // 设备所在ip
	CreateTime     time.Time    `json:"create_time"`                          // 创建时间
	ActivationTime time.Time    `json:"activate_time"`                        // 激活时间
	LastOnLineTime time.Time    `json:"last_online_time"`                     // 最后在线时间
	IotID          string       `gorm:"unique;not null" json:"iot_id"`        // 设备唯一id
	Label          string       `json:"label"`                                // 设备标签
	BatchCreate    bool         `json:"batch_create"`                         // 是否批量创建
	Online         bool         `json:"online"`                               // 是否在线
	//ShadowDevice   ShadowDevice `gorm:"foreignKey:ShadowDeviceID"`            //
	//ShadowDeviceID int          `json:"shadow_device_id"`                     // 影子设备id
}

//func (p *Device) BeforeCreate(scope *gorm.Scope) error {
//	scope.SetColumn("create_time", time.Now())
//	return nil
//}

// 影子设备
//type ShadowDevice struct {
//	gorm.Model
//	DeviceID  int `json:"device_id"`  // 对应设备id
//	MongodbID int `json:"mongodb_id"` // 在mongodb数据库中的id
//}

// 物模型
type Model struct {
	ID          int            `gorm:"primary_key" json:"id"`
	MongodbID   int            `json:"mongodb_id"`                              // 在mongodb数据库中存储id
	Name        string         `json:"name"`                                    // 名称
	Scene       string         `json:"scene"`                                   // 所属场景
	Territory   ModelTerritory `gorm:"foreignKey:TerritoryID" json:"territory"` //
	TerritoryID int            `json:"territory_id"`                            // 所属领域id
	Function    []ModelFunction
}

// 物模型标准功能定义
type ModelFunction struct {
	ID           int    `gorm:"primary_key" json:"id"` // id
	ModelID      int    `json:"model_id"`              // 物模型id
	FunctionType string `json:"function_type"`         // 功能类型
	FunctionName string `json:"function_name"`         // 功能名称
	Tag          string `json:"tag"`                   // 标示符
	DataType     string `json:"data_type"`             // 数据类型
	Required     bool   `json:"required"`              // 是否必选
}

// 物模型领域
type ModelTerritory struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Name string `json:"name"` // 名称
}

// 节点
type NodeType struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Name string `json:"name"` // 名称
}

// 联网方式
type NetworkWay struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Name string `json:"name"` // 名称
}

// 数据格式
type DataFormat struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Name string `json:"name"` // 名称
}

// 认证方式
type AuthMethod struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Name string `json:"name"` // 名称
}

// 设备状态
type DeviceStatus struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Name string `json:"name"` // 名称
}

// 自定义topic类
type CustomTopic struct {
	ID           int             `gorm:"primary_key" json:"id"`
	ProductID    int             `json:"product_id"`                                // 对应产品id
	Permission   TopicPermission `gorm:"foreignKey:PermissionID" json:"permission"` //
	PermissionID int             `json:"permission_id"`                             // 权限id
	Detail       string          `json:"detail"`                                    // 详情
	Describe     string          `json:"desc"`                                      // 描述
}

// topic权限表
type TopicPermission struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}

// mongodb中的物模型
type MongodbModel struct {
	ID             int    `gorm:"primary_key" json:"id"`
	ConciseModelID string `json:"concise_model_id"` // 简洁模型id
	IntactModelID  string `json:"intact_model_id"`  // 完整模型id
}
