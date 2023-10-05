# Getting Started

## Requirements

to get Started your System must have access to the following commands at minimum.
- `make`
- `go`

The Project uses `make` as its Build-System


## Starting a new Project.

1.) You can either copy the contents of this repos [project-template](../project-template/) -Folder
or download the `project-template.tar.gz` from the [Releases-Section](https://github.com/Rocco-Gossmann/GoWas/releases)

2.) Either way, extract the content of that folder into your own project-folder

3.) Run the `make setup` command inside your project-folder. 
This will download the required packages


## Building the Project

To build the project, simply run `make`, that is all.  
The results will be build to the `./docs`-folder inside your project-folder.


## Running the Build

To run it, simply host the build, in any way you like. 
The template comes with a small Server-Script, that can be used for development.
It will host the files via `http` on port `7353` of your system.

So running that script via.
`go run ./server/server.go`

and then opening any Browser and navigating to `http://localhost:7353` should be fine.


## Getting started for Real
remove all the `*.go` files from the project rename the package inside `go.mod`
remove `go.work` and run `make setup` again.

Now your blank slate is ready.

> [!todo] 
> Create a `make` task and tool-script to automate all of this. 


## Optional build Tools

If you have `zsh` and `entr` as available shell commands, you can use the,
following `make` tasks/commands in addition 

```bash
#To start the DevServer it comes with
make run

#To Start the File Watcher and automatically rebuild the project,
#  when files using it change 
make dev

#To Stop the Server started via ` make run `
make stop
```


