package network

type Node struct {
	Name            string
	IP              string
	Memory          int32
	MemoryAllocated int32
	Disk            int32
	DiskAllocated   int32
	TaskAmount      int32
}
