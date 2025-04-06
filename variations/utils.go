package variations

import (
	"fmt"
	"time"
)


// ShowProgress just prints "." to show progress, keeping you engaged (or keep you waiting)
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