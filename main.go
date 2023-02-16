package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("present", "-http", "0.0.0.0:6969", "-orighost", "pre")

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
