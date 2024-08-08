package persistance

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func Initialaze() *firebase.App {

	opt := option.WithCredentialsFile("path/to/refreshToken.json")
	config := &firebase.Config{ProjectID: "my-project-id"}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return app
}
