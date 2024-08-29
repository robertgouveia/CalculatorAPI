package handlers

import (
	"Calculator/structs"
	"encoding/json"
	"fmt"
	"net/http"
)

func Sum(w http.ResponseWriter, r *http.Request) {
	var Data structs.ArrayData
	json.NewDecoder(r.Body).Decode(&Data)

	sum := 0
	for _, num := range Data.Numbers {
		sum += num
	}

	fmt.Fprintf(w, "%d\n", sum)
}
