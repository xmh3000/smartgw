package io

import (
	"bytes"
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"go.uber.org/zap"
	"os"
	"os/exec"
	"path/filepath"
	"smartgw/api/domain"
	"strings"
	"time"
)

type SystemStateTemplate struct {
	MemTotal      string `json:"MemTotal"`
	MemUse        string `json:"MemUse"`
	DiskTotal     string `json:"DiskTotal"`
	DiskUse       string `json:"DiskUse"`
	Name          string `json:"Name"`
	SN            string `json:"SN"`
	HardVer       string `json:"HardVer"`
	SoftVer       string `json:"SoftVer"`
	SystemRTC     string `json:"SystemRTC"`
	RunTime       string `json:"RunTime"` //累计时间
	DeviceTotal   string `json:"DeviceTotal"`
	DeviceOnline  string `json:"DeviceOnline"` //设备在线率
	OpenEveryTime string `json:"OpenEveryTime"`
	//DevicePacketLoss string `json:"DevicePacketLoss"` //设备丢包率
}

var SystemState = SystemStateTemplate{
	MemTotal:      "0",
	MemUse:        "0",
	DiskTotal:     "0",
	DiskUse:       "0",
	Name:          "GreatWall",
	SN:            "2023-02-08",
	HardVer:       "YG210-N485(2)",
	SoftVer:       "V 1.0.7",
	SystemRTC:     "2022-06-2 12:00:00",
	RunTime:       "0",
	DeviceTotal:   "0",
	DeviceOnline:  "0",
	OpenEveryTime: "不支持不同波特率",
	//DevicePacketLoss: "0",
}

var timeStart time.Time

func GetMemState() {
	v, _ := mem.VirtualMemory()
	SystemState.MemTotal = fmt.Sprintf("%d", v.Total/1024/1024)
	SystemState.MemUse = fmt.Sprintf("%5.2f", v.UsedPercent)
}

func GetDiskState() {
	exeCurDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	v, _ := disk.Usage(exeCurDir)
	SystemState.DiskTotal = fmt.Sprintf("%d", v.Total/1024/1024)
	SystemState.DiskUse = fmt.Sprintf("%5.2f", v.UsedPercent)
}

func GetStartTime() {
	timeStart = time.Now()
}

func GetRuntime() {
	elapsed := time.Since(timeStart)
	sec := int64(elapsed.Seconds())
	day := sec / 86400
	hour := sec % 86400 / 3600
	min := sec % 3600 / 60
	sec = sec % 60
	strRunTime := fmt.Sprintf("%d天%d时%d分%d秒", day, hour, min, sec)
	SystemState.SystemRTC = time.Now().Format("2006-01-02 15:04:05")
	SystemState.RunTime = strRunTime
}

func SystemReboot() {
	cmd := exec.Command("reboot")
	var out bytes.Buffer
	cmd.Stdout = &out
	_ = cmd.Start()

	str := out.String()
	fmt.Println(str)
}

func OsCommand(cmd string) string {
	zap.S().Debug(cmd)
	result, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		zap.S().Error(err.Error())
	}
	var s = strings.TrimSpace(string(result))
	zap.S().Debug(s)
	return s
}

func OsCommandSilent(cmd string) string {
	result, _ := exec.Command("/bin/sh", "-c", cmd).Output()
	//if err != nil {
	//	logger.Zap.Error(err.Error())
	//}
	var s = strings.TrimSpace(string(result))
	return s
}

func CmdSetStaticIP(eth *domain.Ethernet) {
	out, err := exec.Command("/bin/sh", "-c",
		fmt.Sprintf("ifconfig %s %s netmask %s", eth.Name, eth.ConfigIP, eth.ConfigNetmask)).Output()
	if err != nil {
		zap.S().Debugf("网卡[%s]设置IP[%s]Netmask[%s]失败 %s %v", eth.Name, eth.ConfigIP, eth.ConfigNetmask, string(out), err)
	} else {
		zap.S().Debugf("网卡[%s]设置IP[%s]Netmask[%s]成功", eth.Name, eth.ConfigIP, eth.ConfigNetmask)
	}

	out, err = exec.Command("/sbin/route", "add", "default", "gw", eth.ConfigGateway).Output()
	if err != nil {
		zap.S().Debugf("网卡[%s]添加默认网关[%s]失败 %s %v", eth.Name, eth.ConfigGateway, string(out), err)
		return
	}
	zap.S().Debugf("网卡[%s]添加默认网关[%s]成功", eth.Name, eth.ConfigGateway)

}

func CmdSetDHCP(eth *domain.Ethernet) (error, string) {

	//非阻塞,动态获取IP有可能不成功
	out, err := exec.Command("/bin/sh", "-c",
		fmt.Sprintf("udhcpc -i %s", eth.Name)).Output()
	if err != nil {
		zap.S().Debug("网卡[%s]动态获取IP失败 %s %v", eth.Name, string(out), err)
		return err, "nil"
	}
	zap.S().Debug("网卡[%s]动态获取IP成功 %s", eth.Name, string(out))

	return nil, string(out)
}
