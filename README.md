# pargs - cross-platform xargs-a-like with argv array batching

# EXAMPLES

```
$ find shakespeare -type f -print | pargs -pool 2 echo 'plays:'
plays: shakespeare/comedies/allswellthatendswell shakespeare/comedies/asyoulikeit
plays: shakespeare/comedies/comedyoferrors shakespeare/comedies/cymbeline
plays: shakespeare/comedies/loveslabourslost shakespeare/comedies/measureforemeasure
plays: shakespeare/comedies/merchantofvenice shakespeare/comedies/merrywivesofwindsor
plays: shakespeare/comedies/midsummersnightsdream shakespeare/comedies/muchadoaboutnothing
plays: shakespeare/comedies/periclesprinceoftyre shakespeare/comedies/tamingoftheshrew
plays: shakespeare/comedies/tempest shakespeare/comedies/troilusandcressida
plays: shakespeare/comedies/twelfthnight shakespeare/comedies/twogentlemenofverona
plays: shakespeare/comedies/winterstale shakespeare/glossary
plays: shakespeare/histories/1kinghenryiv shakespeare/histories/1kinghenryvi
plays: shakespeare/histories/2kinghenryiv shakespeare/histories/2kinghenryvi
plays: shakespeare/histories/3kinghenryvi shakespeare/histories/kinghenryv
plays: shakespeare/histories/kinghenryviii shakespeare/histories/kingjohn
plays: shakespeare/histories/kingrichardii shakespeare/histories/kingrichardiii
plays: shakespeare/poetry/loverscomplaint shakespeare/poetry/rapeoflucrece
plays: shakespeare/poetry/sonnets shakespeare/poetry/various
plays: shakespeare/poetry/venusandadonis shakespeare/README
plays: shakespeare/tragedies/antonyandcleopatra shakespeare/tragedies/coriolanus
plays: shakespeare/tragedies/hamlet shakespeare/tragedies/juliuscaesar
plays: shakespeare/tragedies/kinglear shakespeare/tragedies/macbeth
plays: shakespeare/tragedies/othello shakespeare/tragedies/romeoandjuliet
plays: shakespeare/tragedies/timonofathens shakespeare/tragedies/titusandronicus

$ find shakespeare -type f -print | pargs -pool 10 echo 'plays:'
plays: shakespeare/comedies/allswellthatendswell shakespeare/comedies/asyoulikeit shakespeare/comedies/comedyoferrors shakespeare/comedies/cymbeline shakespeare/comedies/loveslabourslost shakespeare/comedies/measureforemeasure shakespeare/comedies/merchantofvenice shakespeare/comedies/merrywivesofwindsor shakespeare/comedies/midsummersnightsdream shakespeare/comedies/muchadoaboutnothing
plays: shakespeare/comedies/periclesprinceoftyre shakespeare/comedies/tamingoftheshrew shakespeare/comedies/tempest shakespeare/comedies/troilusandcressida shakespeare/comedies/twelfthnight shakespeare/comedies/twogentlemenofverona shakespeare/comedies/winterstale shakespeare/glossary shakespeare/histories/1kinghenryiv shakespeare/histories/1kinghenryvi
plays: shakespeare/histories/2kinghenryiv shakespeare/histories/2kinghenryvi shakespeare/histories/3kinghenryvi shakespeare/histories/kinghenryv shakespeare/histories/kinghenryviii shakespeare/histories/kingjohn shakespeare/histories/kingrichardii shakespeare/histories/kingrichardiii shakespeare/poetry/loverscomplaint shakespeare/poetry/rapeoflucrece
plays: shakespeare/poetry/sonnets shakespeare/poetry/various shakespeare/poetry/venusandadonis shakespeare/README shakespeare/tragedies/antonyandcleopatra shakespeare/tragedies/coriolanus shakespeare/tragedies/hamlet shakespeare/tragedies/juliuscaesar shakespeare/tragedies/kinglear shakespeare/tragedies/macbeth
plays: shakespeare/tragedies/othello shakespeare/tragedies/romeoandjuliet shakespeare/tragedies/timonofathens shakespeare/tragedies/titusandronicus

$ find shakespeare -type f -print | pargs -pool 100 echo 'plays:'
plays: shakespeare/comedies/allswellthatendswell shakespeare/comedies/asyoulikeit shakespeare/comedies/comedyoferrors shakespeare/comedies/cymbeline shakespeare/comedies/loveslabourslost shakespeare/comedies/measureforemeasure shakespeare/comedies/merchantofvenice shakespeare/comedies/merrywivesofwindsor shakespeare/comedies/midsummersnightsdream shakespeare/comedies/muchadoaboutnothing shakespeare/comedies/periclesprinceoftyre shakespeare/comedies/tamingoftheshrew shakespeare/comedies/tempest shakespeare/comedies/troilusandcressida shakespeare/comedies/twelfthnight shakespeare/comedies/twogentlemenofverona shakespeare/comedies/winterstale shakespeare/glossary shakespeare/histories/1kinghenryiv shakespeare/histories/1kinghenryvi shakespeare/histories/2kinghenryiv shakespeare/histories/2kinghenryvi shakespeare/histories/3kinghenryvi shakespeare/histories/kinghenryv shakespeare/histories/kinghenryviii shakespeare/histories/kingjohn shakespeare/histories/kingrichardii shakespeare/histories/kingrichardiii shakespeare/poetry/loverscomplaint shakespeare/poetry/rapeoflucrece shakespeare/poetry/sonnets shakespeare/poetry/various shakespeare/poetry/venusandadonis shakespeare/README shakespeare/tragedies/antonyandcleopatra shakespeare/tragedies/coriolanus shakespeare/tragedies/hamlet shakespeare/tragedies/juliuscaesar shakespeare/tragedies/kinglear shakespeare/tragedies/macbeth shakespeare/tragedies/othello shakespeare/tragedies/romeoandjuliet shakespeare/tragedies/timonofathens shakespeare/tragedies/titusandronicus

$ pargs -help
  -help
        Show usage information
  -pool int
        How many arguments to supply at once. Minimum 1. (default 1000)
  -version
        Show version information
```

# DOWNLOAD

https://github.com/mcandre/pargs/releases

# DOCUMENTATION

https://godoc.org/github.com/mcandre/pargs

# REQUIREMENTS

* [Go](https://golang.org) 1.7+ with [$GOPATH configured](https://gist.github.com/mcandre/ef73fb77a825bd153b7836ddbd9a6ddc)

## Optional

* [coreutils](https://www.gnu.org/software/coreutils/coreutils.html)
* [make](https://www.gnu.org/software/make/)
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) (e.g. `go get golang.org/x/tools/cmd/goimports`)
* [golint](https://github.com/golang/lint) (e.g. `go get github.com/golang/lint/golint`)
* [errcheck](https://github.com/kisielk/errcheck) (e.g. `go get github.com/kisielk/errcheck`)
* [gox](https://github.com/mitchellh/gox) (e.g. `go get github.com/mitchellh/gox`)
* [editorconfig-cli](https://github.com/amyboyd/editorconfig-cli) (e.g. `go get github.com/amyboyd/editorconfig-cli`)
* [flcl](https://github.com/mcandre/flcl) (e.g. `go get github.com/mcandre/flcl/...`)
* [shlint](https://rubygems.org/gems/shlint)
* [shellcheck](http://hackage.haskell.org/package/ShellCheck)

# INSTALL FROM REMOTE GIT REPOSITORY

```
$ go get github.com/mcandre/pargs/...
```

(Yes, include the ellipsis as well, it's the magic Go syntax for downloading, building, and installing all components of a package, including any libraries and command line tools.)

# INSTALL FROM LOCAL GIT REPOSITORY

```
$ mkdir -p $GOPATH/src/github.com/mcandre
$ git clone git@github.com:mcandre/pargs.git $GOPATH/src/github.com/mcandre/pargs
$ cd $GOPATH/src/github.com/mcandre/pargs
$ git submodule update --init --recursive
$ sh -c 'cd cmd/pargs && go install'
```

# TEST

```
$ make integration-test
```

# PORT

```
$ make port
```

# LINT

Keep the code tidy:

```
$ make lint
```

# GIT HOOKS

See `hooks/`.

# CREDITS

Shakespeare examples from [shakesfiles](http://www.compciv.org/practicum/shakefiles/b-downloading-the-shakespeare-zip/)
