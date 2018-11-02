package main

import (
	"google.golang.org/api/option"
	"firebase.google.com/go"
	"context"
	"log"
	"fmt"
)

func main() {
	opt := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil{
		log.Fatalf("error initializing app: %v\n", err)
	}
	fmt.Println(app)
}