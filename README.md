# Cryptographic Password Generator (`pgen`)

A simple and secure console-based password generator written in Go.\
It uses cryptographically strong randomness (`crypto/rand`) and supports
flexible command‑line options.

Author: **Andrew Razuvaev**\
GitHub: https://github.com/posixfan \
Email: posixfan87@yandex.ru

------------------------------------------------------------------------

## Features

-   Generates strong, random passwords using cryptographically secure
    RNG 
-   Supports configurable password length
-   Can generate multiple passwords in a single run
-   Optional flag to exclude punctuation characters
-   Ensures the first character is always a letter or digit
-   Simple command-line interface

------------------------------------------------------------------------

## Installation

``` bash
go install github.com/posixfan/passgen@latest
```

After installation, the executable `pgen` will appear in your
`$GOPATH/bin` or Go bin directory.

------------------------------------------------------------------------

## Usage

    pgen                    # Generate one 14-character password
    pgen <length>           # Generate one password with the specified length
    pgen <length> <count>   # Generate multiple passwords
    pgen --no-punct         # Exclude punctuation characters
    pgen -h, --help         # Show help message

------------------------------------------------------------------------

## Options

  Option           Description
  ---------------- -----------------------------------------------------
  `<length>`       Password length (minimum: 8 characters)
  `<count>`        Number of passwords to generate
  `--no-punct`     Excludes punctuation symbols from the character set
  `-h`, `--help`   Display usage help

------------------------------------------------------------------------

## Examples

``` bash
pgen 16
# Generates one 16-character password

pgen 12 5
# Generates five 12-character passwords

pgen --no-punct
# Generates a password without punctuation

pgen 16 3 --no-punct
# Generates three 16‑character passwords without punctuation

pgen --no-punct 16 3
# Flags may appear anywhere in the command
```

------------------------------------------------------------------------

## How It Works

-   The program builds a character set based on enabled flags
-   The first character is always chosen from letters and digits
-   Remaining characters are selected from the full chosen character
    set
-   Random numbers are generated via `crypto/rand` for high entropy

------------------------------------------------------------------------

## Repository

https://github.com/posixfan/passgen

------------------------------------------------------------------------

## License

GPL-3.0 License\
Copyright (c) Andrew Razuvaev
