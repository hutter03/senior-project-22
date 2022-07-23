# Mutation Machine

This is a polymorphic engine that started off as my highschool senior project, and it is something I would like to continue to develop out of academic interest. This projcet can be used to take a linux binary executable file and create source code for a logically equivalent program that will run on linux.
## Using this code

- GO and Python need to be installed
- Compile main.go
- Your compiled main.go file should be in the same folder as decrypt-exec and payloadhex
- The payloadhex file included in this repo is an example and should be replaced by the one you create using the python script


## Creating the payloadhex file
```bash
  touch filedump.txt
  xxd -p yourBinary > filedump.txt
  touch payloadhex
  python3 payloadhexmaker.py > payloadhex
```


## What I want to do next

- Create more junk code choices
- Make the decrypthex file modular
- Make main.go automatically compile the source code that it generates 
- Replace the Python file with code written in GO
- Potentially change the encryption / encoding scheme
