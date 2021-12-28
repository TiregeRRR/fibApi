package apiGRPC

import (
	"context"
	"errors"
	"testing"
)

var tests = []struct {
	X, Y          int32
	expectedList  []uint64
	expectedError error
}{
	{
		X:             0,
		Y:             5,
		expectedList:  []uint64{0, 1, 1, 2, 3, 5},
		expectedError: nil,
	},
	{
		X:             -1,
		Y:             1,
		expectedError: errors.New("invalid input: x < 0"),
	},
	{
		X:             1,
		Y:             -1,
		expectedError: errors.New("invalid input: y < 0"),
	},
	{
		X:             10,
		Y:             1,
		expectedError: errors.New("invalid input: x > y"),
	},
}

func equal(a, b []uint64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestGRPCWriteFibonacci(t *testing.T) {
	s := GRPCSrv{}
	for _, tc := range tests {
		req := &FibRequest{X: tc.X, Y: tc.Y}
		resp, err := s.GetFib(context.Background(), req)
		if tc.expectedError != nil {
			if err.Error() != tc.expectedError.Error() {
				t.Errorf("GetFib(%v, %v) got %v, expected %d", tc.X, tc.Y, err, tc.expectedError)
			}
		} else if !equal(resp.FibList, tc.expectedList) {
			t.Errorf("GetFib(%v, %v)=%v, wanted %v", tc.X, tc.Y, resp.FibList, tc.expectedList)
		}
	}
}
