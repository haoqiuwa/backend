package request

type EventReq struct {
	RequestId        string      `json:"request_id"`
	EventType        int32       `json:"event_type"`
	RequestTimestamp int64       `json:"request_timestamp"`
	EventTimestamp   int64       `json:"event_timestamp"`
	Data             interface{} `json:"data"`
}

// 场次主视频
type VideoEventReq struct {
	UUID           string `json:"uuid"`
	Court          int32  `json:"court"`
	FilePath       string `json:"file_path"`
	FileName       string `json:"file_name"`
	Hour           int32  `json:"hour"`
	Date           int32  `json:"date"`
	TeamAImgPath   string `json:"team_a_img_path"`
	TeamBImgPath   string `json:"team_b_img_path"`
	HoverImgPath   string `json:"hover_img_path"`
	StartTimestamp int64  `json:"start_timestamp"`
	EndTimestamp   int64  `json:"end_timestamp"`
	Time           int32  `json:"time"`
}

// 场次视频集锦/aigc视频
type VideoClipsEventReq struct {
	UUID         string `json:"uuid"`
	FilePath     string `json:"file_path"`
	VideoType    int32  `json:"vide_type"`
	HoverImgPath string `json:"hover_img_path"`
	Time         int32  `json:"time"`
	Team         string `json:"team"`
}

// 场次视频抽帧图片
type VideoImgEventReq struct {
	UUID         string `json:"uuid"`
	RelativeTime int32  `json:"relative_time"`
	FilePath     string `json:"file_path"`
}