# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

# For full specification on the configuration of this file visit:
# https://github.com/hashicorp/integration-template#metadata-configuration
integration {
  name = "Profitbricks"
  description = "The profitbricks plugin can be used with HashiCorp Packer to create custom images on IONOS Compute Engine."
  identifier = "packer/hashicorp/profitbricks"
  flags = ["archived"]
  component {
    type = "builder"
    name = "ProfitBricks"
    slug = "profitbricks"
  }
}
