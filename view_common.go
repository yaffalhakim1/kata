package main

import (
	"github.com/egoist/mygo/ui"
)

// The window's top bar: what the app is, the two pages, and the appearance
// menu. The pages are a segmented control, so switching is one click.
//
// The row wraps: at the smallest window the controls move to a second line
// rather than the appearance button being clipped at the right edge. The
// brand group grows, so on a wide window the controls still sit at the
// trailing edge.
func (a *app) header(c *ui.Context) {
	t := c.Theme()
	ui.Row(c).Wrap().AlignItems(ui.Center).Gap(14).Padding(pageMargin, 16).Children(func() {
		ui.Row(c).AlignItems(ui.Center).Gap(10).Grow(1).MinWidth(0).Children(func() {
			ui.Text(c, "Kata").FontSize(20).FontWeight(650).Shrink(0)
			ui.Text(c, "日本語 ⇄ Indonesia").FontSize(13).TextColor(t.TextMuted).
				SingleLine().Ellipsis("…").Shrink(1).MinWidth(0)
		})
		ui.Segmented(c, &a.tab, "Study", "Quiz").Label("Page")
		a.appearanceButton(c)
	})
}

// appearanceButton shows what the app is doing now, and opens a menu
// offering all three choices directly, so reaching the desktop's setting
// is one click rather than a cycle back through the others.
func (a *app) appearanceButton(c *ui.Context) {
	icon, label := a.appearanceIcon(), a.appearance.String()
	b := ui.ButtonBase(c).Padding(buttonPadV, buttonPadH).Radius(8).
		Focusable().Shrink(0).
		Background(c.Theme().Surface).Border(1, controlBorder(c.Theme().Dark)).
		Label("Appearance: " + label).
		Tooltip("Appearance: " + label).
		Menu(func(m *ui.Menu) {
			for _, choice := range []Appearance{AppearanceSystem, AppearanceLight, AppearanceDark} {
				if m.Item(choice.String()).Checked(a.appearance == choice).Chosen() {
					a.setAppearance(choice)
				}
			}
		})
	b.Children(func() {
		ui.Row(c).AlignItems(ui.Center).Gap(iconLabelGap).Children(func() {
			ui.Icon(c, icon).Size(14, 14)
			ui.Text(c, label).FontSize(13).Shrink(0)
			ui.Icon(c, iconCaret).Size(11, 11)
		})
	})
	if b.Hovered() {
		b.Background(c.Theme().SurfaceHover)
	}
}

// appearanceIcon is the icon of what the app is doing now: the desktop's
// screen when it follows the system, else the sun or the moon.
func (a *app) appearanceIcon() *ui.SVG {
	switch a.appearance {
	case AppearanceLight:
		return iconSun
	case AppearanceDark:
		return iconMoon
	default:
		return iconSys
	}
}

// setAppearance chooses an appearance and remembers it.
func (a *app) setAppearance(want Appearance) {
	if a.appearance == want {
		return
	}
	a.appearance = want
	a.saveSettings()
}

// pageMargin is the window's horizontal margin, shared by every region so
// their edges line up. buttonPadV and buttonPadH size a control's hit area
// to 32 DIPs tall, clearing the 24x24 floor of WCAG 2.5.8.
const (
	pageMargin   = 20
	buttonPadV   = 7
	buttonPadH   = 12
	iconLabelGap = 8
	contentMax   = 720
)

// card is a raised surface with the brand's hairline border, used for the
// study cards and the quiz panel.
func card(c *ui.Context) ui.Element {
	t := c.Theme()
	return ui.Column(c).Padding(20).Gap(12).Radius(12).
		Background(t.Surface).Border(1, controlBorder(t.Dark))
}

// sectionTitle is a small uppercase label above a group.
func sectionTitle(c *ui.Context, s string) {
	ui.Text(c, s).FontSize(11).FontWeight(600).TextColor(c.Theme().TextMuted)
}

// hint is the muted secondary line under a title.
func hint(c *ui.Context, s string) {
	ui.Text(c, s).FontSize(13).TextColor(c.Theme().TextMuted)
}
