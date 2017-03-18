GCCGO := i686-elf-gccgo
GCC := 	i686-elf-gcc
NASM := nasm
QEMU := qemu-system-i386
MV := mv
GRUB := grub-mkrescue

default: build

build: bootgo.bin
	$(MV) bootgo.bin isodir/boot/bootgo.bin
	$(GRUB) -o bootgo.iso isodir

bootgo.bin: boot.o kernel.o runtime/libgo.so
	$(GCC) -T linker.ld -o bootgo.bin  -ffreestanding -nostdlib boot.o kernel.o runtime/libgo.so -lgcc 

boot.o: boot.asm
	$(NASM) -felf32 boot.asm -o boot.o

kernel.o: kernel.go
	$(GCCGO) -c kernel.go video.go memio.go -fgo-prefix=bootgo -static-libgo

runtime/libgo.so: runtime/libgo.c runtime/go-type-identity.c runtime/go-type-error.c
	cd runtime; \
	$(GCC) -shared libgo.c go-type-identity.c go-type-error.c -o libgo.so -std=gnu99 -ffreestanding

run:
	$(QEMU) -kernel bootgo.bin -vga vmware

clean:
	rm -f *.o
	rm -f **/*.so
	rm -f *.bin
	rm -f *.iso
