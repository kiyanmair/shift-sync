# Contributing to Shift Sync

Thank you for considering contributing to Shift Sync!
This project is in its early stages, and contributions are especially welcome for adding or improving source and destination integrations.

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

1. **Format your code**

    Run `go fmt` to format your Go files for consistency.
    Most IDEs will do this for you if you have Go language features enabled.

1. **Open a pull request**
   
   Once you're ready, open a pull request (PR) in this repository.
   Please describe what your change does in the PR description, and link any issues if relevant.
   The code owners will review your PR as soon as they can.

## Adding a new integration

1. **Create a new integration file**

    Create a new Go file in [internal/integrations/source](./internal/integrations/source/) or [internal/integrations/destination](./internal/integrations/destination/), named with the format `integrationname.go`.
    You should strongly consider duplicating the `example.go` file in the same directory, as there is some boilerplate involved.

1. **Define your required fields**

    Define a `struct` type in your integration file.
    It must have an `ID string` field.
    You can define any other fields you'd like, using `mapstructure` struct tags to indicate their corresponding YAML field name.

1. **Rename your constructor**

    Rename the example constructor to reflect your integration name.
    You should not have to make changes to the constructor implementation.

1. **Implement the interface**

    Implement the relevant interface for your integration.
    For sources, this will be `Source`, requiring an `FetchUsers` method.
    For destinations, this will be `Destination`, requiring an `UpdateUsers` method.

1. **Register your constructor**

    Update the example `init` function body to reference your integration name and constructor function.
    The name should have the format `integration_name` and will match against `type` values in the configuration file.
