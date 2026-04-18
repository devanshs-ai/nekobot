package ascii

const (
	VengeanceRed = "\033[31m"
	CyberTeal    = "\033[36m"
	Reset        = "\033[0m"
)

func GetNeko(cpu float64, active bool) string {
	if cpu > 80 {
		return VengeanceRed + `
   |\---/|    ( OVERHEAT! )
   ( > < )  /
    > ^ <  # #
   /  ~  \
` + Reset
	}
	if !active {
		return "\033[2m" + `
   |\---/|    ( ...lonely )
   ( - - )  /
    v w v 
   /     \
` + Reset
	}
	return CyberTeal + `
   |\---/|    ( Online. )
   ( o o )  /
    > ^ < 
   /     \
` + Reset
}
