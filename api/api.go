package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/dl4ab/timebot/api/slack"
)

// GetRouter returns a root router for everything
// HandleFunc method in Router requires a function in a form of 
//	"func(http.ResponseWriter, *http.Request)" as the 2nd parameter.
// so, healthcheckHandler, app.CommandHandler, and EventHandler should be written in that form.
func GetRouter(app slack.App) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", healthcheckHandler).Methods("GET")
	// Handles a slash command "/time 2019-01-01 PST"
	r.HandleFunc("/api/slack/command", app.CommandHandler).Methods("POST")
	// Handles Slack Event Subscription
	r.HandleFunc("/api/slack/event", app.EventHandler).Methods("POST")
	return r
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
