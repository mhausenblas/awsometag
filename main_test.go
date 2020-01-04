package main

import "testing"

func TestTypeGuess(t *testing.T) {
	res2types := []struct {
		res   string
		rtype string
	}{
		{"arn:aws:s3:::*/*", "s3"},
		{"arn:aws:s3:::abucket/*", "s3"},
		{"arn:aws:s3:::abucket/thing/*", "s3"},
		{"arn:aws:iam::123456789012:user/*", "iam"},
		{"arn:aws:iam::123456789012:user/abc", "iam"},
		{"arn:aws:iam::123456789012:user/dev/some/*", "iam"},
	}

	for _, res2type := range res2types {
		rtype, _ := guesstype(res2type.res)
		if rtype != res2type.rtype {
			t.Errorf("Expected resource type %s but got %s", res2type.rtype, rtype)
		}
	}
}
