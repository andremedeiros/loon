# Loon

This document describes how the development helper tool, `loon`, works. It goes through the different commands the tool should ship with, and what their behaviour should be.

## Introduction

Software complexity and dependencies keep increasing. With this comes a fair amount of infrastructure that supports that software, such as database, cache, and job servers. Also, different software projects require different programming languages, or a lot of times, the same programming language, but a different version.

As the ideal scenario for a developer is to have the closest approximation possible of a production environment in their development machine, dealing with all these different versions and servers becomes a chore.

This document describes a tool that handles, and hides, all that complexity and mess away, so that developers can focus on creating the actual products.

Adoption success metrics:
* Being able to do away with the [Scripts To Rule Them All](https://github.com/github/scripts-to-rule-them-all) for any given project

### Goals

* Infrastructural dependencies
	* Databases (Postgres, MySQL)
	* Caches (Redis, Memcached)
	* Queues (RabbitMQ, Faktory)
	* Monitoring (Zipkin, StatsD, Grafana, Grafite)
* Languages (Go, Ruby, Elm, JavaScript)
* Development to Deployment (Heroku with Terraform)
* Custom tasks per project
* DNS and SSL for development domains

## User Experience

A project's metadata will be defined in a file that lives in the project's root. For now, we'll call it `loon.yml` and it will have the following format:

```yaml
name: awesome-tool
url: https://github.com/andremedeiros/awesome-tool
provider: nix
languages:
  ruby:
    version: 2.7.1
  golang:
    version: 1.13.1
services:
  postgres:
    version: 12.3
  faktory:
    version: 1.4.0
    provider: docker
tools:
  imagemagick:
    version: latest
hostnames:
  - www.awesome-tool.test
  - account.awesome-tool.test
environment:
  THIS_VAR: that value
  SOME_OTHER_VAR: that other value
tasks:
  test:
    description: Runs the test suite
    command: rake test
  race:
    description: Runs race checks
    command: go test -race ./...
```

Monitoring should be included for all projects, unless explicitly disabled.

The tool will assume the latest version of the config is applicable to whatever project payload file it finds.

When invoking the tool, it will traverse up the directory tree until either of these happens:
* A `loon.yml` file is found, and as such, the project root is found
* The invocation fails.

Running the tool will check whether a unix socket is listening in a pre-defined location. If not, the tool will start a server in the background and send commands to that server as it needs to. This server will be responsible for:
* Local DNS for development domains
* Certificate emitting / rotation
* Keeping track of running processes and their health

## Installation

Installation should be done via a simple shell command, much like Homebrew's.

The installation process must:
* Ask for superuser access
* Install its default provider, Nix
* Be explicit about WHAT it's running, and WHY it's running it

## Configuration

The tool must check for `$XDG_CONFIG_HOME/loon/config.yml` for configuration preferences.

### `source_tree`

The template to which the tool checks out repositories. It defaults to:

`$HOME/src/{host}/{owner}/{name}`

Where:

* `host` is one of:
	* github.com
	* gitlab.com
	* bitbucket.org
* `owner` is the repository owner's account or organization
* `name` is the repository or project name

This value must be checked to ensure that, at least, the `repository` placeholder is present.

### `dev_tld`

This configuration value defines the top level domain for development domains. It defaults to `.test` or `.localhost`

### `provider`

This configuration value sets the default provider for services, languages, and tools. Possible values are:
* nix
* docker
* homebrew

## DNS

The DNS server will resolve project hostnames to a project's aliased IP address. This IP address MUST be aliased to 127.0.0.1 to avoid any access outside the developer's machine. The IP address should be semi-random per project, but deterministic (16 bit hash to cover 10.0.X.Y?)

Default domains exposed are:
* `<project>.<dev_tld>`
* `<infrastructure>.srv.<project>.<dev_tld>`

Custom domains can also be defined in the project's metadata.

## Infrastructure

Servers should respect The Twelve-Factor App config principle and expose configuration in the environment.

Services, languages, and tools should accept a version. The tool is responsible for having a catalogue of known good versions and how to install them. Special values for `version` are:
* `latest`, which as the name indicates, will use the latest version available
* `system`, which will check the system for the tool's availability

## Providers

Supported providers:
* Nix
* Docker
* Homebrew (Mac & Linux only)

A language/service/tool and provider pair should check whether the environment is supported. For instance, installing ImageMagick with Homebrew on Windows is not possible as Homebrew does not run on that system.

## Commands

Any command that is invoked, whether it's a task or something the tool runs must:
* Check the user's global configuration
* Inherit the custom environment variables set in the project's metadata
* Ensure the tool's server is running, and if not, run it
* Respect the OS's `stdout` and `stderr` streams
* Support defining a custom UI (plain, colourized, JSON)
* Check the project's task list for any commands that overwrite the tools's commands and show an error
* Check configuration:
	* Emit warning if `dev_tld` is `.dev` (and link to Chrome's asinine

### `$ loon vet` / `$ loon doctor`

The `vet` command will check the current state of the system. It will check:
* Errors
	* Health of the provider(s) (ie. is Nix healthy?)
	* The system's DNS is setup to use our resolver for `dev_tld` domains
	* There is a reasonable amount of disk space available

The `vet` command will also check the goodness of the project's metadata. Things like this will be checked:
* Whether the providers can satisfy the requested versions
* Whether hostnames have the correct TLD

### `$ loon clone`

The `clone` command will check out a git repository into the user's development tree. By default, the development tree follows this convention:

`$HOME/src/<git provider>/<account or org>/project`

The git provider defaults to `github.com` and can be:
- github.com
- gitlab.com
- bitbucket.org

The account or organization defaults to the user's unix name.

When the repository already exists on disk, the tool must fail and exit.

### `$ loon init`

The `init` command asks the user some questions about the project and writes an initial `loon.yml` file.

The project's dependencies can be inferred or suggested based on what files exist. For instance, if there's a `.gemspec` or `gems.rb` file in the project's root, it's safe to assume it's a Ruby project. Moreover, if the `Gemfile` includes `pg`, it's safe to assume the project depends on Postgres.

### `$ loon up` / `$ loon fly`

The `up` command brings all the infrastructure into a state of goodness. This means:
* Downloading and installing any required infrastructure
* Ensuring that infrastructure is running
* Ensuring that any changes to the project's metadata get applied

### `$ loon down` / `$ loon land`

The `down` command brings down all the infrastructure. This means:
* Shutting down all servers that can be running

### `$ loon shell` / `$ loon mux`

The `shell` or `mux` command starts a shell with the project's environment all set up. This is good to further test or experiment with different tasks that can help expand the catalog of the project's task list. This command must:
- Respect the user's chosen shell
- Carry over the existing environment, merged with the project's intended environment
- Check for the presence of tmux when the `mux` v

### `$ loon deploy`

The `deploy` command generates the deployment configuration files for the specific project. This will look at the infrastructure needed and generate these configs using industry best practices.

### `$ loon cd`

The `cd` command will change directories to a project's root, following the convention for `source_tree`:

`$HOME/src/<git provider>/<account or org>/project`

The git provider defaults to `github.com`, the account or organization defaults to the user's unix name.

If the repository does not exist, the tool must offer a snippet that the user can copy to help the user to quickly clone the project.

### `$ loon pr <num>` / `$ loon pull-request <num>`

The `pr` command checks out a pull request so the user can quickly try it out. It should fail if there are changes that are not committed or not staged.

### `$ loon <missing command>`

In the case of a missing command, the tool should check whether there's a task in the project with the same command. This command must check whether the project has been brought into a state of goodness.

### `$ loon versions <language or service or tool>`

The `versions` command generates a list of the catalogue of supported versions for each language, service, and tool that `loon` knows about. If no name is provided, it shows versions for all tools.

## Helpful commands

These are commands that aren't necessarily helpful in the context of a project, but could be useful for a developer regardless:
* `loon qr <string>`
* `loon hash <string or file>`
* `loon weather`
* `loon ip`
* `loon dadjoke`
* `loon commitmsg`
* `loon ts` / `loon timestamp` / `loon now`

## Further Work

- [ ] Investigate non-colliding hashing for project IP attribution
- [ ] Custom languages, services, and tools
