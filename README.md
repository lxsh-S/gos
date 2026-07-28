# GOS

<img width="1267" height="762" alt="gos" src="https://github.com/user-attachments/assets/18cc8a38-405a-4f5d-a5e8-f37c65b79613" />

*gos* is a simple CLI tool that help to make a project structure fast and quick!

# Design choices and goals

Gos is not intended to compete with any other project structuring tool it's a fun hobby project made because of my interest in CLI tools.

# Quick Start

### Golang

```
gos Make-project -l go -t std   --Creates the standard go project structure 

gos Make-project -l go -t api     --Creates a go boilerplate structure for API projects 

gos Make-project -l go -t cli     --Creates a go boilerplate structure for cli projects <lxsh-s approved :D>
```

### C++

```
gos Make-project -l cpp -t std    --Creates the standard cpp project strucutre 

gos Make-project -l cpp -t app    --Creates a cpp boilerplate structure for app/game

gos Make-project -l cpp -t lib    --Creates a cpp boilerplate structure for a reusable library
```

### Typescript

```
gos Make-project -l ts -t std   --Creates the standard Typescript project strucutre

gos Make-project -l ts -t lib   --Creates a Typescript boilerplate for a library

gos Make-project -l ts -t api   --Creates a Typescript boilerplate for a API projects

gos Make-project -l ts -t nxtjs --Creates a standard Typescript biolerplate for Next.js projects
```

### List `Lang and types`

```
gos --list 
```

# Installation

### AUR (Arch User Repository)

```
yay -S gos-bin 
```

### Using Go

```bash
go install github.com/lxsh-S/gos@latest 
```

# Afer Installation

### Gosdir

The current way to make a new empty directory is

```
gos mkdir -m "FolderName"
```

But I recommend to use an alias for making the new directory.
Add this to your shell (The one given below is same that I use i.e .bashrc)

```
alias gosdir='gos mkdir -m ' --Note that space after "-m" is important
or if Using the binary version
alias gosdir='./YOUR_BINARY_NAME mkdir -m '
```

Note:- Only works for versions after `0.8.5` i.e from and beyond version `0.9.0`

# Contribute

```bash
git clone https://github.com/lxsh-S/gos.git
cd gos
go run .
```
