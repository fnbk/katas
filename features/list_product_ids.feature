Feature: list product IDs
  In order to use the products API
  As an API user
  I need to be able to list product IDs

  Scenario: multiple IDs
    Given there are IDs in the product database:
      | ID |
      | 1  |
      | 2  |
    When I send "GET" request to "/products/"
    Then the response code should be 200
    And the response should match json:
      """
      ["1","2"]
      """

  Scenario: no IDs
    Given there are IDs in the product database:
      | ID |
    When I send "GET" request to "/products/"
    Then the response code should be 200
    And the response should match json:
      """
      []
      """
