package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func (p *person) doActivity(activity string, min, max int, wg *sync.WaitGroup) {
	go func() {
		t := time.Second * time.Duration(min+rand.Intn(max-min))
		time.Sleep(t)
		fmt.Println(*p, " spent ", t.Seconds(), "seconds", a)
	}()
}
