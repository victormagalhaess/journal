package date

import "time"

const (
	slashDDMMYY = "02/01/2006"
	dashDDMMYY  = "02-01-2006"
	today       = "today"
	yesterday   = "yesterday"
)

type dateFn = func() time.Time

var specialDates map[string]dateFn = map[string]dateFn{
	today:     func() time.Time { return time.Now() },
	yesterday: func() time.Time { return time.Now().AddDate(0, 0, -1) },
}

func IsDate(input string) bool {
	if _, exists := specialDates[input]; !exists {
		_, err := time.Parse(slashDDMMYY, input)
		if err != nil {
			_, err = time.Parse(dashDDMMYY, input)
			if err != nil {
				return false
			}
		}
	}
	return true
}

func DateParse(input string) string {
	var (
		date time.Time
		err  error
	)

	if _, exists := specialDates[input]; !exists {
		date, err = time.Parse(slashDDMMYY, input)
		if err != nil {
			date, err = time.Parse(dashDDMMYY, input)
			if err != nil {
				date = time.Now()
			}
		}
	} else {
		date = specialDates[input]()
	}

	return dateFormatter(date)
}

func dateFormatter(date time.Time) string {
	return date.Format(slashDDMMYY)
}
