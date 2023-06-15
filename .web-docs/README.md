The profitbricks plugin allows you to create custom images on the IONOS Compute Engine platform.

## Installation

To install this plugin, copy and paste this code into your Packer configuration, then run [`packer init`](https://www.packer.io/docs/commands/init).

```hcl
packer {
  required_plugins {
    profitbricks = {
      source  = "github.com/hashicorp/profitbricks"
      version = "~> 1"
    }
  }
}
```

Alternatively, you can use `packer plugins install` to manage installation of this plugin.

```sh
$ packer plugins install github.com/hashicorp/profitbricks
```

## Components

### Builders

- [profitbricks](/packer/integrations/hashicorp/profitbricks/latest/components/builder/profitbricks) - The ProfitBricks Builder
  is able to create virtual machines for [IONOS Compute Engine](https://cloud.ionos.com/compute).
