package main

import "github.com/made2591/go-patterns/creationals/singleton/logger"

func main() {

	firstLogger := logger.New()
	firstLogger.SetLevel(2)
	firstLogger.Write(nil, "firstLogger level: %d\n", firstLogger.GetLevel())

	secondLogger := logger.New()
	secondLogger.Write(nil, "secondLogger level: %d\n", secondLogger.GetLevel())

}
