package models

import (
	"github.com/gorilla/schema"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Decoder = schema.NewDecoder()


type User struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email         string `json:"email,omitempty" bson:"email,omitempty"`
	Notifications []Transaction `json:"notifications,omitempty" bson:"notifications,omitempty"`
}

type Transaction struct {
	ID            			primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Receiver				string 	`json:"receiver" bson:"receiver"`
	Sender					string 	`json:"sender" bson:"sender"`
	Date					string 	`json:"date" bson:"date"`
	Read 					bool 	`json:"read" bson:"read"`
	Amount 					float64 `json:"amount" bson:"amount"`
}


type MarkMessage struct {
	Email 		string `json:"email"`
	MsgID 		string `json:"msg_id"`
}

type CustomerServiceMessage struct {
	ID            		primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Emailer 			string `json:"emailer,omitempty" bson:"emailer,omitempty"`
	Topic				string `json:"topic,omitempty" bson:"_id,omitempty"`
	Message				string `json:"msg,omitempty" bson:"msg,omitempty"`
	SentAt				string `json:"sentAt,omitempty" bson:"sentAt,omitempty"` 
}