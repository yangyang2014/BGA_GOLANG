package time

import (
	"fmt"
	"testing"
	"time"
)

func Test_timeSub(t *testing.T) {
	begin := time.Now()
	time.Sleep(time.Second)
	end := time.Now()
	fmt.Println(end.Sub(begin).String())
}