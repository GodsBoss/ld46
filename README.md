Your Inner Child (Keep It Alive!)
=================================

Simple tower defense game made as [an entry](https://ldjam.com/events/ludum-dare/46/your-inner-child-keep-it-alive) for [Ludum Dare](https://ldjam.com/) [46](https://ldjam.com/events/ludum-dare/46).

Written in [Go](https://golang.org/) and compiled to [WebAssembly](https://webassembly.org/) to run in the browser.

`make` commands
---------------

### Build

    make

Builds everything into the `./dist` folder.

### Clean

    make clean

Cleanup.

### Unit tests

    make test

Runs unit tests (there is only a few).

### Dev server

    make serve

Run a minimal HTTP server running the game. Defaults to listen at `127.0.0.1:8080`, this can be changed via the environment variable `LISTEN_ADDR`.
