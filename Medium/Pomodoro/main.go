package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	var duration time.Duration

	// Get the desired timer duration from the user
	fmt.Println("Choose the timer duration:")

	var choice int
	fmt.Scan(&choice)

	duration = time.Duration(choice) * time.Minute

	// Start the timer
	timer := time.NewTimer(duration)
	fmt.Printf("Timer set for %s. Press Enter to start.\n", duration)

	// Wait for user input to start the timer
	fmt.Scanln()
	
	// Start the countdown
	fmt.Printf("Started....")

	<-timer.C

	fmt.Println("Time's up!")
	alert()
}

func alert() {

	var cmd *exec.Cmd

	switch os := runtime.GOOS; os {
	case "windows":
		cmd = exec.Command("cmd", "/c", "echo 'Time is up' && pause")
	case "linux":
		cmd = exec.Command("notify-send", "Pomodoro", "Time's Up")
	default:
		fmt.Println("Unsupported operating system")
		return
	}

	err := cmd.Start()
	if err != nil {
		fmt.Println("Failed to execute the command:", err)
		return
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println("Command execution error:", err)
	}
}
