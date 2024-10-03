package dbops

import "log"

func ReadVideoDeletionRecord(count int) ([]string, error) {
	//预编译
	stmtOut, err := dbConn.Prepare("SELECT video_id FROM video_deletion_record LIMIT ?")
	if err != nil {
		log.Printf("Error in ReadVideoDeletionRecord: %v", err)
		return nil, err
	}
	defer stmtOut.Close()
	var videoIDs []string
	rows, err := stmtOut.Query()
	if err != nil {
		log.Printf("ReadVideoDeletionRecord Query Error: %s", err)
		return nil, err
	}
	for rows.Next() {
		var videoID string
		if err = rows.Scan(&videoID); err != nil {
			log.Printf("ReadVideoDeletionRecord Scan Error: %s", err)
		}
		videoIDs = append(videoIDs, videoID)
	}

	return videoIDs, nil
}

func DelVideoDeletionRecord(vid string) error {
	//预编译
	stmtDel, err := dbConn.Prepare("DELETE FROM video_deletion_record WHERE video_id=?")
	if err != nil {
		log.Printf("Error in DelVideoDeletionRecord: %v", err)
		return err
	}
	defer stmtDel.Close()
	_, err = stmtDel.Exec(vid)
	if err != nil {
		log.Printf("Error in DelVideoDeletionRecord: %v", err)
	}
	return nil
}
