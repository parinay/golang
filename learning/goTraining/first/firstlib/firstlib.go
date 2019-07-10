// library to check the guess
package firstlib

// Checkguess checks to see if you won or else returns a message
func Checkguess(guess int, whatnumber int) (string, bool) {
	if guess == whatnumber {
		return "You guessed it!", true
	}
	if guess < whatnumber {
		return "Fizz", false
	} else {
		return "Buzz", false
	}
}
