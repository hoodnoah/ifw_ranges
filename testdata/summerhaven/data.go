package summerhaven

import (
	"io/ioutil"
	"path/filepath"
	"runtime"
)

func ListFileNames() []string {
	return []string{
		"summerhaven_08_03_2023.html",
		"summerhaven_11_28_2023.html",
		"summerhaven_06_23_2024.html",
		"summerhaven_07_16_2024.html",
		// "summerhaven_08_06_2024.html",
		// "summerhaven_09_24_2024.html",
		// "summerhaven_01_13_2025.html",
		// "summerhaven_01_28_2025.html",
		// "summerhaven_03_20_2025.html",
		// "summerhaven_04_03_2025.html",
		// "summerhaven_04_21_2024.html",
		// "summerhaven_04_30_2025.html",
		// "summerhaven_05_07_2025.html",
		"summerhaven_06_09_2025.html",
	}
}

func LoadHTML(filename string) ([]byte, error) {
	_, thisFile, _, _ := runtime.Caller(0)
	baseDir := filepath.Join(filepath.Dir(thisFile), "html")
	path := filepath.Join(baseDir, filename)
	return ioutil.ReadFile(path)
}
