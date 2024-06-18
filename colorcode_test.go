package ascii_art

import "testing"

func TestColorPicker(t *testing.T) {
	type args struct {
		color string
	}
	tests := []struct {
		name          string
		args          args
		wantColorCode string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotColorCode := ColorPicker(tt.args.color); gotColorCode != tt.wantColorCode {
				t.Errorf("ColorPicker() = %v, want %v", gotColorCode, tt.wantColorCode)
			}
		})
	}
}
