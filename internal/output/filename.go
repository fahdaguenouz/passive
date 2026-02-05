package output

import (
	"fmt"
	"os"
	"path/filepath"
)

func NextResultFilename(dir string) (string, error) {
	if err := ensureDirExists(dir); err != nil {
		return "", err
	}

	base := filepath.Join(dir, "result.txt")
	if _, err := os.Stat(base); os.IsNotExist(err) {
		return base, nil
	}

	for i := 2; i < 10000; i++ {
		name := filepath.Join(dir, fmt.Sprintf("result%d.txt", i))
		if _, err := os.Stat(name); os.IsNotExist(err) {
			return name, nil
		}
	}
	return "", fmt.Errorf("too many result files in %s", dir)
}

func ensureDirExists(dir string) error {
	// if dir is ".", it exists; but this keeps it safe for custom paths later
	if dir == "." {
		return nil
	}
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return os.MkdirAll(dir, 0o755)
	}
	return err
}
