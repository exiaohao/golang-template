package model

type Book struct {
	BaseModel
	Title	string	`gorm:"type:varchar(64)"`
	ISSN	string	`gorm:"type:varchar(32);unique_index"`
	Author 	string	`gorm:"type:varchar(64)"`
}

func (b *Book) TableName() string {
	return tablePrefix + "book"
}
