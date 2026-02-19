# nyanmachine

the nyanmachine is an assembler & "virtual machine" that can assemble and run the assembly language provided as part of the [Cambridge International AS & A Level Computer Science syllabus](https://www.cambridgeinternational.org/Images/636089-2024-2025-syllabus.pdf).

as a student studying under this syllabus, i wanted to learn more about the topics we are taught in class which is what has motivated me to persue this project.

see my [rambling](ramble.md) for more info about my approach to this project and some info about how the code works

you are viewing this on the extra branch!

## planned additions

- reworking jump-on-condition instructions:
  - ✅ renaming JPE -> JE
  - ✅ renaming JPN -> JNE
  - ✅ adding JGE (jump greater equal)
  - ✅ adding JLE (jump less equal)
- ✅ add CMX instruction (compare operator with IDX register)
- allow direct reading and manipulation of the PC register
- add MUL, DIV (integer division), MOD instructions
- add more general purpose registers (R1-9)
  - initially will be only able to manipulate them like

    ```asm
    LDM #10
    MUL #5 // 50 currently in ACC
    MOV R1 // moves 50 from ACC to R1 register
    ```

  - planned for future where you could do

    ```asm
    LDM #10
    MUL #5, R1 // multiply ACC by 5 and store result in R1, ACC stays as 10
    ```
