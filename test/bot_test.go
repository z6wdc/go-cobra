package bot_test

import (
	"bytes"
	bot "github.com/z6wdc/go-bot/internal"
	"io"
	"os"
	"testing"
)

func TestBot(t *testing.T) {
	t.Run("Echo", func(t *testing.T) {
		input := "test\n"
		r, w, _ := os.Pipe()

		// Save the original stdin
		originalStdin := os.Stdin
		// Save the original stdout
		originalStdout := os.Stdout
		os.Stdin = r
		os.Stdout = w

		_, err := w.Write([]byte(input))
		if err != nil {
			return
		}

		bot.Echo()

		err = w.Close()
		if err != nil {
			return
		}

		os.Stdin = originalStdin
		os.Stdout = originalStdout

		// Read the output from the reader
		var buf bytes.Buffer
		_, err = io.Copy(&buf, r)
		if err != nil {
			return
		}

		// Convert buffer to string and check the output
		output := buf.String()
		if output != input {
			t.Errorf("expected %q, but got %q", input, output)
		}
	})
}
