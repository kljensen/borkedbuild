package main

import (
	"fmt"
	"github.com/ajstarks/svgo"
	"os"
)

var gray string = "background: #8e8e8e; /  Old browsers / background: -moz-linear-gradient(top, #8e8e8e 1%, #595959 2%, #3d3d3d 98%, #4b4b4b 99%); /  FF3.6+ / background: -webkit-gradient(linear, left top, left bottom, color-stop(1%,#8e8e8e), color-stop(2%,#595959), color-stop(98%,#3d3d3d), color-stop(99%,#4b4b4b)); /  Chrome,Safari4+ / background: -webkit-linear-gradient(top, #8e8e8e 1%,#595959 2%,#3d3d3d 98%,#4b4b4b 99%); /  Chrome10+,Safari5.1+ / background: -o-linear-gradient(top, #8e8e8e 1%,#595959 2%,#3d3d3d 98%,#4b4b4b 99%); /  Opera 11.10+ / background: -ms-linear-gradient(top, #8e8e8e 1%,#595959 2%,#3d3d3d 98%,#4b4b4b 99%); /  IE10+ / background: linear-gradient(to bottom, #8e8e8e 1%,#595959 2%,#3d3d3d 98%,#4b4b4b 99%); /  W3C / filter: progid:DXImageTransform.Microsoft.gradient( startColorstr='#8e8e8e', endColorstr='#4b4b4b',GradientType=0 ); / IE6-9 */"
var green string = "background: #9feb82; /* Old browsers / background: -moz-linear-gradient(top, #9feb82 1%, #50cf2c 2%, #39b524 98%, #489139 99%); /  FF3.6+ / background: -webkit-gradient(linear, left top, left bottom, color-stop(1%,#9feb82), color-stop(2%,#50cf2c), color-stop(98%,#39b524), color-stop(99%,#489139)); /  Chrome,Safari4+ / background: -webkit-linear-gradient(top, #9feb82 1%,#50cf2c 2%,#39b524 98%,#489139 99%); /  Chrome10+,Safari5.1+ / background: -o-linear-gradient(top, #9feb82 1%,#50cf2c 2%,#39b524 98%,#489139 99%); /  Opera 11.10+ / background: -ms-linear-gradient(top, #9feb82 1%,#50cf2c 2%,#39b524 98%,#489139 99%); /  IE10+ / background: linear-gradient(to bottom, #9feb82 1%,#50cf2c 2%,#39b524 98%,#489139 99%); /  W3C / filter: progid:DXImageTransform.Microsoft.gradient( startColorstr='#9feb82', endColorstr='#489139',GradientType=0 ); /  IE6-9 /"

func letterWidth(l rune) int {
	switch l {

	case '!', '(', ')', 'J', 'i', 'j', 'l':
		return 3
	case 'I', 'f', 'r', 't', ' ':
		return 4
	case 'F', 'L', 'c', 'k', 's', 'v', 'y', 'z':
		return 5
	case '$', '*', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'E', 'K', 'P', 'R', 'S', 'T', 'V', 'X', 'Y', 'Z', '^', 'a', 'b', 'd', 'e', 'g', 'h', 'n', 'o', 'p', 'q', 'u', 'x':
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

func makeSvg(lhs, rhs, rhs_color string) {
	buffer := 5
	cornerRadius := 2
	textY := 12

	lhsWidth := stringWidth(lhs)
	rhsWidth := stringWidth(rhs) - cornerRadius
	width := lhsWidth + rhsWidth + 4*buffer + 2*cornerRadius
	height := 18
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
    src: local('Droid Sans'), local('DroidSans'), url(http://themes.googleusercontent.com/static/fonts/droidsans/v3/s-BiyweUPV0v-yRb-cjciC3USBnSvpkopQaUR-2r7iU.ttf) format('truetype');
    }
    ]]>
    </style>
`
	canvas.Def()
	fmt.Fprintf(canvas.Writer, fontStyle)
	canvas.DefEnd()

	fmt.Println(`<linearGradient id="gray" x1="0%" y1="0%" x2="0%" y2="100%" >
        <stop offset="4%" stop-color="#8E8E8E" />
        <stop offset="5%" stop-color="#595959" />
        <stop offset="94%" stop-color="#3d3d3d" />
        <stop offset="95%" stop-color="#2c2c2c" />
        </linearGradient>`)

	switch rhs_color {
	case "green":
		fmt.Println(`<linearGradient id="green" x1="0%" y1="0%" x2="0%" y2="100%" >
            <stop offset="4%" stop-color="#9feb82" />
            <stop offset="5%" stop-color="#50cf2c" />
            <stop offset="94%" stop-color="#39b524" />
            <stop offset="95%" stop-color="#489139" />
            </linearGradient>`)
	case "yellow":
		fmt.Println(`<linearGradient id="yellow" x1="0%" y1="0%" x2="0%" y2="100%" >
            <stop offset="4%" stop-color="#FFEA8E" />
            <stop offset="5%" stop-color="#DCB116" />
            <stop offset="94%" stop-color="#A38310" />
            <stop offset="95%" stop-color="#856B0D" />
            </linearGradient>`)
	case "red":
		fmt.Println(`<linearGradient id="red" x1="0%" y1="0%" x2="0%" y2="100%" >
            <stop offset="4%" stop-color="#FFAC9C" />
            <stop offset="5%" stop-color="#D95A41" />
            <stop offset="94%" stop-color="#A24331" />
            <stop offset="95%" stop-color="#8A3A2A" />
            </linearGradient>`)
	}
	canvas.Gstyle("font-size:7.5pt; font-family:fixed;")

	// Write the left hand side
	midPoint := lhsWidth + 2*buffer
	canvas.Roundrect(0, 0, 5, height, 4, 4, "fill:url(#gray);")
	canvas.Rect(cornerRadius, 0, midPoint, height, "fill:url(#gray);")
	canvas.Text(cornerRadius+buffer, textY+1, lhs, "fill:rgba(0,0,0,0.5);text-anchor:left;font-family:'Droid Sans';")
	canvas.Text(cornerRadius+buffer, textY, lhs, "fill:#F4F4F4;text-anchor:left;font-family:'Droid Sans';")

	// Write the right hand side
	canvas.Rect(midPoint, 0, rhsWidth+2*buffer, height, "fill:url(#"+rhs_color+");")
	canvas.Text(midPoint+buffer, textY+1, rhs, "fill:rgba(0,0,0,0.5);text-anchor:left;font-family:'Droid Sans';")
	canvas.Text(midPoint+buffer, textY, rhs, "fill:#F8FBF7;text-anchor:left;font-family:'Droid Sans';")
	canvas.Roundrect(width-3*cornerRadius, 0, 5, height, 4, 4, "fill:url(#"+rhs_color+");")

	canvas.End()

}

func main() {
	makeSvg("coverage", "100%", "green")
	fmt.Printf("<br><img src='https://coveralls.io/repos/celluloid/celluloid/badge.png?branch=master'>")
	fmt.Println("<br>")
	makeSvg("dependencies", "update", "red")
	fmt.Printf("<br><img src='https://raw.github.com/kljensen/borkedbuild/2a30c54609bfd34a5d8ebba62170c0e369c61dc8/img/red.png'>")
	fmt.Println("<br>")
	makeSvg("dependencies", "out-of-date", "yellow")
	fmt.Printf("<br><img src='https://raw.github.com/kljensen/borkedbuild/2a30c54609bfd34a5d8ebba62170c0e369c61dc8/img/yellow.png'>")
	fmt.Println("<br>")
	makeSvg("build", "passed", "green")
	fmt.Println("<br>")
	makeSvg("guitar volume", "11", "green")

}
