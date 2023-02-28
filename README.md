## jwtauth

Minimal cli to interact with JWT Auth Tokens.

### Usage

_Prerequisites: Install Golang: https://go.dev/doc/install_

Install as a global command:

```bash
go install -v github.com/refcell/jwtauth@latest
```

Then, you can run `jwtauth` from anywhere.

### Reference

```bash
NAME:
   jwtauth - Minimal cli to interact with JWT Auth Tokens

USAGE:
   jwtauth [global options] command [command options] [arguments...]

COMMANDS:
   claims, c         generate claims for a JWT token
   header, h         generate the http auth bearer header for a JWT token
   generate, g, gen  generate a new JWT token
   help, h           Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

### Example

Create a hex encoded jwt token:

```bash
$ jwtauth g
0x66c01996d15563ec70b57ae35e22b739125094d40e5a6853746912b43ae56ce0
```

Generate the claims for the token that can be used in the `Authorization` header:

```bash
$ jwtauth c 0x66c01996d15563ec70b57ae35e22b739125094d40e5a6853746912b43ae56ce0
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2Nzc1NjQwMTh9.5dQj9FjSZ5KWR61ZISBvWMlIj9f79UqrdQeMGPrY9Ns
```

### Building from source

To build from source, run:

```bash
go build -o jwtauth
```


### License

[MIT](https://choosealicense.com/licenses/mit/)
