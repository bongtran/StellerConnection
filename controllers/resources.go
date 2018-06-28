package controllers

import (
	"StellarConnection/model"
)

//Models for JSON resources
type (
	// UserResource For Post - /user/register
	UserResource struct {
		Data UserModel `json:"data"`
	}
	// AuthUserResource Response for authorized user Post - /user/login
	AuthUserResource struct {
		Data AuthUserModel `json:"data"`
	}
	// BookmarkResource For Post/Put - /bookmarks
	// For Get - /bookmarks/id
	BookmarkResource struct {
		Data model.Bookmark `json:"data"`
	}
	// BookmarksResource for Get - /bookmarks
	BookmarksResource struct {
		Data []model.Bookmark `json:"data"`
	}

	TransferResource struct {
		Data model.TransferModel `json:"data"`
	}

	OfferResource struct {
		Data model.OfferModel `json:"data"`
	}

	// UserModel reperesents a user
	UserModel struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}
	// AuthUserModel for authorized user with access token
	AuthUserModel struct {
		User  model.User `json:"user"`
		Token string     `json:"token"`
	}

	ResponseModel struct {
		StatusCode int    `json:"code"`
		Data       string `json:"d"`
		Error      string `json:"error"`
	}
)
