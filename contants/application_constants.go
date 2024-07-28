package constants

const (
	MINUTE = iota
	HOUR
	DOM
	MONTH
	DOW
	COMMAND
)

var (
	months     = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}
	daysOfWeek = []string{"0", "1", "2", "3", "4", "5", "6"}
)
