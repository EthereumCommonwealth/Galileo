package blocks

import "net/http"

func Index(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText("405"), http.StatusMethodNotAllowed)
		return
	}

	bks, err := LastFiveBlocks()
}