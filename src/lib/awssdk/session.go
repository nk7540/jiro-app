package awssdk

import "github.com/aws/aws-sdk-go/aws/session"

type Session struct {
	Session *session.Session
}

func NewSession(profile string) *Session {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile:           profile,
		SharedConfigState: session.SharedConfigEnable,
	}))

	return &Session{sess}
}
