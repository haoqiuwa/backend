package collect

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
	"wxcloudrun-golang/internal/pkg/model"
	"wxcloudrun-golang/internal/pkg/resp"
)

type Service struct {
	CollectDao    *model.Collect
	UserEventDao  *model.UserEvent
	SurveyDao     *model.Survey
	VideoDao      *model.Video
	VideoClipsDao *model.VideoClips
}

func NewService() *Service {
	return &Service{
		CollectDao:    &model.Collect{},
		UserEventDao:  &model.UserEvent{},
		SurveyDao:     &model.Survey{},
		VideoDao:      &model.Video{},
		VideoClipsDao: &model.VideoClips{},
	}
}

func (s *Service) ToggleCollectVideo(openID string, fileID string, picURL string, videoType int32) (*model.Collect,
	error) {
	// 查询是否已经收藏过
	collects, err := s.CollectDao.Gets(&model.Collect{OpenID: openID, FileID: fileID})
	fmt.Println(collects)
	if err != nil {
		return nil, err
	}
	if len(collects) > 0 {
		collect, err := s.CollectDao.Update(&model.Collect{ID: collects[0].ID, Status: collects[0].Status * (-1)})
		if err != nil {
			return nil, err
		}
		return collect, nil
	}
	// 创建收藏
	if videoType == 0 {
		videoType = 1
	}
	collect, err := s.CollectDao.Create(&model.Collect{
		OpenID:      openID,
		FileID:      fileID,
		PicURL:      picURL,
		VideoType:   videoType,
		Status:      1,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return collect, nil
}

func (s *Service) GetCollectByUser(userOpenID string, videoType int32) ([]model.Collect, error) {
	if videoType == 0 {
		videoType = 1
	}
	collects, err := s.CollectDao.Gets(&model.Collect{OpenID: userOpenID, Status: 1, VideoType: videoType})
	if err != nil {
		return nil, err
	}
	// order by created time desc
	sort.Slice(collects, func(i, j int) bool {
		return collects[i].CreatedTime.After(collects[j].CreatedTime)
	})
	return collects, nil
}

func (s *Service) CollectUserEvent(openID string, fileID string, eventType int32, fromPage string,
	videoType int32) (string,
	error) {
	data, err := s.UserEventDao.Create(&model.UserEvent{
		OpenID:    openID,
		FileID:    fileID,
		EventType: eventType,
		FromPage:  fromPage,
		VideoType: videoType,
	})
	if err != nil {
		return "", err
	}
	return data.FileID, nil
}

func (s *Service) CreateSurvey(openID string, content string) (*model.Survey, error) {
	data, err := s.SurveyDao.Create(&model.Survey{OpenID: openID, Content: content, CreatedTime: time.Now(), UpdatedTime: time.Now()})
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *Service) GetUserDownload(openID string) (int32, error) {
	data, err := s.UserEventDao.Gets(&model.UserEvent{OpenID: openID, EventType: 2})
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return int32(len(data)), nil
}

func (s *Service) GetUserDownloadStatus(openID string, fileID string) (bool, error) {
	log.Printf("GetUserDownloadStatus openID:%s ,fileID:%s", openID, fileID)
	data, err := s.UserEventDao.Gets(&model.UserEvent{OpenID: openID, FileID: fileID, EventType: 2})
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return len(data) > 0, nil
}

func (s *Service) GetUserDownloads(openID string, queryType string, page int32, pageSize int32) (resp.PageInfo, error) {
	offset := (page - 1) * pageSize
	pageInfo := resp.PageInfo{}
	pageInfo.Page = page
	data, err := s.UserEventDao.PageGets(openID, queryType, int(offset), int(pageSize))
	if err != nil {
		fmt.Println(err)
		return pageInfo, err
	}
	var result []model.Collect
	for _, v := range data {
		result = append(result, model.Collect{
			ID:          v.ID,
			OpenID:      v.OpenID,
			FileID:      v.FileID,
			PicURL:      s.findHoverImg(v.FileID, v.VideoType),
			VideoType:   v.VideoType,
			CreatedTime: v.CreatedTime,
			UpdatedTime: v.UpdatedTime,
		})
	}
	// order by created time desc
	sort.Slice(result, func(i, j int) bool {
		return result[i].CreatedTime.After(result[j].CreatedTime)
	})
	if len(result) == int(pageSize) {
		pageInfo.HasMore = true
	}
	pageInfo.PageData = result
	return pageInfo, nil
}

func (s *Service) findHoverImg(field string, videoType int32) string {
	if videoType == 6 || videoType == 7 {
		return ""
	}
	if videoType == 2 || videoType == 3 {
		v, r := s.VideoDao.Get(&model.Video{
			FilePath: field,
		})
		if r != nil {
			log.Println("findHoverImg err", r)
		}
		return v.HoverImgPath
	}
	if videoType == 4 || videoType == 5 {
		v, r := s.VideoClipsDao.Get(&model.VideoClips{
			FilePath: field,
		})
		if r != nil {
			log.Println("findHoverImg err", r)
		}
		return v.HoverImgPath
	}
	return ""

}

// get pic link by video link, video link like "highlight/2021/01/01/vxxx.
// mp4" and pic link like "highlight/2021/01/01/pxxx.PNG"
func videoToPicLink(videoLink string) string {
	// 分割视频链接，获取最后一部分
	parts := strings.Split(videoLink, "/")
	lastPart := parts[len(parts)-1]
	// 替换vxxx为pxxx
	lastPart = strings.Replace(lastPart, "v", "p", 1)
	// 替换.mp4为.PNG
	lastPart = strings.Replace(lastPart, ".MP4", ".png", 1)
	// 重新拼接链接
	parts[len(parts)-1] = lastPart
	picLink := strings.Join(parts, "/")
	return picLink
}
