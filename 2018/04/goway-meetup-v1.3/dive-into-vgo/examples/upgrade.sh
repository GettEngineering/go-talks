$ vgo get golang.org/x/text
vgo: finding golang.org/x/text v0.0.0-20180208041248-4e4a3210bb54
vgo: downloading golang.org/x/text v0.0.0-20180208041248-4e4a3210bb54
$ cat go.mod
module "github.com/you/hello"

require (
        "golang.org/x/text" v0.0.0-20171214130843-f21a4dfb5e38
        "rsc.io/quote" v1.5.2
)