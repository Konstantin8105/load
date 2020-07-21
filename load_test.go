package load_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Konstantin8105/binaryexpr"
)

func Test(t *testing.T) {
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(info.Name(), ".go") {
			return nil
		}
		return binaryexpr.Test(path)
	})
	if err != nil {
		t.Error(err)
	}
}
