package monthRus

import (
	"fmt"
	"time"
	"math"
)

var monthRus [12]string = [12]string{"Январь", "февраль", "март", "апрель", "май", "июнь", "июль", "август", "сентябрь", "октябрь", "ноябрь", "декабрь"}

func MonthLocalisation(month time.Month) (string){
	return monthRus[month-1]
}

func FormatDate (date time.Time) (string) {
	return fmt.Sprintf("%v %v %v", MonthLocalisation(date.Month()), date.Day(), date.Year())
}

func DaysUntil(date time.Time) (string){

	timeBefore:= time.Until(date)

	days:= math.Floor(timeBefore.Hours()/24)

	hours:= math.Floor(timeBefore.Hours() - days*24) 

	txt:= fmt.Sprintf("%v дней %v часов", days, hours)

	return txt

}