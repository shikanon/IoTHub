package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shikanon/IoTOrbHub/config"
	"github.com/shikanon/IoTOrbHub/pkg/database"
	"unicode"
)

func ErrResponse(msg string, c *gin.Context) {
	resp := gin.H{
		"status":  "N",
		"message": msg,
		"data":    nil,
	}
	c.JSON(400, resp)
}

func CheckProductNameQualify(name string) (result bool, msg string) {
	var count int
	for _, value := range name {
		if unicode.Is(unicode.Han, value) {
			count += 2
		} else {
			count += 1
		}
	}

	if count < 4 {
		return false, "产品名称短于4"
	}
	if count > 30 {
		return false, "产品名称长于30"
	}

	return true, ""
}

func CheckProductLabelKeyQualify(key string) (result bool, msg string) {
	if len(key) == 0 {
		return false, "标签key为空"
	}
	if len(key) > 30 {
		return false, "标签key长度超过30"
	}
	for _, v := range key {
		if unicode.Is(unicode.Han, v) {
			return false, "标签key不支持中文"
		}
	}
	return true, ""
}

func CheckProductLabelValueQualify(value string) (result bool, msg string) {
	var count int
	for _, value := range value {
		if unicode.Is(unicode.Han, value) {
			count += 2
		} else {
			count += 1
		}
	}
	if count > 128 {
		return false, "标签value长度超过128"
	}
	return true, ""
}

func CheckProductCategoryQualify(category string) (result bool, msg string) {
	if category != "standard_category" && category != "custom_category" {
		return false, "不支持的品类(category)类型"
	}
	return true, ""
}

func CheckModelIDQualify(modelID int) (result bool, msg string) {
	maxNumber := config.GeneralConfig.ModelNumber
	if modelID > maxNumber {
		return false, fmt.Sprintf("model_id超出最大范围。目前最大为%d", maxNumber)
	}
	return true, ""
}

func CheckNodeTypeIDQualify(nodeTypeID int) (result bool, msg string) {
	maxNumber := config.GeneralConfig.NodeTypeNumber
	if nodeTypeID > maxNumber {
		return false, fmt.Sprintf("node_type_id超出最大范围. 目前最大为%d", maxNumber)
	}
	return true, ""
}

func CheckNetworkIDQualify(netWorkID int) (result bool, msg string) {
	maxNumber := config.GeneralConfig.NetWorkNumber
	if netWorkID > maxNumber {
		return false, fmt.Sprintf("network_id超出最大范围。目前最大为%d", maxNumber)
	}
	return true, ""
}

func CheckDataFormatIDQualify(dataFormatID int) (result bool, msg string) {
	maxNumber := config.GeneralConfig.DataFormatNumber
	if dataFormatID > maxNumber {
		return false, fmt.Sprintf("data_format_id超出最大范围。目前最大为%d", maxNumber)
	}
	return true, ""
}

func CheckAuthMethodIDQualify(authMethodID int) (result bool, msg string) {
	maxNumber := config.GeneralConfig.AuthMethodNumber
	if authMethodID > maxNumber {
		return false, fmt.Sprintf("auth_method_id超出最大范围。目前最大为%d", maxNumber)
	}
	return true, ""
}

func CheckProductIDQualify(productID int) (result bool, msg string) {
	ids := database.GetAllProductID()
	for _, value := range ids {
		if productID == value {
			return true, ""
		}
	}
	return false, "产品ID不是有效参数"
}

func CheckProductDescQualify(desc string) (result bool, msg string) {
	var count int
	for range desc {
		count++
	}

	if count > 100 {
		return false, "产品描述超长"
	}

	return true, ""
}

func CheckProductLabelQualify(label []map[string]string) (result bool, msg string) {
	for _, v := range label {
		key := v["key"]
		if keyRes, msg := CheckProductLabelKeyQualify(key); keyRes != true {
			return false, msg
		}
		value := v["value"]
		if valueRes, msg := CheckProductLabelValueQualify(value); valueRes != true {
			return false, msg
		}
	}
	return true, ""
}

func CheckProductTopicOperationQualify(operation int) (result bool, msg string) {
	maxNumber := config.GeneralConfig.TopicOperationNumber
	if operation > maxNumber {
		return false, fmt.Sprintf("operation超出最大范围。目前最大为%d", maxNumber)
	}
	return true, ""
}

func CheckProductTopicDescQualify(desc string) (result bool, msg string) {
	var count int
	for range desc {
		count++
	}

	if count > 100 {
		return false, "Topic描述超长"
	}

	return true, ""
}

func CheckProductTopicNameQualify(name string) (result bool, msg string) {
	for _, v := range name {
		if unicode.Is(unicode.Han, v) {
			return false, "Topic类名不支持中文"
		}
	}
	if len(name) > 64 {
		return false, "Topic类名超长"
	}
	if len(name) == 0 {
		return false, "Topic类名不能为空"
	}
	return true, ""
}

func CheckProductTopicIDQuality(topicID int) (result bool, msg string) {
	ids := database.GetAllProductTopicID()
	for _, value := range ids {
		if topicID == value {
			return true, ""
		}
	}
	return false, "TopicID不是有效参数"
}

func CheckDeviceName(name string) (result bool, msg string) {
	for _, v := range name {
		if unicode.Is(unicode.Han, v) {
			return false, "设备名称不支持中文"
		}
	}
	if len(name) < 4 {
		return false, "设备名称过短"
	}
	if len(name) > 32 {
		return false, "设备名称过长"
	}
	return true, ""
}

func CheckDeviceLabelKeyQuality(key string) (result bool, msg string) {
	if len(key) == 0 {
		return false, "标签key为空"
	}
	if len(key) > 30 {
		return false, "标签key长度超过30"
	}
	for _, v := range key {
		if unicode.Is(unicode.Han, v) {
			return false, "标签key不支持中文"
		}
	}
	return true, ""
}

func CheckDeviceLabelValueQuality(value string) (result bool, msg string) {
	var count int
	for _, value := range value {
		if unicode.Is(unicode.Han, value) {
			count += 2
		} else {
			count += 1
		}
	}
	if count > 128 {
		return false, "标签value长度超过128"
	}
	return true, ""
}

func CheckDeviceRemarkQuality(remark string) (result bool, msg string) {
	var count int
	for _, value := range remark {
		if unicode.Is(unicode.Han, value) {
			count += 2
		} else {
			count += 1
		}
	}

	if count < 4 {
		return false, "设备名称短于4"
	}
	if count > 64 {
		return false, "设备名称长于64"
	}

	return true, ""
}

func CheckAutoAddDeviceNumber(number int)(result bool, msg string) {
	if number > 1000 {
		return false, "一次批量添加不能大于1000个"
	}
	if number == 0 {
		return false, "数量不能为0"
	}
	return true, ""
}

func CheckDeviceIDQuality(id int)(result bool, msg string) {
	ids := database.GetAllDeviceID()
	for _, value := range ids {
		if id == value {
			return true, ""
		}
	}
	return false, "设备ID不是有效参数"
}

func CheckDeviceLabelQualify(label []map[string]string) (result bool, msg string) {
	for _, v := range label {
		key := v["key"]
		if keyRes, msg := CheckDeviceLabelKeyQuality(key); keyRes != true {
			return false, msg
		}
		value := v["value"]
		if valueRes, msg := CheckDeviceLabelValueQuality(value); valueRes != true {
			return false, msg
		}
	}
	return true, ""
}