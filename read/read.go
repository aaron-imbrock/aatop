package read

import (
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func File(f string) string {
	dat, err := os.ReadFile(f)
	check(err)
	return string(dat)
}
