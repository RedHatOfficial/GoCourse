Feature: simple accumulator checks
  An accumulator must be able to add a number to its content

  Scenario Outline: Accumulate multiple values
    Given I have an accumulator with 0
    When I add <amount> to accumulator
    Then the accumulated result should be <accumulated>
    When I add <amount2> to accumulator
    Then the accumulated result should be <accumulated2>

  Examples:
   |amount|accumulated|amount2|accumulated2|
   | 0    | 0         | 0     | 0          |
   | 1    | 1         | 1     | 2          |
   | 2    | 2         | 2     | 4          |
   | 10   | 10        | 10    | 20         |

