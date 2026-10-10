package main

import "github.com/egoist/mygo/ui"

// theme builds the theme for an appearance: the brand's neutral graphite
// surfaces with the brand green as the one accent. The derived values match
// the ports app, so both windows look like one product.
func theme(dark bool) *ui.Theme {
	t := &ui.Theme{Dark: dark}
	if dark {
		t.Background = darkCanvas
		t.Surface = darkRaised
		t.SurfaceHover = darkComposer
		t.SurfacePressed = darkInset
		t.Border = darkBorder
		t.Text = darkText
		t.TextMuted = darkTextSecondary
		t.Selection = brandGreen.Alpha(0.30)
		t.Focus = brandGreen
		t.Inverse = darkText
		t.InverseText = darkCanvas
		t.Scrollbar = darkBorderStrong
		t.Warning = statusWarningDark
		t.Success = statusSuccessDark
		t.Danger = statusDangerDark

		t.Accent = brandGreen
		t.AccentHover = brandGreen
		t.AccentPressed = brandGreen
	} else {
		t.Background = lightCanvas
		t.Surface = lightComposer
		t.SurfaceHover = lightRaised
		t.SurfacePressed = lightInset
		t.Border = lightBorder
		t.Text = lightText
		t.TextMuted = lightTextSecondary
		t.Selection = brandGreenLight.Alpha(0.30)
		t.Focus = brandGreenLightDark
		t.Inverse = lightText
		t.InverseText = lightCanvas
		t.Scrollbar = lightBorderStrong
		t.Warning = statusWarningLight
		t.Success = statusSuccessLight
		t.Danger = statusDangerLight

		t.Accent = brandGreenLight
		t.AccentHover = brandGreenLight
		t.AccentPressed = brandGreenLightDark
	}
	// The brand green is bright in both appearances, so the text on it is
	// the brand's near-black in both: white on it measures under 2:1.
	t.AccentText = onBrandGreen
	t.Radius = 8
	return t
}

// controlBorder is the stroke around a control, stronger than a divider: a
// control has to read as a control. The brand system draws hairlines rather
// than shadows, so this stroke is the only thing that separates a button
// from the surface behind it.
func controlBorder(dark bool) ui.Color {
	if dark {
		return darkBorderStrong
	}
	return lightBorderStrong
}
