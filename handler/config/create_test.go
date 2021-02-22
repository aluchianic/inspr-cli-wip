package config

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	test "inspr-cli/pkg/tests"
	"testing"
)

func TestCreateWorkspace(t *testing.T) {
	workspaceName := "test-workspace"
	cli := fmt.Sprintf("inspr config workspace %s", workspaceName)
	lines, _, err := test.ExecuteCli(cli)
	assert.NoError(t, err)
	fmt.Printf("%v", lines)
}

func TestCreateApp(t *testing.T) {
	appNames := []string{"test-app-1", "test-app-2"}
	cli := fmt.Sprintf("inspr config app %s", appNames)
	lines, _, err := test.ExecuteCli(cli)
	assert.NoError(t, err)
	fmt.Printf("%v", lines)
}
