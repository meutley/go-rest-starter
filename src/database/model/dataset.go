package model

// Queryable is an interface that provides query functionality to DataSet instances.
type Queryable interface {
	Find(id interface{}) (interface{}, error)
	Query(query string) ([]interface{}, error)
	Insert(model interface{}) (interface{}, error)
	Delete(id interface{}) error
}

// DataSet represents a set of data (Queryable objects).
type DataSet interface {
	Queryable
}
