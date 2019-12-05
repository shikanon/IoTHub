package hub

type Device interface {
	GetAttribute(key string) (string,error) //获取属性
	LoadModel(ObjectModel) error//装载物模型
}

type RealDevice sturct {
	ID string //设备ID
	attribute map[string]interface{}
	event	map[string]Notifier
}

func (dev *RealDevice) GetAttribute(key string) (string,error){

}

type ShadowDevice struct {
	Ontology RealDevice //真实设备
	Version string      //版本
	Kind    string      //类型
	Status  interface{} //状态
	Spec    interface{} //期望状态
}
