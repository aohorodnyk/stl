# Standard library for Golang (extended)

[![GitHub Workflow Status](https://github.com/aohorodnyk/stl/actions/workflows/go.yml/badge.svg)](https://github.com/aohorodnyk/stl/actions/workflows/go.yml)
[![License](https://img.shields.io/github/license/aohorodnyk/stl)](https://github.com/aohorodnyk/stl/blob/main/LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/aohorodnyk/stl)](https://github.com/aohorodnyk/stl/issues)
[![GitHub issues](https://img.shields.io/github/issues-pr/aohorodnyk/stl)](https://github.com/aohorodnyk/stl/pulls)
[![The latest release version](https://img.shields.io/github/v/release/aohorodnyk/stl)](https://github.com/aohorodnyk/stl/releases)
[![GoDoc](https://godoc.org/github.com/aohorodnyk/stl?status.svg)](https://pkg.go.dev/github.com/aohorodnyk/stl)

This package contains the library extension and simplification I always missed and could not fill the gap in my every day development. This repository has been created right after Go 1.18 release with the generics support. It was the main problem to build the simple library with good performance.
This package wont try to replace the standard library, but it will try to extend and simplify it.

## Goals
There are the main goals of this library:
* Implement and tests the most common functions and structures of the standard library
* The simplicity and generalization should not affect performance
* The library should be easy to use and understand

## Documentation
To make documentation more standartized and always up-to-date, we use [GoDoc (pkg.go.dev/github.com/aohorodnyk/stl)](https://pkg.go.dev/github.com/aohorodnyk/stl). It is the best documentation tool for Go.

Every feature or feature change should introduce some changes in comments, doc.go, README.md and tests include examples.

## Contributing
The library is open source and you can contribute to it.

Before contrbution, make sure that githook is configured for you and all your commits contain the correct issue tag.

### Branch Name
Before you start the contribution, make sure that you are on the correct branch. Branch name should start from the issue number dash and short explanation with spaces replaced by underscores. Example:
* `1-my_feature`
* `2-fix_bug`
* `234-my_important_pr`

### Git Hook
To configure the git hook, you need to simply run the command: `git config core.hooksPath .githooks`

It will configure the git hook to run the `pre-commit` script. Source code of the hook is in `.githooks/prepare-commit-msg`.

This git-hook will parse branch name and add the issue tag to the commit message.
