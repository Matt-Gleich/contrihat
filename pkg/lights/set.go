package lights

import (
	"github.com/Matt-Gleich/contrihat/pkg/api"
	"github.com/Matt-Gleich/logoru"
	"github.com/nathany/bobblehat/sense/screen"
	"github.com/nathany/bobblehat/sense/screen/color"
	"gopkg.in/go-playground/colors.v1"
)

// Set the lights on the sense hat
func Set(contributions api.Query) {
	fb := screen.NewFrameBuffer()
	days := mergeDays(contributions)
	var (
		x int
		y int
	)
	for _, day := range days {
		if x == 8 {
			y++
			x = 0
		}
		if y == 8 {
			break
		}
		fb.SetPixel(x, y, convert(day))
		x++
	}
	err := screen.Draw(fb)
	if err != nil {
		logoru.Error("Failed to draw screen;", err)
	}
	logoru.Success("Updated lights!")
}

// Merging all weeks
func mergeDays(contributions api.Query) (days []string) {
	for _, week := range contributions.Viewer.ContributionsCollection.ContributionCalendar.Weeks {
		for _, day := range week.ContributionDays {
			days = append([]string{day.Color}, days...)
		}
	}
	return days
}

// Convert the hex code to bobblehat color
func convert(rawHex string) color.Color {
	hex, err := colors.ParseHEX(rawHex)
	if err != nil {
		logoru.Error("Failed to parse hex code;", hex)
	}
	rgb := hex.ToRGB()
	return color.New(rgb.R, rgb.G, rgb.B)
}