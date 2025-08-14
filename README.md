# AUR Status Checker

This is a simple Go application that continuously checks the status of the Arch User Repository (AUR) and displays it in the terminal.

![AUR Status Checker](./cli.png)

## Downloading the Pre-built Binary

You can download the latest pre-built binary for Linux from the GitHub Releases page.

1.  Go to the [latest release page](https://github.com/bjarneo/aur-up/releases/latest).
2.  Download the `aur-up` binary.
3.  (Optional) Download the `aur-up.sha256` file to verify the integrity of the binary.
4.  Make the binary executable:
    ```bash
    chmod +x aur-up
    ```
5.  (Optional) Move the binary to a directory in your `PATH`:
    ```bash
    sudo mv aur-up /usr/local/bin/
    ```

## Building from Source

### Prerequisites

- Go programming language installed on your system.

### Instructions

1.  **Clone the repository or download the source code.**

2.  **Navigate to the project directory:**

    ```bash
    cd aur-up
    ```

3.  **Tidy the dependencies:**

    ```bash
    go mod tidy
    ```

4.  **Run the application:**

    ```bash
    go run main.go
    ```

5.  **Build the application:**

    To build the application into a single executable, run the following command:

    ```bash
    go build -o aur-up main.go
    ```
