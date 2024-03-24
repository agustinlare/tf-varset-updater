# README - Terraform Cloud Variable Updater gin-swagger

This application is a RESTful API built with Go Gin framework and integrated with Swagger for API documentation. It provides endpoints to interact with a Terraform API.

## Prerequisites

Before running this application, ensure you have the following dependencies installed:

- [Docker](https://docs.docker.com/get-docker/)
- [Go](https://golang.org/doc/install) (if you prefer running the application directly without Docker)

## Getting Started

Follow these steps to set up and run the application:

1. Clone the repository:

   ```bash
   git clone <repository_url>
   ```
2. Navigate to the project directory:
    ```bash
    cd <project_directory>
    ```
3. Build the Docker image:
    ```bash
    docker build -t varset-updater .
    ```
4. Run the Docker container:
    ```bash
    docker run -p 8080:8080 varset-updater
    ```
The application swagger will be accesible at `http://localhost:8080`

## API Endpoints

- `GET /api/v1/listvariablesets`: Retrieves a list of variable sets.
- `GET /api/v1/showvariableset/{varset_id}`: Shows details of a specific variable set.
- `GET /api/v1/showvariable/{var_id}`: Shows details of a specific variable.
- `POST /api/v1/appendvariable/{varset_id}`: Appends a variable to a variable set.
- `PATCH /api/v1/updatevariable/{varset_id}`: Updates the value of a variable in a variable set.

## Swagger

Swagger for this API is available at `http://localhost:8080/swagger/index.html`. It provides detailed information about each endpoint, including request parameters, response schemas, and example usage.

## Swagger documentation

The command swag init is part of the swaggo/swag tool, used for integrating Swagger documentation into Go applications. When swag init is executed, the tool scans the Go project for files containing special Swagger-style comments and automatically generates a docs/docs.go file containing Swagger documentation in JSON format. This documentation is then used to generate the Swagger user interface (UI) that describes and documents the API endpoints.

### Terraform API Base URL

The Terraform API base URL used in this application is:

```bash
https://app.terraform.io/api/v2
```

### Curl examples

```
## List variablesets
curl \
  --header "Authorization: Bearer $TOKEN" \
  --header "Content-Type: application/vnd.api+json" \
"https://app.terraform.io/api/v2/organizations/$ORG/varsets"

## Show variablesets
curl \
  --header "Authorization: Bearer $TOKEN" \
  --header "Content-Type: application/vnd.api+json" \
  --request GET \
  https://app.terraform.io/api/v2/varsets/$VARIABLE_SET

## Show variables
curl \
  --header "Authorization: Bearer $TOKEN" \
  --header "Content-Type: application/vnd.api+json" \
  --request GET \
  https://app.terraform.io/api/v2/vars/var-ujiVL4ZdzuicxZT2
```

## Dockerfile

The Dockerfile included in this repository builds the application image in two stages:

1. **Build Stage (`build`)**: Downloads dependencies, initializes Swagger, and builds the application binary.
2. **Runtime Stage**: Copies the built binary from the previous stage and sets up the runtime environment.

### Notes

- Ensure you have proper authentication and authorization to access the Terraform API endpoints.
- Customize the application behavior according to your requirements by modifying the codebase.

Feel free to contribute to this project or report any issues you encounter!
