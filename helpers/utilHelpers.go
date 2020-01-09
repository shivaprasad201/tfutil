package helpers

import "log"

// HandleError logs the error onto the console
func HandleError(e error, msg string) {
	if e != nil {
		log.Fatal(e, msg)
	}
}
