package models

import (
    "time"
    "github.com/jamierocks/webrocks/modules"
)

type Author struct {
    ID int64 `gorm:"primary_key"`
    Name string
    Slug string

    Posts []Post

    CreatedAt time.Time
    UpdatedAt time.Time
}

func (a Author) GetPosts() []Post {
    var posts []Post
    modules.DB.Model(&a).Related(&posts)
    return posts
}
