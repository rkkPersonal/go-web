package file

import (
	"os"
)

func OpenFile() {
	buf := make([]byte, 1024)
	f, _ := os.Open("D:\\debug.log")
	defer f.Close()
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break

		}
		os.Stdout.Write(buf[:n])
	}
}
