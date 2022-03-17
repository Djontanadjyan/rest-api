package usecase

import (
	"context"
	"fmt"

	"api/internal/entity"
)

type userUseCase struct {
	userRepo entity.UserRepository
}

func New(u entity.UserRepository) entity.UserUseCase {
	return &userUseCase{
		userRepo: u,
	}
}

func (u *userUseCase) MakeFriends(ctx context.Context, suid string, tuid string) error {
	return u.userRepo.MakeFriends(ctx, suid, tuid)
}

func (u *userUseCase) Update(ctx context.Context, uid string, na string) error {
	return u.userRepo.Update(ctx, uid, na)
}

func (u *userUseCase) Delete(ctx context.Context, uid string) error {
	err := u.userRepo.Delete(ctx, uid)
	if err != nil {
		return err
	}
	return nil
}

func (u *userUseCase) GetAll(ctx context.Context) ([]*entity.User, error) {
	users, err := u.userRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userUseCase) GetAllFriends(ctx context.Context, uid string) ([]*entity.Friends, error) {
	friends, err := u.userRepo.GetAllFriends(ctx, uid)
	if err != nil {
		return nil, err
	}
	return friends, nil
}

func (u *userUseCase) GetByUUID(ctx context.Context, id string) (*entity.User, error) {
	user, err := u.userRepo.GetByUUID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUseCase) Set(ctx context.Context, user *entity.User) (*entity.User, error) {
	fmt.Println("Set usecase", user.Name, user.Age)
	user, err := u.userRepo.Set(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
