package strings

import (
	"fmt"
	"strings"
	"time"
)

// EnsureFowardSlashAtStringEnd makes such "/" is at the end of a string (useful for file processing).
// Note: this isn't useful for all file systems, but good enough for Unix and Windows-style OSes.
func EnsureFowardSlashAtStringEnd(valDir string) string {
	if !strings.HasSuffix(valDir, "/") {
		valDir += "/"
	}
	return valDir
}

// EnsureFowardSlashAtStringPrefix makes such "/" is at the beginning of a string (useful for file processing).
// Note: this isn't useful for all file systems, but good enough for Unix and Windows-style OSes.
func EnsureFowardSlashAtStringPrefix(valDir string) string {
	if !strings.HasPrefix(valDir, "/") {
		valDir += "/"
	}
	return valDir
}

func MsgDoThenDoneDefault(actionWords string) (msgStart, msgDone func() string) {
	now := time.Now()
	msgStart = func() string {
		return fmt.Sprintf("≈ %v...", actionWords)
	}
	msgDone = func() string {
		end := time.Now()
		elapsed := end.Sub(now)
		return fmt.Sprintf("√ %v • Elapsed: %v", actionWords, elapsed.Round(time.Millisecond))
	}
	return msgStart, msgDone
}
