---
page_title: "doctest Provider"
---

# doctest Provider

An empty provider for testing documentation publishing on the Terraform Registry.
It requires no credentials or configuration.

-> The sample resources are documentation-only. The provider does not implement any resources or data sources.

## Example Usage

```hcl
terraform {
  required_providers {
    doctest = {
      source  = "gerrytan/doctest"
      version = "0.1.0"
    }
  }
}

provider "doctest" {}
```

## Argument Reference

The provider has no configuration arguments.
