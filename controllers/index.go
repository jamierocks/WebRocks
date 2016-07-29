package controllers

import (
    "net/http"
    "gopkg.in/macaron.v1"
)

func GetIndex(ctx *macaron.Context) {
    ctx.Data["title"] = "Jamie Mansfield"
    ctx.HTML(http.StatusOK, "index")
}
