package hw

import "testing"

func TestGeom_CalculateDistance(t *testing.T) {
	tests := []struct {
		name         string
		geom         geom
		wantDistance float64
	}{
		{
			name:         "#1",
			geom:         geom{X1: 1, Y1: 1, X2: 4, Y2: 5},
			wantDistance: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDistance := tt.geom.CalculateDistance(); gotDistance != tt.wantDistance {
				t.Errorf("geom.CalculateDistance() = %v, want %v", gotDistance, tt.wantDistance)
			}
		})
	}
}
