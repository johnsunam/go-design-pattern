package main

import "fmt"

func main() {
	// log := getLoggerInstance()
	// log.SetLogLevel(1)
	// log.Log("this first log message")

	// log = getLoggerInstance()
	// log.SetLogLevel(2)
	// log.Log("this second log message")

	// log = getLoggerInstance()
	// log.SetLogLevel(3)
	// log.Log("this third log message")

	for i := 0; i < 5; i++ {
		go getLoggerInstance()
	}
	fmt.Scanln()
}
