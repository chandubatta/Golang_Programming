# Go `time` Package — Detailed Guide

The Go `time` package is one of the most important standard-library packages to understand because real-world applications constantly deal with timestamps, durations, time zones, deadlines, scheduling, timeouts, logging, database timestamps, JWT expiration, and performance measurement.

## 1. What is the `time` package?

The `time` package is part of Go's standard library and provides functionality for:

- Getting the current date/time
- Creating dates and times
- Formatting dates/times
- Parsing date/time strings
- Calculating differences between times
- Adding/subtracting time
- Working with durations
- Working with time zones
- Creating timers
- Creating periodic tickers
- Implementing timeouts
- Measuring how long code takes
- Working with Unix timestamps
- Handling deadlines and scheduled operations

Import it with:

```go
import "time"
```

The official documentation describes it as providing functionality for measuring and displaying time. Go's calendar calculations use the Gregorian calendar and do not represent leap seconds.

---

# 2. The most important concepts

Before learning individual functions, understand these four concepts:

```text
time.Time
    ↓
A specific point/moment in time

time.Duration
    ↓
An amount of elapsed time

time.Location
    ↓
A time zone/location

time.Timer / time.Ticker
    ↓
Future or repeated execution
```

For example:

```go
now := time.Now()

duration := 2 * time.Hour

india, _ := time.LoadLocation("Asia/Kolkata")

timer := time.NewTimer(5 * time.Second)

ticker := time.NewTicker(1 * time.Second)
```

These represent four very different things.

---

# 3. `time.Time`

`time.Time` represents an instant in time, with nanosecond precision.

```go
var t time.Time
```

For example:

```go
now := time.Now()

fmt.Println(now)
```

Output might look like:

```text
2026-09-08 12:30:45.123456789 +0530 IST
```

A `time.Time` contains information such as:

```text
Year
Month
Day
Hour
Minute
Second
Nanosecond
Location
```

The zero value of `time.Time` is January 1, year 1, 00:00:00 UTC.

You can check whether a `Time` has its zero value using `IsZero()`.

---

# 4. `time.Duration`

`time.Duration` represents an amount of elapsed time.

Internally, it is an `int64` number of nanoseconds.

```go
type Duration int64
```

Go provides these constants:

```go
time.Nanosecond
time.Microsecond
time.Millisecond
time.Second
time.Minute
time.Hour
```

For example:

```go
duration := 5 * time.Second
```

or:

```go
duration := 2 * time.Hour
```

There is deliberately no `time.Day` constant because a calendar day is not always exactly 24 hours in the presence of time-zone/DST transitions.

---

# 5. Simple example

Here is a small example combining several important functions:

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	fmt.Println("Current time:", start)

	time.Sleep(2 * time.Second)

	end := time.Now()

	fmt.Println("End time:", end)

	elapsed := end.Sub(start)

	fmt.Println("Elapsed:", elapsed)
	fmt.Println("Seconds:", elapsed.Seconds())
}
```

Conceptually:

```text
time.Now()
    ↓
start

time.Sleep()
    ↓
wait

time.Now()
    ↓
end

end.Sub(start)
    ↓
Duration
```

One important Go feature is that `time.Now()` can contain both wall-clock and monotonic-clock readings. Operations such as `Sub`, `Since`, `Until`, `Before`, and `After` can use the monotonic clock when appropriate, making elapsed-time measurements robust against wall-clock adjustments.

---

# 6. Package-level functions

## `time.Now()`

Returns the current local time.

```go
now := time.Now()

fmt.Println(now)
```

Use it when you need:

- Current timestamp
- Logging
- Creating records
- Measuring execution time
- Generating expiration times

Example:

```go
createdAt := time.Now()
```

---

## `time.Date()`

Creates a `time.Time` from individual components.

```go
t := time.Date(
	2026,
	time.September,
	8,
	12,
	30,
	0,
	0,
	time.UTC,
)
```

Syntax:

```go
time.Date(
	year,
	month,
	day,
	hour,
	minute,
	second,
	nanosecond,
	location,
)
```

Example:

```go
birthday := time.Date(
	1995,
	time.January,
	15,
	10,
	30,
	0,
	0,
	time.Local,
)

fmt.Println(birthday)
```

This is useful when you need to construct a specific date.

---

# 7. `time.Parse()`

Converts a string into `time.Time`.

```go
t, err := time.Parse(
	"2006-01-02",
	"2026-09-08",
)
```

Go doesn't use `YYYY-MM-DD` as its layout syntax.

Instead, Go uses a special reference date:

```text
Mon Jan 2 15:04:05 MST 2006
```

For example:

```go
layout := "2006-01-02"

t, err := time.Parse(layout, "2026-09-08")
```

The layout system is one of the most important things to learn in the `time` package.

---

# 8. `time.ParseInLocation()`

Similar to `Parse`, but allows you to specify the location.

```go
location, _ := time.LoadLocation("Asia/Kolkata")

t, err := time.ParseInLocation(
	"2006-01-02 15:04:05",
	"2026-09-08 14:30:00",
	location,
)
```

This is important when a string represents local clock time in a particular time zone.

---

# 9. `time.ParseDuration()`

Converts a string such as:

```text
"10s"
"5m"
"2h"
"1h30m"
```

into `time.Duration`.

Example:

```go
d, err := time.ParseDuration("2h30m")

fmt.Println(d)
```

Output:

```text
2h30m0s
```

Valid units include:

```text
ns
us / µs
ms
s
m
h
```

---

# 10. `time.Unix()`

Creates a `time.Time` from Unix seconds.

```go
t := time.Unix(0, 0)

fmt.Println(t)
```

Unix time represents the number of seconds since:

```text
1970-01-01 00:00:00 UTC
```

For example:

```go
timestamp := int64(1757318400)

t := time.Unix(timestamp, 0)
```

Very common when working with:

- APIs
- Databases
- Authentication
- JWT
- Distributed systems

---

# 11. `time.UnixMilli()`

Creates a `Time` from Unix milliseconds.

```go
t := time.UnixMilli(1757318400000)
```

Useful when an API provides timestamps in milliseconds.

---

# 12. `time.UnixMicro()`

Creates a `Time` from Unix microseconds.

```go
t := time.UnixMicro(1757318400000000)
```

---

# 13. `time.Since()`

Returns how much time has passed since a given time.

```go
start := time.Now()

// some operation

elapsed := time.Since(start)

fmt.Println(elapsed)
```

It is essentially shorthand for:

```go
time.Now().Sub(start)
```

Useful for elapsed-time measurement.

---

# 14. `time.Until()`

Returns how much time remains until a specific time.

```go
deadline := time.Now().Add(10 * time.Second)

remaining := time.Until(deadline)

fmt.Println(remaining)
```

Conceptually:

```text
deadline
   ↑
   │ remaining
   │
 now
```

Useful for:

- Token expiration
- Deadlines
- Scheduled operations
- Timeouts

---

# 15. `time.Sleep()`

Pauses the current goroutine for at least the specified duration.

```go
time.Sleep(2 * time.Second)
```

Example:

```go
fmt.Println("Start")

time.Sleep(3 * time.Second)

fmt.Println("End")
```

A zero or negative duration returns immediately.

### Important

Don't think of `Sleep` as a timer that executes something. It simply pauses the current goroutine.

---

# 16. `time.After()`

Returns a channel that receives the current time after a duration.

```go
select {
case <-time.After(5 * time.Second):
	fmt.Println("Timeout")
}
```

This is particularly useful with `select`.

Example:

```go
select {
case result := <-resultChannel:
	fmt.Println("Result:", result)

case <-time.After(5 * time.Second):
	fmt.Println("Request timed out")
}
```

This is a common Go timeout pattern.

In current Go versions, unreferenced timers created for `After` can be recovered by the garbage collector, so the old advice to avoid `time.After` merely for timer-GC reasons is no longer applicable.

---

# 17. `time.Tick()`

Creates a channel that periodically sends times.

```go
ticker := time.Tick(1 * time.Second)

for t := range ticker {
	fmt.Println(t)
}
```

You might see:

```text
12:00:01
12:00:02
12:00:03
12:00:04
...
```

`Tick` is a convenient wrapper around `NewTicker` that gives you only the channel. It returns `nil` for a non-positive duration.

For more control, use `NewTicker`.

---

# 18. `time.NewTicker()`

Creates a ticker that sends events periodically.

```go
ticker := time.NewTicker(1 * time.Second)
defer ticker.Stop()

for t := range ticker.C {
	fmt.Println("Tick:", t)
}
```

Think:

```text
NewTicker
    ↓
every 1 second
    ↓
channel
    ↓
goroutine receives event
```

Very useful for:

- Background jobs
- Monitoring
- Periodic database cleanup
- Health checks
- Metrics
- Polling

---

# 19. `Ticker.Stop()`

Stops a ticker.

```go
ticker := time.NewTicker(time.Second)

defer ticker.Stop()
```

After `Stop`, no more ticks are sent. The channel isn't closed by `Stop`.

---

# 20. `Ticker.Reset()`

Changes the ticker interval.

```go
ticker.Reset(5 * time.Second)
```

The new duration must be greater than zero.

---

# 21. `time.NewTimer()`

Creates a timer that fires once.

```go
timer := time.NewTimer(5 * time.Second)

<-timer.C

fmt.Println("Timer finished")
```

Difference:

```text
Timer
   ↓
fires once

Ticker
   ↓
fires repeatedly
```

---

# 22. `Timer.Stop()`

Stops a timer before it fires.

```go
timer := time.NewTimer(10 * time.Second)

if timer.Stop() {
	fmt.Println("Timer stopped")
}
```

`Stop()` reports whether the timer was successfully stopped before it fired.

---

# 23. `Timer.Reset()`

Resets a timer.

```go
timer := time.NewTimer(5 * time.Second)

timer.Reset(10 * time.Second)
```

This is useful when implementing reusable timeout mechanisms.

---

# 24. `time.AfterFunc()`

Executes a function after a duration.

```go
time.AfterFunc(5*time.Second, func() {
	fmt.Println("Five seconds passed")
})
```

Important difference:

```go
time.Sleep()
```

pauses a goroutine.

Whereas:

```go
time.AfterFunc()
```

schedules a function to run asynchronously.

---

# 25. `Duration` methods

## `Duration.Seconds()`

Converts duration to seconds as `float64`.

```go
d := 1500 * time.Millisecond

fmt.Println(d.Seconds())
```

Output:

```text
1.5
```

---

## `Duration.Minutes()`

```go
d := 90 * time.Second

fmt.Println(d.Minutes())
```

Output:

```text
1.5
```

---

## `Duration.Hours()`

```go
d := 90 * time.Minute

fmt.Println(d.Hours())
```

Output:

```text
1.5
```

---

## `Duration.Milliseconds()`

Returns the duration represented in milliseconds.

```go
d := 2 * time.Second

fmt.Println(d.Milliseconds())
```

Output:

```text
2000
```

---

## `Duration.Microseconds()`

```go
d := time.Second

fmt.Println(d.Microseconds())
```

Output:

```text
1000000
```

---

## `Duration.Nanoseconds()`

```go
d := time.Second

fmt.Println(d.Nanoseconds())
```

Output:

```text
1000000000
```

---

## `Duration.String()`

Converts the duration into a human-readable string.

```go
d := 2*time.Hour + 30*time.Minute

fmt.Println(d.String())
```

Output:

```text
2h30m0s
```

Usually you can simply do:

```go
fmt.Println(d)
```

---

## `Duration.Abs()`

Returns the absolute value.

```go
d := -5 * time.Second

fmt.Println(d.Abs())
```

Result:

```text
5s
```

---

## `Duration.Round()`

Rounds a duration to a specified unit.

```go
d := 90 * time.Millisecond

result := d.Round(100 * time.Millisecond)
```

Conceptually:

```text
90ms
 ↓
nearest 100ms
 ↓
100ms
```

---

## `Duration.Truncate()`

Truncates a duration to a specified unit.

```go
d := 1900 * time.Millisecond

result := d.Truncate(time.Second)
```

Result:

```text
1s
```

Difference:

```text
Round:
1900ms → 2000ms

Truncate:
1900ms → 1000ms
```

---

# 26. `Time` methods

## `t.Add()`

Adds a duration.

```go
now := time.Now()

future := now.Add(2 * time.Hour)
```

Subtract:

```go
past := now.Add(-2 * time.Hour)
```

---

# 27. `t.AddDate()`

Adds years, months and days.

```go
future := now.AddDate(1, 0, 0)
```

One year later.

```go
future := now.AddDate(0, 1, 0)
```

One month later.

```go
future := now.AddDate(0, 0, 7)
```

Seven days later.

### Important distinction

Use:

```go
Add()
```

for elapsed duration.

Use:

```go
AddDate()
```

for calendar calculations.

---

# 28. `t.Sub()`

Calculates the difference between two times.

```go
start := time.Now()

// operation

end := time.Now()

duration := end.Sub(start)
```

Result is a:

```go
time.Duration
```

---

# 29. `t.Before()`

Checks whether one time occurs before another.

```go
if start.Before(end) {
	fmt.Println("start is earlier")
}
```

Returns `bool`.

---

# 30. `t.After()`

Checks whether one time occurs after another.

```go
if deadline.After(now) {
	fmt.Println("deadline hasn't arrived")
}
```

---

# 31. `t.Equal()`

Checks whether two `Time` values represent the same instant.

```go
if t1.Equal(t2) {
	fmt.Println("Same instant")
}
```

### Very important beginner lesson

Prefer:

```go
t1.Equal(t2)
```

instead of:

```go
t1 == t2
```

The `==` operator also considers the location and monotonic-clock information, whereas `Equal` compares the represented instant appropriately.

---

# 32. `t.Compare()`

Compares two times.

```go
result := t1.Compare(t2)
```

Conceptually:

```text
-1 → t1 is before t2
 0 → t1 equals t2
+1 → t1 is after t2
```

---

# 33. `t.Year()`

```go
year := t.Year()
```

Returns the year.

---

# 34. `t.Month()`

```go
month := t.Month()
```

Returns a `time.Month`.

---

# 35. `t.Day()`

Returns the day of the month.

```go
day := t.Day()
```

---

# 36. `t.Hour()`

```go
hour := t.Hour()
```

Returns 0–23.

---

# 37. `t.Minute()`

```go
minute := t.Minute()
```

Returns 0–59.

---

# 38. `t.Second()`

```go
second := t.Second()
```

Returns 0–59.

---

# 39. `t.Nanosecond()`

```go
ns := t.Nanosecond()
```

Returns the nanosecond component.

---

# 40. `t.Date()`

Returns year, month, and day.

```go
year, month, day := t.Date()
```

Instead of calling:

```go
t.Year()
t.Month()
t.Day()
```

separately.

---

# 41. `t.Clock()`

Returns hour, minute, and second.

```go
hour, minute, second := t.Clock()
```

Useful when you only need the clock portion.

---

# 42. `t.Weekday()`

Returns the day of the week.

```go
day := t.Weekday()
```

Possible values:

```text
Sunday
Monday
Tuesday
Wednesday
Thursday
Friday
Saturday
```

---

# 43. `t.YearDay()`

Returns the day number within the year.

```go
day := t.YearDay()
```

For example:

```text
January 1 → 1
January 2 → 2
...
```

---

# 44. `t.ISOWeek()`

Returns the ISO week year and week number.

```go
year, week := t.ISOWeek()

fmt.Println(year, week)
```

Useful for reporting systems and weekly analytics.

---

# 45. `t.Format()`

Converts a `Time` into a string.

```go
now := time.Now()

formatted := now.Format("2006-01-02")

fmt.Println(formatted)
```

Example:

```text
2026-09-08
```

Another:

```go
formatted := now.Format("2006-01-02 15:04:05")
```

Result:

```text
2026-09-08 12:30:45
```

---

# 46. Go's unusual date layout

This is critical:

```go
"2006-01-02 15:04:05"
```

It is based on Go's reference date:

```text
January 2, 2006
15:04:05
```

Memorize:

```text
2006 → year
01   → month
02   → day
15   → 24-hour hour
04   → minute
05   → second
```

Examples:

```go
"2006-01-02"
"15:04:05"
"2006-01-02 15:04:05"
"02/01/2006"
```

---

# 47. Predefined layouts

Go provides many constants:

```go
time.RFC3339
time.RFC3339Nano
time.RFC1123
time.RFC1123Z
time.RFC822
time.RFC822Z
time.RFC850
time.ANSIC
time.UnixDate
time.RubyDate
time.Kitchen
time.DateTime
time.DateOnly
time.TimeOnly
```

For modern APIs, `RFC3339` is particularly useful:

```go
now.Format(time.RFC3339)
```

Example:

```text
2026-09-08T12:30:45+05:30
```

---

# 48. `t.String()`

Returns a standard string representation.

```go
fmt.Println(t.String())
```

Usually:

```go
fmt.Println(t)
```

is enough.

---

# 49. `t.GoString()`

Returns a Go-syntax representation useful for debugging.

```go
fmt.Printf("%#v\n", t)
```

This uses the `GoString()` representation.

---

# 50. `t.UTC()`

Converts the time's location representation to UTC.

```go
utc := t.UTC()
```

Important:

It does not change the actual instant. It changes how that instant is represented.

For example:

```text
India:
12:00 +05:30

UTC:
06:30 +00:00
```

Same instant.

---

# 51. `t.Local()`

Returns the time represented in the machine's local time zone.

```go
local := t.Local()
```

---

# 52. `t.In()`

Converts a time to a specified location.

```go
location, _ := time.LoadLocation("America/New_York")

newYorkTime := t.In(location)
```

This doesn't change the instant. It changes the time-zone representation.

---

# 53. `t.Location()`

Returns the associated location.

```go
location := t.Location()

fmt.Println(location)
```

---

# 54. `t.Zone()`

Returns the zone name and UTC offset.

```go
name, offset := t.Zone()
```

For example:

```text
IST
19800
```

The offset is in seconds east of UTC.

---

# 55. `t.ZoneBounds()`

Returns the start and end of the time-zone period containing the time.

```go
start, end := t.ZoneBounds()
```

This becomes useful when dealing with time-zone transitions and DST boundaries.

---

# 56. `t.IsDST()`

Reports whether the time is in Daylight Saving Time for its configured location.

```go
if t.IsDST() {
	fmt.Println("DST is active")
}
```

---

# 57. `t.IsZero()`

Checks whether the `Time` is the zero time.

```go
var t time.Time

if t.IsZero() {
	fmt.Println("Time not initialized")
}
```

This is very useful when optional timestamps exist.

---

# 58. `t.Round()`

Rounds a time to a duration boundary.

```go
rounded := t.Round(time.Hour)
```

For example, you can round timestamps to:

```text
second
minute
hour
```

---

# 59. `t.Truncate()`

Truncates a time to a duration boundary.

```go
truncated := t.Truncate(time.Hour)
```

Conceptually:

```text
12:47:35
    ↓
truncate to hour
    ↓
12:00:00
```

---

# 60. Unix timestamp methods

You can convert `time.Time` back to Unix timestamps.

## `t.Unix()`

```go
seconds := t.Unix()
```

Returns Unix seconds.

## `t.UnixMilli()`

```go
milliseconds := t.UnixMilli()
```

Returns Unix milliseconds.

## `t.UnixMicro()`

```go
microseconds := t.UnixMicro()
```

Returns Unix microseconds.

## `t.UnixNano()`

```go
nanoseconds := t.UnixNano()
```

Returns Unix nanoseconds.

---

# 61. Serialization methods

`time.Time` supports several encoding formats.

## `MarshalJSON()`

```go
data, err := t.MarshalJSON()
```

Useful when implementing JSON serialization manually.

Normally you don't need to call it yourself because `encoding/json` automatically uses it.

---

## `UnmarshalJSON()`

```go
err := t.UnmarshalJSON(data)
```

Used when decoding JSON manually.

---

## `MarshalText()`

Converts the time to text representation.

```go
data, err := t.MarshalText()
```

---

## `UnmarshalText()`

Reads a text representation back into `Time`.

```go
err := t.UnmarshalText(data)
```

---

## `MarshalBinary()`

Serializes a `Time` into binary representation.

```go
data, err := t.MarshalBinary()
```

---

## `UnmarshalBinary()`

Deserializes binary time data.

```go
err := t.UnmarshalBinary(data)
```

---

## `GobEncode()`

Encodes a time for Go's `encoding/gob`.

```go
data, err := t.GobEncode()
```

---

## `GobDecode()`

Decodes Gob data.

```go
err := t.GobDecode(data)
```

These serialization methods are particularly relevant when implementing custom serialization. Serialized `Time` values do not preserve the monotonic-clock reading.

---

# 62. `AppendText()`

Appends a textual representation to an existing byte slice.

```go
data, err := t.AppendText(buffer)
```

This can be useful when building output efficiently without creating as many intermediate strings.

---

# 63. `AppendBinary()`

Appends the binary representation to an existing byte slice.

```go
data, err := t.AppendBinary(buffer)
```

---

# 64. `AppendFormat()`

Formats the time and appends it to an existing byte slice.

```go
buffer := []byte("Time: ")

buffer = t.AppendFormat(
	buffer,
	"2006-01-02",
)
```

Result:

```text
Time: 2026-09-08
```

---

# 65. `time.Location`

A `Location` represents a time zone.

Examples:

```text
UTC
Asia/Kolkata
America/New_York
Europe/London
Asia/Tokyo
```

---

# 66. `time.LoadLocation()`

Loads a named time zone.

```go
location, err := time.LoadLocation("Asia/Kolkata")
```

Then:

```go
indiaTime := time.Now().In(location)
```

This is the normal approach when working with IANA time zones.

---

# 67. `time.FixedZone()`

Creates a location with a fixed UTC offset.

```go
location := time.FixedZone("IST", 5*60*60+30*60)
```

Then:

```go
t := time.Now().In(location)
```

Use this when you specifically need a fixed offset.

Be careful: a fixed offset is not the same thing as a real geographical time zone with historical/DST rules.

---

# 68. `time.LoadLocationFromTZData()`

Loads a time-zone `Location` directly from TZ database data supplied as bytes.

```go
location, err := time.LoadLocationFromTZData(
	"CustomZone",
	data,
)
```

This is mainly useful for advanced applications that package or supply their own time-zone data.

---

# 69. `Location.String()`

```go
fmt.Println(location.String())
```

Returns the location's name.

---

# 70. `time.Month`

`Month` represents a month.

Constants:

```go
time.January
time.February
time.March
time.April
time.May
time.June
time.July
time.August
time.September
time.October
time.November
time.December
```

Example:

```go
month := time.September

fmt.Println(month)
```

---

# 71. `Month.String()`

```go
fmt.Println(time.September.String())
```

Output:

```text
September
```

---

# 72. `time.Weekday`

Represents a day of the week.

Constants:

```go
time.Sunday
time.Monday
time.Tuesday
time.Wednesday
time.Thursday
time.Friday
time.Saturday
```

---

# 73. `Weekday.String()`

```go
fmt.Println(time.Monday.String())
```

Output:

```text
Monday
```

---

# 74. `time.Timer`

A `Timer` represents a single event that happens after a duration.

Typical pattern:

```go
timer := time.NewTimer(5 * time.Second)

<-timer.C

fmt.Println("Done")
```

The important field is:

```go
timer.C
```

which is a channel that receives the timer event.

---

# 75. `time.Ticker`

A `Ticker` represents repeated time events.

```go
ticker := time.NewTicker(time.Second)
defer ticker.Stop()

for {
	select {
	case t := <-ticker.C:
		fmt.Println(t)
	}
}
```

Think:

```text
Timer:
     ───────●

Ticker:
     ─●─●─●─●─●─●
```

---

# 76. `time.ParseError`

When parsing fails, Go can return a `*time.ParseError`.

Example:

```go
t, err := time.Parse(
	"2006-01-02",
	"wrong-date",
)

if err != nil {
	fmt.Println(err)
}
```

You can inspect the error as a `*time.ParseError` when you need more detailed parsing diagnostics.

---

# 77. A complete practical example

Here's a small program that uses many concepts together:

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	// Current time
	start := time.Now()

	fmt.Println("Start:", start)

	// Format
	fmt.Println(
		"Formatted:",
		start.Format("2006-01-02 15:04:05"),
	)

	// Add duration
	deadline := start.Add(10 * time.Second)

	fmt.Println("Deadline:", deadline)

	// Difference
	fmt.Println(
		"Remaining:",
		time.Until(deadline),
	)

	// Time zone
	location, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		panic(err)
	}

	fmt.Println(
		"India:",
		start.In(location),
	)

	// Sleep
	time.Sleep(2 * time.Second)

	// Measure elapsed time
	elapsed := time.Since(start)

	fmt.Println("Elapsed:", elapsed)
}
```

---

# 78. Three common beginner mistakes

## Mistake 1: Using the wrong format layout

Beginners often write:

```go
time.Now().Format("YYYY-MM-DD")
```

That's incorrect in Go.

Use:

```go
time.Now().Format("2006-01-02")
```

Remember:

```text
2006 → year
01   → month
02   → day
15   → hour
04   → minute
05   → second
```

---

## Mistake 2: Confusing `Duration` and `Time`

These are different:

```go
time.Time
```

means:

> When?

while:

```go
time.Duration
```

means:

> How long?

For example:

```go
deadline := time.Now().Add(30 * time.Second)
```

Here:

```text
time.Now()       → Time
30*time.Second   → Duration
deadline         → Time
```

---

## Mistake 3: Comparing times using `==`

Avoid blindly doing:

```go
if t1 == t2 {
}
```

Prefer:

```go
if t1.Equal(t2) {
}
```

because `==` considers more than just the represented instant, including location and monotonic-clock information.

---

# 79. Two real-world applications

## Application 1: API request timeout

Imagine your Go server calls another service.

You don't want your server to wait forever.

```go
select {
case result := <-response:
	fmt.Println(result)

case <-time.After(5 * time.Second):
	fmt.Println("Request timed out")
}
```

This is extremely common in:

```text
Microservices
REST APIs
Database calls
HTTP clients
External API calls
```

In production code, you'd often use `context.WithTimeout`, but `time` is fundamental to understanding the underlying timing concept.

---

# 80. Application 2: JWT/session expiration

Suppose a token expires one hour after it is created.

```go
createdAt := time.Now()

expiresAt := createdAt.Add(1 * time.Hour)
```

Then:

```go
if time.Now().After(expiresAt) {
	fmt.Println("Token expired")
}
```

The same idea applies to:

```text
Login sessions
Password reset links
OTP expiration
Cache expiration
Temporary URLs
Database records
Scheduled jobs
```

---

# 81. Exercises — progressively challenging

As requested, no solutions.

## Exercise 1 — Beginner: Digital Clock

Create a Go program that:

1. Gets the current time.
2. Prints the date in this format:

```text
YYYY-MM-DD
```

3. Prints the current time in:

```text
HH:MM:SS
```

4. Prints the current weekday.
5. Prints the current month.

### Goal

Practice:

```text
time.Now()
Format()
Year()
Month()
Day()
Weekday()
```

---

# Exercise 2 — Intermediate: Meeting Reminder

Build a program that:

1. Accepts a meeting time from the user in this format:

```text
2026-09-08 15:30:00
```

2. Parses the input into `time.Time`.
3. Compares it with the current time.
4. Prints:

```text
Meeting hasn't started yet.
```

or:

```text
Meeting has already started.
```

5. If the meeting hasn't started, calculate how much time remains.
6. Display the remaining duration.

### Goal

Practice:

```text
Scan input
Parse()
ParseInLocation()
Now()
Before()
After()
Sub()
Until()
Duration
```

---

# Exercise 3 — Advanced: Job Scheduler

Build a small job scheduler.

Your program should:

1. Accept a scheduled execution time.
2. Parse that time.
3. Calculate how long until the job should execute.
4. Wait until the scheduled time.
5. Execute a function when the time arrives.
6. Print the execution timestamp.
7. After the job executes, start a periodic task that runs every 10 seconds.
8. Allow the periodic task to be stopped gracefully.
9. Display how long each execution takes.
10. Handle invalid dates and times correctly.

Your scheduler should conceptually behave like:

```text
User enters schedule
       ↓
Parse time
       ↓
Calculate duration
       ↓
Wait
       ↓
Execute job
       ↓
Start periodic task
       ↓
Every 10 seconds
       ↓
Execute job
       ↓
Measure execution time
       ↓
Graceful shutdown
```

### Goal

This exercise combines:

```text
Time
Duration
Parse
ParseInLocation
Until
Timer
Ticker
Stop
Reset
Since
Goroutines
Channels
Select
Time zones
Error handling
```

---

# 82. The most important mental model

If you're learning Go seriously, remember this:

```text
                    time package
                         │
        ┌────────────────┼────────────────┐
        ↓                ↓                ↓
     time.Time      time.Duration      Location
        │                │                │
      "When?"         "How long?"       "Where?"
        │                │                │
        └────────────────┼────────────────┘
                         ↓
                    Timer / Ticker
                         │
                 "When should code run?"
```

The functions I'd recommend mastering first:

```text
time.Now()
time.Date()
time.Parse()
time.ParseInLocation()
time.ParseDuration()

time.Since()
time.Until()

time.Sleep()
time.After()

time.NewTimer()
time.AfterFunc()

time.NewTicker()
time.Tick()

time.LoadLocation()

t.Add()
t.AddDate()
t.Sub()

t.Before()
t.After()
t.Equal()
t.Compare()

t.Format()
t.In()
t.UTC()
t.Local()

t.Year()
t.Month()
t.Day()
t.Hour()
t.Minute()
t.Second()

t.Unix()
t.UnixMilli()
t.UnixMicro()
t.UnixNano()
```

Once these are comfortable, the rest of the API becomes much easier to understand.

---

# 83. Thought-provoking question

Imagine you are building a **global e-commerce application** where a customer in India places an order at the same moment a customer in New York does.

Both customers see their own local time, the database stores a timestamp, a payment service has a timeout, the order expires after 30 minutes, and a scheduled job runs at midnight for each customer's country.

**How would you design the system so that all of these different notions of "time" remain correct across time zones, daylight-saving changes, server clock changes, and distributed services?**

This takes you beyond simply learning `time.Now()` and into real-world distributed-system time handling, where the Go `time` package becomes particularly interesting.

---

## Official reference

Go `time` package documentation:

https://pkg.go.dev/time
