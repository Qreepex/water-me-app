package services

import (
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"firebase.google.com/go/messaging"

	"google.golang.org/api/option"
)

type FirebaseService struct {
	app             *firebase.App
	authClient      *auth.Client
	messagingClient *messaging.Client
}

func NewFirebaseService() (*FirebaseService, error) {
	// Use GOOGLE_APPLICATION_CREDENTIALS env var if set (for Kubernetes)
	// Otherwise fall back to local path for development
	credPath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if credPath == "" {
		credPath = "./secret/fb.json"
	}

	opt := option.WithCredentialsFile(credPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting Auth client: %v", err)
	}

	messagingClient, err := app.Messaging(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting Messaging client: %v", err)
	}

	return &FirebaseService{
		app:             app,
		authClient:      authClient,
		messagingClient: messagingClient,
	}, nil
}

func (fs *FirebaseService) VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	token, err := fs.authClient.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, fmt.Errorf("error verifying ID token: %v", err)
	}

	return token, nil
}

// SendNotification sends a push notification to a single device token
func (fs *FirebaseService) SendNotification(
	ctx context.Context,
	token, title, body string,
	data map[string]string,
) error {
	message := &messaging.Message{
		Token: token,
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Data: data,
		Android: &messaging.AndroidConfig{
			Priority: "high",
		},
		APNS: &messaging.APNSConfig{
			Headers: map[string]string{
				"apns-priority": "10",
			},
		},
	}

	_, err := fs.messagingClient.Send(ctx, message)
	return err
}

// SendMulticastNotification sends a push notification to multiple device tokens (up to 500)
func (fs *FirebaseService) SendMulticastNotification(
	ctx context.Context,
	tokens []string,
	title, body string,
	data map[string]string,
) (*messaging.BatchResponse, error) {
	if len(tokens) == 0 {
		return nil, fmt.Errorf("no tokens provided")
	}

	if len(tokens) > 500 {
		return nil, fmt.Errorf("too many tokens (max 500)")
	}

	message := &messaging.MulticastMessage{
		Tokens: tokens,
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Data: data,
		Android: &messaging.AndroidConfig{
			Priority: "high",
		},
		APNS: &messaging.APNSConfig{
			Headers: map[string]string{
				"apns-priority": "10",
			},
		},
	}

	return fs.messagingClient.SendMulticast(ctx, message)
}
