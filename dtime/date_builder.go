package dtime

import (
	"fmt"
)

func DateString(year, month, date int) string{
	return fmt.Sprintf("%d-%02d-%02d", year, month, date)
}
