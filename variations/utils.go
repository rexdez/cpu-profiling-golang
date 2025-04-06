package variations

import (
	"fmt"
	"time"
	"os"
	"runtime/pprof"
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

func SaveProfile(filePath, profileType string) {
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating %s: %v\n", filePath, err)
		return
	}
	defer f.Close()

	p := pprof.Lookup(profileType)
	if p == nil {
		fmt.Printf("Profile type %s not found\n", profileType)
		return
	}
	p.WriteTo(f, 0)
}