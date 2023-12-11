package channel

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestPriorityChannel(t *testing.T) {
	channel := NewPriorityChannel()
	channel.SetPriorWorker(func(task any) {
		fmt.Println(task)
	})

	channel.SetNormalWorker(func(task any) {
		fmt.Println(task)
	})

	channel.Start()
	go func() {
		for i := 1; ; i++ {
			channel.DispatchNormalTask("normal task" + strconv.Itoa(i))
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for i := 1; ; i++ {
			channel.DispatchPriorTask("height task" + strconv.Itoa(i))
			time.Sleep(3 * time.Second)
		}
	}()
	fmt.Println("1")
	time.Sleep(30 * time.Second)
	fmt.Println("2")
	channel.Stop()
	time.Sleep(3 * time.Second)
	fmt.Println("3")
}

func TestChannelLen(t *testing.T) {
	channel := make(chan int, 100)

	go func() {
		for {
			i := <-channel
			fmt.Println(i, " ", len(channel))
		}
	}()

	for i := 0; i < 1000; i++ {
		channel <- i
	}

	time.Sleep(5 * time.Second)
}
