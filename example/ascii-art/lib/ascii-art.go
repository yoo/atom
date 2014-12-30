package main

import "github.com/JohannWeging/atom"
import "github.com/CrowdSurge/banner"

func activate() {
	atom.Commands.Add("atom-workspace", "ascii-art:convert", convert)
}

func convert() {
	editor := atom.Workspace.GetActiveTextEditor()
	selections := editor.GetSelections()
	selection := selections[0]
	text := selection.GetText()
	asciiArt := banner.PrintS(text)
	selection.InsertText(asciiArt)
}

func main() {
	atom.Activate(activate)
}
