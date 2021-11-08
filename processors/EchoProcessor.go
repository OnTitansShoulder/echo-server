package processors

const (
	unknownIP = "unknown"
)

type Echo struct {
	IP           string
	FirstHitTime string
	LastHitTime  string

	LastHitTimestamp int64 // for sorting
	Count            int32
}

func TakeEchos(allEchos map[string]Echo, echoChan chan Echo) {
	for {
		echo := <-echoChan
		addEcho(allEchos, echo)
	}
}

func addEcho(allEchos map[string]Echo, echo Echo) {
	if _, ok := allEchos[echo.IP]; ok {
		// seen this source IP before
		echo.FirstHitTime = allEchos[echo.IP].FirstHitTime
		echo.Count = allEchos[echo.IP].Count + 1
	}
	allEchos[echo.IP] = echo
}
