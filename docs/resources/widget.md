---
page_title: "doctest_widget Resource"
subcategory: "Samples"
---

# doctest_widget Resource

A sample widget resource page for testing Terraform Registry documentation publishing.

-> This resource is documentation-only and is not implemented by the provider. The example cannot be applied.

## Example Usage

```hcl
resource "doctest_widget" "example" {
  name = "hello-widget"
}
```

## Argument Reference

* `name` - (Required) The name of the sample widget.

## Attribute Reference

* `id` - The identifier of the sample widget.
