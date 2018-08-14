$ echo 'replace "rsc.io/quote" v1.5.2 => "../quote"' >>go.mod
$ vgo list -m
MODULE                VERSION
github.com/you/hello  -
golang.org/x/text     v0.0.0-20180208041248-4e4a3210bb54
rsc.io/quote          v1.5.2
 => ../quote
rsc.io/sampler        v1.3.1
$ vgo build