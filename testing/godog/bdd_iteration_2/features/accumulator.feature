Feature: simple accumulator checks
  An accumulator must be able to add a number to its content

  Scenario: Accumulate positive integer
    Given I have an accumulator with 0
    When I add 2 to accumulator
    Then the accumulated result should be 2
