package usecase_test

import (
	"testing"

	"github.com/xoltawn/simple-file-storage-file-service/usecase"
)

func TestIsLink(t *testing.T) {
	sut := usecase.NewLinkValidator()
	testCases := []struct {
		desc string
		txt  string
		want bool
	}{
		{
			desc: "starts with http",
			txt:  "http://placehold.co/600x400",
			want: true,
		}, {
			desc: "starts with https",
			txt:  "https://placehold.co/600x400",
			want: true,
		}, {
			desc: "empty url",
			txt:  "",
			want: false,
		}, {
			desc: "space",
			txt:  " ",
			want: false,
		}, {
			desc: "valid url with spaces",
			txt:  "                   https://placehold.co/600x400",
			want: true,
		},
	}

	for _, tc := range testCases {
		if got := sut.IsLink(tc.txt); got != tc.want {
			t.Errorf("IsLink(%q) = %v, want %v", tc.txt, got, tc.want)
		}
	}
}
