# Feature 1 - edit, validate, store

# Iteration 01

[analyse]

First off lets start with a better understanding of the domain. The goal of this feature is to implement a HTTP-Resource for Products. This includes the following actions:

* show a product: (ID) -> (product{})
* edit product attributes: (product{}) -> (product{})
* assign structure to a product: (productID, structureID) -> (product{})
* list product IDs: () -> (ID*)

![architecture overview](images/01_architecture_overview.png)

## Data-Model

**Product**
```
A: {
  ID: "",
  Name: "",
  StructureID: "",
  Attributes: [{
    ID: "",
    Name: "",
    Value: "",
    State: "{active|inactive}",
  }],
  Bs: [{
    ID: "",
    Name: "",
    Attributes: [{
      ID: "",
      Name: "",
      Value: "",
      State: "{active|inactive}",
    }],
    Cs: [{
      ID: "",
      Name: "",
      Attributes: [{
        ID: "",
        Name: "",
        Value: "",
        State: "{active|inactive}",
      }]
    }]
  }]
}
```

**Structure**
```
{
  ID: "",
  Name: "",
  Settings: [{
    ID: "",
    Name: "",
    Tier: {A|B|C},
    Datatype: {string|float32|int32},
  }]
}
```

