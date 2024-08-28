package handlers

import (
	"Calculator/structs"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func Subtract(w http.ResponseWriter, r *http.Request) {
	var Data structs.ResponseData

	json.NewDecoder(r.Body).Decode(&Data)
	result := strconv.Itoa(Data.Number1 - Data.Number2)

	fmt.Fprintf(w, "%s\n", result)
}
