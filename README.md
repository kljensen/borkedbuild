# Borked Build

Borked Build is a free hosted repository status service. It allows developers
to display a custom repo status similar to those seen used in Travis CI's
continuous integration service.

# How it Works

Currently runs locally only. There are two primary ways to view output
at the moment. The first is to use GO's included goplay package, however,
this requires a small tweak to their package to prevent the output from
getting wrapped in `<pre>` tags. The suggested method is to push the output
into a local `.html` file and view in a browser.

	go run ./svg/svg.go > svg_output.html

In `svg.go` you generate a new icon with the following:

	makeSvg("{left hand text}", "{right hand text}", "{color}")

For example, to create an icon that you would see with Travis CI such as
"build passed", run the following:

	makeSvg("build", "passed", "green")

# Options & Configuration

## Text

You can specify any text for both the left and right hand side of the icon.

## Colors

Current color Options

* Green
* Yellow
* Red

## Fonts

Currently only Droid Sans at 10.5px is supported.

