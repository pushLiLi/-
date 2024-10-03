package dbops

import (
	"awesomeProject4/api/defs"
	"database/sql"
	"log"
	"strconv"
	"sync"
)

// 插入sessions
func InsertSessions(sessionsId string, ttl int64, userName string) error {
	ttlString := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare("INSERT INTO sessions (session_id, TTL ,login_name) VALUES (?, ?,?)")
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	_, err = stmtIns.Exec(sessionsId, ttlString, userName)
	if err != nil {
		return err
	}
	return nil
}

// 取session
func RetrieveSession(sessionsId string) (*defs.SimpleSession, error) {
	simpleSession := &defs.SimpleSession{}
	var ttl string
	var userName string
	stmtOut, err := dbConn.Prepare("SELECT TTL ,login_name FROM sessions WHERE session_id = ?")
	if err != nil {
		return nil, err
	}
	defer stmtOut.Close()
	err = stmtOut.QueryRow(sessionsId).Scan(&ttl, &userName)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if ttlRes, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		//代表没有问题的情况
		simpleSession.TTL = ttlRes
		simpleSession.Username = userName
	} else {
		return nil, err
	}

	return simpleSession, nil

}

// 取得所有的sessions
func RetrieveAllSessions() (*sync.Map, error) {
	SessionsMap := &sync.Map{}
	stmt, err := dbConn.Prepare("SELECT * FROM sessions")
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var session *defs.SimpleSession //单个session的实例化  ，，这里可能有问题 ===++
		var session_Id string
		var TTL string
		var login_name string
		if err := rows.Scan(&session_Id, &TTL, &login_name); err != nil {
			return nil, err
		}
		//转化ttl格式
		if ttl, err := strconv.ParseInt(TTL, 10, 64); err == nil {
			session.TTL = ttl
			session.Username = login_name
		}
		SessionsMap.Store(session_Id, session) //加到里面
	}
	return SessionsMap, nil
}

// 删除session
func DeleteSession(sessionId string) error {
	stmt, err := dbConn.Prepare("DELETE FROM sessions WHERE session_id = ?")
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(sessionId)
	if err != nil {
		return err
	}
	return nil

}

//3-16
