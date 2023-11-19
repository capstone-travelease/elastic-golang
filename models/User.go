package models

type User struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
	Age      int    `json:"age"`
}

type TestResponse struct {
	Name string `json:"name"`
}
