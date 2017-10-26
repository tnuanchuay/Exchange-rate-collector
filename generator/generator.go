package generator

import (
	"time"
	"sync"
)

const (
	OK	=	true
	END	=	false
)

type Generator struct{
	isRunning	bool
	Chan		chan interface{}
	Handler		func(c chan interface{}) bool
	Wg		*sync.WaitGroup
}

func (c *Generator) Run(){
	for c.isRunning {
		time.Sleep(10 * time.Second)
		c.Wg.Add(1)
		c.isRunning = c.Handler(c.Chan)
		c.Wg.Done()
	}
}

func (c *Generator) Stop(){
	c.isRunning = false
}
