# Recursive Touch (rtouch)
`touch` with `mkdir` functionality.

## Why?
I use the `touch` command multiple times a day. It's a simple command with only one purpose: to create a new file. The main 'problem' (more of a personal gripe) with this command is that it *only creates a new file*. If you give it a folder that doesn't exist, it will return an error. This means that I have to first run `mkdir` and only then can I use `touch`. I understand that this is by design. GNU/Linux is built around the "do one thing only and do it well" principle (which has been gradually falling apart with the rise of tools such as SystemD), but I personally don't necessarily have to adhere to that rule.<br>


So yeah, use `rtouch` instead of `touch` if you want to skip the `mkdir` command. That's all there is to it.

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
```
`rtouch` will ask you for confirmation when creating new folders.
```
New folder(s) will have to be created. Do you consent? [y/N]
> y
```
Skip confirmation prompt:
```
rtouch -f project/main.go
rtouch --force project/main.go
```

## Todo:
- Allow for multiple filepath args (can only create one file per command right now)
- Include ability to assign permissions

#
##### &copy; Beau Jean van Bemmel, 2022