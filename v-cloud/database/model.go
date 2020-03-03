package database

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 工作空间
type WorkSpace struct {
	ID            int        `gorm:"primary_key" json:"id"`
	Name          string     `json:"name"`                                       // 名称
	SpaceID       string     `gorm:"unique;not null" json:"space_id"`            // 空间ID
	StatusID      int        `json:"status_id"`                                  // 状态ID
	AccessType    AccessType `gorm:"foreignKey:AccessTypeID" json:"access_type"` //
	AccessTypeID  int        `json:"access_type_id"`                             // 接入类型ID
	AccessAddress string     `json:"access_address"`                             // 接入地址
	DeviceCount   int        `json:"device_count"`                               // 接入设备数
	CreateTime    time.Time  `json:"create_time"`                                // 创建时间
}

func (p *WorkSpace) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("create_time", time.Now())
	return nil
}

// 摄像头管理
type CameraManagement struct {
	ID              int       `gorm:"primary_key" json:"id"`
	Sin             string    `json:"sin"`               // sin码
	WorkSpaceID     int       `json:"work_space_id"`     //
	AccessAddress   string    `json:"access_address"`    // 接入地址
	StatusID        int       `json:"status_id"`         // 状态ID
	CreateTime      time.Time `json:"create_time"`       // 创建时间
	LastConnectTime time.Time `json:"last_connect_time"` // 最后连接时间
}

func (p *CameraManagement) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("create_time", time.Now())
	return nil
}

// 工作空间接入类型
type AccessType struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}
