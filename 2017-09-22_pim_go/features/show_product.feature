Feature: show product
  In order to use the products API
  As an API user
  I need to be able to show a product based on an ID

	# P_In:{A1X} S:{A1} P_Out:{A1X}
  Scenario: attribute exists in product and structure
    Given a product with the ID "1" and StructureID "2" exists
    And a structure with the ID "2" exists
		And product and structure with attributes and settings in the database:
      | Product | Structure |
      | A1,X    | A1        |
    When I send "GET" request to "/products/1"
    Then the response code should be 200
    And the response should match json:
      """
			{"Attributes":[{"ID":"","Name":"A1","State":"","Value":"X"}],"Bs":null,"ID":"1","Name":"","StructureID":"2"}
      """

	# P_In:{A1X} S:{} P_Out:{}
  Scenario: attribute exists in product but does not exist in structure
    Given a product with the ID "1" and StructureID "2" exists
    And a structure with the ID "2" exists
		And product and structure with attributes and settings in the database:
      | Product | Structure |
      | A1,X    |           |
    When I send "GET" request to "/products/1"
    Then the response code should be 200
    And the response should match json:
      """
			{"Attributes":[],"Bs":null,"ID":"1","Name":"","StructureID":"2"}
      """

	# P_In:{A1X,B1X} S:{B1} P_Out:{B1X}
  Scenario: two attributes in product, one attribute in structure
    Given a product with the ID "1" and StructureID "2" exists
    And a structure with the ID "2" exists
		And product and structure with attributes and settings in the database:
      | Product | Structure |
      | A1,X    |           |
      | B1,X    | B1        |
    When I send "GET" request to "/products/1"
    Then the response code should be 200
    And the response should match json:
      """
      {"Attributes":[],"Bs":[{"Attributes":[{"ID":"","Name":"B1","State":"","Value":"X"}],"Cs":null,"ID":"","Name":""}],"ID":"1","Name":"","StructureID":"2"}
      """


	# P_In:{A1X} S:{A1,A2} P_Out:{A1X,A2E}
  Scenario: one attribute in product, two attributes in structure
    Given a product with the ID "1" and StructureID "2" exists
    And a structure with the ID "2" exists
		And product and structure with attributes and settings in the database:
      | Product | Structure |
      | A1,X    | A1        |
      |         | A2        |
    When I send "GET" request to "/products/1"
    Then the response code should be 200
    And the response should match json:
      """
			{"Attributes":[{"ID":"","Name":"A1","State":"","Value":"X"},{"ID":"","Name":"A2","State":"","Value":""}],"Bs":null,"ID":"1","Name":"","StructureID":"2"}
      """

	# P_In:{A1X,B1X,C1X} S:{A1,A2,B1,B2,C1,C2} P_Out:{A1X,A2E,B1X,B2E,C1X,C2E}
  Scenario: one attribute in product, two attributes in structure
    Given a product with the ID "1" and StructureID "2" exists
    And a structure with the ID "2" exists
		And product and structure with attributes and settings in the database:
      | Product | Structure |
      | A1,X    | A1        |
      | B1,X    | A2        |
      | C1,X    | B1        |
      |         | B2        |
      |         | B2        |
      |         | C1        |
      |         | C2        |
    When I send "GET" request to "/products/1"
    Then the response code should be 200
    And the response should match json:
      """
			{"Attributes":[{"ID":"","Name":"A1","State":"","Value":"X"},{"ID":"","Name":"A2","State":"","Value":""}],"Bs":[{"Attributes":[{"ID":"","Name":"B1","State":"","Value":"X"},{"ID":"","Name":"B2","State":"","Value":""}],"Cs":[{"Attributes":[{"ID":"","Name":"C1","State":"","Value":"X"},{"ID":"","Name":"C2","State":"","Value":""}],"ID":"","Name":""}],"ID":"","Name":""}],"ID":"1","Name":"","StructureID":"2"}
      """
