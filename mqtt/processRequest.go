package mqtt

import (
	"Innovation/model"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

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
		detectDistribute(serialNum, instruct)
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

type detectWarningPublishInfo struct {
	Timestamp int64  `json:"timestamp"`
	Position  string `json:"position"`
}

func detectWarning(serialNum string) {
	flag, inspect := model.InspectGetInfoBySerialNum(serialNum)
	if flag {
		alarmList := inspect.AlarmEquipmentList
		for _, alarm := range alarmList {
			topic := fmt.Sprintf("device/warning/%s/warning", alarm.SerialNum)
			publishInfo := detectWarningPublishInfo{
				Timestamp: getTimestamp(),
				Position:  "position1",
			}
			marshal, err := json.Marshal(publishInfo)
			if err != nil {
				panic(err)
			}
			println("marshal:", string(marshal))
			publish(topic, string(marshal))
		}

	}
}

func getTimestamp() int64 {
	// 获取当前时间的Unix时间戳（纳秒）
	return time.Now().UnixNano()
}
