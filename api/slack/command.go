package slack

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dl4ab/timebot/timebot"
)

// CommandHandler handles slack slash command
//
// ENDPOINT /api/slack/command
//
// Example Usage
// 	"/time 2018-12-31 21:40 PST" => "2019-01-01 14:40 KST"

// http.ResponseWriter [https://golang.org/pkg/net/http/#ResponseWriter]
//	- to construct an HTTP response
// http.Request [https://golang.org/pkg/net/http/#Request]
//	- request received by a server or to be sent by a client
func (app *App) CommandHandler(w http.ResponseWriter, r *http.Request) {

	// VerifyRequest is defined in http pacakage
	//	- check if X-Slack-Signature is valid by eventually calling "hmac.Equal([]byte(calculatedMAC), []byte(receivedMAC))"
	if !app.TestMode && !VerifyRequest(r, []byte(app.SigningToken)) {
		log.Printf("Slack signature not verifed")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// ??
	err := r.ParseForm()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// get the requested text like "/time 2018-12-31 21:40 PST"
	text := r.PostFormValue("text")

	if text == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// get the output value like "2019-01-01 14:40 KST"
	ret, err := timebot.ParseAndFlipTz(text)

	if err != nil {
		fmt.Fprintf(w, "%v is not a valid date time: %s", text, err)
		return
	}

	resp := Response{
		Text:         ret,
		ResponseType: Ephemeral,
	}

	w.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
