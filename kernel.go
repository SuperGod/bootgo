package kernel

var (
	row, column int
	color       uint8
	buffer      uintptr
)

/*
type MultibootAoutSymbolTable struct {
	TabSize  uint32
	StrSize  uint32
	Addr     uint32
	Reserved uint32
}
type MultibootElfSectionHeaderTable struct {
	Num   uint32
	Size  uint32
	Addr  uint32
	Shndx uint32
}*/
// type MultibootInfo struct {
// 	/* Multiboot info version number */
// 	Flags uint32

// 	/* Available memory from BIOS */
// 	MemLower uint32
// 	MemUpper uint32

// 	/* "root" partition */
// 	BootDevice uint32

// 	/* Kernel command line */
// 	Cmdline uint32

// 	/* Boot-Module list */
// 	ModsCount uint32
// 	ModsAddr  uint32
// 	ElfSec    MultibootElfSectionHeaderTable

// 	/* Memory Mapping buffer */
// 	MmapLength uint32
// 	MmapAddr   uint32

// 	/* Drive Info buffer */
// 	DrivesLength uint32
// 	DrivesAddr   uint32

// 	/* ROM configuration table */
// 	ConfigTable uint32

// 	/* Boot Loader Name */
// 	BootLoaderName uint32

// 	/* APM table */
// 	ApmTable uint32

// 	/* Video */
// 	VbeControlInfo  uint32
// 	VbeModeInfo     uint32
// 	VbeMode         uint32
// 	VbeInterfaceSeg uint16
// 	VbeInterfaceOff uint16
// 	VbeInterfaceLen uint16
// }

func Main(magic, addr uintptr) {
	Init(320, 200, COLOR_WHITE, COLOR_WHITE)
}
