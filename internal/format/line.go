package format

import "github.com/digiconvent/logging/internal/color"

func Line(severityColor string, msg any, noColor bool) string {
	if noColor {
		return now() + ": " + prep(msg)
	}
	return color.White + now() + ":" + severityColor + " " + prep(msg) + " " + color.Reset
}
