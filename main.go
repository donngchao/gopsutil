package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
	// "github.com/shirou/gopsutil/mem"  // to use v2
)

func main() {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	// 计算内存的总量，空闲内存和使用率
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	//打印整个结构体
	fmt.Println(v)

	pids, _ := process.Pids()
	fmt.Printf("%#v", pids)
}
