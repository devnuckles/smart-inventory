package service

import (
	"context"
	"encoding/json"

	"github.com/Tonmoy404/Smart-Inventory/logger"
	"github.com/Tonmoy404/Smart-Inventory/util"
)

type service struct {
	userRepo      UserRepo
	fileRepo      FileRepo
	errRepo       ErrorRepo
	cache         Cache
	authorization Authorization
}

func NewService(
	userRepo UserRepo,
	fileRepo FileRepo,
	errorRepo ErrorRepo,
	cache Cache,
	authorization Authorization,
) Service {
	return &service{
		userRepo:      userRepo,
		fileRepo:      fileRepo,
		errRepo:       errorRepo,
		cache:         cache,
		authorization: authorization,
	}
}

func (s *service) Error(ctx context.Context, code string, description string) *ErrorResponse {
	var errDetail *ErrorDetail

	// get from cache
	errString, err := s.cache.Get(code)
	if err != nil {
		logger.Error(ctx, "cannot get from redis", err)
	}
	if len(errString) > 0 {
		err = json.Unmarshal([]byte(errString), &errDetail)
		if err != nil {
			logger.Error(ctx, "cannot unmarshal error detail", err)
		}
	}

	// found in cache
	if errDetail != nil && len(errDetail.Code) == 0 {
		return &ErrorResponse{
			Timestamp:   util.GetCurrentTimestamp(),
			Description: description,
			Error:       errDetail,
		}
	}

	// not found in cache
	// get from db
	errDetail, err = s.errRepo.GetError(ctx, code)
	if err != nil {
		logger.Error(ctx, "cannot get from db", err)
		return &ErrorResponse{
			Timestamp:   util.GetCurrentTimestamp(),
			Description: description,
			Error: &ErrorDetail{
				Code:      code,
				MessageEn: "Not Set",
				MessageBn: "Not Set",
			},
		}
	}

	errResponse := &ErrorResponse{
		Timestamp:   util.GetCurrentTimestamp(),
		Description: description,
		Error:       errDetail,
	}

	return errResponse
}

func (s *service) Response(ctx context.Context, description string, data interface{}) *ResponseData {
	return &ResponseData{
		Timestamp:   util.GetCurrentTimestamp(),
		Description: description,
		Data:        data,
	}
}

func (s *service) IsPermitted(ctx context.Context, role, action, object string) bool {

	return s.authorization.IsPermitted(ctx, role, action, object)
}
