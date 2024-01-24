package group_anagrams_49

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Case1",
			args: args{strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			want: [][]string{
				{"eat", "tea", "ate"},
				{"tan", "nat"},
				{"bat"},
			},
		},
		{
			name: "Case2",
			args: args{strs: []string{""}},
			want: [][]string{
				{""},
			},
		},
		{
			name: "Case3",
			args: args{strs: []string{"a"}},
			want: [][]string{
				{"a"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
