package cpu

func CalculateUsage(stat CPUStat, prevTotal *uint64, prevIdle *uint64) float64 {
	total := stat.totalTime()
	idle := stat.idleTime()

	if *prevTotal == 0 {
		*prevTotal = total
		*prevIdle = idle
		return 0
	}

	totalDelta := total - *prevTotal
	idleDelta := idle - *prevIdle

	*prevTotal = total
	*prevIdle = idle

	if totalDelta == 0 {
		return 0
	}

	return float64(totalDelta-idleDelta) / float64(totalDelta) * 100 // Usage
}

func (s CPUStat) idleTime() uint64 {
	return s.Idle + s.Iowait
}

func (s CPUStat) totalTime() uint64 {
	return s.User + s.Nice + s.System + s.Idle + s.Iowait + s.IRQ + s.SoftIRQ + s.Steal
}
