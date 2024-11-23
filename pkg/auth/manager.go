package auth

import "github.com/joakimen/ww/pkg/credentials"

type Manager interface {
	Login(credentials.Manager) error
	Logout(credentials.Manager) error
	Show(credentials.Manager) error
	Status(credentials.Manager) error
}
