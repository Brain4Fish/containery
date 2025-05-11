package tui

import (
	"fmt"
	"github.com/Brain4Fish/containery/internal/utils"
	"github.com/charmbracelet/huh"
	"log"
	"os"
)

func RenderedMenu() {
	utils.CleanScreen()
	mainMenuForm()
}

func mainMenuForm() {
	menuItem := ""
	menuSelector := huh.NewSelect[string]().
		Title("=== Main Menu ===").
		Options(
			huh.NewOption("🔁 Transfer Images", "transfer"),
			huh.NewOption("🗂️ Manage Registries", "registries"),
			huh.NewOption("🔐 Auth & Credentials", "auth"),
			huh.NewOption("🕓 Transfer History", "history"),
			huh.NewOption("⚙️ Settings", "settings"),
			huh.NewOption("❌ Exit", "exit"),
		).
		Value(&menuItem).
		WithWidth(100)
	group := huh.NewGroup(menuSelector)
	form := huh.NewForm(group)
	//form.WithTheme(customSenderTheme())
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	mainMenuSelector(menuItem)
}

func mainMenuSelector(selector string) {
	switch selector {
	case "transfer":
		TransferMenu()
	case "registries":
		fmt.Println("registries")
	case "auth":
		RenderedAuthMenu()
	case "settings":
		fmt.Println("settings")
	default:
		fmt.Println("exiting")
		os.Exit(0)
	}
}
