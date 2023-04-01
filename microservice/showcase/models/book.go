package models

import "time"

// initial for gorm struct to create database
type Book struct {
	Id        uint64    `gorm:"column:id;primaryKey" json:"id"`
	NameBook  string    `gorm:"column:name_book" json:"name_book"`
	Author    string    `gorm:"column:author" json:"author"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// request json format
type RequestBook struct {
	NameBook string `json:"name_book"`
	Author   string `json:"author"`
}

// response json format
type ResponseBook struct {
	Id        uint64    `json:"id"`
	NameBook  string    `json:"name_book"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// response json status and message
type BadRequest struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type NotFound struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Deleted struct {
	Message string `json:"message"`
}
