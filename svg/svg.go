package main

import (
    "os"
    "github.com/ajstarks/svgo"
    "fmt"
)
func main() {
    width := 200
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
	src: local('Droid Sans'), local('DroidSans'), url(http://themes.googleusercontent.com/static/fonts/droidsans/v3/s-BiyweUPV0v-yRb-cjciC3USBnSvpkopQaUR-2r7iU.ttf) format('truetype');
	}
	]]>
	</style>
`
	canvas.Def()
	fmt.Fprintf(canvas.Writer, fontStyle)
	canvas.DefEnd()

    canvas.Gstyle("font-size:20pt; font-family:fixed;")
    canvas.Rect(0, 0, width/2, height)
    //canvas.Ellipse(width/2, height, width/2,height/3,"fill:rgb(44,77,232)")             
    canvas.Text(5, 12, "coverage", "fill:white;font-size:10.5px;text-anchor:left;font-family:'Droid Sans';")
    canvas.End()
    fmt.Printf("<br><img src='https://coveralls.io/repos/celluloid/celluloid/badge.png?branch=master'>")
}
