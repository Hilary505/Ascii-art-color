package ascii_art

import "testing"

func TestProcessInput(t *testing.T) {
	type args struct {
		contents  []string
		input     string
		color     string
		subString string
	}
	tests := []struct {
		name       string
		args       args
		wantStrArt string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotStrArt := ProcessInput(tt.args.contents, tt.args.input, tt.args.color, tt.args.subString); gotStrArt != tt.wantStrArt {
				t.Errorf("ProcessInput() = %v, want %v", gotStrArt, tt.wantStrArt)
			}
		})
	}
}
