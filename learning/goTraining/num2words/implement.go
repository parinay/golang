// Num2words gives methods to give a string/words for numeric val
package num2words

var units = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

// Num2words()
func ToWords(number int) (result string) {

	if number >= 0 && number < 10 {
		return units[number]
	}
	return "Unknown number"
}
