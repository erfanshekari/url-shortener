package models

type Model interface {
	Save() error
	ToJson() map[string]interface{}
}
