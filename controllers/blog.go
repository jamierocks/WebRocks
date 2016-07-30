package controllers

import (
    "net/http"
    "gopkg.in/macaron.v1"
    "github.com/jamierocks/webrocks/models"
)

func GetBlog(ctx *macaron.Context) {
    ctx.Data["title"] = "Blog | Jamie Mansfield"
    ctx.Data["posts"] = models.GetPosts()
    ctx.HTML(http.StatusOK, "blog")
}
