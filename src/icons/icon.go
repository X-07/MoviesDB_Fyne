package icon

import (
	_ "embed"

	"fyne.io/fyne/v2"
)

//go:embed logo.png
var logo []byte
var LogoIco = &fyne.StaticResource{
	StaticName:    "logo",
	StaticContent: logo,
}

//go:embed starOn.png
var starOn []byte
var StarOnIco = &fyne.StaticResource{
	StaticName:    "starOn",
	StaticContent: starOn,
}

//go:embed starOnHover.png
var starOnHover []byte
var StarOnHoverIco = &fyne.StaticResource{
	StaticName:    "starOnHover",
	StaticContent: starOnHover,
}

//go:embed starOff.png
var starOff []byte
var StarOffIco = &fyne.StaticResource{
	StaticName:    "starOff",
	StaticContent: starOff,
}

//go:embed starOffHover.png
var starOffHover []byte
var StarOffHoverIco = &fyne.StaticResource{
	StaticName:    "starOffHover",
	StaticContent: starOffHover,
}
