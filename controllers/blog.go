package controllers

import (
    "net/http"
    "gopkg.in/macaron.v1"
    "github.com/jamierocks/webrocks/models"
    "sort"
)

func GetBlog(ctx *macaron.Context) {
    ctx.Data["title"] = "Blog | Jamie Mansfield"

    posts := models.GetPosts()
    sort.Sort(sort.Reverse(models.PostSlice(posts)))
    ctx.Data["posts"] = posts
    ctx.HTML(http.StatusOK, "blog")
}
