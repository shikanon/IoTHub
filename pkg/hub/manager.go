package hub

type DeviceManager interface {
	CreateDevice(ObjectModel) (Device, error) //创建设备
	DeleteDevice(Device) error                //删除设备
	ActiveDevice(Device) error                //激活设备
	MakeDeviceOnline(Device) error            //设备上线
	MakeDeviceOffline(Device) error           //设备下线
	DisableDevice(Device) (Device, error)     //禁用设备
	EnableDevice(string) (Device, error)      //启用设备
}
