# Pass-inator 🔐

A secure, command-line password generator written in Go. Pass-inator generates cryptographically secure passwords with configurable character sets and length.

## Features

- 🔒 Cryptographically secure password generation using `crypto/rand`
- ⚡ Fast and efficient password generation
- 🛠️ Configurable password parameters:
  - Custom password length (minimum 6 characters)
  - Optional character sets:
    - Lowercase letters (a-z)
    - Uppercase letters (A-Z)
    - Numbers (0-9)
    - Special characters (!@#$%^&*()_+-=[]{}|;:,.<>?)
- ✅ Guaranteed inclusion of at least one character from each selected character set
- 🔄 Secure password shuffling using Fisher-Yates algorithm
- 🚫 No dependencies on external packages

## Security Features

- Uses cryptographically secure random number generation
- Implements proper error handling for all security-critical operations
- Ensures uniform distribution of characters
- No predictable patterns in generated passwords
- Secure password shuffling algorithm

## Installation

1. Ensure you have Go installed (version 1.16 or higher recommended)
2. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/pass-inator.git
   cd pass-inator
   ```
3. Build the project:
   ```bash
   go build
   ```

## Usage

Run the program:
```bash
./pass-inator
```

The program will interactively ask for:
1. Password length (minimum 6 characters)
2. Character set preferences (lowercase, uppercase, numbers, special characters)

Example output:
```
Welcome to Pass-inator - Your Secure Password Generator
-----------------------------------------------------
Enter password length (minimum 6): 12
Include lowercase letters? (y/n): y
Include uppercase letters? (y/n): y
Include numbers? (y/n): y
Include special characters? (y/n): y

Your generated password is:
------------------------
6hbIgHsX,R:$
------------------------
```

## Security Considerations

- The program uses Go's `crypto/rand` package for cryptographically secure random number generation
- Each generated password includes at least one character from each selected character set
- Passwords are shuffled using the Fisher-Yates algorithm with secure random numbers
- All random number operations include proper error handling
- No time-based seeding is used, eliminating potential predictability

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Disclaimer

This tool is provided for educational and personal use. While it generates cryptographically secure passwords, the security of your accounts depends on many factors, including:

- How you store and manage your passwords
- The security of the services you use
- Your overall security practices

Always follow security best practices and use a password manager for storing your passwords securely.