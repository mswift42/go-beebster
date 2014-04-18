package main

import (
	"testing"
)

var rstring1 = "987 http://www.bbc.co.uk/iplayer/images/episode/b03wtmlq_150_84.jpg Silk: Series 3 Episode 1"

var rstring2 = "988 http://www.bbc.co.uk/iplayer/images/episode/b03xkk94_150_84.jpg Silk: Series 3 Episode 2"

var oldrec1 = `Matches:
390:	From This Day Forward - -, BBC Two, Drama,Films,Relationships & Romance,TV, default
980:	Shooting Dogs - -, BBC Two, Drama,Films,Guidance,HD,TV,War & Disaster, default,
1048:	Suspicion - -, BBC Two, Drama,Films,TV,Thriller, default
1212:	The Westerner - -, BBC Two, Action & Adventure,Drama,Films,TV, default

INFO: 4 Matching Programmes

These programmes should be deleted:
-----------------------------------
Coco Chanel & Igor Stravinsky - Coco Chanel & Igor Stravinsky, Recorded: 30 days 8 hours ago
-----------------------------------
Do you wish to delete them now (Yes/No) ?`

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

desc:           Live coverage of the announcement of Commons business for the week ahead and questions to leader of the Commons Andrew Lansley.
descmedium:     Live coverage of the announcement of Commons business for the week ahead and questions to leader of the Commons Andrew Lansley.
descshort:      Live coverage of the announcement of Commons business for the week ahead.
index:          234
lastbcastrel:   default: 3 days 0 hours ago
longname:       Business Questions
modes:          default: flashhigh1,flashhigh2,flashlow1,flashlow2,flashstd1,flashstd2,flashvhigh1,flashvhigh2,rtsphigh1,rtsphigh2,rtsplow1,rtsplow2,rtspstd1,rtspstd2,rtspvhigh1,rtspvhigh2
pid:            b03yrkz5
player:         http://www.bbc.co.uk/iplayer/episode/b03yrkz5/Business_Questions_20_03_2014/
senum:          s01e01
seriesnum:      1
thumbfile:      /home/severin/Business_Questions_-_20_03_2014_b03yrkz5_default.jpg
thumbnail4:     http://ichef.bbci.co.uk/programmeimages/p01lcfn2/b03yrkz5_512_288.jpg
timeadded:      2 days 18 hours ago (1395335954)
title:          Business Questions: 20/03/2014
type:           tv
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
func TestIsOldRecording(t *testing.T) {
	t1 := isoldRec(oldrec1)
	if t1 != true {
		t.Error("Exptected false, got: ", t1)
	}
}
func TestOldRecording(t *testing.T) {
	t1 := listOldRecordings(oldrec1)
	if t1 != `These programmes should be deleted:
-----------------------------------
Coco Chanel & Igor Stravinsky - Coco Chanel & Igor Stravinsky, Recorded: 30 days 8 hours ago
-----------------------------------
Do you wish to delete them now (Yes/No) ?` {
		t.Error("Expected hoden, got: ", t1)
	}
}

func TestOldRec(t *testing.T) {
	old := listOldRecordings(oldrec1)
	oldslice := oldrecslice(old)
	if oldslice[0] != "Coco" {
		t.Error("Expected <Coco> , got: ", oldslice)
	}
}
