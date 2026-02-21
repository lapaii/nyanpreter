# nyasm spec

this is the doc for drafting up the new assembly-like language that this project will be based on moving forward.

it is mainly based of the existing language from the cambridge international 9618 cs course.

most instructions will have two operands, unlike just the one in the previous language

## syntax

the syntax for nyasm is a mix of my favourite parts of gas syntax and intel syntax

to define a label, it must be all in one line. allowed characters are A-z, - (hypen / minus) and \_ (underscore)

```asm
label:
  mov %r0, $5 ; ❌ not supported
```

```asm
label: mov %r0, $5 ; ✅ supported
```

convention will be writing mnemonics in lowercase but assembler will also support uppercase

comments are done in with a single ; anywhere in a line to comment out and ignore the rest of that line

parameter order: destination before source (mnemonic, destination, source)
parameter size: all registers and stuff is 32-bit

registers are to be prefixed with a %
immediate values prefixed with $, can define decimal numbers ($95) or hex numbers ($0x5F)

labels are dereferenced by default, i.e. `mov %r0, label` moves the contents of label into r0. to treat a label as a pointer, and only pass an address, you put a `&` at the beginning of the operator, i.e. `mov %r0, &label` would put the address of label into r0

the FLAGS register is a 32-bit register used for storing bits of state about the program. currently this only has 3 flags

bit 0: ZF (Zero Flag) set if the result of an operation is zero
bit 1: SF (Sign Flag) set if the result of an operation is negative
bit 2: OF (Overflow Flag) set if the result of an operation caused an overflow

## instruction table

i dont _see_ this expanding in the future, but you never know!

| opcode | operand             | explanation                                                                 |
| ------ | ------------------- | --------------------------------------------------------------------------- |
| mov    | dest, src           | copy the data from src to dest                                              |
| add    | dest, addend        | addend + dest and stores in dest                                            |
| sub    | minuend, subtrahend | subtrahend - minuend and store result in minuend                            |
| mul    | dest, multiplier    | multiplier \* dest and store result in dest                                 |
| div    | divisor, dividend   | integer divide dividend ÷ divisor and store result in divisor               |
| mod    | divisor, dividend   | dividend mod divisor and store result in divisor                            |
| and    | dest, mask          | bitwise AND both operands and store in dest                                 |
| not    | arg                 | bitwise NOT arg                                                             |
| xor    | dest, flip          | bitwise XOR both operands and stoere in dest                                |
| or     | dest, mask          | bitwise OR both operands and store in dest                                  |
| shr    | dest, cnt           | logical shift dest right by cnt bits                                        |
| shl    | dest, cnt           | logical shift dest left by cnt bits                                         |
| jmp    | loc                 | jump to location                                                            |
| cmp    | minuend, subtrahend | calculates subtrahend - minuend, FLAGS register is set based on the results |
| je     | loc                 | jump to location if ZF = 1                                                  |
| jne    | loc                 | jump to location if ZF = 0                                                  |
| jz     | loc                 | jump to location if ZF = 1 (identical to je)                                |
| jnz    | loc                 | jump to location if ZF = 0 (identical to jne)                               |
| jge    | loc                 | jump to location if SF = OF or ZF = 1                                       |
| jg     | loc                 | jump to location if SF = OF and ZF = 0                                      |
| jle    | loc                 | jump to location if SF != OF or ZF = 1                                      |
| jl     | loc                 | jump to location if SF != OF                                                |
| in     | dest                | save ascii value of a character entered into stdin to dest                  |
| out    | src                 | output ascii value from src to stdout                                       |
| outd   | src                 | output value from src to stdout as a decimal number                         |
| call   | src                 | jump to address, while pushing the current PC to the return stack           |
| ret    |                     | pop from return stack and set PC to that location                           |
| dn     | number              | declare number                                                              |
| ds     | string              | declare string                                                              |
