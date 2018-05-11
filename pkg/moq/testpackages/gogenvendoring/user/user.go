package user

import "github.com/alxerg/somerepo"

//go:generate moq . Service

// Service does something good with computers.
type Service interface {
	DoSomething(somerepo.SomeType) error
}
