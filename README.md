# Recursive Touch (rtouch)
`touch` with `mkdir` functionality.

## Table of Contents:
1. [Why?](#why)
2. [Installation](#installation)
3. [Usage](#usage)
4. [Contribute](#contribute)
5. [Todo](#todo)

## Why?
I use the `touch` command multiple times a day. It's a simple command with only one purpose: to create a new file. The main 'problem' (more of a personal gripe) with this command is that it *only creates a new file*. If you give it a folder that doesn't exist, it will return an error. This means that I have to first run `mkdir` and only then can I use `touch`. I understand that this is by design. GNU/Linux is built around the "do one thing only and do it well" principle (which has been gradually falling apart with the rise of tools such as SystemD), but I personally don't necessarily have to adhere to that rule.<br>


So yeah, use `rtouch` instead of `touch` if you want to skip the `mkdir` command. That's all there is to it.

## Installation:
### Via releases:
- Grab your system's binary from <a href="https://github.com/bjvanbemmel/rtouch/releases/latest">the releases page</a>.
- Move the binary to your path. Linux example:
```bash
sudo chmod +x rtouch_linux_amd64
mv rtouch_linux_amd64 ~/.local/bin/rtouch
```
### From source:
- Clone the repository.
- Move into the newly added directory.
- Execute ``make install`` (or ``make build``).
- (optional) Move newly created executable to your path.
```bash
git clone https://github.com/bjvanbemmel/rtouch.git
cd rtouch
# If ~/.local/bin is within your path.
make install
# Else
make build
mv ./bin/rtouch_linux_amd64 /usr/bin/rtouch
```
*``make install`` will move the executable into the ~/.local/bin folder. If this folder is not part of your path, please either change it or use ``make build`` and manually move it to a folder that's within your path.*

## Usage:
Help / usage:
```
rtouch -h
rtouch --help
```
Create new file:
```
rtouch main.go
```
Create new file within new folder:
```
rtouch project/main.go
rtouch ~/workspace/project/src/main.go
rtouch project/main.go project/user.go
```

## Contribute:
`rtouch` makes use of the MIT license. You're free to create a pull request and/or fork the project. Just make sure to keep the MIT license and copyright notice.

## Todo:
- ~~Allow for multiple filepath args (can only create one file per command right now)~~
- Include ability to assign permissions

#
##### &copy; Beau Jean van Bemmel, 2022