package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// runMain calls main() as if it had been launched with the given arguments,
// and returns everything it printed to stdout.
//
// It never modifies main.go: it just swaps os.Args for the fake one and
// captures os.Stdout through a pipe.
func runMain(t *testing.T, args ...string) string {
	t.Helper()
	oldArgs := os.Args
	os.Args = append([]string{"ascii-art"}, args...)
	defer func() { os.Args = oldArgs }()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer w.Close()

	oldStdout := os.Stdout
	os.Stdout = w
	defer func() { os.Stdout = oldStdout }()

	done := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		done <- buf.String()
	}()

	main()
	w.Close()
	return <-done
}

// loadGolden reads a file from testdata/ as a plain string.
func loadGolden(t *testing.T, name string) string {
	t.Helper()
	content, err := os.ReadFile("testdata/" + name + ".golden")
	if err != nil {
		t.Fatalf("reading golden %q: %v", name, err)
	}
	return string(content)
}

// splitLines splits an output string into lines, ignoring the
// trailing newline that fmt.Println always adds at the end.
func splitLines(out string) []string {
	return strings.Split(strings.TrimSuffix(out, "\n"), "\n")
}

// countBlank counts the empty lines in an output string.
func countBlank(out string) int {
	n := 0
	for _, l := range splitLines(out) {
		if l == "" {
			n++
		}
	}
	return n
}

// An empty argument must print nothing.
func TestEmptyArgument(t *testing.T) {
	if got := runMain(t, ""); got != "" {
		t.Errorf("empty arg: got %q, want %q", got, "")
	}
}

// The literal \n argument must print exactly one blank line.
func TestNewlineOnly(t *testing.T) {
	if got := runMain(t, "\\n"); got != "\n" {
		t.Errorf("backslash-n: got %q, want %q", got, "\n")
	}
}

// The \n edge cases: line counts and blank-line placement.
func TestNewlineCases(t *testing.T) {
	tests := []struct {
		arg        string
		wantLines  int
		wantBlanks int
	}{
		{"Hello\\n", 9, 1},
		{"Hello\\nThere", 16, 0},
		{"Hello\\n\\nThere", 17, 1},
		{"\\nHello", 8, 0},
		{"\\n\\n", 2, 2},
	}
	for _, tt := range tests {
		out := runMain(t, tt.arg)
		lines := splitLines(out)
		if len(lines) != tt.wantLines {
			t.Errorf("runMain(%q): got %d lines, want %d", tt.arg, len(lines), tt.wantLines)
		}
		if blanks := countBlank(out); blanks != tt.wantBlanks {
			t.Errorf("runMain(%q): got %d blank lines, want %d", tt.arg, blanks, tt.wantBlanks)
		}
	}
}

// "Hello" must match the expected output byte for byte.
func TestHelloGolden(t *testing.T) {
	want := loadGolden(t, "hello")
	if got := runMain(t, "Hello"); got != want {
		t.Errorf("Hello mismatch\ngot:  %q\nwant: %q", got, want)
	}
}

// "Hello There" (with a space) must match its expected output.
func TestHelloThereGolden(t *testing.T) {
	want := loadGolden(t, "hellothere")
	if got := runMain(t, "Hello There"); got != want {
		t.Errorf("Hello There mismatch\ngot:  %q\nwant: %q", got, want)
	}
}

// Each banner must render the same text differently, in exactly 8 lines.
func TestBannersDiffer(t *testing.T) {
	standard := runMain(t, "Hello", "standard")
	shadow := runMain(t, "Hello", "shadow")
	thinkertoy := runMain(t, "Hello", "thinkertoy")

	for name, out := range map[string]string{
		"standard":   standard,
		"shadow":     shadow,
		"thinkertoy": thinkertoy,
	} {
		if lines := len(splitLines(out)); lines != 8 {
			t.Errorf("banner %s: got %d lines, want 8", name, lines)
		}
	}
	if standard == shadow || standard == thinkertoy || shadow == thinkertoy {
		t.Error("banners should render differently")
	}
}

// A banner name outside the allowed set must be rejected.
func TestInvalidBanner(t *testing.T) {
	if got := runMain(t, "Hello", "banana"); got != "invalid banner: banana\n" {
		t.Errorf("got %q, want %q", got, "invalid banner: banana\n")
	}
}

// Zero or more than two arguments must be rejected.
func TestInvalidArguments(t *testing.T) {
	want := "Invalid arguments\n"
	if got := runMain(t); got != want {
		t.Errorf("0 args: got %q, want %q", got, want)
	}
	if got := runMain(t, "Hello", "standard", "extra"); got != want {
		t.Errorf("3 args: got %q, want %q", got, want)
	}
}
