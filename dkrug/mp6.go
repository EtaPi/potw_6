// Package mp6 contains a golang solution for the calendar problem of the week,
// possibly WORST CODE EVER! D/K 2015
package main

import "fmt"
import "flag"
import "time"
import "bytes"
import "strconv"

const monthPadding string = "   "

var daysInMonths  = [...]int{ 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31 }
var monthNames    = [...]string{ "January", "February", "March", "April", "May",
  "June", "July", "August", "September", "October", "November", "December" }

func dayOfYearStart(year int) int {

  //Zeller's Congruence
  var adjustedYear int = year - 1
  var k            int = adjustedYear % 100
  var j            int = adjustedYear / 100
  var result       int = (37 + k + (k / 4) + (j / 4) + (5 * j)) % 7

  return convertDayIndex(result)
}

func convertDayIndex(day int) int {

  if day == 0 {
    return 6
  } else {
    return day - 1
  }
}

func computeCalendars(year int) (results [][][]int) {

  var dayIndex            int = dayOfYearStart(year)
  var dayCount            int
  var isMonthComplete     bool

  if year % 4 == 0 && year % 100 != 0 || year % 400 == 0 {
    daysInMonths[1] = 29
  }

  results = make([][][]int, 12)

  for i := range results {
    results[i] = make([][]int, 6)
  }

  for i := range results {
    for j := range results[i] {
      results[i][j] = make([]int, 7)
    }
  }

  for i := 0; i < 12; i++ {

    isMonthComplete = false
    dayCount        = 0

    for j := 0; j < 6; j++ {
      for k := dayIndex; k < 7; k++ {

        if dayCount += 1; dayCount > daysInMonths[i] {
          isMonthComplete = true
          break
        }

        results[i][j][k] = dayCount

        if dayIndex += 1; dayIndex > 6 {
          dayIndex = 0
        }
      }

      if isMonthComplete {
        break
      }
    }
  }

  return
}

func renderCalendar(calendars [][][]int, year int, width int) {

  var buffer bytes.Buffer
  var totalWidth    int = (width * 20) + ((width - 1) * len(monthPadding))
  var rows          int = 12 / width

  if 12 % width > 0 {
    rows += 1
  }

  buffer.WriteString("\n")

  buffer.WriteString(padCentered(strconv.Itoa(year), totalWidth, false, true))

  buffer.WriteString("\n")

  for i := 0; i < rows; i++ {
    buffer.WriteString(createMonthHeaderRow(i, width))
    buffer.WriteString(createDayHeaderRow(i, width))
    buffer.WriteString(createDayRows(calendars, i, width))
    buffer.WriteString("\n")
  }

  fmt.Println(buffer.String())
}

func createMonthHeaderRow(row int, width int) (string) {

  var buffer bytes.Buffer

  for i := 0; i < width; i++ {

    if i > 0 {
        buffer.WriteString(monthPadding)
    }

    if row * width + i < 12 {
      if row == 0 {
        buffer.WriteString(padCentered(monthNames[i], 20, true, false))
      } else {
        buffer.WriteString(padCentered(monthNames[i + (row * width)], 20, true, false))
      }
    }
  }

  buffer.WriteString("\n")

  return buffer.String()
}

func createDayHeaderRow(row int, width int) (string) {

  var buffer bytes.Buffer

  for i := 0; i < width; i++ {

    if i > 0 {
        buffer.WriteString(monthPadding)
    }

    if row * width + i < 12 {
      buffer.WriteString("Su Mo Tu We Th Fr Sa")
    }
  }

  buffer.WriteString("\n")

  return buffer.String()
}

func createDayRows(calendars [][][]int, row int, width int) (string) {

  var buffer bytes.Buffer
  var day int

  for i := 0; i < 6; i++ {
    for j := 0; j < width; j++ {
      if j > 0 {
          buffer.WriteString(monthPadding[:len(monthPadding) - 1])
      }

      if row * width + j < 12 {

        for k := 0; k < 7; k++ {

          if row == 0 {
            day = calendars[j][i][k]
          } else {
            day = calendars[j + (row * width)][i][k]
          }

          if day == 0 {
            buffer.WriteString("   ")
          } else {
              buffer.WriteString(fmt.Sprintf("%2d ", day))
          }
        }
      }
    }

    buffer.WriteString("\n")
  }
  return buffer.String()
}

func padCentered(s string, width int, includePadding bool, includeNewLine bool) (string) {

  var buffer bytes.Buffer
  var spaces int = (width / 2) - (len(s) / 2)

  for i := 0; i < spaces; i++ {
    buffer.WriteString(" ")
  }

  buffer.WriteString(s)

  if includePadding {
    for i := spaces + len(s); i < width; i++ {
      buffer.WriteString(" ")
    }
  }

  if includeNewLine {
    buffer.WriteString("\n")
  }

  return buffer.String()
}

func main() {

  localNow  := time.Now()
  argsYear  := flag.Int("year", localNow.Year(), "Prints calendar for the specified year")
  argsWidth := flag.Int("width", 3, "Number of output columns")

  flag.Parse()

  if *argsWidth > 0 {

    var calendars [][][]int = computeCalendars(*argsYear)

    renderCalendar(calendars, *argsYear, *argsWidth)
  }
}
