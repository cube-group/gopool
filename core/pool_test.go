package core

import (
	"fmt"
	"testing"
	"time"
)

func TestGood(t *testing.T) {
	GoPoolMaxNum = 10
	for i := 0; i < 10; i++ {
		var err = Go(func() {
			time.Sleep(time.Second)
		})
		time.Sleep(time.Millisecond * 500)
		fmt.Println(err)
	}
	time.Sleep(10*time.Second)
}

func TestBad(t *testing.T) {
	GoPoolMaxNum = 2
	GoPoolDebug = true

	for i := 0; i < 100; i++ {
		Go(func() {
			time.Sleep(time.Millisecond*2)
		})
		time.Sleep(time.Millisecond)
	}
	time.Sleep(10*time.Second)
}
