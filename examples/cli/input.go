package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type InputReader struct{}

func NewInputReader() *InputReader {
	return &InputReader{}
}

func (ir *InputReader) ReadInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\033[34m[INPUT] >\033[0m ")
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func (ir *InputReader) ParseInput(input string) (string, []string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return "", nil
	}
	return parts[0], parts[1:]
}

func (ir *InputReader) ConfirmAction(prompt string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s (y/n): ", prompt)
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(strings.ToLower(response))
	return response == "y" || response == "yes"
}