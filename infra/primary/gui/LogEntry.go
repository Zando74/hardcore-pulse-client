package gui

import (
	"fmt"
	"image/color"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type LogWriter struct {
    Entry *widget.Entry
}

func (lw *LogWriter) Write(p []byte) (n int, err error) {
    text := string(p)
	fmt.Println(text)
	if strings.Contains(text, "[DEATHLOG-TRACKER]") {
		maxLines := 10
		lines := strings.Split(lw.Entry.Text, "\n")
		lines = append(lines, text)
		if len(lines) > maxLines {
			lines = lines[len(lines)-maxLines:]
		}
		fyne.Do(func() {
			lw.Entry.SetText(strings.Join(lines, "\n"))
			lw.Entry.Refresh()
		})
		
		return len(p), nil
	}
    return 0, nil
}


func NewLogView() *fyne.Container {
	logEntry := widget.NewMultiLineEntry()
	logEntry.Disable()
	logEntry.Wrapping = fyne.TextWrapWord
	logEntry.TextStyle = fyne.TextStyle{Monospace: true}


	logEntry.Text = ""
	fyne.CurrentApp().Settings().SetTheme(&BlackTextTheme{theme.DefaultTheme()})

	scroll := container.NewVScroll(logEntry)
	scroll.SetMinSize(fyne.NewSize(800, 200))

	content := container.NewStack(scroll)

	log.SetOutput(&LogWriter{Entry: logEntry})

	return content
}

type BlackTextTheme struct {
	fyne.Theme
}

func (t *BlackTextTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameForeground:
		return color.Black
	case theme.ColorNameInputBorder:
		return color.Gray{Y: 80}
	case theme.ColorNameInputBackground:
		return color.Black
	}
	return t.Theme.Color(name, variant)
}
