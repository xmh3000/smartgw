package led

import (
	"smartgw/lib/io"
	"time"
)

const (
	NORMAL = 1000
	ERROR  = 200
)

var runningDelay = NORMAL

func InitLed() {
	//程序运行RUN开始，22引脚根据米尔科技，GPIOX_Y的引脚编号转化公式，得知22引脚对应编号10,1s闪烁
	io.OsCommandSilent("echo 10 > /sys/class/gpio/export")
	io.OsCommandSilent("echo \"out\" > /sys/class/gpio/gpio10/direction")
	SetRunning()

	//初始化LINK状态，连接服务器则常亮，49引脚对应编号9，当程序刚运行时先让灯灭
	io.OsCommandSilent("echo 9 > /sys/class/gpio/export")
	io.OsCommandSilent("echo \"out\" > /sys/class/gpio/gpio9/direction")

	// reset 设置
	io.OsCommandSilent("echo 128 > /sys/class/gpio/export")
	io.OsCommandSilent("echo \"in\" > /sys/class/gpio/gpio128/direction")
	SetReset()
}

func SetRunningNormal() {
	runningDelay = NORMAL
}

func SetRunningSerialErr() {
	runningDelay = ERROR
}

func SetRunning() {
	go func() {
		for {
			TurnOn("10")
			time.Sleep(time.Duration(runningDelay) * time.Millisecond)
			TurnOff("10")
			time.Sleep(time.Duration(runningDelay) * time.Millisecond)
		}
	}()
}

func SetReset() {
	go func() {
		for {
			command := "cat /sys/class/gpio/gpio128/value"
			result := io.OsCommandSilent(command)
			if result == "0" {
				// 恢复网卡设置
				command := "ifconfig eth1 192.168.100.12 up"
				io.OsCommandSilent(command)
			}
			time.Sleep(time.Duration(5) * time.Second)
		}
	}()
}

func TurnOn(pin string) {
	command := "echo 0 > /sys/class/gpio/gpio" + pin + "/value"
	io.OsCommandSilent(command)
}

func TurnOff(pin string) {
	command := "echo 1 > /sys/class/gpio/gpio" + pin + "/value"
	io.OsCommandSilent(command)
}
