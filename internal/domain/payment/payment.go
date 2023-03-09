package payment

import (
	"fmt"
	"time"
)

type Payment struct {
}

func New() *Payment {
	return &Payment{}
}

func (p *Payment) Run() error {
	fmt.Print("Payment is started...")
	time.Sleep(time.Second * 5)
	fmt.Print("Payment is finished...")
	return nil
}
