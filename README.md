# PIM

This Product Information Management (PIM) project is build with the purpose in mind to demonstrate an improved software development process. 

important aspects:

* domain (manage product information including workarounds imposed by the environment)
* technology (HTTP-API, Queue-API, File-Import-API, databases, different programming languages)
* methodology (IODA, Softwarecell, Softwareuniverse, FlowDesign, Slicing, ...)
* psychology (pair programming, brain patterns, clean code principles)
* workflow (agile, git, CI)
* testing (unit tests, module tests, acceptance tests/blackbox testing using the Gherking language)

# Features

**1. edit, validate, store:**
* edit attributes (HTML)
* show validation result
* store product changes

**2. channel in, channel out:**
* import product changes ('fixed attributes')
* export product changes
* write 'channel in' response
* store 'channel out' response
* introduce quality level for 'channeling out' products

**3. change 'structure' for the 'product':**
* deactive 'validation attributes' (if they are not part of the structure anymore)
* validate 'validation attributes'
* show validation results
* store and 'channel out' product changes

**4. workarounds:**
* import product changes (some 'fixed attributes' may turn into 'validation attributes')
* validate 'validation attributes'
* store validation results
* store and 'channel out' product changes

**5. access rights (authorisation):**
* all product changes need the correct access rights based on some products 'fixed attribute'
* access rights: read from DB
* store and 'channel out' responses

**6. change history:**
* store which user changed what data

**7. complex structures:**
* edit structures
* validate products
* channel out product changes
* channel out structe changes

**8. csv import:**
* edit product 'validation attributes'
* store and 'channel out' product changes
* email report: validation, authorisation (scatter/gather)

**9. performance - buffering:**
* buffer changes in complex structures (channel out)

**10. additional features:**
* search
* formats mappers (in/out channels)
* different programming languages (in/out channels)

# Setup and Execution

**download sources**
```
git clone https://github.com/fnbk/pim.git
cd pim
```

**install go**
```
brew install go
```

**run tests**
```
go test
```
