# 🔐 PassGen --- Console-based Cryptographic Password Generator

PassGen is a console-based generator that uses cryptographically secure
random number generation (`crypto/rand`).\
It creates strong, brute‑force‑resistant passwords and supports flexible
configuration of length and quantity.

------------------------------------------------------------------------

## ✨ Features

-   Cryptographically secure password generation
-   Customizable password length
-   Ability to generate multiple passwords
-   First character is always a letter or digit
-   Uses full printable ASCII set
-   Convenient help menu (`-h`)

------------------------------------------------------------------------

## 🛠 Usage

### Basic run

    pgen

Generates a single **14‑character** password.

### Set a custom length

    pgen <length>

Example:

    pgen 16

### Generate multiple passwords

    pgen <length> <count>

Example:

    pgen 12 5

------------------------------------------------------------------------

## 📌 Restrictions

-   Minimum length: **8 characters**
-   All errors are properly handled and shown in the console

------------------------------------------------------------------------

## 🆘 Help

    pgen -h

Displays: - syntax
- examples
- explanations of arguments

------------------------------------------------------------------------

## 🔒 Character Sets

The generator uses:

-   **Letters**: `a-zA-Z`

-   **Digits**: `0-9`

-   **ASCII punctuation**:

        !"#$%&'()*+,-./:;<=>?@[\]^_`{|}~

The first character is selected only from:
`a-zA-Z0-9`

------------------------------------------------------------------------

## 🧩 Technical Details

Random values are generated using:

``` go
rand.Int(rand.Reader, big.NewInt(int64(max)))
```

Characters are selected through a secure helper:

``` go
secureRandomChoice()
```

The program includes: - secure integer generation
- password construction
- CLI argument parsing
- error and help output

------------------------------------------------------------------------

## 📄 Sample Output

    N7gTJ*yWZ2H!pM
    Bv4{3]Amx/YP:8
    qF2G<Z$Hr]9eWf

------------------------------------------------------------------------

## 📚 Build

    go build -o pgen

------------------------------------------------------------------------

## 📥 Run

    ./pgen 20 3
