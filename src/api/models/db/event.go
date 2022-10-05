package db

import "time"

type UserActivityEvent int8

const (
	BlankEvent UserActivityEvent = iota

	// Regular events
	LoginEvent
	RegisterEvent
	PasswordChangeEvent
	ResetPasswordEvent
	LogoutEvent

	// Irregular events
	FailedLoginEvent
	FailedRegisterEvent
	TooManyRequests
)

type EventModel struct {
	Type      int8      `json:"type" bson:"t"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt,omitempty"`
}
