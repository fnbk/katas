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

## I. Iteration: List Product IDs

[analyse]

Display a list of product IDs when making "GET" request to "/products".

![list products overview](images/01_list_products_overview.png)

1) portal: receive HTTP request
1) kernel: `productIDs()`
3) portal: respond HTTP request

[design]

![list products design](images/01_list_products_design.png)


### first increment: implement HTTP-Workflow with fake IDs

[design]

![HTTP request design](images/01_01_http_request_design.png)

[implement]

[test_bed/product_ids.go](test_bed/product_ids.go)

```
go run test_bed/product_ids.go
```

visit [http://localhost:8080/products/](http://localhost:8080/products/)


### next increment: implement productIDs()

[design]

![productIDs design](images/01_02_product_ids_design.png)

[implement]

* [product_provider.go](product_provider.go)
* [product_provider_test.go](product_provider_test.go)

```
go test -v
```



