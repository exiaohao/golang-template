package model

type Book struct {
	BaseModel
	Title	string	`gorm:"type:varchar(64)" json:"title"`
	ISSN	string	`gorm:"type:varchar(32);unique_index" json:"issn" binding:"required"`
	Author 	string	`gorm:"type:varchar(64)" json:"author"`
}

func (b *Book) TableName() string {
	return tablePrefix + "book"
}
