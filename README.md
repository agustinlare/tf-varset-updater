# Module varsetupdate documentation

The package varsetupdate provides functions for interacting with the Terraform Cloud API to update variables within a variable set. It allows users to get information about a variable set, list variables within a set, and create, update or delete variables.

## Types

- `VariableSet`: represents a Terraform Cloud variable set.
- `Variable`: represents a variable in a Terraform Cloud variable set.
- `VarList`: represents a list of variables in a Terraform Cloud variable set.
- `NewVariable`: represents a new variable to be added to a Terraform Cloud variable set.

## Functions

- `getVarset`(varsetid string, token string) (*VariableSet, error): retrieves a variable set from a Terraform Cloud workspace given its ID and an authentication token.
- `getVars`(varsetid string, token string) (*VarList, error): retrieves a list of variables in a variable set from a Terraform Cloud workspace given its ID and an authentication token.
- `appendVar`(varsetid string, token string, newvar NewVariable) (*Variable, error): creates a new variable in a variable set in a Terraform Cloud workspace given its ID, an authentication token, and a NewVariable object.
- `updateVar`(varid string, token string, newvar NewVariable) (*Variable, error): updates a variable in a Terraform Cloud workspace given its ID, an authentication token, and a NewVariable object.

Note: All functions in this package make use of the Terraform Cloud API to perform their operations.