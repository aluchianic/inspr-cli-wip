package test

import (
	"github.com/stretchr/testify/assert"
	"inspr-cli/pkg/tests"

	"fmt"
	"testing"
)

var (
	testWorkspaceName = "inspr"
)

func TestCreateWorkspace(t *testing.T) {
	cli := fmt.Sprintf("inspr-cli init workspace %s", testWorkspaceName)
	lines, _, err := tests.ExecuteCli(cli)
	assert.NoError(t, err)
	fmt.Printf("%v", lines)
}

func TestCreateWorkspace_fail(t *testing.T) {
	cli := fmt.Sprintf("inspr-cli init workspace %s", testWorkspaceName)
	_, _, err := tests.ExecuteCli(cli)
	assert.Errorf(t, err, "already exist")

}
