# shift-sync

`shift-sync` is a tool for syncing on-shift users between arbitrary systems.
For example, you can sync on-call users from a Better Stack schedule to a Slack group.

The guiding principle of `shift-sync`'s design is that adding new integrations should be easy.
Simply specify the configuration details and implement the source or destination interface.
Contributions are very welcome!

**Status**:
This project is in early development.
While the core framework is functional, expect breaking changes as the design is refined.
We follow [Semantic Versioning](https://semver.org/spec/v2.0.0.html) and are in the [major version zero](https://semver.org/spec/v2.0.0.html#spec-item-4).

# Installation

The entrypoint for this tool is the `shiftsync` CLI program.
To run it, you'll need to [install Go](https://go.dev/dl/).

To build `shiftsync`, run `make build` from the root of the repository.
You will need to re-run this command if you make changes to the code.

# Usage

Run `shiftsync` on the config file `syncfile.yaml`:

```bash
./shiftsync run --config-path syncfile.yaml
```
