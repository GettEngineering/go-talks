$ vgo test all
?       github.com/you/hello    [no test files]
?       golang.org/x/text/internal/gen  [no test files]
ok      golang.org/x/text/internal/tag  0.014s
?       golang.org/x/text/internal/testtext     [no test files]
ok      golang.org/x/text/internal/ucd  0.013s
ok      golang.org/x/text/language      0.050s
ok      golang.org/x/text/unicode/cldr  0.059s
--- FAIL: TestHello (0.00s)
        quote_test.go:19: Hello() = "99 bottles of beer on the wall, 99 bottles of beer, ...", want "Hello, world."
FAIL
FAIL    rsc.io/quote    0.011s
--- FAIL: TestHello (0.00s)
        hello_test.go:31: Hello([en-US fr]) = "99 bottles of beer on the wall, 99 bottles of beer, ...", want "Hello, world."
        hello_test.go:31: Hello([fr en-US]) = "99 bottles of beer on the wall, 99 bottles of beer, ...", want "Bonjour le monde."
FAIL
FAIL    rsc.io/sampler  0.009s