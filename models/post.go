package models

import (
    "time"
    "github.com/jamierocks/webrocks/modules"
)

type Post struct {
    ID int64 `gorm:"primary_key"`
    Title string
    Slug string
    Content string `gorm:"type:mediumtext"`

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

func GetPosts() []Post {
    var posts []Post
    modules.DB.Find(&posts)
    return posts
}

type PostSlice []Post

func (p PostSlice) Len() int {
    return len(p)
}

func (p PostSlice) Less(i, j int) bool {
    return p[i].CreatedAt.Before(p[j].CreatedAt)
}

func (p PostSlice) Swap(i, j int) {
    p[i], p[j] = p[j], p[i]
}
