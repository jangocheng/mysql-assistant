package foundations

import "testing"

func TestHidePhone(t *testing.T) {
	type args struct {
		phone string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HidePhone(tt.args.phone); got != tt.want {
				t.Errorf("HidePhone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHidePhoneInt(t *testing.T) {
	type args struct {
		phone int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HidePhoneInt(tt.args.phone); got != tt.want {
				t.Errorf("HidePhoneInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
