package regexp_trans

import (
	"reflect"
	"testing"
)

func TestMergeCharRangeArray(t *testing.T) {
	type args struct {
		a CharRangeArray
		b CharRangeArray
	}
	tests := []struct {
		name string
		args args
		want CharRangeArray
	}{
		{
			name: "0. s1",
			args: args{
				a: CharRangeArray{
					[2]rune{'a', 'b'},
				},
				b: CharRangeArray{},
			},
			want: CharRangeArray{
				[2]rune{'a', 'b'},
			},
		},
		{
			name: "0. s2",
			args: args{
				a: CharRangeArray{},
				b: CharRangeArray{
					[2]rune{'a', 'b'},
				},
			},
			want: CharRangeArray{
				[2]rune{'a', 'b'},
			},
		},
		{
			name: "1. same a and b",
			args: args{
				a: CharRangeArray{
					[2]rune{'a', 'b'},
				},
				b: CharRangeArray{
					[2]rune{'a', 'b'},
				},
			},
			want: CharRangeArray{
				[2]rune{'a', 'b'},
			},
		},
		{
			name: "1. more than one element of a and b",
			args: args{
				a: CharRangeArray{
					[2]rune{'a', 'b'},
					[2]rune{'e', 'f'},
				},
				b: CharRangeArray{
					[2]rune{'a', 'b'},
					[2]rune{'e', 'f'},
				},
			},
			want: CharRangeArray{
				[2]rune{'a', 'b'},
				[2]rune{'e', 'f'},
			},
		},
		{
			name: "2.1. more than one element of a and b",
			args: args{
				a: CharRangeArray{
					[2]rune{'a', 'b'},
				},
				b: CharRangeArray{
					[2]rune{'a', 'b'},
					[2]rune{'e', 'f'},
				},
			},
			want: CharRangeArray{
				[2]rune{'a', 'b'},
				[2]rune{'e', 'f'},
			},
		},
		{
			name: "2.2. more than one element of a and b",
			args: args{
				a: CharRangeArray{
					[2]rune{'a', 'b'},
					[2]rune{'e', 'f'},
				},
				b: CharRangeArray{
					[2]rune{'a', 'b'},
				},
			},
			want: CharRangeArray{
				[2]rune{'a', 'b'},
				[2]rune{'e', 'f'},
			},
		},
		{
			name: "3.1. more than one element of a and b",
			args: args{
				a: CharRangeArray{
					[2]rune{'e', 'f'},
				},
				b: CharRangeArray{
					[2]rune{'a', 'b'},
				},
			},
			want: CharRangeArray{
				[2]rune{'a', 'b'},
				[2]rune{'e', 'f'},
			},
		},
		{
			name: "3.2. more than one element of a and b",
			args: args{
				a: CharRangeArray{
					[2]rune{'a', 'b'},
				},
				b: CharRangeArray{

					[2]rune{'e', 'f'},
				},
			},
			want: CharRangeArray{
				[2]rune{'a', 'b'},
				[2]rune{'e', 'f'},
			},
		},
		{
			name: "4.1. one in other one(not same start and end, is contain)",
			args: args{
				a: CharRangeArray{
					[2]rune{'a', 'f'},
				},
				b: CharRangeArray{

					[2]rune{'b', 'c'},
				},
			},
			want: CharRangeArray{
				[2]rune{'a', 'f'},
			},
		},
		{
			name: "4.2. one in other one(not same start and end, is contain)",
			args: args{
				a: CharRangeArray{
					[2]rune{'b', 'c'},
				},
				b: CharRangeArray{
					[2]rune{'a', 'f'},
				},
			},
			want: CharRangeArray{
				[2]rune{'a', 'f'},
			},
		},
		{
			name: "4.3. one in other one(same start and different end)",
			args: args{
				a: CharRangeArray{
					[2]rune{'a', 'c'},
				},
				b: CharRangeArray{
					[2]rune{'a', 'f'},
				},
			},
			want: CharRangeArray{
				[2]rune{'a', 'f'},
			},
		},
		{
			name: "4.4. one in other one(same start and different end)",
			args: args{
				a: CharRangeArray{
					[2]rune{'a', 'f'},
				},
				b: CharRangeArray{
					[2]rune{'a', 'c'},
				},
			},
			want: CharRangeArray{
				[2]rune{'a', 'f'},
			},
		},
		{
			name: "4.5. one in other one(different start and same end)",
			args: args{
				a: CharRangeArray{
					[2]rune{'c', 'f'},
				},
				b: CharRangeArray{
					[2]rune{'a', 'f'},
				},
			},
			want: CharRangeArray{
				[2]rune{'a', 'f'},
			},
		},
		{
			name: "4.6. one in other one(different start and same end)",
			args: args{
				a: CharRangeArray{
					[2]rune{'a', 'f'},
				},
				b: CharRangeArray{
					[2]rune{'c', 'f'},
				},
			},
			want: CharRangeArray{
				[2]rune{'a', 'f'},
			},
		},
		{
			name: "4.7. one in other one(more than one element)",
			args: args{
				a: CharRangeArray{
					[2]rune{'a', 'f'},
				},
				b: CharRangeArray{
					[2]rune{'c', 'd'},
				},
			},
			want: CharRangeArray{
				[2]rune{'a', 'f'},
			},
		},
		{
			name: "4.8. one in other one(more than one element)",
			args: args{
				a: CharRangeArray{
					[2]rune{'c', 'd'},
				},
				b: CharRangeArray{
					[2]rune{'a', 'f'},
				},
			},
			want: CharRangeArray{
				[2]rune{'a', 'f'},
			},
		},
		{
			name: "5.1. some elements of one in other one(more than one element)",
			args: args{
				a: CharRangeArray{
					[2]rune{'f', 'g'},
				},
				b: CharRangeArray{
					[2]rune{'a', 'f'},
				},
			},
			want: CharRangeArray{
				[2]rune{'a', 'g'},
			},
		},
		{
			name: "5.2. some elements of one in other one(more than one element)",
			args: args{
				a: CharRangeArray{
					[2]rune{'a', 'f'},
				},
				b: CharRangeArray{
					[2]rune{'f', 'g'},
				},
			},
			want: CharRangeArray{
				[2]rune{'a', 'g'},
			},
		},
		{
			name: "6.1. big case 1",
			args: args{
				a: CharRangeArray{
					[2]rune{'a', 'b'},
				},
				b: CharRangeArray{
					[2]rune{'f', 'g'},
				},
			},
			want: CharRangeArray{
				[2]rune{'a', 'b'},
				[2]rune{'f', 'g'},
			},
		},
		{
			name: "6.2. big case 1",
			args: args{
				a: CharRangeArray{
					[2]rune{'f', 'g'},
				},
				b: CharRangeArray{
					[2]rune{'a', 'b'},
				},
			},
			want: CharRangeArray{
				[2]rune{'a', 'b'},
				[2]rune{'f', 'g'},
			},
		},
		{
			name: "6.3. big case 2",
			args: args{
				a: CharRangeArray{
					[2]rune{'a', 'b'},
					[2]rune{'f', 'g'},
				},
				b: CharRangeArray{
					[2]rune{'a', 'c'},
				},
			},
			want: CharRangeArray{
				[2]rune{'a', 'c'},
				[2]rune{'f', 'g'},
			},
		},
		{
			name: "6.4. big case 2",
			args: args{
				a: CharRangeArray{
					[2]rune{'a', 'c'},
				},
				b: CharRangeArray{
					[2]rune{'a', 'b'},
					[2]rune{'f', 'g'},
				},
			},
			want: CharRangeArray{
				[2]rune{'a', 'c'},
				[2]rune{'f', 'g'},
			},
		},
		{
			name: "6.5. big case 3",
			args: args{
				a: CharRangeArray{
					[2]rune{'a', 'b'},
					[2]rune{'d', 'e'},
				},
				b: CharRangeArray{
					[2]rune{'b', 'c'},
					[2]rune{'f', 'h'},
				},
			},
			want: CharRangeArray{
				[2]rune{'a', 'h'},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeCharRangeArray(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeCharRangeArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
