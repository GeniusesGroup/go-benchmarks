/* For license and copyright information please see LEGAL file in repository */

package unix

import (
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	tests := []struct {
		name     string
		wantSec  int64
		wantNsec int32
		wantMono int64
	}{
		{
			name:    "now",
			wantSec: time.Now().Unix(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSec, gotNsec, gotMono := Now()
			if gotSec != tt.wantSec {
				t.Errorf("Now() gotSec = %v, want %v", gotSec, tt.wantSec)
			}
			if gotNsec != tt.wantNsec {
				// t.Errorf("Now() gotNsec = %v, want %v", gotNsec, tt.wantNsec)
			}
			if gotMono != tt.wantMono {
				// t.Errorf("Now() gotMono = %v, want %v", gotMono, tt.wantMono)
			}
		})
	}
}
