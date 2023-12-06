package logic

import (
	"fmt"
	"strings"
)

type charFreq struct {
	Char rune
	Freq int
}

type SlidingWindow struct {
	encodingOut *strings.Builder
	decodingMap []charFreq
}

func (sw *SlidingWindow) Reset() {
	sw.decodingMap = make([]charFreq, 0)
	sw.encodingOut = &strings.Builder{}
}

func (sw *SlidingWindow) GetDecodingMap() []charFreq {
	return sw.decodingMap
}

func NewSlidingWindow() *SlidingWindow {
	return &SlidingWindow{
		encodingOut: new(strings.Builder),
		decodingMap: make([]charFreq, 0),
	}
}

func (sw *SlidingWindow) RunLengthEncode(inp string) string {
	if len(inp) == 0 {
		return ""
	}
	data := []rune(inp)
	counter := 0
	start := 0
	end := 0

	for {
		// log.Printf("start = %v ➜ %v , end = %v ➜ %v\n", start, string(data[start]), end, string(data[end]))
		if end >= len(data) {
			// log.Println("hit not equale with counter = ", counter, "for character {", string(data[start]), "}")
			sw.encodingOut.WriteString(fmt.Sprintf("%v%v", string(data[start]), counter))
			sw.decodingMap = append(sw.decodingMap, charFreq{Char: data[start], Freq: counter})
			break
		}
		if data[start] == data[end] {
			counter++
			end++
		} else {
			// log.Println("hit not equale with counter = ", counter, "for character {", string(data[start]), "}")
			sw.encodingOut.WriteString(fmt.Sprintf("%v%v", string(data[start]), counter))
			// store this info for decoding stage
			sw.decodingMap = append(sw.decodingMap, charFreq{Char: data[start], Freq: counter})
			// reinitalize the counter so we will start a new window
			counter = 0
			// shift the window to avoid repeated work..
			start = end
		}
	}
	return sw.encodingOut.String()
}

func (sw *SlidingWindow) RunLengthDecode(inp string) string {
	sb := strings.Builder{}
	for _, s := range sw.GetDecodingMap() {
		sb.WriteString(strings.Repeat(string(s.Char), s.Freq))
	}
	return sb.String()
}

// a a a a a a a a a a b  b  b  a  x  x  x  x  y  y  y  z  y  x
// 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23

// ====================================================================================
// ====================================================================================
// ====================================================================================
type BruteForce struct{}

func (bf BruteForce) runLengthEncode(inp string) string {
	return ""
}

func (bf BruteForce) runLengthDecode(inp string) string {
	return ""
}
