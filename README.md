# goutils
Small utilities collection for golang

## logger
- DailyLogger

    NewDailyLogger creates a new Logger that logs to a file with name in the format `namef`
    in the given directory `dir`. It swiches to a new file everyday at given hour `h` and minute `m`.
    The `outs` variable denotes other destinations to which log data
    should be written apart from the file. It stops logging when `ctx` is closed.

    `fname` should be an accepted time format. Ex: "2006-Jan-02"

## sch
For some scheduled execution kind of stuff

- PingAt

    PingAt takes a time `hh`-`mm`-`ss`. It sends an empty struct to the channel `ch`
    when system time = given time. Repeats until `ctx` is closed.