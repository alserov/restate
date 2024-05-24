package service

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestSelectAvailableTStampsForMeeting(t *testing.T) {
	tests := []struct {
		Stamps []time.Time

		BeforeStamps []time.Time
		InnerStamps  []time.Time
		OuterStamps  []time.Time
	}{
		{
			Stamps: []time.Time{
				time.Date(2000, 9, 9, 12, 0, 0, 0, time.UTC),
				time.Date(2000, 9, 9, 13, 30, 0, 0, time.UTC)},
			BeforeStamps: []time.Time{
				time.Date(2000, 9, 9, 10, 30, 0, 0, time.UTC),
				time.Date(2000, 9, 9, 9, 0, 0, 0, time.UTC),
			},
			InnerStamps: []time.Time{},
			OuterStamps: []time.Time{
				time.Date(2000, 9, 9, 15, 0, 0, 0, time.UTC),
				time.Date(2000, 9, 9, 16, 30, 0, 0, time.UTC),
			},
		},
		{
			Stamps: []time.Time{
				time.Date(2000, 9, 9, 12, 0, 0, 0, time.UTC),
				time.Date(2000, 9, 9, 15, 00, 0, 0, time.UTC)},
			BeforeStamps: []time.Time{
				time.Date(2000, 9, 9, 10, 30, 0, 0, time.UTC),
				time.Date(2000, 9, 9, 9, 0, 0, 0, time.UTC),
			},
			InnerStamps: []time.Time{time.Date(2000, 9, 9, 13, 30, 0, 0, time.UTC)},
			OuterStamps: []time.Time{time.Date(2000, 9, 9, 16, 30, 0, 0, time.UTC)},
		},
		{
			Stamps: []time.Time{
				time.Date(2000, 9, 9, 12, 0, 0, 0, time.UTC),
				time.Date(2000, 9, 9, 15, 00, 0, 0, time.UTC)},
			BeforeStamps: []time.Time{
				time.Date(2000, 9, 9, 10, 30, 0, 0, time.UTC),
				time.Date(2000, 9, 9, 9, 0, 0, 0, time.UTC),
			},
			InnerStamps: []time.Time{time.Date(2000, 9, 9, 13, 30, 0, 0, time.UTC)},
			OuterStamps: []time.Time{time.Date(2000, 9, 9, 16, 30, 0, 0, time.UTC)},
		},
	}

	for idx, tc := range tests {
		t.Run(fmt.Sprintf("tc: %d", idx), func(t *testing.T) {
			stamps := selectAvailableTStampsForMeeting(tc.Stamps)

			for i, befTStamp := range tc.BeforeStamps {
				require.Equal(t, befTStamp, stamps[i])
			}

			for i, inTStamp := range tc.InnerStamps {
				require.Equal(t, inTStamp, stamps[i+len(tc.BeforeStamps)])
			}

			for i, outTStamp := range tc.OuterStamps {
				require.Equal(t, outTStamp, stamps[i+len(tc.InnerStamps)+len(tc.BeforeStamps)])
			}
		})
	}
}
