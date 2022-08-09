package testing

import "testing"

func TestAbs(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "1 abs",
			args: args{
				x: 1,
			},
			want: 1,
		},
		{
			name: "-1 abs",
			args: args{
				x: -1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.args.x); got != tt.want {
				t.Errorf("Abs( %v) = %v, want %v", tt.args.x, got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "2 > 1",
			args: args{x: 2, y: 1},
			want: 2,
		},
		{
			name: "3 > -1",
			args: args{x: 3, y: 2},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Max(%v,%v) = %v, want %v", tt.args.x, tt.args.y, got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "1 < 2",
			args: args{x: 2, y: 1},
			want: 1,
		},
		{
			name: "-1 < 3",
			args: args{x: 3, y: -1},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Max(%v,%v) = %v, want %v", tt.args.x, tt.args.y, got, tt.want)
			}
		})
	}
}

func TestRandInt(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandInt(); got != tt.want {
				t.Errorf("RandInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
