package main

import (
	"log"
	"saas/internal/domain/billing"
)

func main() {
	b := billing.New()
	err := b.Run()
	if err != nil {
		log.Fatalf("failed to run 2 billing. %v", err)
	}
}
