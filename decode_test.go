package tson

import (
	"testing"
	"time"
)

type Foo struct {
	T *time.Time
}

func TestUnmarshal(t *testing.T) {
	type args struct {
		jsonBytes []byte
		expected  *time.Time
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "null",
			args: args{
				jsonBytes: []byte(`{"T": null}`),
				expected:  nil,
			},
		},
		{
			name: "empty",
			args: args{
				jsonBytes: []byte(`{}`),
				expected:  nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var v Foo

			if err := Unmarshal(tt.args.jsonBytes, &v); (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}

			if v.T == nil {
				if tt.args.expected == nil {
					return
				}

				t.Fatalf("time was nil(expected: %v)", tt.args.expected)
			}

			if tt.args.expected == nil {
				t.Fatalf("time was not nil(actual: %v)", v.T)
			}

			if !tt.args.expected.Equal(*v.T) {
				t.Fatalf("time mismatched(want: %v, got: %v)", tt.args.expected, v.T)
			}
		})
	}
}
