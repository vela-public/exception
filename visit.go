package exception

import (
	"fmt"
)

func Try(e error, protect bool) {
	if e == nil {
		return
	}

	if protect {
		defer func() {
			if cause := recover(); cause == nil {
				return
			} else {
				fmt.Printf("%s\n", StackTrace(0))
			}
		}()
	}

	panic(e)
}
