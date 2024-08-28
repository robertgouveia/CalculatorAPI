package handlers

import (
	"Calculator/structs"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func Add(w http.ResponseWriter, r *http.Request) {
	var Data structs.ResponseData

	err := json.NewDecoder(r.Body).Decode(&Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := strconv.Itoa(Data.Number1 + Data.Number2)

	fmt.Fprintf(w, "%s\n", result)
}
