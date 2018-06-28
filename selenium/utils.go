package selenium

import (
	"fmt"
	"io"
)

func Log(r io.Reader) {
	buf := make([]byte, 80)
	for {
		n, err := r.Read(buf)
		if n > 0 {
			fmt.Println(string(buf[0:n]))
		}
		if err != nil {
			break
		}
	}
}
