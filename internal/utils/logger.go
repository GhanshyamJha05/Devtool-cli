package utils

import "fmt"

// ANSI color codes for terminal output
const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorCyan   = "\033[36m"
	colorGray   = "\033[90m"
)

// PrintSuccess prints a green success message to the terminal.
func PrintSuccess(msg string) {
	fmt.Printf("%s✔ %s%s\n", colorGreen, msg, colorReset)
}

// PrintError prints a red error message to the terminal.
func PrintError(msg string) {
	fmt.Printf("%s✖ %s%s\n", colorRed, msg, colorReset)
}

// PrintInfo prints a cyan informational message to the terminal.
func PrintInfo(msg string) {
	fmt.Printf("%sℹ %s%s\n", colorCyan, msg, colorReset)
}

// PrintWarning prints a yellow warning message to the terminal.
func PrintWarning(msg string) {
	fmt.Printf("%s⚠ %s%s\n", colorYellow, msg, colorReset)
}

// PrintDebug prints a gray debug message, but only when verbose mode is enabled.
// The Verbose flag is checked by the caller before invoking this.
func PrintDebug(msg string) {
	fmt.Printf("%s[DEBUG] %s%s\n", colorGray, msg, colorReset)
}
