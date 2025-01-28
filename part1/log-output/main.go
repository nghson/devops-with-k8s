package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func main() {
	randomString := uuid.New().String()
	for {
		timestamp := time.Now().UTC().Format(time.RFC3339)
		fmt.Printf("%s: %s\n", timestamp, randomString)

		time.Sleep(5 * time.Second)
	}
}

