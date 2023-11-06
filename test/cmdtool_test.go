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

func TestFlagSet_AddFlag(t *testing.T) {
	fs := &cmdtool.FlagSet{}
	flag := &cmdtool.Flag{
		Name:        []string{"flag1", "flag2"},
		Description: "Test flag",
		Value:       "test",
		BoolFlag:    true,
	}

	fs.AddFlag(flag)

	if !reflect.DeepEqual(fs.Flags[0], flag) {
		t.Errorf("Expected flag to be added to cmdtool.FlagSet, but got %v", fs.Flags[0])
	}
}
func TestFlagSet_GetValue(t *testing.T) {
	fs := &cmdtool.FlagSet{}
	flag1 := &cmdtool.Flag{
		Name:        []string{"flag1"},
		Description: "Test flag 1",
		Value:       "test1",
		BoolFlag:    true,
	}
	flag2 := &cmdtool.Flag{
		Name:        []string{"flag2"},
		Description: "Test flag 2",
		Value:       "test2",
		BoolFlag:    true,
	}

	fs.AddFlag(flag1)
	fs.AddFlag(flag2)

	tests := []struct {
		name     string
		flagName string
		want     interface{}
	}{
		{
			name:     "Get flag1 value",
			flagName: "flag1",
			want:     "test1",
		},
		{
			name:     "Get flag2 value",
			flagName: "flag2",
			want:     "test2",
		},
		{
			name:     "Get non-existent flag value",
			flagName: "flag3",
			want:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fs.GetValue(tt.flagName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cmdtool.FlagSet.GetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestFlagSet_IsDefined(t *testing.T) {
	fs := &cmdtool.FlagSet{}
	flag1 := &cmdtool.Flag{Name: []string{"flag1", "f1"}}
	flag2 := &cmdtool.Flag{Name: []string{"flag2", "f2"}}
	fs.AddFlag(flag1)
	fs.AddFlag(flag2)

	tests := []struct {
		name     string
		flagName string
		want     bool
	}{
		{"defined flag", "flag1", true},
		{"undefined flag", "flag3", false},
		{"defined flag with alias", "f2", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fs.IsDefined(tt.flagName); got != tt.want {
				t.Errorf("cmdtool.FlagSet.IsDefined() = %v, want %v", got, tt.want)
			}
		})
	}
}
