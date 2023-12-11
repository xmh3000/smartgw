package watchdog

import (
	"smartgw/lib/config"
	"smartgw/lib/logger"
	"testing"
	"time"
)

func TestWatchDog(t *testing.T) {
	watchDog := NewWatchDog()
	defer func() {
		watchDog.Stop()
		time.Sleep(10 * time.Second)
		if watchDog.Running() {
			t.Error("看门狗没有正确关闭")
		} else {
			t.Log("看门狗正确关闭")
		}
	}()

	cfg := config.NewConfig()
	logger.NewLogger(cfg)

	InitWatchDog(cfg, watchDog)

	time.Sleep(5 * time.Second)

	if !watchDog.Running() {
		t.Error("看门狗没有正确运行")
	}

	time.Sleep(30 * time.Second)
}
