package time_test

import (
	"fmt"
	"testing"
	"time"
)
//定时器的测试
func Test_Timer(t *testing.T) {
	//sleep 实现定时器
	fmt.Println(time.Now())
	time.Sleep(time.Second)
	fmt.Println(time.Now())
	//使用timer 实现定时器
	timer := time.NewTimer(time.Second)

	fmt.Println(<-timer.C)
	timer.Reset(time.Second)
	fmt.Println(timer.Stop())
	fmt.Println(<-timer.C)
	//time after

	fmt.Println(<-time.After(time.Second))
}

func Test_Timer_Self(t *testing.T) {
	// t := Timer1{}
}
