package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"test_cdl/count_palindrom"
	"test_cdl/web_api"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input")
	text, _ := reader.ReadString('\n')
	splitMinMax := strings.Split(strings.TrimSpace(text), " ")
	min, _ := strconv.Atoi(splitMinMax[0])
	max, _ := strconv.Atoi(splitMinMax[1])
	fmt.Println(count_palindrom.CountPalindrom(min, max))

	http.HandleFunc("/palindrom", web_api.GetPalindrom)
	http.ListenAndServe(":8000", nil)
}
