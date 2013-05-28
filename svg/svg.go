package main

import (
	"fmt"
	"github.com/ajstarks/svgo"
	"os"
)

// Return the pixel width of each supported letter / character based on
// precalculation run on Droid Sans 10.5px
func letterWidth(l rune) int {
	switch l {

	case '!', '(', ')', 'J', 'i', 'j', 'l':
		return 3
	case 'I', 'f', 'r', 't', ' ':
		return 4
	case 'F', 'L', 'c', 'k', 's', 'v', 'y', 'z':
		return 5
	case '$', '*', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'E', 'K',
		'P', 'R', 'S', 'T', 'V', 'X', 'Y', 'Z', '^', 'a', 'b', 'd', 'e',
		'g', 'h', 'n', 'o', 'p', 'q', 'u', 'x':
		return 6
	case '#', 'A', 'B', 'C':
		return 7
	case '&', 'D', 'G', 'H', 'N', 'O', 'Q', 'U', 'w':
		return 8
	case '%':
		return 9
	case '@', 'M', 'W', 'm':
		return 10
	}
	return 0
}

// Calculates word length in pixels for
// droid sans fonts at 10.5px font size
func stringWidth(s string) int {
	slen := 0
	for _, l := range s {
		slen += letterWidth(l)
	}
	return slen
}

// Create SVG graphic.
// Requires strings for left hand side text (lhs) and right hand side text (rhs)
// along with string for the color of the right hand side. This is in plain
// english (i.e. "green", "red")
func makeSvg(lhs, rhs, rhs_color string) {
	buffer := 5 // standard pixel padding
	cornerRadius := 2

	lhsWidth := stringWidth(lhs)
	rhsWidth := stringWidth(rhs) - cornerRadius

	// Left side + right side + 4 buffers (both sides of each text) + 2 corners
	width := lhsWidth + rhsWidth + 4*buffer + 2*cornerRadius
	height := 18
	textY := 12 // How far down to push text. 12px based on 10.5px Droid Sans
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)

	fontStyle := `
	    <style type="text/css">
		    <![CDATA[
			    @font-face {
				    font-family: 'Droid Sans';
				    font-style: normal;
				    font-weight: 400;
				    font-size:9px;
				    src: local('Droid Sans'),
				    		local('DroidSans'),
				    		url('http://themes.googleusercontent.com/static/fonts/droidsans/v3/s-BiyweUPV0v-yRb-cjciC3USBnSvpkopQaUR-2r7iU.ttf')
							format('truetype');
			    }
		    ]]>
	    </style>
		`

	canvas.Def()
	fmt.Fprintf(canvas.Writer, fontStyle)
	canvas.DefEnd()

	// Define Linear Gradient style for GRAY left hand side
	fmt.Println(`<linearGradient id="gray" x1="0%" y1="0%" x2="0%" y2="100%" >
        <stop offset="4%" stop-color="#8E8E8E" />
        <stop offset="5%" stop-color="#595959" />
        <stop offset="94%" stop-color="#3d3d3d" />
        <stop offset="95%" stop-color="#2c2c2c" />
        </linearGradient>`)

	// Define Linear Gradient style for right hand side
	//based on rhs_color string
	switch rhs_color {
	case "green":
		fmt.Println(`<linearGradient id="green"
						x1="0%" y1="0%" x2="0%" y2="100%" >
					<stop offset="4%" stop-color="#9feb82" />
					<stop offset="5%" stop-color="#50cf2c" />
					<stop offset="94%" stop-color="#39b524" />
					<stop offset="95%" stop-color="#489139" />
					</linearGradient>`)
	case "yellow":
		fmt.Println(`<linearGradient id="yellow"
						x1="0%" y1="0%" x2="0%" y2="100%" >
					<stop offset="4%" stop-color="#FFEA8E" />
					<stop offset="5%" stop-color="#DCB116" />
					<stop offset="94%" stop-color="#A38310" />
					<stop offset="95%" stop-color="#856B0D" />
					</linearGradient>`)
	case "red":
		fmt.Println(`<linearGradient id="red"
						x1="0%" y1="0%" x2="0%" y2="100%" >
					<stop offset="4%" stop-color="#FFAC9C" />
					<stop offset="5%" stop-color="#D95A41" />
					<stop offset="94%" stop-color="#A24331" />
					<stop offset="95%" stop-color="#8A3A2A" />
					</linearGradient>`)
	}
	canvas.Gstyle("font-size:7.5pt; font-family:fixed;")

	// * * *
	// Write the left hand side
	// * * *
	midPoint := lhsWidth + 2*buffer
	// Left hand cap w/ rounded corners
	canvas.Roundrect(0, 0, 5, height, 4, 4, "fill:url(#gray);")
	// Left hand main rectangle body
	canvas.Rect(cornerRadius, 0, midPoint, height, "fill:url(#gray);")
	// First text layer for shadowing, 1px below textY position
	canvas.Text(cornerRadius+buffer, textY+1, lhs,
		"fill:rgba(0,0,0,0.5);text-anchor:left;font-family:'Droid Sans';")
	// Second text layer (i.e. primary layer on top)
	canvas.Text(cornerRadius+buffer, textY, lhs,
		"fill:#f2f2f2;text-anchor:left;font-family:'Droid Sans';")

	// * * *
	// Write the right hand side
	// * * *
	// Right hand cap with round corners
	canvas.Roundrect(width-3*cornerRadius, 0, 5, height, 4, 4,
		"fill:url(#"+rhs_color+");")
	// Right hand main rectangle body. Color determined by rhs_color
	canvas.Rect(midPoint, 0, rhsWidth+2*buffer,
		height, "fill:url(#"+rhs_color+");")
	// First text layer for shadowing
	canvas.Text(midPoint+buffer, textY+1, rhs,
		"fill:rgba(0,0,0,0.5);text-anchor:left;font-family:'Droid Sans';")
	// Second text layer (i.e. primary layer on top)
	canvas.Text(midPoint+buffer, textY, rhs,
		"fill:#f2f2f2;text-anchor:left;font-family:'Droid Sans';")

	canvas.End()

}

// * * * * * * *
// Create badges
// * * * * * * *
func main() {
	makeSvg("coverage", "100%", "green")
	fmt.Printf("<br /><img src='https://coveralls.io/repos/celluloid/celluloid/badge.png?branch=master'>")
	fmt.Println("<br /><br />")
	makeSvg("dependencies", "update", "red")
	fmt.Printf("<br /><img src='https://raw.github.com/kljensen/borkedbuild/2a30c54609bfd34a5d8ebba62170c0e369c61dc8/img/red.png'>")
	fmt.Println("<br /><br />")
	makeSvg("dependencies", "out-of-date", "yellow")
	fmt.Printf("<br /><img src='https://raw.github.com/kljensen/borkedbuild/2a30c54609bfd34a5d8ebba62170c0e369c61dc8/img/yellow.png'>")
	fmt.Println("<br /><br />")
	fmt.Println("<hr />")
	fmt.Println("<br /><br />")
	makeSvg("build", "passed", "green")
	fmt.Println("<br /><br />")
	makeSvg("guitar volume", "11", "green")
	fmt.Println("<br /><br />")
	makeSvg("build", "borked", "red")

}
