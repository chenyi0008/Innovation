package mqtt

import "strings"

// ProcessRequest 处理mqtt请求
func ProcessRequest(topic, payload string) {
	splitArr := strings.Split(topic, "/")
	//设备类型
	deviceType := splitArr[1]
	//设备序列号
	serialNum := splitArr[2]
	//指令
	instruct := splitArr[3]

	switch deviceType {
	case "warning":
		warningDistribute(serialNum, instruct)
	case "detect":
	}
}

func warningDistribute(serialNum, instruct string) {

}

func detectDistribute(serialNum, instruct string) {
	switch instruct {
	case "warning":
		detectWarning(serialNum)
	case "status":
		println("status")

	}
}

func detectWarning(serialNum string) {

}
