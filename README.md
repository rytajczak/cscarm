## CSCARM

simple assembler for a minified subset of the GNU arm assembler syntax

## Features

- Supports a subset of GNU arm assembler syntax.
- Parses mnemonics, operands, and directives.
- Handles basic arithmetic and logical operations.
- Supports labels and branching instructions.
- Generates machine code targetting the BCM2837 processor (Raspberry Pi 3).

## Pitfalls

- Absolutely zero robustness in mnemonic parsing.

## Installation

```bash
go install github.com/rytajczak/cscarm
```

## Usage

```bash
cscarm [options] <input_file>
```
