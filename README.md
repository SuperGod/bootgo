# bootgo

A barebones OS kernel written in go

The kernel contains some code modified from osdev barebones tutorial http://wiki.osdev.org/Bare_Bones

## setup
To compile bootgo You need a gccgo cross-compiler(my gccgo version is 5.2.0)

1. build a target i386/i686 gcc cross-compiler with go enabled, follow the article http://wiki.osdev.org/GCC_Cross-Compiler (gcc 6.2.0 is recommended)

2. install nasm from your repositories

3. install virtualbox for test

## compiler & run!

1. compile: `make GCC=i686-elf-gcc GCCGO=i686-elf-gccgo`, replace `GCC` and `GCCGO` with your binary name

2. run with virtualbox: Create a virtual machine, add the bootgo.iso to the CDROM,and start the virtual machine

# go fetures:
I copy some c func from gccgo/libgo/runtime from:

https://github.com/gcc-mirror/gcc/tree/master/libgo/runtime

so can support more go fetures:

## supported
### struct
### struct funcs

## donot support:

### new 
### make
### routine
