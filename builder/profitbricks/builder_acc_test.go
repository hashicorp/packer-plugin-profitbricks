// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package profitbricks

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/acctest"
)

func TestBuilderAcc_basic(t *testing.T) {
	testCase := &acctest.PluginTestCase{
		Setup:    testAccPreCheck,
		Template: testBuilderAccBasic,
		Check: func(buildCommand *exec.Cmd, logfile string) error {
			if buildCommand.ProcessState != nil {
				if buildCommand.ProcessState.ExitCode() != 0 {
					return fmt.Errorf("Bad exit code. Logfile: %s", logfile)
				}
			}
			return nil
		},
	}

	acctest.TestPlugin(t, testCase)

}

func testAccPreCheck() error {
	if v := os.Getenv("PROFITBRICKS_USERNAME"); v == "" {
		return fmt.Errorf("PROFITBRICKS_USERNAME must be set for acceptance tests")
	}

	if v := os.Getenv("PROFITBRICKS_PASSWORD"); v == "" {
		return fmt.Errorf("PROFITBRICKS_PASSWORD must be set for acceptance tests")
	}

	return nil

}

const testBuilderAccBasic = `
{
	"builders": [{
	      "image": "Ubuntu-16.04",
	      "password": "password",
	      "username": "username",
	      "snapshot_name": "packer",
	      "type": "profitbricks"
   	}]
}
`
