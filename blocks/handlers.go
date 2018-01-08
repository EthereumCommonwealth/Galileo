package blocks

import "net/http"

//All data fetched by this functions come from database Cassandra

func Index(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	//logic code here
}

func BlockInfo(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	//logic code here
}

func TxInfo(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	//logic code here
}

func AddressId(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	//logic code here
}
