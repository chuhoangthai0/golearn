package service

type Response struct {
	StatusCode int    `json:"statusCode"`
	Body       string `json:"body"`
}

type Books struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

func Getlist() Books {
	books := []Books{
		{Id: 1, Name: "NodeJS", Author: "NodeJS"},
		{Id: 2, Name: "Golang", Author: "Golang"},
	}

	return books[1]
}


