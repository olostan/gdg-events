package gdgreg

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/api/", handler)
	http.HandleFunc("/admin", redirectHandler)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Location", "admin/");
	w.WriteHeader(302);
	fmt.Fprint(w,"Redirecting...");
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
