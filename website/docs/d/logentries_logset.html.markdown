---
layout: "logentries"
page_title: "Logentries: logentries_logset"
sidebar_current: "docs-logentries-datasource-logentries_logset"
description: |-
  Get information on Logentries LogSet.
---

# logentries\_logset

Use this data source to get information (like ID) of already existing Logentries LogSets.

## Example Usage

```hcl
data "logentries_logset" "logset" {
  name = "Example_Logset"
}

output "example_logset_id" {
  value = "data.logentries_logset.logset.id"
}
```

## Attributes Reference

* `name` - Name of the LogSet to query.
