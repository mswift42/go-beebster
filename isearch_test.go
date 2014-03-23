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
