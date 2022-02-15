package services

import (
	context "context"
	time "time"

	bson "go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"

	Database "NoteKeeperAPI/src/database"
	Models "NoteKeeperAPI/src/database/models"
	Helpers "NoteKeeperAPI/src/helpers"
)

var NoteCollection *mongo.Collection = Database.GetCollection(Database.DB, "notes")

type NotesService struct{}

func (v NotesService) GetNotes(UserID string) []Models.Note {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var notes []Models.Note

	cursor, err := NoteCollection.Find(ctx, bson.D{
		{Key: "Author", Value: UserID},
	})
	Helpers.PrintError(err)

	err = cursor.All(ctx, &notes)
	Helpers.PrintError(err)

	return notes
}
