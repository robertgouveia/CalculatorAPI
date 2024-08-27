package handlers

import (
	"Calculator/structs"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var Data structs.ResponseData

	err := json.NewDecoder(r.Body).Decode(&Data)

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	result := strconv.Itoa(Data.Number1 + Data.Number2)

	fmt.Fprintf(w, "%s\n", result)
}
