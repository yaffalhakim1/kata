package main

// Appearance is which of the three looks the app shows. Following the
// desktop is the default, and the third choice on the menu.
type Appearance int

const (
	AppearanceSystem Appearance = iota
	AppearanceLight
	AppearanceDark
)

// String is the appearance's name, as the settings file stores it and as
// the menu shows it.
func (a Appearance) String() string {
	switch a {
	case AppearanceLight:
		return "Light"
	case AppearanceDark:
		return "Dark"
	default:
		return "System"
	}
}

// parseAppearance reads a stored appearance, falling back to following the
// system for anything unrecognised.
func parseAppearance(s string) Appearance {
	switch s {
	case "light", "Light":
		return AppearanceLight
	case "dark", "Dark":
		return AppearanceDark
	default:
		return AppearanceSystem
	}
}

// The pages the window shows, as indices the segmented control sets.
const (
	TabStudy = iota
	TabQuiz
)
