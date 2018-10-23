package database

import "./model"

// InMemory is an in-memory database example, for the sake of testing.
type InMemory struct {
	Users model.UserDataSet
}

// Connect initializes the in-memory database.
func (r InMemory) Connect(params interface{}) error {
	return nil
}

// FindUser finds a User model in the database by id.
func (r InMemory) FindUser(id int) (model.User, error) {
	user, err := r.Users.Find(id)
	if user == nil || err != nil {
		return model.User{}, err
	}

	return user.(model.User), nil
}
