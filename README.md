## CSCASM

#### Pitfalls

- Absolutely zero robustness in mnemonic parsing.
- Branch instruction checks for link based on the second character of the mnemonic, which may not cover all cases.
