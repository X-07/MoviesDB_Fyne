package main

import "fyne.io/fyne/v2/dialog"

func (ui *appUI) downloadFiche() {
	dialog.ShowInformation("Download", ui.headerTitle.Text, ui.win)
}
