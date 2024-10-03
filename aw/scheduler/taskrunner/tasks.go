package taskrunner

import (
	"awesomeProject4/scheduler/dbops"
	"log"
	"os"
	"sync"
)

func deleteVideo(vid string) error {
	err := os.Remove(VIDEO_PATH + vid)
	if err != nil {
		log.Printf("delete video err:%v\n", err)
		return err
	}
	return nil
}

func VideoClearDispatcher(dc dataChan) error {
	res, err := dbops.ReadVideoDeletionRecord(3)
	if err != nil {
		log.Printf("read video deletion record err:%v", err)
		return err
	}
	if len(res) == 0 {
		log.Println("no video deletion record")
		return nil
	}
	for _, id := range res {
		dc <- id
	}
	return nil
}

func VideoClearExecutor(dc dataChan) error {
	errMap := &sync.Map{}
	var err error
forloop:
	for {
		select {
		case vid := <-dc:
			//删除视频源
			go func(id interface{}) {
				if err = deleteVideo(id.(string)); err != nil {
					errMap.Store(id, err)
				}
				//删除数据库video_del_rec中vid
				if err = dbops.DelVideoDeletionRecord(id.(string)); err != nil {
					errMap.Store(id, err)
				}
			}(vid)
		default:
			break forloop
		}
	}
	errMap.Range(func(key, value interface{}) bool {
		err = value.(error)
		if err != nil {
			return false
		}
		return true
	})
	return err
}
