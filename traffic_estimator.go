package gads

import (
	"encoding/xml"
	"fmt"
)

// TrafficEstimatorService .
type TrafficEstimatorService struct {
	Auth
}

//TrafficEstimate .
type TrafficEstimate struct {
}

// NewTrafficEstimatorService .
func NewTrafficEstimatorService(auth *Auth) *TrafficEstimatorService {
	return &TrafficEstimatorService{Auth: *auth}
}

// Get .
func (s *TrafficEstimatorService) Get(selector Selector) (te []TrafficEstimate, totalCount int64, err error) {
	selector.XMLName = xml.Name{"", "serviceSelector"}
	respBody, err := s.Auth.request(
		adServiceUrl,
		"get",
		struct {
			XMLName xml.Name
			Sel     Selector
		}{
			XMLName: xml.Name{
				Space: baseUrl,
				Local: "get",
			},
			Sel: selector,
		},
	)
	if err != nil {
		return te, totalCount, err
	}
	fmt.Printf("%+v\n", respBody)
	return te, 0, err
}
