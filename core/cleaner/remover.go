package cleaner

import (
	"fmt"
	"os"
)

// FilePath is a filename full with path
var FilePath = make(chan string, 1000)

// RmFile is a worker base to remove a file
func RmFile() {
	if len(FilePath) == 0 {
		return
	}
	for f := range FilePath {
		err := os.Remove(f)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
