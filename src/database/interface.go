package database

type Provider interface {
	Connect(params interface{}) error
}
