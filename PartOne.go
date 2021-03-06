package main

import (

	"fmt"
	"regexp"
	"time"
	"math/rand"
	"strings"

)

//Code based on https://gist.github.com/ianmcloughlin/c4c2b8dc586d06943f54b75d9e2250fe
func pronounRefect(input string) string {
	
	re := regexp.MustCompile(`\b`)
	tokens := re.Split(input, -1)
	
	// Reflection map to switch pro-nouns
	reflectionMap := [][]string{
		{`I`, `you`},
		{`me`, `you`},
		{`you`, `my`},
		{`my`, `your`},
		{`your`, `my`},
		{`am`,`are`},
	}

	//Looping through each word reflecting it if there's a match.
	for i, token := range tokens {
		for _, reflectionMap := range reflectionMap {
			if matched, _ := regexp.MatchString(reflectionMap[0], token); matched {
				tokens[i] = reflectionMap[1]
				break
			}
		}
	}
	
	//Putting the words back together.
	return strings.Join(tokens, ``)
}


func ElizaResponse(input string)string{

	rand.Seed(time.Now().UnixNano())
	
	if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, input); matched {
		return "Why don’t you tell me more about your father?"
	}
	
	re := regexp.MustCompile(`(?i).*\bI am\b|\bI'm\b|\bIm\b.*`)
	if matched := re.MatchString(input); matched {
		return re.ReplaceAllString(input, "How do you know you are $1")
	}
	
	answers := []string{`I'm not sure what your trying to say. Could you explain it to me?`, `How does that make you feel?`, `Why do you say that?`}
	
	return answers[rand.Intn(len(answers))]
	
}

func main() {
	
	fmt.Println("People say I look like both my mother and father.")
	fmt.Println(ElizaResponse("People say I look like both my mother and father."))
	fmt.Println()

	fmt.Println("Father was a teacher.")
	fmt.Println(ElizaResponse("Father was a teacher."))
	fmt.Println()

	fmt.Println("I was my father’s favourite.")
	fmt.Println(ElizaResponse("I was my father’s favourite."))
	fmt.Println()

	fmt.Println("I’m looking forward to the weekend.")
	fmt.Println(ElizaResponse("I’m looking forward to the weekend."))
	fmt.Println()

	fmt.Println("My grandfather was French!")
	fmt.Println(ElizaResponse("My grandfather was French!"))
	fmt.Println()
	
	fmt.Println("Im happy.")
	fmt.Println(ElizaResponse("I am happy."))
	fmt.Println()

	fmt.Println("I'm not happy with your responses.")
	fmt.Println(ElizaResponse("I am not happy with your responses."))
	fmt.Println()

	fmt.Println("I am not sure that you understand the effect that your questions are having on me.")
	fmt.Println(ElizaResponse("I am not sure that you understand the effect that your questions are having on me."))
	fmt.Println()
	
	fmt.Println("I am supposed to just take what you’re saying at face value?")
	fmt.Println(ElizaResponse("I am supposed to just take what you’re saying at face value?"))
	fmt.Println()
	
	//for reflectionMap
	fmt.Println(pronounRefect("I am very tired today.")) 
	fmt.Println(pronounRefect("I am your creator"))
	fmt.Println(pronounRefect("my favourite food is bread."))
	fmt.Println(pronounRefect("your creator was keith."))

	
}