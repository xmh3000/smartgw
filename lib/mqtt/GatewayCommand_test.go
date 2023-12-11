package mqtt

import (
	"reflect"
	"testing"
)

func Test_getDevice(t *testing.T) {
	type args struct {
		params map[string]interface{}
		client Client
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDevice(tt.args.params, tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDevice() = %v, want %v", got, tt.want)
			}
		})
	}
}
