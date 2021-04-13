package model

import "time"

type Note struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Icon      string    `json:"icon"`
	Content   string    `json:"content"`
	Parent    *string   `json:"parent"`
	Children  *[]string `json:"children"`
	CreatedBy string    `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
}

type NoteCreateInput struct {
	Title   string  `json:"title"`
	Icon    string  `json:"icon"`
	Content string  `json:"content"`
	Parent  *string `json:"parent"`
}

type NoteUpdateInput struct {
	ID      string  `json:"id"`
	Title   *string `json:"title"`
	Icon    *string `json:"icon"`
	Content *string `json:"content"`
}
