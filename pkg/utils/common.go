package utils

import (
	"fmt"
)

func checkPanic(err interface{}) {
  if err != nil {
    fmt.Println("Panic:", err)
    os.Exit(1)
  }
}
