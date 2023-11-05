package test

import (
	"phil1436/GitGitGo-CLI/pkg/cmdtool"
	"reflect"
	"testing"
)

func TestFlag_Copy(t *testing.T) {
	flag := &cmdtool.Flag{
		Name:        []string{"flag1", "flag2"},
		Description: "Test flag",
		Value:       "test",
		BoolFlag:    true,
	}

	copiedFlag := flag.Copy()

	if !reflect.DeepEqual(flag, copiedFlag) {
		t.Errorf("Expected copied flag to be equal to original flag, but got %v", copiedFlag)
	}

	// Modify the original flag and make sure the copied flag is not affected
	flag.Name[0] = "modified"
	flag.Description = "Modified description"
	flag.Value = "modified"
	flag.BoolFlag = false

	if reflect.DeepEqual(flag, copiedFlag) {
		t.Errorf("Expected copied flag to not be affected by modifications to original flag, but got %v", copiedFlag)
	}
}
