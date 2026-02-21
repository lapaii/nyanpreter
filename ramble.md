# my rambles

## new based swag rambles

## instruction encoding

each instruction begins with a 2 byte header as follows:

| 15 - 9 | 8 - 4  | 3 - 2     | 1 - 0       |
| ------ | ------ | --------- | ----------- |
| unused | opcode | dest type | source type |

### Bit Layout

- bits 0–1: Source operand type
- bits 2–3: Destination operand type
- bits 4–8: Opcode (5 bits, up to 32 instructions, currently have 29)
- bits 9–15: unused

after the header, operand data follows in the order:

`[dest operand bytes][source operand bytes]`

the number of bytes used depends on the operand type. because i now know exactly how many bytes each instruction will take up, i dont need to null delimit them.

---

## operand types

which type an operand is is determined by the syntax of the program

| type | meaning               | encoding size |
| ---- | --------------------- | ------------- |
| 00   | Register              | 1 byte        |
| 01   | Register Dereference  | 1 byte        |
| 10   | Immediate Value       | 4 bytes       |
| 11   | Immediate Dereference | 4 bytes       |

---

## operand syntax examples and their types

### 00 — register

access registers directly

```asm
mov %r0, %r1 ; whatever is in r1 -> r0
```

encodes both operands as type `00`.

operand byte contains the register number.

---

### 01 — register dereference

memory access through a register.

```asm
mov %r0, (%r1) ; value in the memory address stored in r1 copied into r0
```

need to dereference in runtime

---

### 10 — immediate value

immediate value or label address.

```asm
mov %r0, $5 ; 5 now in r0
mov %r0, label ; address of label in r0
```

- `$5` encodes literal `5`
- `label` encodes the resolved address of `label`

---

### 11 — immediate dereference

```asm
mov %r0, ($5)
mov %r0, (label)
```

---

examples:

```asm
mov %r0, label     ; r0 = address of label
mov %r0, (label)   ; r0 = value stored at label

mov %r0, %r1       ; r0 = value in r1
mov %r0, (%r1)     ; r0 = value at memory address stored in r1

mov %r0, $5        ; r0 = 5
mov %r0, ($5)      ; r0 = value at memory address 5
```

## OLD RAMBLES

basically going to cmd+a delete everything

instead of being treated like an interpreter, im going to actually assemble the source code into object programs.

Why am i doing this?

1. because i want to learn more about assemblers and stuff because i think they are Cool.
2. my code is a mess and an interpreter isn't _really_ the most _ideal_ way to do something like this.
3. i want to learn more about assemblers

how it will be implemented next

there will be an additional DAT instruction (DAT #n/Bn/&n), for use with labels for defining data.

example:

```nyan
LDA letterA
OUT // will output "A"

letterA: DAT #65
```

will be a two-pass assembler

on the first pass i:

- go through each line of code
- is it a label definition?
  - yes:
    - store label name and line in a table (address symbol table)
  - no:
    - is it END instruction?
      - yes:
        - go to second pass
      - no:
        - go to next line

on the second pass:

- go through each line
- is instruction valid?
  - yes:
    - is operator valid? (some instructions dont take one or require an address etc)
      - yes:
        - is the operator a label?
          - yes:
            - using the symbol table, replace with label line and save to output
          - no:
            - save to output
      - no:
        - stop and report error to user
  - no:
    - stop and report error to user
- output the assembled program to the binary format

```go
type Instruction struct {
  Opcode int // will be later decoded back into instructions that are useful
  Operator int // Numbr .
}
```

- output binary to file

assembly done!! yayy

running the code:

- basically how i am now, read in file according to the spec i set for the binary file
- run it lol
