# Borked Build

Borked Build is a free hosted repository status service. It allows developers
to display a custom repo status similar to those seen used in Travis CI's
continuous integration service.

# How it Works

Currently runs locally only.

## Start Server

	go run main.go

## Access API
Server is now running on port 3000. The end point is as follows:

	localhost:3000/{left_word}/{right_word}/{color}

Example:

	localhost:3000/Dependencies/Borked/red


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

