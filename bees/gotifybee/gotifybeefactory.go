package gotifybee

import (
	"github.com/muesli/beehive/bees"
)

// GotifyBeeFactory is a factory for GotifyBees.
type GotifyBeeFactory struct {
	bees.BeeFactory
}

// New returns a new Bee instance configured with the supplied options.
func (factory *GotifyBeeFactory) New(name, description string, options bees.BeeOptions) bees.BeeInterface {
	bee := GotifyBee{
		Bee: bees.NewBee(name, factory.ID(), description, options),
	}
	bee.ReloadOptions(options)

	return &bee
}

// ID returns the ID of this Bee.
func (factory *GotifyBeeFactory) ID() string {
	return "gotifybee"
}

// Name returns the name of this Bee.
func (factory *GotifyBeeFactory) Name() string {
	return "Gotify"
}

// Description returns the description of this Bee.
func (factory *GotifyBeeFactory) Description() string {
	return "Sends a message"
}

// Image returns the filename of an image for this Bee.
func (factory *GotifyBeeFactory) Image() string {
	return factory.ID() + ".png"
}

// LogoColor returns the preferred logo background color (used by the admin interface).
func (factory *GotifyBeeFactory) LogoColor() string {
	return "#be1728"
}

// Actions describes the available actions provided by this Bee.
func (factory *GotifyBeeFactory) Actions() []bees.ActionDescriptor {
	actions := []bees.ActionDescriptor{
		{
			Namespace:   factory.Name(),
			Name:        "send",
			Description: "Sends a message",
			Options: []bees.PlaceholderDescriptor{
				{
					Name:        "serverurl",
					Description: "server url",
					Type:        "string",
					Mandatory:   true,
				},
				{
					Name:        "message",
					Description: "message to send",
					Type:        "string",
					Mandatory:   false,
				},
				{
					Name:        "title",
					Description: "message title",
					Type:        "string",
					Mandatory:   false,
				},
				{
					Name:        "token",
					Description: "token",
					Type:        "string",
					Mandatory:   true,
				},
			},
		},
	}
	return actions
}

func init() {
	f := GotifyBeeFactory{}
	bees.RegisterFactory(&f)
}
