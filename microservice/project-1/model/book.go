package model

import "time"

// initial for gorm struct to create database
type Book struct {
	Id        uint64    `gorm:"column:id;primaryKey"`
	NameBook  string    `gorm:"column:name_book"`
	Author    string    `gorm:"column:author"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

// request json format 
type RequestBook struct {
	NameBook  string    `json:"name_book"`
	Author    string    `json:"author"`
}

// response json format
type ResponseBook struct {
	Id        uint64    `json:"id"`
	NameBook  string    `json:"name_book"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}