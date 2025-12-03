package helpers

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
)

var Scanner = bufio.NewScanner(os.Stdin)

// Function to scan user input from console
func ScanInput(promptString string) string {
	fmt.Print(promptString)
	Scanner.Scan()
	return Scanner.Text()
}

// Function to pause execution until user presses Enter
func PressEnter() {
	fmt.Print("Press [Enter] to continue...")

	if _, err := fmt.Scanln(); err != nil {
		if err != io.EOF {
			fmt.Println("Error reading input:", err)
		}
	}
}

// Function to clear console screen
func ClearScreen() {
	// Clear screen for Unix systems
	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		// Clear screen for Windows
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to clear screen:", err)
	}
}
