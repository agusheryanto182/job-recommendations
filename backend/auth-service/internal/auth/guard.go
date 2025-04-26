package auth

import (
	"time"
)

type Guard string

const (
	Admin              Guard = "admin"
	User               Guard = "user"
	UserResetPassword  Guard = "user-reset-password"
	AdminResetPassword Guard = "admin-reset-password"
)

// Expire time each guard in milisecond
func (g Guard) ExpireTime() time.Duration {
	switch g {
	case Admin:
		return (time.Hour * 24) * 7 // 1 week
	case User:
		return (time.Hour * 24) * 7 // 1 week
	case UserResetPassword:
		return (time.Hour * 1) * 1 // 1 hour
	case AdminResetPassword:
		return (time.Hour * 1) // 1 hour
	}
	return (time.Hour * 0) // 1 week
}
