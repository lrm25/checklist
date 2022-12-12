package filecopy

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const FILE_HEADER = "****************************** COPY ******************************\n"

type FileCopy struct {
	originalPath string
	copyPath     string
}

func NewFileCopy(originalPath string) *FileCopy {
	return &FileCopy{
		originalPath: originalPath,
		copyPath:     "",
	}
}

func (f *FileCopy) Create() error {

	fileStrings, err := ioutil.ReadFile(f.originalPath)
	if err != nil {
		return err
	}

	newFileStrings := []string{FILE_HEADER, ""}
	for _, fileString := range fileStrings {
		newFileStrings = append(newFileStrings, string(fileString))
	}

	fileSplit := strings.Split(f.originalPath, string(os.PathSeparator))
	origFileName := fileSplit[len(fileSplit)-1]

	nowString := time.Now().Format("20060102150405")
	copyPath := filepath.Join(os.TempDir(), origFileName+"."+nowString)

	copyFile, err := os.Create(copyPath)
	if err != nil {
		return err
	}
	defer copyFile.Close()

	for _, s := range newFileStrings {
		newString := strings.Replace(s, "\r", "", -1)
		if _, err := copyFile.WriteString(newString); err != nil {
			return err
		}
	}
	f.copyPath = copyFile.Name()

	return nil
}

func (f *FileCopy) Open() error {
	cmd := exec.Command("gvim", f.copyPath)
	err := cmd.Start()
	if err != nil {
		return err
	}
	return nil
}
