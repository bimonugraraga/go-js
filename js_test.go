package gojs

import (
	"errors"
	"reflect"
	"testing"
)

func Test_Ternary(t *testing.T) {
	type args struct {
		condition bool
		result1   interface{}
		result2   interface{}
	}
	testingString := "testing"
	testingInt := 100
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "Success_one_condition_bool",
			args: args{
				condition: true,
				result1:   "Success",
				result2:   "Fail",
			},
			want: "Success",
		},
		{
			name: "Success_one_condition_string",
			args: args{
				condition: testingString == "testing",
				result1:   "Success",
				result2:   "Fail",
			},
			want: "Success",
		},
		{
			name: "Success_one_condition_int",
			args: args{
				condition: testingInt == 100,
				result1:   "Success",
				result2:   "Fail",
			},
			want: "Success",
		},
		{
			name: "Success_two_condition_bool_string",
			args: args{
				condition: true && testingString == "testing",
				result1:   "Success",
				result2:   "Fail",
			},
			want: "Success",
		},
		{
			name: "Success_two_condition_bool_int",
			args: args{
				condition: true && testingInt == 100,
				result1:   "Success",
				result2:   "Fail",
			},
			want: "Success",
		},
		{
			name: "Success_two_condition_int_string",
			args: args{
				condition: testingInt == 100 && testingString == "testing",
				result1:   "Success",
				result2:   "Fail",
			},
			want: "Success",
		},
		{
			name: "Success_output_int",
			args: args{
				condition: true,
				result1:   100,
				result2:   "Fail",
			},
			want: 100,
		},
		{
			name: "Success_output_interface",
			args: args{
				condition: true,
				result1:   interface{}("Success"),
				result2:   "Fail",
			},
			want: interface{}("Success"),
		},
		{
			name: "Success_output_map",
			args: args{
				condition: true && testingInt == 100,
				result1:   map[int]string{1: "Success"},
				result2:   "Fail",
			},
			want: map[int]string{1: "Success"},
		},
		{
			name: "Fail_output_int",
			args: args{
				condition: false,
				result1:   100,
				result2:   0,
			},
			want: 0,
		},
		{
			name: "Fail_output_interface",
			args: args{
				condition: false,
				result1:   interface{}("Success"),
				result2:   interface{}("Fail"),
			},
			want: interface{}("Fail"),
		},
		{
			name: "Fail_output_map",
			args: args{
				condition: false,
				result1:   map[int]string{1: "Success"},
				result2:   map[int]string{1: "Fail"},
			},
			want: map[int]string{1: "Fail"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ternary(tt.args.condition, tt.args.result1, tt.args.result2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ternary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_IsFalsy(t *testing.T) {
	type args struct {
		params interface{}
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Falsy_input_empty_string",
			args: args{
				params: "",
			},
			want: true,
		},
		{
			name: "Falsy_input_nil",
			args: args{
				params: nil,
			},
			want: true,
		},
		{
			name: "Falsy_input_0",
			args: args{
				params: 0,
			},
			want: true,
		},
		{
			name: "Not_Falsy_input_string",
			args: args{
				params: "testing",
			},
			want: false,
		},
		{
			name: "Not_Falsy_input_map",
			args: args{
				params: make(map[int]string),
			},
			want: false,
		},
		{
			name: "Not_Falsy_input_int",
			args: args{
				params: 100,
			},
			want: false,
		},
		{
			name: "Falsy_input_error",
			args: args{
				params: errors.New("Err"),
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFalsy(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ternary() = %v, want %v", got, tt.want)
			}
		})
	}
}
