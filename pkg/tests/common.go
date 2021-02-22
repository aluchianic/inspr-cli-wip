package tests

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"

	"inspr-cli/cmd"
	"inspr-cli/pkg/util"
)

// Returns a buffer for stdout, stderr, and a function.
// The function should be used as a defer to restore the state
// See example at handle/config/create_test.go
func InsprTestSetupCli(args string) (*bytes.Buffer, *bytes.Buffer, Restorer) {
	// Save
	oldargs := os.Args
	oldStdout := util.Stdout
	oldStderr := util.Stderr

	// Create new buffers
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)

	// Set buffers
	util.Stdout = stdout
	util.Stderr = stderr
	os.Args = strings.Split(args, " ")

	return stdout, stderr, func() {
		os.Args = oldargs
		util.Stdout = oldStdout
		util.Stderr = oldStderr
	}
}

// runInspr runs a command saved in os.Args and returns the error if any
func RunInspr() error {
	return cmd.Main()
}

// Execute cli and return the buffers of standard out, standard error,
// and error.
// Callers must managage global variables using Patch from pkg/tests
func executeCliRaw(cli string) (*bytes.Buffer, *bytes.Buffer, error) {
	so, se, r := InsprTestSetupCli(cli)

	// Defer to cleanup state
	defer r()

	// Start the CLI
	err := RunInspr()

	return so, se, err
}

// Execute cli and return string lists of standard out, standard error,
// and error if any.
// Callers must manage global variables using Patch from pkg/tests
func ExecuteCli(cli string) ([]string, []string, error) {
	so, se, err := executeCliRaw(cli)

	return strings.Split(so.String(), "\n"),
		strings.Split(se.String(), "\n"),
		err
}

// Takes a volume name and size. Returns the created volume id.
// For some reason our tests container only recoganizes id and not name for some calls.
func TestCreateWorkspace(t *testing.T, workspaceName string) {
	cli := fmt.Sprintf("inspr create %s", workspaceName)
	lines, _, err := ExecuteCli(cli)
	assert.NoError(t, err)
	fmt.Printf("%v", lines)
}
