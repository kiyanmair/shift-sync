# Contributing to Shift Sync

Thank you for considering contributing to Shift Sync!
This project is in its early stages, and contributions are especially welcome for adding or improving integrations.

## General guidelines

1. **Fork the repository**

    Start by forking this repository and cloning your fork locally.
    For more guidance, see GitHub's documentation on [contributing to a project](https://docs.github.com/en/get-started/exploring-projects-on-github/contributing-to-a-project).

1. **Set up your environment**

    If you haven't already, [install Go](https://go.dev/dl/).
    Install dependencies by running `go mod tidy` in the repository root.

1. **Test your changes**

    Ensure that your changes work how you expect them to.
    At this point in time, there are no automated tests.
    However, you are welcome to contribute them.

1. **Format the code**

    Run `go fmt` to format the Go files for consistency.
    Most IDEs will do this for you if you have Go language features enabled.

1. **Open a pull request**
   
   Once you're ready, open a pull request (PR) in this repository.
   Please describe what your change does in the PR description, and link any issues if relevant.
   The code owners will review your PR as soon as they can.

## Adding a new integration

1. **Create a new integration directory**

    Create a new directory in [internal/integrations](./internal/integrations/), named with the format `integrationname`.
    You should strongly consider duplicating [foochat](./internal/integrations/foochat/), as there is some boilerplate involved.
    The integration code will have to be in `.go` files, which you can name according to their contents.

1. **Name the Go package**

    Provide the package name at the top of all of the integration's `.go` files.
    Use the same format as the directory name, i.e. `package integrationname`.

1. **Define the supplied fields**

    Define a `struct` type for the integration's supplied fields, named with the format `IntegrationName`.
    Use `mapstructure` struct tags to indicate each struct field's corresponding TOML field name.
    Don't include the `type` configuration field; this indicates the corresponding integration, but does not populate the struct.

1. **Update the constructor**

    Rename the duplicated constructor with the format `NewIntegrationName`.
    Change the type of the `i` variable to your struct type.
    You should not have to make additional changes to the constructor implementation.

1. **Register the constructor**

    Update the duplicated `init` function body to reference the integration name and constructor function.
    The name should have the format `integration_name` and will match against `type` values in the configuration file.
    Add a blank import to [importer.go](./internal/integrations/importer.go) for the integration directory.

1. **Implement the interface**

    Implement the relevant interfaces for the integration.
    Source systems will require a `FetchUsers` method.
    Destination systems will require an `UpdateUsers` method.
    Bidirectional systems will require both of these methods.
    Please do not use integration-specific external packages like SDKs, to minimise dependencies.

1. **Implement struct validation**

    Implement the `Validate` method, returning a descriptive error for any invalid configuration values.
    Bidirectional integrations may require different validity rules depending on whether the integration is being used as a source or destination.
    In this case, check the value of `direction` before performing direction-specific validation.
