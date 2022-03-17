package repository

import (
	"api/internal/entity"

	"context"
	"fmt"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoUserRepository struct {
	mu  sync.Mutex
	col *mongo.Collection
}

func New(collection *mongo.Collection) *mongoUserRepository {
	return &mongoUserRepository{col: collection}
}

func (r *mongoUserRepository) Set(ctx context.Context, u *entity.User) (user *entity.User, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	res, err := r.col.InsertOne(ctx, u)
	if err != nil {
		return nil, err
	}
	u.UUID = res.InsertedID.(primitive.ObjectID)
	user = u
	return user, nil
}

func (r *mongoUserRepository) Delete(ctx context.Context, uid string) (err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	id, _ := primitive.ObjectIDFromHex(uid)

	var user entity.User

	err = r.col.FindOne(ctx, bson.D{{"_id", id}}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return err
		}
		log.Fatal("FindOne ", err)
		return err
	}
	for _, friend := range user.Friends {
		id, _ := primitive.ObjectIDFromHex(friend)
		fmt.Println(id)
		_, err := r.col.UpdateOne(ctx, bson.D{{"_id", id}}, bson.D{{"$pull", bson.D{{"friends", uid}}}})
		if err != nil {
			if err == mongo.ErrNoDocuments {
			}
			log.Fatal("UpdateOne", err)
		}

	}

	if _, err := r.col.DeleteOne(ctx, bson.D{{"_id", id}}); err != nil {
		return err
	}
	return nil
}

func (r *mongoUserRepository) Update(ctx context.Context, uid string, na string) (err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	id, _ := primitive.ObjectIDFromHex(uid)
	_, err = r.col.UpdateOne(ctx, bson.D{{"_id", id}}, bson.D{{"$set", bson.D{{"age", na}}}})
	if err != nil {
		return err
	}

	return nil

}

func (r *mongoUserRepository) MakeFriends(ctx context.Context, suid string, tuid string) (err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	sid, _ := primitive.ObjectIDFromHex(suid)

	tid, _ := primitive.ObjectIDFromHex(tuid)

	updateRes, err := r.col.UpdateOne(ctx, bson.M{"_id": sid}, bson.M{"$addToSet": bson.M{"friends": tuid}})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(updateRes.UpsertedID)
	_, err = r.col.UpdateOne(ctx, bson.D{{"_id", tid}}, bson.D{{"$addToSet", bson.D{{"friends", suid}}}})

	return nil
}

func (r *mongoUserRepository) GetByUUID(ctx context.Context, uid string) (*entity.User, error) {

	id, _ := primitive.ObjectIDFromHex(uid)

	var user entity.User

	err := r.col.FindOne(ctx, bson.D{{"_id", id}}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		log.Fatal(err)
		return nil, err
	}

	return &user, nil
}

func (r *mongoUserRepository) GetAll(ctx context.Context) ([]*entity.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var users []*entity.User

	cursor, err := r.col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user entity.User
		if err = cursor.Decode(&user); err != nil {
			log.Fatal(err)
		}
		users = append(users, &user)
	}
	return users, nil
}

func (r *mongoUserRepository) GetAllFriends(ctx context.Context, uid string) ([]*entity.Friends, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var friends []*entity.Friends
	user, err := r.GetByUUID(ctx, uid)
	for _, val := range user.Friends {
		friend, _ := r.GetByUUID(ctx, val)
		result := entity.Friends{UUID: friend.UUID, Name: friend.Name, Age: friend.Age}
		friends = append(friends, &result)
	}
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return friends, nil
}

/*
func remove(su map[uint64]*entity.User, tuid uint64) {
	for i := range su {
		for j := range su[i].Friends {
			if su[i].Friends[j].UUID == tuid {
				su[i].Friends[j] = su[i].Friends[len(su[i].Friends)-1]
				su[i].Friends[len(su[i].Friends)-1] = nil
				su[i].Friends = su[i].Friends[:len(su[i].Friends)-1]
				fmt.Println(true)
				break
			}
		}
	}
}
*/
