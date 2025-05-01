package tui

import (
	"fmt"
	"github.com/Brain4Fish/containery/internal/utils"
	"github.com/charmbracelet/huh"
	"log"
)

func TransferMenu() {
	utils.CleanScreen()
	transferMenuForm()
}

func transferMenuForm() {
	var transferItems []string
	actionItem := ""
	transferSelector := huh.NewMultiSelect[string]().
		Title("=== Transfer Menu ===").
		Options(
			huh.NewOption("keke", "ololo"),
			huh.NewOption("keke1", "ololo1"),
			huh.NewOption("keke2", "ololo2"),
		).
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

func transferMenuSelector(selectedValues []string, action string) {
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
