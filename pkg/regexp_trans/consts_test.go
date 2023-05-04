package regexp_trans

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
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

func TestRandomRangeChar(t *testing.T) {
	defaultRan := rand.New(rand.NewSource(time.Now().UnixNano()))
	type args struct {
		ran         *rand.Rand
		charClasses CharRangeArray
		count       int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "1. common", args: args{ran: defaultRan, charClasses: CharClassRangeAllLetters, count: 100}},
		{name: "2. numbers", args: args{ran: defaultRan, charClasses: CharClassRangeNumbers, count: 100}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomRangeChar(defaultRan, tt.args.charClasses, tt.args.count)
			if len(got) != tt.args.count {
				t.Errorf("want count: %d, but: %d", tt.args.count, len(got))
			}
			for _, item := range got {
				if !isRuneInSlice(item, tt.args.charClasses) {
					t.Errorf("got [%v] not in slice: [%v]", item, tt.args.charClasses)
				}
			}
		})
	}
}

func isRuneInSlice(r rune, sli CharRangeArray) bool {
	for _, item := range sli {
		if r >= item[0] && r <= item[1] {
			return true
		}
	}

	return false
}

func TestParseCharRangeArray(t *testing.T) {
	type args struct {
		rs []rune
	}
	tests := []struct {
		name    string
		args    args
		want    CharRangeArray
		wantErr bool
	}{
		{name: "1.common test", args: args{rs: []rune{'a', 'b'}}, want: CharRangeArray{[2]rune{'a', 'b'}}, wantErr: false},
		{name: "2.common error", args: args{rs: []rune{'a'}}, want: nil, wantErr: true},
		{name: "1.common test 2", args: args{rs: []rune{'a', 'b', 'c', 'd'}}, want: CharRangeArray{[2]rune{'a', 'b'}, [2]rune{'c', 'd'}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseCharRangeArray(tt.args.rs)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseCharRangeArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseCharRangeArray() got = %v, want %v", got, tt.want)
			}
		})
	}
}
