package services

import (
	context "context"
	time "time"

	bson "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (v NotesService) GetSingleNote(NoteID string) *Models.Note {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ID, err := primitive.ObjectIDFromHex(NoteID)
	Helpers.PrintError(err)

	var note Models.Note
	err = NoteCollection.FindOne(ctx, bson.M{"_id": ID}).Decode(&note)
	if err == mongo.ErrNoDocuments {
		return nil
	}

	return &note
}
