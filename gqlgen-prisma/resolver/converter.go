package resolver

import (
	"gqlgen-prisma/db"
	"gqlgen-prisma/graph/model"
)

func convertToUserResponse(user *db.UserModel) *model.User {
	if user == nil {
		return nil
	}

	return &model.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func convertToPostResponse(post *db.PostModel) *model.Post {
	if post == nil {
		return nil
	}

	return &model.Post{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
	}
}
