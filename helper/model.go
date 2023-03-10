package helper

import (
	"pznrestfulapi/model/domain"
	"pznrestfulapi/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToLoginResponse(user domain.User) web.LoginResponse {
	return web.LoginResponse{
		Id:       user.Id,
		Username: user.Username,
		Token:    user.Token,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}
