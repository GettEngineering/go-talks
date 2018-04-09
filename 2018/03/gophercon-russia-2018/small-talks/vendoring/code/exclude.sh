$ echo 'exclude "rsc.io/sampler" v1.99.99' >>go.mod
$ vgo list -t rsc.io/sampler
rsc.io/sampler
	v1.0.0
	v1.2.0
	v1.2.1
	v1.3.0
	v1.3.1
	v1.99.99 # excluded
$ vgo get -u
vgo: finding golang.org/x/text latest
vgo: finding rsc.io/quote latest
vgo: finding rsc.io/sampler latest
vgo: finding rsc.io/sampler latest
vgo: finding golang.org/x/text latest
$ vgo list -m
MODULE                VERSION
github.com/you/hello  -
golang.org/x/text     v0.0.0-20171214130843-f21a4dfb5e38
rsc.io/quote          v1.5.2
rsc.io/sampler        v1.3.1