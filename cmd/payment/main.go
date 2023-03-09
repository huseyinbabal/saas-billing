package main

import (
	"log"
	"saas/internal/domain/payment"
)

func main() {
	p := payment.New()
	err := p.Run()
	if err != nil {
		log.Fatalf("failed to run payment. %v", err)
	}
}
