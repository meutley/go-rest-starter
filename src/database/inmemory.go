package database

import (
	"./model"
)

// InMemory is an in-memory database example, for the sake of testing.
type InMemory struct {
	Users model.DataSet
}

// In-memory dataset is initialized here
var usersDataSet = model.UserDataSet{
	Models: []model.User{
		model.User{
			Id:        1,
			FirstName: "John",
			LastName:  "Doe",
		},
	},
}

// Connect initializes the in-memory database.
func (r *InMemory) Connect(params interface{}) error {
	r.Users = usersDataSet
	return nil
}

// FindUser finds a User model in the database by id.
func (r *InMemory) FindUser(id int) (model.User, error) {
	user, err := r.Users.Find(id)
	if user == nil || err != nil {
		return model.User{}, err
	}

	return user.(model.User), nil
}
