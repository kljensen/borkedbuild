package main

import (
	"fmt"
	"github.com/ajstarks/svgo"
	"os"
)

func writeLetter(l rune, y int, canvas *svg.SVG) {
	repeat := 20
	runeArray := make([]rune, repeat)
	for i := 0; i < repeat; i++ {
		runeArray[i] = l
	}
	canvas.Text(5, y, string(runeArray), "font-size:10.5px;text-anchor:left;font-family:'Droid Sans';")
}

func main() {
	width := 500
	height := 500
	fmt.Println("<html><body>")
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)

	fontStyle := `
    <style type="text/css">
    <![CDATA[
    @font-face {
    font-family: 'Droid Sans';
    font-style: normal;
    font-weight: 400;
    src: local('Droid Sans'), local('DroidSans'), url(http://themes.googleusercontent.com/static/fonts/droidsans/v3/s-BiyweUPV0v-yRb-cjciC3USBnSvpkopQaUR-2r7iU.ttf) format('truetype');
    }
    ]]>
    </style>
`
	canvas.Def()
	fmt.Fprintf(canvas.Writer, fontStyle)
	canvas.DefEnd()

	canvas.Gstyle("font-size:20pt; font-family:fixed;")
	//canvas.Ellipse(width/2, height, width/2,height/3,"fill:rgb(44,77,232)")
	y := 15
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()"
	for _, c := range chars {
		writeLetter(c, y, canvas)
		y += 15
	}
	canvas.End()
	fmt.Println("</body></html>")
}
