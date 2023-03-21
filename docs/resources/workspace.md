# Workspace Resource

The Resource to create a Workspace inside Xata.io.

## Example Usage

```hcl
resource "xata_workspace" "ruakworkspace" {
  name = "Ranca Upas"
  slug = "ruak"
}
```

## Attribute Reference

* `name` - (Required) The name for this workspace.
* `slug` - (Required) The shorthand for xata to identify this workspace.
