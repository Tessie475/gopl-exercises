// editor.go opens the user's preferred text editor so they can type an issue's
// title and body, rather than passing them on the command line.
package main

import (
	"os"
	"os/exec"
	"strings"
)

// editText writes initial into a temp file, opens the editor on it, waits for
// the editor to exit, and returns the edited contents. The editor is taken
// from $EDITOR, then $VISUAL, falling back to vi.
func editText(initial string) (string, error) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = os.Getenv("VISUAL")
	}
	if editor == "" {
		editor = "vi"
	}

	// A temp file the editor will open. os.Remove runs on every return path.
	f, err := os.CreateTemp("", "issue-*.md")
	if err != nil {
		return "", err
	}
	name := f.Name()
	defer os.Remove(name)

	if _, err := f.WriteString(initial); err != nil {
		f.Close()
		return "", err
	}
	if err := f.Close(); err != nil {
		return "", err
	}

	// Give the editor our terminal so the user can interact with it, then
	// block until they save and quit.
	cmd := exec.Command(editor, name)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return "", err
	}

	data, err := os.ReadFile(name)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// splitTitleBody reads edited text as an issue: the first line is the title,
// everything after it is the body.
func splitTitleBody(text string) (title, body string) {
	text = strings.TrimLeft(text, "\n")
	if i := strings.IndexByte(text, '\n'); i >= 0 {
		return strings.TrimSpace(text[:i]), strings.TrimSpace(text[i+1:])
	}
	return strings.TrimSpace(text), ""
}
