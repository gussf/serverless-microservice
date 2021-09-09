package models

import "testing"

func TestCar_Available(t *testing.T) {
	tests := []struct {
		name string
		car  Car
		want bool
	}{
		{
			"Should return that car is available",
			Car{available: true, name: "Name", brand: "Brand", price_cents: 1},
			true,
		},
		{
			"Should return that car is not available",
			Car{available: false, name: "Name", brand: "Brand", price_cents: 2},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.car.Available(); got != tt.want {
				t.Errorf("Car.Available() = %v, want %v", got, tt.want)
			}
		})
	}
}
