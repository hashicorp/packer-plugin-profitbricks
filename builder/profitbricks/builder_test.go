// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package profitbricks

import (
	"fmt"
	"testing"

	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

func testConfig() map[string]interface{} {
	return map[string]interface{}{
		"image":         "Ubuntu-16.04",
		"password":      "password",
		"username":      "username",
		"snapshot_name": "packer",
		"type":          "profitbricks",
	}
}

func TestImplementsBuilder(t *testing.T) {
	var raw interface{} = &Builder{}
	if _, ok := raw.(packersdk.Builder); !ok {
		t.Fatalf("Builder should be a builder")
	}
}

func TestBuilder_Prepare_BadType(t *testing.T) {
	b := &Builder{}
	c := map[string]interface{}{
		"api_key": []string{},
	}

	_, warns, err := b.Prepare(c)
	if len(warns) > 0 {
		t.Fatalf("bad: %#v", warns)
	}
	if err == nil {
		fmt.Println(err)
		fmt.Println(warns)
		t.Fatalf("prepare should fail")
	}
}

func TestBuilderPrepare_InvalidKey(t *testing.T) {
	var b Builder
	config := testConfig()

	config["i_should_not_be_valid"] = true
	_, warnings, err := b.Prepare(config)
	if len(warnings) > 0 {
		t.Fatalf("bad: %#v", warnings)
	}
	if err == nil {
		t.Fatal("should have error")
	}
}
