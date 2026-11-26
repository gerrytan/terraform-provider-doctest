---
page_title: "doctest_gadget Resource"
subcategory: "Samples"
---

# doctest_gadget Resource

A second sample resource page for testing Terraform Registry navigation and Markdown rendering.

-> This resource is documentation-only and is not implemented by the provider. The example cannot be applied.

## Example Usage

```hcl
resource "doctest_gadget" "example" {
  name        = "hello-gadget"
  description = "A sample gadget"
}
```

## Argument Reference

* `name` - (Required) The name of the sample gadget.
* `description` - (Optional) A description of the sample gadget.

## Attribute Reference

* `id` - The identifier of the sample gadget.
