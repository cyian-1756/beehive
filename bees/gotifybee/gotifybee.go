// Package gotifybee is a Bee that can interact with a gotify server.
package gotifybee

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/muesli/beehive/bees"
)

// GotifyBee is a Bee that can launch external processes.
type GotifyBee struct {
	bees.Bee

	eventChan chan bees.Event
}

// Action triggers the action passed to it.
func (mod *GotifyBee) Action(action bees.Action) []bees.Placeholder {
	outs := []bees.Placeholder{}

	switch action.Name {
	case "send":
		var title string
		var message string
		var serverurl string
		var token string
		action.Options.Bind("title", &title)
		action.Options.Bind("message", &message)
		action.Options.Bind("serverurl", &serverurl)
		action.Options.Bind("token", &token)
		mod.Logln("Server:" + serverurl)
		response, err := http.PostForm(serverurl+"/message?token="+token, url.Values{
			"title":   {title},
			"message": {message}})
		if err != nil {
			mod.LogErrorf("Error: %s", err)
			return outs
		}
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)

		if err != nil {
			mod.LogErrorf("Error: %s", err)
			return outs
		}

		fmt.Printf("%s\n", string(body))

	default:
		panic("Unknown action triggered in " + mod.Name() + ": " + action.Name)
	}

	return outs
}

// Run executes the Bee's event loop.
func (mod *GotifyBee) Run(eventChan chan bees.Event) {
	mod.eventChan = eventChan
}

// ReloadOptions parses the config options and initializes the Bee.
func (mod *GotifyBee) ReloadOptions(options bees.BeeOptions) {
	mod.SetOptions(options)
}
