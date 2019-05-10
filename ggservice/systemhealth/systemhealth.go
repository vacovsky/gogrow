package systemhealth

import (
	"os"
	"time"

	"github.com/vacovsky/gogrow/ggmodels"
	"github.com/vacovsky/gogrow/ggservice/systemhealth/device"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"

	// "github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

// GetDBSize : returns the size of the database in bytes
func GetDBSize() {}

// GatherDeviceInfo returns health metrics on device
func GatherDeviceInfo() ggmodels.DeviceState {
	loadStat, _ := load.Avg()
	memStats, _ := mem.VirtualMemory()
	diskStats, _ := disk.Usage("/")
	cpuStats, _ := cpu.Info()
	cpuTimes, _ := cpu.Times(false)
	cpuTemp, humidity, ambTemp, thum, ttemp := device.GetThermalInfo()
	processes, _ := process.Processes()
	// internalIP, _ := net.InterfaceAddr

	ls, _ := host.BootTime()
	up, _ := host.Uptime()

	d := ggmodels.DeviceState{
		Uptime:   up,
		BootTime: ls,

		TotalMemMB:     memStats.Total / 1000000,
		FreeMemPercent: (float64(memStats.Free) / float64(memStats.Total)) * 100,
		UsedMemPercent: (float64(memStats.Used) / float64(memStats.Total) * 100),

		CPUUserPercent:   cpuTimes[0].User,
		CPUKernelPercent: cpuTimes[0].System,
		CPUIdlePercent:   cpuTimes[0].Idle,

		CPUCoreCount: cpuStats[0].Cores,
		CPULoad1:     loadStat.Load1,
		CPULoad5:     loadStat.Load5,
		CPULoad15:    loadStat.Load15,

		// ExternalIP: "",
		// InternalIP: "",

		DiskSize: diskStats.Total,
		DiskUsed: diskStats.Used,
		DiskFree: diskStats.Free,

		ProcessCount: len(processes),

		CPUTemperatureC: cpuTemp,

		AmbientHumidity:    humidity,
		AmbientTemperature: ambTemp,

		TentHumidity:    thum,
		TentTemperature: ttemp,

		TimeStamp:       time.Now(),
		PlatformVersion: os.Getenv("GG_PLATFORM"),
	}

	return d
}
