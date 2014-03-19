package main

import (
        "testing"
)

var rstring1 = "987 http://www.bbc.co.uk/iplayer/images/episode/b03wtmlq_150_84.jpg Silk: Series 3 Episode 1"

var rstring2 = "988 http://www.bbc.co.uk/iplayer/images/episode/b03xkk94_150_84.jpg Silk: Series 3 Episode 2"

// func TestIndex(t *testing.T) {
//         t1 := &Ipsearch{searchterm: rstring1}
//         t2 := &Ipsearch{searchterm: rstring2}
//         indext1 := t1.Index()[0]
//         if indext1 != "987" {
//                 t.Error("Expted \"987\" got: ", indext1)
//         } else if indext2 := t2.Index()[0]; indext2 != "988" {
//                 t.Error("expteced \"988\" got: ", indext2)
//         }
// }
func TestIndex(t *testing.T) {
        t1 := index(rstring1)
        t2 := index(rstring2)
        if t1 != "987" {
                t.Error("Expected 987 got: ", t1)
        } else if t2 != "988" {
                t.Error("Expected 988 got: ", t2)
        }
}

// test Thumbmail method
// func TestThumbnail(t *testing.T) {
//         t1 := &Ipsearch{searchterm: rstring1}
//         thumbt1 := t1.ThumbNail()[0]
//         if thumbt1 != "http://www.bbc.co.uk/iplayer/images/episode/b03wtmlq_150_84.jpg" {
//                 t.Error("Expected some thumbnail, got ", thumbt1)
//         }
// }
