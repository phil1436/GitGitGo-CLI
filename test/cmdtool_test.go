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

func TestFlag_ToString(t *testing.T) {
	tests := []struct {
		name       string
		flag       *cmdtool.Flag
		wantResult string
	}{
		{
			name: "single letter flag",
			flag: &cmdtool.Flag{
				Name:        []string{"f"},
				Description: "Test flag",
				Value:       nil,
			},
			wantResult: "   -f Test flag",
		},
		{
			name: "multi-letter flag",
			flag: &cmdtool.Flag{
				Name:        []string{"flag1", "flag2"},
				Description: "Test flag",
				Value:       nil,
			},
			wantResult: "   [-f|flag2] Test flag",
		},
		{
			name: "flag with value",
			flag: &cmdtool.Flag{
				Name:        []string{"f"},
				Description: "Test flag",
				Value:       "test",
			},
			wantResult: "   -f Test flag (optional)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := tt.flag.ToString(); gotResult != tt.wantResult {
				t.Errorf("Flag.ToString() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
