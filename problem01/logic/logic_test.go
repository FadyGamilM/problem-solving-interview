package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindrome(t *testing.T) {
	testCases := []struct {
		inp         string
		expectedOut bool
	}{
		{
			inp:         "",
			expectedOut: true,
		},
		{
			inp:         "levl",
			expectedOut: false,
		},
		{
			inp:         "level",
			expectedOut: true,
		},
		{
			inp:         "x",
			expectedOut: true,
		},
		{
			inp:         "anna",
			expectedOut: true,
		},
		{
			inp:         "ana",
			expectedOut: true,
		},
		{
			inp:         "rotator",
			expectedOut: true,
		},
		{
			inp:         "Dog doo? Good God!", // -> Dogdoo?GoodGod! -> dogdoogoodgod
			expectedOut: true,
		},
		{
			inp:         "apple",
			expectedOut: false,
		},
		{
			inp:         "aabbccddeeffgghhiijjkkllmmnnooppqqrrssttuuvvwwxxyyzzzzyyxxwwvvuuttssrrqqppoonnmmllkkjjiihhggffeeddccbbaa",
			expectedOut: true,
		},
		{
			inp:         "!@#$XFFX$#@?",
			expectedOut: true,
		},
		{
			inp:         "Aba",
			expectedOut: true,
		},
		{
			inp:         "A man, a plan, a canal: Panama",
			expectedOut: true,
		},
		{
			inp:         "åbba",
			expectedOut: false,
		},
	}
	for i, test := range testCases {
		actualOut := Palindrome(test.inp)
		t.Logf("start test case no.{%v} with input = {%v}", i, test.inp)
		succeed := assert.Equal(t, test.expectedOut, actualOut)
		if succeed {
			t.Logf("test case no.{%v} has succeed ", i)
		}
	}
}
