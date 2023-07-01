package version

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/mattn/go-runewidth"
)

var (
	Version string
	Date    string
)

func init() {
	yellow := color.New(color.FgYellow).SprintFunc()

	// Calculate the number of leading spaces needed for center alignment
	terminalWidth := 80 // Assuming a terminal width of 80 characters
	message := fmt.Sprintf("Version: %s", yellow(Version))
	spaces := (terminalWidth - runewidth.StringWidth(message)) / 2

	// Create a string with the required number of spaces
	padding := strings.Repeat(" ", spaces)

	// Print the version and build date with centered alignment
	fmt.Printf("%s%s\n", padding, message)
	fmt.Printf("%sBuild date: %s\n", padding, yellow(Date))
}
