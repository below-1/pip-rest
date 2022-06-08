package siswa

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List siswa")
}
