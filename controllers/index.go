package controllers

import (
    "net/http"
    "gopkg.in/macaron.v1"
)

func GetIndex(ctx *macaron.Context) {
    ctx.Data["title"] = "Jamie Mansfield"
    ctx.HTML(http.StatusOK, "index")
}

func NotFound(ctx *macaron.Context) {
    ctx.Data["title"] = "404 Not Found - Jamie Mansfield"
    ctx.HTML(http.StatusNotFound, "404")
}
