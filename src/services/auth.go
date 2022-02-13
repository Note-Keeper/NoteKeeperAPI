package services

import (
	context "context"
	fmt "fmt"
	time "time"

	validator "github.com/go-playground/validator/v10"
	bson "go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"

	Database "NoteKeeperAPI/src/database"
	// Models "NoteKeeperAPI/src/database/models"
	Helpers "NoteKeeperAPI/src/helpers"
	Requests "NoteKeeperAPI/src/types/requests"
)

var AccountCollection *mongo.Collection = Database.GetCollection(Database.DB, "accounts")
var Validate = validator.New()

type AuthService struct{}

func (v AuthService) Register(userData Requests.Register) Requests.Register {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newUser, err := AccountCollection.InsertOne(ctx, bson.D{
		{Key: "Name", Value: userData.Name},
		{Key: "Email", Value: userData.Email},
		{Key: "Login", Value: userData.Login},
		{Key: "Password", Value: userData.Password},
	})
	Helpers.PrintError(err)

	fmt.Println(newUser)

	return userData
}
