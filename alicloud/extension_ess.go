package alicloud

import "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"

type AdjustmentType string

const (
	QuantityChangeInCapacity = AdjustmentType("QuantityChangeInCapacity")
	PercentChangeInCapacity  = AdjustmentType("PercentChangeInCapacity")
	TotalCapacity            = AdjustmentType("TotalCapacity")
)

type RecurrenceType string

const (
	Daily   = RecurrenceType("Daily")
	Weekly  = RecurrenceType("Weekly")
	Monthly = RecurrenceType("Monthly")
)

type InstanceCreationType string

const (
	AutoCreated = InstanceCreationType("AutoCreated")
	Attached    = InstanceCreationType("Attached")
)

func EssCommonRequestInit(region string, code ServiceCode, domain CommonRequestDomain) *requests.CommonRequest {
	request := CommonRequestInit(region, code, domain)
	request.Version = ApiVersion20140828
	return request
}
