package web_api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"test_cdl/count_palindrom"
)

type Palindrom struct {
	Min    int    `json:"min"`
	Max    int    `json:"max"`
	Result string `json:"result"`
}

var data = []Palindrom{
	{1, 10, ""},
}

// func main() {
// 	http.HandleFunc("/palindrom", GetPalindrom)
// 	http.ListenAndServe(":8000", nil)
// }

func GetPalindrom(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	if request.Method == "POST" {
		min := request.FormValue("min")
		max := request.FormValue("max")
		data[0].Min, _ = strconv.Atoi(min)
		data[0].Max, _ = strconv.Atoi(max)
		data[0].Result = count_palindrom.CountPalindrom(data[0].Min, data[0].Max)
		res, err := json.Marshal(data)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Write(res)
		return
	}
	http.Error(response, "400 bad request", http.StatusBadRequest)
}
