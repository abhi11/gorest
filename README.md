# gorest
A simple rest api interface to query and post logs


## Current APIs:
* GET APIs
  * /logs : Fetches all logs
  * /logs/twist_device/{id} : Fetches all logs for the given device id
  * /logs/mobile_device/{id} : Fetches all logs for the given mobile id
* POST APIs
  * /logs : Takes a single log entry and inserts db
  * /logs/batch_insert : Takes a list of log entries and inserts them in db

* Filters for GET APIs:
  * log_level : string value (eg Error, Debug, Fatal, Warn)
  * after : Unix time, for logs after the given time(more than equal to)
  * before : Unix time, for logs before the given time(less than)
  * limit : integer, limit the count of logs

**Note** : All the filters could be combined with each other.

**Pre-Caution** : If both after and before are used together, make sure before > after

## Usage
```
GET /logs?before=1235700&after=1235000&log_level=Fatal
```

```
GET /logs/twist_device/twist007?before=1235700&after=123500
```

**For POST use curl**

**For /logs**

```
curl -i -X POST -H 'Content-Type: application/json' -d '{"log":"New log message", "log_level":"Error","timestamp":1235000, "twist_device_id":"twist12"}' http://localhost:8080/logs
```

**For /logs/batch_insert**

```
curl -i -X POST -H 'Content-Type: application/json'-d '[{"log":"Another twist log", "log_level":"Debug","timestamp":1235700, "twist_device_id":"twist100"}, {"log":"Log message from nexus 6", "log_level":"Fatal","timestamp":1235600, "mobile_device_id":"nexus6"}]' http://localhost:8080/logs/batch_insert
```

### Note
**Does not support de-duplication of log entries**

**No Authentication support yet**

### Testing
* Use the logs.json file to import the data into mongodb

```
mongoimport --jsonArray --db test --collection logs --file logs.json
```
* The above helps in testing the GET APIs only
* For POST APIs use the curl command


## Running on localhost
You would need mongodb to run gorest.
And run it on the default port.

```
$ mongod
```

```
$ go get github.com/abhi11/gorest
$ ./gorest
```

## TODO
* Refactoring, make separate packages as appropriate
* Better error handling
* Add unit tests
* Add support to have customizable monogdb host
* Write a script to start gorest [assuming mongodb is present]
* Adding docker support
