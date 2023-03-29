package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

// pour entrer seulement un nombre de 0 à 10
type ratingEntry struct {
	widget.Entry
}

func newRatingEntry() *ratingEntry {
	entry := &ratingEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}

func (e *ratingEntry) TypedRune(r rune) {
	if r >= '0' && r <= '9' {
		e.Entry.TypedRune(r)
	}
}

// Icon cliquable
type clickIcon struct {
	widget.Icon
	OnTapped func() `json:"-"`
}

func newClickIcon() *clickIcon {
	icon := &clickIcon{}
	icon.ExtendBaseWidget(icon)
	return icon
}

func (i *clickIcon) Tapped(_ *fyne.PointEvent) {
	if i.OnTapped != nil {
		i.OnTapped()
	}
}

// extIcon
type extIcon struct {
	desktop.Hoverable
	widget.Icon
	indx         int
	iconOff      fyne.Resource
	iconOffHover fyne.Resource
	iconOn       fyne.Resource
	iconOnHover  fyne.Resource
	tapped       bool
	disabled     bool
	OnTapped     func(i int) `json:"-"`
	OnMouseIn    func(i int) `json:"-"`
	OnMouseOut   func(i int) `json:"-"`
}

// MouseIn is called when a desktop pointer enters the widget
func (i *extIcon) MouseIn(*desktop.MouseEvent) {
	i.tapped = false
	if !i.disabled {
		if i.OnMouseIn != nil {
			i.OnMouseIn(i.indx)
		}
		i.MouseHoverIn()
	}
}

// MouseHoverIn is called for all widgets when a desktop pointer enters one of widgets
func (i *extIcon) MouseHoverIn() {
	if i.Resource.Name() == i.iconOn.Name() {
		i.SetResource(i.iconOnHover)
	} else {
		i.SetResource(i.iconOffHover)
	}
	i.Refresh()
}

// MouseIn is called when a desktop pointer enters the widget
func (i *extIcon) MouseMoved(*desktop.MouseEvent) {
}

// MouseOut is called when a desktop pointer exits the widget
func (i *extIcon) MouseOut() {
	if !i.tapped {
		if i.OnMouseOut != nil {
			i.OnMouseOut(i.indx)
		}
		i.MouseHoverOut()
	}
}

// MouseHoverOut is called for all widgets when a desktop pointer exits all widgets
func (i *extIcon) MouseHoverOut() {
	if i.Resource.Name() == i.iconOnHover.Name() {
		i.SetResource(i.iconOn)
	} else {
		i.SetResource(i.iconOff)
	}
	i.Refresh()
}

func (i *extIcon) Tapped(_ *fyne.PointEvent) {
	if i.OnTapped != nil {
		i.OnTapped(i.indx)
	}
}

func (i *extIcon) Enable() {
	i.disabled = false
}

func (i *extIcon) Disable() {
	i.disabled = true
}

func newExtIcon(iconOn, iconOnHover, iconOff, iconOffHover fyne.Resource) *extIcon {
	icon := &extIcon{}
	icon.iconOn = iconOn
	icon.iconOnHover = iconOnHover
	icon.iconOff = iconOff
	icon.iconOffHover = iconOffHover
	icon.disabled = false
	icon.SetResource(iconOff)

	icon.ExtendBaseWidget(icon)
	return icon
}

// // contextButton
// type contextButton struct {
// 	widget.Button
// 	popup    *widget.PopUp
// 	OnTapped func() `json:"-"`
// }

// func (b *contextButton) Tapped(e *fyne.PointEvent) {
// 	if b.OnTapped != nil {
// 		b.OnTapped()
// 		b.popup.ShowAtPosition(e.AbsolutePosition)
// 	}
// }

// func newContextButton(label string, tapped func()) *contextButton {
// 	button := &contextButton{}
// 	button.SetText(label)
// 	button.OnTapped = tapped

// 	button.ExtendBaseWidget(button)
// 	return button
// }

// func newContextButtonWithIcon(label string, icon fyne.Resource, tapped func()) *contextButton {
// 	button := &contextButton{}
// 	button.SetText(label)
// 	button.SetIcon(icon)
// 	button.OnTapped = tapped

// 	button.ExtendBaseWidget(button)
// 	return button
// }

// // hoverButton
// type hoverButton struct {
// 	widget.Button
// 	iconOut fyne.Resource
// 	iconIn  fyne.Resource
// }

// // MouseIn is called when a desktop pointer enters the widget
// func (b *hoverButton) MouseIn(*desktop.MouseEvent) {
// 	b.Icon = b.iconIn
// 	b.Refresh()
// }

// // MouseOut is called when a desktop pointer exits the widget
// func (b *hoverButton) MouseOut() {
// 	b.Icon = b.iconOut
// 	b.Refresh()
// }

// func newHoverButtonWithIcons(label string, iconIn fyne.Resource, iconOut fyne.Resource, tapped func()) *hoverButton {
// 	button := &hoverButton{}
// 	button.iconIn = iconIn
// 	button.iconOut = iconOut
// 	button.SetText(label)
// 	button.SetIcon(iconOut)
// 	button.OnTapped = tapped

// 	button.ExtendBaseWidget(button)
// 	return button
// }

// // numEntry
// type numEntry struct {
// 	widget.Entry
// }

// func newNumEntry() *numEntry {
// 	e := &numEntry{}
// 	e.ExtendBaseWidget(e)
// 	e.Validator = validation.NewRegexp(`\d`, "Must contain a number")
// 	return e
// }
