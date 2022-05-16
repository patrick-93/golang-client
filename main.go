package main

import (
	"client/linux"
	"client/windows"
	"fmt"
	"runtime"
)

func main() {
	switch runtime.GOOS {
	case "linux":
		linux.LinuxClient()
	case "windows":
		windows.Windows_client()
	default:
		fmt.Println("Found an unsupported OS, exiting")
	}
}
