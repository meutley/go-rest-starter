package model

// User represents a user model in the database.
type User struct {
	Id        int
	FirstName string
	LastName  string
}

// UserDataSet represents a set of User models.
type UserDataSet struct {
	DataSet
	Models []User
}

// Find will query the database to find a User model with the given id value.
func (u UserDataSet) Find(id interface{}) (interface{}, error) {
	for _, u := range u.Models {
		if u.Id == id {
			return u, nil
		}
	}

	return nil, nil
}
