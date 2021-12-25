package handlers

import (
	"fmt"
	"log"
)

// CheckErrPanic Generic panic if err != nil
func CheckErrPanic(e error) {
	if e != nil {
		panic(e)
	}
}

// CheckErrLog Generic log if err != nil
func CheckErrLog(e error, additionalMsg string) {
	msg := fmt.Sprintf("Err: %s Handler msg: %s", e, additionalMsg)
	if e != nil {
		log.Println(msg)
	}
}
