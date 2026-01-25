package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Confirm(message string) bool {
	// Ask for confirmation
	fmt.Print(message)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(strings.ToLower(input))
	// We check for "y", "yes", or empty string (if you want Enter to mean Yes)
	if input == "y" || input == "yes" || input == "" {
		return true
	}

	return false
}
