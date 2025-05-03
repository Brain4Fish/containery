package main

import (
	"github.com/Brain4Fish/containery/internal/db"
	"github.com/Brain4Fish/containery/internal/tui"
)

func main() {

	tui.RenderedMenu()
	dbConn := db.Database{DatabasePath: "settings.db"}
	dbConn.CreateConnection()
	dbConn.ReadData()
}
