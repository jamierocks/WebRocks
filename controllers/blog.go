package controllers

import (
    "sort"
    "net/http"
    "gopkg.in/macaron.v1"
    "github.com/jamierocks/webrocks/models"
)

func GetBlog(ctx *macaron.Context) {
    ctx.Data["title"] = "Blog | Jamie Mansfield"

    posts := models.GetPosts()
    sort.Sort(sort.Reverse(models.PostSlice(posts)))
    ctx.Data["posts"] = posts
    ctx.HTML(http.StatusOK, "blog")
}

func GetPost(ctx *macaron.Context) {
    post := models.GetPost(ctx.Params("slug"))

    if post.ID != 0 {
        ctx.Data["title"] = post.Title
        ctx.Data["post"] = post
        ctx.HTML(http.StatusOK, "post")
    } else {
        ctx.Data["title"] = "404 Not Found - Jamie Mansfield"
        ctx.HTML(http.StatusOK, "404")
    }
}
