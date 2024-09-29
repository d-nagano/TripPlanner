package models

type Prefecture struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Prefectures *[]Prefecture
