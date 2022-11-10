package core

import (
	"errors"
	"log"
	"sync"
)

type PoolFunc func()

type PoolGoItem struct {
	id int
	ch chan PoolFunc
}

func newPoolGoItem(index int, f PoolFunc) *PoolGoItem {
	item := &PoolGoItem{
		id: index,
		ch: make(chan PoolFunc, 1),
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				poolDebug("Fatal", index, err)
			}
		}()
		for {
			f := <-item.ch
			f()
		}
	}()
	item.run(f)
	return item
}

func (t *PoolGoItem) IsRunning() bool {
	return len(t.ch) > 0
}

func (t *PoolGoItem) run(f PoolFunc) {
	t.ch <- f
}

var GoPoolMaxNum = 5
var GoPoolDebug bool
var poolMaps sync.Map

func Go(f PoolFunc) error {
	var done bool
	var currentLength int
	poolMaps.Range(func(key, value any) bool {
		item := value.(*PoolGoItem)
		if !item.IsRunning() {
			poolDebug("pool", item.id, "reused")
			item.run(f)
			done = true
		}
		currentLength++
		return true
	})

	if !done {
		if currentLength >= GoPoolMaxNum {
			poolDebug("pool full")
			return errors.New("pool full")
		} else {
			poolDebug("new pool")
			poolMaps.Store(currentLength, newPoolGoItem(currentLength, f))
		}
	}
	return nil
}

func poolDebug(v ...interface{}) {
	if !GoPoolDebug {
		return
	}
	v = append([]interface{}{"Debug"}, v...)
	log.Println(v...)
}
