package notification

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsHolyday(t *testing.T) {
	tests := []struct {
		input    time.Time
		expected bool
	}{
		{time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC), true},    // 2021-01-01 元日
		{time.Date(2021, time.February, 23, 0, 0, 0, 0, time.UTC), true},  // 2021-02-23 天皇誕生日
		{time.Date(2021, time.July, 22, 0, 0, 0, 0, time.UTC), true},      // 2021-07-22 海の日
		{time.Date(2021, time.July, 23, 0, 0, 0, 0, time.UTC), true},      // 2021-07-23 スポーツの日
		{time.Date(2021, time.August, 8, 0, 0, 0, 0, time.UTC), true},     // 2021-08-08 山の日
		{time.Date(2021, time.August, 9, 0, 0, 0, 0, time.UTC), true},     // 2021-08-09 山の日振替休日
		{time.Date(2021, time.December, 23, 0, 0, 0, 0, time.UTC), false}, // 2021-12-23 天皇誕生日ではない
		{time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC), true},    // 2022-01-01 元日
		{time.Date(2022, time.July, 18, 0, 0, 0, 0, time.UTC), true},      // 2022-07-18 海の日
		{time.Date(2022, time.July, 22, 0, 0, 0, 0, time.UTC), false},     // 2022-07-22 海の日ではない
		{time.Date(2022, time.July, 23, 0, 0, 0, 0, time.UTC), false},     // 2022-07-23 スポーツの日ではない
		{time.Date(2022, time.August, 8, 0, 0, 0, 0, time.UTC), false},    // 2022-08-08 山の日ではない
		{time.Date(2022, time.August, 11, 0, 0, 0, 0, time.UTC), true},    // 2022-08-11 山の日
		{time.Date(2022, time.October, 10, 0, 0, 0, 0, time.UTC), true},   // 2022-10-10 スポーツの日
	}

	for _, tt := range tests {
		assert.Equal(t, isHolyday(tt.input), tt.expected)
	}
}
