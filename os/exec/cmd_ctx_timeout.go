//
package main

import (
	"context"
	"log"
	"os/exec"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := exec.CommandContext(ctx, "sleep", "5").Run(); err != nil {
		log.Println("error:", err)
		// This will fail after 100 milliseconds. The 5 second sleep
		// will be interrupted.
	}
	// block wen ctx timeout or cmd exec finished
	log.Println("sleep finished!")
}
