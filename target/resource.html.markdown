---
subcategory: "Contract"
layout: "aci"
page_title: "Aci: aci_contract"
sidebar_current: "docs-aci-resource-contract"
description: |-
  Resource for Aci Contract
---

# aci_contract #
Resource for Aci Contract

## Example Usage ##

```hcl
	resource "aci_contract" "democontract" {
        tenant_dn = "[TODO]: Write values here"
        name = "[TODO]: Write values here"
        prio = "[TODO]: Write values here"
}
```
## Argument Reference ##
* `tenant_dn` - (Required) tenant DN.

* `name` - (Required) contract name.

* `my_map` - (Required) My map for testing.

* `prio` - (Optional) prio.


* `cast` - (Required) [TODO]: Write your Description here

* `filter` - (Required) filter list.
* `filter_name` - (Required) name of filter.

* `id` - (Optional) id of filter.


* `description` - (Optional) description of filter.


* `filter_entry` - (Required) list of filter_entry.
* `entry_next` - (Required) [TODO]: Write your Description here
* `entry_next_name` - (Required) [TODO]: Write your Description here



* `cast` - (Required) [TODO]: Write your Description here

* `filter_entry_name` - (Required) name of filter entry.

* `id` - (Optional) id of filter entry.


* `apply_to_frag` - (Optional) apply to fragment.






## Attribute Reference

The only attribute that this resource exports is the `id`, which is set to the
Dn of the Contract.
* `filter.id` - Exports this attribute for filter object. Set to the Dn for the filter managed by the contract.
* `filter.filter_entry.id` - Exports this attribute for filter entry object of filter object. Set to the Dn for the filter entry managed by the contract.

## Importing ##

An existing Contract can be [imported][docs-import] into this resource via its , via the following command: [docs-import]:
[docs-import]: https://www.terraform.io/docs/import/

```
terraform import aci_contract.example ```