package protometry

import (
    "github.com/louis030195/protometry/internal/utils"
    "reflect"
    "testing"
)


func TestBox_Fit(t *testing.T) {
	a := NewBoxOfSize(0.5, 0.5, 0.5, 1)
	b := NewBoxOfSize(0.5, 0.5, 0.5, 1)

	// contains equal Box, symmetrically
	utils.Equals(t, true, a.Fit(*b))

	utils.Equals(t, true, b.Fit(*b))

	// contained on edge
	b = NewBoxMinMax(0, 0, 0, 0.5, 1, 1)

	utils.Equals(t, true, b.Fit(*a))

	utils.Equals(t, false, a.Fit(*b))

	// contained away from edges
	b = NewBoxMinMax(0.1, 0.1, 0.1, 0.9, 0.9, 0.9)
	utils.Equals(t, true, b.Fit(*a))

	utils.Equals(t, false, a.Fit(*b))

	// 1 corner Fit
	b = NewBoxMinMax(-0.1, -0.1, -0.1, 0.9, 0.9, 0.9)

	utils.Equals(t, false, b.Fit(*a))

	utils.Equals(t, false, a.Fit(*b))

	b = NewBoxMinMax(0.9, 0.9, 0.9, 1.1, 1.1, 1.1)

	utils.Equals(t, false, b.Fit(*a))

	utils.Equals(t, false, a.Fit(*b))
}

func TestBox_Intersects(t *testing.T) {
	a := NewBoxMinMax(0, 0, 0, 1, 1, 1)

	b := NewBoxMinMax(1.1, 0, 0, 2, 1, 1)


	// not intersecting area above or below in each dimension
	utils.Equals(t, false, a.Intersects(*b))

	b = NewBoxMinMax(-1, 0, 0, -0.1, 1, 1)

	utils.Equals(t, false, a.Intersects(*b))

	b = NewBoxMinMax(0, 1.1, 0, 1, 2, 1)

	utils.Equals(t, false, a.Intersects(*b))

	b = NewBoxMinMax(0, -1, 0, 1, -0.1, 1)

	utils.Equals(t, false, a.Intersects(*b))

	b = NewBoxMinMax(0, 0, 1.1, 1, 1, 2)

	utils.Equals(t, false, a.Intersects(*b))

	b = NewBoxMinMax(0, 0, -1, 1, 1, -0.1)

	utils.Equals(t, false, a.Intersects(*b))

	// intersects equal Box, symmetrically
	b = NewBoxMinMax(0, 0, 0, 1, 1, 1)

	utils.Equals(t, true, a.Intersects(*b))

	// intersects containing and contained
	b = NewBoxMinMax(0.1, 0.1, 0.1, 0.9, 0.9, 0.9)

	utils.Equals(t, true, a.Intersects(*b))

	// intersects partial containment on each corner
	b = NewBoxMinMax(0.9, 0.9, 0.9, 2, 2, 2)

	utils.Equals(t, true, a.Intersects(*b))

	b = NewBoxMinMax(-1, 0.9, 0.9, 1, 2, 2)

	utils.Equals(t, true, a.Intersects(*b))

	b = NewBoxMinMax(0.9, -1, 0.9, 2, 0.1, 2)

	utils.Equals(t, true, a.Intersects(*b))

	b = NewBoxMinMax(-1, -1, 0.9, 0.1, 0.1, 2)

	utils.Equals(t, true, a.Intersects(*b))

	b = NewBoxMinMax(0.9, 0.9, -1, 2, 2, 0.1)

	utils.Equals(t, true, a.Intersects(*b))

	b = NewBoxMinMax(-1, 0.9, -1, 0.1, 2, 0.1)

	utils.Equals(t, true, a.Intersects(*b))

	b = NewBoxMinMax(0.9, -1, -1, 2, 0.1, 0.1)

	utils.Equals(t, true, a.Intersects(*b))

	b = NewBoxMinMax(-1, -1, -1, 0.1, 0.1, 0.1)

	utils.Equals(t, true, a.Intersects(*b))

	// intersects 'beam'; where no corners Fit
	// other but some contained
	b = NewBoxMinMax(-1, 0.1, 0.1, 2, 0.9, 0.9)

	utils.Equals(t, true, a.Intersects(*b))

	b = NewBoxMinMax(0.1, -1, 0.1, 0.9, 2, 0.9)

	utils.Equals(t, true, a.Intersects(*b))

	b = NewBoxMinMax(0.1, 0.1, -1, 0.9, 0.9, 2)

	utils.Equals(t, true, a.Intersects(*b))

	// Other
	b = NewBoxMinMax(1, 1, 1, 4, 4, 4)

	b = NewBoxMinMax(0, 0, 0, 1, 1, 1)

	utils.Equals(t, true, b.Intersects(*a))
	utils.Equals(t, true, a.Intersects(*b))
	b = NewBoxMinMax(1, 1, 1, 1, 1, 1)

	utils.Equals(t, true, b.Intersects(*a))
	b = NewBoxMinMax(1, 1, 1, 4, 4, 4)

	utils.Equals(t, true, b.Intersects(*a))
}

func TestBox_Split(t *testing.T) {
	/*
		What we want to achieve (in 2D):
		 _ _ _ _
		|		|
		|_ _ _ _|

		->
		 _ _ _ _
		|_ _|_ _|
		|_ _|_ _|

	 */

	/*
		Representation (ignoring the Z axis cause it's hard to draw in 3D ;))
			centered at 0.5,0.5,0
			0.5,0.5,0 extent (square)
			--->
		 _ _ _ _
		|		|
		|_ _ _ _|

		------->
		1 size
		min 0,0,0
		max 1,1,0

		So in theory the split-ed box would be

		Each sub-box extent: 0.25,0.25,0
		 ->
		 _ _ _ _
		|_A_|_B_|
		|_D_|_C_|

		--->
		Each sub-box size: 0.5,0.5,0

		- - - - - - - - - - - - - - - - - - - - - - - - - - |
		B:						|		D:					|
		center: 0.25,0.75,0		|		center: 0.75,0.75,0 |
		min: 0,0.5,0			|		min: 0.5,0.5,0		|
		max: 0.5,1,0			|		max: 1,1,0			|
		- - - - - - - - - - - - - - - - - - - - - - - - - - |
		A:						|		C:					|
		center: 0.25,0.25,0		|		center: 0.75,0.25,0 |
		min: 0,0,0				|		min: 0.5,0,0		|
		max: 0.5,0.5,0			|		max: 1,0.5,0		|
		- - - - - - - - - - - - - - - - - - - - - - - - - - |
		...
	 */
	b := Box{
		Min: NewVector3(0, 0, 0),
		Max: NewVector3(1, 1, 0),
	}
	got := b.Split()


	/*
	 *    3____7
	 *  2/___6/|
	 *  | 1__|_5
	 *  0/___4/
	*/
	want := [8]*Box{
		{Min: NewVector3(0, 0, 0), Max: NewVector3(0.5, 0.5, 0)}, // A
		{Min: NewVector3(0, 0, 0), Max: NewVector3(0.5, 0.5, 0)}, // A
		{Min: NewVector3(0, 0.5, 0), Max: NewVector3(0.5, 1, 0)}, // B
		{Min: NewVector3(0, 0.5, 0), Max: NewVector3(0.5, 1, 0)}, // B

		{Min: NewVector3(0.5, 0, 0), Max: NewVector3(1, 0.5, 0)}, // C
		{Min: NewVector3(0.5, 0, 0), Max: NewVector3(1, 0.5, 0)}, // C
		{Min: NewVector3(0.5, 0.5, 0), Max: NewVector3(1, 1, 0)}, // D
		{Min: NewVector3(0.5, 0.5, 0), Max: NewVector3(1, 1, 0)}, // D
	}

	tester := func(got, want [8]*Box) {
		t.Logf("\nBefore split \n%v", b)
		for i := range want {
			t.Logf("\nMin: {%v}\nWant: {%v}\n\nMax: {%v}\nWant: {%v}",
				got[i].Min,
				want[i].Min,
				got[i].Max,
				want[i].Max,
			)
			utils.Equals(t, want[i].Min, got[i].Min)
			utils.Equals(t, want[i].Max, got[i].Max)
		}
	}
	tester(got, want)
}

func BenchmarkArea_NewBoxMinMax(b *testing.B) {
	size := float64(b.N)
	b.ResetTimer()
	for i := 0.; i < size; i++ {
		NewBoxMinMax(i, i, i, i*2, i*2, i*2)
	}
}

func BenchmarkArea_NewBoxOfSize(b *testing.B) {
	size := float64(b.N)
	b.ResetTimer()
	for i := 0.; i < size; i++ {
		NewBoxOfSize(i, i, i, i)
	}
}

func BenchmarkArea_In(b *testing.B) {
	size := float64(b.N)
	b.ResetTimer()
	for i := 0.; i < size; i++ {
		NewVector3(i, i, i).In(*NewBoxOfSize(i, i, i, 1))
	}
}

func BenchmarkArea_Fit(b *testing.B) {
	size := float64(b.N)
	b.ResetTimer()
	for i := 0.; i < size; i++ {
		NewBoxOfSize(i, i, i, 0.1).Fit(*NewBoxOfSize(i, i, i, 1))
	}
}

func BenchmarkArea_Intersects(b *testing.B) {
	size := float64(b.N)
	b.ResetTimer()
	for i := 0.; i < size; i++ {
		NewBoxOfSize(i, i, i, 1).Intersects(*NewBoxOfSize(i, i, i, 1))
	}
}

func BenchmarkArea_Split(b *testing.B) {
	size := float64(b.N)
	b.ResetTimer()
	for i := 0.; i < size; i++ {
		NewBoxOfSize(i, i, i, 1).Split()
	}
}





























/** Generated tests **/
func TestBox_Equal(t *testing.T) {
	type fields struct {
		Min                  *Vector3
		Max                  *Vector3
	}
	type args struct {
		other Box
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			fields:fields{
				Min: NewVector3(0, 0, 0),
				Max: NewVector3(1, 1, 1),
			},
			args:args{other: *NewBoxMinMax(0, 0, 0, 1, 1, 1)},
			want: true,
		},
		{
			fields:fields{
				Min: NewVector3(0, 0, 0),
				Max: NewVector3(1.1, 1, 1),
			},
			args:args{other: *NewBoxMinMax(0, 0, 0, 1, 1, 1)},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Box{
				Min:                  tt.fields.Min,
				Max:                  tt.fields.Max,
			}
			if got := b.Equal(tt.args.other); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBox_Fit1(t *testing.T) {
	type fields struct {
		Min                  *Vector3
		Max                  *Vector3
	}
	type args struct {
		o Box
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			fields:fields{
				Min: NewVector3(0.5, 0.5, 0.5),
				Max: NewVector3(0.6, 0.6, 0.6),
			},
			args:args{o: *NewBoxMinMax(0.5, 0.5, 0.5, 1, 1, 1)},
			want: true,
		},
		{
			fields:fields{
				Min: NewVector3(5.85, -3.9, 5.2),
				Max: NewVector3(6.85, -2.9, 6.2),
			},
			args:args{o: *NewBoxMinMax(2.92, -1.9, 2.59, 10.3, -4.4, 9.3)},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Box{
				Min:                  tt.fields.Min,
				Max:                  tt.fields.Max,
			}
			if got := b.Fit(tt.args.o); got != tt.want {
				t.Errorf("Fit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBox_GetCenter(t *testing.T) {
	type fields struct {
		Min                  *Vector3
		Max                  *Vector3
	}
	tests := []struct {
		name   string
		fields fields
		want   Vector3
	}{
		// TODO: Add test cases.
		{
			fields:fields{
				Min: NewVector3(0.5, 0.5, 0.5),
				Max: NewVector3(0.6, 0.6, 0.6),
			},
			want: *NewVector3(0.55, 0.55, 0.55),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Box{
				Min:                  tt.fields.Min,
				Max:                  tt.fields.Max,
			}
			if got := b.GetCenter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCenter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBox_GetSize(t *testing.T) {
	type fields struct {
		Min                  *Vector3
		Max                  *Vector3
	}
	tests := []struct {
		name   string
		fields fields
		want   Vector3
	}{
		// TODO: Add test cases.
		{
			fields:fields{
				Min: NewVector3(0.5, 0.5, 0.5),
				Max: NewVector3(0.6, 0.6, 0.6),
			},
			want: *NewVector3(0.1, 0.1, 0.1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Box{
				Min:                  tt.fields.Min,
				Max:                  tt.fields.Max,
			}
			if got := b.GetSize(); !got.Equal(tt.want) {
				t.Errorf("GetSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBox_Intersects1(t *testing.T) {
	type fields struct {
		Min                  *Vector3
		Max                  *Vector3
	}
	type args struct {
		b2 Box
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			fields:fields{
				Min: NewVector3(0.5, 0.5, 0.5),
				Max: NewVector3(0.6, 0.6, 0.6),
			},
			args:args{b2: *NewBoxMinMax(0.5, 0.5, 0.5, 1, 1, 1)},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Box{
				Min:                  tt.fields.Min,
				Max:                  tt.fields.Max,
			}
			if got := b.Intersects(tt.args.b2); got != tt.want {
				t.Errorf("Intersects() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func TestBox_Split1(t *testing.T) {
//	type fields struct {
//		Min                  *Vector3
//		Max                  *Vector3
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   [8]*Box
//	}{
//		// TODO: Add test cases.
//		{
//			fields:fields{
//				Min:                  NewVector3(0, 0, 0),
//				Max:                  NewVector3(1, 1, 1),
//			},
//			want: [8]*Box{
//				// TODO
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			b := &Box{
//				Min:                  tt.fields.Min,
//				Max:                  tt.fields.Max,
//			}
//			if got := b.Split(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Split() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestNewBoxMinMax(t *testing.T) {
	type args struct {
		minX float64
		minY float64
		minZ float64
		maxX float64
		maxY float64
		maxZ float64
	}
	tests := []struct {
		name string
		args args
		want *Box
	}{
		{
			args:args{0.5, 0.5, 0.5, 1, 1, 1},
			want: &Box{
				Min: NewVector3(0.5, 0.5, 0.5),
				Max: NewVector3(1, 1, 1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBoxMinMax(tt.args.minX, tt.args.minY, tt.args.minZ, tt.args.maxX, tt.args.maxY, tt.args.maxZ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBoxMinMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBoxOfSize(t *testing.T) {
	type args struct {
		x    float64
		y    float64
		z    float64
		size float64
	}
	tests := []struct {
		name string
		args args
		want *Box
	}{
		// TODO: Add test cases.
		{
			args: args{0, 0, 0, 1},
			want: NewBoxMinMax(-0.5, -0.5, -0.5, 0.5, 0.5, 0.5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBoxOfSize(tt.args.x, tt.args.y, tt.args.z, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBoxOfSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBox_EncapsulateBox(t *testing.T) {
    type fields struct {
        Min                  *Vector3
        Max                  *Vector3
        XXX_NoUnkeyedLiteral struct{}
        XXX_unrecognized     []byte
        XXX_sizecache        int32
    }
    type args struct {
        o Box
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Box
    }{
        // TODO: Add test cases.
        {
            fields: fields{
                Min: &Vector3{
                    X:                    0,
                    Y:                    0,
                    Z:                    0,
                },
                Max: &Vector3{
                    X:                    1,
                    Y:                    1,
                    Z:                    1,
                },
            },
            args:args{
                o: Box{
                Min: &Vector3{
                    X:                    1,
                    Y:                    1,
                    Z:                    1,
                },
                Max: &Vector3{
                    X:                    2,
                    Y:                    2,
                    Z:                    2,
                },
            },
            },
            want: NewBoxMinMax(0, 0, 0, 2, 2, 2),
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            b := &Box{
                Min:                  tt.fields.Min,
                Max:                  tt.fields.Max,
                XXX_NoUnkeyedLiteral: tt.fields.XXX_NoUnkeyedLiteral,
                XXX_unrecognized:     tt.fields.XXX_unrecognized,
                XXX_sizecache:        tt.fields.XXX_sizecache,
            }
            if got := b.EncapsulateBox(tt.args.o); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("EncapsulateBox() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestBox_EncapsulatePoint(t *testing.T) {
    type fields struct {
        Min                  *Vector3
        Max                  *Vector3
        XXX_NoUnkeyedLiteral struct{}
        XXX_unrecognized     []byte
        XXX_sizecache        int32
    }
    type args struct {
        o Vector3
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Box
    }{
        // TODO: Add test cases.
        {
            fields: fields{
                Min: &Vector3{
                    X:                    0,
                    Y:                    0,
                    Z:                    0,
                },
                Max: &Vector3{
                    X:                    1,
                    Y:                    1,
                    Z:                    1,
                },
            },
            args:args{
                o: Vector3{
                        X:                    2,
                        Y:                    2,
                        Z:                    2,
                    },
            },
            want: NewBoxMinMax(0, 0, 0, 2, 2, 2),
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            b := &Box{
                Min:                  tt.fields.Min,
                Max:                  tt.fields.Max,
                XXX_NoUnkeyedLiteral: tt.fields.XXX_NoUnkeyedLiteral,
                XXX_unrecognized:     tt.fields.XXX_unrecognized,
                XXX_sizecache:        tt.fields.XXX_sizecache,
            }
            if got := b.EncapsulatePoint(tt.args.o); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("EncapsulatePoint() = %v, want %v", got, tt.want)
            }
        })
    }
}