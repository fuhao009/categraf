package cpu

import (
	"log"

	cpuUtil "github.com/shirou/gopsutil/v3/cpu"

	"flashcat.cloud/categraf/config"
	"flashcat.cloud/categraf/inputs"
	"flashcat.cloud/categraf/inputs/system"
	"flashcat.cloud/categraf/types"
)

const inputName = "cpu"

type CPUStats struct {
	ps        system.PS
	lastStats map[string]cpuUtil.TimesStat

	config.PluginConfig
	CollectPerCPU bool `toml:"collect_per_cpu"`
}

func init() {
	inputs.Add(inputName, func() inputs.Input {
		return &CPUStats{
			ps: system.NewSystemPS(),
		}
	})
}

func (c *CPUStats) Clone() inputs.Input {
	return &CPUStats{
		ps: system.NewSystemPS(),
	}
}

func (c *CPUStats) Name() string {
	return inputName
}

func (c *CPUStats) Gather(slist *types.SampleList) {
	times, err := c.ps.CPUTimes(c.CollectPerCPU, true)
	if err != nil {
		log.Println("E! failed to get cpu metrics:", err)
		return
	}

	for _, cts := range times {
		tags := map[string]string{
			"cpu": cts.CPU,
		}

		total := totalCPUTime(cts)
		active := activeCPUTime(cts)

		// Add in percentage
		if len(c.lastStats) == 0 {
			// If it's the 1st gather, can't get CPU Usage stats yet
			continue
		}

		lastCts, ok := c.lastStats[cts.CPU]
		if !ok {
			continue
		}

		lastTotal := totalCPUTime(lastCts)
		lastActive := activeCPUTime(lastCts)
		totalDelta := total - lastTotal

		if totalDelta < 0 {
			log.Println("W! current total CPU time is less than previous total CPU time")
			break
		}

		if totalDelta == 0 {
			continue
		}

		// 定义一个包含CPU使用率字段和对应值的map
		fields := map[string]interface{}{
			// 用户态使用CPU的时间比例，计算公式为 (当前用户态时间 - 上次用户态时间 - (当前虚拟机用户态时间 - 上次虚拟机用户态时间)) / 总增量
			"user": 100 * (cts.User - lastCts.User - (cts.Guest - lastCts.Guest)) / totalDelta,
			// 系统态使用CPU的时间比例，计算公式为 (当前系统态时间 - 上次系统态时间) / 总增量
			"system": 100 * (cts.System - lastCts.System) / totalDelta,
			// CPU处于空闲状态的时间比例，计算公式为 (当前空闲时间 - 上次空闲时间) / 总增量
			"idle": 100 * (cts.Idle - lastCts.Idle) / totalDelta,
			// 用户态使用CPU的时间比例（低优先级），计算公式为 (当前低优先级用户态时间 - 上次低优先级用户态时间 - (当前虚拟机低优先级用户态时间 - 上次虚拟机低优先级用户态时间)) / 总增量
			"nice": 100 * (cts.Nice - lastCts.Nice - (cts.GuestNice - lastCts.GuestNice)) / totalDelta,
			// CPU等待I/O完成的时间比例，计算公式为 (当前I/O等待时间 - 上次I/O等待时间) / 总增量
			"iowait": 100 * (cts.Iowait - lastCts.Iowait) / totalDelta,
			// 处理硬件中断的时间比例，计算公式为 (当前硬件中断时间 - 上次硬件中断时间) / 总增量
			"irq": 100 * (cts.Irq - lastCts.Irq) / totalDelta,
			// 处理软件中断的时间比例，计算公式为 (当前软件中断时间 - 上次软件中断时间) / 总增量
			"softirq": 100 * (cts.Softirq - lastCts.Softirq) / totalDelta,
			// 抢占时间比例，计算公式为 (当前抢占时间 - 上次抢占时间) / 总增量
			"steal": 100 * (cts.Steal - lastCts.Steal) / totalDelta,
			// 虚拟机运行在非特权级别的时间比例，计算公式为 (当前虚拟机非特权态时间 - 上次虚拟机非特权态时间) / 总增量
			"guest": 100 * (cts.Guest - lastCts.Guest) / totalDelta,
			// 虚拟机运行在低优先级非特权级别的时间比例，计算公式为 (当前虚拟机低优先级非特权态时间 - 上次虚拟机低优先级非特权态时间) / 总增量
			"guest_nice": 100 * (cts.GuestNice - lastCts.GuestNice) / totalDelta,
			// 活跃CPU时间的比例，计算公式为 (当前活跃时间 - 上次活跃时间) / 总增量
			"active": 100 * (active - lastActive) / totalDelta,
		}

		slist.PushSamples("cpu_usage", fields, tags)
	}

	c.lastStats = make(map[string]cpuUtil.TimesStat)
	for _, cts := range times {
		c.lastStats[cts.CPU] = cts
	}
}

func totalCPUTime(t cpuUtil.TimesStat) float64 {
	total := t.User + t.System + t.Nice + t.Iowait + t.Irq + t.Softirq + t.Steal + t.Idle
	return total
}

func activeCPUTime(t cpuUtil.TimesStat) float64 {
	active := totalCPUTime(t) - t.Idle
	return active
}
