package main

import (
	"google.golang.org/api/option"
	"firebase.google.com/go"
	"context"
	"log"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"google.golang.org/api/iterator"
	"encoding/json"
	"github.com/mitchellh/mapstructure"
)
type Book struct {
	Name string `firestore:"name"`
	Created_at time.Time `firestore:"Created_at"`
	Owner string `firestore:"owner"`
}

func MapToStruct(m map[string]interface{}, val interface{}) error {
	tmp, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tmp, val)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	opt := option.WithCredentialsFile("serviceAccountKey.json")
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil{
		log.Fatalf("error initializing app: %v\n", err)
	}
	fmt.Println("initializing app successed.")

	client, err := app.Firestore(ctx)
	if err != nil{
		log.Fatalf("error initializng Firestore: %v\n",err)
	}
	defer client.Close()
	fmt.Printf("initializing Firestore succcess: %v\n",client)


	var books []Book
	var book Book
	item := client.Collection("books").Documents(ctx)
	for {
		doc, err := item.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v<br>", err)
		}
		//fmt.Printf("%v<br>", doc.Data())
		err = mapstructure.Decode(doc.Data(),&book)
		if err != nil {
			panic(err)
		}

		//fmt.Printf("%#v", book)
		books = append(books,book)
	}
	fmt.Printf("%v\n",books)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.tmpl")
	//router.GET("/index",getIndex)
	router.Run(":8080")

}

