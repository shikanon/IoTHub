package api

import (
	"github.com/gin-gonic/gin"
	"github.com/shikanon/IoTOrbHub/v-cloud/database"
)

func Home(c *gin.Context) {
	db, msg := database.DbConn()
	if db == nil {
		DbErrorResponse(msg, c)
		return
	}
	defer db.Close()
	//db.AutoMigrate(&database.CameraManagement{}, &database.AccessType{}, &database.WorkSpace{})

	//var access database.AccessType
	//access.Name = "HTTP-FLV"
	//id, msg := database.MysqlInsertOneData(&access)
	//fmt.Println(id)
	//fmt.Println(msg)

	resp := gin.H{
		"status":  "Y",
		"message": "asasasas",
		"data":    nil,
	}
	c.JSON(200, resp)
}


