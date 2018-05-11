package user

import "github.com/alxerg/somerepo"

// Service does something good with computers.
type Service interface {
	DoSomething(somerepo.SomeType) error
}
