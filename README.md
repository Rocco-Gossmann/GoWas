# GoWas - a (GO) + (W)eb(as)sembly Engine
A Go and WASM based Engine.  
It is heavily inspired by Engines like Pico-8 and the OneLoneCoders PixelGameEngine
As well as various GameConsoles of the 80s and 90s
(https://github.com/OneLoneCoder/olcPixelGameEngine)


# How to set it up.
In this project, you'll find the " **project-template** " folder.
It contains everything to get started with your own project.

1. Copy it to your workspace

2. Run 
```bash
make setup
```


# Your first Build
The Project template is working out of the box. It doubles as a Demo in that regard.

To create a build in the templates `./docs` directory, run:
```bash
make
```
You can test your project via the `server/server.go` file, that the template comes with.
```bash
go run ./server/server.go
```
> [!Warning] 
> This server should not be used in any kind of production build ever. It only 
> exists for testing and development purposes.


Then open [http://localhost:7353](http://localhost:7353)  and you should see a 
Red square that can be controlled by your Mouse.

Once you confirmed, that everything works,  
use the following if you want to delete all build files
```bash
# To remove all build files and build everything from scratch
make clean
```


# Get Started for real
Before getting started, make sure to change the go.mod files `module` and `go` Version
to the onse your project needs. Otherwise everything is the same as for *Your first build*


# Additional Development tools
If you have `ZSH` and `entr` as an available shell commands, you can use the,
following commands in addition 

```bash
#To start the DevServer it comes with
make run

#To Start the File Watcher and automatically rebuild the project,
#  when files using it change 
make dev

#To Stop the Server started via ` make run `
make stop
```

The Build be be in the `./docs` Folder

