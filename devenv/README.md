# Devenv

Devenv is a simple tool to manage the setup of consistent developer environments
for projects containing multiple languages.

It's primary motivation is to enable the local developer workflow to be replicated
under CI/CD, and vice versa.

(It will also always only be a single binary install for maximum portability)

## Design

Devenv makes heavy use of existing language package managers and similar tools,
thanks to:

- Homebrew[^1] - for system wide dependencies and recommendations
- Trunk.io - for hermetically managed project specific dependencies and code
  quality automation
- go-task - for implementing scripts to perform common project tasks

[^1]: Also works on Linux, previously known as [Linuxbrew](https://docs.brew.sh/Homebrew-on-Linux)

## Usage

You can download the binary directly from the releases page, but it's prefered
distribution is via Homebrew, which will include it's core dependencies:

`brew install metafeather/tools/devenv`

You will need run `devenv init` to install it's embedded assets before use.

It is then recommended to run `devenv doctor` to see if there is anything else
required.

For use in managing git project dependencies and linters you will most likely
need be prompted to install `trunk-io`[^2]

[^2]: Trunk Check runs 100+ idiomatic code-checking tools for every language and technology, locally https://docs.trunk.io/check/cli
