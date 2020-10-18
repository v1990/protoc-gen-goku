package helper

import "testing"

func TestFormat(t *testing.T) {
	tests := []struct {
		name    string
		args    Args
		want    string
		wantErr bool
	}{
		{"a", Args{1, "%d"}, "1", false},
		//{"b", Args{1,2, "%d"}, "1", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Format(tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Format() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Format() got = %v, want = %v", got, tt.want)
			}
		})
	}
}
