package daysago

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetDateStringFromWordString(t *testing.T) {
	tests := []struct {
		inputTime   time.Time
		inputString string
		expected    string
	}{
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "yesterday", expected: "2009-11-09"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "today", expected: "2009-11-10"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "tomorrow", expected: "2009-11-11"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "day after tomorrow", expected: "2009-11-12"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "day before yesterday", expected: "2009-11-08"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "2 days ago", expected: "2009-11-08"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "2 months ago", expected: "2009-09-10"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "2 years ago", expected: "2007-11-10"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "last 2 days", expected: "2009-11-08"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "last 2 months", expected: "2009-09-10"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "last 2 years", expected: "2007-11-10"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "last year", expected: "2008-01-01"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "last month", expected: "2009-10-01"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "past year", expected: "2008-01-01"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "past month", expected: "2009-10-01"},
		{inputTime: time.Date(2017, time.September, 10, 23, 0, 0, 0, time.UTC), inputString: "past week", expected: "2017-09-03"},
		{inputTime: time.Date(2017, time.September, 10, 23, 0, 0, 0, time.UTC), inputString: "last week", expected: "2017-09-03"},
	}

	for _, eachTest := range tests {
		got := GetDateStringFromWordString(eachTest.inputTime, eachTest.inputString)
		assert.Equal(t, eachTest.expected, got, fmt.Sprintf("%s", eachTest.inputString))
	}
}

func TestGetDateStringRangeFromWordString(t *testing.T) {
	tests := []struct {
		inputTime     time.Time
		inputString   string
		expectedStart string
		expectedEnd   string
	}{
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "yesterday", expectedStart: "2009-11-09", expectedEnd: "2009-11-10"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "today", expectedStart: "2009-11-10", expectedEnd: "2009-11-10"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "tomorrow", expectedStart: "2009-11-10", expectedEnd: "2009-11-11"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "day after tomorrow", expectedStart: "2009-11-10", expectedEnd: "2009-11-12"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "day before yesterday", expectedStart: "2009-11-08", expectedEnd: "2009-11-10"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "2 days ago", expectedStart: "2009-11-08", expectedEnd: "2009-11-10"},

		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "2 months ago", expectedStart: "2009-09-01", expectedEnd: "2009-10-31"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "2 years ago", expectedStart: "2007-01-01", expectedEnd: "2008-12-31"},

		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "last 2 days", expectedStart: "2009-11-08", expectedEnd: "2009-11-10"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "last 2 months", expectedStart: "2009-09-01", expectedEnd: "2009-10-31"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "last 2 years", expectedStart: "2007-01-01", expectedEnd: "2008-12-31"},

		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "last year", expectedStart: "2008-01-01", expectedEnd: "2008-12-31"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "last month", expectedStart: "2009-10-01", expectedEnd: "2009-10-31"},

		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "past year", expectedStart: "2008-01-01", expectedEnd: "2008-12-31"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "past month", expectedStart: "2009-10-01", expectedEnd: "2009-10-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "past week", expectedStart: "2017-09-03", expectedEnd: "2017-09-09"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "last week", expectedStart: "2017-09-03", expectedEnd: "2017-09-09"},

		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "this week", expectedStart: "2017-09-10", expectedEnd: "2017-09-16"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "this month", expectedStart: "2017-09-01", expectedEnd: "2017-09-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "this year", expectedStart: "2017-01-01", expectedEnd: "2017-12-31"},

		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "october", expectedStart: "2017-10-01", expectedEnd: "2017-10-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "january", expectedStart: "2017-01-01", expectedEnd: "2017-01-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "jan", expectedStart: "2017-01-01", expectedEnd: "2017-01-31"},

		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "february", expectedStart: "2017-02-01", expectedEnd: "2017-02-28"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "feb", expectedStart: "2017-02-01", expectedEnd: "2017-02-28"},

		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "march", expectedStart: "2017-03-01", expectedEnd: "2017-03-31"},

		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "april", expectedStart: "2017-04-01", expectedEnd: "2017-04-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "apr", expectedStart: "2017-04-01", expectedEnd: "2017-04-30"},

		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "may", expectedStart: "2017-05-01", expectedEnd: "2017-05-31"},

		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "june", expectedStart: "2017-06-01", expectedEnd: "2017-06-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "jun", expectedStart: "2017-06-01", expectedEnd: "2017-06-30"},

		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "july", expectedStart: "2017-07-01", expectedEnd: "2017-07-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "jul", expectedStart: "2017-07-01", expectedEnd: "2017-07-31"},

		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "august", expectedStart: "2017-08-01", expectedEnd: "2017-08-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "aug", expectedStart: "2017-08-01", expectedEnd: "2017-08-31"},

		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "september", expectedStart: "2017-09-01", expectedEnd: "2017-09-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "sep", expectedStart: "2017-09-01", expectedEnd: "2017-09-30"},

		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "october", expectedStart: "2017-10-01", expectedEnd: "2017-10-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "oct", expectedStart: "2017-10-01", expectedEnd: "2017-10-31"},

		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "november", expectedStart: "2017-11-01", expectedEnd: "2017-11-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "nov", expectedStart: "2017-11-01", expectedEnd: "2017-11-30"},

		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "december", expectedStart: "2017-12-01", expectedEnd: "2017-12-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "dec", expectedStart: "2017-12-01", expectedEnd: "2017-12-31"},
	}

	for _, eachTest := range tests {
		gotStart, gotEnd := GetDateStringRangeFromWordString(eachTest.inputTime, eachTest.inputString)
		assert.Equal(t, eachTest.expectedStart, gotStart, fmt.Sprintf("%v Start", eachTest.inputString))
		assert.Equal(t, eachTest.expectedEnd, gotEnd, fmt.Sprintf("%v End", eachTest.inputString))
	}
}
