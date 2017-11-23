//David Clarke G00335563
package hello
//imports
import (
    //"fmt"
    "math/rand"
    //"bufio"
    //"time"
    "regexp"
	//"os"
	"strings"	
)

func ElizaOutput(inputStr string) string{

	//converting input to string
	input := inputStr

	//Regular Expressions
	if strings.Contains(strings.ToLower(input), "hello"){
		return "Hi!! How are you??"
	}

	if matched,_ := regexp.MatchString(`(?i). *\bkevin\b.*`,input); matched {
		return "That guy is stupid..."
	}



	//array of random answers
	answers := []string{
		"Alright. Interesting....",
		"How does that make you feel?",
		"Why do you say that?",
		"Huh??",
		"I'm not even going to pretend that I know what you're trying to say...",
		 

	}
		//return answer if input does not match expressions
		return answers[rand.Intn(len(answers))]
}