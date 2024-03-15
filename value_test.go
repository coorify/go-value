package value

import (
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	type object0 struct {
		A string
		B int
	}

	type object1 struct {
		C string
	}

	type object2 struct {
		D string
	}

	type object3 struct {
		object1

		o0 object0
		O0 object0

		O2 []object2
		O3 map[string]string
		O4 map[int]string
	}

	src := object3{
		object1: object1{
			C: "cccc",
		},
		o0: object0{
			A: "aaaa",
			B: -9999,
		},
		O0: object0{
			A: "AAAA",
			B: 999,
		},
		O2: []object2{
			{
				D: "D1",
			},
			{
				D: "D2",
			},
		},
		O3: map[string]string{
			"1": "O3 1",
		},
		O4: map[int]string{
			2: "O4 1",
		},
	}

	type args struct {
		src  interface{}
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "struct exported field",
			args: args{
				src:  src,
				path: "O0",
			},
			want:    src.O0,
			wantErr: false,
		},
		{
			name: "struct unexported field",
			args: args{
				src:  src,
				path: "o0",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "struct field O0.A",
			args: args{
				src:  src,
				path: "O0.A",
			},
			want:    src.O0.A,
			wantErr: false,
		},
		{
			name: "struct field O0.B",
			args: args{
				src:  src,
				path: "O0.B",
			},
			want:    src.O0.B,
			wantErr: false,
		},
		{
			name: "struct embed 01",
			args: args{
				src:  src,
				path: "C",
			},
			want:    src.C,
			wantErr: false,
		},
		{
			name: "struct embed 02",
			args: args{
				src:  src,
				path: "object1.C",
			},
			want:    src.object1.C,
			wantErr: false,
		},
		{
			name: "struct array",
			args: args{
				src:  src,
				path: "O2.0.D",
			},
			want:    src.O2[0].D,
			wantErr: false,
		},
		{
			name: "struct map 01",
			args: args{
				src:  src,
				path: "O3.1",
			},
			want:    src.O3["1"],
			wantErr: false,
		},
		{
			name: "struct map 02",
			args: args{
				src:  src,
				path: "O4.2",
			},
			want:    src.O4[2],
			wantErr: false,
		},
		{
			name: "struct nil",
			args: args{
				src:  src,
				path: "O0.C",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "struct nil",
			args: args{
				src:  src,
				path: "O0.C.D",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "struct nil",
			args: args{
				src:  src,
				path: "O0.A.D",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.args.src, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("Get() = %#v, want %#v", got, tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
