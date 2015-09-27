### gorest
A simple rest api interface to query and post logs


## Current APIs:
* GET APIs
  * /logs : Fetches all logs
  * /logs/{timestamp} : Fetches all logs with the given timestamp
  * /logs/after/{timestamp} : Fetches all logs after the given timestamp (unix time)
  * /logs/before/{timestamp} : Fetches all logs before the given timestamp (unix time)
* POST APIs
  * /log/insert : Takes a single log entry and inserts db
  * /logs/batchinsert : Takes a list of log entries and inserts them in db

### Note
**Does not support de-duplication of log entries**

## Running on localhost
You would need mongodb to run gorest
```
$ mongod
```

```
$ go get github.com/abhi11/gorest
$ ./gorest
```
## TODO
* Better routing and refactoring
* Better erros handling
* write a script to start gorest [assuming mongodb is present]
* Adding docker support
