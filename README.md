userdir
=======

[![Build Status](https://travis-ci.org/vrischmann/userdir.svg?branch=master)](https://travis-ci.org/vrischmann/userdir)
[![GoDoc](https://godoc.org/github.com/vrischmann/userdir?status.svg)](https://godoc.org/github.com/vrischmann/userdir)

userdir is a small library which allows you to get user directories according to the operating system conventions.

Simple example:

```go
d := userdir.GetDataHome()
fmt.Println(d)
// Outputs something like this
//   /home/vincent/.local/share        - on Linux
//   C:/Users/vincent/AppData/Roaming  - on Windows
```

Windows specificities
---------------------

On Windows *userdir* will use the operating systems function [SHGetKnownFolderPath](https://msdn.microsoft.com/en-us/library/windows/desktop/bb762188(v=vs.85).aspx).

Linux specificities
-------------------

On Linux *userdir* will use the [XDG Base Directory Specification](http://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html).

Further support
---------------

Still needs support for other Unixes (might be as simple as using XDG there too) and Mac OS X.
