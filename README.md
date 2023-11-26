<div align="center">
  <br />
  <img src="assets/logo.png" alt="GitGitGo-CLILogo" width="30%"/>
  <h1>GitGitGo-CLI</h1>
  <p>
     GitGitGo-CLI is a CLI tool to help you manage your git repositories.
  </p>
</div>

<!-- Badges -->
<div align="center">
   <a href="https://github.com/phil1436/GitGitGo-CLI/releases">
       <img src="https://img.shields.io/github/v/release/phil1436/GitGitGo-CLI?display_name=tag" alt="current realease" />
   </a>
   <a href="https://github.com/phil1436/GitGitGo-CLI/blob/main/LICENSE">
       <img src="https://img.shields.io/github/license/phil1436/GitGitGo-CLI" alt="license" />
   </a>
   <a href="https://github.com/phil1436/GitGitGo-CLI/stargazers">
       <img src="https://img.shields.io/github/stars/phil1436/GitGitGo-CLI" alt="stars" />
   </a>
   <a href="https://github.com/phil1436/GitGitGo-CLI/commits/main">
       <img src="https://img.shields.io/github/last-commit/phil1436/GitGitGo-CLI" alt="last commit" />
   </a>
</div>

---

-   [Concept](#concept)
-   [Installation](#installation)
-   [Usage](#usage)
    -   [Commands](#commands)
    -   [Global flags](#global-flags)
    -   [Parameters](#parameters)
-   [Examples](#examples)
-   [Bugs](#bugs)
-   [Release Notes](#release-notes)

---

## Concept

Working with git is essential for every developer. But sometimes keeping track of all files you need in your repository can be a hassle. GitGitGo-CLI is a CLI tool to help you manage your git repositories. It allows you to create a repository with a predefined structure and to add predefined files to it dynamically. This way you can create a repository with all the files you need in a matter of seconds and start working on your project right away!

### .gitgitgo

To store your templates you need to create a `.gitgitgo` repository (use [this](https://github.com/phil1436/.gitgitgo) template repository). This repository will contain all your templates in a json file. Click [here](https://github.com/phil1436/.gitgitgo) to get more information on how to create your own `.gitgitgo` repository.

### Provider

GitGitGo-CLI needs a provider (basically a GitHub user- / organisation name) to fetch your templates from. That can be useful when you need different file templates in different contexts (e.g. work and private).

---

## Installation

### Binaries

1. Download the binary for your OS from the [latest release](https://github.com/phil1436/GitGitGo-CLI/releases/latest)
2. Add the binary to your PATH
3. Run `gitgitgo` in your terminal to check if it works

### From source

1. Clone the repository
2. Run `go build` in the root directory
3. Add the binary to your PATH
4. Run `gitgitgo` in your terminal to check if it works

### .gitgitgo

Follow the instructions [here](https://github.com/phil1436/.gitgitgo) to create your own `.gitgitgo` repository.

---

## Usage

### Commands

#### `version`

```bash
gitgitgo version
```

Prints the current version of GitGitGo-CLI.

#### `help`

```bash
gitgitgo help <command>
```

Prints the help page for the given command. If no command is given, the help page for the `gitgitgo` command will be printed.

#### `init`

```bash
gitgitgo init <keywords> <flags>
```

Initializes a new git repository with a predefined structure and files.

You can add a comma separated list of keywords to the command. The keywords will be used to filter the templates in the `.gitgitgo` repository. Only templates that contain all keywords will be used. If no keywords are given, all templates that have the property `oninit` set to `true` will be used.

Flags:

-   `[-destination|-d] <path>`: The path where to initialize the repository. Defaults to the current working directory.
-   `-force`: Forces the initialization of the repository. If the files already exist, they will be overwritten.
-   `[-dry-run|-dryrun|-dr]`: If set, the repository will not be initialized. Instead, the files that would be created will be printed to the console.
-   `[-ignore|-i] <file1,file2, ... >`: A comma separated list of files to ignore. If a file is ignored, it will not be created in the repository.

#### `add`

```bash
gitgitgo add <filename> <flags>
```

Adds the given file to the current git repository.

Flags:

-   `[-destination|-d] <path>`: The path where to initialize the repository. Defaults to the current working directory.
-   `-as`: The name of the file in the repository. Defaults to the name of the file.
-   `-force`: Forces the initialization of the repository. If the files already exist, they will be overwritten.
-   `[-dry-run|-dryrun|-dr]`: If set, the repository will not be initialized. Instead, the files that would be created will be printed to the console.

#### `print`

```bash
gitgitgo print <filename | parametername> <flags>
```

If a [parameter](#parameters) name is given, prints the value of the parameter. If a filename is given, prints the content of the file. When no parameter or filename is given, prints all files in the current `.gitgitgo` repository.

Flags:

-   `[-name|-n]`: Prints the name of the files instead of the content.

#### `shell`

```bash
gitgitgo shell
```

Start a GitGitGo-CLI shell. This allows you to run multiple commands without having to restart the CLI every time. This commands adds more commands to navigate the shell. Please run `help` in the shell to get more information.

#### `run`

```bash
gitgitgo run <filename>
```

Runs the given file. This is useful if you want to run a script that is stored in your `.gitgitgo` repository. Take a closer look at the [examples section](#examples) to see how this can be used.

Flags:

-   `-force`: If enabled the execution will continue even if an error occurs

#### `set`

```bash
gitgitgo set <name> <value>
```

Sets a [parameter](#parameters) to the given value.

### Global flags

This flags can be used with every command:

-   `[-quiet|-q]` : If enabled no output is printed
-   `[-provider|-p] <name>` : Change the provider of the .gitgitgo repository
-   `[-githubname|-gname|-gn] <name>` : Change the github name that will be used
-   `[-fullname|-fname|-fn] <name>` : Change the full name that will be used
-   `[-reponame|-rname|-rn] <name>` : Change the repo name that will be used
-   `-dev` : Run in dev mode [internal use only]

### Parameters

#### `provider`

The provider is the name of the user or organisation that owns the `.gitgitgo` repository.

> You can use `p` as a shortcut for `provider`.

#### `githubname`

The github name is the name that will be used in the file templates. The value will be used for the `${GITHUBNAME}` placeholder in the templates.

> You can use `gname` or `gn` as a shortcut for `githubname`.

#### `fullname`

The full name is the name that will be used in the file templates. The value will be used for the `${FULLNAME}` placeholder in the templates.

> You can use `fname` or `fn` as a shortcut for `fullname`.

#### `reponame`

The repo name is the name that will be used in the file templates. The value will be used for the `${REPONAME}` placeholder in the templates. The default is the name of the current working directory.

> You can use `rname` or `rn` as a shortcut for `reponame`.

---

## Config file

GitGitGo-CLI can use a config file to store the current provider and the current parameters. The config file is located in the home directory of the current user. The name of the file is **`.gitgitgoc`**. The file should has the following structure:

```
# This is a comment

# Set the provider
provider=phil1436

# Set your full name
fullname=Philipp B.

# Set your github name
githubname=phil1436
```

---

## Examples

View and run the files in the [examples](examples/) directory to get a better understanding of how GitGitGo-CLI works.

---

## Bugs

-   _no known bugs_

---

## [Release Notes](https://github.com/phil1436/GitGitGo-CLI/blob/main/CHANGELOG.md)

### [v0.0.2](https://github.com/phil1436/GitGitGo-CLI/tree/0.0.2)

-   Bug fixes
-   Description and Keywords support for `print` command
-   `.gitgitgoc` file support
-   keyword support for `init` command

---

by [Philipp B.](https://github.com/phil1436)
