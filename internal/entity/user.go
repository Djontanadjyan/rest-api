package entity

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	UUID    primitive.ObjectID `json:"id" bson:"_id,omitempty" :"UUID"`
	Name    string             `json:"name" bson:"name"  :"name"`
	Age     string             `json:"age" bson:"age" :"age"`
	Friends []string           `json:"friends,omitempty" bson:"friends,omitempty"  :"friends"`
}

type Friends struct {
	UUID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name,omitempty" bson:"name,omitempty"`
	Age  string             `json:"age,omitempty" bson:"age,omitempty"`
}

type FriendPool struct {
	SourceID string `json:"sourceID,omitempty"`
	TargetID string `json:"targetID,omitempty"`
}

type UserUseCase interface {
	GetByUUID(ctx context.Context, uid string) (*User, error)
	Update(ctx context.Context, u string, na string) error
	Delete(ctx context.Context, uid string) error
	GetAll(ctx context.Context) ([]*User, error)
	Set(ctx context.Context, u *User) (*User, error)
	MakeFriends(ctx context.Context, suid string, tuid string) error
	GetAllFriends(ctx context.Context, uid string) ([]*Friends, error)
}

type UserRepository interface {
	GetByUUID(ctx context.Context, uid string) (*User, error)
	Set(ctx context.Context, u *User) (*User, error)
	GetAll(ctx context.Context) ([]*User, error)
	Delete(ctx context.Context, uid string) error
	Update(ctx context.Context, u string, na string) error
	MakeFriends(ctx context.Context, suid string, tuid string) error
	GetAllFriends(ctx context.Context, uid string) ([]*Friends, error)
}
