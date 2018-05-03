Feature: Clean local branches
  In order to have a clean git working environment
  As a developer
  I need to be able to prune all branches except for master

  Scenario: Clean local non-master branches
    Given I am in the go-erl/gitclean project repo
    And I have multiple local branches
    When I run gitclean in that repo
    Then I should only be left with the master branch
