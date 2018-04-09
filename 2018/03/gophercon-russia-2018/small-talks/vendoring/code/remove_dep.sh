$ vgo get rsc.io/sampler@none
vgo: downloading rsc.io/quote v1.4.0
vgo: finding rsc.io/quote v1.3.0
$ vgo list -m
MODULE                VERSION
github.com/you/hello  -
golang.org/x/text     v0.0.0-20180208041248-4e4a3210bb54
rsc.io/quote          v1.3.0
$ cat go.mod
module "github.com/you/hello"

require (
	"golang.org/x/text" v0.0.0-20180208041248-4e4a3210bb54
	"rsc.io/quote" v1.3.0
)
$ vgo test all
vgo: downloading rsc.io/quote v1.3.0
?   	github.com/you/hello	[no test files]
ok  	rsc.io/quote	0.014s