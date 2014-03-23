package main

import (
        "testing"
)

var rstring1 = "987 http://www.bbc.co.uk/iplayer/images/episode/b03wtmlq_150_84.jpg Silk: Series 3 Episode 1"

var rstring2 = "988 http://www.bbc.co.uk/iplayer/images/episode/b03xkk94_150_84.jpg Silk: Series 3 Episode 2"

func TestIndex(t *testing.T) {
        t1 := index(rstring1)
        t2 := index(rstring2)
        if t1 != "/info?index=987" {
                t.Error("Expected /info?index=987 got: ", t1)
        } else if t2 != "/info?index=988" {
                t.Error("Expected ?info?index=988 got: ", t2)
        }
}

func TestThumbnail(t *testing.T) {
        t1 := thumbnail(rstring1)
        t2 := thumbnail(rstring2)

        if t1 != "http://www.bbc.co.uk/iplayer/images/episode/b03wtmlq_150_84.jpg" {
                t.Error("Expected http://www.bbc.co.uk/iplayer/images/episode/b03wtmlq_150_84.jpg, got ", t1)
        } else if t2 != "http://www.bbc.co.uk/iplayer/images/episode/b03xkk94_150_84.jpg" {
                t.Error("Expected http://www.bbc.co.uk/iplayer/images/episode/b03xkk94_150_84.jpg, got ", t2)
        }
}

func TestTitle(t *testing.T) {
        t1 := title(rstring1)
        t2 := title(rstring2)
        if t1 != "Silk: Series 3 Episode 1" {
                t.Error("Expeced Silk: Series 3 Episode 1, got: ", t1)
        } else if t2 != "Silk: Series 3 Episode 2" {
                t.Error("Expected Silk: Series 3 Episode 2, got: ", t2)
        }
}

var infostring = `get_iplayer v2.83, Copyright (C) 2008-2010 Phil Lewis
  This program comes with ABSOLUTELY NO WARRANTY; for details use --warranty.
  This is free software, and you are welcome to redistribute it under certain
  conditions; use --conditions for details.

Matches:
234:	Business Questions - 20/03/2014, BBC Parliament, Factual,News,Politics,TV, default,
INFO: File name prefix = Business_Questions_-_20_03_2014_b03yrkz5_default

available:      Unknown
categories:     News,Factual,Politics
channel:        BBC Parliament
desc:           Live coverage of the announcement of Commons business for the week ahead and questions to leader of the Commons Andrew Lansley.
descmedium:     Live coverage of the announcement of Commons business for the week ahead and questions to leader of the Commons Andrew Lansley.
descshort:      Live coverage of the announcement of Commons business for the week ahead.
dir:            /home/severin
dldate:         2014-03-23
dltime:         12:18:19
duration:       2700
durations:      default: 2700
episode:        20/03/2014
episodenum:     1
episodeshort:   20/03/2014
expiry:         2014-03-27T12:14:00Z
expiryrel:      in 3 days 23 hours
ext:            EXT
filename:       /home/severin/Business_Questions_-_20_03_2014_b03yrkz5_default.EXT
filepart:       /home/severin/Business_Questions_-_20_03_2014_b03yrkz5_default.partial.EXT
fileprefix:     Business_Questions_-_20_03_2014_b03yrkz5_default
firstbcast:     default: 2014-03-20T11:30:00Z
firstbcastrel:  default: 3 days 0 hours ago
index:          234
lastbcast:      default: 2014-03-20T11:30:00Z
lastbcastrel:   default: 3 days 0 hours ago
longname:       Business Questions
modes:          default: flashhigh1,flashhigh2,flashlow1,flashlow2,flashstd1,flashstd2,flashvhigh1,flashvhigh2,rtsphigh1,rtsphigh2,rtsplow1,rtsplow2,rtspstd1,rtspstd2,rtspvhigh1,rtspvhigh2
modesizes:      default: flashhigh1=262MB,flashhigh2=262MB,flashlow1=131MB,flashlow2=131MB,flashstd1=158MB,flashstd2=158MB,flashvhigh1=494MB,flashvhigh2=494MB,rtsphigh1=262MB,rtsphigh2=262MB,rtsplow1=131MB,rtsplow2=131MB,rtspstd1=158MB,rtspstd2=158MB,rtspvhigh1=494MB,rtspvhigh2=494MB
name:           Business Questions
nameshort:      Business Questions
pid:            b03yrkz5
player:         http://www.bbc.co.uk/iplayer/episode/b03yrkz5/Business_Questions_20_03_2014/
senum:          s01e01
seriesnum:      1
thumbfile:      /home/severin/Business_Questions_-_20_03_2014_b03yrkz5_default.jpg
thumbnail:      http://www.bbc.co.uk/iplayer/images/episode/b03yrkz5_150_84.jpg
thumbnail1:     http://ichef.bbci.co.uk/programmeimages/p01lcfn2/b03yrkz5_86_48.jpg
thumbnail2:     http://ichef.bbci.co.uk/programmeimages/p01lcfn2/b03yrkz5_150_84.jpg
thumbnail3:     http://ichef.bbci.co.uk/programmeimages/p01lcfn2/b03yrkz5_178_100.jpg
thumbnail4:     http://ichef.bbci.co.uk/programmeimages/p01lcfn2/b03yrkz5_512_288.jpg
thumbnail5:     http://ichef.bbci.co.uk/programmeimages/p01lcfn2/b03yrkz5_528_297.jpg
thumbnail6:     http://ichef.bbci.co.uk/programmeimages/p01lcfn2/b03yrkz5_640_360.jpg
timeadded:      2 days 18 hours ago (1395335954)
title:          Business Questions: 20/03/2014
type:           tv
verpids:        default: b03yrkmx
version:        default
versions:       default
web:            http://www.bbc.co.uk/programmes/b03yrkz5.html


INFO: 1 Matching Programmes
`

func TestThumb4(t *testing.T) {
        t1 := IplayerIndex{infostring}
        if t1.Thumb4() != "http://ichef.bbci.co.uk/programmeimages/p01lcfn2/b03yrkz5_512_288.jpg" {
                t.Error("Expected http://ichef.bbci.co.uk/programmeimages/p01lcfn2/b03yrkz5_512_288.jpg, got: ", t1.Thumb4())
        }
}
func TestDescription(t *testing.T) {
        t1 := IplayerIndex{infostring}
        if t1.Description() != "Live coverage of the announcement of Commons business for the week ahead and questions to leader of the Commons Andrew Lansley." {
                t.Error("Expected Live coverage of the announcement of Commons business for the week ahead and questions to leader of the Commons Andrew Lansley., got: ", t1.Description())
        }
}
func TestInfoTitle(t *testing.T) {
        t1 := IplayerIndex{infostring}
        if t1.Title() != "Business Questions: 20/03/2014" {
                t.Error("Expected Business Questions: 20/03/2014, got: ", t1.Title())
        }
}
