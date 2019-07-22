package main

import (
	"log"
	"net/http" // http package
	"os"

	"github.com/dl4ab/timebot/api"
	"github.com/dl4ab/timebot/api/slack"
)

/*
	look up environment variable existence
*/
func mustLookupEnv(env string) string {
	ret, ok := os.LookupEnv(env)

	if !ok {
		log.Fatalf("Environment Variable %v is not available!\n", env)
	}

	return ret
}

func main() {
	// The env PORT is needed for Heroku
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	slackToken := mustLookupEnv("SLACK_SIGNING_SECRET")
	slackBotOAuthToken := mustLookupEnv("SLACK_BOT_OAUTH_TOKEN")

	// slack is just a struct to hold slack related tokens
	app := slack.New(slackToken, slackBotOAuthToken)

	// http.ListenAndServe creates HTTP server
	// 	- need to specify a port to listen to
	//	- need to specify http.NewServeMux() instance to map current request to the designated route 
	// check api package in github.com/dl4ab/timebot/ for more details
	// check more about it [https://golang.org/pkg/net/http/]
	log.Printf("[MAIN] The server is running at 0.0.0.0:%v\n", port)
	log.Println("[MAIN]", http.ListenAndServe(":"+port, api.GetRouter(app)))
}
