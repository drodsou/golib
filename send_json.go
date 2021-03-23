package lib

import (
	"fmt"
	"net/http"
	"strconv"
)

const api_version = 1

// SendJSON fn
func SendJSON(w http.ResponseWriter, status int, statusDesc string, data string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status) // important: last one after other headerotros header

	j := `{Status: ` + strconv.Itoa(status) + `, StatusDesc:` + statusDesc + `, Version:` + api_version + `, Data:` + data + `}`

	fmt.Fprintf(w, j)
}
