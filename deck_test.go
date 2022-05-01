package main

import "testing"

func TestNewDesk(t *testing.T){
    d := newDeck()

    if len(d) != 16 {
        t.Errorf("Expected deck length of 16, got %d", len(d))
    }
}