package api

import (
	"context"

	"github.com/shellbear/go-graphql-example/configs"
	"github.com/shellbear/go-graphql-example/models"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input models.NewUser) (*models.User, error) {
	user := models.User{Name: input.Name}

	err := configs.DB.Create(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input models.NewTodo) (*models.Todo, error) {
	userID, err := models.UnmarshalID(input.UserID)

	if err != nil {
		return nil, err
	}

	var user models.User

	err = configs.DB.Find(&user, userID).Error

	if err != nil {
		return nil, err
	}

	todo := models.Todo{Text: input.Text, Done: false, User: &user}

	err = configs.DB.Create(&todo).Error

	if err != nil {
		return nil, err
	}

	return &todo, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	var todods []*models.Todo
	err := configs.DB.Find(&todods).Error

	return todods, err
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	err := configs.DB.Find(&users).Error

	return users, err
}
