# Journal
## Day 1: gowc, a wc clone in Go (Nov 1-3)
**Built:** a CLI that counts lines, words and bytes from files or stdin, with -l/-w/-c flags,
a total line for multiple files, errors on stderr and exit code 1 on failure. Table-driven tests pass.
**Broke:** `go run` said "cannot find main module" because I ran it from the repo root instead of
the folder with go.mod. Also the missing-file error named the file twice because os.Open already
includes it and I wrapped it again.
**Learned:** io.Reader lets the same Count function read a string, a file or stdin. stderr (stream 2)
is separate from stdout, so `2>/dev/null` hides errors but the exit code (`echo $?`) still shows failure.
A method needs a pointer receiver if it changes the value.
**Next / unsure about:** how Go decides when an error should be wrapped with %w.