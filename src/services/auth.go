package services

import (
	context "context"
	time "time"

	bson "go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"

	Database "NoteKeeperAPI/src/database"
	// Models "NoteKeeperAPI/src/database/models"
	Helpers "NoteKeeperAPI/src/helpers"
	Requests "NoteKeeperAPI/src/types/requests"
)

var AccountCollection *mongo.Collection = Database.GetCollection(Database.DB, "accounts")

type AuthService struct{}

func (v AuthService) Register(userData Requests.Register) Requests.Register {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := AccountCollection.InsertOne(ctx, bson.D{
		{Key: "Name", Value: userData.Name},
		{Key: "Email", Value: userData.Email},
		{Key: "Login", Value: userData.Login},
		{Key: "Password", Value: userData.Password},
	})
	Helpers.PrintError(err)

	return userData
}

func (v AuthService) ValidateRegister(userData Requests.Register) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	checkEmail, err := AccountCollection.CountDocuments(ctx, bson.M{"Email": userData.Email})
	Helpers.PrintError(err)

	checkLogin, err := AccountCollection.CountDocuments(ctx, bson.M{"Login": userData.Login})
	Helpers.PrintError(err)

	return checkEmail == 0 && checkLogin == 0
}
