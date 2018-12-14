package main

import "testing"

func TestParse(t *testing.T) {
	cases := [] struct {
		in string
		want Message
	}{
		{"<#C3MSN3ANA|general> hello world", Message{"C3MSN3ANA", "hello world"}},
		{"<#C3MSN3ANA|general> hello", Message{"C3MSN3ANA", "hello"}},
		//{"should break", Message{"C3MSN3ANA", "hello"}}, //TODO throw some exception alike
	}
	for _, c := range cases {
		got := parse(c.in)
		if got != c.want {
			t.Errorf("parse(%q) == %q, want %q", c.in, got, c.want)
		}
	}

}
