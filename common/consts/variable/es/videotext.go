package es

/*
"video_id":  msg.VideoID,
"title":     v.Title,
"user_id":   v.Userid,
"play_url":  v.Playurl,
"content":   v.Content.String,
"create_at": v.CreatedAt.Format("2006-01-02 15:04:05"),
*/
var (
	VideoTextIndex   = "video_text"
	VideoTextMapping = map[string]interface{}{
		"settings": map[string]interface{}{
			"number_of_shards":   1, // 设置分片数量
			"number_of_replicas": 0, // 设置副本数量
		},
		"mappings": map[string]interface{}{
			"properties": map[string]interface{}{
				"video_id": map[string]interface{}{
					"type": "integer",
				},
				"title": map[string]interface{}{
					"type": "text",
				},
				"user_id": map[string]interface{}{
					"type": "integer",
				},
				"play_url": map[string]interface{}{
					"type":  "keyword",
					"index": false,
				},
				"content": map[string]interface{}{
					"type":     "text",
					"analyzer": "standard",
				},
				"create_at": map[string]interface{}{
					"type":   "date",
					"format": "yyyy-MM-dd HH:mm:ss||epoch_millis",
				},
			},
		},
	}
)
