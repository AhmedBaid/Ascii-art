# Ascii-Art

## Objectives

**Ascii-Art** is a program that converts a string into a graphical representation using ASCII characters. This project involves handling inputs with letters, numbers, spaces, special characters, and newlines (`\n`), transforming them into a visual format based on predefined ASCII templates.

## Features

- **Graphical Representation:** Converts input strings into ASCII art.
- **Character Support:** Handles letters, numbers, spaces, special characters, and newlines.
- **Banner Templates:** Uses predefined ASCII templates for characters.
- **Custom Layouts:** Supports different banner styles like `shadow`, `standard`, and `thinkertoy`.
- **Multiple Lines:** Processes inputs with line breaks for multi-line ASCII art.
- **Error Handling:** Gracefully manages invalid or unsupported inputs.

## Banner Format

Each character is represented as an 8-line block, separated by a newline (`\n`). Below are examples of the ASCII representations for `' '`, `'!'`, and `'"'`:

### Example:

#### Space (`' '`)
```
......
......
......
......
......
......
......
......
```

#### Exclamation (`'!'`)
```
._..
|.|.
|.|.
|.|.
|_|.
(_).
....
....
```

#### Double Quote (`'"'`)
```
._._..
(.|.).
.V.V..
......
......
......
......
......
```

## Instructions

Your project must be written in **Go** and adhere to Go's best practices. It is also recommended to include unit tests for validation.

### Allowed Packages

- Only the standard Go packages are allowed.

### Usage

The program receives an input string and outputs its ASCII representation to the terminal. Here are some examples:

```bash
$ go run . "Hello"
 _    _          _   _         
| |  | |        | | | |        
| |__| |   ___  | | | |   ___  
|  _ \   / _ \ | | | |  / _ \ 
| | | | |  __/ | | | | | (_) |
|_| |_|  \___| |_| |_|  \___/ 
                              
                              
```

```bash
$ go run . "Hello\nThere"
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
 _______   _                           $
|__   __| | |                          $
   | |    | |__     ___   _ __    ___  $
   | |    |  _ \   / _ \ | '__|  / _ \ $
   | |    | | | | |  __/ | |    |  __/ $
   |_|    |_| |_|  \___| |_|     \___| $
                                       $
                                       $
```

### Edge Cases

The program also handles:
- **Empty Inputs:** Produces no output for an empty string.
- **Newlines (`\n`):** Generates ASCII art for multi-line text.
- **Unsupported Characters:** Skips any characters not defined in the templates.
- **Mixed Cases:** Renders uppercase and lowercase characters appropriately.

## Project Guidelines

### Templates

The project includes three ASCII templates for styling the output:
1. `shadow`
2. `standard`
3. `thinkertoy`

These templates must not be modified and should be loaded as-is.

### File Handling

- **Input:** Read the input string from the program arguments.
- **Templates:** Load banner templates from predefined files.
- **Output:** Write the rendered ASCII art directly to the terminal.

### Error Handling

1. **Invalid Inputs:** Skips characters that are not found in the templates.
2. **Missing Templates:** Exits with an appropriate error message if a required template file is missing or unreadable.
3. **Empty Inputs:** Outputs nothing but handles the case gracefully.

### Testing

Unit tests should validate:
- Correct ASCII rendering for all supported characters.
- Proper handling of newlines and multi-line strings.
- Skipping unsupported characters without crashing.
- Handling edge cases like empty input or missing template files.

## Learning Outcomes

This project will help you learn and apply:
- **Go File System API:** Read and manipulate files in Go.
- **String Manipulation:** Process and transform strings for ASCII art.
- **Error Handling:** Write robust code with clear error messages.
- **Unit Testing:** Implement and validate test cases for your program.
- **Code Practices:** Follow Go best practices for readability and maintainability.

---

Happy coding! 🚀
