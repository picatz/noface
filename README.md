# noface

An almost practical network interface tool for the command-line.

![noface](https://i.imgur.com/YTnFXRg.gif)

### What does it do?

`noface` will simply return the first network interface if it has an ip address.

##### Example

If the network interface has an ip address ( connected to the network ).
```shell
$ noface
en0
```

Otherwise, if things don't go as planned, `noface` will resond with a failed exit code.
```shell
$ noface # no output
$ echo $?
1
```

### I named my network interface "noface"

Plz don't do that. 

Also, the `noface` command-line tool with return a failed exit code which can be checked in a shell script, for example.

## Installation

```shell
# go get noface
$ go get github.com/picatz/noface
# or use git
$ git clone https://github.com/picatz/noface.git

# change into the directory where noface has been downloaded
$ cd noface

# get glide for dependencies
$ make get_glide

# install dependencies
$ make install_dependencies

# build osx binary
$ make build_osx

# build linux binary
$ make build_linux
```
