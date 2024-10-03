package dbops

import "log"

func AddVideoDeletionRecord(vid string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO video_del_rec (video_id) values (?) ")
	if err != nil {
		log.Printf("stmtIns err : %v", err)
		return err
	}
	defer stmtIns.Close()
	_, err = stmtIns.Exec()
	if err != nil {
		log.Printf("stmtIns err : %v", err)
		return err
	}
	return nil
}
