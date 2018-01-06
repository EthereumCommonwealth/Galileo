package main

import (
	"net/http"
	"github.com/eduartua/callisto/Galileo/blocks"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/blocks", blocks.Index)
}

func index(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/blocks", http.StatusSeeOther)
}