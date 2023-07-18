# arch-sample-job
Sample application for building backend jobs

## Todo
# How to build & extend the sample CLI application

How this was made
- go mod init github.com/chef/arch-sample-cli
- go get -u github.com/spf13/cobra@latest
- go install github.com/spf13/cobra-cli@latest
- go mod tidy
- cobra-cli init (creates the cmd folder and blows away main.go)
- go get github.com/spf13/viper@v1.8.1 (add viper ref to root.go)
- (OPTIONAL) add files for license, folders for bin, coverage, dev-docs, docs-chef.io

At this point add your own stuff
- clone the repo
- go install github.com/spf13/cobra-cli@latest
- test install with cobra-cli
- for each command line flag, cobra-cli add flag (adds flag.go to cmd)

Then
- go mod vendor (creates vendor subdirectory files)
- copy old logic to right spots, like main()
- go vet (check for folder, dep, syntax issues)

To build (command-line terminal in VSCode)
- go build -o bin/arch-sample-cli .
- go run .\main.go -- or -- ./bin/arch-sample-cli 
- to do a flag, go run . get (get called)

Or you can use the github workflows

Go Doc
- go install -v golang.org/x/tools/cmd/godoc@latest
- ~/go/bin/godoc -http=:6060 (start the server)
- browse to https://localhost:6060
