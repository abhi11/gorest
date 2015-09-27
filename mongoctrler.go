package main


import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func DBGetSession() *mgo.Session {
	session, err := mgo.Dial("localhost")
	if err != nil {
		// log error
		panic(err)
	}
	return session
}

func DBGetAllLogs() LogMessages {
	session := DBGetSession()
	col := session.DB("test").C("logs")
	logs := LogMessages{}

	err := col.Find(nil).All(&logs)

	if err != nil {
		panic(err)
	}

	return logs
}

func DBGetLogsBefore(timestamp int64) LogMessages {
	session := DBGetSession()
	col := session.DB("test").C("logs")
	logs := LogMessages{}

	err := col.Find(bson.M{"timestamp" : bson.M{"$lt" : timestamp}}).All(&logs)

	if err != nil {
		panic(err)
	}

	return logs
}

func DBGetLogsAfter(timestamp int64) LogMessages {
	session := DBGetSession()
	col := session.DB("test").C("logs")
	logs := LogMessages{}

	err := col.Find(bson.M{"timestamp" : bson.M{"$gt" : timestamp}}).All(&logs)

	if err != nil {
		panic(err)
	}

	return logs

}

func DBGetLogs(timestamp int64) LogMessages {
	session := DBGetSession()
	col := session.DB("test").C("logs")
	logs := LogMessages{}

	err := col.Find(bson.M{"timestamp" : timestamp}).All(&logs)

	if err != nil {
		panic(err)
	}

	return logs
}

func DBPostLog(logEntry LogMessage) int {
	session := DBGetSession()
	col := session.DB("test").C("logs")
	err := col.Insert(logEntry)

	if err != nil {
		//log error
		return 1
	}
	// log for success
	return 0
}

func DBPostLogsBatch(logEntries LogMessages) int {
	session := DBGetSession()
	col := session.DB("test").C("logs")

	for _, logEntry := range logEntries {
		err := col.Insert(logEntry)
		if err != nil {
			//log error
			return 1
		}
	}
	// log for success
	return 0
}
