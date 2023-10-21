
go-safari
=========
Read-only access to Safari bookmarks, Reading List and history, plus live interaction with windows and tabs.

# Author & Contributors
Originally created by [Dean Jackson][dean] in 2016 for use with [Alfred][alfred]. Other contributors include [Troels Thomsen][tt]. [John-Mason Shackelford][jps] made minor updates after Dean archived the project in 2020.

# Usage
To install the command-line utility:
```
$ brew install go
$ go install github.com/deanishe/go-safari/cmd/safari@latest
```
The executabile will be installed in $GOBIN or in $HOME/go/bin.
```
$ $GOBIN/safari help
```
To search history use the domain of the site of interest. For example:
```
$ safari history google.com
```

# Development
Once you've cloned the repo, download dependencies and run the tests.
```
$ cd go-safari
$ go install github.com/cortesi/modd/cmd/modd@latest
$ modd
```
Check out the [modd docs][modd] for more. Tests will run automatically on touched files as you make changes.

John-Mason won't be of much use for questions as he is both a go newb and hasn't worked much with this code base either. Fork away.

# Licensing
This code is released under the [MIT License][mit].

[godoc]: https://godoc.org/pkg/github.com/deanishe/go-safari
[mit]: ./LICENCE.txt
[modd]: https://github.com/cortesi/modd
[dean]: https://github.com/deanishe
[tt]: https://github.com/tt
[jps]: https://github.com/jpshackelford
[alfred]: https://www.alfredapp.com
