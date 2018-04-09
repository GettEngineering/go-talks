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