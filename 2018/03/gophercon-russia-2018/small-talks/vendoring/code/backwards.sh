$ vgo vendor
$ ls vendor
golang.org rsc.io     vgo.list
$ cat vendor/vgo.list
MODULE                VERSION
github.com/you/hello  -
golang.org/x/text     v0.0.0-20171214130843-f21a4dfb5e38
rsc.io/quote          v1.5.2
rsc.io/sampler        v1.3.1
$ mkdir -p $GOPATH/src/github.com/you
$ cp -a . $GOPATH/src/github.com/you/hello
$ go build -o vhello github.com/you/hello
$ LANG=es ./vhello
Hola Mundo.
$ go tool nm hello | grep sampler.hello
 116c900 B rsc.io/sampler.hello
$ go tool nm vhello | grep sampler.hello
 11718e8 B github.com/you/hello/vendor/rsc.io/sampler.hello