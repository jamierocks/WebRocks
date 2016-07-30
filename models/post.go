package models

import (
    "time"
    "github.com/jamierocks/webrocks/modules"
)

type Post struct {
    ID int64 `gorm:"primary_key"`
    Title string
    Slug string
    Content string

    Author Author
    AuthorID int64

    CreatedAt time.Time
    UpdatedAt time.Time
}

func (b Post) GetAuthor() Author {
    var author Author
    modules.DB.Model(&b).Related(&author)
    return author
}
