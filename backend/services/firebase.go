package services

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"

	"google.golang.org/api/option"
)

type FirebaseService struct {
	// Add Firebase client fields here
	app        *firebase.App
	authClient *auth.Client
}

func NewFirebaseService() (*FirebaseService, error) {
	opt := option.WithAuthCredentialsFile(option.ServiceAccount, "./secret/fb.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting Auth client: %v", err)
	}

	// Initialize Firebase client here
	return &FirebaseService{
		app:        app,
		authClient: authClient,
	}, nil
}

func (fs *FirebaseService) VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	token, err := fs.authClient.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, fmt.Errorf("error verifying ID token: %v", err)
	}

	return token, nil
}
