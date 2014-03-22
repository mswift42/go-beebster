package main

import (
        "testing"
)

var rstring1 = "987 http://www.bbc.co.uk/iplayer/images/episode/b03wtmlq_150_84.jpg Silk: Series 3 Episode 1"

var rstring2 = "988 http://www.bbc.co.uk/iplayer/images/episode/b03xkk94_150_84.jpg Silk: Series 3 Episode 2"

func TestIndex(t *testing.T) {
        t1 := index(rstring1)
        t2 := index(rstring2)
        if t1 != "987" {
                t.Error("Expected 987 got: ", t1)
        } else if t2 != "988" {
                t.Error("Expected 988 got: ", t2)
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

func TestCategory(t *testing.T) {
        var t1 = Categoryinit()
        t2 := Categoryinit()
        t1name := t1[0]
        t2first := t2[0]
        if t1name.Name != "popular" {
                t.Error("Expected <popular>, got: ", t1name.Name)
        } else if t2first.Url != "/?category=popular" {
                t.Error("Expected <?category=popular>, got: ", t2first.Url)
        }
}
