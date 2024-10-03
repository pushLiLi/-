package session

import (
	"awesomeProject4/api/dbops"
	"awesomeProject4/api/defs"
	"awesomeProject4/api/utils"
	"sync"
	"time"
)

var sessionsMap *sync.Map

func NowTimeInMilliSecond() int64 {
	//获取当前时间，以毫秒的方式
	return time.Now().UnixNano() / 1000000
}

func init() {
	sessionsMap = &sync.Map{}
}

func LoadSessionsFromDB() {
	//RetrieveAllSessions的实现方法，在这里存
	res, err := dbops.RetrieveAllSessions()
	if err != nil {
		panic(err)
	}
	res.Range(func(key, value interface{}) bool {
		session := value.(*defs.SimpleSession)
		sessionsMap.Store(key, session)
		return true
	})
}

func GenerateNewSessionId(username string) string {
	//产生新的sessionId
	id, _ := utils.NewUUID()
	currentTime := NowTimeInMilliSecond()
	ttl := currentTime * 30 * 60 * 1000 //过期时间设置30min
	session := &defs.SimpleSession{Username: username, TTL: ttl}
	sessionsMap.Store(id, session)
	dbops.InsertSessions(id, ttl, username)

	return id
}

func deleteSession(id string) {
	//删除过期sessionId
	sessionsMap.Delete(id)
	dbops.DeleteSession(id)
}

func IsSessionExpired(sessionId string) (string, bool) {
	//返回userName和是否过期
	session, ok := sessionsMap.Load(sessionId)
	if ok {
		currentTime := NowTimeInMilliSecond()
		if currentTime > session.(*defs.SimpleSession).TTL {
			deleteSession(sessionId)
			//当前过期
			return "", true
		}
		return session.(*defs.SimpleSession).Username, false
	}
	return "", true
}
