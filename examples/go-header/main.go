package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	h := NewAnonymouse()

	// Automatically populate from incoming HTTP headers
	h.PopulateFromHTTP(r.Header)

	pos, err := h.Position()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("Received Position: %+v\n", pos)

	// Modify headers
	pos.X += 1
	pos.Y += 2
	pos.Z += 3
	h.SetPosition(pos)

	h.WriteToHTTP(w)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Headers processed\n"))
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
