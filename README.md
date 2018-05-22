# Burrow

Project inspired in django and buffalo.
Burrow is a Go API First eco-system (not a framework jet). Design to focus on
writing your app without needing to reinvent the wheel.

## Installation

```bash
$ go get -u -v github.com/jorgechato/burrow/burrow
```

## Start a new Project

It's a piece of cake, this is what you should think of start an API First
project in Go. The generator is the key of success.

```bash
$ burrow startproject <name>
```

## Start an App

Modularity is what makes a pleasure to work in different projects. The
applications are single packages you can use not only in your current project
but export and use them in as many projects as you want. Type once and use it
anywhere.

```bash
$ burrow startapp <name>
```
