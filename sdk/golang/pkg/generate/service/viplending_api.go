// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

package service

import (
	"github.com/FuYuanDe2024/kucoin-universal-sdk/sdk/golang/internal/interfaces"
	"github.com/FuYuanDe2024/kucoin-universal-sdk/sdk/golang/pkg/generate/viplending/viplending"
)

type ViplendingService interface {

	// Get VIPLendingAPI API
	GetVIPLendingAPI() viplending.VIPLendingAPI
}

type ViplendingServiceImpl struct {
	VIPLendingAPI viplending.VIPLendingAPI
}

func NewViplendingService(transport interfaces.Transport) ViplendingService {
	api := &ViplendingServiceImpl{}
	api.VIPLendingAPI = viplending.NewVIPLendingAPIImp(transport)
	return api
}

func (impl *ViplendingServiceImpl) GetVIPLendingAPI() viplending.VIPLendingAPI {
	return impl.VIPLendingAPI
}
