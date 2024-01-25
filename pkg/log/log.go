package log

import "log"

const (
	infoLevel  = "INFO"
	warnLevel  = "WARN"
	errorLevel = "ERROR"
	fatalLevel = "FATAL"
)

func Info(v ...any) {
	print(infoLevel, v...)
}

func Infof(format string, v ...any) {
	printf(infoLevel, format, v...)
}

func Warn(v ...any) {
	print(warnLevel, v...)
}

func Warnf(format string, v ...any) {
	printf(warnLevel, format, v...)
}

func Error(v ...any) {
	print(errorLevel, v...)
}

func Errorf(format string, v ...any) {
	printf(errorLevel, format, v...)
}

func Fatal(v ...any) {
	v = append([]any{fatalLevel}, v...)
	log.Fatal(v...)
}

func Fatalf(format string, v ...any) {
	log.Fatalf(fatalLevel+" "+format, v...)
}

func print(level string, v ...any) {
	v = append([]any{level + " "}, v...)
	log.Print(v...)
}

func printf(level, format string, v ...any) {
	log.Printf(level+" "+format, v...)
}
