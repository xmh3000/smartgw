package worker

import (
	"encoding/json"
	"fmt"
	"smartgw/api/domain"
	"testing"
	"time"
)

func TestDelayRead(t *testing.T) {
	//stop := make(chan struct{}, 1)
	timer := time.NewTimer(3 * time.Second)
	//go func() {
	for {
		select {
		//case <-stop:
		case <-timer.C:
			fmt.Println("停止线程...")
			return
		default:
			var x int
			fmt.Println("请输入一个数字")
			fmt.Scan(&x)
			time.Sleep(100 * time.Millisecond)
		}
	}
	//}()
	fmt.Println("等待数据读取...")
	//<-timer.C
	//stop <- struct{}{}
	time.Sleep(5 * time.Second)
	fmt.Println("结束")
}

func TestDeviceCopy(t *testing.T) {
	device := domain.Device{
		Name: "aaa",
		Type: domain.DeviceType{
			Name:   "bbb",
			Driver: "ccc",
			Properties: []domain.DeviceProperty{{
				Name:        "power",
				Description: "aaa",
				Type:        "int",
				Length:      13,
				Decimal:     4,
				Unit:        "ss",
				Value:       90.0,
				Reported:    false,
			}},
		},
		Address:        "",
		Collector:      domain.Collector{},
		Alone:          false,
		Serial:         domain.Serial{},
		Online:         false,
		CollectTime:    time.Time{},
		CollectTotal:   0,
		CollectSuccess: 0,
		ReportTime:     time.Time{},
		ReportTotal:    0,
		ReportSuccess:  0,
	}

	TempDeviceCopy(device)
}

func TempDeviceCopy(task any) {
	device := task.(domain.Device)
	deviceCopy := domain.Device{}

	temp, _ := json.Marshal(device)
	_ = json.Unmarshal(temp, &deviceCopy)

	device.Type.Properties[0].Value = 20

	fmt.Println(device.Type.Properties[0].Value, deviceCopy.Type.Properties[0].Value)
}
