package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

/*
Based on the code from https://developers.google.com/drive/v3/web/quickstart/go,
which is lisenced under Apache 2.0 License.
*/

func main() {
	var clientSecret string
	var scope string
	flag.StringVar(&clientSecret, "c", "client_secret.json", "JSON file for client secret")
	flag.StringVar(&scope, "s", "https://www.googleapis.com/auth/drive.file", "scope you want to have")

	flag.Parse()

	cred, err := ioutil.ReadFile(clientSecret)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(cred, scope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	// token from web
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Fprintf(os.Stderr, "Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}

	json.NewEncoder(os.Stdout).Encode(tok)
}
