# gophercises
go practice

* Main file in `project/cmd/executable_name/main.go`.
* to run: `go run path_to_main/main.go args`.
* to build: `go build` in main dir. Will make executable in place.
* to install `GOBIN=/usr/local/bin/ go install`. Will build and move to specified dir. `go/bin` is the default. Or `go build -o /path/binary-name`. If installing to `go/bin` make sure that `/Users/username/go/bin` added to `$PATH`. 
* to test `go test -v -run TestName`

