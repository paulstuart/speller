package speller

import (
	"reflect"
	"sort"
	"testing"
)

func cmp(t *testing.T, expects, got []string) {
	t.Helper()
	sort.Strings(expects)
	sort.Strings(got)
	if !reflect.DeepEqual(got, expects) {
		t.Fatalf("\nexpected: %v\ngot: %v\n", expects, got)
	}
}
func TestSplits(t *testing.T) {
	expects := splits{
		{"", "hello"},
		{"h", "ello"},
		{"he", "llo"},
		{"hel", "lo"},
		{"hell", "o"},
		{"hello", ""},
	}
	got := cleaves("hello")

	if !reflect.DeepEqual(got, expects) {
		t.Fatalf("expected: %v -- got: %v\n", expects, got)
	}
}

func TestSplits2(t *testing.T) {
	expects := splits{
		{"", "no"},
		{"n", "o"},
		{"no", ""},
	}
	got := cleaves("no")
	if !reflect.DeepEqual(got, expects) {
		t.Fatalf("expected: %v -- got: %v\n", expects, got)
	}
}

func TestDeletes(t *testing.T) {
	expects := []string{"o", "n"}
	got := deletes(cleaves("no"))
	cmp(t, expects, got)
}

func TestTransposes(t *testing.T) {
	expects := []string{"on"}
	got := transposes(cleaves("no"))
	cmp(t, expects, got)
}

func TestReplaces(t *testing.T) {
	expects := []string{
		"ao", "bo", "co", "do", "eo", "fo", "go", "ho", "io", "jo", "ko", "lo", "mo", "no", "oo", "po", "qo", "ro", "so", "to", "uo", "vo", "wo", "xo", "yo", "zo", "na", "nb", "nc", "nd", "ne", "nf", "ng", "nh", "ni", "nj", "nk", "nl", "nm", "nn", "no", "np", "nq", "nr", "ns", "nt", "nu", "nv", "nw", "nx", "ny", "nz",
	}
	got := replaces(cleaves("no"))
	cmp(t, expects, got)
}

func TestEdits1(t *testing.T) {
	expects := []string{"nao", "not", "mno", "go", "ng", "jno", "nlo", "ono", "lo", "ngo", "to", "bno", "nol", "do", "non", "noo", "rno", "noi", "yo", "nok", "nod", "noe", "nof", "nog", "noa", "nob", "noc", "nox", "noy", "noz", "nvo", "nou", "nov", "now", "nop", "noq", "nor", "nos", "yno", "nfo", "qo", "eo", "zo", "njo", "nuo", "lno", "nbo", "neo", "nwo", "cno", "ro", "wo", "bo", "jo", "nto", "nmo", "qno", "oo", "on", "ndo", "o", "fno", "xno", "uno", "co", "xo", "nso", "nho", "dno", "po", "nro", "wno", "nko", "ho", "pno", "nco", "mo", "n", "uo", "tno", "nzo", "gno", "nno", "nom", "ko", "ano", "noj", "ao", "vo", "ino", "io", "noh", "nqo", "nh", "ni", "nj", "nk", "nl", "nm", "nn", "no", "na", "nb", "nc", "nd", "ne", "nf", "eno", "nx", "ny", "nz", "nyo", "np", "nq", "nr", "ns", "nt", "nu", "nv", "nw", "npo", "nio", "vno", "sno", "fo", "kno", "zno", "hno", "so", "nxo"}
	got := edits1("no")
	cmp(t, expects, got)
}
