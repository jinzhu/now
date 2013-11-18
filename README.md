## Now

Now is a time toolkit for golang

#### Why the project named `Now`?

```go
now.BeginningOfDay()
```
`now` is quite readable, aha?

#### But `now` is so common I can't search the project with my favorite search engine

* Star it in github [https://github.com/jinzhu/now](https://github.com/jinzhu/now)
* Search it with [http://godoc.org](http://godoc.org)

### Usage

```go
import "github.com/jinzhu/now"

time.Now() // 2013-11-18 17:51:49.123456789 Mon

now.BeginningOfMinute()   // 2013-11-18 17:51:00 Mon
now.BeginningOfHour()     // 2013-11-18 17:00:00 Mon
now.BeginningOfDay()      // 2013-11-18 00:00:00 Mon
now.BeginningOfWeek()     // 2013-11-17 00:00:00 Sun
now.FirstDayMonday = true // Set Monday as first day, default is Sunday
now.BeginningOfWeek()     // 2013-11-18 00:00:00 Mon
now.BeginningOfMonth()    // 2013-11-01 00:00:00 Fri
now.BeginningOfYear()     // 2013-01-01 00:00:00 Tue

now.EndOfMinute()         // 2013-11-18 17:51:59.999999999 Mon
now.EndOfHour()           // 2013-11-18 17:59:59.999999999 Mon
now.EndOfDay()            // 2013-11-18 23:59:59.999999999 Mon
now.EndOfWeek()           // 2013-11-23 23:59:59.999999999 Sat
now.FirstDayMonday = true // Set Monday as first day, default is Sunday
now.EndOfWeek()           // 2013-11-24 23:59:59.999999999 Sun
now.EndOfMonth()          // 2013-11-30 23:59:59.999999999 Sat
now.EndOfYear()           // 2013-12-31 23:59:59.999999999 Tue

// Use another time
t := time.Date(2013, 02, 18, 17, 51, 49, 123456789, time.Now().Location())
now.New(t).EndOfMonth()   // 2013-02-28 23:59:59.999999999 Thu
```
