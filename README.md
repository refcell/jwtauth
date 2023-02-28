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

Create a hex encoded 256 bit jwt secret string (with a `0x` prefix):

```bash
$ jwtauth g 0x66c01996d15563ec70b57ae35e22b739125094d40e5a6853746912b43ae56ce0
```

Generate the claims for the token that can be used in the `Authorization` header:

```bash
$ jwtauth c 0x66c01996d15563ec70b57ae35e22b739125094d40e5a6853746912b43ae56ce0
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2Nzc1NjQwMTh9.5dQj9FjSZ5KWR61ZISBvWMlIj9f79UqrdQeMGPrY9Ns
```

You can also set the token expiry in both the `claims` and `header` subcommands using the `--expiry` flag, providing a [duration string](https://pkg.go.dev/time#ParseDuration). For example:

```bash
$ jwtauth c --expiry 100s 0x66c01996d15563ec70b57ae35e22b739125094d40e5a6853746912b43ae56ce0
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Nzc2MjI1NzgsImlhdCI6MTY3NzYyMjQ3OH0.u_K0Zvv3niL5dIiSdBAebe3oWgtjamtZKT5kdcmoGvI
```

The `header` subcommand simply prefixes `Authorization: Bearer ` to the JWT Token like so.

```bash
$ jwtauth h --expiry 100s 0x66c01996d15563ec70b57ae35e22b739125094d40e5a6853746912b43ae56ce0
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Nzc2MjI2NjAsImlhdCI6MTY3NzYyMjU2MH0.fssZEUdKAQtPc1bwLPZ_2WIbpmDde8T7iy2oXLFBnQM
```

### Building from source

To build from source, run:

```bash
go build -o jwtauth
```


### License

[MIT](https://choosealicense.com/licenses/mit/)
