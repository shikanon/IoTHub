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

// controller 是一个控制器
type DeviceController struct {
	runFlag bool // 结束标记
}

func (ctl *DeviceController) run() error {
	for runFlag {
		// TODO watch ShadowDevice status and spec
	}
}

func (ctl *DeviceController) Start() error {
	ctl.runFlag = true
	ctl.run()
}

func (ctl *DeviceController) Stop() error {
	ctl.runFlag = false
}

func (ctl *DeviceController) CreateDevice(model ObjectModel) (Device, error) {

}

func (ctl *DeviceController) DeleteDevice(dev Device) error {

}

func (ctl *DeviceController) ActiveDevice(dev Device) error {

}
