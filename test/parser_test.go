package test

import (
	"cron-parser/pkg/cronparser"
	"testing"
)

func TestParseCronExpression(t *testing.T) {
	expr := "*/15 0 1,15 * 1-5 /usr/bin/find"
	expected := `minute        */15
hour          0
day of month  1 15
month         1 2 3 4 5 6 7 8 9 10 11 12
day of week   1-5
command       /usr/bin/find`
	runTest(t, expr, expected)
}

func TestParseSpecificValueCronExpression(t *testing.T) {
	expr := "30 6 1 1 1 /usr/bin/find"
	expected := `minute        30
hour          6
day of month  1
month         1
day of week   1
command       /usr/bin/find`
	runTest(t, expr, expected)
}

func TestParseRangeCronExpression(t *testing.T) {
	expr := "10-20 1-5 1-15 1-6 1-3 /usr/bin/find"
	expected := `minute        10-20
hour          1-5
day of month  1-15
month         1-6
day of week   1-3
command       /usr/bin/find`
	runTest(t, expr, expected)
}

func runTest(t *testing.T, expr string, expected string) {
	parser := cronparser.NewParser(expr)
	result, err := parser.Parse(expr)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result != expected {
		t.Fatalf("Expected %v, got %v", expected, result)
	}
}
