package main

import (
	"bufio"
	"cron-parser/pkg/cronparser"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cronparser \"<cron_expression>\"")
		fmt.Println("Example: cronparser \"* * * * *\"")
		fmt.Println("No cron expression provided. Entering interactive mode...")
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter cron expression: ")
		cronExpr, _ := reader.ReadString('\n')
		parseAndPrint(cronExpr)
	} else {
		cronExpr := os.Args[1]
		parseAndPrint(cronExpr)
	}
}

func parseAndPrint(cronExpr string) {
	parser := cronparser.NewParser(cronExpr)
	result, err := parser.Parse(cronExpr)
	if err != nil {
		fmt.Println("Error parsing cron expression:", err)
		fmt.Println("Please ensure your cron expression is valid. For example: \"* * * * *\"")
		os.Exit(1)
	}
	fmt.Println(result)
}
