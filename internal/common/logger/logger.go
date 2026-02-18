package logger

import "fmt"

func Info(text string) {
	fmt.Printf("[info] %s", text)
}
func Error(text string) {
	fmt.Printf("[error] %s", text)
}
func Fatal(text string) {
	fmt.Printf("[fatal] %s", text)
}
