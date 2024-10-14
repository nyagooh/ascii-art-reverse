# ASCII Art Reverse

This project is a Go-based tool that reverses an ASCII art image, flipping ascii art to its original state. It processes ASCII characters to create a mirrored version of the original art.It is build on other ascii art projects meaning you can colour the ascii,write the ascii specifying the banner file and writing the ascii in a file.

## Features

- Flip ASCII art to its original state
- Write text in  ascii format with or without the banner file being specified.
- Write text in ascii format in a specified file.
- Colour text which is written in ascii format.

## Prerequisites

Before using this tool, ensure that you have the following installed:

- [Go](https://golang.org/dl/) (version 1.19 or higher)

## Installation

1. Clone the repository:

    ```bash
    git clone https://learn.zone01kisumu.ke/git/kada/ascii-art-reverse
    ```

2. Navigate into the project directory:

    ```bash
    cd ascii-art-reverse
    ```

3. Build the project:

    ```bash
    go build
    ```

## Usage

To reverse an ASCII art:

1. Prepare your ASCII art in a text file, for example, `art.txt`.
2. Run the following command:

    ```bash
    go run . --reverse=art.txt
    ```

3. The output will be the reversed ASCII art printed in the terminal.
```bash
  Hello
 ```
### Example 1: Simple Face

Given the following input in `art.txt`:
```bash
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                

```

These example show how  ASCII art is flipped .

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests to help improve the tool.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.


