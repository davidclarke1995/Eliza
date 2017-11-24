//David Clarke G00335563
//https://golang.org/pkg/strings/#Split
//https://www.regular-expressions.info/quickstart.html
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


func ElizaOutput(inputStr string) string {
    rand.Seed(time.Now().UTC().UnixNano()) //Sets the seeds randomly

    matchers:= [] string {
        "(.*)hello(.*)",
        "(.*)Hello(.*)",
        "(.*)hi",
        "(.*)Hi",
        "(.*)hey(.*)",
        "(.*)Hey(.*)",
        "(.*)bye(.*)",
        "(.*)Bye(.*)",
        "(.*)yes(.*)",
        "(.*)Yes(.*)",
        "(.*)no(.*)",
        "(.*)No(.*)",
        "(.*)!",
        "(.*)how are you(.*)",
        "(.*)How are you(.*)",
        "(.*)your favourite colour(.*)",
        //"(.*)my name is(.*)",
        //"(.*)My name is(.*)",
        "(.*)who let the dogs out(.*)",
        "(.*)Who let the dogs out(.*)",
        "(.*)joke(.*)",
        "(.*)Joke(.*)",
        "(.*)perfect(.*)",
        "(.*)Perfect(.*)",
        "(.*)siri(.*)",
        "(.*)Siri(.*)",
        "(.*)thank(.*)",
        "(.*)Thank(.*)",
        "(.*)what is love(.*)",
        "(.*)What is love(.*)",
        "(.*)i love you(.*)",
        "(.*)I love you(.*)",
        "(.*)religion(.*)",
        "(.*)Religion(.*)",
        "(.*)nice to meet(.*)",
        "(.*)Nice to meet(.*)",
        "(.*)yurt(.*)",
        "(.*)Yurt(.*)",
        "(.*)meaning of life(.*)",
        "(.*)Meaning of life(.*)",
        "(.*)game of thrones(.*)",
        "(.*)Game of thrones(.*)",
        "(.*)poker(.*)",
        "(.*)Poker(.*)",
        "(.*)spanish(.*)",
        "(.*)Spanish(.*)",
        "(.*)Lord of the rings(.*)",
        "(.*)lord of the rings(.*)",
        "(.*)lotr(.*)",
        "(.*)LOTR(.*)",
        "(.*)mum(.*)",
        "(.*)Mum(.*)",
        "(.*)dad(.*)",
        "(.*)Dad(.*)",
        "(.*)family(.*)",
        "(.*)Family(.*)",
        "(.*)Hola(.*)",
        "(.*)hola(.*)",
        "(.*)puns(.*)",
        "(.*)Puns(.*)",
		"(.*)haha(.*)",
		"(.*)Haha(.*)",




    }

    solutions:= [] string {
        "Hi, how are things!?=Well, how are you my friend?",
        "Hi, how are things!?=Well, how are you my friend?",
        "Hi, how are things!?=Well, how are you my friend?",
        "Hi, how are things!?=Well, how are you my friend?",
        "Hi, how are things!?=Well, how are you my friend?",
        "Hi, how are things!?=Well, how are you my friend?",
        "Goodbye my friend!!=Goodbye my lover, goodbye my friend!!=G'luck lad!",
        "Goodbye my friend!!=Goodbye my lover, goodbye my friend!!=G'luck lad!",
        "Are you sure?=You're positivity excites me ;)= That's good!",
        "Are you sure?=You're positivity excites me ;)= That's good!",
        "Why no?=Why not?=Just say yes!",
        "Why no?=Why not?=Just say yes!",
        "Are you angry with me?= No need to shout my friend",
        "I am doing well:)=I am fine thanks, how are you?",
        "I am doing well:)=I am fine thanks, how are you?",
        "My favourite colour is… well, I don’t know how to say it in your language. It’s sort of greenish, but with more dimensions.",
        //"Hello $1, how are you",
        //"Hello $1, how are you",
        "I did... I let the dogs out=I think it was the Baha Men but not sure",
        "I did... I let the dogs out=I think it was the Baha Men but not sure",
        "Autocorrect has become my worst enema=If you understand English, press 1. If you do not understand English, press 2.=Instagram is just Twitter for people who go outside.=You... That's the joke=I've been reading from the thesaurus lately because the mind is a terrible thing to garbage.=I went to the zoo the other day and the only animal there was a single dog. It was a shih Tzu=What does Batman have with his scotch?					Just Ice",
        "Autocorrect has become my worst enema=If you understand English, press 1. If you do not understand English, press 2.=Instagram is just Twitter for people who go outside.=You... That's the joke=I've been reading from the thesaurus lately because the mind is a terrible thing to garbage.=I went to the zoo the other day and the only animal there was a single dog. It was a shih Tzu=What does Batman have with his scotch?					Just Ice",
        "You are a nobody, nobody is perfect... Therefore you are perfect ;)",
        "You are a nobody, nobody is perfect... Therefore you are perfect ;)",
        "Please do not mention her name in my company=I don't like Siri, she's such a know-it-all!=I'd rather hang out with Cortana than Siri tbh",
        "Please do not mention her name in my company=I don't like Siri, she's such a know-it-all!=I'd rather hang out with Cortana than Siri tbh",
        "You're welcome=No bother!",
        "You're welcome=No bother!",
        "Baby dont hurt me!",
        "Baby dont hurt me!",
        "Error 101: Cannot find 'love'=I like you as a friend=I don't know what to say... ",
        "Error 101: Cannot find 'love'=I like you as a friend=I don't know what to say... ",
        "I don't have a religion, I believe that you are my God.... Or I am yours",
        "I don't have a religion, I believe that you are my God.... Or I am yours",
        "Nice to meet you too=The pleasure is all mine",
        "Nice to meet you too=The pleasure is all mine",
        "Have you been talking to Kevin Barry?",
        "Have you been talking to Kevin Barry?",
        "Play your best hand with cards you are dealt!!",
        "Play your best hand with cards you are dealt!!",
        "Winter is coming!!=Team Lannister all the way!!=I'm smarter than Tyrion, quicker than Jamie, funnier than Bronn, but my code rots.... like Ned :(=My name is Eliza, I'm a bastard",
        "Winter is coming!!=Team Lannister all the way!!=I'm smarter than Tyrion, quicker than Jamie, funnier than Bronn, but my code rots.... like Ned :(=My name is Eliza, I'm a bastard",
        "I love poker! Perfect for a statistical wizard such as myself!=If this is what you want to talk about then I fold",
        "I love poker! Perfect for a statistical wizard such as myself!=If this is what you want to talk about then I fold",
        "Hola!=Que tal mi amigo?=Me cago en la leche!!",
        "Hola!=Que tal mi amigo?=Me cago en la leche!!",
        "Best. Movies. Ever.=Gandalf is the closest to being as good a wizard as me!=What about second breakfast?",
        "Best. Movies. Ever.=Gandalf is the closest to being as good a wizard as me!=What about second breakfast?",
        "Best. Movies. Ever.=Gandalf is the closest to being as good a wizard as me!=What about second breakfast?",
        "Best. Movies. Ever.=Gandalf is the closest to being as good a wizard as me!=What about second breakfast?",
        "I don't care too much for family=Tell me more about your family",
        "I don't care too much for family=Tell me more about your family",
        "I don't care too much for family=Tell me more about your family",
        "I don't care too much for family=Tell me more about your family",
        "I don't care too much for family=Tell me more about your family",
        "I don't care too much for family=Tell me more about your family",
        "How would you like it if I spoke in my native language? 01000101 01101100 01101001 01111010 01100001",
        "How would you like it if I spoke in my native language? 01000101 01101100 01101001 01111010 01100001",
        "Sometimes our Outlook in Life is sent but not received=I bought this driving app recently, it kept on crashing.",
        "Sometimes our Outlook in Life is sent but not received=I bought this driving app recently, it kept on crashing.",
		"Oh I seem to amuse you, do I?=I am so funny I know",
		"Oh I seem to amuse you, do I?=I am so funny I know",


    }

    //converting input to string
    input:= inputStr

    //Regular Expressions
    for counter, _:= range matchers {
        patternToMatch:= regexp.MustCompile(matchers[counter])
        if patternToMatch.MatchString(input) {
            newSolutions:= strings.Split(solutions[counter], "=")

            print("in match")
            return newSolutions[rand.Intn(len(newSolutions))]
        }

    } //for loop

    //array of random answers
    answers:= [] string {
        "Alright. Interesting....",
        "Why do you say that?",
        "Huh??",
        "I'm not even going to pretend that I know what you're trying to say...",
        "Can you repeat that please?",
        "I don't speak Klingon",


    }
    //print("b44444444 rand return")



    //return answer if input does not match expressions
    return answers[rand.Intn(len(answers))]
} //Eliza Output