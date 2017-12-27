package models

import (
    "github.com/fiscaluno/fiscaluno-api/database"
)

var db = database.GetInstance()

// Model contract
type Model interface {
    All()
    Find()
    Save()
    Delete()
    Update()
}
