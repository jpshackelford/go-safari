
go-safari
=========
Read-only access to Safari bookmarks, Reading List and history, plus live interaction with windows and tabs.

# Author & Contributors
Originally created by [Dean Jackson][dean] in 2016 for use with [Alfred][alfred]. Other contributors include [Troels Thomsen][tt]. [John-Mason Shackelford][jps] made minor updates after Dean archived the project in 2020.

# Usage
To install the command-line utility:
```console
brew install go
go install github.com/jpshackelford/go-safari/cmd/safari@latest
```
The executabile will be installed in $GOBIN or in $HOME/go/bin.
```console
$GOBIN/safari help
```
To search history use the domain of the site of interest. For example:
```console
safari history google.com
```

# Development
Once you've cloned the repo, download dependencies and run the tests.
```console
cd go-safari
go install github.com/cortesi/modd/cmd/modd@latest
modd
```
Check out the [modd docs][modd] for more. Tests will run automatically on touched files as you make changes.

To install the development version
```console
go build ./cmd/safari && go install ./cmd/safari
```

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
