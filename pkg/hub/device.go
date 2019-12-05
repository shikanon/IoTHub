package hub

type Device interface {
	Active()
	GoOnline()
	GoOffline()
}


type RealDevice struct {
	ID string //设备ID
}

type ShadowDevice struct {
	Ontology RealDevice //真实设备
	Version string      //版本
	Kind    string      //类型
	Status  interface{} //状态
	Spec    interface{} //期望状态
}
