package main

import (
	"log"
	"os"
	"os/exec"
	"runtime"
)

func clearConsole() {
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		log.Fatal("unsupported platform. cannot clear console.")
	}
}
