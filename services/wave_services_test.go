package services

import "testing"

func TestIsValidWaveStatus(t *testing.T) {

	tests := []struct {
		name   string
		status string
		want   bool
	}{
		{
			name:   "valid waiting status",
			status: "WAITING",
			want:   true,
		},
		{
			name:   "valid processing status",
			status: "PROCESSING",
			want:   true,
		},
		{
			name:   "valid completed status",
			status: "COMPLETED",
			want:   true,
		},
		{
			name:   "invalid cancelled status",
			status: "CANCELLED",
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidWaveStatus(tt.status)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
