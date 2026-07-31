<p align="center">
  <img width="400" alt="mascot" src="https://github.com/user-attachments/assets/d297eb68-8065-4016-8f27-8bd6127f7946" />
</p>
# GOS
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
```

Or if Using the binary version add:

```
alias gosdir='./YOUR_BINARY_NAME mkdir -m '
```

Then the command gets simplified to

```
gosdir FolderName
```

### Gosadd

The current way to add a new custom template is

```
gosadd -a "FolderName"
```

But because it feels so bad to type that big command we'll use this (.bashrc)

```
alias gosadd='gos gosadd -a ' --Space is important here too 
```

Or if your using the binary version

```
alias gosdir='./YOUR_BINARY_NAME gosadd -a '
```

Then the command gets simplified to

```
goadd projectName
```

### Gosget

The current way to import an saved user template is
Here:-

- `projectName` is the name you wanna give the new folder
- `templateName` is the templates name that you have saved in `.gos`

```
gos gosget -p projectName -e templateName 
```

But he recommend way is to add this to your (.bashrc)

```
alias gosget='gos gosget'
```

or if using the binary version

```
alias gosget='./YOUR_BINARY_NAME gosget'
```

Then the command gets simplified to

```
gosget -p projectName -e templateName
```

# Contribute

```bash
git clone https://github.com/lxsh-S/gos.git
cd gos
go run .
```
