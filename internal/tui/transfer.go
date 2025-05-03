package tui

import (
	"fmt"
	"github.com/Brain4Fish/containery/internal/db"
	"github.com/Brain4Fish/containery/internal/utils"
	"github.com/charmbracelet/huh"
	"log"
)

type imageParams struct {
	id       string
	name     string
	selected bool
}

func TransferMenu() {
	utils.CleanScreen()
	transferMenuForm()
}

// Stopped here. I must implement
func transferMenuForm() {
	var transferItems []imageParams
	actionItem := ""
	transferSelector := huh.NewMultiSelect[imageParams]().
		Title("=== Transfer Menu ===").
		Options(optionsList()...).
		Value(&transferItems).
		WithWidth(100)

	actionSelector := huh.NewSelect[string]().
		Title("Actions:").
		Options(
			huh.NewOption("[📝] Add New Image", "new"),
			huh.NewOption("[🧹] Clear All Selections", "clear"),
			huh.NewOption("[💾] Save Image List", "save"),
			huh.NewOption("[🚀] Start Transfer", "start"),
			huh.NewOption("[⬅️] Back", "back"),
		).
		Value(&actionItem).
		WithWidth(100)
	group := huh.NewGroup(transferSelector, actionSelector)
	form := huh.NewForm(group)
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	transferMenuSelector(transferItems, actionItem)
}

func transferMenuSelector(selectedValues []imageParams, action string) {
	switch action {
	case "new":
		fmt.Println("new ", selectedValues)
	case "clear":
		fmt.Println("clear ", selectedValues)
	case "save":
		fmt.Println("save ", selectedValues)
	case "start":
		fmt.Println("start ", selectedValues)
	default:
		fmt.Println("back ", selectedValues)
		RenderedMenu() // Going back to main menu
	}
}

func optionsList() []huh.Option[imageParams] {
	dbConn := db.Database{DatabasePath: "settings.db"}
	dbConn.CreateConnection()
	rows, _ := dbConn.ReadData()
	var images []huh.Option[imageParams]
	for _, row := range rows {
		image := imageParams{id: row.ID, name: fmt.Sprintf("%s:%s", row.Name, row.Tag), selected: row.Selected}
		option := huh.NewOption[imageParams](image.name, image).Selected(row.Selected)
		images = append(images, option)
	}
	return images
}
