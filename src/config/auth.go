package config

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

type AuthConfig struct {
	*auth.Client
	CGPServiceKeyJSON string `env:"GCP_SERVICE_KEY_JSON"`
}

func (c *AuthConfig) Setup() {
	opt := option.WithCredentialsJSON([]byte(c.CGPServiceKeyJSON))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(err)
	}
	ac, err := app.Auth(context.Background())
	if err != nil {
		panic(err)
	}

	c.Client = ac
}

// VerifyIDToken - IDトークンを確認して、デコードされたトークンからデバイスのuidを取得
func (c *AuthConfig) VerifyIDToken(ctx context.Context, idToken string) (string, error) {
	t, err := c.Client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return "", err
	}

	return t.UID, nil
}

// GetUserByUID - UIDによるユーザー情報の取得
func (c *AuthConfig) GetUserByUID(ctx context.Context, uid string) (*auth.UserRecord, error) {
	u, err := c.Client.GetUser(ctx, uid)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// GetUIDByEmail - メールアドレスによるユーザーUIDの取得
func (c *AuthConfig) GetUIDByEmail(ctx context.Context, email string) (string, error) {
	u, err := c.Client.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	return u.UID, nil
}

// CreateUser - 新しいFirebase Authenticationユーザーを作成
func (c *AuthConfig) CreateUser(ctx context.Context, uid string, email string, password string) (string, error) {
	params := (&auth.UserToCreate{}).
		UID(uid).
		Email(email).
		EmailVerified(false).
		Password(password).
		Disabled(false)

	u, err := c.Client.CreateUser(ctx, params)
	if err != nil {
		return "", err
	}

	return u.UID, nil
}

// UpdateEmail - メールアドレスの変更
func (c *AuthConfig) UpdateEmail(ctx context.Context, uid string, email string) error {
	params := (&auth.UserToUpdate{}).
		Email(email).
		EmailVerified(emailVerified(ctx, c, uid, email))

	if _, err := c.Client.UpdateUser(ctx, uid, params); err != nil {
		return err
	}

	return nil
}

// UpdatePassword - Passwordを変更
func (c *AuthConfig) UpdatePassword(ctx context.Context, uid string, password string) error {
	params := (&auth.UserToUpdate{}).
		Password(password)

	if _, err := c.Client.UpdateUser(ctx, uid, params); err != nil {
		return err
	}

	return nil
}

// UpdateActivated - アカウントの状態を変更
func (c *AuthConfig) UpdateActivated(ctx context.Context, uid string, disabled bool) error {
	params := (&auth.UserToUpdate{}).
		Disabled(disabled)

	if _, err := c.Client.UpdateUser(ctx, uid, params); err != nil {
		return err
	}

	return nil
}

// DeleteUser - 既存のユーザーをuidで削除
func (c *AuthConfig) DeleteUser(ctx context.Context, uid string) error {
	return c.Client.DeleteUser(ctx, uid)
}

func emailVerified(ctx context.Context, c *AuthConfig, uid string, email string) bool {
	// uid == "" -> 新規ユーザー
	if uid == "" {
		return false
	}

	// uid != "" -> 既存ユーザー
	u, err := c.Client.GetUserByEmail(ctx, uid)
	if err != nil {
		return false // TODO: エラー処理
	}

	return email == u.Email
}
