package parser

import (
	"cron-parser/constants"
	"errors"
	"fmt"
	"strings"
)

func ParseStdCronExpression(expr string) (string, error) {
	fields := strings.Fields(expr)
	if len(fields) != 6 {
		return "", errors.New("invalid cron expression")
	}

	minute := expandField(fields[constants.MINUTE], 0, 59)
	hour := expandField(fields[constants.HOUR], 0, 23)
	dayOfMonth := expandField(fields[constants.DOM], 1, 31)
	month := expandField(fields[constants.MONTH], 1, 12)
	dayOfWeek := expandField(fields[constants.DOW], 0, 6)
	command := fields[constants.COMMAND]

	result := fmt.Sprintf("minute        %s\nhour          %s\nday of month  %s\nmonth         %s\nday of week   %s\ncommand       %s",
		strings.Join(minute, " "),
		strings.Join(hour, " "),
		strings.Join(dayOfMonth, " "),
		strings.Join(month, " "),
		strings.Join(dayOfWeek, " "),
		command)

	return result, nil
}

func expandField(field string, min, max int) []string {
	if field == "*" {
		return generateRange(min, max)
	}
	// More parsing logic for other formats like "*/15", "1-5", etc.
	return strings.Split(field, ",")
}

func generateRange(min, max int) []string {
	var result []string
	for i := min; i <= max; i++ {
		result = append(result, fmt.Sprintf("%d", i))
	}
	return result
}
