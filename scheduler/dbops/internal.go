package dbops

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// ReadVideoDeletionRecord 从数据库中读取要删除的视频ID
func ReadVideoDeletionRecord(count int) ([]string, error) {
	var ids []string
	stmtOut, err := dbConn.Prepare("SELECT video_id FROM video_del_rec LIMIT ?")
	if err != nil {
		return ids, err
	}

	rows, err := stmtOut.Query(count)
	if err != nil {
		log.Printf("Query VideoDeletionRecord error: %v", err)
		return ids, err
	}

	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return ids, err
		}
		// 添加新写入的id到 ids 切片中
		ids = append(ids, id)
	}
	defer stmtOut.Close()
	return ids, nil
}

func DelVideoDeletionRecord(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_del_rec WHERE video_id=?")
	if err != nil {
		return err
	}

	_, err = stmtDel.Exec(vid)
	if err != nil {
		log.Printf("Deleting VideoDeletionRecord error: %v", err)
		return err
	}

	defer stmtDel.Close()
	return nil
}
