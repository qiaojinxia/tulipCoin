package utils

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
)

/**
 * Created by @CaomaoBoy on 2021/5/22.
 *  email:<115882934@qq.com>
 */



var Wg = &sync.WaitGroup{}
func GO (do func()){
	go func() {
		Wg.Add(1)
		do()
		Wg.Done()
	}()

}

func GO_Func(fn func()){
	Wg.Add(1)
	go func() {
		select {
		case <-stopChanForGo:
			return
		default:
			Try(fn).CatchAll(func(err error) {
				log.Printf("Error : %v",err)
			})

		}
		Wg.Done()
	}()
}





var waitAll sync.WaitGroup
var goid uint32
var gocount int32
var PoolSize int32 = 10
var stopChanForGo = make(chan struct{})

func GO_Func1(fn func()) {
	pc := PoolSize
	select {
	case poolChan <- fn:
		return
	default:
		pc = atomic.AddInt32(&poolGoCount, 1)
		if pc > PoolSize {
			atomic.AddInt32(&poolGoCount, -1)
		}
	}
	waitAll.Add(1)
	//id := atomic.AddUint32(&goid, 1)
	c := atomic.AddInt32(&gocount, 1)
	go func() {
		//Try(fn, nil)
		for pc <= PoolSize {
			select {
			case <-stopChanForGo:
				pc = PoolSize + 1
			case nfn := <-poolChan:
				Try(nfn).CatchAll(func(err error) {
					log.Printf("Error : %v",err)
				})
			}
		}
		waitAll.Done()
		c = atomic.AddInt32(&gocount, -1)
		fmt.Println(c)
	}()
}


var poolChan  = make(chan func(),1024)
var poolGoCount int32
func Go(fn func()) {
	var pc int32
	select {
	case poolChan <- fn:
		return
	default:
		pc = atomic.AddInt32(&poolGoCount, 1)
		if pc > PoolSize {
			atomic.AddInt32(&poolGoCount, -1)
		}
	}
	Wg.Add(1)
	atomic.AddInt32(&gocount, 1)
	go func() {
		for pc <= PoolSize{
			select {
			case <-stopChanForGo:
				return
			case nfn := <- poolChan:
				Try(nfn).CatchAll(func(err error) {
					log.Printf("Error : %v",err)
				})
			}
		}
	}()
	Wg.Done()
	atomic.AddInt32(&gocount, -1)
}