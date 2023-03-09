package billing

import (
	"fmt"
	"time"
)

type Billing struct {
}

func New() *Billing {
	return &Billing{}
}

func (b *Billing) Run() error {
	fmt.Print("Billing is started...")
	time.Sleep(time.Second * 5)
	fmt.Print("Billing is finished...")
	return nil
}
