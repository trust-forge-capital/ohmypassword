package ui

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mattn/go-runewidth"
)

var spinnerChars = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

type Spinner struct {
	message      string
	stopped      bool
	spinnerIndex int
}

func NewSpinner(message string) *Spinner {
	return &Spinner{message: message}
}

func (s *Spinner) Start() {
	go func() {
		for !s.stopped {
			fmt.Printf("\r%s %s", spinnerChars[s.spinnerIndex], s.message)
			s.spinnerIndex = (s.spinnerIndex + 1) % len(spinnerChars)
			time.Sleep(100 * time.Millisecond)
		}
	}()
}

func (s *Spinner) Stop() {
	s.stopped = true
	fmt.Printf("\r%s\n", strings.Repeat(" ", runewidth.StringWidth(s.message)+2))
}

func PrintSuccess(message string) {
	printWithColor(message, 32)
}

func PrintError(message string) {
	printWithColor(message, 31)
}

func PrintWarning(message string) {
	printWithColor(message, 33)
}

func PrintInfo(message string) {
	printWithColor(message, 36)
}

func printWithColor(message string, colorCode int) {
	fmt.Fprintf(os.Stderr, "\033[%dm%s\033[0m\n", colorCode, message)
}