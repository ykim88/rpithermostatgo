package filesystem

import (
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"path"
	"testing"
)

func TestOnFileModified(t *testing.T) {
	fullPath := CreateTestFile()
	watcher := NewFsWatcher(fullPath)
	defer watcher.Stop()
	changeEvent, err := watcher.Start()
	go exec.Command("touch", fullPath).Run()

	event := <-changeEvent

	assert.Nil(t, err)
	assert.NotNil(t, event)
	Clean(fullPath)
}

func Clean(fullPath string) {
	os.Remove(fullPath)
}

func CreateTestFile() string {
	directory := os.TempDir()
	fullPath := path.Join(directory, "FileTest")
	os.Create(fullPath)
	return fullPath
}
