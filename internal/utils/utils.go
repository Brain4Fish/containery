package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func CleanScreen() {
	// TODO: support commands for different archs
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	if cmd.Run() != nil {
		fmt.Printf("There is an error when cleaning screen")
	}
}
