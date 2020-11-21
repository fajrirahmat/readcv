package search

import (
	"testing"
)

func TestIsFoundGolang(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"Test text contains golang", args{"I have experience in golang"}, true, false},
		{"Test text not contains golang", args{"I have experience in c#"}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsFoundGolang(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsFoundGolang() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsFoundGolang() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFoundKeyword(t *testing.T) {
	type args struct {
		text     string
		keywords []string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"Test text contains golang", args{"I have experience in golang", []string{"golang"}}, true, false},
		{"Test text not contains golang", args{"I have experience in c#", []string{"golang"}}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FoundKeyword(tt.args.text, tt.args.keywords)
			if (err != nil) != tt.wantErr {
				t.Errorf("FoundKeyword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FoundKeyword() = %v, want %v", got, tt.want)
			}
		})
	}
}
