package variations

import (
	"fmt"
	"time"
)


// ShowProgress just prints "." , so that you won't get bored
func ShowProgress(done <- chan bool) {
	for {
		select {
		case <- done:
			fmt.Println()
			return
		default:
			time.Sleep(time.Second)
			fmt.Print(".")
		}
	}
}