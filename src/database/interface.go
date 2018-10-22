package database

import "./model"

// Provider represents a database provider.
type Provider interface {
	Connect(params interface{}) error
}

// DatabaseContract should contain all the methods available via the database/repository, as it is
//	injected via a Context as the interface rather than a concrete implementation.
type DatabaseContract interface {
	Provider
	FindUser(id int) (model.User, error)
}
