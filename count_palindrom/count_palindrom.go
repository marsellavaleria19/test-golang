package count_palindrom

import (
	"math"
	"strconv"
)

/**
function to get palindrom value
*/
func PalindromHandle(digit int) int {
	result := 0 // to save palindrom value
	currentDigit := digit
	for currentDigit > 0 { // while current digit grather than 0
		lastDigit := currentDigit % 10     // last digit mod 10
		currentDigit = currentDigit / 10   // current digit divide 10
		result = (result * 10) + lastDigit // var result 10 and plus last digit
	}
	return result
}

/**
function fot count palindrom
*/
func CountPalindrom(min int, max int) string {
	count := 0                                // to save palindrom count
	if min > 0 && max <= int(math.Pow10(9)) { // if min grather than 0 dan max smaller than 10 pow 9
		for i := min; i <= max; i++ {
			palindrom := PalindromHandle(i) // all i value become palindrom with call palindrom function
			if i == palindrom {             // check if i same with palindrom
				count++ // if true , then count become count + 1
			}
		}
		return strconv.Itoa(count)
	} else {
		return ""
	}
}

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Println("Input")
// 	text, _ := reader.ReadString('\n')
// 	splitMinMax := strings.Split(strings.TrimSpace(text), " ")
// 	min, _ := strconv.Atoi(splitMinMax[0])
// 	max, _ := strconv.Atoi(splitMinMax[1])
// 	fmt.Println(CountPalindrom(min, max))
// }
