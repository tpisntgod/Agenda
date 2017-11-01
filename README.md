# Agenda
[![Build Status](https://travis-ci.org/tpisntgod/Agenda.svg?branch=master)](https://travis-ci.org/tpisntgod/Agenda)

## Team member
- 胡子昂 15331111 
- 侯培中 15331105
- 柯永基 15331135

## 项目分工

- 胡子昂
  1.entity/user 
  2.log
- 侯培中
  1.entity/meeting
  2.travis
- 柯永基
  1.cmd

## 运行

```
go get github.com/tpisntgod/Agenda
$GOPATH/bin/Agenda
```

## Usage
```
Usage:
  agenda [command]

Available Commands:
  help        Help about any command
  register    to create a new account
  login       to sign in
  logout      to sign out
  usrSch      to list all the users
  usrDel      to delete the current user
  mc          to create a new meeting
  ap          to add some participators to a meeting
  dp          to delete some partipators from a meeting
  ms          to list all meetings you take part in according to the time you provide
  mq          to quit a meeting
  mc          to cancel a meeting
  mclr        to clear all the meetings you host

Flags:
  -h, --help   help for agenda

Use "agenda [command] --help" for more information about a command.
```
