$ vgo get -u
vgo: finding golang.org/x/text latest
vgo: finding rsc.io/quote latest
vgo: finding rsc.io/sampler latest
vgo: finding rsc.io/sampler v1.99.99
vgo: finding golang.org/x/text latest
vgo: downloading rsc.io/sampler v1.99.99
$ cat go.mod
module "github.com/you/hello"

require (
        "golang.org/x/text" v0.0.0-20171214130843-f21a4dfb5e38
        "rsc.io/quote" v1.5.2
        "rsc.io/sampler" v1.99.99
)