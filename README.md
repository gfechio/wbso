# WBSO Command Line Tool
Wet Bevordering Speur- en Ontwikkelingswerk

WBSO is a command-line tool designed to streamline workflow management, integrating with Git version control and YouTrack issue tracking. This tool provides functionalities to facilitate common operations during software development.

## Features

- **Setup**: Initialize a new Git repository and set up the project for version control.

- **Associate**: Associate a task or issue from YouTrack with the current branch.

- **Checkout**: Switch to a specified branch in the repository.

- **Commit**: Commit changes to the current branch in the repository.

- **Push**: Push the committed changes to the remote repository.

- **Close**: Close an issue or task in YouTrack, marking it as completed.

- **Alias**: Set up Git aliases for common Git commands.

## Installation

1. Clone the repository:

    ```bash
    git clone <repository_url>
    ```

2. Navigate to the project directory:

    ```bash
    cd WBSO
    ```

3. Build the binary:

    ```bash
    go build -o WBSO main.go
    ```

4. Add the binary to your system path (optional):

    ```bash
    export PATH=$PATH:/path/to/WBSO
    ```

## Usage

To use the WBSO command-line tool, run the executable and specify a subcommand.

```bash
./WBSO <subcommand>
```

Replace `<subcommand>` with one of the available features: `setup`, `associate`, `checkout`, `commit`, `push`, `close`, or `alias`.

For more information and options on each subcommand, refer to the appropriate section below.

### Subcommands

- **Setup**: Initialize a new Git repository and set up the project for version control.

- **Associate**: Associate a task or issue from YouTrack with the current branch.

- **Checkout**: Switch to a specified branch in the repository.

- **Commit**: Commit changes to the current branch in the repository.

- **Push**: Push the committed changes to the remote repository.

- **Close**: Close an issue or task in YouTrack, marking it as completed.

- **Alias**: Set up Git aliases for common Git commands.

## Contribution

Contributions to the WBSO command-line tool are welcome! Feel free to open issues for bugs, feature requests, or general improvements. If you'd like to contribute, fork the repository and submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).

---

Feel free to customize and enhance this README according to the specific functionalities and details of your `WBSO` command-line tool integrating Git and YouTrack.
