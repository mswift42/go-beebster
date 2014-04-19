Go-Beebster
=============================

Go-Beebster is a Web Gui frontend for [get_iplayer](http://www.infradead.org/get_iplayer/html/get_iplayer.html).

Introduction
============

get-iplayer is a wonderful, wonderful program. However, it can be a bit tedious to constantly type <code>get-iplayer --category crime</code>, or <code>get-iplayer -i 345 | grep desc:</code>.

Thus beebster was born. You can search for a programme, or pick a category (e.g. crime) from the navbar.

For downloading a programme, you can pick the desired quality,
available modes are `flashhd` (highest, but only available for some programmes),
`flashvhigh`, `flashhigh` to `flashlow` with the lowest quality.


Installation:
=============

Install [go](http://golang.org/)

Install [get-iplayer](https://github.com/dinkypumpkin/get_iplayer).

Install the [martini](https://github.com/codegangsta/martini) dependencies:

<code>go get github.com/codegangsta/martini</code>
<code>go get github.com/martini-contrib/render</code>

Clone this repository:

<code>git clone https://github.com/mswift42/go-beebster/</code>

To build go-beebster, cd into the directory and:

<code>go build</code>

To run the programme, change into the directory and

<code>./go-beebster</code>


Screenshots:
============

![category](https://github.com/mswift42/go-beebster/raw/master/Screenshot-cat.png)

![info](https://github.com/mswift42/go-beebster/raw/master/Screenshot-info.png)
