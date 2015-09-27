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

func DBGetLogs(query bson.M, limit int) LogMessages {
	session := DBGetSession()
	col := session.DB("test").C("logs")
	logs := LogMessages{}

	if limit < 0 {
		if err := col.Find(query).All(&logs); err != nil {
			panic(err)
		}
	} else {
		err := col.Find(query).Limit(limit).All(&logs)
		if err != nil {
			panic(err)
		}
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
