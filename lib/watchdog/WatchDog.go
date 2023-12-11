package watchdog

import (
	"go.uber.org/zap"
	"os"
	"smartgw/lib/config"
	"time"
)

type (
	WatchDog interface {
		Start()
		Stop()
		Running() bool
	}

	watchDog struct {
		file *os.File
		stop chan bool
	}
)

var _ WatchDog = (*watchDog)(nil)

func NewWatchDog() WatchDog {
	return &watchDog{
		file: nil,
		stop: make(chan bool, 1),
	}
}

func InitWatchDog(config *config.Config, dog WatchDog) {
	if config.Mode == "release" && !dog.Running() {
		dog.Start()
	}
}

func (w *watchDog) Start() {
	go w.Watch()
}

func (w *watchDog) Stop() {
	w.stop <- true
}

func (w *watchDog) Watch() {
	err := error(nil)
	w.file, err = os.OpenFile("/dev/watchdog", os.O_WRONLY, 0666)
	if err != nil {
		zap.S().Errorf("打开看门狗失败: %v", err)
		w.file = nil
		return
	}
	defer func() {
		w.file.Close()
		w.file = nil
	}()

	buffer := make([]byte, 1)
	buffer[0] = 0
	for {
		select {
		case <-w.stop:
			return
		default:
			_, err := w.file.Write(buffer)
			if err != nil {
				zap.S().Errorf("看门狗写入失败: %v", err)
			}
			time.Sleep(30 * time.Second)
		}
	}
}

func (w *watchDog) Running() bool {
	return w.file != nil
}
