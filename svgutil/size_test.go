package svgutil_test

import (
	"testing"

	"github.com/mathpresso/go-utils/svgutil"
)

func TestSize(t *testing.T) {

	tests := []struct {
		name       string
		strSvg     string
		wantWidth  int
		wantHeight int
		wantErr    bool
	}{
		{
			"empty",
			"",
			0,
			0,
			true,
		},
		{
			"has width height",
			`
<?xml version="1.0" encoding="UTF-8" standalone="no"?>

<svg
   width="799.68781"
   height="1050.2272"
   viewBox="0 0 211.58407 277.8726"
   version="1.1"></svg>
`,
			799,
			1050,
			false,
		},
		{
			"in view box",
			`<svg id="_레이어_2" data-name="레이어 2" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 749.26 604.52"></svg>`,
			749,
			604,
			false,
		},
		{
			"inch",
			`<svg width="100.00001in" height="200.00002in"></svg>`,
			9600,
			19200,
			false,
		},
		{
			"inch in view box",
			`<svg id="_레이어_2" data-name="레이어 2" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 100.00001in 200.00002in"></svg>`,
			9600,
			19200,
			false,
		},
		{
			"cm",
			`<svg width="100.00001cm" height="200.00002cm"></svg>`,
			3779,
			7558,
			false,
		},
		{
			"cm in view box",
			`<svg id="_레이어_2" data-name="레이어 2" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 100.00001cm 200.00002cm"></svg>`,
			3779,
			7558,
			false,
		},
		{
			"mm",
			`<svg width="100.00001mm" height="200.00002mm"></svg>`,
			377,
			755,
			false,
		},
		{
			"mm int view box",
			`<svg id="_레이어_2" data-name="레이어 2" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 100.00001mm 200.00002mm"></svg>`,
			377,
			755,
			false,
		},
		{
			"pt",
			`<svg width="100.00001pt" height="200.00002pt"></svg>`,
			133,
			266,
			false,
		},
		{
			"pt in view box",
			`<svg id="_레이어_2" data-name="레이어 2" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 100.00001pt 200.00002pt"></svg>`,
			133,
			266,
			false,
		},
		{
			"pc",
			`<svg width="100.00001pc" height="200.00002pc"></svg>`,
			1600,
			3200,
			false,
		},
		{
			"pc view box",
			`<svg id="_레이어_2" data-name="레이어 2" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 100.00001pc 200.00002pc"></svg>`,
			1600,
			3200,
			false,
		},
		{
			"nested svg",
			`
<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" version="1.1" height="639"
    width="505" viewBox="(0, 0, 505, 639)">
    <defs />
    <svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" version="1.1" height="639"
        width="505" viewBox="0 -639 505 639">
        <defs />
    </svg>
</svg>`,
			505,
			639,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWidth, gotHeight, err := svgutil.Size(tt.strSvg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Size() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotWidth != tt.wantWidth {
				t.Errorf("Size() gotWidth = %v, want %v", gotWidth, tt.wantWidth)
			}
			if gotHeight != tt.wantHeight {
				t.Errorf("Size() gotHeight = %v, want %v", gotHeight, tt.wantHeight)
			}
		})
	}
}
