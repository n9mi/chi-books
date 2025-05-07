package entity

type Book struct {
	BaseEntity
	Title         string
	Author        string
	PublishedYear int
}
