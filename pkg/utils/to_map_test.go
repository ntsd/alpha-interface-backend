package utils

import (
	"reflect"
	"testing"

	"github.com/AlekSi/pointer"
)

type TestStruct struct {
	TestString        string  `form:"string,limit=10"`
	TestPointerString *string `form:"pointerString,required"`
	TestInt           int     `form:"int"`
	TestPointerInt    *int    `form:"pointerInt"`
}

func TestToMapString(t *testing.T) {
	type args struct {
		in  interface{}
		tag string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			name: "success struct",
			args: args{
				in: TestStruct{
					TestString:        "foo",
					TestPointerString: pointer.ToString("bar"),
					TestInt:           10,
					TestPointerInt:    pointer.ToInt(15),
				},
				tag: "form",
			},
			want: map[string]string{
				"string":        "foo",
				"pointerString": "bar",
				"int":           "10",
				"pointerInt":    "15",
			},
			wantErr: false,
		},
		{
			name: "success pointer struct",
			args: args{
				in: &TestStruct{
					TestString:        "foo",
					TestPointerString: pointer.ToString("bar"),
					TestInt:           10,
					TestPointerInt:    pointer.ToInt(15),
				},
				tag: "form",
			},
			want: map[string]string{
				"string":        "foo",
				"pointerString": "bar",
				"int":           "10",
				"pointerInt":    "15",
			},
			wantErr: false,
		},
		{
			name: "success without string",
			args: args{
				in: &TestStruct{
					TestPointerString: pointer.ToString("bar"),
					TestInt:           10,
					TestPointerInt:    pointer.ToInt(15),
				},
				tag: "form",
			},
			want: map[string]string{
				"string":        "",
				"pointerString": "bar",
				"int":           "10",
				"pointerInt":    "15",
			},
			wantErr: false,
		},
		{
			name: "success without pointer string",
			args: args{
				in: &TestStruct{
					TestString:        "foo",
					TestPointerString: nil,
					TestInt:           10,
					TestPointerInt:    pointer.ToInt(15),
				},
				tag: "form",
			},
			want: map[string]string{
				"string":     "foo",
				"int":        "10",
				"pointerInt": "15",
			},
			wantErr: false,
		},
		{
			name: "success without int",
			args: args{
				in: &TestStruct{
					TestString:        "foo",
					TestPointerString: pointer.ToString("bar"),
					TestPointerInt:    pointer.ToInt(15),
				},
				tag: "form",
			},
			want: map[string]string{
				"string":        "foo",
				"pointerString": "bar",
				"int":           "0",
				"pointerInt":    "15",
			},
			wantErr: false,
		},
		{
			name: "success without pointer int",
			args: args{
				in: &TestStruct{
					TestString:        "foo",
					TestPointerString: pointer.ToString("bar"),
					TestInt:           10,
					TestPointerInt:    nil,
				},
				tag: "form",
			},
			want: map[string]string{
				"string":        "foo",
				"pointerString": "bar",
				"int":           "10",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToMapString(tt.args.in, tt.args.tag)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToMapString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMapString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
