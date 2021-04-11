package views

import (
	"time"

	"slashbase.com/backend/models/user"
)

type UserView struct {
	ID              string    `json:"id"`
	Email           string    `json:"email"`
	Name            *string   `json:"name"`
	ProfileImageURL *string   `json:"profileImageUrl"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

type UserSessionView struct {
	ID        string    `json:"id"`
	User      UserView  `json:"user"`
	Token     string    `json:"token"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func BuildUser(usr *user.User) UserView {
	userView := UserView{
		ID:              usr.ID,
		Name:            nil,
		Email:           usr.Email,
		ProfileImageURL: nil,
		CreatedAt:       usr.CreatedAt,
		UpdatedAt:       usr.UpdatedAt,
	}
	if usr.ProfileImageURL.Valid {
		userView.ProfileImageURL = &usr.ProfileImageURL.String
	}
	if usr.FullName.Valid {
		userView.Name = &usr.FullName.String
	}
	return userView
}

func BuildUserSession(userSession *user.UserSession) UserSessionView {
	userSessView := UserSessionView{
		ID:        userSession.ID,
		User:      BuildUser(&userSession.User),
		Token:     userSession.GetAuthToken(),
		IsActive:  userSession.IsActive,
		CreatedAt: userSession.CreatedAt,
		UpdatedAt: userSession.UpdatedAt,
	}
	return userSessView
}
