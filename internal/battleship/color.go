package battleship

import "fmt"

type Color string

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	purple = "\033[35m"
	cyan   = "\033[36m"
	white  = "\033[37m"
)

func RedPrintf(format string, a ...interface{}) {
	fmt.Print(Red(format, a...))
}

func GreenPrintf(format string, a ...interface{}) {
	fmt.Print(Green(format, a...))
}

func YellowPrintf(format string, a ...interface{}) {
	fmt.Print(Yellow(format, a...))
}

func BluePrintf(format string, a ...interface{}) {
	fmt.Print(Blue(format, a...))
}

func PurplePrintf(format string, a ...interface{}) {
	fmt.Print(Purple(format, a...))
}

func CyanPrintf(format string, a ...interface{}) {
	fmt.Print(Cyan(format, a...))
}

func WhitePrintf(format string, a ...interface{}) {
	fmt.Print(White(format, a...))
}

func Red(format string, a ...interface{}) string {
	return colorize(red, fmt.Sprintf(format, a...))
}

func Green(format string, a ...interface{}) string {
	return colorize(green, fmt.Sprintf(format, a...))
}

func Yellow(format string, a ...interface{}) string {
	return colorize(yellow, fmt.Sprintf(format, a...))
}

func Blue(format string, a ...interface{}) string {
	return colorize(blue, fmt.Sprintf(format, a...))
}

func Purple(format string, a ...interface{}) string {
	return colorize(purple, fmt.Sprintf(format, a...))
}

func Cyan(format string, a ...interface{}) string {
	return colorize(cyan, fmt.Sprintf(format, a...))
}

func White(format string, a ...interface{}) string {
	return colorize(white, fmt.Sprintf(format, a...))
}

func colorize(color Color, input string) string {
	return string(color) + input + string(reset)
}
