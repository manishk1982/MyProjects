Feature: User information
  Scenario: Check login functionality
    Given user enters name and password
             """
             Tutorialspoint Behave
             Topic – Multiline Text
             """
    Then user should be logged in


#Feature: showing off behave
  Scenario: run a simple test
    Given we have behave installed
    And todo
    When we implement a test
    Then behave will test it for us!


  Scenario: test
    Given Manish
    When Shweta
    Then Hitarth
    And Sushrut