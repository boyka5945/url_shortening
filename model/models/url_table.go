package model

type UrlTable struct {
	ID             int64 `gorm:"primary_key"`
	OriginalUrl    string
	ShortUrl 	   string
	Offset		   int64
}