package log_service

import (
	"fmt"
	"server/entity"
	"server/pkg/param"
	"strings"
	"time"
)

type Config struct {
	LogLanguage string `json:"log_language" default:"fa"`
}

type LogRepo interface {
	Insert(logEntry *entity.Log) error
	GetByTitle(title string) ([]*entity.Log, error)
	Gets(offset, limit int16) ([]*entity.Log, error)
	GetLicenseCheckByTitle(customerID, productID string) ([]*entity.Log, error)
	CountRecord() (int16, int16, int16)
	GetLogTemplateByKeyAndLanguage(key, language string) (*entity.LogTemplate, error)
}

type Service struct {
	repo LogRepo
	cfg  Config
}

func New(repo LogRepo, cfg Config) *Service {
	return &Service{repo: repo, cfg: cfg}
}

func (s *Service) CreateActivityLog(templateKey string, logParams []any) error {

	lo := new(entity.Log)

	var strLogParams []string
	for _, item := range logParams {
		strLogParams = append(strLogParams, fmt.Sprintf("%v", item))
	}

	// add params to template value
	lo.Title = templateKey
	lo.Message = strings.Join(strLogParams, "&")
	lo.CreatedAt = time.Now().Unix()

	// save log
	if err := s.repo.Insert(lo); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetActivityLogByTitle(req *param.GetActivityLogsByTitleRequest) ([]*entity.Log, error) {

	logs, err := s.repo.GetByTitle(req.Title)
	if err != nil {
		return nil, err
	}

	// get log template by key
	temp, err := s.repo.GetLogTemplateByKeyAndLanguage(req.Title, req.Language)
	if err != nil {
		return nil, err
	}

	if temp == nil {

		return nil, fmt.Errorf("error in load logs template")
	}

	result := make([]*entity.Log, 0, len(logs))
	for _, lo := range logs {

		logMsgStr := strings.Split(lo.Message, "&")
		anySlice := make([]any, len(logMsgStr))
		for i, v := range logMsgStr {
			anySlice[i] = v
		}

		lo.Message = fmt.Sprintf(temp.Value, anySlice...)

		result = append(result, lo)
	}

	return result, nil
}

func (s *Service) GetLicenseCheckActivityLogByTitle(req *param.GetLicenseCheckActivityLogsByTitleRequest) ([]*entity.Log, error) {

	logs, err := s.repo.GetLicenseCheckByTitle(req.CustomerID, req.ProductID)
	if err != nil {
		return nil, err
	}

	// get log template by key
	temp, err := s.repo.GetLogTemplateByKeyAndLanguage("checkLicense", req.Language)
	if err != nil {
		return nil, err
	}

	if temp == nil {

		return nil, fmt.Errorf("error in load logs template")
	}

	result := make([]*entity.Log, 0, len(logs))
	for _, lo := range logs {

		logMsgStr := strings.Split(lo.Message, "&")
		anySlice := make([]any, len(logMsgStr))
		for i, v := range logMsgStr {
			anySlice[i] = v
		}

		lo.Message = fmt.Sprintf(temp.Value, anySlice...)

		result = append(result, lo)
	}

	return result, nil
}

func (s *Service) GetActivityLogs(req *param.GetActivityLogsRequest) ([]*entity.Log, error) {

	logs, err := s.repo.Gets(req.Offset, req.Limit)
	if err != nil {
		return nil, err
	}

	result := make([]*entity.Log, 0, len(logs))
	for _, lo := range logs {

		// get log template by key
		temp, err := s.repo.GetLogTemplateByKeyAndLanguage(lo.Title, req.Language)
		if err != nil {
			return nil, err
		}

		if temp == nil {

			return nil, fmt.Errorf("error in load logs template")
		}

		logMsgStr := strings.Split(lo.Message, "&")
		anySlice := make([]any, len(logMsgStr))
		for i, v := range logMsgStr {
			anySlice[i] = v
		}

		lo.Message = fmt.Sprintf(temp.Value, anySlice...)

		result = append(result, lo)
	}

	return result, nil
}

func (s *Service) CountRecord() (int16, int16, int16) {
	return s.repo.CountRecord()
}
