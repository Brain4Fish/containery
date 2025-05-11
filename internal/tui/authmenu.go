package tui

import (
	"fmt"
	"github.com/Brain4Fish/containery/internal/db"
	"github.com/Brain4Fish/containery/internal/utils"
	"github.com/charmbracelet/huh"
	"log"
)

type credentialsData struct {
	id       uint
	username string
	password string
}

func RenderedAuthMenu() {
	utils.CleanScreen()
	authMenuForm()
}

func authMenuForm() {
	var transferItems credentialsData
	actionItem := ""
	credsSelector := huh.NewSelect[credentialsData]().
		Title("=== Transfer Menu ===").
		Options(credentialsList()...).
		Value(&transferItems).
		WithWidth(100)

	actionSelector := huh.NewSelect[string]().
		Title("Actions:").
		Options(
			huh.NewOption("[➕] Add New Credentials", "new"),
			huh.NewOption("[📝] Edit Credentials", "edit"),
			huh.NewOption("[🗑️] Delete Selected Credential", "delete"),
			huh.NewOption("[⬅️] Back", "back"),
		).
		Value(&actionItem).
		WithWidth(100)
	group := huh.NewGroup(credsSelector, actionSelector)
	form := huh.NewForm(group)
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	authMenuSelector(transferItems, actionItem)
}

func authMenuSelector(selectedValues credentialsData, action string) {
	switch action {
	case "new":
		fmt.Println("new -> ", selectedValues)
	case "edit":
		fmt.Println("edit -> ", selectedValues)
	case "auth":
		fmt.Println("delete -> ", selectedValues)
	default:
		fmt.Println("back")
		RenderedMenu() // Going back to main menu
	}
}

func credentialsList() []huh.Option[credentialsData] {
	dbConn := db.Database{DatabasePath: "settings.db"}
	dbConn.CreateConnection()
	rows, _ := dbConn.ReadCredentialsData()
	var creds []huh.Option[credentialsData]
	for _, row := range rows {
		cred := credentialsData{id: row.ID, username: row.Username, password: row.Password}
		option := huh.NewOption[credentialsData](cred.username, cred)
		creds = append(creds, option)
	}
	return creds
}
