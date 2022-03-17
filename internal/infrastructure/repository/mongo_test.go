package repository

import (
	"api/internal/entity"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"testing"
)

var userCollection *mongo.Collection

func NewTest(collection *mongo.Collection) *mongoUserRepository {
	return &mongoUserRepository{col: collection}
}

func TestInsertOne(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	ctx := context.Background()
	r := NewTest(userCollection)
	mt.Run("success", func(mt *mtest.T) {
		r.col = mt.Coll
		id := primitive.NewObjectID()
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		insertedUser, err := r.Set(ctx, &entity.User{
			UUID: id,
			Name: "john",
			Age:  "35",
		})
		assert.Nil(t, err)
		assert.Equal(t, &entity.User{
			UUID: id,
			Name: "john",
			Age:  "35",
		}, insertedUser)
	})

	mt.Run("custom error duplicate", func(mt *mtest.T) {
		userCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "duplicate key error",
		}))

		insertedUser, err := r.Set(ctx, &entity.User{})

		assert.Nil(t, insertedUser)
		assert.NotNil(t, err)
	})

	mt.Run("simple error", func(mt *mtest.T) {
		userCollection = mt.Coll
		mt.AddMockResponses(bson.D{{"ok", 0}})

		insertedUser, err := r.Set(ctx, &entity.User{})

		assert.Nil(t, insertedUser)
		assert.NotNil(t, err)
	})
}

func TestFindOne(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	ctx := context.Background()
	r := NewTest(userCollection)
	mt.Run("success", func(mt *mtest.T) {
		r.col = mt.Coll
		expectedUser := entity.User{
			UUID: primitive.NewObjectID(),
			Name: "john",
			Age:  "35",
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{"_id", expectedUser.UUID},
			{"name", expectedUser.Name},
			{"age", expectedUser.Age},
		}))

		fmt.Println(expectedUser.UUID.Hex())
		userResponse, err := r.GetByUUID(ctx, expectedUser.UUID.Hex())
		fmt.Println(userResponse.UUID.Hex())
		assert.Nil(t, err)
		assert.Equal(t, &expectedUser, userResponse)
	})
}
