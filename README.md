![banner](https://user-images.githubusercontent.com/9689/85356663-8035c100-b4dd-11ea-99ec-4e969ccc87e2.png)

Made with :heart: by [Andre Medeiros](https://andre.cool)

## Show, don't tell

[![asciicast](https://asciinema.org/a/3R3uiG2jTnW6Pb1wo6gFwRnah.png)](https://asciinema.org/a/3R3uiG2jTnW6Pb1wo6gFwRnah)

## What is it?

Loon is a developer tool that aims to facilitate getting started on various projects.

## Getting started

```
$ go get -u github.com/andremedeiros/loon
$ loon shellrc >> ~/.bashrc # if you're a bash user
$ loon shellrc >> ~/.zshrc  # if you're a zsh user
```

### A simple payload

```yaml
# ~/src/github.com/you/awesome-project/loon.yml
name: Your Awesome Project
url: https://github.com/you/awesome-project
deps:
  - postgres
  - redis
  - ruby: 2.7.1
  - node: latest
tasks:
  setup:
    description: Fetches dependencies and runs database migrations
    command: |
      bundle check || bundle install
      rake db:setup
      rake db:migrate
  server:
    description: Starts the web server
    command: bundle exec rails server
  jobs:
    description: Starts the job worker
    command: bundle exec sidekiq
  test:
    description: Runs the test suite
    command: bundle exec rails test
```

With this example, when you run:

```
$ loon up
```

The tool will do a few things, in order:

- It will generate a Nix derivation that is specific for your project's combination of services and languages.
- It ensures the derivation can be satisfied, downloads packages if built, or sources if not.
- It generates a unique IP for your project that is bound to `localhost`
- It initializes services (`initdb` for Postgres, `mysqld --initialize-insecure` for MySQL)
- It starts services bound to that IP and their default ports

After that, you can start running commands. In this case, by running `loon setup`, the tool:

- Generates an environment for the project, by injecting [Twelve-Factor](https://12factor.net) friendly config
- Executes the specific version of the invoked tool, or the latest supported if not specified

While you're developing, `loon test` runs your tests, `loon server` starts the Rails development server, and `loon jobs` starts a job runner.

Finally, `loon down` shuts the services down.

### Services

**Service**|**Versions**
-----|-----
Memcached|1.6.5, 1.6.6
MySQL|8.0.17
Postgresql|9.5.22, 9.6.18, 10.13, 11.8, 12.3
Redis|6.0.4

### Languages

**Language**|**Versions**
-----|-----
Crystal|0.35.1
Golang|1.13.12, 1.14.4
Node|12.18.1, 14.4.0
Ruby|2.6.6, 2.7.1

## Acknowledgements

* [Mark Imbriaco](https://github.com/imbriaco) for gently guiding me on how to pass options to Postgres
* [Burke Libbey](https://github.com/burke) for giving me the finalizers idea
* [Jon Pulsifer](https://github.com/j0npulsifer) for showing me how to add network aliases on Linux