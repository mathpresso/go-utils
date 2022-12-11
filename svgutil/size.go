package svgutil

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type unit struct {
	notation     string
	pixelPerUnit float64
}

var (
	pixel      = unit{notation: "px", pixelPerUnit: 1}
	inch       = unit{notation: "in", pixelPerUnit: 96}
	centimeter = unit{notation: "cm", pixelPerUnit: 37.79}
	millimeter = unit{notation: "mm", pixelPerUnit: 3.779}
	point      = unit{notation: "pt", pixelPerUnit: 1.33}
	pica       = unit{notation: "pc", pixelPerUnit: 16}

	units = []unit{
		pixel,
		inch,
		centimeter,
		millimeter,
		point,
		pica,
	}
)

func (u unit) toValue(s string) (value, error) {
	if !strings.HasSuffix(s, u.notation) {
		return value{}, fmt.Errorf("not %s: %s", u.notation, s)
	}
	trimmed := strings.TrimSuffix(s, u.notation)
	f, err := strconv.ParseFloat(trimmed, 64)
	if err != nil {
		return value{}, fmt.Errorf("parse float: %w", err)
	}
	return value{
		value: f,
		unit:  u,
	}, nil
}

type value struct {
	value float64
	unit  unit
}

func (s value) toPixel() value {
	return value{
		value: s.value * s.unit.pixelPerUnit,
		unit:  pixel,
	}
}

type svg struct {
	Width   string `xml:"width,attr"`
	Height  string `xml:"height,attr"`
	ViewBox string `xml:"viewBox,attr"`
}

func Size(str string) (width, height int, err error) {
	parsed := svg{}
	err = xml.Unmarshal([]byte(str), &parsed)
	if err != nil {
		return 0, 0, fmt.Errorf("parse failed: %w", err)
	}

	w, h := parsed.Width, parsed.Height

	if w == "" && h == "" {
		w, h, err = whFromViewBox(parsed.ViewBox)
		if err != nil {
			return 0, 0, err
		}
	}
	if unicode.IsNumber(rune(w[len(w)-1])) {
		w += pixel.notation
	}
	if unicode.IsNumber(rune(h[len(h)-1])) {
		h += pixel.notation
	}

	for _, unit := range units {
		if width == 0 {
			wv, err := unit.toValue(w)
			if err == nil {
				width = int(wv.toPixel().value)
			}
		}
		if height == 0 {
			hv, err := unit.toValue(h)
			if err == nil {
				height = int(hv.toPixel().value)
			}
		}
		if width != 0 && height != 0 {
			return width, height, nil
		}
	}
	return 0, 0, fmt.Errorf("invalid length value: (%s, %s)", w, h)
}

func whFromViewBox(vb string) (width string, height string, err error) {
	split := strings.Split(vb, " ")
	if len(split) != 4 {
		return "", "", fmt.Errorf("invalid view box: %s", vb)
	}
	return split[2], split[3], nil
}
