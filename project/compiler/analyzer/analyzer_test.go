package analyzer

import (
	"compiler/utils"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTokenFilePath(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("/home/user/some/HogeT.xml",
		GetTokenFilePath("/home/user/some/Hoge.jack"))

	assert.Equal("../../../HogeT.xml",
		GetTokenFilePath("../../../Hoge.jack"))
}

func TestTokenizer(t *testing.T) {
	assert := assert.New(t)

	testNames := []string{
		"ArrayTest",
		"Square",
		"ExpressionLessSquare",
	}
	originalDirBase := "testdata/tokenizer/original"
	targetDirBase := "testdata/tokenizer/target"

	for _, testName := range testNames {
		analyzer := NewAnalyzer()
		err := analyzer.Analyze(filepath.Join(targetDirBase, testName))
		assert.Nil(err)

		// compare outputs
		originalFiles := utils.Must2(utils.GetFilePathsInDir(
			filepath.Join(originalDirBase, testName),
		))
		for _, originalFile := range originalFiles {
			if filepath.Ext(originalFile) != ".xml" {
				continue
			}
			if originalFile[len(originalFile)-5] != 'T' {
				continue
			}

			targetFile := filepath.Join(targetDirBase, testName, filepath.Base(originalFile))
			targetFileFD := utils.Must2(os.Open(targetFile))
			targetData := utils.Must2(io.ReadAll(targetFileFD))

			originalFileFD := utils.Must2(os.Open(originalFile))
			originalData := utils.Must2(io.ReadAll(originalFileFD))

			assert.Equal(string(originalData), string(targetData), originalFile)
		}
	}
}
