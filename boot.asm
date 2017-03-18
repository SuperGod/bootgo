section .data
  ALIGNED equ 1<<0
  MEMINFO equ 1<<1
  VIDEOINFO equ 1<<2
  MODETYPE equ 0
  WIDTH equ 320
  HEIGHT equ 200
  DEPTH equ 24
  MAGIC equ 0x1BADB002
  FLAGS equ ALIGNED|MEMINFO|VIDEOINFO
  CHECKSUM equ -(MAGIC + FLAGS)
  ZERO equ 0
section .multiboot
align 4
  dd MAGIC
  dd FLAGS
  dd CHECKSUM
  dd ZERO
  dd ZERO
  dd ZERO
  dd ZERO
  dd ZERO		
  dd MODETYPE
  dd WIDTH
  dd HEIGHT
  dd DEPTH

	
section .boot_stack
	align 4
stack_bottom:
resb 16384
stack_top:

section .text
global _start
_start:

  extern bootgo.kernel.Main
  push ebx
  push eax
  call bootgo.kernel.Main
  cli

.hang:
  hlt
  jmp .hang
