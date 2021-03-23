// derived from https://github.com/jeromer/mumbojumbo

package goconceal

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unsafe"
)

// CreateFile fn
func CreateFile(pkg, str, fileName string) {

	fileContent := fmt.Sprintf(`
		package %s
		import "unsafe"
		func eax() uint8{
			return uint8(unsafe.Sizeof(true))
		}
		// Unconceal fn
		func Unconceal() string {
			%s
		}
  	`, pkg, obfuscate(str))

	f, _ := os.Create(fileName)
	defer f.Close()
	_, _ = f.WriteString(fileContent)
}

// ---- HELPERS

// https://github.com/GH0st3rs/obfus/blob/master/obfus.go

func obfuscate(txt string) string {
	lines := []string{}

	for _, item := range []byte(txt) {
		lines = append(
			lines, getNumber(item),
		)
	}

	return fmt.Sprintf(
		"return string(\n[]byte{\n%s,\n},\n)",
		strings.Join(lines, ",\n"),
	)
}

const (
	// EAX const
	EAX = uint8(unsafe.Sizeof(true))
	// ONE const
	ONE = "eax()"
)

func getNumber(n byte) (buf string) {
	var arr []byte
	var x uint8

	for n > EAX {
		x = 0

		if n%2 == EAX {
			x = EAX
		}

		arr = append(arr, x)

		n = n >> EAX
	}

	buf = ONE

	rand.Seed(
		time.Now().Unix(),
	)

	for i := len(arr) - 1; i >= 0; i-- {
		buf = fmt.Sprintf(
			"%s<<%s", buf, ONE,
		)

		if arr[i] == EAX {
			op := "(%s|%s)"

			if rand.Intn(2) == 0 {
				op = "(%s^%s)"
			}

			buf = fmt.Sprintf(
				op, buf, ONE,
			)
		}
	}

	return buf
}
