package util

import (
	"fmt"
	"log"
)

func CheckError(err error, msg string) {
	if err != nil {
		log.Fatal(fmt.Sprintf("Msg: %v -> Error: %v", msg, err.Error()))
	}
}
