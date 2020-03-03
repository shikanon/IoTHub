package database

func (ws *WorkSpace) SaveWorkSpace() (id int, msg string) {

	// TODO 假数据，缺少相关功能代码
	ws.SpaceID = "aaaaaaaaaaaaaaaaa"
	ws.AccessAddress = "192.168.0.1:1935"

	id, msg = MysqlInsertOneData(ws)
	if id == 0 {
		return 0, msg
	}
	return id, ""
}

func (ca *CameraManagement) SaveCamera() (id int, msg string) {

	// TODO 假数据，缺少相关功能代码
	ca.AccessAddress = "192.168.0.1:1935/masdffdaff"

	id, msg = MysqlInsertOneData(ca)
	if id == 0 {
		return 0, msg
	}
	return id, ""
}
