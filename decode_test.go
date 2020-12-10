package tson

import (
	"reflect"
	"testing"
	"time"
)

type String string

type Foo struct {
	T       *time.Time
	Strings []String
	Ints    []int
}

func TestUnmarshal(t *testing.T) {
	type args struct {
		jsonBytes []byte
		expected  *time.Time
	}
	tests := []struct {
		name    string
		args    args
		wantFoo Foo
		wantErr bool
	}{
		{
			name: "null",
			args: args{
				jsonBytes: []byte(`{"T": null, "Strings": ["a", "b"], "Ints": [1, 2, 3]}`),
				expected:  nil,
			},
			wantFoo: Foo{
				T:       nil,
				Strings: []String{"a", "b"},
				Ints:    []int{1, 2, 3},
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

			if tt.wantFoo.T != v.T {
				t.Fatalf("expected = %v, actual = %v", tt.wantFoo.T, v.T)
			}

			if !reflect.DeepEqual(tt.wantFoo.Strings, v.Strings) {
				t.Fatalf("expected = %v, actual = %v", tt.wantFoo.Strings, v.Strings)
			}

			if !reflect.DeepEqual(tt.wantFoo.Ints, v.Ints) {
				t.Fatalf("expected = %v, actual = %v", tt.wantFoo.Ints, v.Ints)
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
