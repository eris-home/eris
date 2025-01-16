package cli

import (
	"fmt"
	"strings"
)

type OutputFormatter struct{}

func NewOutputFormatter() *OutputFormatter {
	return &OutputFormatter{}
}

func (of *OutputFormatter) DisplaySuccess(message string) {
	fmt.Printf("\033[32m[SUCCESS]\033[0m %s\n", message)
}

func (of *OutputFormatter) DisplayError(message string) {
	fmt.Printf("\033[31m[ERROR]\033[0m %s\n", message)
}

func (of *OutputFormatter) DisplayInfo(message string) {
	fmt.Printf("\033[34m[INFO]\033[0m %s\n", message)
}

func (of *OutputFormatter) FormatList(items []string) string {
	var builder strings.Builder
	builder.WriteString("\nAvailable Commands:\n")
	for _, item := range items {
		builder.WriteString(fmt.Sprintf("- %s\n", item))
	}
	return builder.String()
}

func (of *OutputFormatter) FormatTable(headers []string, rows [][]string) {
	fmt.Println(strings.Repeat("=", 40))
	fmt.Printf("%-15s%-15s\n", headers[0], headers[1])
	fmt.Println(strings.Repeat("=", 40))

	for _, row := range rows {
		fmt.Printf("%-15s%-15s\n", row[0], row[1])
	}
	fmt.Println(strings.Repeat("=", 40))
}