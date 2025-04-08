package main

import (
	"context"
	"testing"

	calculatorv1 "ConnectRPC-Demo/gen/calculator/v1"

	"connectrpc.com/connect"
)

func TestCalculatorServer_Subtract(t *testing.T) {
	server := &calculatorServer{}
	ctx := context.Background()

	tests := []struct {
		name    string
		number1 int32
		number2 int32
		want    int32
	}{
		{
			name:    "Positive numbers",
			number1: 10,
			number2: 5,
			want:    5,
		},
		{
			name:    "Negative result",
			number1: 5,
			number2: 10,
			want:    -5,
		},
		{
			name:    "Zero result",
			number1: 5,
			number2: 5,
			want:    0,
		},
		{
			name:    "With negative numbers",
			number1: -10,
			number2: -5,
			want:    -5,
		},
		{
			name:    "Large numbers",
			number1: 1000000,
			number2: 999999,
			want:    1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := connect.NewRequest(&calculatorv1.OperateRequest{
				Number1: float64(tt.number1),
				Number2: float64(tt.number2),
			})

			resp, err := server.Subtract(ctx, req)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if int32(resp.Msg.Result) != tt.want {
				t.Errorf("Subtract() = %v, want %v", resp.Msg.Result, tt.want)
			}
		})
	}
}
