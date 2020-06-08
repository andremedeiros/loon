---
title: User Configuration
author: Andre Medeiros
timestamp: 20200527142905
status: Proposal
---

# User Configuration

This document proposes a specification for the user configuration aspect of the tool. It explains where the configuration lives, what is possible to configure, and what those values should look like.

## Parameters

### Defaults

| Configuration | Default |
|:--|:--|
| `dev_tld` | `.test` |
| `provider` | `nix` |
| `silence_dev_dev_tld` | false |
| `source_tree` | `$HOME/src/{host}/{owner}/{name}` |

### `dev_tld`

This is the Top Level Domain for the development domains. It MUST start with a `.` and conform to [RFC 1035, Section 3.1](https://tools.ietf.org/html/rfc1035#section-3.1)

When this value is changed, the tool MUST ensure that there is a resolver configuration available for the new TLD.

Special cases:
* `.dev` MUST warn the user that Google Chrome forces HTTPS on that TLD, unless `silence_dev_dev_tld` is enabled. It MUST explain why this is a problem[^1]

### `provider`

This sets the default provider for the user's environment. Possible values are:
* `nix`
* `docker`
* `homebrew`

A provider MUST implement checks to verify whether the environment is set up in a way where it can satisfy a project's dependencies. For instance:
* `nix` should check whether there's a running nix environment, and it should know how to install itself
* `homebrew` should check whether the user is running on Windows, and fail accordingly

### `silence_dev_dev_tld`

This configuration value silences warnings about using the `.dev` development TLD (see above.)

### `source_tree`

This contains the template on which source trees will be calculated. This template will allow for the following variables:

| Variable | Default | Required |
|:--|:--|:--|
| `host` | `github.com` | |
| `owner` | The current UNIX username | |
| `name` | The name of the repository | X |

The template MUST validate that the resulting path is valid.

Once the template is rendered, the resulting stream MUST resolve the environment variables that are present. For instance, `$HOME/src/github.com/andremedeiros/loon` would have to resolve `$HOME`.

## Reading

The expected location for the configuration is `$XDG_CONFIG_HOME/loon/config.yml`. The configuration API MUST handle the following cases:
* `$XDG_CONFIG_HOME`, `$XDG_CONFIG_HOME/loon`, `$XDG_CONFIG_HOME/loon/config.yml` not existing, in which case the default configuration will be returned
* Invalid formatting, in which case an error will be shown on screen and the tool will exit
* Invalid value for a specific configuration value, in which case an error will be shown on screen and the tool will exit

[^1]: https://webdevstudios.com/2017/12/12/google-chrome-63/
