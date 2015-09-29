package mongo

import (
	"github.com/abhi11/gorest/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func DBSession() *mgo.Session {
	session, err := mgo.Dial("localhost")
	if err != nil {
		// log error
		panic(err)
	}
	return session
}

func DBGetLogs(query bson.M, caps map[string]int) model.LogMessages {
	session := DBSession()
	col := session.DB("test").C("logs")
	logs := model.LogMessages{}
	limit := caps["limit"]
	offset := caps["offset"]

	err := col.Find(query).Skip(offset).Limit(limit).All(&logs)

	if err != nil {
		panic(err)
	}

	return logs
}

func DBPostLog(logEntry model.LogMessage) int {
	session := DBSession()
	col := session.DB("test").C("logs")
	err := col.Insert(logEntry)

	if err != nil {
		//log error
		return 1
	}
	// log for success
	return 0
}

func DBPostLogsBatch(logEntries model.LogMessages) int {
	session := DBSession()
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
