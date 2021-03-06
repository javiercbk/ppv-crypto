package auth

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/javiercbk/ppv-crypto/server/http/security"
	"github.com/javiercbk/ppv-crypto/server/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"golang.org/x/crypto/bcrypt"
)

// ErrDuplicatedUser is returned when attempting to create a duplicated user
type ErrDuplicatedUser string

func (e ErrDuplicatedUser) Error() string {
	return string(e)
}

// NameTooLongErr is returned when attempting to store a first or last name that is too long
type NameTooLongErr string

func (e NameTooLongErr) Error() string {
	return string(e)
}

// EmailTooLongErr is returned when attempting to store an email that is too long
type EmailTooLongErr string

func (e EmailTooLongErr) Error() string {
	return string(e)
}

// BadCredentialsErr is returned when attempting to login with invalid credentials
type BadCredentialsErr string

func (e BadCredentialsErr) Error() string {
	return string(e)
}

// UserNotExistErr is thrown when the user does not exist anymore
type UserNotExistErr string

func (e UserNotExistErr) Error() string {
	return string(e)
}

const (
	nameConstraint        = "user_name_cnst"
	emailConstraint       = "user_email_cnst"
	emailUniqueConstraint = "users_email_idx"
	// ErrNameTooLong is returned when attempting to store a first or last name that is too long
	ErrNameTooLong NameTooLongErr = "either first or last name is too long"
	// ErrEmailTooLong is returned when attempting to store an email that is too long
	ErrEmailTooLong EmailTooLongErr = "email is too long"
	// ErrBadCredentials is returned when attempting to login with invalid credentials
	ErrBadCredentials BadCredentialsErr = "bad credentials"
	// ErrUserNotExist is thrown when the user does not exist anymore
	ErrUserNotExist UserNotExistErr = "bad credentials"
)

// Credentials has all the data necesary to authenticate an admin
type Credentials struct {
	Email    string `json:"email" validate:"required,gt=0,lte=256"`
	Password string `json:"password" validate:"required,gt=0"`
}

// TokenResponse contains a jwt token
type TokenResponse struct {
	User  VisibleUser `json:"user"`
	Token string      `json:"token"`
}

// VisibleUser is the public data of an admin
type VisibleUser struct {
	ID          int64                       `json:"id"`
	FirstName   string                      `json:"firstName"`
	LastName    string                      `json:"lastName"`
	Expiry      time.Time                   `json:"validUntil"`
	Permissions models.PermissionsUserSlice `json:"permissions"`
	CreatedAt   time.Time                   `json:"createdAt"`
	UpdatedAt   *time.Time                  `json:"updatedAt,omitempty"`
}

// APIFactory is a function capable of creating an Auth API
type APIFactory func(logger *log.Logger, db *sql.DB, jwtSecret string) API

// API is authentication API interface
type API interface {
	AuthenticateUser(ctx context.Context, credentials Credentials) (TokenResponse, error)
	UserInfo(ctx context.Context, jwtUser security.JWTUser, user *VisibleUser) error
}

type api struct {
	logger    *log.Logger
	db        *sql.DB
	jwtSecret string
}

// NewAPI creates a new authentication API
func NewAPI(logger *log.Logger, db *sql.DB, jwtSecret string) API {
	return api{
		logger:    logger,
		db:        db,
		jwtSecret: jwtSecret,
	}
}

// AuthenticateUser authenticates a user and returns a token
func (api api) AuthenticateUser(ctx context.Context, credentials Credentials) (TokenResponse, error) {
	tokenResponse := TokenResponse{}
	user, err := models.Users(
		qm.Where("email = ?", credentials.Email),
		qm.Load(models.UserRels.PermissionsUsers),
	).One(ctx, api.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// if the admin does not exist, in this API method return unvalid credentials
			return tokenResponse, ErrBadCredentials
		}
		api.logger.Printf("error retrieving admin for update: %v\n", err)
		return tokenResponse, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		return tokenResponse, ErrBadCredentials
	}
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)
	jwtUser := security.JWTUser{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Permissions: security.ToPermissionsMap(user.R.PermissionsUsers),
	}
	// session last an hour
	token.Claims = security.JWTEncode(jwtUser, time.Hour*1)
	t, err := token.SignedString([]byte(api.jwtSecret))
	if err != nil {
		api.logger.Printf("error signing token %v\n", err)
		return tokenResponse, errors.New("error creating token")
	}
	tokenResponse.Token = t
	tokenResponse.User = VisibleUser{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Permissions: user.R.PermissionsUsers,
		CreatedAt:   user.CreatedAt.Time,
		UpdatedAt:   user.UpdatedAt.Ptr(),
	}
	return tokenResponse, nil
}

// UserInfo returns a visible user from a userID
func (api api) UserInfo(ctx context.Context, jwtUser security.JWTUser, visibleUser *VisibleUser) error {
	user, err := models.Users(qm.Where("id = ?", jwtUser.ID), qm.Load(models.UserRels.PermissionsUsers)).One(ctx, api.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			api.logger.Printf("user %d does not exist\n", jwtUser.ID)
			return ErrUserNotExist
		}
		api.logger.Printf("error fiding user %v\n", err)
		return err
	}
	visibleUser.ID = user.ID
	visibleUser.FirstName = user.FirstName
	visibleUser.LastName = user.LastName
	visibleUser.Expiry = jwtUser.Expiry
	visibleUser.Permissions = user.R.PermissionsUsers
	visibleUser.CreatedAt = user.CreatedAt.Time
	visibleUser.UpdatedAt = user.UpdatedAt.Ptr()
	return nil
}
