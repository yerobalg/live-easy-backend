package infrastructure

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	firebase_auth "firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

type Firebase struct {
	Auth *firebase_auth.Client
}

func InitFirebase() *Firebase {
	opt := option.WithCredentialsFile("firebase_auth_credential.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(err)
	}

	firebaseAuth, err := app.Auth(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to firebase!")

	return &Firebase{
		Auth: firebaseAuth,
	}
}

func (f *Firebase) GetUser(ctx context.Context, uid string) (*firebase_auth.UserRecord, error) {
	user, err := f.Auth.GetUser(ctx, uid)
	if err != nil {
		return nil, err
	}

	return user, nil
}
