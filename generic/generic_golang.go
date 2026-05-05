package generic

import "fmt"

func Print(s []string) {
	for _, v := range s {
		fmt.Print(v)
	}
}
