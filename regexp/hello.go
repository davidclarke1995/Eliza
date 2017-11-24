//David Clarke G00335563
//https://golang.org/pkg/strings/#Split
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
		"(.*)no(.*)",
		"(.*)!",
		"how are you(.*)",
		"(.*)your favourite colour(.*)",
		"(.*)my name is(.*)",
		"(.*)who let the dogs out(.*)",
		"(.*)joke(.*)",
		"(.*)perfect(.*)",
		"(.*)siri(.*)",
		"(.*)thank(.*)",
		
		

		


	}

	solutions := [] string{
		"Hi, how are things!?=Well, how are you my friend?",
		"Goodbye my friend!!=Goodbye my lover, goodbye my friend!!=G'luck lad!",
		"Are you sure?=You're positivity excites me ;)= That's good!",
		"Why no?=Why not?=Just say yes!",
		"Are you angry with me?= No need to shout my friend",
		"I am good=I am fine thanks, how are you?",
		"My favourite colour is… well, I don’t know how to say it in your language. It’s sort of greenish, but with more dimensions.",
		"Hello $1, how are you",
		"I did... I let the dogs out=I think it was the Baha Men but not sure",
		"Autocorrect has become my worst enema=If you understand English, press 1. If you do not understand English, press 2.=Instagram is just Twitter for people who go outside.",
		"You are a nobody, nobody is perfect... Therefore you are perfect ;)",
		"Please do not mention her name in my company=I don't like Siri, she's such a know-it-all!=I'd rather hang out with Cortana than Siri tbh",
		"You're welcome=No bother!",
	}

	//converting input to string
	input := inputStr

	//Regular Expressions
	for counter,_ :=range matchers {
		patternToMatch := regexp.MustCompile(matchers[counter])
	if patternToMatch.MatchString(input) {
		newSolutions:= strings.Split(solutions[counter], "=")

		print("in match")
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
			//print("b44444444 rand return")


		
		//return answer if input does not match expressions
		return answers[rand.Intn(len(answers))]
}//Eliza Output