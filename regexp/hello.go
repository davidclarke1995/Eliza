//David Clarke G00335563
package hello
//imports
import (
    //"fmt"
    "math/rand"
    //"bufio"
    "time"
    "regexp"
	//"os"
	"strings"	
)




func ElizaOutput(inputStr string) string{
rand.Seed(time.Now().UTC().UnixNano())//Sets the seeds randomly

	matchers := [] string{
		"(.*)hello(.*)",
		"(.*)bye(.*)",
		"(.*)yes(.*)",
	}

	solutions := [] string{
		"Hi, how are things!?=Well, how are you my friend?",
		"Goodbye my friend!!=Goodbye my lover, goodbye my friend!!=G'luck lad!",
		"Are you sure?=You're positivity excites me ;)",
	}

	//converting input to string
	input := inputStr


	//Regular Expressions
	for counter,_ :=range matchers {
		patternToMatch := regexp.MustCompile(matchers[counter])
	if patternToMatch.MatchString(input) {
		newSolutions:= strings.Split(solutions[counter], "=")

		print("iinnn maaaaaaaaaaaaaaaaaaaaaaaaaatccccch")
		return newSolutions[rand.Intn(len(newSolutions))]
		}

	}//for loop

	//array of random answers
	answers := []string{
		"Alright. Interesting....",
		"How does that make you feel?",
		"Why do you say that?",
		"Huh??",
		"I'm not even going to pretend that I know what you're trying to say...",
		 

	}
			print("b44444444 rand return")


		
		//return answer if input does not match expressions
		return answers[rand.Intn(len(answers))]
}