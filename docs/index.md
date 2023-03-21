# Xata Provider

Xata Terraform Provider.

## Example Usage

```hcl
terraform {
  required_providers {
    xata = {
      source = "mrkresnofatih/xata"
      version = "1.0.3-dev" # check for the latest version
    }
  }
}

provider "xata" {
  api_key = "1234567890abcdefghijklmnopqrstuvwxyz"
}
```

## Argument Reference

* List any arguments for the provider block.
