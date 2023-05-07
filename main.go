package main

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/sethvargo/go-envconfig"
	"github.com/shoriwe/message-api/common/sqlite"
	"github.com/shoriwe/message-api/controller"
	"github.com/shoriwe/message-api/router"
	"github.com/shoriwe/message-api/session"
	"google.golang.org/api/option"
)

func main() {
	var env Environment
	pErr := envconfig.Process(context.Background(), &env)
	if pErr != nil {
		log.Fatal(pErr)
	}
	db := sqlite.New(env.Database)
	j := session.New([]byte(env.Secret))
	app, aErr := firebase.NewApp(
		context.Background(),
		&firebase.Config{ProjectID: env.Firebase.ProjectID},
		option.WithCredentialsFile(env.Firebase.Configuration),
	)
	if aErr != nil {
		log.Fatal(aErr)
	}
	c := controller.New(db, j, app)
	r := router.New(c, gin.Default())
	rErr := r.Run(os.Args[1:]...)
	if rErr != nil {
		log.Fatal(rErr)
	}
}
