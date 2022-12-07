package svgutil

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

const (
	pixel = "px"

	inch      = "in"
	pxPerInch = float64(96)

	centimeter      = "cm"
	pxPerCentimeter = float64(37.79)

	millimeter      = "mm"
	pxPerMillimeter = float64(3.779)

	point      = "pt"
	pxPerPoint = float64(1.33)

	pica      = "pc"
	pxPerPica = float64(16)
)

type svg struct {
	Width   string `xml:"width,attr"`
	Height  string `xml:"height,attr"`
	ViewBox string `xml:"viewBox,attr"`
}

func Size(str string) (width, height int, err error) {
	parsed := svg{}
	err = xml.Unmarshal([]byte(str), &parsed)
	if err != nil {
		return 0, 0, fmt.Errorf("%w: parse svg", err)
	}

	w, h := parsed.Width, parsed.Height

	if w == "" && h == "" {
		w, h, err = whFromViewBox(parsed.ViewBox)
		if err != nil {
			return 0, 0, err
		}
	}

	wf, wfErr := strconv.ParseFloat(w, 64)
	hf, hFErr := strconv.ParseFloat(h, 64)
	if wfErr == nil && hFErr == nil {
		return int(wf), int(hf), nil
	}

	if hasUnit(w, h, pixel) {
		return convert(w, h, pixel, 1)
	} else if hasUnit(w, h, inch) {
		return convert(w, h, inch, pxPerInch)
	} else if hasUnit(w, h, centimeter) {
		return convert(w, h, centimeter, pxPerCentimeter)
	} else if hasUnit(w, h, millimeter) {
		return convert(w, h, millimeter, pxPerMillimeter)
	} else if hasUnit(w, h, point) {
		return convert(w, h, point, pxPerPoint)
	} else if hasUnit(w, h, pica) {
		return convert(w, h, pica, pxPerPica)
	}
	return 0, 0, fmt.Errorf("not surpported unit: %s, %s", w, h)
}

func whFromViewBox(vb string) (width string, height string, err error) {
	split := strings.Split(vb, " ")
	if len(split) != 4 {
		return "", "", fmt.Errorf("invalid view box: %s", vb)
	}
	return split[2], split[3], nil
}

func convert(w, h, unit string, pxPerUnit float64) (width int, height int, err error) {
	w, h = strings.TrimSuffix(w, unit), strings.TrimRight(h, unit)
	wf, err := strconv.ParseFloat(w, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("%w: parse float: %s", err, w)
	}
	hf, err := strconv.ParseFloat(h, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("%w: parse float: %s", err, h)
	}
	return int(wf * pxPerUnit), int(hf * pxPerUnit), nil
}

func hasUnit(w, h, unit string) bool {
	return strings.HasSuffix(w, unit) && strings.HasSuffix(h, unit)
}
