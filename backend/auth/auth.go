package auth

import (
	"context"
	"fmt"
	"math/rand"

	"encore.app/backend/db"
	"encore.app/backend/user"
	"encore.dev/rlog"
)

const TOKEN = "dummy-token"

type LoginRequest struct {
	Token string `json:"token"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type Data struct {
	Email string
	Name  string
}

//encore:api public method=POST path=/auth/login
func Login(ctx context.Context, params *LoginRequest) (*LoginResponse, error) {
	// Validate the email and password, for example by calling Firebase Auth: https://encore.dev/docs/how-to/firebase-auth
	rlog.Info("User Login", "params", params)
	rlog.Info("User login", "token", params.Token)
	pool, err := db.Get(context.Background())
	if err != nil {
		return nil, err
	}
	email, err := user.GetFireBaseEmail(ctx, params.Token)
	if err != nil {
		return nil, err
	}
	rows, err := pool.Query(ctx, "select * from users where email = $1", email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() == false {
		// generate random ID

		// insert user
		_, err := pool.Exec(ctx, "insert into users (email, id) values ($1, $2)", email, fmt.Sprintf("%d", rand.Intn(1000000)))
		if err != nil {
			return nil, err
		}

	}
	return &LoginResponse{Token: params.Token}, nil
}
