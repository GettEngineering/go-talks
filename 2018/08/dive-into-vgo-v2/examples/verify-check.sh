# E.g. change `rsc.io/quote v1.5.2` hash in go.modverify from `w` to `v`

$ vgo build
vgo: verifying rsc.io/quote v1.5.2: module hash mismatch
        downloaded:   h1:w5fcysjrx7yqtD/aO+QwRjYZOKnaM9Uh2b40tElTs3Y=
        go.modverify: h1:v5fcysjrx7yqtD/aO+QwRjYZOKnaM9Uh2b40tElTs3Y=