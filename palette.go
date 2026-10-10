package main

import "github.com/egoist/mygo/ui"

// The colour system, copied from the ports brand: neutral graphite surfaces
// with hue reserved for meaning, and the brand green as the one accent. The
// values are the same ones the other clients use, so the products stay one
// family.
var (
	// Dark: graphite canvas, ink text.
	darkCanvas   = ui.Hex("#1A1A1A")
	darkRaised   = ui.Hex("#232323")
	darkComposer = ui.Hex("#212121")
	darkInset    = ui.Hex("#151515")
	// The hairlines: white 7% and 14% over the canvas, resolved.
	darkBorder       = ui.Hex("#2A2A2A")
	darkBorderStrong = ui.Hex("#3A3A3A")

	darkText          = ui.Hex("#E2E2E2")
	darkTextSecondary = ui.Hex("#A3A3A3")

	// Light: paper canvas, ink text.
	lightCanvas   = ui.Hex("#F6F5F6")
	lightRaised   = ui.Hex("#ECECEC")
	lightComposer = ui.Hex("#FFFFFF")
	lightInset    = ui.Hex("#E6E6E6")
	// hsla(220,10%,12%,.08) and .15 over the canvas, resolved.
	lightBorder       = ui.Hex("#E5E4E5")
	lightBorderStrong = ui.Hex("#D5D5D7")

	lightText          = ui.Hex("#242424")
	lightTextSecondary = ui.Hex("#666666")

	// The brand green, and the near-black that sits on it. The green is
	// bright in both appearances, so its text is dark in both.
	brandGreen      = ui.Hex("#26E085")
	brandGreenLight = ui.Hex("#00DA7D")
	onBrandGreen    = ui.Hex("#17181C")

	// Status colours, the only other saturated hues, darkened as text so
	// each passes 4.5:1 on its canvas at 13px.
	statusSuccessDark  = ui.Hex("#62C987")
	statusSuccessLight = ui.Hex("#1E7A40")
	statusWarningDark  = ui.Hex("#E0B36A")
	statusWarningLight = ui.Hex("#8A5714")
	statusDangerDark   = ui.Hex("#E2726A")
	statusDangerLight  = ui.Hex("#B03A33")

	// The focus ring on the light canvas: the light brand green measures
	// 1.71:1 there, and a focus indicator needs 3:1.
	brandGreenLightDark = ui.Hex("#00814A")
)
