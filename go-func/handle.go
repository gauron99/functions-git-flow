package function

import (
	"fmt"
	"net/http"
)

// Handle an HTTP Request.
func Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TESTING")
	fmt.Fprintf(w, "TESTING SEND")
}
