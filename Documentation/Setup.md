# Getting Started
<!-- TOC depthfrom:2 -->

- [Requirements](#requirements)
- [Starting a new Project.](#starting-a-new-project)
- [Building the Project](#building-the-project)
- [Running the Build](#running-the-build)
- [Getting Started for Real](#getting-started-for-real)
- [Using entr as an auto-building tool](#using-entr-as-an-auto-building-tool)
        - [For Fedora Linux](#for-fedora-linux)
        - [For Debian/Ubuntu based Linux](#for-debianubuntu-based-linux)
        - [for MacOS use Homebrew](#for-macos-use-homebrew)

<!-- /TOC -->
## Requirements

to get Started your System must have access to the following commands at
minimum.

- `make`
- `go`

The Project uses `make` as its Build-System

## Starting a new Project.

1.) You can either copy the contents of this repos
[project-template](../project-template/) -Folder or download the
`project-template.tar.gz` from the
[Releases-Section](https://github.com/Rocco-Gossmann/GoWas/releases)

2.) Either way, extract the content of that folder into your project folder

3.) Run the `make setup` command inside your project folder. This will download
the required packages

## Building the Project

To build the project, simply run `make`, that is all.\
The results will be built into the `./docs`-folder inside your project folder.

## Running the Build

To run it, simply host the build, in any way you like. The template comes with a
small Server-Script, that can be used for development. It will host the files
via `http` on port `7353` of your system.

You can run that script either via.

```bash
go run ./.tools/server/server.go
```
and then opening any Browser and navigating to `http://localhost:7353` should be
fine.
Alternatively, you can use the 
```bash
make open
```
command to open the URL in your Default Browser


## Getting Started for Real

The default project template comes with some demo code, to check if the build was successful.
To clear it and make room for your own code you can run the
```bash
make reset
```
command to clear everything and revert the Project-Template to a completely blank (yet still working) state.


Now your blank slate is ready.


## Using entr as an auto-building tool 

If you have `entr` as an available shell command, you can use the,
following `make` tasks/commands in addition

```bash
#To Start the file watcher and automatically rebuild the project,
#  when files using it change
make dev
```

To install entr use:

#### For Fedora Linux
```bash
dnf install entr
```
#### For Debian/Ubuntu based Linux
```bash
apt install entr
```

#### for MacOS use Homebrew
```bash
brew install entr
```


