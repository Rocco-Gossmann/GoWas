# GoWas - a (GO) + (W)eb(as)sembly Engine
A Go and WASM based Engine to build small applications heavily focused on pixel based screen manipulation.  
It is heavily inspired by Engines like Pico-8 and the OneLoneCoders PixelGameEngine
(https://github.com/OneLoneCoder/olcPixelGameEngine)



# How to set it up.
In this project, you'll find the " **project-template** " folder.
It contains everything to get started with your own project.

1. Copy it to your workspace
2. run 
```bash
# To Update and download a all required dependencys
go mod tidy

# To Compile the project
make

# To remove all build files and build everything from scratch
make clean
```

If you have `ZSH` and `entr` as an available shell commands, you can use the,
following commands in addition 

```bash
#To start the DevServer it comes with
make run

#To Start the File Watcher and automatically rebuild the project,
#  when files using it change 
make dev
```

The Build be be in the `./docs` Folder

