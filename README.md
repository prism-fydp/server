# server
The server subsystem for the Prism fourth-year design project.

# Setup
1. Make sure that you have Go installed, this can be done via `$ brew install go`. Note that if you install Go via brew, you may also need to install `gcc` via brew as well.
2. Run `go mod download` to install all dependencies listed in the `go.mod` file.

## Verifying setup
1. Start the server through running ```bashmake run``` (this will also build the project).
2. Open a new terminal window and run `curl localhost:8080/ping`. If the everything is running correctly, you should see an output of `{"message": "pong"}`