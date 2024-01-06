package user

import (
	"context"
	"fmt"

	"encore.dev/beta/auth"
	"encore.dev/rlog"
	firebase "firebase.google.com/go/v4"
	fbauth "firebase.google.com/go/v4/auth"
	"go4.org/syncutil"
	"google.golang.org/api/option"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var userMap = map[string]string{
	"1": "Alice",
	"2": "Bob",
	"3": "Caroline",
	"4": "Dave",
}

//encore:api public method=GET path=/users
func ListUsers(ctx context.Context) (*ListResponse, error) {
	var users []User
	for k, v := range userMap {
		users = append(users, User{ID: k, Name: v})
	}

	return &ListResponse{Users: users}, nil
}

type ListResponse struct {
	Users []User `json:"users"`
}

//encore:api auth method=GET path=/users/:id
func GetUser(ctx context.Context, id string) (*UserResponse, error) {
	if v, ok := userMap[id]; ok {
		return &UserResponse{User: User{ID: id, Name: v}}, nil
	}
	return nil, fmt.Errorf("user with id %s not found", id)
}

type UserResponse struct {
	User User `json:"user"`
}

var (
	fbAuth    *fbauth.Client
	setupOnce syncutil.Once
)

func setupFB() error {
	return setupOnce.Do(func() error {
		opt := option.WithCredentialsJSON([]byte(secrets.FirebasePrivateKey))
		app, err := firebase.NewApp(context.Background(), nil, opt)

		if err == nil {
			fbAuth, err = app.Auth(context.Background())

		}

		return err
	})
}

var secrets struct {
	// FirebasePrivateKey is the JSON credentials for calling Firebase.
	FirebasePrivateKey string
}

type Data struct {
	Email string
	Name  string
}

//encore:authhandler
func AuthHandler(ctx context.Context, token string) (auth.UID, *Data, error) {
	// Validate the token and look up the user id, for example by calling Firebase Auth: https://encore.dev/docs/how-to/firebase-auth
	rlog.Info("Authing Jobs", "Token", token)
	if err := setupFB(); err != nil {
		return "", nil, err
	}
	rlog.Info("Authing Jobs", "FbAuth", fbAuth)
	tok, err := fbAuth.VerifyIDToken(ctx, token)
	if err != nil {
		return "", nil, err
	}

	email, _ := tok.Claims["email"].(string)
	name, _ := tok.Claims["name"].(string)

	uid := auth.UID(tok.UID)

	usr := &Data{
		Email: email,
		Name:  name,
	}
	return uid, usr, nil
}

func GetFireBaseEmail(ctx context.Context, token string) (string, error) {
	if err := setupFB(); err != nil {
		return "", err
	}
	tok, err := fbAuth.VerifyIDToken(ctx, token)
	if err != nil {
		return "", err
	}

	email, _ := tok.Claims["email"].(string)
	return email, nil
}
