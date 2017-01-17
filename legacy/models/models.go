package models

import (
    "github.com/jamierocks/webrocks/modules"
)

func AutoMigrate() {
    modules.DB.AutoMigrate(
        &Post{},
        &Author{},
    )
}
