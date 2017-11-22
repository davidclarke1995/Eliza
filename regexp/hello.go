package hello

import (
    //"fmt"
    "math/rand"
    //"bufio"
    "time"
    "regexp"
	//"os"
	//"strings"	
)

func ElizaOutput(inputStr string) string{


	input := inputStr
	//pattern := `.*father.*`
	pattern := `(?!).*\bfather\b.*`
	//pattern2 := []string{`.*i am(.*)`, `.*I AM(.*)`, `.*I'm(.*)`, `.*i'm(.*)`, `.*im(.*)`, `.*I am(.*)`}
	output := "Why dont you tell me more about your father?"
	//response := ""
	rand.Seed(time.Now().UTC().UnixNano())



		answers := []string{//responses
        	"I’m not sure what you’re trying to say. Could you explain it to me?",
        	"How does that make you feel?",
        	"Why do you say that?",
        }

        if matched, _ := regexp.MatchString(pattern, input); matched {
        //returning the string below
        	output = "Input: " + input + "\nOutput :" + output

		} else {
			output = "Input: " + input + " \nOutput :" + answers[rand.Intn(len(answers))]
		}

		return answers[rand.Intn(len(answers))]
}