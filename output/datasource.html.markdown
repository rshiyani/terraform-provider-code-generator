---
subcategory: "Contract"
layout: "aci"
page_title: "Aci: aci_contract"
sidebar_current: "docs-aci-data-source-contract"
description: |-
  Data source for Aci Contract
---

# aci_contract #
Data source for Aci Contract

## Example Usage ##

```hcl
data "aci_contract" "example" {
  tenant_dn = [TODO]: Write values here
  name = [TODO]: Write values here
}
```
## Argument Reference ##
* `tenant_dn` - (Required) tenant DN.

* `name` - (Required) contract name.





## Attribute Reference



* `prio` - (Optional) prio.



