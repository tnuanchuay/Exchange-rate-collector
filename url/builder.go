package url

import (
	"fmt"
	"ecr/dtime"
)

func HistoryUrl(endpoint string, year, month, date int) string{
	sdate := dtime.DateString(year, month, date)
	return fmt.Sprintf("%s/%s", endpoint, sdate)
}
