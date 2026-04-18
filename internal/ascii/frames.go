package ascii

const (
	Pink  = "\033[38;5;205m"
	Teal  = "\033[38;5;51m"
	Red   = "\033[31m"
	Dim   = "\033[2m"
	Reset = "\033[0m"
)

func GetNeko(cpu float64, mem float64, active bool, uptime uint64) string {
	if cpu > 70 {
		return Red + `
       /\___/\
      ( > x < )  !!!!
      /  vvv  \  CPU CRITICAL
     (  |   |  )
      \_|___|_/ ` + Reset
	}
	if mem > 80 {
		return Pink + `
       /\___/\
      ( @ o @ )  *heavy*
      /  ---  \  RAM BLOAT
     (  |   |  )
      \_|___|_/ ` + Reset
	}
	if !active {
		return Dim + `
       /\___/\
      ( - . - )  ...
      /  ---  \  LONELY
     (  |   |  )
      \_|___|_/ ` + Reset
	}
	// Default "Cyber-Brain" Active State
	return Teal + `
       /\___/\
      ( ^ . ^ )  ONLINE
      /  > <  \  SYNCED
     (  |   |  )
      \_|___|_/ ` + Reset
}
