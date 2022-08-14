package parser

import (
	"reflect"
	"testing"
)

func Test_splitIntoParts(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"basic", args{"hoge"}, []string{"hoge"}},
		{"basic", args{"hoge fuga"}, []string{"hoge", "fuga"}},
		{"more space", args{"  hoge   fuga "}, []string{"hoge", "fuga"}},
		{"empty", args{""}, []string{}},
		{"comment out", args{"// hogehoge"}, []string{}},
		{"elements with comment out", args{"hoge fuga piyo // hogehoge"}, []string{"hoge", "fuga", "piyo"}},
		{"elements with comment out", args{"hoge fuga piyo// hogehoge"}, []string{"hoge", "fuga", "piyo"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitIntoParts(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitIntoParts() = %v, want %v", got, tt.want)
			}
		})
	}
}
