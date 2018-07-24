package integration

import (
	"io/ioutil"
	"os"
	"fmt"
	"path/filepath"
	"github.com/pmezard/go-difflib/difflib"
	. "github.com/onsi/gomega"
)


// CompareDir returns false if the two directories have different contents
func CompareDir(expected, actual string) (bool, error) {
	expectedDir, err := ioutil.ReadDir(expected)
	Expect(err).NotTo(HaveOccurred())

	actualDir, err := ioutil.ReadDir(actual)
	Expect(err).NotTo(HaveOccurred())

	expectedMap := make(map[string]os.FileInfo)
	expectedFilenamesMap := make(map[string]struct{})
	for _, file := range expectedDir {
		expectedMap[file.Name()] = file
		expectedFilenamesMap[file.Name()] = struct{}{}
	}

	actualMap := make(map[string]os.FileInfo)
	actualFilenamesMap := make(map[string]struct{})
	for _, file := range actualDir {
		actualMap[file.Name()] = file
		actualFilenamesMap[file.Name()] = struct{}{}
	}

	Expect(actualFilenamesMap).To(Equal(expectedFilenamesMap), fmt.Sprintf("Contents of directories %s (expected) and %s (actual) did not match", expected, actual))

	for name, expectedFile := range expectedMap {
		actualFile, ok := actualMap[name]
		Expect(ok).To(BeTrue())
		Expect(actualFile.IsDir()).To(Equal(expectedFile.IsDir()))

		expectedFilePath := filepath.Join(expected, expectedFile.Name())
		actualFilePath := filepath.Join(actual, actualFile.Name())

		if expectedFile.IsDir() {
			// compare child items
			result, err := CompareDir(expectedFilePath, actualFilePath)
			if !result || err != nil {
				return result, err
			}
		} else {
			// compare expectedFile contents
			expectedContents, err := ioutil.ReadFile(expectedFilePath)
			Expect(err).NotTo(HaveOccurred())
			actualContents, err := ioutil.ReadFile(actualFilePath)
			Expect(err).NotTo(HaveOccurred())

			diff := difflib.UnifiedDiff{
				A:        difflib.SplitLines(string(expectedContents)),
				B:        difflib.SplitLines(string(actualContents)),
				FromFile: "expected contents",
				ToFile:   "actual contents",
				Context:  3,
			}

			diffText, err := difflib.GetUnifiedDiffString(diff)
			Expect(err).NotTo(HaveOccurred())
			Expect(diffText).To(BeEmpty(), fmt.Sprintf("Contents of files %s (expected) and %s (actual) did not match", expectedFilePath, actualFilePath))
		}
	}

	return true, nil
}
