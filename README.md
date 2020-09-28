# GO SETUP

-   GOPATH has two segments.

    -   The first is `$PROJ/golang` which stores all downloaded third party libs
    -   The second is `$PROJ/goworkspace` which has all my go projects under `/src`

-   `$PROJ/goworkspace/src` follows the usual repository mirroring convention of `/src/github.com/zeuslawyer/<repo name>`. This is also the folder where all my packages are rooted.

# `PROJECT LOCATION`

`$PROJ/goworkspace/src/github.com/zeuslawyer/go_with_tests` . This should be `$PWD` for `git`. From there, `cd` into the relevant package/dir to run code inside them.

# `godoc` documentation server

run `godoc -http:=6060` and open up `http://localhost:6060/pkg/` to inspect the doc server. Under `Third Party` you can see `github.com/zeuslawyer/go_with_tests/hello`!
