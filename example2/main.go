package main

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

var (
	subtle    = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	textStyle = lipgloss.NewStyle()
)

func main() {
	//physicalWidth, _, _ := term.GetSize(int(os.Stdout.Fd()))
	doc := strings.Builder{}
	doc.WriteString(textStyle.Copy().Background(subtle).Render("Lipgloss"))
	fmt.Println(doc.String())

}
