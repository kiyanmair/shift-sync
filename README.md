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
