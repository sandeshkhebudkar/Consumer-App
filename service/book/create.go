package book

import (
	"encoding/json"
	"fmt"

	"github.com/sandeshkhebudkar/Consumer-App/model"
	"go.mongodb.org/mongo-driver/mongo"
)

//Add book
func Add(client *mongo.Client, book []byte) {

	//request tp middleware

	var bookObj model.DBLayer

	err1 := json.Unmarshal(book, &bookObj)
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	fmt.Println(bookObj)
	//Call to save the book
	doc, err := model.Save(client, bookObj)
	if err != nil {

		fmt.Println("message: Sorry for incovenicne ", err)
	}

	fmt.Println("Document created id :", doc)

}
