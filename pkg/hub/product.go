package hub

import (
	"github.com/shikanon/IoTOrbHub/pkg/auth"
)

type Product interface {
	AddDevice(name, remarks string) (Device, error)
}

type IoTProduct struct {
	Key         string            // 产品key
	Name        string            // 产品名称
	NetworkMode string            // 联网方式
	AuthMethod  auth.DeviceAuthen // 认证方式
	NodeType    string            // 节点类型
	DataFormat  string            // 数据格式,默认支持Alink json
	Describe    string            // 产品描述
}

func NewIoTProduct(ObjectModel) (IoTProduct, error) {}

func (prod *IoTProduct) AddDevice(name, remarks string) (Device, error) {}
