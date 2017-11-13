package main

import (
	"net/http"
	"log"
	"fmt"
	"regexp"
	"time"
	"math/rand"

)

type Response struct {
	pattern         *regexp.Regexp
	possibleAnswers []string
}

func Eliza(w http.ResponseWriter, r *http.Request){

	rand.Seed(time.Now().UnixNano())
	userGuess := r.URL.Query().Get("value")
	
	re := regexp.MustCompile("(.*?)")

	
	answers := []string{`I'm not sure what your trying to say. Could you explain it to me?`, `How does that make you feel?`, `Why do you say that?`}
	
	likeResponse := Response{re, answers} // Responses could be read from files.


	if likeResponse.pattern.MatchString(userGuess) {
		index := rand.Intn(len(likeResponse.possibleAnswers)) // index of random possible answer
		chosenResponse := likeResponse.possibleAnswers[index] // pick the answer
		finalResponse := fmt.Sprintf(chosenResponse)   // sub the "topic" into the response - since the response has a %s for this purpose
		fmt.Fprintf(w,finalResponse)                            // we actually want to display this in the html page or template or something.

		// The output will be "How long have you liked me?" or "Why do you like me?"
	} else {
		fmt.Fprintf(w,"I don't know what you said.") // there was no regex match so just say some default answer.
	}
	
	/*
	response := ""
	if (userGuess == "Hello"){
		//fmt.Fprintf(w, "Hello, %s!", userGuess)
		response = "Hello Keith, how are you?"
		
	}else{
		response = "Try saying Hello!"
		
	}
	if(strings.Contains(userGuess,"bad")){
		response = "Be happy"
		
	}else if(strings.Contains(userGuess,"good")){
		response = "Maaa G, Stress Less I love it hombre"
	}
	if(userGuess=="lets play"){
		
		response = ""
		fmt.Fprintf(w,"Hahah see what i Did: %s",gameGuess)
	
	}
	
	fmt.Fprintf(w,"%s",response)
	*/
	
}

func main() {
	
	//Handles static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/Eliza", Eliza)
    log.Println("Preparing guessing game , enter this in your web browser - Localhost:8080")
    http.ListenAndServe(":8080", nil)	
	
}