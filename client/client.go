// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ProblemLevelGroupValue struct {
	ChildRuleRelation     *int32  `json:"childRuleRelation,omitempty" xml:"childRuleRelation,omitempty"`
	MatchCount            *int64  `json:"matchCount,omitempty" xml:"matchCount,omitempty"`
	TimeWindow            *int64  `json:"timeWindow,omitempty" xml:"timeWindow,omitempty"`
	TimeWindowUnit        *string `json:"timeWindowUnit,omitempty" xml:"timeWindowUnit,omitempty"`
	EnableUpgrade         *bool   `json:"enableUpgrade,omitempty" xml:"enableUpgrade,omitempty"`
	UpgradeTimeWindow     *int64  `json:"upgradeTimeWindow,omitempty" xml:"upgradeTimeWindow,omitempty"`
	UpgradeTimeWindowUnit *string `json:"upgradeTimeWindowUnit,omitempty" xml:"upgradeTimeWindowUnit,omitempty"`
}

func (s ProblemLevelGroupValue) String() string {
	return tea.Prettify(s)
}

func (s ProblemLevelGroupValue) GoString() string {
	return s.String()
}

func (s *ProblemLevelGroupValue) SetChildRuleRelation(v int32) *ProblemLevelGroupValue {
	s.ChildRuleRelation = &v
	return s
}

func (s *ProblemLevelGroupValue) SetMatchCount(v int64) *ProblemLevelGroupValue {
	s.MatchCount = &v
	return s
}

func (s *ProblemLevelGroupValue) SetTimeWindow(v int64) *ProblemLevelGroupValue {
	s.TimeWindow = &v
	return s
}

func (s *ProblemLevelGroupValue) SetTimeWindowUnit(v string) *ProblemLevelGroupValue {
	s.TimeWindowUnit = &v
	return s
}

func (s *ProblemLevelGroupValue) SetEnableUpgrade(v bool) *ProblemLevelGroupValue {
	s.EnableUpgrade = &v
	return s
}

func (s *ProblemLevelGroupValue) SetUpgradeTimeWindow(v int64) *ProblemLevelGroupValue {
	s.UpgradeTimeWindow = &v
	return s
}

func (s *ProblemLevelGroupValue) SetUpgradeTimeWindowUnit(v string) *ProblemLevelGroupValue {
	s.UpgradeTimeWindowUnit = &v
	return s
}

type DataProblemLevelGroupValue struct {
	ChildRuleRelation     *int32  `json:"childRuleRelation,omitempty" xml:"childRuleRelation,omitempty"`
	MatchCount            *int64  `json:"matchCount,omitempty" xml:"matchCount,omitempty"`
	TimeWindow            *int64  `json:"timeWindow,omitempty" xml:"timeWindow,omitempty"`
	TimeWindowUnit        *string `json:"timeWindowUnit,omitempty" xml:"timeWindowUnit,omitempty"`
	EnableUpgrade         *bool   `json:"enableUpgrade,omitempty" xml:"enableUpgrade,omitempty"`
	UpgradeTimeWindow     *int64  `json:"upgradeTimeWindow,omitempty" xml:"upgradeTimeWindow,omitempty"`
	UpgradeTimeWindowUnit *string `json:"upgradeTimeWindowUnit,omitempty" xml:"upgradeTimeWindowUnit,omitempty"`
}

func (s DataProblemLevelGroupValue) String() string {
	return tea.Prettify(s)
}

func (s DataProblemLevelGroupValue) GoString() string {
	return s.String()
}

func (s *DataProblemLevelGroupValue) SetChildRuleRelation(v int32) *DataProblemLevelGroupValue {
	s.ChildRuleRelation = &v
	return s
}

func (s *DataProblemLevelGroupValue) SetMatchCount(v int64) *DataProblemLevelGroupValue {
	s.MatchCount = &v
	return s
}

func (s *DataProblemLevelGroupValue) SetTimeWindow(v int64) *DataProblemLevelGroupValue {
	s.TimeWindow = &v
	return s
}

func (s *DataProblemLevelGroupValue) SetTimeWindowUnit(v string) *DataProblemLevelGroupValue {
	s.TimeWindowUnit = &v
	return s
}

func (s *DataProblemLevelGroupValue) SetEnableUpgrade(v bool) *DataProblemLevelGroupValue {
	s.EnableUpgrade = &v
	return s
}

func (s *DataProblemLevelGroupValue) SetUpgradeTimeWindow(v int64) *DataProblemLevelGroupValue {
	s.UpgradeTimeWindow = &v
	return s
}

func (s *DataProblemLevelGroupValue) SetUpgradeTimeWindowUnit(v string) *DataProblemLevelGroupValue {
	s.UpgradeTimeWindowUnit = &v
	return s
}

type DataValue struct {
	Code              *string `json:"code,omitempty" xml:"code,omitempty"`
	Description       *string `json:"description,omitempty" xml:"description,omitempty"`
	ConfigDescription *string `json:"configDescription,omitempty" xml:"configDescription,omitempty"`
	ConfigCode        *string `json:"configCode,omitempty" xml:"configCode,omitempty"`
	ParentCode        *string `json:"parentCode,omitempty" xml:"parentCode,omitempty"`
	ConfigKey         *string `json:"configKey,omitempty" xml:"configKey,omitempty"`
	ConfigValue       *string `json:"configValue,omitempty" xml:"configValue,omitempty"`
	Requirement       *bool   `json:"requirement,omitempty" xml:"requirement,omitempty"`
}

func (s DataValue) String() string {
	return tea.Prettify(s)
}

func (s DataValue) GoString() string {
	return s.String()
}

func (s *DataValue) SetCode(v string) *DataValue {
	s.Code = &v
	return s
}

func (s *DataValue) SetDescription(v string) *DataValue {
	s.Description = &v
	return s
}

func (s *DataValue) SetConfigDescription(v string) *DataValue {
	s.ConfigDescription = &v
	return s
}

func (s *DataValue) SetConfigCode(v string) *DataValue {
	s.ConfigCode = &v
	return s
}

func (s *DataValue) SetParentCode(v string) *DataValue {
	s.ParentCode = &v
	return s
}

func (s *DataValue) SetConfigKey(v string) *DataValue {
	s.ConfigKey = &v
	return s
}

func (s *DataValue) SetConfigValue(v string) *DataValue {
	s.ConfigValue = &v
	return s
}

func (s *DataValue) SetRequirement(v bool) *DataValue {
	s.Requirement = &v
	return s
}

type AddProblemServiceGroupRequest struct {
	ProblemId       *int64   `json:"problemId,omitempty" xml:"problemId,omitempty"`
	ServiceGroupIds []*int64 `json:"serviceGroupIds,omitempty" xml:"serviceGroupIds,omitempty" type:"Repeated"`
}

func (s AddProblemServiceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddProblemServiceGroupRequest) GoString() string {
	return s.String()
}

func (s *AddProblemServiceGroupRequest) SetProblemId(v int64) *AddProblemServiceGroupRequest {
	s.ProblemId = &v
	return s
}

func (s *AddProblemServiceGroupRequest) SetServiceGroupIds(v []*int64) *AddProblemServiceGroupRequest {
	s.ServiceGroupIds = v
	return s
}

type AddProblemServiceGroupResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AddProblemServiceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddProblemServiceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddProblemServiceGroupResponseBody) SetRequestId(v string) *AddProblemServiceGroupResponseBody {
	s.RequestId = &v
	return s
}

type AddProblemServiceGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddProblemServiceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddProblemServiceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddProblemServiceGroupResponse) GoString() string {
	return s.String()
}

func (s *AddProblemServiceGroupResponse) SetHeaders(v map[string]*string) *AddProblemServiceGroupResponse {
	s.Headers = v
	return s
}

func (s *AddProblemServiceGroupResponse) SetStatusCode(v int32) *AddProblemServiceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddProblemServiceGroupResponse) SetBody(v *AddProblemServiceGroupResponseBody) *AddProblemServiceGroupResponse {
	s.Body = v
	return s
}

type BillingStatisticsResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *BillingStatisticsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s BillingStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BillingStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *BillingStatisticsResponseBody) SetRequestId(v string) *BillingStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BillingStatisticsResponseBody) SetData(v *BillingStatisticsResponseBodyData) *BillingStatisticsResponseBody {
	s.Data = v
	return s
}

type BillingStatisticsResponseBodyData struct {
	AppUserCount                     *int64 `json:"appUserCount,omitempty" xml:"appUserCount,omitempty"`
	AppUserCountFree                 *int64 `json:"appUserCountFree,omitempty" xml:"appUserCountFree,omitempty"`
	EmailSend                        *int64 `json:"emailSend,omitempty" xml:"emailSend,omitempty"`
	EmailSendFree                    *int64 `json:"emailSendFree,omitempty" xml:"emailSendFree,omitempty"`
	EscalationPlanCount              *int64 `json:"escalationPlanCount,omitempty" xml:"escalationPlanCount,omitempty"`
	EscalationPlanCountFree          *int64 `json:"escalationPlanCountFree,omitempty" xml:"escalationPlanCountFree,omitempty"`
	EventReportApi                   *int64 `json:"eventReportApi,omitempty" xml:"eventReportApi,omitempty"`
	EventReportApiFree               *int64 `json:"eventReportApiFree,omitempty" xml:"eventReportApiFree,omitempty"`
	HasScheduleServiceGroupCount     *int64 `json:"hasScheduleServiceGroupCount,omitempty" xml:"hasScheduleServiceGroupCount,omitempty"`
	HasScheduleServiceGroupCountFree *int64 `json:"hasScheduleServiceGroupCountFree,omitempty" xml:"hasScheduleServiceGroupCountFree,omitempty"`
	ImMsgSend                        *int64 `json:"imMsgSend,omitempty" xml:"imMsgSend,omitempty"`
	ImMsgSendFree                    *int64 `json:"imMsgSendFree,omitempty" xml:"imMsgSendFree,omitempty"`
	RuleCount                        *int64 `json:"ruleCount,omitempty" xml:"ruleCount,omitempty"`
	RuleCountFree                    *int64 `json:"ruleCountFree,omitempty" xml:"ruleCountFree,omitempty"`
	SmsSend                          *int64 `json:"smsSend,omitempty" xml:"smsSend,omitempty"`
	SmsSendFree                      *int64 `json:"smsSendFree,omitempty" xml:"smsSendFree,omitempty"`
	SubscriptionNotifyCount          *int64 `json:"subscriptionNotifyCount,omitempty" xml:"subscriptionNotifyCount,omitempty"`
	SubscriptionNotifyCountFree      *int64 `json:"subscriptionNotifyCountFree,omitempty" xml:"subscriptionNotifyCountFree,omitempty"`
	Type                             *bool  `json:"type,omitempty" xml:"type,omitempty"`
	VoiceSend                        *int64 `json:"voiceSend,omitempty" xml:"voiceSend,omitempty"`
	VoiceSendFree                    *int64 `json:"voiceSendFree,omitempty" xml:"voiceSendFree,omitempty"`
}

func (s BillingStatisticsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BillingStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *BillingStatisticsResponseBodyData) SetAppUserCount(v int64) *BillingStatisticsResponseBodyData {
	s.AppUserCount = &v
	return s
}

func (s *BillingStatisticsResponseBodyData) SetAppUserCountFree(v int64) *BillingStatisticsResponseBodyData {
	s.AppUserCountFree = &v
	return s
}

func (s *BillingStatisticsResponseBodyData) SetEmailSend(v int64) *BillingStatisticsResponseBodyData {
	s.EmailSend = &v
	return s
}

func (s *BillingStatisticsResponseBodyData) SetEmailSendFree(v int64) *BillingStatisticsResponseBodyData {
	s.EmailSendFree = &v
	return s
}

func (s *BillingStatisticsResponseBodyData) SetEscalationPlanCount(v int64) *BillingStatisticsResponseBodyData {
	s.EscalationPlanCount = &v
	return s
}

func (s *BillingStatisticsResponseBodyData) SetEscalationPlanCountFree(v int64) *BillingStatisticsResponseBodyData {
	s.EscalationPlanCountFree = &v
	return s
}

func (s *BillingStatisticsResponseBodyData) SetEventReportApi(v int64) *BillingStatisticsResponseBodyData {
	s.EventReportApi = &v
	return s
}

func (s *BillingStatisticsResponseBodyData) SetEventReportApiFree(v int64) *BillingStatisticsResponseBodyData {
	s.EventReportApiFree = &v
	return s
}

func (s *BillingStatisticsResponseBodyData) SetHasScheduleServiceGroupCount(v int64) *BillingStatisticsResponseBodyData {
	s.HasScheduleServiceGroupCount = &v
	return s
}

func (s *BillingStatisticsResponseBodyData) SetHasScheduleServiceGroupCountFree(v int64) *BillingStatisticsResponseBodyData {
	s.HasScheduleServiceGroupCountFree = &v
	return s
}

func (s *BillingStatisticsResponseBodyData) SetImMsgSend(v int64) *BillingStatisticsResponseBodyData {
	s.ImMsgSend = &v
	return s
}

func (s *BillingStatisticsResponseBodyData) SetImMsgSendFree(v int64) *BillingStatisticsResponseBodyData {
	s.ImMsgSendFree = &v
	return s
}

func (s *BillingStatisticsResponseBodyData) SetRuleCount(v int64) *BillingStatisticsResponseBodyData {
	s.RuleCount = &v
	return s
}

func (s *BillingStatisticsResponseBodyData) SetRuleCountFree(v int64) *BillingStatisticsResponseBodyData {
	s.RuleCountFree = &v
	return s
}

func (s *BillingStatisticsResponseBodyData) SetSmsSend(v int64) *BillingStatisticsResponseBodyData {
	s.SmsSend = &v
	return s
}

func (s *BillingStatisticsResponseBodyData) SetSmsSendFree(v int64) *BillingStatisticsResponseBodyData {
	s.SmsSendFree = &v
	return s
}

func (s *BillingStatisticsResponseBodyData) SetSubscriptionNotifyCount(v int64) *BillingStatisticsResponseBodyData {
	s.SubscriptionNotifyCount = &v
	return s
}

func (s *BillingStatisticsResponseBodyData) SetSubscriptionNotifyCountFree(v int64) *BillingStatisticsResponseBodyData {
	s.SubscriptionNotifyCountFree = &v
	return s
}

func (s *BillingStatisticsResponseBodyData) SetType(v bool) *BillingStatisticsResponseBodyData {
	s.Type = &v
	return s
}

func (s *BillingStatisticsResponseBodyData) SetVoiceSend(v int64) *BillingStatisticsResponseBodyData {
	s.VoiceSend = &v
	return s
}

func (s *BillingStatisticsResponseBodyData) SetVoiceSendFree(v int64) *BillingStatisticsResponseBodyData {
	s.VoiceSendFree = &v
	return s
}

type BillingStatisticsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BillingStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BillingStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s BillingStatisticsResponse) GoString() string {
	return s.String()
}

func (s *BillingStatisticsResponse) SetHeaders(v map[string]*string) *BillingStatisticsResponse {
	s.Headers = v
	return s
}

func (s *BillingStatisticsResponse) SetStatusCode(v int32) *BillingStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *BillingStatisticsResponse) SetBody(v *BillingStatisticsResponseBody) *BillingStatisticsResponse {
	s.Body = v
	return s
}

type CancelProblemRequest struct {
	CancelReason            *int64  `json:"cancelReason,omitempty" xml:"cancelReason,omitempty"`
	CancelReasonDescription *string `json:"cancelReasonDescription,omitempty" xml:"cancelReasonDescription,omitempty"`
	ClientToken             *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	ProblemId               *int64  `json:"problemId,omitempty" xml:"problemId,omitempty"`
	ProblemNotifyType       *int64  `json:"problemNotifyType,omitempty" xml:"problemNotifyType,omitempty"`
}

func (s CancelProblemRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelProblemRequest) GoString() string {
	return s.String()
}

func (s *CancelProblemRequest) SetCancelReason(v int64) *CancelProblemRequest {
	s.CancelReason = &v
	return s
}

func (s *CancelProblemRequest) SetCancelReasonDescription(v string) *CancelProblemRequest {
	s.CancelReasonDescription = &v
	return s
}

func (s *CancelProblemRequest) SetClientToken(v string) *CancelProblemRequest {
	s.ClientToken = &v
	return s
}

func (s *CancelProblemRequest) SetProblemId(v int64) *CancelProblemRequest {
	s.ProblemId = &v
	return s
}

func (s *CancelProblemRequest) SetProblemNotifyType(v int64) *CancelProblemRequest {
	s.ProblemNotifyType = &v
	return s
}

type CancelProblemResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CancelProblemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelProblemResponseBody) GoString() string {
	return s.String()
}

func (s *CancelProblemResponseBody) SetRequestId(v string) *CancelProblemResponseBody {
	s.RequestId = &v
	return s
}

type CancelProblemResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelProblemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelProblemResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelProblemResponse) GoString() string {
	return s.String()
}

func (s *CancelProblemResponse) SetHeaders(v map[string]*string) *CancelProblemResponse {
	s.Headers = v
	return s
}

func (s *CancelProblemResponse) SetStatusCode(v int32) *CancelProblemResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelProblemResponse) SetBody(v *CancelProblemResponseBody) *CancelProblemResponse {
	s.Body = v
	return s
}

type CheckWebhookRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	Webhook     *string `json:"webhook,omitempty" xml:"webhook,omitempty"`
	WebhookType *string `json:"webhookType,omitempty" xml:"webhookType,omitempty"`
}

func (s CheckWebhookRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckWebhookRequest) GoString() string {
	return s.String()
}

func (s *CheckWebhookRequest) SetClientToken(v string) *CheckWebhookRequest {
	s.ClientToken = &v
	return s
}

func (s *CheckWebhookRequest) SetWebhook(v string) *CheckWebhookRequest {
	s.Webhook = &v
	return s
}

func (s *CheckWebhookRequest) SetWebhookType(v string) *CheckWebhookRequest {
	s.WebhookType = &v
	return s
}

type CheckWebhookResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CheckWebhookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *CheckWebhookResponseBody) SetRequestId(v string) *CheckWebhookResponseBody {
	s.RequestId = &v
	return s
}

type CheckWebhookResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckWebhookResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckWebhookResponse) GoString() string {
	return s.String()
}

func (s *CheckWebhookResponse) SetHeaders(v map[string]*string) *CheckWebhookResponse {
	s.Headers = v
	return s
}

func (s *CheckWebhookResponse) SetStatusCode(v int32) *CheckWebhookResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckWebhookResponse) SetBody(v *CheckWebhookResponseBody) *CheckWebhookResponse {
	s.Body = v
	return s
}

type ConfirmIntegrationConfigRequest struct {
	ClientToken         *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	IntegrationConfigId *int64  `json:"integrationConfigId,omitempty" xml:"integrationConfigId,omitempty"`
}

func (s ConfirmIntegrationConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfirmIntegrationConfigRequest) GoString() string {
	return s.String()
}

func (s *ConfirmIntegrationConfigRequest) SetClientToken(v string) *ConfirmIntegrationConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *ConfirmIntegrationConfigRequest) SetIntegrationConfigId(v int64) *ConfirmIntegrationConfigRequest {
	s.IntegrationConfigId = &v
	return s
}

type ConfirmIntegrationConfigResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ConfirmIntegrationConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfirmIntegrationConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ConfirmIntegrationConfigResponseBody) SetRequestId(v string) *ConfirmIntegrationConfigResponseBody {
	s.RequestId = &v
	return s
}

type ConfirmIntegrationConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ConfirmIntegrationConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfirmIntegrationConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfirmIntegrationConfigResponse) GoString() string {
	return s.String()
}

func (s *ConfirmIntegrationConfigResponse) SetHeaders(v map[string]*string) *ConfirmIntegrationConfigResponse {
	s.Headers = v
	return s
}

func (s *ConfirmIntegrationConfigResponse) SetStatusCode(v int32) *ConfirmIntegrationConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfirmIntegrationConfigResponse) SetBody(v *ConfirmIntegrationConfigResponseBody) *ConfirmIntegrationConfigResponse {
	s.Body = v
	return s
}

type CreateEscalationPlanRequest struct {
	// clientToken
	ClientToken                *string                                                  `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	EscalationPlanDescription  *string                                                  `json:"escalationPlanDescription,omitempty" xml:"escalationPlanDescription,omitempty"`
	EscalationPlanName         *string                                                  `json:"escalationPlanName,omitempty" xml:"escalationPlanName,omitempty"`
	EscalationPlanRules        []*CreateEscalationPlanRequestEscalationPlanRules        `json:"escalationPlanRules,omitempty" xml:"escalationPlanRules,omitempty" type:"Repeated"`
	EscalationPlanScopeObjects []*CreateEscalationPlanRequestEscalationPlanScopeObjects `json:"escalationPlanScopeObjects,omitempty" xml:"escalationPlanScopeObjects,omitempty" type:"Repeated"`
}

func (s CreateEscalationPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEscalationPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateEscalationPlanRequest) SetClientToken(v string) *CreateEscalationPlanRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEscalationPlanRequest) SetEscalationPlanDescription(v string) *CreateEscalationPlanRequest {
	s.EscalationPlanDescription = &v
	return s
}

func (s *CreateEscalationPlanRequest) SetEscalationPlanName(v string) *CreateEscalationPlanRequest {
	s.EscalationPlanName = &v
	return s
}

func (s *CreateEscalationPlanRequest) SetEscalationPlanRules(v []*CreateEscalationPlanRequestEscalationPlanRules) *CreateEscalationPlanRequest {
	s.EscalationPlanRules = v
	return s
}

func (s *CreateEscalationPlanRequest) SetEscalationPlanScopeObjects(v []*CreateEscalationPlanRequestEscalationPlanScopeObjects) *CreateEscalationPlanRequest {
	s.EscalationPlanScopeObjects = v
	return s
}

type CreateEscalationPlanRequestEscalationPlanRules struct {
	EscalationPlanConditions []*CreateEscalationPlanRequestEscalationPlanRulesEscalationPlanConditions `json:"escalationPlanConditions,omitempty" xml:"escalationPlanConditions,omitempty" type:"Repeated"`
	EscalationPlanStrategies []*CreateEscalationPlanRequestEscalationPlanRulesEscalationPlanStrategies `json:"escalationPlanStrategies,omitempty" xml:"escalationPlanStrategies,omitempty" type:"Repeated"`
	EscalationPlanType       *string                                                                   `json:"escalationPlanType,omitempty" xml:"escalationPlanType,omitempty"`
}

func (s CreateEscalationPlanRequestEscalationPlanRules) String() string {
	return tea.Prettify(s)
}

func (s CreateEscalationPlanRequestEscalationPlanRules) GoString() string {
	return s.String()
}

func (s *CreateEscalationPlanRequestEscalationPlanRules) SetEscalationPlanConditions(v []*CreateEscalationPlanRequestEscalationPlanRulesEscalationPlanConditions) *CreateEscalationPlanRequestEscalationPlanRules {
	s.EscalationPlanConditions = v
	return s
}

func (s *CreateEscalationPlanRequestEscalationPlanRules) SetEscalationPlanStrategies(v []*CreateEscalationPlanRequestEscalationPlanRulesEscalationPlanStrategies) *CreateEscalationPlanRequestEscalationPlanRules {
	s.EscalationPlanStrategies = v
	return s
}

func (s *CreateEscalationPlanRequestEscalationPlanRules) SetEscalationPlanType(v string) *CreateEscalationPlanRequestEscalationPlanRules {
	s.EscalationPlanType = &v
	return s
}

type CreateEscalationPlanRequestEscalationPlanRulesEscalationPlanConditions struct {
	Effection *string `json:"effection,omitempty" xml:"effection,omitempty"`
	Level     *string `json:"level,omitempty" xml:"level,omitempty"`
}

func (s CreateEscalationPlanRequestEscalationPlanRulesEscalationPlanConditions) String() string {
	return tea.Prettify(s)
}

func (s CreateEscalationPlanRequestEscalationPlanRulesEscalationPlanConditions) GoString() string {
	return s.String()
}

func (s *CreateEscalationPlanRequestEscalationPlanRulesEscalationPlanConditions) SetEffection(v string) *CreateEscalationPlanRequestEscalationPlanRulesEscalationPlanConditions {
	s.Effection = &v
	return s
}

func (s *CreateEscalationPlanRequestEscalationPlanRulesEscalationPlanConditions) SetLevel(v string) *CreateEscalationPlanRequestEscalationPlanRulesEscalationPlanConditions {
	s.Level = &v
	return s
}

type CreateEscalationPlanRequestEscalationPlanRulesEscalationPlanStrategies struct {
	EnableWebhook   *bool     `json:"enableWebhook,omitempty" xml:"enableWebhook,omitempty"`
	NoticeChannels  []*string `json:"noticeChannels,omitempty" xml:"noticeChannels,omitempty" type:"Repeated"`
	NoticeObjects   []*int64  `json:"noticeObjects,omitempty" xml:"noticeObjects,omitempty" type:"Repeated"`
	NoticeTime      *string   `json:"noticeTime,omitempty" xml:"noticeTime,omitempty"`
	ServiceGroupIds []*int64  `json:"serviceGroupIds,omitempty" xml:"serviceGroupIds,omitempty" type:"Repeated"`
}

func (s CreateEscalationPlanRequestEscalationPlanRulesEscalationPlanStrategies) String() string {
	return tea.Prettify(s)
}

func (s CreateEscalationPlanRequestEscalationPlanRulesEscalationPlanStrategies) GoString() string {
	return s.String()
}

func (s *CreateEscalationPlanRequestEscalationPlanRulesEscalationPlanStrategies) SetEnableWebhook(v bool) *CreateEscalationPlanRequestEscalationPlanRulesEscalationPlanStrategies {
	s.EnableWebhook = &v
	return s
}

func (s *CreateEscalationPlanRequestEscalationPlanRulesEscalationPlanStrategies) SetNoticeChannels(v []*string) *CreateEscalationPlanRequestEscalationPlanRulesEscalationPlanStrategies {
	s.NoticeChannels = v
	return s
}

func (s *CreateEscalationPlanRequestEscalationPlanRulesEscalationPlanStrategies) SetNoticeObjects(v []*int64) *CreateEscalationPlanRequestEscalationPlanRulesEscalationPlanStrategies {
	s.NoticeObjects = v
	return s
}

func (s *CreateEscalationPlanRequestEscalationPlanRulesEscalationPlanStrategies) SetNoticeTime(v string) *CreateEscalationPlanRequestEscalationPlanRulesEscalationPlanStrategies {
	s.NoticeTime = &v
	return s
}

func (s *CreateEscalationPlanRequestEscalationPlanRulesEscalationPlanStrategies) SetServiceGroupIds(v []*int64) *CreateEscalationPlanRequestEscalationPlanRulesEscalationPlanStrategies {
	s.ServiceGroupIds = v
	return s
}

type CreateEscalationPlanRequestEscalationPlanScopeObjects struct {
	Scope         *string `json:"scope,omitempty" xml:"scope,omitempty"`
	ScopeObjectId *int64  `json:"scopeObjectId,omitempty" xml:"scopeObjectId,omitempty"`
}

func (s CreateEscalationPlanRequestEscalationPlanScopeObjects) String() string {
	return tea.Prettify(s)
}

func (s CreateEscalationPlanRequestEscalationPlanScopeObjects) GoString() string {
	return s.String()
}

func (s *CreateEscalationPlanRequestEscalationPlanScopeObjects) SetScope(v string) *CreateEscalationPlanRequestEscalationPlanScopeObjects {
	s.Scope = &v
	return s
}

func (s *CreateEscalationPlanRequestEscalationPlanScopeObjects) SetScopeObjectId(v int64) *CreateEscalationPlanRequestEscalationPlanScopeObjects {
	s.ScopeObjectId = &v
	return s
}

type CreateEscalationPlanResponseBody struct {
	// data
	Data *CreateEscalationPlanResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateEscalationPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEscalationPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEscalationPlanResponseBody) SetData(v *CreateEscalationPlanResponseBodyData) *CreateEscalationPlanResponseBody {
	s.Data = v
	return s
}

func (s *CreateEscalationPlanResponseBody) SetRequestId(v string) *CreateEscalationPlanResponseBody {
	s.RequestId = &v
	return s
}

type CreateEscalationPlanResponseBodyData struct {
	EscalationPlanId *int64 `json:"escalationPlanId,omitempty" xml:"escalationPlanId,omitempty"`
}

func (s CreateEscalationPlanResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateEscalationPlanResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateEscalationPlanResponseBodyData) SetEscalationPlanId(v int64) *CreateEscalationPlanResponseBodyData {
	s.EscalationPlanId = &v
	return s
}

type CreateEscalationPlanResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateEscalationPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateEscalationPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEscalationPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateEscalationPlanResponse) SetHeaders(v map[string]*string) *CreateEscalationPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateEscalationPlanResponse) SetStatusCode(v int32) *CreateEscalationPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEscalationPlanResponse) SetBody(v *CreateEscalationPlanResponseBody) *CreateEscalationPlanResponse {
	s.Body = v
	return s
}

type CreateIncidentRequest struct {
	AssignUserId        *int64    `json:"assignUserId,omitempty" xml:"assignUserId,omitempty"`
	Channels            []*string `json:"channels,omitempty" xml:"channels,omitempty" type:"Repeated"`
	ClientToken         *string   `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	Effect              *string   `json:"effect,omitempty" xml:"effect,omitempty"`
	IncidentDescription *string   `json:"incidentDescription,omitempty" xml:"incidentDescription,omitempty"`
	IncidentLevel       *string   `json:"incidentLevel,omitempty" xml:"incidentLevel,omitempty"`
	IncidentTitle       *string   `json:"incidentTitle,omitempty" xml:"incidentTitle,omitempty"`
	RelatedServiceId    *int64    `json:"relatedServiceId,omitempty" xml:"relatedServiceId,omitempty"`
	// 12000
	ServiceGroupId *int64 `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
}

func (s CreateIncidentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIncidentRequest) GoString() string {
	return s.String()
}

func (s *CreateIncidentRequest) SetAssignUserId(v int64) *CreateIncidentRequest {
	s.AssignUserId = &v
	return s
}

func (s *CreateIncidentRequest) SetChannels(v []*string) *CreateIncidentRequest {
	s.Channels = v
	return s
}

func (s *CreateIncidentRequest) SetClientToken(v string) *CreateIncidentRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateIncidentRequest) SetEffect(v string) *CreateIncidentRequest {
	s.Effect = &v
	return s
}

func (s *CreateIncidentRequest) SetIncidentDescription(v string) *CreateIncidentRequest {
	s.IncidentDescription = &v
	return s
}

func (s *CreateIncidentRequest) SetIncidentLevel(v string) *CreateIncidentRequest {
	s.IncidentLevel = &v
	return s
}

func (s *CreateIncidentRequest) SetIncidentTitle(v string) *CreateIncidentRequest {
	s.IncidentTitle = &v
	return s
}

func (s *CreateIncidentRequest) SetRelatedServiceId(v int64) *CreateIncidentRequest {
	s.RelatedServiceId = &v
	return s
}

func (s *CreateIncidentRequest) SetServiceGroupId(v int64) *CreateIncidentRequest {
	s.ServiceGroupId = &v
	return s
}

type CreateIncidentResponseBody struct {
	// Id of the request
	Data *CreateIncidentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateIncidentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIncidentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIncidentResponseBody) SetData(v *CreateIncidentResponseBodyData) *CreateIncidentResponseBody {
	s.Data = v
	return s
}

func (s *CreateIncidentResponseBody) SetRequestId(v string) *CreateIncidentResponseBody {
	s.RequestId = &v
	return s
}

type CreateIncidentResponseBodyData struct {
	IncidentId *int64 `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
}

func (s CreateIncidentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateIncidentResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateIncidentResponseBodyData) SetIncidentId(v int64) *CreateIncidentResponseBodyData {
	s.IncidentId = &v
	return s
}

type CreateIncidentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateIncidentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateIncidentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIncidentResponse) GoString() string {
	return s.String()
}

func (s *CreateIncidentResponse) SetHeaders(v map[string]*string) *CreateIncidentResponse {
	s.Headers = v
	return s
}

func (s *CreateIncidentResponse) SetStatusCode(v int32) *CreateIncidentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIncidentResponse) SetBody(v *CreateIncidentResponseBody) *CreateIncidentResponse {
	s.Body = v
	return s
}

type CreateIncidentSubtotalRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	IncidentId  *int64  `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
}

func (s CreateIncidentSubtotalRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIncidentSubtotalRequest) GoString() string {
	return s.String()
}

func (s *CreateIncidentSubtotalRequest) SetClientToken(v string) *CreateIncidentSubtotalRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateIncidentSubtotalRequest) SetDescription(v string) *CreateIncidentSubtotalRequest {
	s.Description = &v
	return s
}

func (s *CreateIncidentSubtotalRequest) SetIncidentId(v int64) *CreateIncidentSubtotalRequest {
	s.IncidentId = &v
	return s
}

type CreateIncidentSubtotalResponseBody struct {
	// data
	Data *CreateIncidentSubtotalResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateIncidentSubtotalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIncidentSubtotalResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIncidentSubtotalResponseBody) SetData(v *CreateIncidentSubtotalResponseBodyData) *CreateIncidentSubtotalResponseBody {
	s.Data = v
	return s
}

func (s *CreateIncidentSubtotalResponseBody) SetRequestId(v string) *CreateIncidentSubtotalResponseBody {
	s.RequestId = &v
	return s
}

type CreateIncidentSubtotalResponseBodyData struct {
	SubtotalId *int64 `json:"subtotalId,omitempty" xml:"subtotalId,omitempty"`
}

func (s CreateIncidentSubtotalResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateIncidentSubtotalResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateIncidentSubtotalResponseBodyData) SetSubtotalId(v int64) *CreateIncidentSubtotalResponseBodyData {
	s.SubtotalId = &v
	return s
}

type CreateIncidentSubtotalResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateIncidentSubtotalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateIncidentSubtotalResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIncidentSubtotalResponse) GoString() string {
	return s.String()
}

func (s *CreateIncidentSubtotalResponse) SetHeaders(v map[string]*string) *CreateIncidentSubtotalResponse {
	s.Headers = v
	return s
}

func (s *CreateIncidentSubtotalResponse) SetStatusCode(v int32) *CreateIncidentSubtotalResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIncidentSubtotalResponse) SetBody(v *CreateIncidentSubtotalResponseBody) *CreateIncidentSubtotalResponse {
	s.Body = v
	return s
}

type CreateIntegrationConfigRequest struct {
	ClientToken     *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	MonitorSourceId *int64  `json:"monitorSourceId,omitempty" xml:"monitorSourceId,omitempty"`
}

func (s CreateIntegrationConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIntegrationConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateIntegrationConfigRequest) SetClientToken(v string) *CreateIntegrationConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateIntegrationConfigRequest) SetMonitorSourceId(v int64) *CreateIntegrationConfigRequest {
	s.MonitorSourceId = &v
	return s
}

type CreateIntegrationConfigResponseBody struct {
	// data
	Data *CreateIntegrationConfigResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateIntegrationConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIntegrationConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIntegrationConfigResponseBody) SetData(v *CreateIntegrationConfigResponseBodyData) *CreateIntegrationConfigResponseBody {
	s.Data = v
	return s
}

func (s *CreateIntegrationConfigResponseBody) SetRequestId(v string) *CreateIntegrationConfigResponseBody {
	s.RequestId = &v
	return s
}

type CreateIntegrationConfigResponseBodyData struct {
	IntegrationConfigId *int64 `json:"integrationConfigId,omitempty" xml:"integrationConfigId,omitempty"`
}

func (s CreateIntegrationConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateIntegrationConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateIntegrationConfigResponseBodyData) SetIntegrationConfigId(v int64) *CreateIntegrationConfigResponseBodyData {
	s.IntegrationConfigId = &v
	return s
}

type CreateIntegrationConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateIntegrationConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateIntegrationConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIntegrationConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateIntegrationConfigResponse) SetHeaders(v map[string]*string) *CreateIntegrationConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateIntegrationConfigResponse) SetStatusCode(v int32) *CreateIntegrationConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIntegrationConfigResponse) SetBody(v *CreateIntegrationConfigResponseBody) *CreateIntegrationConfigResponse {
	s.Body = v
	return s
}

type CreateProblemRequest struct {
	AffectServiceIds          []*int64 `json:"affectServiceIds,omitempty" xml:"affectServiceIds,omitempty" type:"Repeated"`
	ClientToken               *string  `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	DiscoverTime              *string  `json:"discoverTime,omitempty" xml:"discoverTime,omitempty"`
	IncidentId                *int64   `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
	MainHandlerId             *int64   `json:"mainHandlerId,omitempty" xml:"mainHandlerId,omitempty"`
	PreliminaryReason         *string  `json:"preliminaryReason,omitempty" xml:"preliminaryReason,omitempty"`
	ProblemLevel              *string  `json:"problemLevel,omitempty" xml:"problemLevel,omitempty"`
	ProblemName               *string  `json:"problemName,omitempty" xml:"problemName,omitempty"`
	ProblemNotifyType         *string  `json:"problemNotifyType,omitempty" xml:"problemNotifyType,omitempty"`
	ProblemStatus             *string  `json:"problemStatus,omitempty" xml:"problemStatus,omitempty"`
	ProgressSummary           *string  `json:"progressSummary,omitempty" xml:"progressSummary,omitempty"`
	ProgressSummaryRichTextId *int64   `json:"progressSummaryRichTextId,omitempty" xml:"progressSummaryRichTextId,omitempty"`
	RecoveryTime              *string  `json:"recoveryTime,omitempty" xml:"recoveryTime,omitempty"`
	RelatedServiceId          *int64   `json:"relatedServiceId,omitempty" xml:"relatedServiceId,omitempty"`
	ServiceGroupIds           []*int64 `json:"serviceGroupIds,omitempty" xml:"serviceGroupIds,omitempty" type:"Repeated"`
}

func (s CreateProblemRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProblemRequest) GoString() string {
	return s.String()
}

func (s *CreateProblemRequest) SetAffectServiceIds(v []*int64) *CreateProblemRequest {
	s.AffectServiceIds = v
	return s
}

func (s *CreateProblemRequest) SetClientToken(v string) *CreateProblemRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateProblemRequest) SetDiscoverTime(v string) *CreateProblemRequest {
	s.DiscoverTime = &v
	return s
}

func (s *CreateProblemRequest) SetIncidentId(v int64) *CreateProblemRequest {
	s.IncidentId = &v
	return s
}

func (s *CreateProblemRequest) SetMainHandlerId(v int64) *CreateProblemRequest {
	s.MainHandlerId = &v
	return s
}

func (s *CreateProblemRequest) SetPreliminaryReason(v string) *CreateProblemRequest {
	s.PreliminaryReason = &v
	return s
}

func (s *CreateProblemRequest) SetProblemLevel(v string) *CreateProblemRequest {
	s.ProblemLevel = &v
	return s
}

func (s *CreateProblemRequest) SetProblemName(v string) *CreateProblemRequest {
	s.ProblemName = &v
	return s
}

func (s *CreateProblemRequest) SetProblemNotifyType(v string) *CreateProblemRequest {
	s.ProblemNotifyType = &v
	return s
}

func (s *CreateProblemRequest) SetProblemStatus(v string) *CreateProblemRequest {
	s.ProblemStatus = &v
	return s
}

func (s *CreateProblemRequest) SetProgressSummary(v string) *CreateProblemRequest {
	s.ProgressSummary = &v
	return s
}

func (s *CreateProblemRequest) SetProgressSummaryRichTextId(v int64) *CreateProblemRequest {
	s.ProgressSummaryRichTextId = &v
	return s
}

func (s *CreateProblemRequest) SetRecoveryTime(v string) *CreateProblemRequest {
	s.RecoveryTime = &v
	return s
}

func (s *CreateProblemRequest) SetRelatedServiceId(v int64) *CreateProblemRequest {
	s.RelatedServiceId = &v
	return s
}

func (s *CreateProblemRequest) SetServiceGroupIds(v []*int64) *CreateProblemRequest {
	s.ServiceGroupIds = v
	return s
}

type CreateProblemResponseBody struct {
	// data
	Data *CreateProblemResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateProblemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProblemResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProblemResponseBody) SetData(v *CreateProblemResponseBodyData) *CreateProblemResponseBody {
	s.Data = v
	return s
}

func (s *CreateProblemResponseBody) SetRequestId(v string) *CreateProblemResponseBody {
	s.RequestId = &v
	return s
}

type CreateProblemResponseBodyData struct {
	ProblemId *int64 `json:"problemId,omitempty" xml:"problemId,omitempty"`
}

func (s CreateProblemResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateProblemResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateProblemResponseBodyData) SetProblemId(v int64) *CreateProblemResponseBodyData {
	s.ProblemId = &v
	return s
}

type CreateProblemResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateProblemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateProblemResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProblemResponse) GoString() string {
	return s.String()
}

func (s *CreateProblemResponse) SetHeaders(v map[string]*string) *CreateProblemResponse {
	s.Headers = v
	return s
}

func (s *CreateProblemResponse) SetStatusCode(v int32) *CreateProblemResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProblemResponse) SetBody(v *CreateProblemResponseBody) *CreateProblemResponse {
	s.Body = v
	return s
}

type CreateProblemEffectionServiceRequest struct {
	// clientToken
	ClientToken *string   `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	Description *string   `json:"description,omitempty" xml:"description,omitempty"`
	Level       *string   `json:"level,omitempty" xml:"level,omitempty"`
	PictureUrl  []*string `json:"pictureUrl,omitempty" xml:"pictureUrl,omitempty" type:"Repeated"`
	ProblemId   *int64    `json:"problemId,omitempty" xml:"problemId,omitempty"`
	ServiceId   *int64    `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	Status      *string   `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateProblemEffectionServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProblemEffectionServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateProblemEffectionServiceRequest) SetClientToken(v string) *CreateProblemEffectionServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateProblemEffectionServiceRequest) SetDescription(v string) *CreateProblemEffectionServiceRequest {
	s.Description = &v
	return s
}

func (s *CreateProblemEffectionServiceRequest) SetLevel(v string) *CreateProblemEffectionServiceRequest {
	s.Level = &v
	return s
}

func (s *CreateProblemEffectionServiceRequest) SetPictureUrl(v []*string) *CreateProblemEffectionServiceRequest {
	s.PictureUrl = v
	return s
}

func (s *CreateProblemEffectionServiceRequest) SetProblemId(v int64) *CreateProblemEffectionServiceRequest {
	s.ProblemId = &v
	return s
}

func (s *CreateProblemEffectionServiceRequest) SetServiceId(v int64) *CreateProblemEffectionServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateProblemEffectionServiceRequest) SetStatus(v string) *CreateProblemEffectionServiceRequest {
	s.Status = &v
	return s
}

type CreateProblemEffectionServiceResponseBody struct {
	// data
	Data *CreateProblemEffectionServiceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateProblemEffectionServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProblemEffectionServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProblemEffectionServiceResponseBody) SetData(v *CreateProblemEffectionServiceResponseBodyData) *CreateProblemEffectionServiceResponseBody {
	s.Data = v
	return s
}

func (s *CreateProblemEffectionServiceResponseBody) SetRequestId(v string) *CreateProblemEffectionServiceResponseBody {
	s.RequestId = &v
	return s
}

type CreateProblemEffectionServiceResponseBodyData struct {
	EffectionServiceId *int64 `json:"effectionServiceId,omitempty" xml:"effectionServiceId,omitempty"`
}

func (s CreateProblemEffectionServiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateProblemEffectionServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateProblemEffectionServiceResponseBodyData) SetEffectionServiceId(v int64) *CreateProblemEffectionServiceResponseBodyData {
	s.EffectionServiceId = &v
	return s
}

type CreateProblemEffectionServiceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateProblemEffectionServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateProblemEffectionServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProblemEffectionServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateProblemEffectionServiceResponse) SetHeaders(v map[string]*string) *CreateProblemEffectionServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateProblemEffectionServiceResponse) SetStatusCode(v int32) *CreateProblemEffectionServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProblemEffectionServiceResponse) SetBody(v *CreateProblemEffectionServiceResponseBody) *CreateProblemEffectionServiceResponse {
	s.Body = v
	return s
}

type CreateProblemMeasureRequest struct {
	CheckStandard  *string `json:"checkStandard,omitempty" xml:"checkStandard,omitempty"`
	CheckUserId    *int64  `json:"checkUserId,omitempty" xml:"checkUserId,omitempty"`
	ClientToken    *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	Content        *string `json:"content,omitempty" xml:"content,omitempty"`
	DirectorId     *int64  `json:"directorId,omitempty" xml:"directorId,omitempty"`
	PlanFinishTime *string `json:"planFinishTime,omitempty" xml:"planFinishTime,omitempty"`
	ProblemId      *int64  `json:"problemId,omitempty" xml:"problemId,omitempty"`
	StalkerId      *int64  `json:"stalkerId,omitempty" xml:"stalkerId,omitempty"`
	Status         *string `json:"status,omitempty" xml:"status,omitempty"`
	Type           *int32  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateProblemMeasureRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProblemMeasureRequest) GoString() string {
	return s.String()
}

func (s *CreateProblemMeasureRequest) SetCheckStandard(v string) *CreateProblemMeasureRequest {
	s.CheckStandard = &v
	return s
}

func (s *CreateProblemMeasureRequest) SetCheckUserId(v int64) *CreateProblemMeasureRequest {
	s.CheckUserId = &v
	return s
}

func (s *CreateProblemMeasureRequest) SetClientToken(v string) *CreateProblemMeasureRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateProblemMeasureRequest) SetContent(v string) *CreateProblemMeasureRequest {
	s.Content = &v
	return s
}

func (s *CreateProblemMeasureRequest) SetDirectorId(v int64) *CreateProblemMeasureRequest {
	s.DirectorId = &v
	return s
}

func (s *CreateProblemMeasureRequest) SetPlanFinishTime(v string) *CreateProblemMeasureRequest {
	s.PlanFinishTime = &v
	return s
}

func (s *CreateProblemMeasureRequest) SetProblemId(v int64) *CreateProblemMeasureRequest {
	s.ProblemId = &v
	return s
}

func (s *CreateProblemMeasureRequest) SetStalkerId(v int64) *CreateProblemMeasureRequest {
	s.StalkerId = &v
	return s
}

func (s *CreateProblemMeasureRequest) SetStatus(v string) *CreateProblemMeasureRequest {
	s.Status = &v
	return s
}

func (s *CreateProblemMeasureRequest) SetType(v int32) *CreateProblemMeasureRequest {
	s.Type = &v
	return s
}

type CreateProblemMeasureResponseBody struct {
	// data
	Data *CreateProblemMeasureResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateProblemMeasureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProblemMeasureResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProblemMeasureResponseBody) SetData(v *CreateProblemMeasureResponseBodyData) *CreateProblemMeasureResponseBody {
	s.Data = v
	return s
}

func (s *CreateProblemMeasureResponseBody) SetRequestId(v string) *CreateProblemMeasureResponseBody {
	s.RequestId = &v
	return s
}

type CreateProblemMeasureResponseBodyData struct {
	MeasureId *int64 `json:"measureId,omitempty" xml:"measureId,omitempty"`
}

func (s CreateProblemMeasureResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateProblemMeasureResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateProblemMeasureResponseBodyData) SetMeasureId(v int64) *CreateProblemMeasureResponseBodyData {
	s.MeasureId = &v
	return s
}

type CreateProblemMeasureResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateProblemMeasureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateProblemMeasureResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProblemMeasureResponse) GoString() string {
	return s.String()
}

func (s *CreateProblemMeasureResponse) SetHeaders(v map[string]*string) *CreateProblemMeasureResponse {
	s.Headers = v
	return s
}

func (s *CreateProblemMeasureResponse) SetStatusCode(v int32) *CreateProblemMeasureResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProblemMeasureResponse) SetBody(v *CreateProblemMeasureResponseBody) *CreateProblemMeasureResponse {
	s.Body = v
	return s
}

type CreateProblemSubtotalRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	ProblemId   *int64  `json:"problemId,omitempty" xml:"problemId,omitempty"`
}

func (s CreateProblemSubtotalRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProblemSubtotalRequest) GoString() string {
	return s.String()
}

func (s *CreateProblemSubtotalRequest) SetClientToken(v string) *CreateProblemSubtotalRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateProblemSubtotalRequest) SetDescription(v string) *CreateProblemSubtotalRequest {
	s.Description = &v
	return s
}

func (s *CreateProblemSubtotalRequest) SetProblemId(v int64) *CreateProblemSubtotalRequest {
	s.ProblemId = &v
	return s
}

type CreateProblemSubtotalResponseBody struct {
	// object
	Data *CreateProblemSubtotalResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateProblemSubtotalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProblemSubtotalResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProblemSubtotalResponseBody) SetData(v *CreateProblemSubtotalResponseBodyData) *CreateProblemSubtotalResponseBody {
	s.Data = v
	return s
}

func (s *CreateProblemSubtotalResponseBody) SetRequestId(v string) *CreateProblemSubtotalResponseBody {
	s.RequestId = &v
	return s
}

type CreateProblemSubtotalResponseBodyData struct {
	SubtotalId *int64 `json:"subtotalId,omitempty" xml:"subtotalId,omitempty"`
}

func (s CreateProblemSubtotalResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateProblemSubtotalResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateProblemSubtotalResponseBodyData) SetSubtotalId(v int64) *CreateProblemSubtotalResponseBodyData {
	s.SubtotalId = &v
	return s
}

type CreateProblemSubtotalResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateProblemSubtotalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateProblemSubtotalResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProblemSubtotalResponse) GoString() string {
	return s.String()
}

func (s *CreateProblemSubtotalResponse) SetHeaders(v map[string]*string) *CreateProblemSubtotalResponse {
	s.Headers = v
	return s
}

func (s *CreateProblemSubtotalResponse) SetStatusCode(v int32) *CreateProblemSubtotalResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProblemSubtotalResponse) SetBody(v *CreateProblemSubtotalResponseBody) *CreateProblemSubtotalResponse {
	s.Body = v
	return s
}

type CreateProblemTimelineRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	Content     *string `json:"content,omitempty" xml:"content,omitempty"`
	KeyNode     *string `json:"keyNode,omitempty" xml:"keyNode,omitempty"`
	ProblemId   *int64  `json:"problemId,omitempty" xml:"problemId,omitempty"`
	Time        *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s CreateProblemTimelineRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProblemTimelineRequest) GoString() string {
	return s.String()
}

func (s *CreateProblemTimelineRequest) SetClientToken(v string) *CreateProblemTimelineRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateProblemTimelineRequest) SetContent(v string) *CreateProblemTimelineRequest {
	s.Content = &v
	return s
}

func (s *CreateProblemTimelineRequest) SetKeyNode(v string) *CreateProblemTimelineRequest {
	s.KeyNode = &v
	return s
}

func (s *CreateProblemTimelineRequest) SetProblemId(v int64) *CreateProblemTimelineRequest {
	s.ProblemId = &v
	return s
}

func (s *CreateProblemTimelineRequest) SetTime(v string) *CreateProblemTimelineRequest {
	s.Time = &v
	return s
}

type CreateProblemTimelineResponseBody struct {
	// data
	Data *CreateProblemTimelineResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateProblemTimelineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProblemTimelineResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProblemTimelineResponseBody) SetData(v *CreateProblemTimelineResponseBodyData) *CreateProblemTimelineResponseBody {
	s.Data = v
	return s
}

func (s *CreateProblemTimelineResponseBody) SetRequestId(v string) *CreateProblemTimelineResponseBody {
	s.RequestId = &v
	return s
}

type CreateProblemTimelineResponseBodyData struct {
	ProblemTimelineId *int64 `json:"problemTimelineId,omitempty" xml:"problemTimelineId,omitempty"`
}

func (s CreateProblemTimelineResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateProblemTimelineResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateProblemTimelineResponseBodyData) SetProblemTimelineId(v int64) *CreateProblemTimelineResponseBodyData {
	s.ProblemTimelineId = &v
	return s
}

type CreateProblemTimelineResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateProblemTimelineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateProblemTimelineResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProblemTimelineResponse) GoString() string {
	return s.String()
}

func (s *CreateProblemTimelineResponse) SetHeaders(v map[string]*string) *CreateProblemTimelineResponse {
	s.Headers = v
	return s
}

func (s *CreateProblemTimelineResponse) SetStatusCode(v int32) *CreateProblemTimelineResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProblemTimelineResponse) SetBody(v *CreateProblemTimelineResponseBody) *CreateProblemTimelineResponse {
	s.Body = v
	return s
}

type CreateProblemTimelinesRequest struct {
	// clientToken
	ClientToken   *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	ProblemId     *int64  `json:"problemId,omitempty" xml:"problemId,omitempty"`
	TimelineNodes *string `json:"timelineNodes,omitempty" xml:"timelineNodes,omitempty"`
}

func (s CreateProblemTimelinesRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProblemTimelinesRequest) GoString() string {
	return s.String()
}

func (s *CreateProblemTimelinesRequest) SetClientToken(v string) *CreateProblemTimelinesRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateProblemTimelinesRequest) SetProblemId(v int64) *CreateProblemTimelinesRequest {
	s.ProblemId = &v
	return s
}

func (s *CreateProblemTimelinesRequest) SetTimelineNodes(v string) *CreateProblemTimelinesRequest {
	s.TimelineNodes = &v
	return s
}

type CreateProblemTimelinesResponseBody struct {
	// data
	Data *CreateProblemTimelinesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateProblemTimelinesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProblemTimelinesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProblemTimelinesResponseBody) SetData(v *CreateProblemTimelinesResponseBodyData) *CreateProblemTimelinesResponseBody {
	s.Data = v
	return s
}

func (s *CreateProblemTimelinesResponseBody) SetRequestId(v string) *CreateProblemTimelinesResponseBody {
	s.RequestId = &v
	return s
}

type CreateProblemTimelinesResponseBodyData struct {
	ProblemTimelineIds []*int64 `json:"problemTimelineIds,omitempty" xml:"problemTimelineIds,omitempty" type:"Repeated"`
}

func (s CreateProblemTimelinesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateProblemTimelinesResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateProblemTimelinesResponseBodyData) SetProblemTimelineIds(v []*int64) *CreateProblemTimelinesResponseBodyData {
	s.ProblemTimelineIds = v
	return s
}

type CreateProblemTimelinesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateProblemTimelinesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateProblemTimelinesResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProblemTimelinesResponse) GoString() string {
	return s.String()
}

func (s *CreateProblemTimelinesResponse) SetHeaders(v map[string]*string) *CreateProblemTimelinesResponse {
	s.Headers = v
	return s
}

func (s *CreateProblemTimelinesResponse) SetStatusCode(v int32) *CreateProblemTimelinesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProblemTimelinesResponse) SetBody(v *CreateProblemTimelinesResponseBody) *CreateProblemTimelinesResponse {
	s.Body = v
	return s
}

type CreateRichTextRequest struct {
	InstanceId   *int64  `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	RichText     *string `json:"richText,omitempty" xml:"richText,omitempty"`
}

func (s CreateRichTextRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRichTextRequest) GoString() string {
	return s.String()
}

func (s *CreateRichTextRequest) SetInstanceId(v int64) *CreateRichTextRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateRichTextRequest) SetInstanceType(v string) *CreateRichTextRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateRichTextRequest) SetRichText(v string) *CreateRichTextRequest {
	s.RichText = &v
	return s
}

type CreateRichTextResponseBody struct {
	// data
	Data *CreateRichTextResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateRichTextResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRichTextResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRichTextResponseBody) SetData(v *CreateRichTextResponseBodyData) *CreateRichTextResponseBody {
	s.Data = v
	return s
}

func (s *CreateRichTextResponseBody) SetRequestId(v string) *CreateRichTextResponseBody {
	s.RequestId = &v
	return s
}

type CreateRichTextResponseBodyData struct {
	InstanceId   *int64  `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	InstanceType *int64  `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	RichText     *string `json:"richText,omitempty" xml:"richText,omitempty"`
}

func (s CreateRichTextResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateRichTextResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateRichTextResponseBodyData) SetInstanceId(v int64) *CreateRichTextResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *CreateRichTextResponseBodyData) SetInstanceType(v int64) *CreateRichTextResponseBodyData {
	s.InstanceType = &v
	return s
}

func (s *CreateRichTextResponseBodyData) SetRichText(v string) *CreateRichTextResponseBodyData {
	s.RichText = &v
	return s
}

type CreateRichTextResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRichTextResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRichTextResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRichTextResponse) GoString() string {
	return s.String()
}

func (s *CreateRichTextResponse) SetHeaders(v map[string]*string) *CreateRichTextResponse {
	s.Headers = v
	return s
}

func (s *CreateRichTextResponse) SetStatusCode(v int32) *CreateRichTextResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRichTextResponse) SetBody(v *CreateRichTextResponseBody) *CreateRichTextResponse {
	s.Body = v
	return s
}

type CreateRouteRuleRequest struct {
	AssignObjectId           *int64                                   `json:"assignObjectId,omitempty" xml:"assignObjectId,omitempty"`
	AssignObjectType         *string                                  `json:"assignObjectType,omitempty" xml:"assignObjectType,omitempty"`
	ChildRuleRelation        *string                                  `json:"childRuleRelation,omitempty" xml:"childRuleRelation,omitempty"`
	ClientToken              *string                                  `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	CoverageProblemLevels    []*string                                `json:"coverageProblemLevels,omitempty" xml:"coverageProblemLevels,omitempty" type:"Repeated"`
	Effection                *string                                  `json:"effection,omitempty" xml:"effection,omitempty"`
	EnableStatus             *string                                  `json:"enableStatus,omitempty" xml:"enableStatus,omitempty"`
	IncidentLevel            *string                                  `json:"incidentLevel,omitempty" xml:"incidentLevel,omitempty"`
	MatchCount               *int32                                   `json:"matchCount,omitempty" xml:"matchCount,omitempty"`
	NotifyChannels           []*string                                `json:"notifyChannels,omitempty" xml:"notifyChannels,omitempty" type:"Repeated"`
	ProblemEffectionServices []*int64                                 `json:"problemEffectionServices,omitempty" xml:"problemEffectionServices,omitempty" type:"Repeated"`
	ProblemLevelGroup        map[string]*ProblemLevelGroupValue       `json:"problemLevelGroup,omitempty" xml:"problemLevelGroup,omitempty"`
	RelatedServiceId         *int64                                   `json:"relatedServiceId,omitempty" xml:"relatedServiceId,omitempty"`
	RouteChildRules          []*CreateRouteRuleRequestRouteChildRules `json:"routeChildRules,omitempty" xml:"routeChildRules,omitempty" type:"Repeated"`
	RouteType                *string                                  `json:"routeType,omitempty" xml:"routeType,omitempty"`
	RuleName                 *string                                  `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	TimeWindow               *int64                                   `json:"timeWindow,omitempty" xml:"timeWindow,omitempty"`
	TimeWindowUnit           *string                                  `json:"timeWindowUnit,omitempty" xml:"timeWindowUnit,omitempty"`
}

func (s CreateRouteRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRouteRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateRouteRuleRequest) SetAssignObjectId(v int64) *CreateRouteRuleRequest {
	s.AssignObjectId = &v
	return s
}

func (s *CreateRouteRuleRequest) SetAssignObjectType(v string) *CreateRouteRuleRequest {
	s.AssignObjectType = &v
	return s
}

func (s *CreateRouteRuleRequest) SetChildRuleRelation(v string) *CreateRouteRuleRequest {
	s.ChildRuleRelation = &v
	return s
}

func (s *CreateRouteRuleRequest) SetClientToken(v string) *CreateRouteRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRouteRuleRequest) SetCoverageProblemLevels(v []*string) *CreateRouteRuleRequest {
	s.CoverageProblemLevels = v
	return s
}

func (s *CreateRouteRuleRequest) SetEffection(v string) *CreateRouteRuleRequest {
	s.Effection = &v
	return s
}

func (s *CreateRouteRuleRequest) SetEnableStatus(v string) *CreateRouteRuleRequest {
	s.EnableStatus = &v
	return s
}

func (s *CreateRouteRuleRequest) SetIncidentLevel(v string) *CreateRouteRuleRequest {
	s.IncidentLevel = &v
	return s
}

func (s *CreateRouteRuleRequest) SetMatchCount(v int32) *CreateRouteRuleRequest {
	s.MatchCount = &v
	return s
}

func (s *CreateRouteRuleRequest) SetNotifyChannels(v []*string) *CreateRouteRuleRequest {
	s.NotifyChannels = v
	return s
}

func (s *CreateRouteRuleRequest) SetProblemEffectionServices(v []*int64) *CreateRouteRuleRequest {
	s.ProblemEffectionServices = v
	return s
}

func (s *CreateRouteRuleRequest) SetProblemLevelGroup(v map[string]*ProblemLevelGroupValue) *CreateRouteRuleRequest {
	s.ProblemLevelGroup = v
	return s
}

func (s *CreateRouteRuleRequest) SetRelatedServiceId(v int64) *CreateRouteRuleRequest {
	s.RelatedServiceId = &v
	return s
}

func (s *CreateRouteRuleRequest) SetRouteChildRules(v []*CreateRouteRuleRequestRouteChildRules) *CreateRouteRuleRequest {
	s.RouteChildRules = v
	return s
}

func (s *CreateRouteRuleRequest) SetRouteType(v string) *CreateRouteRuleRequest {
	s.RouteType = &v
	return s
}

func (s *CreateRouteRuleRequest) SetRuleName(v string) *CreateRouteRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateRouteRuleRequest) SetTimeWindow(v int64) *CreateRouteRuleRequest {
	s.TimeWindow = &v
	return s
}

func (s *CreateRouteRuleRequest) SetTimeWindowUnit(v string) *CreateRouteRuleRequest {
	s.TimeWindowUnit = &v
	return s
}

type CreateRouteRuleRequestRouteChildRules struct {
	ChildConditionRelation *int64                                             `json:"childConditionRelation,omitempty" xml:"childConditionRelation,omitempty"`
	Conditions             []*CreateRouteRuleRequestRouteChildRulesConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	MonitorSourceId        *int64                                             `json:"monitorSourceId,omitempty" xml:"monitorSourceId,omitempty"`
	ProblemLevel           *string                                            `json:"problemLevel,omitempty" xml:"problemLevel,omitempty"`
}

func (s CreateRouteRuleRequestRouteChildRules) String() string {
	return tea.Prettify(s)
}

func (s CreateRouteRuleRequestRouteChildRules) GoString() string {
	return s.String()
}

func (s *CreateRouteRuleRequestRouteChildRules) SetChildConditionRelation(v int64) *CreateRouteRuleRequestRouteChildRules {
	s.ChildConditionRelation = &v
	return s
}

func (s *CreateRouteRuleRequestRouteChildRules) SetConditions(v []*CreateRouteRuleRequestRouteChildRulesConditions) *CreateRouteRuleRequestRouteChildRules {
	s.Conditions = v
	return s
}

func (s *CreateRouteRuleRequestRouteChildRules) SetMonitorSourceId(v int64) *CreateRouteRuleRequestRouteChildRules {
	s.MonitorSourceId = &v
	return s
}

func (s *CreateRouteRuleRequestRouteChildRules) SetProblemLevel(v string) *CreateRouteRuleRequestRouteChildRules {
	s.ProblemLevel = &v
	return s
}

type CreateRouteRuleRequestRouteChildRulesConditions struct {
	Key             *string `json:"key,omitempty" xml:"key,omitempty"`
	OperationSymbol *string `json:"operationSymbol,omitempty" xml:"operationSymbol,omitempty"`
	Value           *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateRouteRuleRequestRouteChildRulesConditions) String() string {
	return tea.Prettify(s)
}

func (s CreateRouteRuleRequestRouteChildRulesConditions) GoString() string {
	return s.String()
}

func (s *CreateRouteRuleRequestRouteChildRulesConditions) SetKey(v string) *CreateRouteRuleRequestRouteChildRulesConditions {
	s.Key = &v
	return s
}

func (s *CreateRouteRuleRequestRouteChildRulesConditions) SetOperationSymbol(v string) *CreateRouteRuleRequestRouteChildRulesConditions {
	s.OperationSymbol = &v
	return s
}

func (s *CreateRouteRuleRequestRouteChildRulesConditions) SetValue(v string) *CreateRouteRuleRequestRouteChildRulesConditions {
	s.Value = &v
	return s
}

type CreateRouteRuleResponseBody struct {
	Data      *CreateRouteRuleResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	RequestId *string                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateRouteRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRouteRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRouteRuleResponseBody) SetData(v *CreateRouteRuleResponseBodyData) *CreateRouteRuleResponseBody {
	s.Data = v
	return s
}

func (s *CreateRouteRuleResponseBody) SetRequestId(v string) *CreateRouteRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateRouteRuleResponseBodyData struct {
	RouteRuleId *int64 `json:"routeRuleId,omitempty" xml:"routeRuleId,omitempty"`
}

func (s CreateRouteRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateRouteRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateRouteRuleResponseBodyData) SetRouteRuleId(v int64) *CreateRouteRuleResponseBodyData {
	s.RouteRuleId = &v
	return s
}

type CreateRouteRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRouteRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRouteRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRouteRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateRouteRuleResponse) SetHeaders(v map[string]*string) *CreateRouteRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateRouteRuleResponse) SetStatusCode(v int32) *CreateRouteRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRouteRuleResponse) SetBody(v *CreateRouteRuleResponseBody) *CreateRouteRuleResponse {
	s.Body = v
	return s
}

type CreateServiceRequest struct {
	ClientToken        *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	ServiceDescription *string `json:"serviceDescription,omitempty" xml:"serviceDescription,omitempty"`
	ServiceName        *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
}

func (s CreateServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceRequest) SetClientToken(v string) *CreateServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceRequest) SetServiceDescription(v string) *CreateServiceRequest {
	s.ServiceDescription = &v
	return s
}

func (s *CreateServiceRequest) SetServiceName(v string) *CreateServiceRequest {
	s.ServiceName = &v
	return s
}

type CreateServiceResponseBody struct {
	Data *CreateServiceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceResponseBody) SetData(v *CreateServiceResponseBodyData) *CreateServiceResponseBody {
	s.Data = v
	return s
}

func (s *CreateServiceResponseBody) SetRequestId(v string) *CreateServiceResponseBody {
	s.RequestId = &v
	return s
}

type CreateServiceResponseBodyData struct {
	ServiceId *int64 `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
}

func (s CreateServiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateServiceResponseBodyData) SetServiceId(v int64) *CreateServiceResponseBodyData {
	s.ServiceId = &v
	return s
}

type CreateServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceResponse) SetHeaders(v map[string]*string) *CreateServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceResponse) SetStatusCode(v int32) *CreateServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceResponse) SetBody(v *CreateServiceResponseBody) *CreateServiceResponse {
	s.Body = v
	return s
}

type CreateServiceGroupRequest struct {
	ClientToken             *string                                            `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	EnableWebhook           *string                                            `json:"enableWebhook,omitempty" xml:"enableWebhook,omitempty"`
	MonitorSourceTemplates  []*CreateServiceGroupRequestMonitorSourceTemplates `json:"monitorSourceTemplates,omitempty" xml:"monitorSourceTemplates,omitempty" type:"Repeated"`
	ServiceGroupDescription *string                                            `json:"serviceGroupDescription,omitempty" xml:"serviceGroupDescription,omitempty"`
	ServiceGroupName        *string                                            `json:"serviceGroupName,omitempty" xml:"serviceGroupName,omitempty"`
	UserIds                 []*int64                                           `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
	// webhooklink
	WebhookLink *string `json:"webhookLink,omitempty" xml:"webhookLink,omitempty"`
	WebhookType *string `json:"webhookType,omitempty" xml:"webhookType,omitempty"`
}

func (s CreateServiceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceGroupRequest) SetClientToken(v string) *CreateServiceGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceGroupRequest) SetEnableWebhook(v string) *CreateServiceGroupRequest {
	s.EnableWebhook = &v
	return s
}

func (s *CreateServiceGroupRequest) SetMonitorSourceTemplates(v []*CreateServiceGroupRequestMonitorSourceTemplates) *CreateServiceGroupRequest {
	s.MonitorSourceTemplates = v
	return s
}

func (s *CreateServiceGroupRequest) SetServiceGroupDescription(v string) *CreateServiceGroupRequest {
	s.ServiceGroupDescription = &v
	return s
}

func (s *CreateServiceGroupRequest) SetServiceGroupName(v string) *CreateServiceGroupRequest {
	s.ServiceGroupName = &v
	return s
}

func (s *CreateServiceGroupRequest) SetUserIds(v []*int64) *CreateServiceGroupRequest {
	s.UserIds = v
	return s
}

func (s *CreateServiceGroupRequest) SetWebhookLink(v string) *CreateServiceGroupRequest {
	s.WebhookLink = &v
	return s
}

func (s *CreateServiceGroupRequest) SetWebhookType(v string) *CreateServiceGroupRequest {
	s.WebhookType = &v
	return s
}

type CreateServiceGroupRequestMonitorSourceTemplates struct {
	MonitorSourceId   *int64  `json:"monitorSourceId,omitempty" xml:"monitorSourceId,omitempty"`
	MonitorSourceName *string `json:"monitorSourceName,omitempty" xml:"monitorSourceName,omitempty"`
	TemplateContent   *string `json:"templateContent,omitempty" xml:"templateContent,omitempty"`
	TemplateId        *int64  `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s CreateServiceGroupRequestMonitorSourceTemplates) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceGroupRequestMonitorSourceTemplates) GoString() string {
	return s.String()
}

func (s *CreateServiceGroupRequestMonitorSourceTemplates) SetMonitorSourceId(v int64) *CreateServiceGroupRequestMonitorSourceTemplates {
	s.MonitorSourceId = &v
	return s
}

func (s *CreateServiceGroupRequestMonitorSourceTemplates) SetMonitorSourceName(v string) *CreateServiceGroupRequestMonitorSourceTemplates {
	s.MonitorSourceName = &v
	return s
}

func (s *CreateServiceGroupRequestMonitorSourceTemplates) SetTemplateContent(v string) *CreateServiceGroupRequestMonitorSourceTemplates {
	s.TemplateContent = &v
	return s
}

func (s *CreateServiceGroupRequestMonitorSourceTemplates) SetTemplateId(v int64) *CreateServiceGroupRequestMonitorSourceTemplates {
	s.TemplateId = &v
	return s
}

type CreateServiceGroupResponseBody struct {
	Data *CreateServiceGroupResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateServiceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceGroupResponseBody) SetData(v *CreateServiceGroupResponseBodyData) *CreateServiceGroupResponseBody {
	s.Data = v
	return s
}

func (s *CreateServiceGroupResponseBody) SetRequestId(v string) *CreateServiceGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateServiceGroupResponseBodyData struct {
	ServiceGroupId *int64 `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
}

func (s CreateServiceGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateServiceGroupResponseBodyData) SetServiceGroupId(v int64) *CreateServiceGroupResponseBodyData {
	s.ServiceGroupId = &v
	return s
}

type CreateServiceGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateServiceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateServiceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceGroupResponse) SetHeaders(v map[string]*string) *CreateServiceGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceGroupResponse) SetStatusCode(v int32) *CreateServiceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceGroupResponse) SetBody(v *CreateServiceGroupResponseBody) *CreateServiceGroupResponse {
	s.Body = v
	return s
}

type CreateServiceGroupSchedulingRequest struct {
	ClientToken    *string                                            `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	FastScheduling *CreateServiceGroupSchedulingRequestFastScheduling `json:"fastScheduling,omitempty" xml:"fastScheduling,omitempty" type:"Struct"`
	FineScheduling *CreateServiceGroupSchedulingRequestFineScheduling `json:"fineScheduling,omitempty" xml:"fineScheduling,omitempty" type:"Struct"`
	SchedulingWay  *string                                            `json:"schedulingWay,omitempty" xml:"schedulingWay,omitempty"`
	ServiceGroupId *int64                                             `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
}

func (s CreateServiceGroupSchedulingRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceGroupSchedulingRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceGroupSchedulingRequest) SetClientToken(v string) *CreateServiceGroupSchedulingRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceGroupSchedulingRequest) SetFastScheduling(v *CreateServiceGroupSchedulingRequestFastScheduling) *CreateServiceGroupSchedulingRequest {
	s.FastScheduling = v
	return s
}

func (s *CreateServiceGroupSchedulingRequest) SetFineScheduling(v *CreateServiceGroupSchedulingRequestFineScheduling) *CreateServiceGroupSchedulingRequest {
	s.FineScheduling = v
	return s
}

func (s *CreateServiceGroupSchedulingRequest) SetSchedulingWay(v string) *CreateServiceGroupSchedulingRequest {
	s.SchedulingWay = &v
	return s
}

func (s *CreateServiceGroupSchedulingRequest) SetServiceGroupId(v int64) *CreateServiceGroupSchedulingRequest {
	s.ServiceGroupId = &v
	return s
}

type CreateServiceGroupSchedulingRequestFastScheduling struct {
	DutyPlan           *string                                                             `json:"dutyPlan,omitempty" xml:"dutyPlan,omitempty"`
	SchedulingUsers    []*CreateServiceGroupSchedulingRequestFastSchedulingSchedulingUsers `json:"schedulingUsers,omitempty" xml:"schedulingUsers,omitempty" type:"Repeated"`
	SingleDuration     *int32                                                              `json:"singleDuration,omitempty" xml:"singleDuration,omitempty"`
	SingleDurationUnit *string                                                             `json:"singleDurationUnit,omitempty" xml:"singleDurationUnit,omitempty"`
}

func (s CreateServiceGroupSchedulingRequestFastScheduling) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceGroupSchedulingRequestFastScheduling) GoString() string {
	return s.String()
}

func (s *CreateServiceGroupSchedulingRequestFastScheduling) SetDutyPlan(v string) *CreateServiceGroupSchedulingRequestFastScheduling {
	s.DutyPlan = &v
	return s
}

func (s *CreateServiceGroupSchedulingRequestFastScheduling) SetSchedulingUsers(v []*CreateServiceGroupSchedulingRequestFastSchedulingSchedulingUsers) *CreateServiceGroupSchedulingRequestFastScheduling {
	s.SchedulingUsers = v
	return s
}

func (s *CreateServiceGroupSchedulingRequestFastScheduling) SetSingleDuration(v int32) *CreateServiceGroupSchedulingRequestFastScheduling {
	s.SingleDuration = &v
	return s
}

func (s *CreateServiceGroupSchedulingRequestFastScheduling) SetSingleDurationUnit(v string) *CreateServiceGroupSchedulingRequestFastScheduling {
	s.SingleDurationUnit = &v
	return s
}

type CreateServiceGroupSchedulingRequestFastSchedulingSchedulingUsers struct {
	SchedulingOrder  *int64 `json:"schedulingOrder,omitempty" xml:"schedulingOrder,omitempty"`
	SchedulingUserId *int64 `json:"schedulingUserId,omitempty" xml:"schedulingUserId,omitempty"`
}

func (s CreateServiceGroupSchedulingRequestFastSchedulingSchedulingUsers) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceGroupSchedulingRequestFastSchedulingSchedulingUsers) GoString() string {
	return s.String()
}

func (s *CreateServiceGroupSchedulingRequestFastSchedulingSchedulingUsers) SetSchedulingOrder(v int64) *CreateServiceGroupSchedulingRequestFastSchedulingSchedulingUsers {
	s.SchedulingOrder = &v
	return s
}

func (s *CreateServiceGroupSchedulingRequestFastSchedulingSchedulingUsers) SetSchedulingUserId(v int64) *CreateServiceGroupSchedulingRequestFastSchedulingSchedulingUsers {
	s.SchedulingUserId = &v
	return s
}

type CreateServiceGroupSchedulingRequestFineScheduling struct {
	Period                       *int32                                                                           `json:"period,omitempty" xml:"period,omitempty"`
	PeriodUnit                   *string                                                                          `json:"periodUnit,omitempty" xml:"periodUnit,omitempty"`
	SchedulingFineShifts         []*CreateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts         `json:"schedulingFineShifts,omitempty" xml:"schedulingFineShifts,omitempty" type:"Repeated"`
	SchedulingTemplateFineShifts []*CreateServiceGroupSchedulingRequestFineSchedulingSchedulingTemplateFineShifts `json:"schedulingTemplateFineShifts,omitempty" xml:"schedulingTemplateFineShifts,omitempty" type:"Repeated"`
	ShiftType                    *string                                                                          `json:"shiftType,omitempty" xml:"shiftType,omitempty"`
}

func (s CreateServiceGroupSchedulingRequestFineScheduling) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceGroupSchedulingRequestFineScheduling) GoString() string {
	return s.String()
}

func (s *CreateServiceGroupSchedulingRequestFineScheduling) SetPeriod(v int32) *CreateServiceGroupSchedulingRequestFineScheduling {
	s.Period = &v
	return s
}

func (s *CreateServiceGroupSchedulingRequestFineScheduling) SetPeriodUnit(v string) *CreateServiceGroupSchedulingRequestFineScheduling {
	s.PeriodUnit = &v
	return s
}

func (s *CreateServiceGroupSchedulingRequestFineScheduling) SetSchedulingFineShifts(v []*CreateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts) *CreateServiceGroupSchedulingRequestFineScheduling {
	s.SchedulingFineShifts = v
	return s
}

func (s *CreateServiceGroupSchedulingRequestFineScheduling) SetSchedulingTemplateFineShifts(v []*CreateServiceGroupSchedulingRequestFineSchedulingSchedulingTemplateFineShifts) *CreateServiceGroupSchedulingRequestFineScheduling {
	s.SchedulingTemplateFineShifts = v
	return s
}

func (s *CreateServiceGroupSchedulingRequestFineScheduling) SetShiftType(v string) *CreateServiceGroupSchedulingRequestFineScheduling {
	s.ShiftType = &v
	return s
}

type CreateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts struct {
	CycleOrder          *int32  `json:"cycleOrder,omitempty" xml:"cycleOrder,omitempty"`
	SchedulingEndTime   *string `json:"schedulingEndTime,omitempty" xml:"schedulingEndTime,omitempty"`
	SchedulingOrder     *int32  `json:"schedulingOrder,omitempty" xml:"schedulingOrder,omitempty"`
	SchedulingStartTime *string `json:"schedulingStartTime,omitempty" xml:"schedulingStartTime,omitempty"`
	SchedulingUserId    *int64  `json:"schedulingUserId,omitempty" xml:"schedulingUserId,omitempty"`
	ShiftName           *string `json:"shiftName,omitempty" xml:"shiftName,omitempty"`
	SkipOneDay          *bool   `json:"skipOneDay,omitempty" xml:"skipOneDay,omitempty"`
}

func (s CreateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts) GoString() string {
	return s.String()
}

func (s *CreateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts) SetCycleOrder(v int32) *CreateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts {
	s.CycleOrder = &v
	return s
}

func (s *CreateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts) SetSchedulingEndTime(v string) *CreateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts {
	s.SchedulingEndTime = &v
	return s
}

func (s *CreateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts) SetSchedulingOrder(v int32) *CreateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts {
	s.SchedulingOrder = &v
	return s
}

func (s *CreateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts) SetSchedulingStartTime(v string) *CreateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts {
	s.SchedulingStartTime = &v
	return s
}

func (s *CreateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts) SetSchedulingUserId(v int64) *CreateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts {
	s.SchedulingUserId = &v
	return s
}

func (s *CreateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts) SetShiftName(v string) *CreateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts {
	s.ShiftName = &v
	return s
}

func (s *CreateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts) SetSkipOneDay(v bool) *CreateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts {
	s.SkipOneDay = &v
	return s
}

type CreateServiceGroupSchedulingRequestFineSchedulingSchedulingTemplateFineShifts struct {
	SchedulingEndTime   *string `json:"schedulingEndTime,omitempty" xml:"schedulingEndTime,omitempty"`
	SchedulingOrder     *int64  `json:"schedulingOrder,omitempty" xml:"schedulingOrder,omitempty"`
	SchedulingStartTime *string `json:"schedulingStartTime,omitempty" xml:"schedulingStartTime,omitempty"`
	SchedulingUserId    *int64  `json:"schedulingUserId,omitempty" xml:"schedulingUserId,omitempty"`
	SchedulingUserName  *string `json:"schedulingUserName,omitempty" xml:"schedulingUserName,omitempty"`
	SkipOneDay          *bool   `json:"skipOneDay,omitempty" xml:"skipOneDay,omitempty"`
}

func (s CreateServiceGroupSchedulingRequestFineSchedulingSchedulingTemplateFineShifts) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceGroupSchedulingRequestFineSchedulingSchedulingTemplateFineShifts) GoString() string {
	return s.String()
}

func (s *CreateServiceGroupSchedulingRequestFineSchedulingSchedulingTemplateFineShifts) SetSchedulingEndTime(v string) *CreateServiceGroupSchedulingRequestFineSchedulingSchedulingTemplateFineShifts {
	s.SchedulingEndTime = &v
	return s
}

func (s *CreateServiceGroupSchedulingRequestFineSchedulingSchedulingTemplateFineShifts) SetSchedulingOrder(v int64) *CreateServiceGroupSchedulingRequestFineSchedulingSchedulingTemplateFineShifts {
	s.SchedulingOrder = &v
	return s
}

func (s *CreateServiceGroupSchedulingRequestFineSchedulingSchedulingTemplateFineShifts) SetSchedulingStartTime(v string) *CreateServiceGroupSchedulingRequestFineSchedulingSchedulingTemplateFineShifts {
	s.SchedulingStartTime = &v
	return s
}

func (s *CreateServiceGroupSchedulingRequestFineSchedulingSchedulingTemplateFineShifts) SetSchedulingUserId(v int64) *CreateServiceGroupSchedulingRequestFineSchedulingSchedulingTemplateFineShifts {
	s.SchedulingUserId = &v
	return s
}

func (s *CreateServiceGroupSchedulingRequestFineSchedulingSchedulingTemplateFineShifts) SetSchedulingUserName(v string) *CreateServiceGroupSchedulingRequestFineSchedulingSchedulingTemplateFineShifts {
	s.SchedulingUserName = &v
	return s
}

func (s *CreateServiceGroupSchedulingRequestFineSchedulingSchedulingTemplateFineShifts) SetSkipOneDay(v bool) *CreateServiceGroupSchedulingRequestFineSchedulingSchedulingTemplateFineShifts {
	s.SkipOneDay = &v
	return s
}

type CreateServiceGroupSchedulingResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateServiceGroupSchedulingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceGroupSchedulingResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceGroupSchedulingResponseBody) SetRequestId(v string) *CreateServiceGroupSchedulingResponseBody {
	s.RequestId = &v
	return s
}

type CreateServiceGroupSchedulingResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateServiceGroupSchedulingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateServiceGroupSchedulingResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceGroupSchedulingResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceGroupSchedulingResponse) SetHeaders(v map[string]*string) *CreateServiceGroupSchedulingResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceGroupSchedulingResponse) SetStatusCode(v int32) *CreateServiceGroupSchedulingResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceGroupSchedulingResponse) SetBody(v *CreateServiceGroupSchedulingResponseBody) *CreateServiceGroupSchedulingResponse {
	s.Body = v
	return s
}

type CreateSubscriptionRequest struct {
	ClientToken        *string                                        `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	EndTime            *string                                        `json:"endTime,omitempty" xml:"endTime,omitempty"`
	ExpiredType        *int64                                         `json:"expiredType,omitempty" xml:"expiredType,omitempty"`
	NotifyObjectList   []*CreateSubscriptionRequestNotifyObjectList   `json:"notifyObjectList,omitempty" xml:"notifyObjectList,omitempty" type:"Repeated"`
	NotifyObjectType   *int64                                         `json:"notifyObjectType,omitempty" xml:"notifyObjectType,omitempty"`
	NotifyStrategyList []*CreateSubscriptionRequestNotifyStrategyList `json:"notifyStrategyList,omitempty" xml:"notifyStrategyList,omitempty" type:"Repeated"`
	Period             *string                                        `json:"period,omitempty" xml:"period,omitempty"`
	Scope              *int64                                         `json:"scope,omitempty" xml:"scope,omitempty"`
	ScopeObjectList    []*CreateSubscriptionRequestScopeObjectList    `json:"scopeObjectList,omitempty" xml:"scopeObjectList,omitempty" type:"Repeated"`
	StartTime          *string                                        `json:"startTime,omitempty" xml:"startTime,omitempty"`
	SubscriptionTitle  *string                                        `json:"subscriptionTitle,omitempty" xml:"subscriptionTitle,omitempty"`
}

func (s CreateSubscriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionRequest) SetClientToken(v string) *CreateSubscriptionRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSubscriptionRequest) SetEndTime(v string) *CreateSubscriptionRequest {
	s.EndTime = &v
	return s
}

func (s *CreateSubscriptionRequest) SetExpiredType(v int64) *CreateSubscriptionRequest {
	s.ExpiredType = &v
	return s
}

func (s *CreateSubscriptionRequest) SetNotifyObjectList(v []*CreateSubscriptionRequestNotifyObjectList) *CreateSubscriptionRequest {
	s.NotifyObjectList = v
	return s
}

func (s *CreateSubscriptionRequest) SetNotifyObjectType(v int64) *CreateSubscriptionRequest {
	s.NotifyObjectType = &v
	return s
}

func (s *CreateSubscriptionRequest) SetNotifyStrategyList(v []*CreateSubscriptionRequestNotifyStrategyList) *CreateSubscriptionRequest {
	s.NotifyStrategyList = v
	return s
}

func (s *CreateSubscriptionRequest) SetPeriod(v string) *CreateSubscriptionRequest {
	s.Period = &v
	return s
}

func (s *CreateSubscriptionRequest) SetScope(v int64) *CreateSubscriptionRequest {
	s.Scope = &v
	return s
}

func (s *CreateSubscriptionRequest) SetScopeObjectList(v []*CreateSubscriptionRequestScopeObjectList) *CreateSubscriptionRequest {
	s.ScopeObjectList = v
	return s
}

func (s *CreateSubscriptionRequest) SetStartTime(v string) *CreateSubscriptionRequest {
	s.StartTime = &v
	return s
}

func (s *CreateSubscriptionRequest) SetSubscriptionTitle(v string) *CreateSubscriptionRequest {
	s.SubscriptionTitle = &v
	return s
}

type CreateSubscriptionRequestNotifyObjectList struct {
	NotifyObjectId *int64 `json:"notifyObjectId,omitempty" xml:"notifyObjectId,omitempty"`
}

func (s CreateSubscriptionRequestNotifyObjectList) String() string {
	return tea.Prettify(s)
}

func (s CreateSubscriptionRequestNotifyObjectList) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionRequestNotifyObjectList) SetNotifyObjectId(v int64) *CreateSubscriptionRequestNotifyObjectList {
	s.NotifyObjectId = &v
	return s
}

type CreateSubscriptionRequestNotifyStrategyList struct {
	Channels      *string                                                   `json:"channels,omitempty" xml:"channels,omitempty"`
	InstanceType  *int64                                                    `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	PeriodChannel *CreateSubscriptionRequestNotifyStrategyListPeriodChannel `json:"periodChannel,omitempty" xml:"periodChannel,omitempty" type:"Struct"`
	Strategies    []*CreateSubscriptionRequestNotifyStrategyListStrategies  `json:"strategies,omitempty" xml:"strategies,omitempty" type:"Repeated"`
}

func (s CreateSubscriptionRequestNotifyStrategyList) String() string {
	return tea.Prettify(s)
}

func (s CreateSubscriptionRequestNotifyStrategyList) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionRequestNotifyStrategyList) SetChannels(v string) *CreateSubscriptionRequestNotifyStrategyList {
	s.Channels = &v
	return s
}

func (s *CreateSubscriptionRequestNotifyStrategyList) SetInstanceType(v int64) *CreateSubscriptionRequestNotifyStrategyList {
	s.InstanceType = &v
	return s
}

func (s *CreateSubscriptionRequestNotifyStrategyList) SetPeriodChannel(v *CreateSubscriptionRequestNotifyStrategyListPeriodChannel) *CreateSubscriptionRequestNotifyStrategyList {
	s.PeriodChannel = v
	return s
}

func (s *CreateSubscriptionRequestNotifyStrategyList) SetStrategies(v []*CreateSubscriptionRequestNotifyStrategyListStrategies) *CreateSubscriptionRequestNotifyStrategyList {
	s.Strategies = v
	return s
}

type CreateSubscriptionRequestNotifyStrategyListPeriodChannel struct {
	NonWorkday *string `json:"nonWorkday,omitempty" xml:"nonWorkday,omitempty"`
	Workday    *string `json:"workday,omitempty" xml:"workday,omitempty"`
}

func (s CreateSubscriptionRequestNotifyStrategyListPeriodChannel) String() string {
	return tea.Prettify(s)
}

func (s CreateSubscriptionRequestNotifyStrategyListPeriodChannel) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionRequestNotifyStrategyListPeriodChannel) SetNonWorkday(v string) *CreateSubscriptionRequestNotifyStrategyListPeriodChannel {
	s.NonWorkday = &v
	return s
}

func (s *CreateSubscriptionRequestNotifyStrategyListPeriodChannel) SetWorkday(v string) *CreateSubscriptionRequestNotifyStrategyListPeriodChannel {
	s.Workday = &v
	return s
}

type CreateSubscriptionRequestNotifyStrategyListStrategies struct {
	Conditions []*CreateSubscriptionRequestNotifyStrategyListStrategiesConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
}

func (s CreateSubscriptionRequestNotifyStrategyListStrategies) String() string {
	return tea.Prettify(s)
}

func (s CreateSubscriptionRequestNotifyStrategyListStrategies) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionRequestNotifyStrategyListStrategies) SetConditions(v []*CreateSubscriptionRequestNotifyStrategyListStrategiesConditions) *CreateSubscriptionRequestNotifyStrategyListStrategies {
	s.Conditions = v
	return s
}

type CreateSubscriptionRequestNotifyStrategyListStrategiesConditions struct {
	Action            *string `json:"action,omitempty" xml:"action,omitempty"`
	Effection         *string `json:"effection,omitempty" xml:"effection,omitempty"`
	Level             *string `json:"level,omitempty" xml:"level,omitempty"`
	ProblemNotifyType *string `json:"problemNotifyType,omitempty" xml:"problemNotifyType,omitempty"`
}

func (s CreateSubscriptionRequestNotifyStrategyListStrategiesConditions) String() string {
	return tea.Prettify(s)
}

func (s CreateSubscriptionRequestNotifyStrategyListStrategiesConditions) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionRequestNotifyStrategyListStrategiesConditions) SetAction(v string) *CreateSubscriptionRequestNotifyStrategyListStrategiesConditions {
	s.Action = &v
	return s
}

func (s *CreateSubscriptionRequestNotifyStrategyListStrategiesConditions) SetEffection(v string) *CreateSubscriptionRequestNotifyStrategyListStrategiesConditions {
	s.Effection = &v
	return s
}

func (s *CreateSubscriptionRequestNotifyStrategyListStrategiesConditions) SetLevel(v string) *CreateSubscriptionRequestNotifyStrategyListStrategiesConditions {
	s.Level = &v
	return s
}

func (s *CreateSubscriptionRequestNotifyStrategyListStrategiesConditions) SetProblemNotifyType(v string) *CreateSubscriptionRequestNotifyStrategyListStrategiesConditions {
	s.ProblemNotifyType = &v
	return s
}

type CreateSubscriptionRequestScopeObjectList struct {
	ScopeObjectId *int64 `json:"scopeObjectId,omitempty" xml:"scopeObjectId,omitempty"`
}

func (s CreateSubscriptionRequestScopeObjectList) String() string {
	return tea.Prettify(s)
}

func (s CreateSubscriptionRequestScopeObjectList) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionRequestScopeObjectList) SetScopeObjectId(v int64) *CreateSubscriptionRequestScopeObjectList {
	s.ScopeObjectId = &v
	return s
}

type CreateSubscriptionResponseBody struct {
	// data
	Data *CreateSubscriptionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// request id
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateSubscriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionResponseBody) SetData(v *CreateSubscriptionResponseBodyData) *CreateSubscriptionResponseBody {
	s.Data = v
	return s
}

func (s *CreateSubscriptionResponseBody) SetRequestId(v string) *CreateSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

type CreateSubscriptionResponseBodyData struct {
	SubscriptionId *int64 `json:"subscriptionId,omitempty" xml:"subscriptionId,omitempty"`
}

func (s CreateSubscriptionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateSubscriptionResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionResponseBodyData) SetSubscriptionId(v int64) *CreateSubscriptionResponseBodyData {
	s.SubscriptionId = &v
	return s
}

type CreateSubscriptionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSubscriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionResponse) SetHeaders(v map[string]*string) *CreateSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *CreateSubscriptionResponse) SetStatusCode(v int32) *CreateSubscriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSubscriptionResponse) SetBody(v *CreateSubscriptionResponseBody) *CreateSubscriptionResponse {
	s.Body = v
	return s
}

type CreateTenantApplicationRequest struct {
	Channel     *string `json:"channel,omitempty" xml:"channel,omitempty"`
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateTenantApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTenantApplicationRequest) GoString() string {
	return s.String()
}

func (s *CreateTenantApplicationRequest) SetChannel(v string) *CreateTenantApplicationRequest {
	s.Channel = &v
	return s
}

func (s *CreateTenantApplicationRequest) SetClientToken(v string) *CreateTenantApplicationRequest {
	s.ClientToken = &v
	return s
}

type CreateTenantApplicationResponseBody struct {
	// data
	Data *CreateTenantApplicationResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// id of the req
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateTenantApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTenantApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTenantApplicationResponseBody) SetData(v *CreateTenantApplicationResponseBodyData) *CreateTenantApplicationResponseBody {
	s.Data = v
	return s
}

func (s *CreateTenantApplicationResponseBody) SetRequestId(v string) *CreateTenantApplicationResponseBody {
	s.RequestId = &v
	return s
}

type CreateTenantApplicationResponseBodyData struct {
	OpenUrl  *string `json:"openUrl,omitempty" xml:"openUrl,omitempty"`
	Progress *string `json:"progress,omitempty" xml:"progress,omitempty"`
}

func (s CreateTenantApplicationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateTenantApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTenantApplicationResponseBodyData) SetOpenUrl(v string) *CreateTenantApplicationResponseBodyData {
	s.OpenUrl = &v
	return s
}

func (s *CreateTenantApplicationResponseBodyData) SetProgress(v string) *CreateTenantApplicationResponseBodyData {
	s.Progress = &v
	return s
}

type CreateTenantApplicationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTenantApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTenantApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTenantApplicationResponse) GoString() string {
	return s.String()
}

func (s *CreateTenantApplicationResponse) SetHeaders(v map[string]*string) *CreateTenantApplicationResponse {
	s.Headers = v
	return s
}

func (s *CreateTenantApplicationResponse) SetStatusCode(v int32) *CreateTenantApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTenantApplicationResponse) SetBody(v *CreateTenantApplicationResponseBody) *CreateTenantApplicationResponse {
	s.Body = v
	return s
}

type CreateUserRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	Email       *string `json:"email,omitempty" xml:"email,omitempty"`
	Phone       *string `json:"phone,omitempty" xml:"phone,omitempty"`
	RamId       *int64  `json:"ramId,omitempty" xml:"ramId,omitempty"`
	Username    *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s CreateUserRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserRequest) GoString() string {
	return s.String()
}

func (s *CreateUserRequest) SetClientToken(v string) *CreateUserRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateUserRequest) SetEmail(v string) *CreateUserRequest {
	s.Email = &v
	return s
}

func (s *CreateUserRequest) SetPhone(v string) *CreateUserRequest {
	s.Phone = &v
	return s
}

func (s *CreateUserRequest) SetRamId(v int64) *CreateUserRequest {
	s.RamId = &v
	return s
}

func (s *CreateUserRequest) SetUsername(v string) *CreateUserRequest {
	s.Username = &v
	return s
}

type CreateUserResponseBody struct {
	// data
	Data *CreateUserResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBody) SetData(v *CreateUserResponseBodyData) *CreateUserResponseBody {
	s.Data = v
	return s
}

func (s *CreateUserResponseBody) SetRequestId(v string) *CreateUserResponseBody {
	s.RequestId = &v
	return s
}

type CreateUserResponseBodyData struct {
	UserId *int64 `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateUserResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBodyData) SetUserId(v int64) *CreateUserResponseBodyData {
	s.UserId = &v
	return s
}

type CreateUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUserResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponse) GoString() string {
	return s.String()
}

func (s *CreateUserResponse) SetHeaders(v map[string]*string) *CreateUserResponse {
	s.Headers = v
	return s
}

func (s *CreateUserResponse) SetStatusCode(v int32) *CreateUserResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserResponse) SetBody(v *CreateUserResponseBody) *CreateUserResponse {
	s.Body = v
	return s
}

type DeleteEscalationPlanRequest struct {
	ClientToken      *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	EscalationPlanId *int64  `json:"escalationPlanId,omitempty" xml:"escalationPlanId,omitempty"`
}

func (s DeleteEscalationPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEscalationPlanRequest) GoString() string {
	return s.String()
}

func (s *DeleteEscalationPlanRequest) SetClientToken(v string) *DeleteEscalationPlanRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteEscalationPlanRequest) SetEscalationPlanId(v int64) *DeleteEscalationPlanRequest {
	s.EscalationPlanId = &v
	return s
}

type DeleteEscalationPlanResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteEscalationPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEscalationPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEscalationPlanResponseBody) SetRequestId(v string) *DeleteEscalationPlanResponseBody {
	s.RequestId = &v
	return s
}

type DeleteEscalationPlanResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteEscalationPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteEscalationPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEscalationPlanResponse) GoString() string {
	return s.String()
}

func (s *DeleteEscalationPlanResponse) SetHeaders(v map[string]*string) *DeleteEscalationPlanResponse {
	s.Headers = v
	return s
}

func (s *DeleteEscalationPlanResponse) SetStatusCode(v int32) *DeleteEscalationPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEscalationPlanResponse) SetBody(v *DeleteEscalationPlanResponseBody) *DeleteEscalationPlanResponse {
	s.Body = v
	return s
}

type DeleteIncidentRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	IncidentId  *int64  `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
}

func (s DeleteIncidentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIncidentRequest) GoString() string {
	return s.String()
}

func (s *DeleteIncidentRequest) SetClientToken(v string) *DeleteIncidentRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteIncidentRequest) SetIncidentId(v int64) *DeleteIncidentRequest {
	s.IncidentId = &v
	return s
}

type DeleteIncidentResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteIncidentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIncidentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIncidentResponseBody) SetRequestId(v string) *DeleteIncidentResponseBody {
	s.RequestId = &v
	return s
}

type DeleteIncidentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteIncidentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteIncidentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIncidentResponse) GoString() string {
	return s.String()
}

func (s *DeleteIncidentResponse) SetHeaders(v map[string]*string) *DeleteIncidentResponse {
	s.Headers = v
	return s
}

func (s *DeleteIncidentResponse) SetStatusCode(v int32) *DeleteIncidentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIncidentResponse) SetBody(v *DeleteIncidentResponseBody) *DeleteIncidentResponse {
	s.Body = v
	return s
}

type DeleteIntegrationConfigRequest struct {
	ClientToken         *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	IntegrationConfigId *int64  `json:"integrationConfigId,omitempty" xml:"integrationConfigId,omitempty"`
}

func (s DeleteIntegrationConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIntegrationConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteIntegrationConfigRequest) SetClientToken(v string) *DeleteIntegrationConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteIntegrationConfigRequest) SetIntegrationConfigId(v int64) *DeleteIntegrationConfigRequest {
	s.IntegrationConfigId = &v
	return s
}

type DeleteIntegrationConfigResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteIntegrationConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIntegrationConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIntegrationConfigResponseBody) SetRequestId(v string) *DeleteIntegrationConfigResponseBody {
	s.RequestId = &v
	return s
}

type DeleteIntegrationConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteIntegrationConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteIntegrationConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIntegrationConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteIntegrationConfigResponse) SetHeaders(v map[string]*string) *DeleteIntegrationConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteIntegrationConfigResponse) SetStatusCode(v int32) *DeleteIntegrationConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIntegrationConfigResponse) SetBody(v *DeleteIntegrationConfigResponseBody) *DeleteIntegrationConfigResponse {
	s.Body = v
	return s
}

type DeleteProblemRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	ProblemId   *int64  `json:"problemId,omitempty" xml:"problemId,omitempty"`
}

func (s DeleteProblemRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProblemRequest) GoString() string {
	return s.String()
}

func (s *DeleteProblemRequest) SetClientToken(v string) *DeleteProblemRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteProblemRequest) SetProblemId(v int64) *DeleteProblemRequest {
	s.ProblemId = &v
	return s
}

type DeleteProblemResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteProblemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProblemResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProblemResponseBody) SetRequestId(v string) *DeleteProblemResponseBody {
	s.RequestId = &v
	return s
}

type DeleteProblemResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteProblemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProblemResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProblemResponse) GoString() string {
	return s.String()
}

func (s *DeleteProblemResponse) SetHeaders(v map[string]*string) *DeleteProblemResponse {
	s.Headers = v
	return s
}

func (s *DeleteProblemResponse) SetStatusCode(v int32) *DeleteProblemResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProblemResponse) SetBody(v *DeleteProblemResponseBody) *DeleteProblemResponse {
	s.Body = v
	return s
}

type DeleteProblemEffectionServiceRequest struct {
	// clientToken
	ClientToken        *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	EffectionServiceId *int64  `json:"effectionServiceId,omitempty" xml:"effectionServiceId,omitempty"`
	ProblemId          *int64  `json:"problemId,omitempty" xml:"problemId,omitempty"`
}

func (s DeleteProblemEffectionServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProblemEffectionServiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteProblemEffectionServiceRequest) SetClientToken(v string) *DeleteProblemEffectionServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteProblemEffectionServiceRequest) SetEffectionServiceId(v int64) *DeleteProblemEffectionServiceRequest {
	s.EffectionServiceId = &v
	return s
}

func (s *DeleteProblemEffectionServiceRequest) SetProblemId(v int64) *DeleteProblemEffectionServiceRequest {
	s.ProblemId = &v
	return s
}

type DeleteProblemEffectionServiceResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteProblemEffectionServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProblemEffectionServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProblemEffectionServiceResponseBody) SetRequestId(v string) *DeleteProblemEffectionServiceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteProblemEffectionServiceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteProblemEffectionServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProblemEffectionServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProblemEffectionServiceResponse) GoString() string {
	return s.String()
}

func (s *DeleteProblemEffectionServiceResponse) SetHeaders(v map[string]*string) *DeleteProblemEffectionServiceResponse {
	s.Headers = v
	return s
}

func (s *DeleteProblemEffectionServiceResponse) SetStatusCode(v int32) *DeleteProblemEffectionServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProblemEffectionServiceResponse) SetBody(v *DeleteProblemEffectionServiceResponseBody) *DeleteProblemEffectionServiceResponse {
	s.Body = v
	return s
}

type DeleteProblemMeasureRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	MeasureId   *int64  `json:"measureId,omitempty" xml:"measureId,omitempty"`
	ProblemId   *string `json:"problemId,omitempty" xml:"problemId,omitempty"`
}

func (s DeleteProblemMeasureRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProblemMeasureRequest) GoString() string {
	return s.String()
}

func (s *DeleteProblemMeasureRequest) SetClientToken(v string) *DeleteProblemMeasureRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteProblemMeasureRequest) SetMeasureId(v int64) *DeleteProblemMeasureRequest {
	s.MeasureId = &v
	return s
}

func (s *DeleteProblemMeasureRequest) SetProblemId(v string) *DeleteProblemMeasureRequest {
	s.ProblemId = &v
	return s
}

type DeleteProblemMeasureResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteProblemMeasureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProblemMeasureResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProblemMeasureResponseBody) SetRequestId(v string) *DeleteProblemMeasureResponseBody {
	s.RequestId = &v
	return s
}

type DeleteProblemMeasureResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteProblemMeasureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProblemMeasureResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProblemMeasureResponse) GoString() string {
	return s.String()
}

func (s *DeleteProblemMeasureResponse) SetHeaders(v map[string]*string) *DeleteProblemMeasureResponse {
	s.Headers = v
	return s
}

func (s *DeleteProblemMeasureResponse) SetStatusCode(v int32) *DeleteProblemMeasureResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProblemMeasureResponse) SetBody(v *DeleteProblemMeasureResponseBody) *DeleteProblemMeasureResponse {
	s.Body = v
	return s
}

type DeleteProblemTimelineRequest struct {
	ClientToken       *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	ProblemId         *int64  `json:"problemId,omitempty" xml:"problemId,omitempty"`
	ProblemTimelineId *int64  `json:"problemTimelineId,omitempty" xml:"problemTimelineId,omitempty"`
}

func (s DeleteProblemTimelineRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProblemTimelineRequest) GoString() string {
	return s.String()
}

func (s *DeleteProblemTimelineRequest) SetClientToken(v string) *DeleteProblemTimelineRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteProblemTimelineRequest) SetProblemId(v int64) *DeleteProblemTimelineRequest {
	s.ProblemId = &v
	return s
}

func (s *DeleteProblemTimelineRequest) SetProblemTimelineId(v int64) *DeleteProblemTimelineRequest {
	s.ProblemTimelineId = &v
	return s
}

type DeleteProblemTimelineResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteProblemTimelineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProblemTimelineResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProblemTimelineResponseBody) SetRequestId(v string) *DeleteProblemTimelineResponseBody {
	s.RequestId = &v
	return s
}

type DeleteProblemTimelineResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteProblemTimelineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProblemTimelineResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProblemTimelineResponse) GoString() string {
	return s.String()
}

func (s *DeleteProblemTimelineResponse) SetHeaders(v map[string]*string) *DeleteProblemTimelineResponse {
	s.Headers = v
	return s
}

func (s *DeleteProblemTimelineResponse) SetStatusCode(v int32) *DeleteProblemTimelineResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProblemTimelineResponse) SetBody(v *DeleteProblemTimelineResponseBody) *DeleteProblemTimelineResponse {
	s.Body = v
	return s
}

type DeleteRouteRuleRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	RouteRuleId *int64  `json:"routeRuleId,omitempty" xml:"routeRuleId,omitempty"`
}

func (s DeleteRouteRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRouteRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRouteRuleRequest) SetClientToken(v string) *DeleteRouteRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteRouteRuleRequest) SetRouteRuleId(v int64) *DeleteRouteRuleRequest {
	s.RouteRuleId = &v
	return s
}

type DeleteRouteRuleResponseBody struct {
	RequestId *int64 `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteRouteRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRouteRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRouteRuleResponseBody) SetRequestId(v int64) *DeleteRouteRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRouteRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteRouteRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRouteRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRouteRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteRouteRuleResponse) SetHeaders(v map[string]*string) *DeleteRouteRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteRouteRuleResponse) SetStatusCode(v int32) *DeleteRouteRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRouteRuleResponse) SetBody(v *DeleteRouteRuleResponseBody) *DeleteRouteRuleResponse {
	s.Body = v
	return s
}

type DeleteServiceRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	ServiceId   *int64  `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
}

func (s DeleteServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceRequest) SetClientToken(v string) *DeleteServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteServiceRequest) SetServiceId(v int64) *DeleteServiceRequest {
	s.ServiceId = &v
	return s
}

type DeleteServiceResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceResponseBody) SetRequestId(v string) *DeleteServiceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceResponse) SetHeaders(v map[string]*string) *DeleteServiceResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceResponse) SetStatusCode(v int32) *DeleteServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceResponse) SetBody(v *DeleteServiceResponseBody) *DeleteServiceResponse {
	s.Body = v
	return s
}

type DeleteServiceGroupRequest struct {
	ClientToken    *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	ServiceGroupId *int64  `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
}

func (s DeleteServiceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceGroupRequest) SetClientToken(v string) *DeleteServiceGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteServiceGroupRequest) SetServiceGroupId(v int64) *DeleteServiceGroupRequest {
	s.ServiceGroupId = &v
	return s
}

type DeleteServiceGroupResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteServiceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceGroupResponseBody) SetRequestId(v string) *DeleteServiceGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServiceGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteServiceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteServiceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceGroupResponse) SetHeaders(v map[string]*string) *DeleteServiceGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceGroupResponse) SetStatusCode(v int32) *DeleteServiceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceGroupResponse) SetBody(v *DeleteServiceGroupResponseBody) *DeleteServiceGroupResponse {
	s.Body = v
	return s
}

type DeleteServiceGroupSchedulingResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteServiceGroupSchedulingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceGroupSchedulingResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceGroupSchedulingResponseBody) SetRequestId(v string) *DeleteServiceGroupSchedulingResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServiceGroupSchedulingResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteServiceGroupSchedulingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteServiceGroupSchedulingResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceGroupSchedulingResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceGroupSchedulingResponse) SetHeaders(v map[string]*string) *DeleteServiceGroupSchedulingResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceGroupSchedulingResponse) SetStatusCode(v int32) *DeleteServiceGroupSchedulingResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceGroupSchedulingResponse) SetBody(v *DeleteServiceGroupSchedulingResponseBody) *DeleteServiceGroupSchedulingResponse {
	s.Body = v
	return s
}

type DeleteServiceGroupUserRequest struct {
	ClientToken    *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	NewUserId      *int64  `json:"newUserId,omitempty" xml:"newUserId,omitempty"`
	OldUserId      *int64  `json:"oldUserId,omitempty" xml:"oldUserId,omitempty"`
	RemoveUser     *bool   `json:"removeUser,omitempty" xml:"removeUser,omitempty"`
	ServiceGroupId *int64  `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
}

func (s DeleteServiceGroupUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceGroupUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceGroupUserRequest) SetClientToken(v string) *DeleteServiceGroupUserRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteServiceGroupUserRequest) SetNewUserId(v int64) *DeleteServiceGroupUserRequest {
	s.NewUserId = &v
	return s
}

func (s *DeleteServiceGroupUserRequest) SetOldUserId(v int64) *DeleteServiceGroupUserRequest {
	s.OldUserId = &v
	return s
}

func (s *DeleteServiceGroupUserRequest) SetRemoveUser(v bool) *DeleteServiceGroupUserRequest {
	s.RemoveUser = &v
	return s
}

func (s *DeleteServiceGroupUserRequest) SetServiceGroupId(v int64) *DeleteServiceGroupUserRequest {
	s.ServiceGroupId = &v
	return s
}

type DeleteServiceGroupUserResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteServiceGroupUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceGroupUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceGroupUserResponseBody) SetRequestId(v string) *DeleteServiceGroupUserResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServiceGroupUserResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteServiceGroupUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteServiceGroupUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceGroupUserResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceGroupUserResponse) SetHeaders(v map[string]*string) *DeleteServiceGroupUserResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceGroupUserResponse) SetStatusCode(v int32) *DeleteServiceGroupUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceGroupUserResponse) SetBody(v *DeleteServiceGroupUserResponseBody) *DeleteServiceGroupUserResponse {
	s.Body = v
	return s
}

type DeleteSubscriptionRequest struct {
	SubscriptionId *int64 `json:"subscriptionId,omitempty" xml:"subscriptionId,omitempty"`
}

func (s DeleteSubscriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *DeleteSubscriptionRequest) SetSubscriptionId(v int64) *DeleteSubscriptionRequest {
	s.SubscriptionId = &v
	return s
}

type DeleteSubscriptionResponseBody struct {
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteSubscriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSubscriptionResponseBody) SetRequestId(v string) *DeleteSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSubscriptionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSubscriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *DeleteSubscriptionResponse) SetHeaders(v map[string]*string) *DeleteSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *DeleteSubscriptionResponse) SetStatusCode(v int32) *DeleteSubscriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSubscriptionResponse) SetBody(v *DeleteSubscriptionResponseBody) *DeleteSubscriptionResponse {
	s.Body = v
	return s
}

type DeleteUserRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	UserId      *int64  `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeleteUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserRequest) SetClientToken(v string) *DeleteUserRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteUserRequest) SetUserId(v int64) *DeleteUserRequest {
	s.UserId = &v
	return s
}

type DeleteUserResponseBody struct {
	// id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserResponseBody) SetRequestId(v string) *DeleteUserResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserResponse) SetHeaders(v map[string]*string) *DeleteUserResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserResponse) SetStatusCode(v int32) *DeleteUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserResponse) SetBody(v *DeleteUserResponseBody) *DeleteUserResponse {
	s.Body = v
	return s
}

type DeliverIncidentRequest struct {
	AssignUserId *int64  `json:"assignUserId,omitempty" xml:"assignUserId,omitempty"`
	ClientToken  *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	IncidentId   *int64  `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
}

func (s DeliverIncidentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeliverIncidentRequest) GoString() string {
	return s.String()
}

func (s *DeliverIncidentRequest) SetAssignUserId(v int64) *DeliverIncidentRequest {
	s.AssignUserId = &v
	return s
}

func (s *DeliverIncidentRequest) SetClientToken(v string) *DeliverIncidentRequest {
	s.ClientToken = &v
	return s
}

func (s *DeliverIncidentRequest) SetIncidentId(v int64) *DeliverIncidentRequest {
	s.IncidentId = &v
	return s
}

type DeliverIncidentResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeliverIncidentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeliverIncidentResponseBody) GoString() string {
	return s.String()
}

func (s *DeliverIncidentResponseBody) SetRequestId(v string) *DeliverIncidentResponseBody {
	s.RequestId = &v
	return s
}

type DeliverIncidentResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeliverIncidentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeliverIncidentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeliverIncidentResponse) GoString() string {
	return s.String()
}

func (s *DeliverIncidentResponse) SetHeaders(v map[string]*string) *DeliverIncidentResponse {
	s.Headers = v
	return s
}

func (s *DeliverIncidentResponse) SetStatusCode(v int32) *DeliverIncidentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeliverIncidentResponse) SetBody(v *DeliverIncidentResponseBody) *DeliverIncidentResponse {
	s.Body = v
	return s
}

type DisableEscalationPlanRequest struct {
	ClientToken      *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	EscalationPlanId *int64  `json:"escalationPlanId,omitempty" xml:"escalationPlanId,omitempty"`
}

func (s DisableEscalationPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableEscalationPlanRequest) GoString() string {
	return s.String()
}

func (s *DisableEscalationPlanRequest) SetClientToken(v string) *DisableEscalationPlanRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableEscalationPlanRequest) SetEscalationPlanId(v int64) *DisableEscalationPlanRequest {
	s.EscalationPlanId = &v
	return s
}

type DisableEscalationPlanResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DisableEscalationPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableEscalationPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DisableEscalationPlanResponseBody) SetRequestId(v string) *DisableEscalationPlanResponseBody {
	s.RequestId = &v
	return s
}

type DisableEscalationPlanResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableEscalationPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableEscalationPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableEscalationPlanResponse) GoString() string {
	return s.String()
}

func (s *DisableEscalationPlanResponse) SetHeaders(v map[string]*string) *DisableEscalationPlanResponse {
	s.Headers = v
	return s
}

func (s *DisableEscalationPlanResponse) SetStatusCode(v int32) *DisableEscalationPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableEscalationPlanResponse) SetBody(v *DisableEscalationPlanResponseBody) *DisableEscalationPlanResponse {
	s.Body = v
	return s
}

type DisableIntegrationConfigRequest struct {
	ClientToken         *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	IntegrationConfigId *int64  `json:"integrationConfigId,omitempty" xml:"integrationConfigId,omitempty"`
}

func (s DisableIntegrationConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableIntegrationConfigRequest) GoString() string {
	return s.String()
}

func (s *DisableIntegrationConfigRequest) SetClientToken(v string) *DisableIntegrationConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableIntegrationConfigRequest) SetIntegrationConfigId(v int64) *DisableIntegrationConfigRequest {
	s.IntegrationConfigId = &v
	return s
}

type DisableIntegrationConfigResponseBody struct {
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DisableIntegrationConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableIntegrationConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DisableIntegrationConfigResponseBody) SetRequestId(v string) *DisableIntegrationConfigResponseBody {
	s.RequestId = &v
	return s
}

type DisableIntegrationConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableIntegrationConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableIntegrationConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableIntegrationConfigResponse) GoString() string {
	return s.String()
}

func (s *DisableIntegrationConfigResponse) SetHeaders(v map[string]*string) *DisableIntegrationConfigResponse {
	s.Headers = v
	return s
}

func (s *DisableIntegrationConfigResponse) SetStatusCode(v int32) *DisableIntegrationConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableIntegrationConfigResponse) SetBody(v *DisableIntegrationConfigResponseBody) *DisableIntegrationConfigResponse {
	s.Body = v
	return s
}

type DisableRouteRuleRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	RouteRuleId *int64  `json:"routeRuleId,omitempty" xml:"routeRuleId,omitempty"`
}

func (s DisableRouteRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableRouteRuleRequest) GoString() string {
	return s.String()
}

func (s *DisableRouteRuleRequest) SetClientToken(v string) *DisableRouteRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableRouteRuleRequest) SetRouteRuleId(v int64) *DisableRouteRuleRequest {
	s.RouteRuleId = &v
	return s
}

type DisableRouteRuleResponseBody struct {
	// C4BE3837-1A13-413B-A225-2C88188E8A43
	Data      *int64  `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DisableRouteRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableRouteRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DisableRouteRuleResponseBody) SetData(v int64) *DisableRouteRuleResponseBody {
	s.Data = &v
	return s
}

func (s *DisableRouteRuleResponseBody) SetRequestId(v string) *DisableRouteRuleResponseBody {
	s.RequestId = &v
	return s
}

type DisableRouteRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableRouteRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableRouteRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableRouteRuleResponse) GoString() string {
	return s.String()
}

func (s *DisableRouteRuleResponse) SetHeaders(v map[string]*string) *DisableRouteRuleResponse {
	s.Headers = v
	return s
}

func (s *DisableRouteRuleResponse) SetStatusCode(v int32) *DisableRouteRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableRouteRuleResponse) SetBody(v *DisableRouteRuleResponseBody) *DisableRouteRuleResponse {
	s.Body = v
	return s
}

type DisableServiceGroupWebhookRequest struct {
	ClientToken    *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	ServiceGroupId *int64  `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
}

func (s DisableServiceGroupWebhookRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableServiceGroupWebhookRequest) GoString() string {
	return s.String()
}

func (s *DisableServiceGroupWebhookRequest) SetClientToken(v string) *DisableServiceGroupWebhookRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableServiceGroupWebhookRequest) SetServiceGroupId(v int64) *DisableServiceGroupWebhookRequest {
	s.ServiceGroupId = &v
	return s
}

type DisableServiceGroupWebhookResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DisableServiceGroupWebhookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableServiceGroupWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *DisableServiceGroupWebhookResponseBody) SetRequestId(v string) *DisableServiceGroupWebhookResponseBody {
	s.RequestId = &v
	return s
}

type DisableServiceGroupWebhookResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableServiceGroupWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableServiceGroupWebhookResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableServiceGroupWebhookResponse) GoString() string {
	return s.String()
}

func (s *DisableServiceGroupWebhookResponse) SetHeaders(v map[string]*string) *DisableServiceGroupWebhookResponse {
	s.Headers = v
	return s
}

func (s *DisableServiceGroupWebhookResponse) SetStatusCode(v int32) *DisableServiceGroupWebhookResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableServiceGroupWebhookResponse) SetBody(v *DisableServiceGroupWebhookResponseBody) *DisableServiceGroupWebhookResponse {
	s.Body = v
	return s
}

type DisableSubscriptionRequest struct {
	SubscriptionId *int64 `json:"subscriptionId,omitempty" xml:"subscriptionId,omitempty"`
}

func (s DisableSubscriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *DisableSubscriptionRequest) SetSubscriptionId(v int64) *DisableSubscriptionRequest {
	s.SubscriptionId = &v
	return s
}

type DisableSubscriptionResponseBody struct {
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DisableSubscriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *DisableSubscriptionResponseBody) SetRequestId(v string) *DisableSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

type DisableSubscriptionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableSubscriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *DisableSubscriptionResponse) SetHeaders(v map[string]*string) *DisableSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *DisableSubscriptionResponse) SetStatusCode(v int32) *DisableSubscriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableSubscriptionResponse) SetBody(v *DisableSubscriptionResponseBody) *DisableSubscriptionResponse {
	s.Body = v
	return s
}

type EnableEscalationPlanRequest struct {
	ClientToken      *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	EscalationPlanId *int64  `json:"escalationPlanId,omitempty" xml:"escalationPlanId,omitempty"`
}

func (s EnableEscalationPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableEscalationPlanRequest) GoString() string {
	return s.String()
}

func (s *EnableEscalationPlanRequest) SetClientToken(v string) *EnableEscalationPlanRequest {
	s.ClientToken = &v
	return s
}

func (s *EnableEscalationPlanRequest) SetEscalationPlanId(v int64) *EnableEscalationPlanRequest {
	s.EscalationPlanId = &v
	return s
}

type EnableEscalationPlanResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s EnableEscalationPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableEscalationPlanResponseBody) GoString() string {
	return s.String()
}

func (s *EnableEscalationPlanResponseBody) SetRequestId(v string) *EnableEscalationPlanResponseBody {
	s.RequestId = &v
	return s
}

type EnableEscalationPlanResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableEscalationPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableEscalationPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableEscalationPlanResponse) GoString() string {
	return s.String()
}

func (s *EnableEscalationPlanResponse) SetHeaders(v map[string]*string) *EnableEscalationPlanResponse {
	s.Headers = v
	return s
}

func (s *EnableEscalationPlanResponse) SetStatusCode(v int32) *EnableEscalationPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableEscalationPlanResponse) SetBody(v *EnableEscalationPlanResponseBody) *EnableEscalationPlanResponse {
	s.Body = v
	return s
}

type EnableIntegrationConfigRequest struct {
	ClientToken         *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	IntegrationConfigId *int64  `json:"integrationConfigId,omitempty" xml:"integrationConfigId,omitempty"`
}

func (s EnableIntegrationConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableIntegrationConfigRequest) GoString() string {
	return s.String()
}

func (s *EnableIntegrationConfigRequest) SetClientToken(v string) *EnableIntegrationConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *EnableIntegrationConfigRequest) SetIntegrationConfigId(v int64) *EnableIntegrationConfigRequest {
	s.IntegrationConfigId = &v
	return s
}

type EnableIntegrationConfigResponseBody struct {
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s EnableIntegrationConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableIntegrationConfigResponseBody) GoString() string {
	return s.String()
}

func (s *EnableIntegrationConfigResponseBody) SetRequestId(v string) *EnableIntegrationConfigResponseBody {
	s.RequestId = &v
	return s
}

type EnableIntegrationConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableIntegrationConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableIntegrationConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableIntegrationConfigResponse) GoString() string {
	return s.String()
}

func (s *EnableIntegrationConfigResponse) SetHeaders(v map[string]*string) *EnableIntegrationConfigResponse {
	s.Headers = v
	return s
}

func (s *EnableIntegrationConfigResponse) SetStatusCode(v int32) *EnableIntegrationConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableIntegrationConfigResponse) SetBody(v *EnableIntegrationConfigResponseBody) *EnableIntegrationConfigResponse {
	s.Body = v
	return s
}

type EnableRouteRuleRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	RouteRuleId *int64  `json:"routeRuleId,omitempty" xml:"routeRuleId,omitempty"`
}

func (s EnableRouteRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableRouteRuleRequest) GoString() string {
	return s.String()
}

func (s *EnableRouteRuleRequest) SetClientToken(v string) *EnableRouteRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *EnableRouteRuleRequest) SetRouteRuleId(v int64) *EnableRouteRuleRequest {
	s.RouteRuleId = &v
	return s
}

type EnableRouteRuleResponseBody struct {
	Data      *int32  `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s EnableRouteRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableRouteRuleResponseBody) GoString() string {
	return s.String()
}

func (s *EnableRouteRuleResponseBody) SetData(v int32) *EnableRouteRuleResponseBody {
	s.Data = &v
	return s
}

func (s *EnableRouteRuleResponseBody) SetRequestId(v string) *EnableRouteRuleResponseBody {
	s.RequestId = &v
	return s
}

type EnableRouteRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableRouteRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableRouteRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableRouteRuleResponse) GoString() string {
	return s.String()
}

func (s *EnableRouteRuleResponse) SetHeaders(v map[string]*string) *EnableRouteRuleResponse {
	s.Headers = v
	return s
}

func (s *EnableRouteRuleResponse) SetStatusCode(v int32) *EnableRouteRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableRouteRuleResponse) SetBody(v *EnableRouteRuleResponseBody) *EnableRouteRuleResponse {
	s.Body = v
	return s
}

type EnableServiceGroupWebhookRequest struct {
	ClientToken    *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	ServiceGroupId *int64  `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
}

func (s EnableServiceGroupWebhookRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableServiceGroupWebhookRequest) GoString() string {
	return s.String()
}

func (s *EnableServiceGroupWebhookRequest) SetClientToken(v string) *EnableServiceGroupWebhookRequest {
	s.ClientToken = &v
	return s
}

func (s *EnableServiceGroupWebhookRequest) SetServiceGroupId(v int64) *EnableServiceGroupWebhookRequest {
	s.ServiceGroupId = &v
	return s
}

type EnableServiceGroupWebhookResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s EnableServiceGroupWebhookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableServiceGroupWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *EnableServiceGroupWebhookResponseBody) SetRequestId(v string) *EnableServiceGroupWebhookResponseBody {
	s.RequestId = &v
	return s
}

type EnableServiceGroupWebhookResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableServiceGroupWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableServiceGroupWebhookResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableServiceGroupWebhookResponse) GoString() string {
	return s.String()
}

func (s *EnableServiceGroupWebhookResponse) SetHeaders(v map[string]*string) *EnableServiceGroupWebhookResponse {
	s.Headers = v
	return s
}

func (s *EnableServiceGroupWebhookResponse) SetStatusCode(v int32) *EnableServiceGroupWebhookResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableServiceGroupWebhookResponse) SetBody(v *EnableServiceGroupWebhookResponseBody) *EnableServiceGroupWebhookResponse {
	s.Body = v
	return s
}

type EnableSubscriptionRequest struct {
	SubscriptionId *int64 `json:"subscriptionId,omitempty" xml:"subscriptionId,omitempty"`
}

func (s EnableSubscriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *EnableSubscriptionRequest) SetSubscriptionId(v int64) *EnableSubscriptionRequest {
	s.SubscriptionId = &v
	return s
}

type EnableSubscriptionResponseBody struct {
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s EnableSubscriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *EnableSubscriptionResponseBody) SetRequestId(v string) *EnableSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

type EnableSubscriptionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableSubscriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *EnableSubscriptionResponse) SetHeaders(v map[string]*string) *EnableSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *EnableSubscriptionResponse) SetStatusCode(v int32) *EnableSubscriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableSubscriptionResponse) SetBody(v *EnableSubscriptionResponseBody) *EnableSubscriptionResponse {
	s.Body = v
	return s
}

type FinishIncidentRequest struct {
	ClientToken                       *string  `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	IncidentFinishReason              *int32   `json:"incidentFinishReason,omitempty" xml:"incidentFinishReason,omitempty"`
	IncidentFinishReasonDescription   *string  `json:"incidentFinishReasonDescription,omitempty" xml:"incidentFinishReasonDescription,omitempty"`
	IncidentFinishSolution            *int32   `json:"incidentFinishSolution,omitempty" xml:"incidentFinishSolution,omitempty"`
	IncidentFinishSolutionDescription *string  `json:"incidentFinishSolutionDescription,omitempty" xml:"incidentFinishSolutionDescription,omitempty"`
	IncidentIds                       []*int64 `json:"incidentIds,omitempty" xml:"incidentIds,omitempty" type:"Repeated"`
}

func (s FinishIncidentRequest) String() string {
	return tea.Prettify(s)
}

func (s FinishIncidentRequest) GoString() string {
	return s.String()
}

func (s *FinishIncidentRequest) SetClientToken(v string) *FinishIncidentRequest {
	s.ClientToken = &v
	return s
}

func (s *FinishIncidentRequest) SetIncidentFinishReason(v int32) *FinishIncidentRequest {
	s.IncidentFinishReason = &v
	return s
}

func (s *FinishIncidentRequest) SetIncidentFinishReasonDescription(v string) *FinishIncidentRequest {
	s.IncidentFinishReasonDescription = &v
	return s
}

func (s *FinishIncidentRequest) SetIncidentFinishSolution(v int32) *FinishIncidentRequest {
	s.IncidentFinishSolution = &v
	return s
}

func (s *FinishIncidentRequest) SetIncidentFinishSolutionDescription(v string) *FinishIncidentRequest {
	s.IncidentFinishSolutionDescription = &v
	return s
}

func (s *FinishIncidentRequest) SetIncidentIds(v []*int64) *FinishIncidentRequest {
	s.IncidentIds = v
	return s
}

type FinishIncidentResponseBody struct {
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s FinishIncidentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FinishIncidentResponseBody) GoString() string {
	return s.String()
}

func (s *FinishIncidentResponseBody) SetRequestId(v string) *FinishIncidentResponseBody {
	s.RequestId = &v
	return s
}

type FinishIncidentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FinishIncidentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FinishIncidentResponse) String() string {
	return tea.Prettify(s)
}

func (s FinishIncidentResponse) GoString() string {
	return s.String()
}

func (s *FinishIncidentResponse) SetHeaders(v map[string]*string) *FinishIncidentResponse {
	s.Headers = v
	return s
}

func (s *FinishIncidentResponse) SetStatusCode(v int32) *FinishIncidentResponse {
	s.StatusCode = &v
	return s
}

func (s *FinishIncidentResponse) SetBody(v *FinishIncidentResponseBody) *FinishIncidentResponse {
	s.Body = v
	return s
}

type FinishProblemRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	ProblemId   *int64  `json:"problemId,omitempty" xml:"problemId,omitempty"`
}

func (s FinishProblemRequest) String() string {
	return tea.Prettify(s)
}

func (s FinishProblemRequest) GoString() string {
	return s.String()
}

func (s *FinishProblemRequest) SetClientToken(v string) *FinishProblemRequest {
	s.ClientToken = &v
	return s
}

func (s *FinishProblemRequest) SetProblemId(v int64) *FinishProblemRequest {
	s.ProblemId = &v
	return s
}

type FinishProblemResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s FinishProblemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FinishProblemResponseBody) GoString() string {
	return s.String()
}

func (s *FinishProblemResponseBody) SetRequestId(v string) *FinishProblemResponseBody {
	s.RequestId = &v
	return s
}

type FinishProblemResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FinishProblemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FinishProblemResponse) String() string {
	return tea.Prettify(s)
}

func (s FinishProblemResponse) GoString() string {
	return s.String()
}

func (s *FinishProblemResponse) SetHeaders(v map[string]*string) *FinishProblemResponse {
	s.Headers = v
	return s
}

func (s *FinishProblemResponse) SetStatusCode(v int32) *FinishProblemResponse {
	s.StatusCode = &v
	return s
}

func (s *FinishProblemResponse) SetBody(v *FinishProblemResponseBody) *FinishProblemResponse {
	s.Body = v
	return s
}

type GeneratePictureLinkRequest struct {
	// keys
	Keys      []*string `json:"keys,omitempty" xml:"keys,omitempty" type:"Repeated"`
	ProblemId *int64    `json:"problemId,omitempty" xml:"problemId,omitempty"`
}

func (s GeneratePictureLinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GeneratePictureLinkRequest) GoString() string {
	return s.String()
}

func (s *GeneratePictureLinkRequest) SetKeys(v []*string) *GeneratePictureLinkRequest {
	s.Keys = v
	return s
}

func (s *GeneratePictureLinkRequest) SetProblemId(v int64) *GeneratePictureLinkRequest {
	s.ProblemId = &v
	return s
}

type GeneratePictureLinkResponseBody struct {
	// data
	Data *GeneratePictureLinkResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GeneratePictureLinkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GeneratePictureLinkResponseBody) GoString() string {
	return s.String()
}

func (s *GeneratePictureLinkResponseBody) SetData(v *GeneratePictureLinkResponseBodyData) *GeneratePictureLinkResponseBody {
	s.Data = v
	return s
}

func (s *GeneratePictureLinkResponseBody) SetRequestId(v string) *GeneratePictureLinkResponseBody {
	s.RequestId = &v
	return s
}

type GeneratePictureLinkResponseBodyData struct {
	// array
	Links []*GeneratePictureLinkResponseBodyDataLinks `json:"links,omitempty" xml:"links,omitempty" type:"Repeated"`
}

func (s GeneratePictureLinkResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GeneratePictureLinkResponseBodyData) GoString() string {
	return s.String()
}

func (s *GeneratePictureLinkResponseBodyData) SetLinks(v []*GeneratePictureLinkResponseBodyDataLinks) *GeneratePictureLinkResponseBodyData {
	s.Links = v
	return s
}

type GeneratePictureLinkResponseBodyDataLinks struct {
	// oss key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// url
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
}

func (s GeneratePictureLinkResponseBodyDataLinks) String() string {
	return tea.Prettify(s)
}

func (s GeneratePictureLinkResponseBodyDataLinks) GoString() string {
	return s.String()
}

func (s *GeneratePictureLinkResponseBodyDataLinks) SetKey(v string) *GeneratePictureLinkResponseBodyDataLinks {
	s.Key = &v
	return s
}

func (s *GeneratePictureLinkResponseBodyDataLinks) SetLink(v string) *GeneratePictureLinkResponseBodyDataLinks {
	s.Link = &v
	return s
}

type GeneratePictureLinkResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GeneratePictureLinkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GeneratePictureLinkResponse) String() string {
	return tea.Prettify(s)
}

func (s GeneratePictureLinkResponse) GoString() string {
	return s.String()
}

func (s *GeneratePictureLinkResponse) SetHeaders(v map[string]*string) *GeneratePictureLinkResponse {
	s.Headers = v
	return s
}

func (s *GeneratePictureLinkResponse) SetStatusCode(v int32) *GeneratePictureLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *GeneratePictureLinkResponse) SetBody(v *GeneratePictureLinkResponseBody) *GeneratePictureLinkResponse {
	s.Body = v
	return s
}

type GeneratePictureUploadSignRequest struct {
	Files        []*GeneratePictureUploadSignRequestFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	InstanceId   *int64                                   `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	InstanceType *string                                  `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
}

func (s GeneratePictureUploadSignRequest) String() string {
	return tea.Prettify(s)
}

func (s GeneratePictureUploadSignRequest) GoString() string {
	return s.String()
}

func (s *GeneratePictureUploadSignRequest) SetFiles(v []*GeneratePictureUploadSignRequestFiles) *GeneratePictureUploadSignRequest {
	s.Files = v
	return s
}

func (s *GeneratePictureUploadSignRequest) SetInstanceId(v int64) *GeneratePictureUploadSignRequest {
	s.InstanceId = &v
	return s
}

func (s *GeneratePictureUploadSignRequest) SetInstanceType(v string) *GeneratePictureUploadSignRequest {
	s.InstanceType = &v
	return s
}

type GeneratePictureUploadSignRequestFiles struct {
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileSize *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
}

func (s GeneratePictureUploadSignRequestFiles) String() string {
	return tea.Prettify(s)
}

func (s GeneratePictureUploadSignRequestFiles) GoString() string {
	return s.String()
}

func (s *GeneratePictureUploadSignRequestFiles) SetFileName(v string) *GeneratePictureUploadSignRequestFiles {
	s.FileName = &v
	return s
}

func (s *GeneratePictureUploadSignRequestFiles) SetFileSize(v int64) *GeneratePictureUploadSignRequestFiles {
	s.FileSize = &v
	return s
}

func (s *GeneratePictureUploadSignRequestFiles) SetFileType(v string) *GeneratePictureUploadSignRequestFiles {
	s.FileType = &v
	return s
}

type GeneratePictureUploadSignResponseBody struct {
	// data
	Data *GeneratePictureUploadSignResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GeneratePictureUploadSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GeneratePictureUploadSignResponseBody) GoString() string {
	return s.String()
}

func (s *GeneratePictureUploadSignResponseBody) SetData(v *GeneratePictureUploadSignResponseBodyData) *GeneratePictureUploadSignResponseBody {
	s.Data = v
	return s
}

func (s *GeneratePictureUploadSignResponseBody) SetRequestId(v string) *GeneratePictureUploadSignResponseBody {
	s.RequestId = &v
	return s
}

type GeneratePictureUploadSignResponseBodyData struct {
	// accessKeyId
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// oss bucket name
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// files
	Files []*GeneratePictureUploadSignResponseBodyDataFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// policy
	Policy *string `json:"policy,omitempty" xml:"policy,omitempty"`
	// signature
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
	// url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GeneratePictureUploadSignResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GeneratePictureUploadSignResponseBodyData) GoString() string {
	return s.String()
}

func (s *GeneratePictureUploadSignResponseBodyData) SetAccessKeyId(v string) *GeneratePictureUploadSignResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GeneratePictureUploadSignResponseBodyData) SetBucketName(v string) *GeneratePictureUploadSignResponseBodyData {
	s.BucketName = &v
	return s
}

func (s *GeneratePictureUploadSignResponseBodyData) SetFiles(v []*GeneratePictureUploadSignResponseBodyDataFiles) *GeneratePictureUploadSignResponseBodyData {
	s.Files = v
	return s
}

func (s *GeneratePictureUploadSignResponseBodyData) SetPolicy(v string) *GeneratePictureUploadSignResponseBodyData {
	s.Policy = &v
	return s
}

func (s *GeneratePictureUploadSignResponseBodyData) SetSignature(v string) *GeneratePictureUploadSignResponseBodyData {
	s.Signature = &v
	return s
}

func (s *GeneratePictureUploadSignResponseBodyData) SetUrl(v string) *GeneratePictureUploadSignResponseBodyData {
	s.Url = &v
	return s
}

type GeneratePictureUploadSignResponseBodyDataFiles struct {
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileSize *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// oss key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s GeneratePictureUploadSignResponseBodyDataFiles) String() string {
	return tea.Prettify(s)
}

func (s GeneratePictureUploadSignResponseBodyDataFiles) GoString() string {
	return s.String()
}

func (s *GeneratePictureUploadSignResponseBodyDataFiles) SetFileName(v string) *GeneratePictureUploadSignResponseBodyDataFiles {
	s.FileName = &v
	return s
}

func (s *GeneratePictureUploadSignResponseBodyDataFiles) SetFileSize(v int64) *GeneratePictureUploadSignResponseBodyDataFiles {
	s.FileSize = &v
	return s
}

func (s *GeneratePictureUploadSignResponseBodyDataFiles) SetFileType(v string) *GeneratePictureUploadSignResponseBodyDataFiles {
	s.FileType = &v
	return s
}

func (s *GeneratePictureUploadSignResponseBodyDataFiles) SetKey(v string) *GeneratePictureUploadSignResponseBodyDataFiles {
	s.Key = &v
	return s
}

type GeneratePictureUploadSignResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GeneratePictureUploadSignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GeneratePictureUploadSignResponse) String() string {
	return tea.Prettify(s)
}

func (s GeneratePictureUploadSignResponse) GoString() string {
	return s.String()
}

func (s *GeneratePictureUploadSignResponse) SetHeaders(v map[string]*string) *GeneratePictureUploadSignResponse {
	s.Headers = v
	return s
}

func (s *GeneratePictureUploadSignResponse) SetStatusCode(v int32) *GeneratePictureUploadSignResponse {
	s.StatusCode = &v
	return s
}

func (s *GeneratePictureUploadSignResponse) SetBody(v *GeneratePictureUploadSignResponseBody) *GeneratePictureUploadSignResponse {
	s.Body = v
	return s
}

type GenerateProblemPictureLinkRequest struct {
	// oss key
	Keys      []*string `json:"keys,omitempty" xml:"keys,omitempty" type:"Repeated"`
	ProblemId *string   `json:"problemId,omitempty" xml:"problemId,omitempty"`
}

func (s GenerateProblemPictureLinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateProblemPictureLinkRequest) GoString() string {
	return s.String()
}

func (s *GenerateProblemPictureLinkRequest) SetKeys(v []*string) *GenerateProblemPictureLinkRequest {
	s.Keys = v
	return s
}

func (s *GenerateProblemPictureLinkRequest) SetProblemId(v string) *GenerateProblemPictureLinkRequest {
	s.ProblemId = &v
	return s
}

type GenerateProblemPictureLinkResponseBody struct {
	// data
	Data *GenerateProblemPictureLinkResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GenerateProblemPictureLinkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateProblemPictureLinkResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateProblemPictureLinkResponseBody) SetData(v *GenerateProblemPictureLinkResponseBodyData) *GenerateProblemPictureLinkResponseBody {
	s.Data = v
	return s
}

func (s *GenerateProblemPictureLinkResponseBody) SetRequestId(v string) *GenerateProblemPictureLinkResponseBody {
	s.RequestId = &v
	return s
}

type GenerateProblemPictureLinkResponseBodyData struct {
	Links []*GenerateProblemPictureLinkResponseBodyDataLinks `json:"links,omitempty" xml:"links,omitempty" type:"Repeated"`
}

func (s GenerateProblemPictureLinkResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateProblemPictureLinkResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateProblemPictureLinkResponseBodyData) SetLinks(v []*GenerateProblemPictureLinkResponseBodyDataLinks) *GenerateProblemPictureLinkResponseBodyData {
	s.Links = v
	return s
}

type GenerateProblemPictureLinkResponseBodyDataLinks struct {
	// oss key
	Key  *string `json:"key,omitempty" xml:"key,omitempty"`
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
}

func (s GenerateProblemPictureLinkResponseBodyDataLinks) String() string {
	return tea.Prettify(s)
}

func (s GenerateProblemPictureLinkResponseBodyDataLinks) GoString() string {
	return s.String()
}

func (s *GenerateProblemPictureLinkResponseBodyDataLinks) SetKey(v string) *GenerateProblemPictureLinkResponseBodyDataLinks {
	s.Key = &v
	return s
}

func (s *GenerateProblemPictureLinkResponseBodyDataLinks) SetLink(v string) *GenerateProblemPictureLinkResponseBodyDataLinks {
	s.Link = &v
	return s
}

type GenerateProblemPictureLinkResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateProblemPictureLinkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateProblemPictureLinkResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateProblemPictureLinkResponse) GoString() string {
	return s.String()
}

func (s *GenerateProblemPictureLinkResponse) SetHeaders(v map[string]*string) *GenerateProblemPictureLinkResponse {
	s.Headers = v
	return s
}

func (s *GenerateProblemPictureLinkResponse) SetStatusCode(v int32) *GenerateProblemPictureLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateProblemPictureLinkResponse) SetBody(v *GenerateProblemPictureLinkResponseBody) *GenerateProblemPictureLinkResponse {
	s.Body = v
	return s
}

type GenerateProblemPictureUploadSignRequest struct {
	FileName  *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileSize  *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	FileType  *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	ProblemId *int64  `json:"problemId,omitempty" xml:"problemId,omitempty"`
}

func (s GenerateProblemPictureUploadSignRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateProblemPictureUploadSignRequest) GoString() string {
	return s.String()
}

func (s *GenerateProblemPictureUploadSignRequest) SetFileName(v string) *GenerateProblemPictureUploadSignRequest {
	s.FileName = &v
	return s
}

func (s *GenerateProblemPictureUploadSignRequest) SetFileSize(v int64) *GenerateProblemPictureUploadSignRequest {
	s.FileSize = &v
	return s
}

func (s *GenerateProblemPictureUploadSignRequest) SetFileType(v string) *GenerateProblemPictureUploadSignRequest {
	s.FileType = &v
	return s
}

func (s *GenerateProblemPictureUploadSignRequest) SetProblemId(v int64) *GenerateProblemPictureUploadSignRequest {
	s.ProblemId = &v
	return s
}

type GenerateProblemPictureUploadSignResponseBody struct {
	// data
	Data *GenerateProblemPictureUploadSignResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GenerateProblemPictureUploadSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateProblemPictureUploadSignResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateProblemPictureUploadSignResponseBody) SetData(v *GenerateProblemPictureUploadSignResponseBodyData) *GenerateProblemPictureUploadSignResponseBody {
	s.Data = v
	return s
}

func (s *GenerateProblemPictureUploadSignResponseBody) SetRequestId(v string) *GenerateProblemPictureUploadSignResponseBody {
	s.RequestId = &v
	return s
}

type GenerateProblemPictureUploadSignResponseBodyData struct {
	// ossaccessKeyId
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// oss bucket name
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// oss key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// policy
	Policy *string `json:"policy,omitempty" xml:"policy,omitempty"`
	// signature
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
	// url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GenerateProblemPictureUploadSignResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateProblemPictureUploadSignResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateProblemPictureUploadSignResponseBodyData) SetAccessKeyId(v string) *GenerateProblemPictureUploadSignResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GenerateProblemPictureUploadSignResponseBodyData) SetBucketName(v string) *GenerateProblemPictureUploadSignResponseBodyData {
	s.BucketName = &v
	return s
}

func (s *GenerateProblemPictureUploadSignResponseBodyData) SetKey(v string) *GenerateProblemPictureUploadSignResponseBodyData {
	s.Key = &v
	return s
}

func (s *GenerateProblemPictureUploadSignResponseBodyData) SetPolicy(v string) *GenerateProblemPictureUploadSignResponseBodyData {
	s.Policy = &v
	return s
}

func (s *GenerateProblemPictureUploadSignResponseBodyData) SetSignature(v string) *GenerateProblemPictureUploadSignResponseBodyData {
	s.Signature = &v
	return s
}

func (s *GenerateProblemPictureUploadSignResponseBodyData) SetUrl(v string) *GenerateProblemPictureUploadSignResponseBodyData {
	s.Url = &v
	return s
}

type GenerateProblemPictureUploadSignResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateProblemPictureUploadSignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateProblemPictureUploadSignResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateProblemPictureUploadSignResponse) GoString() string {
	return s.String()
}

func (s *GenerateProblemPictureUploadSignResponse) SetHeaders(v map[string]*string) *GenerateProblemPictureUploadSignResponse {
	s.Headers = v
	return s
}

func (s *GenerateProblemPictureUploadSignResponse) SetStatusCode(v int32) *GenerateProblemPictureUploadSignResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateProblemPictureUploadSignResponse) SetBody(v *GenerateProblemPictureUploadSignResponseBody) *GenerateProblemPictureUploadSignResponse {
	s.Body = v
	return s
}

type GetEscalationPlanRequest struct {
	ClientToken      *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	EscalationPlanId *int64  `json:"escalationPlanId,omitempty" xml:"escalationPlanId,omitempty"`
}

func (s GetEscalationPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEscalationPlanRequest) GoString() string {
	return s.String()
}

func (s *GetEscalationPlanRequest) SetClientToken(v string) *GetEscalationPlanRequest {
	s.ClientToken = &v
	return s
}

func (s *GetEscalationPlanRequest) SetEscalationPlanId(v int64) *GetEscalationPlanRequest {
	s.EscalationPlanId = &v
	return s
}

type GetEscalationPlanResponseBody struct {
	// data
	Data *GetEscalationPlanResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetEscalationPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEscalationPlanResponseBody) GoString() string {
	return s.String()
}

func (s *GetEscalationPlanResponseBody) SetData(v *GetEscalationPlanResponseBodyData) *GetEscalationPlanResponseBody {
	s.Data = v
	return s
}

func (s *GetEscalationPlanResponseBody) SetRequestId(v string) *GetEscalationPlanResponseBody {
	s.RequestId = &v
	return s
}

type GetEscalationPlanResponseBodyData struct {
	CreateTime                 *string                                                        `json:"createTime,omitempty" xml:"createTime,omitempty"`
	EscalationPlanDescription  *string                                                        `json:"escalationPlanDescription,omitempty" xml:"escalationPlanDescription,omitempty"`
	EscalationPlanId           *int64                                                         `json:"escalationPlanId,omitempty" xml:"escalationPlanId,omitempty"`
	EscalationPlanName         *string                                                        `json:"escalationPlanName,omitempty" xml:"escalationPlanName,omitempty"`
	EscalationPlanRules        []*GetEscalationPlanResponseBodyDataEscalationPlanRules        `json:"escalationPlanRules,omitempty" xml:"escalationPlanRules,omitempty" type:"Repeated"`
	EscalationPlanScopeObjects []*GetEscalationPlanResponseBodyDataEscalationPlanScopeObjects `json:"escalationPlanScopeObjects,omitempty" xml:"escalationPlanScopeObjects,omitempty" type:"Repeated"`
}

func (s GetEscalationPlanResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetEscalationPlanResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEscalationPlanResponseBodyData) SetCreateTime(v string) *GetEscalationPlanResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetEscalationPlanResponseBodyData) SetEscalationPlanDescription(v string) *GetEscalationPlanResponseBodyData {
	s.EscalationPlanDescription = &v
	return s
}

func (s *GetEscalationPlanResponseBodyData) SetEscalationPlanId(v int64) *GetEscalationPlanResponseBodyData {
	s.EscalationPlanId = &v
	return s
}

func (s *GetEscalationPlanResponseBodyData) SetEscalationPlanName(v string) *GetEscalationPlanResponseBodyData {
	s.EscalationPlanName = &v
	return s
}

func (s *GetEscalationPlanResponseBodyData) SetEscalationPlanRules(v []*GetEscalationPlanResponseBodyDataEscalationPlanRules) *GetEscalationPlanResponseBodyData {
	s.EscalationPlanRules = v
	return s
}

func (s *GetEscalationPlanResponseBodyData) SetEscalationPlanScopeObjects(v []*GetEscalationPlanResponseBodyDataEscalationPlanScopeObjects) *GetEscalationPlanResponseBodyData {
	s.EscalationPlanScopeObjects = v
	return s
}

type GetEscalationPlanResponseBodyDataEscalationPlanRules struct {
	EscalationPlanConditions []*GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanConditions `json:"escalationPlanConditions,omitempty" xml:"escalationPlanConditions,omitempty" type:"Repeated"`
	EscalationPlanRuleId     *int64                                                                          `json:"escalationPlanRuleId,omitempty" xml:"escalationPlanRuleId,omitempty"`
	EscalationPlanStrategies []*GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanStrategies `json:"escalationPlanStrategies,omitempty" xml:"escalationPlanStrategies,omitempty" type:"Repeated"`
}

func (s GetEscalationPlanResponseBodyDataEscalationPlanRules) String() string {
	return tea.Prettify(s)
}

func (s GetEscalationPlanResponseBodyDataEscalationPlanRules) GoString() string {
	return s.String()
}

func (s *GetEscalationPlanResponseBodyDataEscalationPlanRules) SetEscalationPlanConditions(v []*GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanConditions) *GetEscalationPlanResponseBodyDataEscalationPlanRules {
	s.EscalationPlanConditions = v
	return s
}

func (s *GetEscalationPlanResponseBodyDataEscalationPlanRules) SetEscalationPlanRuleId(v int64) *GetEscalationPlanResponseBodyDataEscalationPlanRules {
	s.EscalationPlanRuleId = &v
	return s
}

func (s *GetEscalationPlanResponseBodyDataEscalationPlanRules) SetEscalationPlanStrategies(v []*GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanStrategies) *GetEscalationPlanResponseBodyDataEscalationPlanRules {
	s.EscalationPlanStrategies = v
	return s
}

type GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanConditions struct {
	Effection *string `json:"effection,omitempty" xml:"effection,omitempty"`
	Level     *string `json:"level,omitempty" xml:"level,omitempty"`
}

func (s GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanConditions) String() string {
	return tea.Prettify(s)
}

func (s GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanConditions) GoString() string {
	return s.String()
}

func (s *GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanConditions) SetEffection(v string) *GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanConditions {
	s.Effection = &v
	return s
}

func (s *GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanConditions) SetLevel(v string) *GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanConditions {
	s.Level = &v
	return s
}

type GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanStrategies struct {
	EnableWebhook      *bool                                                                                           `json:"enableWebhook,omitempty" xml:"enableWebhook,omitempty"`
	EscalationPlanType *string                                                                                         `json:"escalationPlanType,omitempty" xml:"escalationPlanType,omitempty"`
	NoticeChannels     *string                                                                                         `json:"noticeChannels,omitempty" xml:"noticeChannels,omitempty"`
	NoticeObjectList   []*GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanStrategiesNoticeObjectList `json:"noticeObjectList,omitempty" xml:"noticeObjectList,omitempty" type:"Repeated"`
	NoticeTime         *int64                                                                                          `json:"noticeTime,omitempty" xml:"noticeTime,omitempty"`
	ServiceGroups      []*GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanStrategiesServiceGroups    `json:"serviceGroups,omitempty" xml:"serviceGroups,omitempty" type:"Repeated"`
}

func (s GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanStrategies) String() string {
	return tea.Prettify(s)
}

func (s GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanStrategies) GoString() string {
	return s.String()
}

func (s *GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanStrategies) SetEnableWebhook(v bool) *GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanStrategies {
	s.EnableWebhook = &v
	return s
}

func (s *GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanStrategies) SetEscalationPlanType(v string) *GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanStrategies {
	s.EscalationPlanType = &v
	return s
}

func (s *GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanStrategies) SetNoticeChannels(v string) *GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanStrategies {
	s.NoticeChannels = &v
	return s
}

func (s *GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanStrategies) SetNoticeObjectList(v []*GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanStrategiesNoticeObjectList) *GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanStrategies {
	s.NoticeObjectList = v
	return s
}

func (s *GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanStrategies) SetNoticeTime(v int64) *GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanStrategies {
	s.NoticeTime = &v
	return s
}

func (s *GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanStrategies) SetServiceGroups(v []*GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanStrategiesServiceGroups) *GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanStrategies {
	s.ServiceGroups = v
	return s
}

type GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanStrategiesNoticeObjectList struct {
	NoticeObjectId   *int64  `json:"noticeObjectId,omitempty" xml:"noticeObjectId,omitempty"`
	NoticeObjectName *string `json:"noticeObjectName,omitempty" xml:"noticeObjectName,omitempty"`
}

func (s GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanStrategiesNoticeObjectList) String() string {
	return tea.Prettify(s)
}

func (s GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanStrategiesNoticeObjectList) GoString() string {
	return s.String()
}

func (s *GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanStrategiesNoticeObjectList) SetNoticeObjectId(v int64) *GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanStrategiesNoticeObjectList {
	s.NoticeObjectId = &v
	return s
}

func (s *GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanStrategiesNoticeObjectList) SetNoticeObjectName(v string) *GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanStrategiesNoticeObjectList {
	s.NoticeObjectName = &v
	return s
}

type GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanStrategiesServiceGroups struct {
	Id               *int64  `json:"id,omitempty" xml:"id,omitempty"`
	ServiceGroupName *string `json:"serviceGroupName,omitempty" xml:"serviceGroupName,omitempty"`
}

func (s GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanStrategiesServiceGroups) String() string {
	return tea.Prettify(s)
}

func (s GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanStrategiesServiceGroups) GoString() string {
	return s.String()
}

func (s *GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanStrategiesServiceGroups) SetId(v int64) *GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanStrategiesServiceGroups {
	s.Id = &v
	return s
}

func (s *GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanStrategiesServiceGroups) SetServiceGroupName(v string) *GetEscalationPlanResponseBodyDataEscalationPlanRulesEscalationPlanStrategiesServiceGroups {
	s.ServiceGroupName = &v
	return s
}

type GetEscalationPlanResponseBodyDataEscalationPlanScopeObjects struct {
	EscalationPlanScopeObjects *int32  `json:"escalationPlanScopeObjects,omitempty" xml:"escalationPlanScopeObjects,omitempty"`
	Scope                      *string `json:"scope,omitempty" xml:"scope,omitempty"`
	ScopeObjectDeletedType     *int32  `json:"scopeObjectDeletedType,omitempty" xml:"scopeObjectDeletedType,omitempty"`
	ScopeObjectId              *int64  `json:"scopeObjectId,omitempty" xml:"scopeObjectId,omitempty"`
	ScopeObjectName            *string `json:"scopeObjectName,omitempty" xml:"scopeObjectName,omitempty"`
}

func (s GetEscalationPlanResponseBodyDataEscalationPlanScopeObjects) String() string {
	return tea.Prettify(s)
}

func (s GetEscalationPlanResponseBodyDataEscalationPlanScopeObjects) GoString() string {
	return s.String()
}

func (s *GetEscalationPlanResponseBodyDataEscalationPlanScopeObjects) SetEscalationPlanScopeObjects(v int32) *GetEscalationPlanResponseBodyDataEscalationPlanScopeObjects {
	s.EscalationPlanScopeObjects = &v
	return s
}

func (s *GetEscalationPlanResponseBodyDataEscalationPlanScopeObjects) SetScope(v string) *GetEscalationPlanResponseBodyDataEscalationPlanScopeObjects {
	s.Scope = &v
	return s
}

func (s *GetEscalationPlanResponseBodyDataEscalationPlanScopeObjects) SetScopeObjectDeletedType(v int32) *GetEscalationPlanResponseBodyDataEscalationPlanScopeObjects {
	s.ScopeObjectDeletedType = &v
	return s
}

func (s *GetEscalationPlanResponseBodyDataEscalationPlanScopeObjects) SetScopeObjectId(v int64) *GetEscalationPlanResponseBodyDataEscalationPlanScopeObjects {
	s.ScopeObjectId = &v
	return s
}

func (s *GetEscalationPlanResponseBodyDataEscalationPlanScopeObjects) SetScopeObjectName(v string) *GetEscalationPlanResponseBodyDataEscalationPlanScopeObjects {
	s.ScopeObjectName = &v
	return s
}

type GetEscalationPlanResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetEscalationPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEscalationPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEscalationPlanResponse) GoString() string {
	return s.String()
}

func (s *GetEscalationPlanResponse) SetHeaders(v map[string]*string) *GetEscalationPlanResponse {
	s.Headers = v
	return s
}

func (s *GetEscalationPlanResponse) SetStatusCode(v int32) *GetEscalationPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEscalationPlanResponse) SetBody(v *GetEscalationPlanResponseBody) *GetEscalationPlanResponse {
	s.Body = v
	return s
}

type GetEventRequest struct {
	MonitorSourceId *int64 `json:"monitorSourceId,omitempty" xml:"monitorSourceId,omitempty"`
}

func (s GetEventRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEventRequest) GoString() string {
	return s.String()
}

func (s *GetEventRequest) SetMonitorSourceId(v int64) *GetEventRequest {
	s.MonitorSourceId = &v
	return s
}

type GetEventResponseBody struct {
	Data *GetEventResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEventResponseBody) GoString() string {
	return s.String()
}

func (s *GetEventResponseBody) SetData(v *GetEventResponseBodyData) *GetEventResponseBody {
	s.Data = v
	return s
}

func (s *GetEventResponseBody) SetRequestId(v string) *GetEventResponseBody {
	s.RequestId = &v
	return s
}

type GetEventResponseBodyData struct {
	EventJson         *string `json:"eventJson,omitempty" xml:"eventJson,omitempty"`
	EventTime         *string `json:"eventTime,omitempty" xml:"eventTime,omitempty"`
	MonitorSourceId   *int64  `json:"monitorSourceId,omitempty" xml:"monitorSourceId,omitempty"`
	MonitorSourceName *string `json:"monitorSourceName,omitempty" xml:"monitorSourceName,omitempty"`
}

func (s GetEventResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetEventResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEventResponseBodyData) SetEventJson(v string) *GetEventResponseBodyData {
	s.EventJson = &v
	return s
}

func (s *GetEventResponseBodyData) SetEventTime(v string) *GetEventResponseBodyData {
	s.EventTime = &v
	return s
}

func (s *GetEventResponseBodyData) SetMonitorSourceId(v int64) *GetEventResponseBodyData {
	s.MonitorSourceId = &v
	return s
}

func (s *GetEventResponseBodyData) SetMonitorSourceName(v string) *GetEventResponseBodyData {
	s.MonitorSourceName = &v
	return s
}

type GetEventResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetEventResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEventResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEventResponse) GoString() string {
	return s.String()
}

func (s *GetEventResponse) SetHeaders(v map[string]*string) *GetEventResponse {
	s.Headers = v
	return s
}

func (s *GetEventResponse) SetStatusCode(v int32) *GetEventResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEventResponse) SetBody(v *GetEventResponseBody) *GetEventResponse {
	s.Body = v
	return s
}

type GetHomePageGuidanceRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s GetHomePageGuidanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHomePageGuidanceRequest) GoString() string {
	return s.String()
}

func (s *GetHomePageGuidanceRequest) SetClientToken(v string) *GetHomePageGuidanceRequest {
	s.ClientToken = &v
	return s
}

type GetHomePageGuidanceResponseBody struct {
	Data *GetHomePageGuidanceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetHomePageGuidanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHomePageGuidanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetHomePageGuidanceResponseBody) SetData(v *GetHomePageGuidanceResponseBodyData) *GetHomePageGuidanceResponseBody {
	s.Data = v
	return s
}

func (s *GetHomePageGuidanceResponseBody) SetRequestId(v string) *GetHomePageGuidanceResponseBody {
	s.RequestId = &v
	return s
}

type GetHomePageGuidanceResponseBodyData struct {
	NotifySubscriptionStatus *bool `json:"notifySubscriptionStatus,omitempty" xml:"notifySubscriptionStatus,omitempty"`
	ServiceGroupStatus       *bool `json:"serviceGroupStatus,omitempty" xml:"serviceGroupStatus,omitempty"`
	ServiceStatus            *bool `json:"serviceStatus,omitempty" xml:"serviceStatus,omitempty"`
	UsersStatus              *bool `json:"usersStatus,omitempty" xml:"usersStatus,omitempty"`
}

func (s GetHomePageGuidanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHomePageGuidanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHomePageGuidanceResponseBodyData) SetNotifySubscriptionStatus(v bool) *GetHomePageGuidanceResponseBodyData {
	s.NotifySubscriptionStatus = &v
	return s
}

func (s *GetHomePageGuidanceResponseBodyData) SetServiceGroupStatus(v bool) *GetHomePageGuidanceResponseBodyData {
	s.ServiceGroupStatus = &v
	return s
}

func (s *GetHomePageGuidanceResponseBodyData) SetServiceStatus(v bool) *GetHomePageGuidanceResponseBodyData {
	s.ServiceStatus = &v
	return s
}

func (s *GetHomePageGuidanceResponseBodyData) SetUsersStatus(v bool) *GetHomePageGuidanceResponseBodyData {
	s.UsersStatus = &v
	return s
}

type GetHomePageGuidanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetHomePageGuidanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHomePageGuidanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHomePageGuidanceResponse) GoString() string {
	return s.String()
}

func (s *GetHomePageGuidanceResponse) SetHeaders(v map[string]*string) *GetHomePageGuidanceResponse {
	s.Headers = v
	return s
}

func (s *GetHomePageGuidanceResponse) SetStatusCode(v int32) *GetHomePageGuidanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHomePageGuidanceResponse) SetBody(v *GetHomePageGuidanceResponseBody) *GetHomePageGuidanceResponse {
	s.Body = v
	return s
}

type GetIncidentRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	IncidentId  *int64  `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
}

func (s GetIncidentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetIncidentRequest) GoString() string {
	return s.String()
}

func (s *GetIncidentRequest) SetClientToken(v string) *GetIncidentRequest {
	s.ClientToken = &v
	return s
}

func (s *GetIncidentRequest) SetIncidentId(v int64) *GetIncidentRequest {
	s.IncidentId = &v
	return s
}

type GetIncidentResponseBody struct {
	// data
	Data *GetIncidentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetIncidentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIncidentResponseBody) GoString() string {
	return s.String()
}

func (s *GetIncidentResponseBody) SetData(v *GetIncidentResponseBodyData) *GetIncidentResponseBody {
	s.Data = v
	return s
}

func (s *GetIncidentResponseBody) SetRequestId(v string) *GetIncidentResponseBody {
	s.RequestId = &v
	return s
}

type GetIncidentResponseBodyData struct {
	AssignToWhoIsValid        *int32    `json:"assignToWhoIsValid,omitempty" xml:"assignToWhoIsValid,omitempty"`
	AssignUserId              *int64    `json:"assignUserId,omitempty" xml:"assignUserId,omitempty"`
	AssignUserName            *string   `json:"assignUserName,omitempty" xml:"assignUserName,omitempty"`
	AssignUserPhone           *string   `json:"assignUserPhone,omitempty" xml:"assignUserPhone,omitempty"`
	CreateTime                *string   `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DefaultAssignToWho        *int32    `json:"defaultAssignToWho,omitempty" xml:"defaultAssignToWho,omitempty"`
	DefaultAssignToWhoIsValid *int32    `json:"defaultAssignToWhoIsValid,omitempty" xml:"defaultAssignToWhoIsValid,omitempty"`
	DefaultAssignToWhoName    *string   `json:"defaultAssignToWhoName,omitempty" xml:"defaultAssignToWhoName,omitempty"`
	DurationTime              *int64    `json:"durationTime,omitempty" xml:"durationTime,omitempty"`
	Effect                    *string   `json:"effect,omitempty" xml:"effect,omitempty"`
	IncidentDescription       *string   `json:"incidentDescription,omitempty" xml:"incidentDescription,omitempty"`
	IncidentId                *int64    `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
	IncidentLevel             *string   `json:"incidentLevel,omitempty" xml:"incidentLevel,omitempty"`
	IncidentNumber            *string   `json:"incidentNumber,omitempty" xml:"incidentNumber,omitempty"`
	IncidentStatus            *string   `json:"incidentStatus,omitempty" xml:"incidentStatus,omitempty"`
	IncidentTitle             *string   `json:"incidentTitle,omitempty" xml:"incidentTitle,omitempty"`
	IsManual                  *bool     `json:"isManual,omitempty" xml:"isManual,omitempty"`
	IsUpgrade                 *bool     `json:"isUpgrade,omitempty" xml:"isUpgrade,omitempty"`
	NotifyChannels            []*string `json:"notifyChannels,omitempty" xml:"notifyChannels,omitempty" type:"Repeated"`
	ProblemId                 *int64    `json:"problemId,omitempty" xml:"problemId,omitempty"`
	ProblemNumber             *string   `json:"problemNumber,omitempty" xml:"problemNumber,omitempty"`
	RelRouteRuleDeleteType    *int32    `json:"relRouteRuleDeleteType,omitempty" xml:"relRouteRuleDeleteType,omitempty"`
	RelServiceDeleteType      *int32    `json:"relServiceDeleteType,omitempty" xml:"relServiceDeleteType,omitempty"`
	RelServiceGroupIsValid    *int32    `json:"relServiceGroupIsValid,omitempty" xml:"relServiceGroupIsValid,omitempty"`
	RelatedServiceDescription *string   `json:"relatedServiceDescription,omitempty" xml:"relatedServiceDescription,omitempty"`
	RelatedServiceGroupId     *int64    `json:"relatedServiceGroupId,omitempty" xml:"relatedServiceGroupId,omitempty"`
	RelatedServiceGroupName   *string   `json:"relatedServiceGroupName,omitempty" xml:"relatedServiceGroupName,omitempty"`
	RelatedServiceId          *int64    `json:"relatedServiceId,omitempty" xml:"relatedServiceId,omitempty"`
	RelatedServiceName        *string   `json:"relatedServiceName,omitempty" xml:"relatedServiceName,omitempty"`
	RouteRuleId               *int64    `json:"routeRuleId,omitempty" xml:"routeRuleId,omitempty"`
	RouteRuleName             *string   `json:"routeRuleName,omitempty" xml:"routeRuleName,omitempty"`
}

func (s GetIncidentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetIncidentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetIncidentResponseBodyData) SetAssignToWhoIsValid(v int32) *GetIncidentResponseBodyData {
	s.AssignToWhoIsValid = &v
	return s
}

func (s *GetIncidentResponseBodyData) SetAssignUserId(v int64) *GetIncidentResponseBodyData {
	s.AssignUserId = &v
	return s
}

func (s *GetIncidentResponseBodyData) SetAssignUserName(v string) *GetIncidentResponseBodyData {
	s.AssignUserName = &v
	return s
}

func (s *GetIncidentResponseBodyData) SetAssignUserPhone(v string) *GetIncidentResponseBodyData {
	s.AssignUserPhone = &v
	return s
}

func (s *GetIncidentResponseBodyData) SetCreateTime(v string) *GetIncidentResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetIncidentResponseBodyData) SetDefaultAssignToWho(v int32) *GetIncidentResponseBodyData {
	s.DefaultAssignToWho = &v
	return s
}

func (s *GetIncidentResponseBodyData) SetDefaultAssignToWhoIsValid(v int32) *GetIncidentResponseBodyData {
	s.DefaultAssignToWhoIsValid = &v
	return s
}

func (s *GetIncidentResponseBodyData) SetDefaultAssignToWhoName(v string) *GetIncidentResponseBodyData {
	s.DefaultAssignToWhoName = &v
	return s
}

func (s *GetIncidentResponseBodyData) SetDurationTime(v int64) *GetIncidentResponseBodyData {
	s.DurationTime = &v
	return s
}

func (s *GetIncidentResponseBodyData) SetEffect(v string) *GetIncidentResponseBodyData {
	s.Effect = &v
	return s
}

func (s *GetIncidentResponseBodyData) SetIncidentDescription(v string) *GetIncidentResponseBodyData {
	s.IncidentDescription = &v
	return s
}

func (s *GetIncidentResponseBodyData) SetIncidentId(v int64) *GetIncidentResponseBodyData {
	s.IncidentId = &v
	return s
}

func (s *GetIncidentResponseBodyData) SetIncidentLevel(v string) *GetIncidentResponseBodyData {
	s.IncidentLevel = &v
	return s
}

func (s *GetIncidentResponseBodyData) SetIncidentNumber(v string) *GetIncidentResponseBodyData {
	s.IncidentNumber = &v
	return s
}

func (s *GetIncidentResponseBodyData) SetIncidentStatus(v string) *GetIncidentResponseBodyData {
	s.IncidentStatus = &v
	return s
}

func (s *GetIncidentResponseBodyData) SetIncidentTitle(v string) *GetIncidentResponseBodyData {
	s.IncidentTitle = &v
	return s
}

func (s *GetIncidentResponseBodyData) SetIsManual(v bool) *GetIncidentResponseBodyData {
	s.IsManual = &v
	return s
}

func (s *GetIncidentResponseBodyData) SetIsUpgrade(v bool) *GetIncidentResponseBodyData {
	s.IsUpgrade = &v
	return s
}

func (s *GetIncidentResponseBodyData) SetNotifyChannels(v []*string) *GetIncidentResponseBodyData {
	s.NotifyChannels = v
	return s
}

func (s *GetIncidentResponseBodyData) SetProblemId(v int64) *GetIncidentResponseBodyData {
	s.ProblemId = &v
	return s
}

func (s *GetIncidentResponseBodyData) SetProblemNumber(v string) *GetIncidentResponseBodyData {
	s.ProblemNumber = &v
	return s
}

func (s *GetIncidentResponseBodyData) SetRelRouteRuleDeleteType(v int32) *GetIncidentResponseBodyData {
	s.RelRouteRuleDeleteType = &v
	return s
}

func (s *GetIncidentResponseBodyData) SetRelServiceDeleteType(v int32) *GetIncidentResponseBodyData {
	s.RelServiceDeleteType = &v
	return s
}

func (s *GetIncidentResponseBodyData) SetRelServiceGroupIsValid(v int32) *GetIncidentResponseBodyData {
	s.RelServiceGroupIsValid = &v
	return s
}

func (s *GetIncidentResponseBodyData) SetRelatedServiceDescription(v string) *GetIncidentResponseBodyData {
	s.RelatedServiceDescription = &v
	return s
}

func (s *GetIncidentResponseBodyData) SetRelatedServiceGroupId(v int64) *GetIncidentResponseBodyData {
	s.RelatedServiceGroupId = &v
	return s
}

func (s *GetIncidentResponseBodyData) SetRelatedServiceGroupName(v string) *GetIncidentResponseBodyData {
	s.RelatedServiceGroupName = &v
	return s
}

func (s *GetIncidentResponseBodyData) SetRelatedServiceId(v int64) *GetIncidentResponseBodyData {
	s.RelatedServiceId = &v
	return s
}

func (s *GetIncidentResponseBodyData) SetRelatedServiceName(v string) *GetIncidentResponseBodyData {
	s.RelatedServiceName = &v
	return s
}

func (s *GetIncidentResponseBodyData) SetRouteRuleId(v int64) *GetIncidentResponseBodyData {
	s.RouteRuleId = &v
	return s
}

func (s *GetIncidentResponseBodyData) SetRouteRuleName(v string) *GetIncidentResponseBodyData {
	s.RouteRuleName = &v
	return s
}

type GetIncidentResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetIncidentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetIncidentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIncidentResponse) GoString() string {
	return s.String()
}

func (s *GetIncidentResponse) SetHeaders(v map[string]*string) *GetIncidentResponse {
	s.Headers = v
	return s
}

func (s *GetIncidentResponse) SetStatusCode(v int32) *GetIncidentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIncidentResponse) SetBody(v *GetIncidentResponseBody) *GetIncidentResponse {
	s.Body = v
	return s
}

type GetIncidentStatisticsRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s GetIncidentStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetIncidentStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetIncidentStatisticsRequest) SetClientToken(v string) *GetIncidentStatisticsRequest {
	s.ClientToken = &v
	return s
}

type GetIncidentStatisticsResponseBody struct {
	// data
	Data *GetIncidentStatisticsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetIncidentStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIncidentStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetIncidentStatisticsResponseBody) SetData(v *GetIncidentStatisticsResponseBodyData) *GetIncidentStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *GetIncidentStatisticsResponseBody) SetRequestId(v string) *GetIncidentStatisticsResponseBody {
	s.RequestId = &v
	return s
}

type GetIncidentStatisticsResponseBodyData struct {
	AllFinish   *int32 `json:"allFinish,omitempty" xml:"allFinish,omitempty"`
	AllResponse *int32 `json:"allResponse,omitempty" xml:"allResponse,omitempty"`
	MyFinish    *int32 `json:"myFinish,omitempty" xml:"myFinish,omitempty"`
	MyResponse  *int32 `json:"myResponse,omitempty" xml:"myResponse,omitempty"`
}

func (s GetIncidentStatisticsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetIncidentStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetIncidentStatisticsResponseBodyData) SetAllFinish(v int32) *GetIncidentStatisticsResponseBodyData {
	s.AllFinish = &v
	return s
}

func (s *GetIncidentStatisticsResponseBodyData) SetAllResponse(v int32) *GetIncidentStatisticsResponseBodyData {
	s.AllResponse = &v
	return s
}

func (s *GetIncidentStatisticsResponseBodyData) SetMyFinish(v int32) *GetIncidentStatisticsResponseBodyData {
	s.MyFinish = &v
	return s
}

func (s *GetIncidentStatisticsResponseBodyData) SetMyResponse(v int32) *GetIncidentStatisticsResponseBodyData {
	s.MyResponse = &v
	return s
}

type GetIncidentStatisticsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetIncidentStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetIncidentStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIncidentStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetIncidentStatisticsResponse) SetHeaders(v map[string]*string) *GetIncidentStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetIncidentStatisticsResponse) SetStatusCode(v int32) *GetIncidentStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIncidentStatisticsResponse) SetBody(v *GetIncidentStatisticsResponseBody) *GetIncidentStatisticsResponse {
	s.Body = v
	return s
}

type GetIncidentSubtotalCountRequest struct {
	ClientToken *string  `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	IncidentIds []*int64 `json:"incidentIds,omitempty" xml:"incidentIds,omitempty" type:"Repeated"`
}

func (s GetIncidentSubtotalCountRequest) String() string {
	return tea.Prettify(s)
}

func (s GetIncidentSubtotalCountRequest) GoString() string {
	return s.String()
}

func (s *GetIncidentSubtotalCountRequest) SetClientToken(v string) *GetIncidentSubtotalCountRequest {
	s.ClientToken = &v
	return s
}

func (s *GetIncidentSubtotalCountRequest) SetIncidentIds(v []*int64) *GetIncidentSubtotalCountRequest {
	s.IncidentIds = v
	return s
}

type GetIncidentSubtotalCountResponseBody struct {
	// data
	Data *GetIncidentSubtotalCountResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s GetIncidentSubtotalCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIncidentSubtotalCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetIncidentSubtotalCountResponseBody) SetData(v *GetIncidentSubtotalCountResponseBodyData) *GetIncidentSubtotalCountResponseBody {
	s.Data = v
	return s
}

type GetIncidentSubtotalCountResponseBodyData struct {
	// id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// map
	SubtotalCount map[string]interface{} `json:"subtotalCount,omitempty" xml:"subtotalCount,omitempty"`
}

func (s GetIncidentSubtotalCountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetIncidentSubtotalCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetIncidentSubtotalCountResponseBodyData) SetRequestId(v string) *GetIncidentSubtotalCountResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *GetIncidentSubtotalCountResponseBodyData) SetSubtotalCount(v map[string]interface{}) *GetIncidentSubtotalCountResponseBodyData {
	s.SubtotalCount = v
	return s
}

type GetIncidentSubtotalCountResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetIncidentSubtotalCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetIncidentSubtotalCountResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIncidentSubtotalCountResponse) GoString() string {
	return s.String()
}

func (s *GetIncidentSubtotalCountResponse) SetHeaders(v map[string]*string) *GetIncidentSubtotalCountResponse {
	s.Headers = v
	return s
}

func (s *GetIncidentSubtotalCountResponse) SetStatusCode(v int32) *GetIncidentSubtotalCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIncidentSubtotalCountResponse) SetBody(v *GetIncidentSubtotalCountResponseBody) *GetIncidentSubtotalCountResponse {
	s.Body = v
	return s
}

type GetIntegrationConfigRequest struct {
	ClientToken         *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	IntegrationConfigId *int64  `json:"integrationConfigId,omitempty" xml:"integrationConfigId,omitempty"`
}

func (s GetIntegrationConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetIntegrationConfigRequest) GoString() string {
	return s.String()
}

func (s *GetIntegrationConfigRequest) SetClientToken(v string) *GetIntegrationConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *GetIntegrationConfigRequest) SetIntegrationConfigId(v int64) *GetIntegrationConfigRequest {
	s.IntegrationConfigId = &v
	return s
}

type GetIntegrationConfigResponseBody struct {
	// data
	Data *GetIntegrationConfigResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetIntegrationConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIntegrationConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetIntegrationConfigResponseBody) SetData(v *GetIntegrationConfigResponseBodyData) *GetIntegrationConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetIntegrationConfigResponseBody) SetRequestId(v string) *GetIntegrationConfigResponseBody {
	s.RequestId = &v
	return s
}

type GetIntegrationConfigResponseBodyData struct {
	AccessKey              *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	IntegrationConfigId    *int64  `json:"integrationConfigId,omitempty" xml:"integrationConfigId,omitempty"`
	IsReceivedEvent        *bool   `json:"isReceivedEvent,omitempty" xml:"isReceivedEvent,omitempty"`
	MonitorSourceId        *int64  `json:"monitorSourceId,omitempty" xml:"monitorSourceId,omitempty"`
	MonitorSourceName      *string `json:"monitorSourceName,omitempty" xml:"monitorSourceName,omitempty"`
	MonitorSourceShortName *string `json:"monitorSourceShortName,omitempty" xml:"monitorSourceShortName,omitempty"`
	Status                 *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetIntegrationConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetIntegrationConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetIntegrationConfigResponseBodyData) SetAccessKey(v string) *GetIntegrationConfigResponseBodyData {
	s.AccessKey = &v
	return s
}

func (s *GetIntegrationConfigResponseBodyData) SetIntegrationConfigId(v int64) *GetIntegrationConfigResponseBodyData {
	s.IntegrationConfigId = &v
	return s
}

func (s *GetIntegrationConfigResponseBodyData) SetIsReceivedEvent(v bool) *GetIntegrationConfigResponseBodyData {
	s.IsReceivedEvent = &v
	return s
}

func (s *GetIntegrationConfigResponseBodyData) SetMonitorSourceId(v int64) *GetIntegrationConfigResponseBodyData {
	s.MonitorSourceId = &v
	return s
}

func (s *GetIntegrationConfigResponseBodyData) SetMonitorSourceName(v string) *GetIntegrationConfigResponseBodyData {
	s.MonitorSourceName = &v
	return s
}

func (s *GetIntegrationConfigResponseBodyData) SetMonitorSourceShortName(v string) *GetIntegrationConfigResponseBodyData {
	s.MonitorSourceShortName = &v
	return s
}

func (s *GetIntegrationConfigResponseBodyData) SetStatus(v string) *GetIntegrationConfigResponseBodyData {
	s.Status = &v
	return s
}

type GetIntegrationConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetIntegrationConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetIntegrationConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIntegrationConfigResponse) GoString() string {
	return s.String()
}

func (s *GetIntegrationConfigResponse) SetHeaders(v map[string]*string) *GetIntegrationConfigResponse {
	s.Headers = v
	return s
}

func (s *GetIntegrationConfigResponse) SetStatusCode(v int32) *GetIntegrationConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIntegrationConfigResponse) SetBody(v *GetIntegrationConfigResponseBody) *GetIntegrationConfigResponse {
	s.Body = v
	return s
}

type GetProblemRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	ProblemId   *int64  `json:"problemId,omitempty" xml:"problemId,omitempty"`
}

func (s GetProblemRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProblemRequest) GoString() string {
	return s.String()
}

func (s *GetProblemRequest) SetClientToken(v string) *GetProblemRequest {
	s.ClientToken = &v
	return s
}

func (s *GetProblemRequest) SetProblemId(v int64) *GetProblemRequest {
	s.ProblemId = &v
	return s
}

type GetProblemResponseBody struct {
	Data      *GetProblemResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	RequestId *string                     `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetProblemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProblemResponseBody) GoString() string {
	return s.String()
}

func (s *GetProblemResponseBody) SetData(v *GetProblemResponseBodyData) *GetProblemResponseBody {
	s.Data = v
	return s
}

func (s *GetProblemResponseBody) SetRequestId(v string) *GetProblemResponseBody {
	s.RequestId = &v
	return s
}

type GetProblemResponseBodyData struct {
	CancelProblemOperateLogs  []*GetProblemResponseBodyDataCancelProblemOperateLogs  `json:"cancelProblemOperateLogs,omitempty" xml:"cancelProblemOperateLogs,omitempty" type:"Repeated"`
	CancelReason              *int64                                                 `json:"cancelReason,omitempty" xml:"cancelReason,omitempty"`
	CancelReasonDescription   *string                                                `json:"cancelReasonDescription,omitempty" xml:"cancelReasonDescription,omitempty"`
	CoordinationGroups        []*GetProblemResponseBodyDataCoordinationGroups        `json:"coordinationGroups,omitempty" xml:"coordinationGroups,omitempty" type:"Repeated"`
	CreateTime                *string                                                `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DiscoverTime              *string                                                `json:"discoverTime,omitempty" xml:"discoverTime,omitempty"`
	DurationTime              *int64                                                 `json:"durationTime,omitempty" xml:"durationTime,omitempty"`
	EffectionServices         []*GetProblemResponseBodyDataEffectionServices         `json:"effectionServices,omitempty" xml:"effectionServices,omitempty" type:"Repeated"`
	Feedback                  *string                                                `json:"feedback,omitempty" xml:"feedback,omitempty"`
	HandingProblemOperateLogs []*GetProblemResponseBodyDataHandingProblemOperateLogs `json:"handingProblemOperateLogs,omitempty" xml:"handingProblemOperateLogs,omitempty" type:"Repeated"`
	IncidentId                *int64                                                 `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
	IncidentNumber            *string                                                `json:"incidentNumber,omitempty" xml:"incidentNumber,omitempty"`
	MainHandler               *int64                                                 `json:"mainHandler,omitempty" xml:"mainHandler,omitempty"`
	MainHandlerId             *int64                                                 `json:"mainHandlerId,omitempty" xml:"mainHandlerId,omitempty"`
	MainHandlerIsValid        *int64                                                 `json:"mainHandlerIsValid,omitempty" xml:"mainHandlerIsValid,omitempty"`
	MainHandlerPhone          *string                                                `json:"mainHandlerPhone,omitempty" xml:"mainHandlerPhone,omitempty"`
	PreliminaryReason         *string                                                `json:"preliminaryReason,omitempty" xml:"preliminaryReason,omitempty"`
	// ID
	ProblemId                   *int64                                                   `json:"problemId,omitempty" xml:"problemId,omitempty"`
	ProblemLevel                *int32                                                   `json:"problemLevel,omitempty" xml:"problemLevel,omitempty"`
	ProblemName                 *string                                                  `json:"problemName,omitempty" xml:"problemName,omitempty"`
	ProblemNumber               *string                                                  `json:"problemNumber,omitempty" xml:"problemNumber,omitempty"`
	ProblemStatus               *int32                                                   `json:"problemStatus,omitempty" xml:"problemStatus,omitempty"`
	ProgressSummary             *string                                                  `json:"progressSummary,omitempty" xml:"progressSummary,omitempty"`
	ProgressSummaryRichTextId   *int64                                                   `json:"progressSummaryRichTextId,omitempty" xml:"progressSummaryRichTextId,omitempty"`
	RecoveryTime                *string                                                  `json:"recoveryTime,omitempty" xml:"recoveryTime,omitempty"`
	RelatedServiceId            *int64                                                   `json:"relatedServiceId,omitempty" xml:"relatedServiceId,omitempty"`
	ReplayProblemOperateLogs    []*GetProblemResponseBodyDataReplayProblemOperateLogs    `json:"replayProblemOperateLogs,omitempty" xml:"replayProblemOperateLogs,omitempty" type:"Repeated"`
	ReplayingProblemOperateLogs []*GetProblemResponseBodyDataReplayingProblemOperateLogs `json:"replayingProblemOperateLogs,omitempty" xml:"replayingProblemOperateLogs,omitempty" type:"Repeated"`
	RestoredProblemOperateLogs  []*GetProblemResponseBodyDataRestoredProblemOperateLogs  `json:"restoredProblemOperateLogs,omitempty" xml:"restoredProblemOperateLogs,omitempty" type:"Repeated"`
	// serviceDeleteType
	ServiceDeleteType *int32                                 `json:"serviceDeleteType,omitempty" xml:"serviceDeleteType,omitempty"`
	ServiceName       *string                                `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	Timelines         []*GetProblemResponseBodyDataTimelines `json:"timelines,omitempty" xml:"timelines,omitempty" type:"Repeated"`
}

func (s GetProblemResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetProblemResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProblemResponseBodyData) SetCancelProblemOperateLogs(v []*GetProblemResponseBodyDataCancelProblemOperateLogs) *GetProblemResponseBodyData {
	s.CancelProblemOperateLogs = v
	return s
}

func (s *GetProblemResponseBodyData) SetCancelReason(v int64) *GetProblemResponseBodyData {
	s.CancelReason = &v
	return s
}

func (s *GetProblemResponseBodyData) SetCancelReasonDescription(v string) *GetProblemResponseBodyData {
	s.CancelReasonDescription = &v
	return s
}

func (s *GetProblemResponseBodyData) SetCoordinationGroups(v []*GetProblemResponseBodyDataCoordinationGroups) *GetProblemResponseBodyData {
	s.CoordinationGroups = v
	return s
}

func (s *GetProblemResponseBodyData) SetCreateTime(v string) *GetProblemResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetProblemResponseBodyData) SetDiscoverTime(v string) *GetProblemResponseBodyData {
	s.DiscoverTime = &v
	return s
}

func (s *GetProblemResponseBodyData) SetDurationTime(v int64) *GetProblemResponseBodyData {
	s.DurationTime = &v
	return s
}

func (s *GetProblemResponseBodyData) SetEffectionServices(v []*GetProblemResponseBodyDataEffectionServices) *GetProblemResponseBodyData {
	s.EffectionServices = v
	return s
}

func (s *GetProblemResponseBodyData) SetFeedback(v string) *GetProblemResponseBodyData {
	s.Feedback = &v
	return s
}

func (s *GetProblemResponseBodyData) SetHandingProblemOperateLogs(v []*GetProblemResponseBodyDataHandingProblemOperateLogs) *GetProblemResponseBodyData {
	s.HandingProblemOperateLogs = v
	return s
}

func (s *GetProblemResponseBodyData) SetIncidentId(v int64) *GetProblemResponseBodyData {
	s.IncidentId = &v
	return s
}

func (s *GetProblemResponseBodyData) SetIncidentNumber(v string) *GetProblemResponseBodyData {
	s.IncidentNumber = &v
	return s
}

func (s *GetProblemResponseBodyData) SetMainHandler(v int64) *GetProblemResponseBodyData {
	s.MainHandler = &v
	return s
}

func (s *GetProblemResponseBodyData) SetMainHandlerId(v int64) *GetProblemResponseBodyData {
	s.MainHandlerId = &v
	return s
}

func (s *GetProblemResponseBodyData) SetMainHandlerIsValid(v int64) *GetProblemResponseBodyData {
	s.MainHandlerIsValid = &v
	return s
}

func (s *GetProblemResponseBodyData) SetMainHandlerPhone(v string) *GetProblemResponseBodyData {
	s.MainHandlerPhone = &v
	return s
}

func (s *GetProblemResponseBodyData) SetPreliminaryReason(v string) *GetProblemResponseBodyData {
	s.PreliminaryReason = &v
	return s
}

func (s *GetProblemResponseBodyData) SetProblemId(v int64) *GetProblemResponseBodyData {
	s.ProblemId = &v
	return s
}

func (s *GetProblemResponseBodyData) SetProblemLevel(v int32) *GetProblemResponseBodyData {
	s.ProblemLevel = &v
	return s
}

func (s *GetProblemResponseBodyData) SetProblemName(v string) *GetProblemResponseBodyData {
	s.ProblemName = &v
	return s
}

func (s *GetProblemResponseBodyData) SetProblemNumber(v string) *GetProblemResponseBodyData {
	s.ProblemNumber = &v
	return s
}

func (s *GetProblemResponseBodyData) SetProblemStatus(v int32) *GetProblemResponseBodyData {
	s.ProblemStatus = &v
	return s
}

func (s *GetProblemResponseBodyData) SetProgressSummary(v string) *GetProblemResponseBodyData {
	s.ProgressSummary = &v
	return s
}

func (s *GetProblemResponseBodyData) SetProgressSummaryRichTextId(v int64) *GetProblemResponseBodyData {
	s.ProgressSummaryRichTextId = &v
	return s
}

func (s *GetProblemResponseBodyData) SetRecoveryTime(v string) *GetProblemResponseBodyData {
	s.RecoveryTime = &v
	return s
}

func (s *GetProblemResponseBodyData) SetRelatedServiceId(v int64) *GetProblemResponseBodyData {
	s.RelatedServiceId = &v
	return s
}

func (s *GetProblemResponseBodyData) SetReplayProblemOperateLogs(v []*GetProblemResponseBodyDataReplayProblemOperateLogs) *GetProblemResponseBodyData {
	s.ReplayProblemOperateLogs = v
	return s
}

func (s *GetProblemResponseBodyData) SetReplayingProblemOperateLogs(v []*GetProblemResponseBodyDataReplayingProblemOperateLogs) *GetProblemResponseBodyData {
	s.ReplayingProblemOperateLogs = v
	return s
}

func (s *GetProblemResponseBodyData) SetRestoredProblemOperateLogs(v []*GetProblemResponseBodyDataRestoredProblemOperateLogs) *GetProblemResponseBodyData {
	s.RestoredProblemOperateLogs = v
	return s
}

func (s *GetProblemResponseBodyData) SetServiceDeleteType(v int32) *GetProblemResponseBodyData {
	s.ServiceDeleteType = &v
	return s
}

func (s *GetProblemResponseBodyData) SetServiceName(v string) *GetProblemResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *GetProblemResponseBodyData) SetTimelines(v []*GetProblemResponseBodyDataTimelines) *GetProblemResponseBodyData {
	s.Timelines = v
	return s
}

type GetProblemResponseBodyDataCancelProblemOperateLogs struct {
	ActionName *string `json:"actionName,omitempty" xml:"actionName,omitempty"`
	ActionTime *string `json:"actionTime,omitempty" xml:"actionTime,omitempty"`
	Operator   *string `json:"operator,omitempty" xml:"operator,omitempty"`
	UserId     *int64  `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetProblemResponseBodyDataCancelProblemOperateLogs) String() string {
	return tea.Prettify(s)
}

func (s GetProblemResponseBodyDataCancelProblemOperateLogs) GoString() string {
	return s.String()
}

func (s *GetProblemResponseBodyDataCancelProblemOperateLogs) SetActionName(v string) *GetProblemResponseBodyDataCancelProblemOperateLogs {
	s.ActionName = &v
	return s
}

func (s *GetProblemResponseBodyDataCancelProblemOperateLogs) SetActionTime(v string) *GetProblemResponseBodyDataCancelProblemOperateLogs {
	s.ActionTime = &v
	return s
}

func (s *GetProblemResponseBodyDataCancelProblemOperateLogs) SetOperator(v string) *GetProblemResponseBodyDataCancelProblemOperateLogs {
	s.Operator = &v
	return s
}

func (s *GetProblemResponseBodyDataCancelProblemOperateLogs) SetUserId(v int64) *GetProblemResponseBodyDataCancelProblemOperateLogs {
	s.UserId = &v
	return s
}

type GetProblemResponseBodyDataCoordinationGroups struct {
	IsValid          *int64  `json:"isValid,omitempty" xml:"isValid,omitempty"`
	ServiceGroupId   *int64  `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
	ServiceGroupName *string `json:"serviceGroupName,omitempty" xml:"serviceGroupName,omitempty"`
}

func (s GetProblemResponseBodyDataCoordinationGroups) String() string {
	return tea.Prettify(s)
}

func (s GetProblemResponseBodyDataCoordinationGroups) GoString() string {
	return s.String()
}

func (s *GetProblemResponseBodyDataCoordinationGroups) SetIsValid(v int64) *GetProblemResponseBodyDataCoordinationGroups {
	s.IsValid = &v
	return s
}

func (s *GetProblemResponseBodyDataCoordinationGroups) SetServiceGroupId(v int64) *GetProblemResponseBodyDataCoordinationGroups {
	s.ServiceGroupId = &v
	return s
}

func (s *GetProblemResponseBodyDataCoordinationGroups) SetServiceGroupName(v string) *GetProblemResponseBodyDataCoordinationGroups {
	s.ServiceGroupName = &v
	return s
}

type GetProblemResponseBodyDataEffectionServices struct {
	Description        *string `json:"description,omitempty" xml:"description,omitempty"`
	EffectionLevel     *int64  `json:"effectionLevel,omitempty" xml:"effectionLevel,omitempty"`
	EffectionServiceId *int64  `json:"effectionServiceId,omitempty" xml:"effectionServiceId,omitempty"`
	EffectionStatus    *int32  `json:"effectionStatus,omitempty" xml:"effectionStatus,omitempty"`
	ServiceDeleteType  *int32  `json:"serviceDeleteType,omitempty" xml:"serviceDeleteType,omitempty"`
	ServiceName        *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
}

func (s GetProblemResponseBodyDataEffectionServices) String() string {
	return tea.Prettify(s)
}

func (s GetProblemResponseBodyDataEffectionServices) GoString() string {
	return s.String()
}

func (s *GetProblemResponseBodyDataEffectionServices) SetDescription(v string) *GetProblemResponseBodyDataEffectionServices {
	s.Description = &v
	return s
}

func (s *GetProblemResponseBodyDataEffectionServices) SetEffectionLevel(v int64) *GetProblemResponseBodyDataEffectionServices {
	s.EffectionLevel = &v
	return s
}

func (s *GetProblemResponseBodyDataEffectionServices) SetEffectionServiceId(v int64) *GetProblemResponseBodyDataEffectionServices {
	s.EffectionServiceId = &v
	return s
}

func (s *GetProblemResponseBodyDataEffectionServices) SetEffectionStatus(v int32) *GetProblemResponseBodyDataEffectionServices {
	s.EffectionStatus = &v
	return s
}

func (s *GetProblemResponseBodyDataEffectionServices) SetServiceDeleteType(v int32) *GetProblemResponseBodyDataEffectionServices {
	s.ServiceDeleteType = &v
	return s
}

func (s *GetProblemResponseBodyDataEffectionServices) SetServiceName(v string) *GetProblemResponseBodyDataEffectionServices {
	s.ServiceName = &v
	return s
}

type GetProblemResponseBodyDataHandingProblemOperateLogs struct {
	ActionName  *string `json:"actionName,omitempty" xml:"actionName,omitempty"`
	ActionTime  *string `json:"actionTime,omitempty" xml:"actionTime,omitempty"`
	Operator    *string `json:"operator,omitempty" xml:"operator,omitempty"`
	UserId      *int64  `json:"userId,omitempty" xml:"userId,omitempty"`
	UserIsValid *int64  `json:"userIsValid,omitempty" xml:"userIsValid,omitempty"`
}

func (s GetProblemResponseBodyDataHandingProblemOperateLogs) String() string {
	return tea.Prettify(s)
}

func (s GetProblemResponseBodyDataHandingProblemOperateLogs) GoString() string {
	return s.String()
}

func (s *GetProblemResponseBodyDataHandingProblemOperateLogs) SetActionName(v string) *GetProblemResponseBodyDataHandingProblemOperateLogs {
	s.ActionName = &v
	return s
}

func (s *GetProblemResponseBodyDataHandingProblemOperateLogs) SetActionTime(v string) *GetProblemResponseBodyDataHandingProblemOperateLogs {
	s.ActionTime = &v
	return s
}

func (s *GetProblemResponseBodyDataHandingProblemOperateLogs) SetOperator(v string) *GetProblemResponseBodyDataHandingProblemOperateLogs {
	s.Operator = &v
	return s
}

func (s *GetProblemResponseBodyDataHandingProblemOperateLogs) SetUserId(v int64) *GetProblemResponseBodyDataHandingProblemOperateLogs {
	s.UserId = &v
	return s
}

func (s *GetProblemResponseBodyDataHandingProblemOperateLogs) SetUserIsValid(v int64) *GetProblemResponseBodyDataHandingProblemOperateLogs {
	s.UserIsValid = &v
	return s
}

type GetProblemResponseBodyDataReplayProblemOperateLogs struct {
	ActionName  *string `json:"actionName,omitempty" xml:"actionName,omitempty"`
	ActionTime  *string `json:"actionTime,omitempty" xml:"actionTime,omitempty"`
	Operator    *string `json:"operator,omitempty" xml:"operator,omitempty"`
	UserId      *int64  `json:"userId,omitempty" xml:"userId,omitempty"`
	UserIsValid *int64  `json:"userIsValid,omitempty" xml:"userIsValid,omitempty"`
}

func (s GetProblemResponseBodyDataReplayProblemOperateLogs) String() string {
	return tea.Prettify(s)
}

func (s GetProblemResponseBodyDataReplayProblemOperateLogs) GoString() string {
	return s.String()
}

func (s *GetProblemResponseBodyDataReplayProblemOperateLogs) SetActionName(v string) *GetProblemResponseBodyDataReplayProblemOperateLogs {
	s.ActionName = &v
	return s
}

func (s *GetProblemResponseBodyDataReplayProblemOperateLogs) SetActionTime(v string) *GetProblemResponseBodyDataReplayProblemOperateLogs {
	s.ActionTime = &v
	return s
}

func (s *GetProblemResponseBodyDataReplayProblemOperateLogs) SetOperator(v string) *GetProblemResponseBodyDataReplayProblemOperateLogs {
	s.Operator = &v
	return s
}

func (s *GetProblemResponseBodyDataReplayProblemOperateLogs) SetUserId(v int64) *GetProblemResponseBodyDataReplayProblemOperateLogs {
	s.UserId = &v
	return s
}

func (s *GetProblemResponseBodyDataReplayProblemOperateLogs) SetUserIsValid(v int64) *GetProblemResponseBodyDataReplayProblemOperateLogs {
	s.UserIsValid = &v
	return s
}

type GetProblemResponseBodyDataReplayingProblemOperateLogs struct {
	ActionName  *string `json:"actionName,omitempty" xml:"actionName,omitempty"`
	ActionTime  *string `json:"actionTime,omitempty" xml:"actionTime,omitempty"`
	Operator    *string `json:"operator,omitempty" xml:"operator,omitempty"`
	UserId      *int64  `json:"userId,omitempty" xml:"userId,omitempty"`
	UserIsValid *int64  `json:"userIsValid,omitempty" xml:"userIsValid,omitempty"`
}

func (s GetProblemResponseBodyDataReplayingProblemOperateLogs) String() string {
	return tea.Prettify(s)
}

func (s GetProblemResponseBodyDataReplayingProblemOperateLogs) GoString() string {
	return s.String()
}

func (s *GetProblemResponseBodyDataReplayingProblemOperateLogs) SetActionName(v string) *GetProblemResponseBodyDataReplayingProblemOperateLogs {
	s.ActionName = &v
	return s
}

func (s *GetProblemResponseBodyDataReplayingProblemOperateLogs) SetActionTime(v string) *GetProblemResponseBodyDataReplayingProblemOperateLogs {
	s.ActionTime = &v
	return s
}

func (s *GetProblemResponseBodyDataReplayingProblemOperateLogs) SetOperator(v string) *GetProblemResponseBodyDataReplayingProblemOperateLogs {
	s.Operator = &v
	return s
}

func (s *GetProblemResponseBodyDataReplayingProblemOperateLogs) SetUserId(v int64) *GetProblemResponseBodyDataReplayingProblemOperateLogs {
	s.UserId = &v
	return s
}

func (s *GetProblemResponseBodyDataReplayingProblemOperateLogs) SetUserIsValid(v int64) *GetProblemResponseBodyDataReplayingProblemOperateLogs {
	s.UserIsValid = &v
	return s
}

type GetProblemResponseBodyDataRestoredProblemOperateLogs struct {
	ActionName  *string `json:"actionName,omitempty" xml:"actionName,omitempty"`
	ActionTime  *string `json:"actionTime,omitempty" xml:"actionTime,omitempty"`
	Operator    *string `json:"operator,omitempty" xml:"operator,omitempty"`
	UserId      *int64  `json:"userId,omitempty" xml:"userId,omitempty"`
	UserIsValid *int64  `json:"userIsValid,omitempty" xml:"userIsValid,omitempty"`
}

func (s GetProblemResponseBodyDataRestoredProblemOperateLogs) String() string {
	return tea.Prettify(s)
}

func (s GetProblemResponseBodyDataRestoredProblemOperateLogs) GoString() string {
	return s.String()
}

func (s *GetProblemResponseBodyDataRestoredProblemOperateLogs) SetActionName(v string) *GetProblemResponseBodyDataRestoredProblemOperateLogs {
	s.ActionName = &v
	return s
}

func (s *GetProblemResponseBodyDataRestoredProblemOperateLogs) SetActionTime(v string) *GetProblemResponseBodyDataRestoredProblemOperateLogs {
	s.ActionTime = &v
	return s
}

func (s *GetProblemResponseBodyDataRestoredProblemOperateLogs) SetOperator(v string) *GetProblemResponseBodyDataRestoredProblemOperateLogs {
	s.Operator = &v
	return s
}

func (s *GetProblemResponseBodyDataRestoredProblemOperateLogs) SetUserId(v int64) *GetProblemResponseBodyDataRestoredProblemOperateLogs {
	s.UserId = &v
	return s
}

func (s *GetProblemResponseBodyDataRestoredProblemOperateLogs) SetUserIsValid(v int64) *GetProblemResponseBodyDataRestoredProblemOperateLogs {
	s.UserIsValid = &v
	return s
}

type GetProblemResponseBodyDataTimelines struct {
	KeyNode *string `json:"keyNode,omitempty" xml:"keyNode,omitempty"`
}

func (s GetProblemResponseBodyDataTimelines) String() string {
	return tea.Prettify(s)
}

func (s GetProblemResponseBodyDataTimelines) GoString() string {
	return s.String()
}

func (s *GetProblemResponseBodyDataTimelines) SetKeyNode(v string) *GetProblemResponseBodyDataTimelines {
	s.KeyNode = &v
	return s
}

type GetProblemResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetProblemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProblemResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProblemResponse) GoString() string {
	return s.String()
}

func (s *GetProblemResponse) SetHeaders(v map[string]*string) *GetProblemResponse {
	s.Headers = v
	return s
}

func (s *GetProblemResponse) SetStatusCode(v int32) *GetProblemResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProblemResponse) SetBody(v *GetProblemResponseBody) *GetProblemResponse {
	s.Body = v
	return s
}

type GetProblemEffectionServiceRequest struct {
	// clientToken
	ClientToken        *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	EffectionServiceId *int64  `json:"effectionServiceId,omitempty" xml:"effectionServiceId,omitempty"`
	ProblemId          *int64  `json:"problemId,omitempty" xml:"problemId,omitempty"`
}

func (s GetProblemEffectionServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProblemEffectionServiceRequest) GoString() string {
	return s.String()
}

func (s *GetProblemEffectionServiceRequest) SetClientToken(v string) *GetProblemEffectionServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *GetProblemEffectionServiceRequest) SetEffectionServiceId(v int64) *GetProblemEffectionServiceRequest {
	s.EffectionServiceId = &v
	return s
}

func (s *GetProblemEffectionServiceRequest) SetProblemId(v int64) *GetProblemEffectionServiceRequest {
	s.ProblemId = &v
	return s
}

type GetProblemEffectionServiceResponseBody struct {
	// data
	Data *GetProblemEffectionServiceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetProblemEffectionServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProblemEffectionServiceResponseBody) GoString() string {
	return s.String()
}

func (s *GetProblemEffectionServiceResponseBody) SetData(v *GetProblemEffectionServiceResponseBodyData) *GetProblemEffectionServiceResponseBody {
	s.Data = v
	return s
}

func (s *GetProblemEffectionServiceResponseBody) SetRequestId(v string) *GetProblemEffectionServiceResponseBody {
	s.RequestId = &v
	return s
}

type GetProblemEffectionServiceResponseBodyData struct {
	Description        *string   `json:"description,omitempty" xml:"description,omitempty"`
	EffectionServiceId *int64    `json:"effectionServiceId,omitempty" xml:"effectionServiceId,omitempty"`
	Level              *int64    `json:"level,omitempty" xml:"level,omitempty"`
	PicUrl             []*string `json:"picUrl,omitempty" xml:"picUrl,omitempty" type:"Repeated"`
	ServiceId          *int64    `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	ServiceName        *string   `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	Status             *int64    `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetProblemEffectionServiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetProblemEffectionServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProblemEffectionServiceResponseBodyData) SetDescription(v string) *GetProblemEffectionServiceResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetProblemEffectionServiceResponseBodyData) SetEffectionServiceId(v int64) *GetProblemEffectionServiceResponseBodyData {
	s.EffectionServiceId = &v
	return s
}

func (s *GetProblemEffectionServiceResponseBodyData) SetLevel(v int64) *GetProblemEffectionServiceResponseBodyData {
	s.Level = &v
	return s
}

func (s *GetProblemEffectionServiceResponseBodyData) SetPicUrl(v []*string) *GetProblemEffectionServiceResponseBodyData {
	s.PicUrl = v
	return s
}

func (s *GetProblemEffectionServiceResponseBodyData) SetServiceId(v int64) *GetProblemEffectionServiceResponseBodyData {
	s.ServiceId = &v
	return s
}

func (s *GetProblemEffectionServiceResponseBodyData) SetServiceName(v string) *GetProblemEffectionServiceResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *GetProblemEffectionServiceResponseBodyData) SetStatus(v int64) *GetProblemEffectionServiceResponseBodyData {
	s.Status = &v
	return s
}

type GetProblemEffectionServiceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetProblemEffectionServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProblemEffectionServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProblemEffectionServiceResponse) GoString() string {
	return s.String()
}

func (s *GetProblemEffectionServiceResponse) SetHeaders(v map[string]*string) *GetProblemEffectionServiceResponse {
	s.Headers = v
	return s
}

func (s *GetProblemEffectionServiceResponse) SetStatusCode(v int32) *GetProblemEffectionServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProblemEffectionServiceResponse) SetBody(v *GetProblemEffectionServiceResponseBody) *GetProblemEffectionServiceResponse {
	s.Body = v
	return s
}

type GetProblemImprovementRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	ProblemId   *string `json:"problemId,omitempty" xml:"problemId,omitempty"`
}

func (s GetProblemImprovementRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProblemImprovementRequest) GoString() string {
	return s.String()
}

func (s *GetProblemImprovementRequest) SetClientToken(v string) *GetProblemImprovementRequest {
	s.ClientToken = &v
	return s
}

func (s *GetProblemImprovementRequest) SetProblemId(v string) *GetProblemImprovementRequest {
	s.ProblemId = &v
	return s
}

type GetProblemImprovementResponseBody struct {
	// data
	Data *GetProblemImprovementResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetProblemImprovementResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProblemImprovementResponseBody) GoString() string {
	return s.String()
}

func (s *GetProblemImprovementResponseBody) SetData(v *GetProblemImprovementResponseBodyData) *GetProblemImprovementResponseBody {
	s.Data = v
	return s
}

func (s *GetProblemImprovementResponseBody) SetRequestId(v string) *GetProblemImprovementResponseBody {
	s.RequestId = &v
	return s
}

type GetProblemImprovementResponseBodyData struct {
	CustomProblemReason   *string                                             `json:"customProblemReason,omitempty" xml:"customProblemReason,omitempty"`
	DiscoverSource        *string                                             `json:"discoverSource,omitempty" xml:"discoverSource,omitempty"`
	DutyDepartmentId      *string                                             `json:"dutyDepartmentId,omitempty" xml:"dutyDepartmentId,omitempty"`
	DutyDepartmentName    *string                                             `json:"dutyDepartmentName,omitempty" xml:"dutyDepartmentName,omitempty"`
	DutyUserId            *int64                                              `json:"dutyUserId,omitempty" xml:"dutyUserId,omitempty"`
	DutyUserIsValid       *int64                                              `json:"dutyUserIsValid,omitempty" xml:"dutyUserIsValid,omitempty"`
	DutyUserName          *string                                             `json:"dutyUserName,omitempty" xml:"dutyUserName,omitempty"`
	DutyUserPhone         *string                                             `json:"dutyUserPhone,omitempty" xml:"dutyUserPhone,omitempty"`
	InjectionMode         *string                                             `json:"injectionMode,omitempty" xml:"injectionMode,omitempty"`
	IsManual              *bool                                               `json:"isManual,omitempty" xml:"isManual,omitempty"`
	MeasureList           []*GetProblemImprovementResponseBodyDataMeasureList `json:"measureList,omitempty" xml:"measureList,omitempty" type:"Repeated"`
	MonitorSourceName     *string                                             `json:"monitorSourceName,omitempty" xml:"monitorSourceName,omitempty"`
	ProblemId             *string                                             `json:"problemId,omitempty" xml:"problemId,omitempty"`
	ProblemReason         *string                                             `json:"problemReason,omitempty" xml:"problemReason,omitempty"`
	RecentActivity        *string                                             `json:"recentActivity,omitempty" xml:"recentActivity,omitempty"`
	RecoveryMode          *string                                             `json:"recoveryMode,omitempty" xml:"recoveryMode,omitempty"`
	RelationChanges       *string                                             `json:"relationChanges,omitempty" xml:"relationChanges,omitempty"`
	Remark                *string                                             `json:"remark,omitempty" xml:"remark,omitempty"`
	ReplayDutyUserId      *int64                                              `json:"replayDutyUserId,omitempty" xml:"replayDutyUserId,omitempty"`
	ReplayDutyUserIsValid *int64                                              `json:"replayDutyUserIsValid,omitempty" xml:"replayDutyUserIsValid,omitempty"`
	ReplayDutyUserName    *string                                             `json:"replayDutyUserName,omitempty" xml:"replayDutyUserName,omitempty"`
	ReplayDutyUserPhone   *string                                             `json:"replayDutyUserPhone,omitempty" xml:"replayDutyUserPhone,omitempty"`
	UserReport            *int64                                              `json:"userReport,omitempty" xml:"userReport,omitempty"`
}

func (s GetProblemImprovementResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetProblemImprovementResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProblemImprovementResponseBodyData) SetCustomProblemReason(v string) *GetProblemImprovementResponseBodyData {
	s.CustomProblemReason = &v
	return s
}

func (s *GetProblemImprovementResponseBodyData) SetDiscoverSource(v string) *GetProblemImprovementResponseBodyData {
	s.DiscoverSource = &v
	return s
}

func (s *GetProblemImprovementResponseBodyData) SetDutyDepartmentId(v string) *GetProblemImprovementResponseBodyData {
	s.DutyDepartmentId = &v
	return s
}

func (s *GetProblemImprovementResponseBodyData) SetDutyDepartmentName(v string) *GetProblemImprovementResponseBodyData {
	s.DutyDepartmentName = &v
	return s
}

func (s *GetProblemImprovementResponseBodyData) SetDutyUserId(v int64) *GetProblemImprovementResponseBodyData {
	s.DutyUserId = &v
	return s
}

func (s *GetProblemImprovementResponseBodyData) SetDutyUserIsValid(v int64) *GetProblemImprovementResponseBodyData {
	s.DutyUserIsValid = &v
	return s
}

func (s *GetProblemImprovementResponseBodyData) SetDutyUserName(v string) *GetProblemImprovementResponseBodyData {
	s.DutyUserName = &v
	return s
}

func (s *GetProblemImprovementResponseBodyData) SetDutyUserPhone(v string) *GetProblemImprovementResponseBodyData {
	s.DutyUserPhone = &v
	return s
}

func (s *GetProblemImprovementResponseBodyData) SetInjectionMode(v string) *GetProblemImprovementResponseBodyData {
	s.InjectionMode = &v
	return s
}

func (s *GetProblemImprovementResponseBodyData) SetIsManual(v bool) *GetProblemImprovementResponseBodyData {
	s.IsManual = &v
	return s
}

func (s *GetProblemImprovementResponseBodyData) SetMeasureList(v []*GetProblemImprovementResponseBodyDataMeasureList) *GetProblemImprovementResponseBodyData {
	s.MeasureList = v
	return s
}

func (s *GetProblemImprovementResponseBodyData) SetMonitorSourceName(v string) *GetProblemImprovementResponseBodyData {
	s.MonitorSourceName = &v
	return s
}

func (s *GetProblemImprovementResponseBodyData) SetProblemId(v string) *GetProblemImprovementResponseBodyData {
	s.ProblemId = &v
	return s
}

func (s *GetProblemImprovementResponseBodyData) SetProblemReason(v string) *GetProblemImprovementResponseBodyData {
	s.ProblemReason = &v
	return s
}

func (s *GetProblemImprovementResponseBodyData) SetRecentActivity(v string) *GetProblemImprovementResponseBodyData {
	s.RecentActivity = &v
	return s
}

func (s *GetProblemImprovementResponseBodyData) SetRecoveryMode(v string) *GetProblemImprovementResponseBodyData {
	s.RecoveryMode = &v
	return s
}

func (s *GetProblemImprovementResponseBodyData) SetRelationChanges(v string) *GetProblemImprovementResponseBodyData {
	s.RelationChanges = &v
	return s
}

func (s *GetProblemImprovementResponseBodyData) SetRemark(v string) *GetProblemImprovementResponseBodyData {
	s.Remark = &v
	return s
}

func (s *GetProblemImprovementResponseBodyData) SetReplayDutyUserId(v int64) *GetProblemImprovementResponseBodyData {
	s.ReplayDutyUserId = &v
	return s
}

func (s *GetProblemImprovementResponseBodyData) SetReplayDutyUserIsValid(v int64) *GetProblemImprovementResponseBodyData {
	s.ReplayDutyUserIsValid = &v
	return s
}

func (s *GetProblemImprovementResponseBodyData) SetReplayDutyUserName(v string) *GetProblemImprovementResponseBodyData {
	s.ReplayDutyUserName = &v
	return s
}

func (s *GetProblemImprovementResponseBodyData) SetReplayDutyUserPhone(v string) *GetProblemImprovementResponseBodyData {
	s.ReplayDutyUserPhone = &v
	return s
}

func (s *GetProblemImprovementResponseBodyData) SetUserReport(v int64) *GetProblemImprovementResponseBodyData {
	s.UserReport = &v
	return s
}

type GetProblemImprovementResponseBodyDataMeasureList struct {
	CheckStandard    *string `json:"checkStandard,omitempty" xml:"checkStandard,omitempty"`
	CheckUserId      *int64  `json:"checkUserId,omitempty" xml:"checkUserId,omitempty"`
	CheckUserIsValid *int32  `json:"checkUserIsValid,omitempty" xml:"checkUserIsValid,omitempty"`
	CheckUserName    *string `json:"checkUserName,omitempty" xml:"checkUserName,omitempty"`
	Content          *string `json:"content,omitempty" xml:"content,omitempty"`
	DirectorId       *int64  `json:"directorId,omitempty" xml:"directorId,omitempty"`
	DirectorIsValid  *int32  `json:"directorIsValid,omitempty" xml:"directorIsValid,omitempty"`
	DirectorName     *string `json:"directorName,omitempty" xml:"directorName,omitempty"`
	MeasureId        *int64  `json:"measureId,omitempty" xml:"measureId,omitempty"`
	PlanFinishTime   *string `json:"planFinishTime,omitempty" xml:"planFinishTime,omitempty"`
	StalkerId        *int64  `json:"stalkerId,omitempty" xml:"stalkerId,omitempty"`
	StalkerIsValid   *int32  `json:"stalkerIsValid,omitempty" xml:"stalkerIsValid,omitempty"`
	StalkerName      *string `json:"stalkerName,omitempty" xml:"stalkerName,omitempty"`
	Status           *string `json:"status,omitempty" xml:"status,omitempty"`
	Type             *int64  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetProblemImprovementResponseBodyDataMeasureList) String() string {
	return tea.Prettify(s)
}

func (s GetProblemImprovementResponseBodyDataMeasureList) GoString() string {
	return s.String()
}

func (s *GetProblemImprovementResponseBodyDataMeasureList) SetCheckStandard(v string) *GetProblemImprovementResponseBodyDataMeasureList {
	s.CheckStandard = &v
	return s
}

func (s *GetProblemImprovementResponseBodyDataMeasureList) SetCheckUserId(v int64) *GetProblemImprovementResponseBodyDataMeasureList {
	s.CheckUserId = &v
	return s
}

func (s *GetProblemImprovementResponseBodyDataMeasureList) SetCheckUserIsValid(v int32) *GetProblemImprovementResponseBodyDataMeasureList {
	s.CheckUserIsValid = &v
	return s
}

func (s *GetProblemImprovementResponseBodyDataMeasureList) SetCheckUserName(v string) *GetProblemImprovementResponseBodyDataMeasureList {
	s.CheckUserName = &v
	return s
}

func (s *GetProblemImprovementResponseBodyDataMeasureList) SetContent(v string) *GetProblemImprovementResponseBodyDataMeasureList {
	s.Content = &v
	return s
}

func (s *GetProblemImprovementResponseBodyDataMeasureList) SetDirectorId(v int64) *GetProblemImprovementResponseBodyDataMeasureList {
	s.DirectorId = &v
	return s
}

func (s *GetProblemImprovementResponseBodyDataMeasureList) SetDirectorIsValid(v int32) *GetProblemImprovementResponseBodyDataMeasureList {
	s.DirectorIsValid = &v
	return s
}

func (s *GetProblemImprovementResponseBodyDataMeasureList) SetDirectorName(v string) *GetProblemImprovementResponseBodyDataMeasureList {
	s.DirectorName = &v
	return s
}

func (s *GetProblemImprovementResponseBodyDataMeasureList) SetMeasureId(v int64) *GetProblemImprovementResponseBodyDataMeasureList {
	s.MeasureId = &v
	return s
}

func (s *GetProblemImprovementResponseBodyDataMeasureList) SetPlanFinishTime(v string) *GetProblemImprovementResponseBodyDataMeasureList {
	s.PlanFinishTime = &v
	return s
}

func (s *GetProblemImprovementResponseBodyDataMeasureList) SetStalkerId(v int64) *GetProblemImprovementResponseBodyDataMeasureList {
	s.StalkerId = &v
	return s
}

func (s *GetProblemImprovementResponseBodyDataMeasureList) SetStalkerIsValid(v int32) *GetProblemImprovementResponseBodyDataMeasureList {
	s.StalkerIsValid = &v
	return s
}

func (s *GetProblemImprovementResponseBodyDataMeasureList) SetStalkerName(v string) *GetProblemImprovementResponseBodyDataMeasureList {
	s.StalkerName = &v
	return s
}

func (s *GetProblemImprovementResponseBodyDataMeasureList) SetStatus(v string) *GetProblemImprovementResponseBodyDataMeasureList {
	s.Status = &v
	return s
}

func (s *GetProblemImprovementResponseBodyDataMeasureList) SetType(v int64) *GetProblemImprovementResponseBodyDataMeasureList {
	s.Type = &v
	return s
}

type GetProblemImprovementResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetProblemImprovementResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProblemImprovementResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProblemImprovementResponse) GoString() string {
	return s.String()
}

func (s *GetProblemImprovementResponse) SetHeaders(v map[string]*string) *GetProblemImprovementResponse {
	s.Headers = v
	return s
}

func (s *GetProblemImprovementResponse) SetStatusCode(v int32) *GetProblemImprovementResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProblemImprovementResponse) SetBody(v *GetProblemImprovementResponseBody) *GetProblemImprovementResponse {
	s.Body = v
	return s
}

type GetProblemPreviewRequest struct {
	ClientToken       *string  `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	EffectServiceIds  []*int64 `json:"effectServiceIds,omitempty" xml:"effectServiceIds,omitempty" type:"Repeated"`
	IncidentId        *int64   `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
	ProblemId         *int64   `json:"problemId,omitempty" xml:"problemId,omitempty"`
	ProblemLevel      *string  `json:"problemLevel,omitempty" xml:"problemLevel,omitempty"`
	ProblemNotifyType *string  `json:"problemNotifyType,omitempty" xml:"problemNotifyType,omitempty"`
	RelatedServiceId  *int64   `json:"relatedServiceId,omitempty" xml:"relatedServiceId,omitempty"`
	ServiceGroupIds   []*int64 `json:"serviceGroupIds,omitempty" xml:"serviceGroupIds,omitempty" type:"Repeated"`
}

func (s GetProblemPreviewRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProblemPreviewRequest) GoString() string {
	return s.String()
}

func (s *GetProblemPreviewRequest) SetClientToken(v string) *GetProblemPreviewRequest {
	s.ClientToken = &v
	return s
}

func (s *GetProblemPreviewRequest) SetEffectServiceIds(v []*int64) *GetProblemPreviewRequest {
	s.EffectServiceIds = v
	return s
}

func (s *GetProblemPreviewRequest) SetIncidentId(v int64) *GetProblemPreviewRequest {
	s.IncidentId = &v
	return s
}

func (s *GetProblemPreviewRequest) SetProblemId(v int64) *GetProblemPreviewRequest {
	s.ProblemId = &v
	return s
}

func (s *GetProblemPreviewRequest) SetProblemLevel(v string) *GetProblemPreviewRequest {
	s.ProblemLevel = &v
	return s
}

func (s *GetProblemPreviewRequest) SetProblemNotifyType(v string) *GetProblemPreviewRequest {
	s.ProblemNotifyType = &v
	return s
}

func (s *GetProblemPreviewRequest) SetRelatedServiceId(v int64) *GetProblemPreviewRequest {
	s.RelatedServiceId = &v
	return s
}

func (s *GetProblemPreviewRequest) SetServiceGroupIds(v []*int64) *GetProblemPreviewRequest {
	s.ServiceGroupIds = v
	return s
}

type GetProblemPreviewResponseBody struct {
	// data
	Data *GetProblemPreviewResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetProblemPreviewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProblemPreviewResponseBody) GoString() string {
	return s.String()
}

func (s *GetProblemPreviewResponseBody) SetData(v *GetProblemPreviewResponseBodyData) *GetProblemPreviewResponseBody {
	s.Data = v
	return s
}

func (s *GetProblemPreviewResponseBody) SetRequestId(v string) *GetProblemPreviewResponseBody {
	s.RequestId = &v
	return s
}

type GetProblemPreviewResponseBodyData struct {
	DeAfterData  *string                                `json:"deAfterData,omitempty" xml:"deAfterData,omitempty"`
	DeBeforeData *string                                `json:"deBeforeData,omitempty" xml:"deBeforeData,omitempty"`
	Mail         *GetProblemPreviewResponseBodyDataMail `json:"mail,omitempty" xml:"mail,omitempty" type:"Struct"`
	// object
	Problem      *GetProblemPreviewResponseBodyDataProblem `json:"problem,omitempty" xml:"problem,omitempty" type:"Struct"`
	Sms          *GetProblemPreviewResponseBodyDataSms     `json:"sms,omitempty" xml:"sms,omitempty" type:"Struct"`
	UpAfterData  *string                                   `json:"upAfterData,omitempty" xml:"upAfterData,omitempty"`
	UpBeforeData *string                                   `json:"upBeforeData,omitempty" xml:"upBeforeData,omitempty"`
	Voice        *GetProblemPreviewResponseBodyDataVoice   `json:"voice,omitempty" xml:"voice,omitempty" type:"Struct"`
	// webhook
	Webhook *GetProblemPreviewResponseBodyDataWebhook `json:"webhook,omitempty" xml:"webhook,omitempty" type:"Struct"`
}

func (s GetProblemPreviewResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetProblemPreviewResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProblemPreviewResponseBodyData) SetDeAfterData(v string) *GetProblemPreviewResponseBodyData {
	s.DeAfterData = &v
	return s
}

func (s *GetProblemPreviewResponseBodyData) SetDeBeforeData(v string) *GetProblemPreviewResponseBodyData {
	s.DeBeforeData = &v
	return s
}

func (s *GetProblemPreviewResponseBodyData) SetMail(v *GetProblemPreviewResponseBodyDataMail) *GetProblemPreviewResponseBodyData {
	s.Mail = v
	return s
}

func (s *GetProblemPreviewResponseBodyData) SetProblem(v *GetProblemPreviewResponseBodyDataProblem) *GetProblemPreviewResponseBodyData {
	s.Problem = v
	return s
}

func (s *GetProblemPreviewResponseBodyData) SetSms(v *GetProblemPreviewResponseBodyDataSms) *GetProblemPreviewResponseBodyData {
	s.Sms = v
	return s
}

func (s *GetProblemPreviewResponseBodyData) SetUpAfterData(v string) *GetProblemPreviewResponseBodyData {
	s.UpAfterData = &v
	return s
}

func (s *GetProblemPreviewResponseBodyData) SetUpBeforeData(v string) *GetProblemPreviewResponseBodyData {
	s.UpBeforeData = &v
	return s
}

func (s *GetProblemPreviewResponseBodyData) SetVoice(v *GetProblemPreviewResponseBodyDataVoice) *GetProblemPreviewResponseBodyData {
	s.Voice = v
	return s
}

func (s *GetProblemPreviewResponseBodyData) SetWebhook(v *GetProblemPreviewResponseBodyDataWebhook) *GetProblemPreviewResponseBodyData {
	s.Webhook = v
	return s
}

type GetProblemPreviewResponseBodyDataMail struct {
	Count *int64                                        `json:"count,omitempty" xml:"count,omitempty"`
	Users []*GetProblemPreviewResponseBodyDataMailUsers `json:"users,omitempty" xml:"users,omitempty" type:"Repeated"`
}

func (s GetProblemPreviewResponseBodyDataMail) String() string {
	return tea.Prettify(s)
}

func (s GetProblemPreviewResponseBodyDataMail) GoString() string {
	return s.String()
}

func (s *GetProblemPreviewResponseBodyDataMail) SetCount(v int64) *GetProblemPreviewResponseBodyDataMail {
	s.Count = &v
	return s
}

func (s *GetProblemPreviewResponseBodyDataMail) SetUsers(v []*GetProblemPreviewResponseBodyDataMailUsers) *GetProblemPreviewResponseBodyDataMail {
	s.Users = v
	return s
}

type GetProblemPreviewResponseBodyDataMailUsers struct {
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s GetProblemPreviewResponseBodyDataMailUsers) String() string {
	return tea.Prettify(s)
}

func (s GetProblemPreviewResponseBodyDataMailUsers) GoString() string {
	return s.String()
}

func (s *GetProblemPreviewResponseBodyDataMailUsers) SetUsername(v string) *GetProblemPreviewResponseBodyDataMailUsers {
	s.Username = &v
	return s
}

type GetProblemPreviewResponseBodyDataProblem struct {
	CoordinationGroups        []*GetProblemPreviewResponseBodyDataProblemCoordinationGroups `json:"coordinationGroups,omitempty" xml:"coordinationGroups,omitempty" type:"Repeated"`
	CreateTime                *string                                                       `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DiscoverTime              *string                                                       `json:"discoverTime,omitempty" xml:"discoverTime,omitempty"`
	EffectionServices         []*GetProblemPreviewResponseBodyDataProblemEffectionServices  `json:"effectionServices,omitempty" xml:"effectionServices,omitempty" type:"Repeated"`
	IsManual                  *bool                                                         `json:"isManual,omitempty" xml:"isManual,omitempty"`
	IsUpgrade                 *bool                                                         `json:"isUpgrade,omitempty" xml:"isUpgrade,omitempty"`
	MainHandlerId             *string                                                       `json:"mainHandlerId,omitempty" xml:"mainHandlerId,omitempty"`
	MainHandlerName           *string                                                       `json:"mainHandlerName,omitempty" xml:"mainHandlerName,omitempty"`
	PreliminaryReason         *string                                                       `json:"preliminaryReason,omitempty" xml:"preliminaryReason,omitempty"`
	ProblemId                 *int64                                                        `json:"problemId,omitempty" xml:"problemId,omitempty"`
	ProblemLevel              *string                                                       `json:"problemLevel,omitempty" xml:"problemLevel,omitempty"`
	ProblemName               *string                                                       `json:"problemName,omitempty" xml:"problemName,omitempty"`
	ProblemStatus             *string                                                       `json:"problemStatus,omitempty" xml:"problemStatus,omitempty"`
	ProgressSummary           *string                                                       `json:"progressSummary,omitempty" xml:"progressSummary,omitempty"`
	ProgressSummaryRichTextId *int64                                                        `json:"progressSummaryRichTextId,omitempty" xml:"progressSummaryRichTextId,omitempty"`
	RecoveryTime              *string                                                       `json:"recoveryTime,omitempty" xml:"recoveryTime,omitempty"`
	RelatedServiceId          *int64                                                        `json:"relatedServiceId,omitempty" xml:"relatedServiceId,omitempty"`
	ServiceName               *string                                                       `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
}

func (s GetProblemPreviewResponseBodyDataProblem) String() string {
	return tea.Prettify(s)
}

func (s GetProblemPreviewResponseBodyDataProblem) GoString() string {
	return s.String()
}

func (s *GetProblemPreviewResponseBodyDataProblem) SetCoordinationGroups(v []*GetProblemPreviewResponseBodyDataProblemCoordinationGroups) *GetProblemPreviewResponseBodyDataProblem {
	s.CoordinationGroups = v
	return s
}

func (s *GetProblemPreviewResponseBodyDataProblem) SetCreateTime(v string) *GetProblemPreviewResponseBodyDataProblem {
	s.CreateTime = &v
	return s
}

func (s *GetProblemPreviewResponseBodyDataProblem) SetDiscoverTime(v string) *GetProblemPreviewResponseBodyDataProblem {
	s.DiscoverTime = &v
	return s
}

func (s *GetProblemPreviewResponseBodyDataProblem) SetEffectionServices(v []*GetProblemPreviewResponseBodyDataProblemEffectionServices) *GetProblemPreviewResponseBodyDataProblem {
	s.EffectionServices = v
	return s
}

func (s *GetProblemPreviewResponseBodyDataProblem) SetIsManual(v bool) *GetProblemPreviewResponseBodyDataProblem {
	s.IsManual = &v
	return s
}

func (s *GetProblemPreviewResponseBodyDataProblem) SetIsUpgrade(v bool) *GetProblemPreviewResponseBodyDataProblem {
	s.IsUpgrade = &v
	return s
}

func (s *GetProblemPreviewResponseBodyDataProblem) SetMainHandlerId(v string) *GetProblemPreviewResponseBodyDataProblem {
	s.MainHandlerId = &v
	return s
}

func (s *GetProblemPreviewResponseBodyDataProblem) SetMainHandlerName(v string) *GetProblemPreviewResponseBodyDataProblem {
	s.MainHandlerName = &v
	return s
}

func (s *GetProblemPreviewResponseBodyDataProblem) SetPreliminaryReason(v string) *GetProblemPreviewResponseBodyDataProblem {
	s.PreliminaryReason = &v
	return s
}

func (s *GetProblemPreviewResponseBodyDataProblem) SetProblemId(v int64) *GetProblemPreviewResponseBodyDataProblem {
	s.ProblemId = &v
	return s
}

func (s *GetProblemPreviewResponseBodyDataProblem) SetProblemLevel(v string) *GetProblemPreviewResponseBodyDataProblem {
	s.ProblemLevel = &v
	return s
}

func (s *GetProblemPreviewResponseBodyDataProblem) SetProblemName(v string) *GetProblemPreviewResponseBodyDataProblem {
	s.ProblemName = &v
	return s
}

func (s *GetProblemPreviewResponseBodyDataProblem) SetProblemStatus(v string) *GetProblemPreviewResponseBodyDataProblem {
	s.ProblemStatus = &v
	return s
}

func (s *GetProblemPreviewResponseBodyDataProblem) SetProgressSummary(v string) *GetProblemPreviewResponseBodyDataProblem {
	s.ProgressSummary = &v
	return s
}

func (s *GetProblemPreviewResponseBodyDataProblem) SetProgressSummaryRichTextId(v int64) *GetProblemPreviewResponseBodyDataProblem {
	s.ProgressSummaryRichTextId = &v
	return s
}

func (s *GetProblemPreviewResponseBodyDataProblem) SetRecoveryTime(v string) *GetProblemPreviewResponseBodyDataProblem {
	s.RecoveryTime = &v
	return s
}

func (s *GetProblemPreviewResponseBodyDataProblem) SetRelatedServiceId(v int64) *GetProblemPreviewResponseBodyDataProblem {
	s.RelatedServiceId = &v
	return s
}

func (s *GetProblemPreviewResponseBodyDataProblem) SetServiceName(v string) *GetProblemPreviewResponseBodyDataProblem {
	s.ServiceName = &v
	return s
}

type GetProblemPreviewResponseBodyDataProblemCoordinationGroups struct {
	ServiceGroupDescription *string `json:"serviceGroupDescription,omitempty" xml:"serviceGroupDescription,omitempty"`
	ServiceGroupId          *int64  `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
	ServiceGroupName        *string `json:"serviceGroupName,omitempty" xml:"serviceGroupName,omitempty"`
}

func (s GetProblemPreviewResponseBodyDataProblemCoordinationGroups) String() string {
	return tea.Prettify(s)
}

func (s GetProblemPreviewResponseBodyDataProblemCoordinationGroups) GoString() string {
	return s.String()
}

func (s *GetProblemPreviewResponseBodyDataProblemCoordinationGroups) SetServiceGroupDescription(v string) *GetProblemPreviewResponseBodyDataProblemCoordinationGroups {
	s.ServiceGroupDescription = &v
	return s
}

func (s *GetProblemPreviewResponseBodyDataProblemCoordinationGroups) SetServiceGroupId(v int64) *GetProblemPreviewResponseBodyDataProblemCoordinationGroups {
	s.ServiceGroupId = &v
	return s
}

func (s *GetProblemPreviewResponseBodyDataProblemCoordinationGroups) SetServiceGroupName(v string) *GetProblemPreviewResponseBodyDataProblemCoordinationGroups {
	s.ServiceGroupName = &v
	return s
}

type GetProblemPreviewResponseBodyDataProblemEffectionServices struct {
	ServiceId   *int64  `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
}

func (s GetProblemPreviewResponseBodyDataProblemEffectionServices) String() string {
	return tea.Prettify(s)
}

func (s GetProblemPreviewResponseBodyDataProblemEffectionServices) GoString() string {
	return s.String()
}

func (s *GetProblemPreviewResponseBodyDataProblemEffectionServices) SetServiceId(v int64) *GetProblemPreviewResponseBodyDataProblemEffectionServices {
	s.ServiceId = &v
	return s
}

func (s *GetProblemPreviewResponseBodyDataProblemEffectionServices) SetServiceName(v string) *GetProblemPreviewResponseBodyDataProblemEffectionServices {
	s.ServiceName = &v
	return s
}

type GetProblemPreviewResponseBodyDataSms struct {
	Count *int64                                       `json:"count,omitempty" xml:"count,omitempty"`
	Users []*GetProblemPreviewResponseBodyDataSmsUsers `json:"users,omitempty" xml:"users,omitempty" type:"Repeated"`
}

func (s GetProblemPreviewResponseBodyDataSms) String() string {
	return tea.Prettify(s)
}

func (s GetProblemPreviewResponseBodyDataSms) GoString() string {
	return s.String()
}

func (s *GetProblemPreviewResponseBodyDataSms) SetCount(v int64) *GetProblemPreviewResponseBodyDataSms {
	s.Count = &v
	return s
}

func (s *GetProblemPreviewResponseBodyDataSms) SetUsers(v []*GetProblemPreviewResponseBodyDataSmsUsers) *GetProblemPreviewResponseBodyDataSms {
	s.Users = v
	return s
}

type GetProblemPreviewResponseBodyDataSmsUsers struct {
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s GetProblemPreviewResponseBodyDataSmsUsers) String() string {
	return tea.Prettify(s)
}

func (s GetProblemPreviewResponseBodyDataSmsUsers) GoString() string {
	return s.String()
}

func (s *GetProblemPreviewResponseBodyDataSmsUsers) SetUsername(v string) *GetProblemPreviewResponseBodyDataSmsUsers {
	s.Username = &v
	return s
}

type GetProblemPreviewResponseBodyDataVoice struct {
	Count *int64                                         `json:"count,omitempty" xml:"count,omitempty"`
	Users []*GetProblemPreviewResponseBodyDataVoiceUsers `json:"users,omitempty" xml:"users,omitempty" type:"Repeated"`
}

func (s GetProblemPreviewResponseBodyDataVoice) String() string {
	return tea.Prettify(s)
}

func (s GetProblemPreviewResponseBodyDataVoice) GoString() string {
	return s.String()
}

func (s *GetProblemPreviewResponseBodyDataVoice) SetCount(v int64) *GetProblemPreviewResponseBodyDataVoice {
	s.Count = &v
	return s
}

func (s *GetProblemPreviewResponseBodyDataVoice) SetUsers(v []*GetProblemPreviewResponseBodyDataVoiceUsers) *GetProblemPreviewResponseBodyDataVoice {
	s.Users = v
	return s
}

type GetProblemPreviewResponseBodyDataVoiceUsers struct {
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s GetProblemPreviewResponseBodyDataVoiceUsers) String() string {
	return tea.Prettify(s)
}

func (s GetProblemPreviewResponseBodyDataVoiceUsers) GoString() string {
	return s.String()
}

func (s *GetProblemPreviewResponseBodyDataVoiceUsers) SetUsername(v string) *GetProblemPreviewResponseBodyDataVoiceUsers {
	s.Username = &v
	return s
}

type GetProblemPreviewResponseBodyDataWebhook struct {
	Count         *int64                                                   `json:"count,omitempty" xml:"count,omitempty"`
	ServiceGroups []*GetProblemPreviewResponseBodyDataWebhookServiceGroups `json:"serviceGroups,omitempty" xml:"serviceGroups,omitempty" type:"Repeated"`
}

func (s GetProblemPreviewResponseBodyDataWebhook) String() string {
	return tea.Prettify(s)
}

func (s GetProblemPreviewResponseBodyDataWebhook) GoString() string {
	return s.String()
}

func (s *GetProblemPreviewResponseBodyDataWebhook) SetCount(v int64) *GetProblemPreviewResponseBodyDataWebhook {
	s.Count = &v
	return s
}

func (s *GetProblemPreviewResponseBodyDataWebhook) SetServiceGroups(v []*GetProblemPreviewResponseBodyDataWebhookServiceGroups) *GetProblemPreviewResponseBodyDataWebhook {
	s.ServiceGroups = v
	return s
}

type GetProblemPreviewResponseBodyDataWebhookServiceGroups struct {
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
}

func (s GetProblemPreviewResponseBodyDataWebhookServiceGroups) String() string {
	return tea.Prettify(s)
}

func (s GetProblemPreviewResponseBodyDataWebhookServiceGroups) GoString() string {
	return s.String()
}

func (s *GetProblemPreviewResponseBodyDataWebhookServiceGroups) SetServiceName(v string) *GetProblemPreviewResponseBodyDataWebhookServiceGroups {
	s.ServiceName = &v
	return s
}

type GetProblemPreviewResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetProblemPreviewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProblemPreviewResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProblemPreviewResponse) GoString() string {
	return s.String()
}

func (s *GetProblemPreviewResponse) SetHeaders(v map[string]*string) *GetProblemPreviewResponse {
	s.Headers = v
	return s
}

func (s *GetProblemPreviewResponse) SetStatusCode(v int32) *GetProblemPreviewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProblemPreviewResponse) SetBody(v *GetProblemPreviewResponseBody) *GetProblemPreviewResponse {
	s.Body = v
	return s
}

type GetResourceStatisticsRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s GetResourceStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetResourceStatisticsRequest) SetClientToken(v string) *GetResourceStatisticsRequest {
	s.ClientToken = &v
	return s
}

type GetResourceStatisticsResponseBody struct {
	// data
	Data *GetResourceStatisticsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetResourceStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceStatisticsResponseBody) SetData(v *GetResourceStatisticsResponseBodyData) *GetResourceStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *GetResourceStatisticsResponseBody) SetRequestId(v string) *GetResourceStatisticsResponseBody {
	s.RequestId = &v
	return s
}

type GetResourceStatisticsResponseBodyData struct {
	AlertCount       *int32 `json:"alertCount,omitempty" xml:"alertCount,omitempty"`
	IncidentCount    *int32 `json:"incidentCount,omitempty" xml:"incidentCount,omitempty"`
	IntegrationCount *int32 `json:"integrationCount,omitempty" xml:"integrationCount,omitempty"`
	ProblemCount     *int32 `json:"problemCount,omitempty" xml:"problemCount,omitempty"`
}

func (s GetResourceStatisticsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetResourceStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetResourceStatisticsResponseBodyData) SetAlertCount(v int32) *GetResourceStatisticsResponseBodyData {
	s.AlertCount = &v
	return s
}

func (s *GetResourceStatisticsResponseBodyData) SetIncidentCount(v int32) *GetResourceStatisticsResponseBodyData {
	s.IncidentCount = &v
	return s
}

func (s *GetResourceStatisticsResponseBodyData) SetIntegrationCount(v int32) *GetResourceStatisticsResponseBodyData {
	s.IntegrationCount = &v
	return s
}

func (s *GetResourceStatisticsResponseBodyData) SetProblemCount(v int32) *GetResourceStatisticsResponseBodyData {
	s.ProblemCount = &v
	return s
}

type GetResourceStatisticsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetResourceStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetResourceStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetResourceStatisticsResponse) SetHeaders(v map[string]*string) *GetResourceStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetResourceStatisticsResponse) SetStatusCode(v int32) *GetResourceStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceStatisticsResponse) SetBody(v *GetResourceStatisticsResponseBody) *GetResourceStatisticsResponse {
	s.Body = v
	return s
}

type GetRichTextRequest struct {
	InstanceId   *int64  `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	RichTextId   *int64  `json:"richTextId,omitempty" xml:"richTextId,omitempty"`
}

func (s GetRichTextRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRichTextRequest) GoString() string {
	return s.String()
}

func (s *GetRichTextRequest) SetInstanceId(v int64) *GetRichTextRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRichTextRequest) SetInstanceType(v string) *GetRichTextRequest {
	s.InstanceType = &v
	return s
}

func (s *GetRichTextRequest) SetRichTextId(v int64) *GetRichTextRequest {
	s.RichTextId = &v
	return s
}

type GetRichTextResponseBody struct {
	// data
	Data *GetRichTextResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetRichTextResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRichTextResponseBody) GoString() string {
	return s.String()
}

func (s *GetRichTextResponseBody) SetData(v *GetRichTextResponseBodyData) *GetRichTextResponseBody {
	s.Data = v
	return s
}

func (s *GetRichTextResponseBody) SetRequestId(v string) *GetRichTextResponseBody {
	s.RequestId = &v
	return s
}

type GetRichTextResponseBodyData struct {
	InstanceId   *int64  `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	InstanceType *int64  `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	RichText     *string `json:"richText,omitempty" xml:"richText,omitempty"`
}

func (s GetRichTextResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRichTextResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRichTextResponseBodyData) SetInstanceId(v int64) *GetRichTextResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetRichTextResponseBodyData) SetInstanceType(v int64) *GetRichTextResponseBodyData {
	s.InstanceType = &v
	return s
}

func (s *GetRichTextResponseBodyData) SetRichText(v string) *GetRichTextResponseBodyData {
	s.RichText = &v
	return s
}

type GetRichTextResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRichTextResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRichTextResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRichTextResponse) GoString() string {
	return s.String()
}

func (s *GetRichTextResponse) SetHeaders(v map[string]*string) *GetRichTextResponse {
	s.Headers = v
	return s
}

func (s *GetRichTextResponse) SetStatusCode(v int32) *GetRichTextResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRichTextResponse) SetBody(v *GetRichTextResponseBody) *GetRichTextResponse {
	s.Body = v
	return s
}

type GetRouteRuleRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	RouteRuleId *int64  `json:"routeRuleId,omitempty" xml:"routeRuleId,omitempty"`
}

func (s GetRouteRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRouteRuleRequest) GoString() string {
	return s.String()
}

func (s *GetRouteRuleRequest) SetClientToken(v string) *GetRouteRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *GetRouteRuleRequest) SetRouteRuleId(v int64) *GetRouteRuleRequest {
	s.RouteRuleId = &v
	return s
}

type GetRouteRuleResponseBody struct {
	Data      *GetRouteRuleResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	RequestId *string                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetRouteRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRouteRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetRouteRuleResponseBody) SetData(v *GetRouteRuleResponseBodyData) *GetRouteRuleResponseBody {
	s.Data = v
	return s
}

func (s *GetRouteRuleResponseBody) SetRequestId(v string) *GetRouteRuleResponseBody {
	s.RequestId = &v
	return s
}

type GetRouteRuleResponseBodyData struct {
	AssignObjectId           *int64                                              `json:"assignObjectId,omitempty" xml:"assignObjectId,omitempty"`
	AssignObjectName         *string                                             `json:"assignObjectName,omitempty" xml:"assignObjectName,omitempty"`
	AssignObjectType         *string                                             `json:"assignObjectType,omitempty" xml:"assignObjectType,omitempty"`
	ChildRuleRelation        *string                                             `json:"childRuleRelation,omitempty" xml:"childRuleRelation,omitempty"`
	CoverageProblemLevels    []*string                                           `json:"coverageProblemLevels,omitempty" xml:"coverageProblemLevels,omitempty" type:"Repeated"`
	CreateTime               *string                                             `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Effection                *string                                             `json:"effection,omitempty" xml:"effection,omitempty"`
	EnableStatus             *string                                             `json:"enableStatus,omitempty" xml:"enableStatus,omitempty"`
	EventRouteChildRules     []*GetRouteRuleResponseBodyDataEventRouteChildRules `json:"eventRouteChildRules,omitempty" xml:"eventRouteChildRules,omitempty" type:"Repeated"`
	IncidentLevel            *string                                             `json:"incidentLevel,omitempty" xml:"incidentLevel,omitempty"`
	MatchCount               *int64                                              `json:"matchCount,omitempty" xml:"matchCount,omitempty"`
	NotifyChannelNames       []*string                                           `json:"notifyChannelNames,omitempty" xml:"notifyChannelNames,omitempty" type:"Repeated"`
	NotifyChannels           []*string                                           `json:"notifyChannels,omitempty" xml:"notifyChannels,omitempty" type:"Repeated"`
	ProblemEffectionServices []*int64                                            `json:"problemEffectionServices,omitempty" xml:"problemEffectionServices,omitempty" type:"Repeated"`
	ProblemLevelGroup        map[string]*DataProblemLevelGroupValue              `json:"problemLevelGroup,omitempty" xml:"problemLevelGroup,omitempty"`
	RelServiceDeleteType     *int32                                              `json:"relServiceDeleteType,omitempty" xml:"relServiceDeleteType,omitempty"`
	RelatedServiceId         *int64                                              `json:"relatedServiceId,omitempty" xml:"relatedServiceId,omitempty"`
	RelatedServiceName       *string                                             `json:"relatedServiceName,omitempty" xml:"relatedServiceName,omitempty"`
	RouteRuleId              *int64                                              `json:"routeRuleId,omitempty" xml:"routeRuleId,omitempty"`
	RouteType                *string                                             `json:"routeType,omitempty" xml:"routeType,omitempty"`
	RuleName                 *string                                             `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	TimeWindow               *int32                                              `json:"timeWindow,omitempty" xml:"timeWindow,omitempty"`
	UpdateTime               *string                                             `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetRouteRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRouteRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRouteRuleResponseBodyData) SetAssignObjectId(v int64) *GetRouteRuleResponseBodyData {
	s.AssignObjectId = &v
	return s
}

func (s *GetRouteRuleResponseBodyData) SetAssignObjectName(v string) *GetRouteRuleResponseBodyData {
	s.AssignObjectName = &v
	return s
}

func (s *GetRouteRuleResponseBodyData) SetAssignObjectType(v string) *GetRouteRuleResponseBodyData {
	s.AssignObjectType = &v
	return s
}

func (s *GetRouteRuleResponseBodyData) SetChildRuleRelation(v string) *GetRouteRuleResponseBodyData {
	s.ChildRuleRelation = &v
	return s
}

func (s *GetRouteRuleResponseBodyData) SetCoverageProblemLevels(v []*string) *GetRouteRuleResponseBodyData {
	s.CoverageProblemLevels = v
	return s
}

func (s *GetRouteRuleResponseBodyData) SetCreateTime(v string) *GetRouteRuleResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetRouteRuleResponseBodyData) SetEffection(v string) *GetRouteRuleResponseBodyData {
	s.Effection = &v
	return s
}

func (s *GetRouteRuleResponseBodyData) SetEnableStatus(v string) *GetRouteRuleResponseBodyData {
	s.EnableStatus = &v
	return s
}

func (s *GetRouteRuleResponseBodyData) SetEventRouteChildRules(v []*GetRouteRuleResponseBodyDataEventRouteChildRules) *GetRouteRuleResponseBodyData {
	s.EventRouteChildRules = v
	return s
}

func (s *GetRouteRuleResponseBodyData) SetIncidentLevel(v string) *GetRouteRuleResponseBodyData {
	s.IncidentLevel = &v
	return s
}

func (s *GetRouteRuleResponseBodyData) SetMatchCount(v int64) *GetRouteRuleResponseBodyData {
	s.MatchCount = &v
	return s
}

func (s *GetRouteRuleResponseBodyData) SetNotifyChannelNames(v []*string) *GetRouteRuleResponseBodyData {
	s.NotifyChannelNames = v
	return s
}

func (s *GetRouteRuleResponseBodyData) SetNotifyChannels(v []*string) *GetRouteRuleResponseBodyData {
	s.NotifyChannels = v
	return s
}

func (s *GetRouteRuleResponseBodyData) SetProblemEffectionServices(v []*int64) *GetRouteRuleResponseBodyData {
	s.ProblemEffectionServices = v
	return s
}

func (s *GetRouteRuleResponseBodyData) SetProblemLevelGroup(v map[string]*DataProblemLevelGroupValue) *GetRouteRuleResponseBodyData {
	s.ProblemLevelGroup = v
	return s
}

func (s *GetRouteRuleResponseBodyData) SetRelServiceDeleteType(v int32) *GetRouteRuleResponseBodyData {
	s.RelServiceDeleteType = &v
	return s
}

func (s *GetRouteRuleResponseBodyData) SetRelatedServiceId(v int64) *GetRouteRuleResponseBodyData {
	s.RelatedServiceId = &v
	return s
}

func (s *GetRouteRuleResponseBodyData) SetRelatedServiceName(v string) *GetRouteRuleResponseBodyData {
	s.RelatedServiceName = &v
	return s
}

func (s *GetRouteRuleResponseBodyData) SetRouteRuleId(v int64) *GetRouteRuleResponseBodyData {
	s.RouteRuleId = &v
	return s
}

func (s *GetRouteRuleResponseBodyData) SetRouteType(v string) *GetRouteRuleResponseBodyData {
	s.RouteType = &v
	return s
}

func (s *GetRouteRuleResponseBodyData) SetRuleName(v string) *GetRouteRuleResponseBodyData {
	s.RuleName = &v
	return s
}

func (s *GetRouteRuleResponseBodyData) SetTimeWindow(v int32) *GetRouteRuleResponseBodyData {
	s.TimeWindow = &v
	return s
}

func (s *GetRouteRuleResponseBodyData) SetUpdateTime(v string) *GetRouteRuleResponseBodyData {
	s.UpdateTime = &v
	return s
}

type GetRouteRuleResponseBodyDataEventRouteChildRules struct {
	ChildConditionRelation     *int64                                                        `json:"childConditionRelation,omitempty" xml:"childConditionRelation,omitempty"`
	ChildRouteRuleId           *int64                                                        `json:"childRouteRuleId,omitempty" xml:"childRouteRuleId,omitempty"`
	Conditions                 []*GetRouteRuleResponseBodyDataEventRouteChildRulesConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	IsValidChildRule           *bool                                                         `json:"isValidChildRule,omitempty" xml:"isValidChildRule,omitempty"`
	MonitorIntegrationConfigId *int64                                                        `json:"monitorIntegrationConfigId,omitempty" xml:"monitorIntegrationConfigId,omitempty"`
	MonitorSourceId            *int64                                                        `json:"monitorSourceId,omitempty" xml:"monitorSourceId,omitempty"`
	MonitorSourceName          *string                                                       `json:"monitorSourceName,omitempty" xml:"monitorSourceName,omitempty"`
	ParentRuleId               *int64                                                        `json:"parentRuleId,omitempty" xml:"parentRuleId,omitempty"`
	ProblemLevel               *string                                                       `json:"problemLevel,omitempty" xml:"problemLevel,omitempty"`
}

func (s GetRouteRuleResponseBodyDataEventRouteChildRules) String() string {
	return tea.Prettify(s)
}

func (s GetRouteRuleResponseBodyDataEventRouteChildRules) GoString() string {
	return s.String()
}

func (s *GetRouteRuleResponseBodyDataEventRouteChildRules) SetChildConditionRelation(v int64) *GetRouteRuleResponseBodyDataEventRouteChildRules {
	s.ChildConditionRelation = &v
	return s
}

func (s *GetRouteRuleResponseBodyDataEventRouteChildRules) SetChildRouteRuleId(v int64) *GetRouteRuleResponseBodyDataEventRouteChildRules {
	s.ChildRouteRuleId = &v
	return s
}

func (s *GetRouteRuleResponseBodyDataEventRouteChildRules) SetConditions(v []*GetRouteRuleResponseBodyDataEventRouteChildRulesConditions) *GetRouteRuleResponseBodyDataEventRouteChildRules {
	s.Conditions = v
	return s
}

func (s *GetRouteRuleResponseBodyDataEventRouteChildRules) SetIsValidChildRule(v bool) *GetRouteRuleResponseBodyDataEventRouteChildRules {
	s.IsValidChildRule = &v
	return s
}

func (s *GetRouteRuleResponseBodyDataEventRouteChildRules) SetMonitorIntegrationConfigId(v int64) *GetRouteRuleResponseBodyDataEventRouteChildRules {
	s.MonitorIntegrationConfigId = &v
	return s
}

func (s *GetRouteRuleResponseBodyDataEventRouteChildRules) SetMonitorSourceId(v int64) *GetRouteRuleResponseBodyDataEventRouteChildRules {
	s.MonitorSourceId = &v
	return s
}

func (s *GetRouteRuleResponseBodyDataEventRouteChildRules) SetMonitorSourceName(v string) *GetRouteRuleResponseBodyDataEventRouteChildRules {
	s.MonitorSourceName = &v
	return s
}

func (s *GetRouteRuleResponseBodyDataEventRouteChildRules) SetParentRuleId(v int64) *GetRouteRuleResponseBodyDataEventRouteChildRules {
	s.ParentRuleId = &v
	return s
}

func (s *GetRouteRuleResponseBodyDataEventRouteChildRules) SetProblemLevel(v string) *GetRouteRuleResponseBodyDataEventRouteChildRules {
	s.ProblemLevel = &v
	return s
}

type GetRouteRuleResponseBodyDataEventRouteChildRulesConditions struct {
	Key             *string `json:"key,omitempty" xml:"key,omitempty"`
	OperationSymbol *string `json:"operationSymbol,omitempty" xml:"operationSymbol,omitempty"`
	Value           *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetRouteRuleResponseBodyDataEventRouteChildRulesConditions) String() string {
	return tea.Prettify(s)
}

func (s GetRouteRuleResponseBodyDataEventRouteChildRulesConditions) GoString() string {
	return s.String()
}

func (s *GetRouteRuleResponseBodyDataEventRouteChildRulesConditions) SetKey(v string) *GetRouteRuleResponseBodyDataEventRouteChildRulesConditions {
	s.Key = &v
	return s
}

func (s *GetRouteRuleResponseBodyDataEventRouteChildRulesConditions) SetOperationSymbol(v string) *GetRouteRuleResponseBodyDataEventRouteChildRulesConditions {
	s.OperationSymbol = &v
	return s
}

func (s *GetRouteRuleResponseBodyDataEventRouteChildRulesConditions) SetValue(v string) *GetRouteRuleResponseBodyDataEventRouteChildRulesConditions {
	s.Value = &v
	return s
}

type GetRouteRuleResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRouteRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRouteRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRouteRuleResponse) GoString() string {
	return s.String()
}

func (s *GetRouteRuleResponse) SetHeaders(v map[string]*string) *GetRouteRuleResponse {
	s.Headers = v
	return s
}

func (s *GetRouteRuleResponse) SetStatusCode(v int32) *GetRouteRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRouteRuleResponse) SetBody(v *GetRouteRuleResponseBody) *GetRouteRuleResponse {
	s.Body = v
	return s
}

type GetServiceRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	ServiceId   *int64  `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
}

func (s GetServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceRequest) GoString() string {
	return s.String()
}

func (s *GetServiceRequest) SetClientToken(v string) *GetServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *GetServiceRequest) SetServiceId(v int64) *GetServiceRequest {
	s.ServiceId = &v
	return s
}

type GetServiceResponseBody struct {
	Data *GetServiceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBody) SetData(v *GetServiceResponseBodyData) *GetServiceResponseBody {
	s.Data = v
	return s
}

func (s *GetServiceResponseBody) SetRequestId(v string) *GetServiceResponseBody {
	s.RequestId = &v
	return s
}

type GetServiceResponseBodyData struct {
	ServiceDescription *string `json:"serviceDescription,omitempty" xml:"serviceDescription,omitempty"`
	ServiceId          *int64  `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	ServiceName        *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	UpdateTime         *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetServiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyData) SetServiceDescription(v string) *GetServiceResponseBodyData {
	s.ServiceDescription = &v
	return s
}

func (s *GetServiceResponseBodyData) SetServiceId(v int64) *GetServiceResponseBodyData {
	s.ServiceId = &v
	return s
}

func (s *GetServiceResponseBodyData) SetServiceName(v string) *GetServiceResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *GetServiceResponseBodyData) SetUpdateTime(v string) *GetServiceResponseBodyData {
	s.UpdateTime = &v
	return s
}

type GetServiceResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponse) GoString() string {
	return s.String()
}

func (s *GetServiceResponse) SetHeaders(v map[string]*string) *GetServiceResponse {
	s.Headers = v
	return s
}

func (s *GetServiceResponse) SetStatusCode(v int32) *GetServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceResponse) SetBody(v *GetServiceResponseBody) *GetServiceResponse {
	s.Body = v
	return s
}

type GetServiceGroupRequest struct {
	ClientToken    *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	ServiceGroupId *int64  `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
}

func (s GetServiceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceGroupRequest) GoString() string {
	return s.String()
}

func (s *GetServiceGroupRequest) SetClientToken(v string) *GetServiceGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *GetServiceGroupRequest) SetServiceGroupId(v int64) *GetServiceGroupRequest {
	s.ServiceGroupId = &v
	return s
}

type GetServiceGroupResponseBody struct {
	Data      *GetServiceGroupResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	RequestId *string                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetServiceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceGroupResponseBody) SetData(v *GetServiceGroupResponseBodyData) *GetServiceGroupResponseBody {
	s.Data = v
	return s
}

func (s *GetServiceGroupResponseBody) SetRequestId(v string) *GetServiceGroupResponseBody {
	s.RequestId = &v
	return s
}

type GetServiceGroupResponseBodyData struct {
	CreateTime              *string                                 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	EnableWebhook           *string                                 `json:"enableWebhook,omitempty" xml:"enableWebhook,omitempty"`
	ServiceGroupDescription *string                                 `json:"serviceGroupDescription,omitempty" xml:"serviceGroupDescription,omitempty"`
	ServiceGroupId          *int64                                  `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
	ServiceGroupName        *string                                 `json:"serviceGroupName,omitempty" xml:"serviceGroupName,omitempty"`
	UpdateTime              *string                                 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	Users                   []*GetServiceGroupResponseBodyDataUsers `json:"users,omitempty" xml:"users,omitempty" type:"Repeated"`
	WebhookLink             *string                                 `json:"webhookLink,omitempty" xml:"webhookLink,omitempty"`
	WebhookType             *string                                 `json:"webhookType,omitempty" xml:"webhookType,omitempty"`
}

func (s GetServiceGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetServiceGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetServiceGroupResponseBodyData) SetCreateTime(v string) *GetServiceGroupResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetServiceGroupResponseBodyData) SetEnableWebhook(v string) *GetServiceGroupResponseBodyData {
	s.EnableWebhook = &v
	return s
}

func (s *GetServiceGroupResponseBodyData) SetServiceGroupDescription(v string) *GetServiceGroupResponseBodyData {
	s.ServiceGroupDescription = &v
	return s
}

func (s *GetServiceGroupResponseBodyData) SetServiceGroupId(v int64) *GetServiceGroupResponseBodyData {
	s.ServiceGroupId = &v
	return s
}

func (s *GetServiceGroupResponseBodyData) SetServiceGroupName(v string) *GetServiceGroupResponseBodyData {
	s.ServiceGroupName = &v
	return s
}

func (s *GetServiceGroupResponseBodyData) SetUpdateTime(v string) *GetServiceGroupResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetServiceGroupResponseBodyData) SetUsers(v []*GetServiceGroupResponseBodyDataUsers) *GetServiceGroupResponseBodyData {
	s.Users = v
	return s
}

func (s *GetServiceGroupResponseBodyData) SetWebhookLink(v string) *GetServiceGroupResponseBodyData {
	s.WebhookLink = &v
	return s
}

func (s *GetServiceGroupResponseBodyData) SetWebhookType(v string) *GetServiceGroupResponseBodyData {
	s.WebhookType = &v
	return s
}

type GetServiceGroupResponseBodyDataUsers struct {
	Phone          *string `json:"phone,omitempty" xml:"phone,omitempty"`
	ServiceGroupId *int64  `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
	UserId         *int64  `json:"userId,omitempty" xml:"userId,omitempty"`
	UserName       *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s GetServiceGroupResponseBodyDataUsers) String() string {
	return tea.Prettify(s)
}

func (s GetServiceGroupResponseBodyDataUsers) GoString() string {
	return s.String()
}

func (s *GetServiceGroupResponseBodyDataUsers) SetPhone(v string) *GetServiceGroupResponseBodyDataUsers {
	s.Phone = &v
	return s
}

func (s *GetServiceGroupResponseBodyDataUsers) SetServiceGroupId(v int64) *GetServiceGroupResponseBodyDataUsers {
	s.ServiceGroupId = &v
	return s
}

func (s *GetServiceGroupResponseBodyDataUsers) SetUserId(v int64) *GetServiceGroupResponseBodyDataUsers {
	s.UserId = &v
	return s
}

func (s *GetServiceGroupResponseBodyDataUsers) SetUserName(v string) *GetServiceGroupResponseBodyDataUsers {
	s.UserName = &v
	return s
}

type GetServiceGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetServiceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetServiceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceGroupResponse) GoString() string {
	return s.String()
}

func (s *GetServiceGroupResponse) SetHeaders(v map[string]*string) *GetServiceGroupResponse {
	s.Headers = v
	return s
}

func (s *GetServiceGroupResponse) SetStatusCode(v int32) *GetServiceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceGroupResponse) SetBody(v *GetServiceGroupResponseBody) *GetServiceGroupResponse {
	s.Body = v
	return s
}

type GetServiceGroupPersonSchedulingRequest struct {
	ClientToken    *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	EndTime        *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	ServiceGroupId *int64  `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
	StartTime      *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	UserId         *int64  `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetServiceGroupPersonSchedulingRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceGroupPersonSchedulingRequest) GoString() string {
	return s.String()
}

func (s *GetServiceGroupPersonSchedulingRequest) SetClientToken(v string) *GetServiceGroupPersonSchedulingRequest {
	s.ClientToken = &v
	return s
}

func (s *GetServiceGroupPersonSchedulingRequest) SetEndTime(v string) *GetServiceGroupPersonSchedulingRequest {
	s.EndTime = &v
	return s
}

func (s *GetServiceGroupPersonSchedulingRequest) SetServiceGroupId(v int64) *GetServiceGroupPersonSchedulingRequest {
	s.ServiceGroupId = &v
	return s
}

func (s *GetServiceGroupPersonSchedulingRequest) SetStartTime(v string) *GetServiceGroupPersonSchedulingRequest {
	s.StartTime = &v
	return s
}

func (s *GetServiceGroupPersonSchedulingRequest) SetUserId(v int64) *GetServiceGroupPersonSchedulingRequest {
	s.UserId = &v
	return s
}

type GetServiceGroupPersonSchedulingResponseBody struct {
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetServiceGroupPersonSchedulingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceGroupPersonSchedulingResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceGroupPersonSchedulingResponseBody) SetData(v map[string]interface{}) *GetServiceGroupPersonSchedulingResponseBody {
	s.Data = v
	return s
}

func (s *GetServiceGroupPersonSchedulingResponseBody) SetRequestId(v string) *GetServiceGroupPersonSchedulingResponseBody {
	s.RequestId = &v
	return s
}

type GetServiceGroupPersonSchedulingResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetServiceGroupPersonSchedulingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetServiceGroupPersonSchedulingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceGroupPersonSchedulingResponse) GoString() string {
	return s.String()
}

func (s *GetServiceGroupPersonSchedulingResponse) SetHeaders(v map[string]*string) *GetServiceGroupPersonSchedulingResponse {
	s.Headers = v
	return s
}

func (s *GetServiceGroupPersonSchedulingResponse) SetStatusCode(v int32) *GetServiceGroupPersonSchedulingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceGroupPersonSchedulingResponse) SetBody(v *GetServiceGroupPersonSchedulingResponseBody) *GetServiceGroupPersonSchedulingResponse {
	s.Body = v
	return s
}

type GetServiceGroupSchedulingRequest struct {
	ClientToken    *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	ServiceGroupId *int64  `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
}

func (s GetServiceGroupSchedulingRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceGroupSchedulingRequest) GoString() string {
	return s.String()
}

func (s *GetServiceGroupSchedulingRequest) SetClientToken(v string) *GetServiceGroupSchedulingRequest {
	s.ClientToken = &v
	return s
}

func (s *GetServiceGroupSchedulingRequest) SetServiceGroupId(v int64) *GetServiceGroupSchedulingRequest {
	s.ServiceGroupId = &v
	return s
}

type GetServiceGroupSchedulingResponseBody struct {
	Data *GetServiceGroupSchedulingResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetServiceGroupSchedulingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceGroupSchedulingResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceGroupSchedulingResponseBody) SetData(v *GetServiceGroupSchedulingResponseBodyData) *GetServiceGroupSchedulingResponseBody {
	s.Data = v
	return s
}

func (s *GetServiceGroupSchedulingResponseBody) SetRequestId(v string) *GetServiceGroupSchedulingResponseBody {
	s.RequestId = &v
	return s
}

type GetServiceGroupSchedulingResponseBodyData struct {
	FastScheduling *GetServiceGroupSchedulingResponseBodyDataFastScheduling `json:"fastScheduling,omitempty" xml:"fastScheduling,omitempty" type:"Struct"`
	FineScheduling *GetServiceGroupSchedulingResponseBodyDataFineScheduling `json:"fineScheduling,omitempty" xml:"fineScheduling,omitempty" type:"Struct"`
	SchedulingWay  *string                                                  `json:"schedulingWay,omitempty" xml:"schedulingWay,omitempty"`
	ServiceGroupId *int64                                                   `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
	Users          []*GetServiceGroupSchedulingResponseBodyDataUsers        `json:"users,omitempty" xml:"users,omitempty" type:"Repeated"`
}

func (s GetServiceGroupSchedulingResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetServiceGroupSchedulingResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetServiceGroupSchedulingResponseBodyData) SetFastScheduling(v *GetServiceGroupSchedulingResponseBodyDataFastScheduling) *GetServiceGroupSchedulingResponseBodyData {
	s.FastScheduling = v
	return s
}

func (s *GetServiceGroupSchedulingResponseBodyData) SetFineScheduling(v *GetServiceGroupSchedulingResponseBodyDataFineScheduling) *GetServiceGroupSchedulingResponseBodyData {
	s.FineScheduling = v
	return s
}

func (s *GetServiceGroupSchedulingResponseBodyData) SetSchedulingWay(v string) *GetServiceGroupSchedulingResponseBodyData {
	s.SchedulingWay = &v
	return s
}

func (s *GetServiceGroupSchedulingResponseBodyData) SetServiceGroupId(v int64) *GetServiceGroupSchedulingResponseBodyData {
	s.ServiceGroupId = &v
	return s
}

func (s *GetServiceGroupSchedulingResponseBodyData) SetUsers(v []*GetServiceGroupSchedulingResponseBodyDataUsers) *GetServiceGroupSchedulingResponseBodyData {
	s.Users = v
	return s
}

type GetServiceGroupSchedulingResponseBodyDataFastScheduling struct {
	DutyPlan           *string                                                                   `json:"dutyPlan,omitempty" xml:"dutyPlan,omitempty"`
	Id                 *int64                                                                    `json:"id,omitempty" xml:"id,omitempty"`
	SchedulingUsers    []*GetServiceGroupSchedulingResponseBodyDataFastSchedulingSchedulingUsers `json:"schedulingUsers,omitempty" xml:"schedulingUsers,omitempty" type:"Repeated"`
	SingleDuration     *int32                                                                    `json:"singleDuration,omitempty" xml:"singleDuration,omitempty"`
	SingleDurationUnit *string                                                                   `json:"singleDurationUnit,omitempty" xml:"singleDurationUnit,omitempty"`
}

func (s GetServiceGroupSchedulingResponseBodyDataFastScheduling) String() string {
	return tea.Prettify(s)
}

func (s GetServiceGroupSchedulingResponseBodyDataFastScheduling) GoString() string {
	return s.String()
}

func (s *GetServiceGroupSchedulingResponseBodyDataFastScheduling) SetDutyPlan(v string) *GetServiceGroupSchedulingResponseBodyDataFastScheduling {
	s.DutyPlan = &v
	return s
}

func (s *GetServiceGroupSchedulingResponseBodyDataFastScheduling) SetId(v int64) *GetServiceGroupSchedulingResponseBodyDataFastScheduling {
	s.Id = &v
	return s
}

func (s *GetServiceGroupSchedulingResponseBodyDataFastScheduling) SetSchedulingUsers(v []*GetServiceGroupSchedulingResponseBodyDataFastSchedulingSchedulingUsers) *GetServiceGroupSchedulingResponseBodyDataFastScheduling {
	s.SchedulingUsers = v
	return s
}

func (s *GetServiceGroupSchedulingResponseBodyDataFastScheduling) SetSingleDuration(v int32) *GetServiceGroupSchedulingResponseBodyDataFastScheduling {
	s.SingleDuration = &v
	return s
}

func (s *GetServiceGroupSchedulingResponseBodyDataFastScheduling) SetSingleDurationUnit(v string) *GetServiceGroupSchedulingResponseBodyDataFastScheduling {
	s.SingleDurationUnit = &v
	return s
}

type GetServiceGroupSchedulingResponseBodyDataFastSchedulingSchedulingUsers struct {
	SchedulingOrder    *int32  `json:"schedulingOrder,omitempty" xml:"schedulingOrder,omitempty"`
	SchedulingUserId   *int64  `json:"schedulingUserId,omitempty" xml:"schedulingUserId,omitempty"`
	SchedulingUserName *string `json:"schedulingUserName,omitempty" xml:"schedulingUserName,omitempty"`
}

func (s GetServiceGroupSchedulingResponseBodyDataFastSchedulingSchedulingUsers) String() string {
	return tea.Prettify(s)
}

func (s GetServiceGroupSchedulingResponseBodyDataFastSchedulingSchedulingUsers) GoString() string {
	return s.String()
}

func (s *GetServiceGroupSchedulingResponseBodyDataFastSchedulingSchedulingUsers) SetSchedulingOrder(v int32) *GetServiceGroupSchedulingResponseBodyDataFastSchedulingSchedulingUsers {
	s.SchedulingOrder = &v
	return s
}

func (s *GetServiceGroupSchedulingResponseBodyDataFastSchedulingSchedulingUsers) SetSchedulingUserId(v int64) *GetServiceGroupSchedulingResponseBodyDataFastSchedulingSchedulingUsers {
	s.SchedulingUserId = &v
	return s
}

func (s *GetServiceGroupSchedulingResponseBodyDataFastSchedulingSchedulingUsers) SetSchedulingUserName(v string) *GetServiceGroupSchedulingResponseBodyDataFastSchedulingSchedulingUsers {
	s.SchedulingUserName = &v
	return s
}

type GetServiceGroupSchedulingResponseBodyDataFineScheduling struct {
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 1
	Period                       *int32                                                                                 `json:"period,omitempty" xml:"period,omitempty"`
	PeriodUnit                   *string                                                                                `json:"periodUnit,omitempty" xml:"periodUnit,omitempty"`
	SchedulingFineShifts         []*GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingFineShifts         `json:"schedulingFineShifts,omitempty" xml:"schedulingFineShifts,omitempty" type:"Repeated"`
	SchedulingTemplateFineShifts []*GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingTemplateFineShifts `json:"schedulingTemplateFineShifts,omitempty" xml:"schedulingTemplateFineShifts,omitempty" type:"Repeated"`
	ShiftType                    *string                                                                                `json:"shiftType,omitempty" xml:"shiftType,omitempty"`
}

func (s GetServiceGroupSchedulingResponseBodyDataFineScheduling) String() string {
	return tea.Prettify(s)
}

func (s GetServiceGroupSchedulingResponseBodyDataFineScheduling) GoString() string {
	return s.String()
}

func (s *GetServiceGroupSchedulingResponseBodyDataFineScheduling) SetId(v int64) *GetServiceGroupSchedulingResponseBodyDataFineScheduling {
	s.Id = &v
	return s
}

func (s *GetServiceGroupSchedulingResponseBodyDataFineScheduling) SetPeriod(v int32) *GetServiceGroupSchedulingResponseBodyDataFineScheduling {
	s.Period = &v
	return s
}

func (s *GetServiceGroupSchedulingResponseBodyDataFineScheduling) SetPeriodUnit(v string) *GetServiceGroupSchedulingResponseBodyDataFineScheduling {
	s.PeriodUnit = &v
	return s
}

func (s *GetServiceGroupSchedulingResponseBodyDataFineScheduling) SetSchedulingFineShifts(v []*GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingFineShifts) *GetServiceGroupSchedulingResponseBodyDataFineScheduling {
	s.SchedulingFineShifts = v
	return s
}

func (s *GetServiceGroupSchedulingResponseBodyDataFineScheduling) SetSchedulingTemplateFineShifts(v []*GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingTemplateFineShifts) *GetServiceGroupSchedulingResponseBodyDataFineScheduling {
	s.SchedulingTemplateFineShifts = v
	return s
}

func (s *GetServiceGroupSchedulingResponseBodyDataFineScheduling) SetShiftType(v string) *GetServiceGroupSchedulingResponseBodyDataFineScheduling {
	s.ShiftType = &v
	return s
}

type GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingFineShifts struct {
	CycleOrder          *int64  `json:"cycleOrder,omitempty" xml:"cycleOrder,omitempty"`
	SchedulingEndTime   *string `json:"schedulingEndTime,omitempty" xml:"schedulingEndTime,omitempty"`
	SchedulingOrder     *int32  `json:"schedulingOrder,omitempty" xml:"schedulingOrder,omitempty"`
	SchedulingStartTime *string `json:"schedulingStartTime,omitempty" xml:"schedulingStartTime,omitempty"`
	SchedulingUserId    *int64  `json:"schedulingUserId,omitempty" xml:"schedulingUserId,omitempty"`
	SchedulingUserName  *string `json:"schedulingUserName,omitempty" xml:"schedulingUserName,omitempty"`
	ShiftName           *string `json:"shiftName,omitempty" xml:"shiftName,omitempty"`
	SkipOneDay          *bool   `json:"skipOneDay,omitempty" xml:"skipOneDay,omitempty"`
}

func (s GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingFineShifts) String() string {
	return tea.Prettify(s)
}

func (s GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingFineShifts) GoString() string {
	return s.String()
}

func (s *GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingFineShifts) SetCycleOrder(v int64) *GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingFineShifts {
	s.CycleOrder = &v
	return s
}

func (s *GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingFineShifts) SetSchedulingEndTime(v string) *GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingFineShifts {
	s.SchedulingEndTime = &v
	return s
}

func (s *GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingFineShifts) SetSchedulingOrder(v int32) *GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingFineShifts {
	s.SchedulingOrder = &v
	return s
}

func (s *GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingFineShifts) SetSchedulingStartTime(v string) *GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingFineShifts {
	s.SchedulingStartTime = &v
	return s
}

func (s *GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingFineShifts) SetSchedulingUserId(v int64) *GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingFineShifts {
	s.SchedulingUserId = &v
	return s
}

func (s *GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingFineShifts) SetSchedulingUserName(v string) *GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingFineShifts {
	s.SchedulingUserName = &v
	return s
}

func (s *GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingFineShifts) SetShiftName(v string) *GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingFineShifts {
	s.ShiftName = &v
	return s
}

func (s *GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingFineShifts) SetSkipOneDay(v bool) *GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingFineShifts {
	s.SkipOneDay = &v
	return s
}

type GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingTemplateFineShifts struct {
	SchedulingEndTime   *string `json:"schedulingEndTime,omitempty" xml:"schedulingEndTime,omitempty"`
	SchedulingOrder     *int64  `json:"schedulingOrder,omitempty" xml:"schedulingOrder,omitempty"`
	SchedulingStartTime *string `json:"schedulingStartTime,omitempty" xml:"schedulingStartTime,omitempty"`
	SchedulingUserId    *string `json:"schedulingUserId,omitempty" xml:"schedulingUserId,omitempty"`
	SchedulingUserName  *string `json:"schedulingUserName,omitempty" xml:"schedulingUserName,omitempty"`
	ShiftName           *string `json:"shiftName,omitempty" xml:"shiftName,omitempty"`
	SkipOneDay          *bool   `json:"skipOneDay,omitempty" xml:"skipOneDay,omitempty"`
}

func (s GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingTemplateFineShifts) String() string {
	return tea.Prettify(s)
}

func (s GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingTemplateFineShifts) GoString() string {
	return s.String()
}

func (s *GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingTemplateFineShifts) SetSchedulingEndTime(v string) *GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingTemplateFineShifts {
	s.SchedulingEndTime = &v
	return s
}

func (s *GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingTemplateFineShifts) SetSchedulingOrder(v int64) *GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingTemplateFineShifts {
	s.SchedulingOrder = &v
	return s
}

func (s *GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingTemplateFineShifts) SetSchedulingStartTime(v string) *GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingTemplateFineShifts {
	s.SchedulingStartTime = &v
	return s
}

func (s *GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingTemplateFineShifts) SetSchedulingUserId(v string) *GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingTemplateFineShifts {
	s.SchedulingUserId = &v
	return s
}

func (s *GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingTemplateFineShifts) SetSchedulingUserName(v string) *GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingTemplateFineShifts {
	s.SchedulingUserName = &v
	return s
}

func (s *GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingTemplateFineShifts) SetShiftName(v string) *GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingTemplateFineShifts {
	s.ShiftName = &v
	return s
}

func (s *GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingTemplateFineShifts) SetSkipOneDay(v bool) *GetServiceGroupSchedulingResponseBodyDataFineSchedulingSchedulingTemplateFineShifts {
	s.SkipOneDay = &v
	return s
}

type GetServiceGroupSchedulingResponseBodyDataUsers struct {
	UserId   *int64  `json:"userId,omitempty" xml:"userId,omitempty"`
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s GetServiceGroupSchedulingResponseBodyDataUsers) String() string {
	return tea.Prettify(s)
}

func (s GetServiceGroupSchedulingResponseBodyDataUsers) GoString() string {
	return s.String()
}

func (s *GetServiceGroupSchedulingResponseBodyDataUsers) SetUserId(v int64) *GetServiceGroupSchedulingResponseBodyDataUsers {
	s.UserId = &v
	return s
}

func (s *GetServiceGroupSchedulingResponseBodyDataUsers) SetUserName(v string) *GetServiceGroupSchedulingResponseBodyDataUsers {
	s.UserName = &v
	return s
}

type GetServiceGroupSchedulingResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetServiceGroupSchedulingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetServiceGroupSchedulingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceGroupSchedulingResponse) GoString() string {
	return s.String()
}

func (s *GetServiceGroupSchedulingResponse) SetHeaders(v map[string]*string) *GetServiceGroupSchedulingResponse {
	s.Headers = v
	return s
}

func (s *GetServiceGroupSchedulingResponse) SetStatusCode(v int32) *GetServiceGroupSchedulingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceGroupSchedulingResponse) SetBody(v *GetServiceGroupSchedulingResponseBody) *GetServiceGroupSchedulingResponse {
	s.Body = v
	return s
}

type GetServiceGroupSchedulingPreviewRequest struct {
	ClientToken    *string                                                `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	EndTime        *string                                                `json:"endTime,omitempty" xml:"endTime,omitempty"`
	FastScheduling *GetServiceGroupSchedulingPreviewRequestFastScheduling `json:"fastScheduling,omitempty" xml:"fastScheduling,omitempty" type:"Struct"`
	FineScheduling *GetServiceGroupSchedulingPreviewRequestFineScheduling `json:"fineScheduling,omitempty" xml:"fineScheduling,omitempty" type:"Struct"`
	SchedulingWay  *string                                                `json:"schedulingWay,omitempty" xml:"schedulingWay,omitempty"`
	ServiceGroupId *int64                                                 `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
	StartTime      *string                                                `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetServiceGroupSchedulingPreviewRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceGroupSchedulingPreviewRequest) GoString() string {
	return s.String()
}

func (s *GetServiceGroupSchedulingPreviewRequest) SetClientToken(v string) *GetServiceGroupSchedulingPreviewRequest {
	s.ClientToken = &v
	return s
}

func (s *GetServiceGroupSchedulingPreviewRequest) SetEndTime(v string) *GetServiceGroupSchedulingPreviewRequest {
	s.EndTime = &v
	return s
}

func (s *GetServiceGroupSchedulingPreviewRequest) SetFastScheduling(v *GetServiceGroupSchedulingPreviewRequestFastScheduling) *GetServiceGroupSchedulingPreviewRequest {
	s.FastScheduling = v
	return s
}

func (s *GetServiceGroupSchedulingPreviewRequest) SetFineScheduling(v *GetServiceGroupSchedulingPreviewRequestFineScheduling) *GetServiceGroupSchedulingPreviewRequest {
	s.FineScheduling = v
	return s
}

func (s *GetServiceGroupSchedulingPreviewRequest) SetSchedulingWay(v string) *GetServiceGroupSchedulingPreviewRequest {
	s.SchedulingWay = &v
	return s
}

func (s *GetServiceGroupSchedulingPreviewRequest) SetServiceGroupId(v int64) *GetServiceGroupSchedulingPreviewRequest {
	s.ServiceGroupId = &v
	return s
}

func (s *GetServiceGroupSchedulingPreviewRequest) SetStartTime(v string) *GetServiceGroupSchedulingPreviewRequest {
	s.StartTime = &v
	return s
}

type GetServiceGroupSchedulingPreviewRequestFastScheduling struct {
	// FAST_CHOICE
	DutyPlan        *string                                                                 `json:"dutyPlan,omitempty" xml:"dutyPlan,omitempty"`
	SchedulingUsers []*GetServiceGroupSchedulingPreviewRequestFastSchedulingSchedulingUsers `json:"schedulingUsers,omitempty" xml:"schedulingUsers,omitempty" type:"Repeated"`
	SingleDuration  *int32                                                                  `json:"singleDuration,omitempty" xml:"singleDuration,omitempty"`
	// DAY
	SingleDurationUnit *string `json:"singleDurationUnit,omitempty" xml:"singleDurationUnit,omitempty"`
}

func (s GetServiceGroupSchedulingPreviewRequestFastScheduling) String() string {
	return tea.Prettify(s)
}

func (s GetServiceGroupSchedulingPreviewRequestFastScheduling) GoString() string {
	return s.String()
}

func (s *GetServiceGroupSchedulingPreviewRequestFastScheduling) SetDutyPlan(v string) *GetServiceGroupSchedulingPreviewRequestFastScheduling {
	s.DutyPlan = &v
	return s
}

func (s *GetServiceGroupSchedulingPreviewRequestFastScheduling) SetSchedulingUsers(v []*GetServiceGroupSchedulingPreviewRequestFastSchedulingSchedulingUsers) *GetServiceGroupSchedulingPreviewRequestFastScheduling {
	s.SchedulingUsers = v
	return s
}

func (s *GetServiceGroupSchedulingPreviewRequestFastScheduling) SetSingleDuration(v int32) *GetServiceGroupSchedulingPreviewRequestFastScheduling {
	s.SingleDuration = &v
	return s
}

func (s *GetServiceGroupSchedulingPreviewRequestFastScheduling) SetSingleDurationUnit(v string) *GetServiceGroupSchedulingPreviewRequestFastScheduling {
	s.SingleDurationUnit = &v
	return s
}

type GetServiceGroupSchedulingPreviewRequestFastSchedulingSchedulingUsers struct {
	SchedulingOrder  *int32 `json:"schedulingOrder,omitempty" xml:"schedulingOrder,omitempty"`
	SchedulingUserId *int64 `json:"schedulingUserId,omitempty" xml:"schedulingUserId,omitempty"`
}

func (s GetServiceGroupSchedulingPreviewRequestFastSchedulingSchedulingUsers) String() string {
	return tea.Prettify(s)
}

func (s GetServiceGroupSchedulingPreviewRequestFastSchedulingSchedulingUsers) GoString() string {
	return s.String()
}

func (s *GetServiceGroupSchedulingPreviewRequestFastSchedulingSchedulingUsers) SetSchedulingOrder(v int32) *GetServiceGroupSchedulingPreviewRequestFastSchedulingSchedulingUsers {
	s.SchedulingOrder = &v
	return s
}

func (s *GetServiceGroupSchedulingPreviewRequestFastSchedulingSchedulingUsers) SetSchedulingUserId(v int64) *GetServiceGroupSchedulingPreviewRequestFastSchedulingSchedulingUsers {
	s.SchedulingUserId = &v
	return s
}

type GetServiceGroupSchedulingPreviewRequestFineScheduling struct {
	Period               *int32                                                                       `json:"period,omitempty" xml:"period,omitempty"`
	PeriodUnit           *string                                                                      `json:"periodUnit,omitempty" xml:"periodUnit,omitempty"`
	SchedulingFineShifts []*GetServiceGroupSchedulingPreviewRequestFineSchedulingSchedulingFineShifts `json:"schedulingFineShifts,omitempty" xml:"schedulingFineShifts,omitempty" type:"Repeated"`
	ShiftType            *string                                                                      `json:"shiftType,omitempty" xml:"shiftType,omitempty"`
}

func (s GetServiceGroupSchedulingPreviewRequestFineScheduling) String() string {
	return tea.Prettify(s)
}

func (s GetServiceGroupSchedulingPreviewRequestFineScheduling) GoString() string {
	return s.String()
}

func (s *GetServiceGroupSchedulingPreviewRequestFineScheduling) SetPeriod(v int32) *GetServiceGroupSchedulingPreviewRequestFineScheduling {
	s.Period = &v
	return s
}

func (s *GetServiceGroupSchedulingPreviewRequestFineScheduling) SetPeriodUnit(v string) *GetServiceGroupSchedulingPreviewRequestFineScheduling {
	s.PeriodUnit = &v
	return s
}

func (s *GetServiceGroupSchedulingPreviewRequestFineScheduling) SetSchedulingFineShifts(v []*GetServiceGroupSchedulingPreviewRequestFineSchedulingSchedulingFineShifts) *GetServiceGroupSchedulingPreviewRequestFineScheduling {
	s.SchedulingFineShifts = v
	return s
}

func (s *GetServiceGroupSchedulingPreviewRequestFineScheduling) SetShiftType(v string) *GetServiceGroupSchedulingPreviewRequestFineScheduling {
	s.ShiftType = &v
	return s
}

type GetServiceGroupSchedulingPreviewRequestFineSchedulingSchedulingFineShifts struct {
	SchedulingEndTime   *string `json:"schedulingEndTime,omitempty" xml:"schedulingEndTime,omitempty"`
	SchedulingOrder     *int64  `json:"schedulingOrder,omitempty" xml:"schedulingOrder,omitempty"`
	SchedulingStartTime *string `json:"schedulingStartTime,omitempty" xml:"schedulingStartTime,omitempty"`
	ShiftName           *string `json:"shiftName,omitempty" xml:"shiftName,omitempty"`
}

func (s GetServiceGroupSchedulingPreviewRequestFineSchedulingSchedulingFineShifts) String() string {
	return tea.Prettify(s)
}

func (s GetServiceGroupSchedulingPreviewRequestFineSchedulingSchedulingFineShifts) GoString() string {
	return s.String()
}

func (s *GetServiceGroupSchedulingPreviewRequestFineSchedulingSchedulingFineShifts) SetSchedulingEndTime(v string) *GetServiceGroupSchedulingPreviewRequestFineSchedulingSchedulingFineShifts {
	s.SchedulingEndTime = &v
	return s
}

func (s *GetServiceGroupSchedulingPreviewRequestFineSchedulingSchedulingFineShifts) SetSchedulingOrder(v int64) *GetServiceGroupSchedulingPreviewRequestFineSchedulingSchedulingFineShifts {
	s.SchedulingOrder = &v
	return s
}

func (s *GetServiceGroupSchedulingPreviewRequestFineSchedulingSchedulingFineShifts) SetSchedulingStartTime(v string) *GetServiceGroupSchedulingPreviewRequestFineSchedulingSchedulingFineShifts {
	s.SchedulingStartTime = &v
	return s
}

func (s *GetServiceGroupSchedulingPreviewRequestFineSchedulingSchedulingFineShifts) SetShiftName(v string) *GetServiceGroupSchedulingPreviewRequestFineSchedulingSchedulingFineShifts {
	s.ShiftName = &v
	return s
}

type GetServiceGroupSchedulingPreviewResponseBody struct {
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetServiceGroupSchedulingPreviewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceGroupSchedulingPreviewResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceGroupSchedulingPreviewResponseBody) SetData(v map[string]interface{}) *GetServiceGroupSchedulingPreviewResponseBody {
	s.Data = v
	return s
}

func (s *GetServiceGroupSchedulingPreviewResponseBody) SetRequestId(v string) *GetServiceGroupSchedulingPreviewResponseBody {
	s.RequestId = &v
	return s
}

type GetServiceGroupSchedulingPreviewResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetServiceGroupSchedulingPreviewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetServiceGroupSchedulingPreviewResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceGroupSchedulingPreviewResponse) GoString() string {
	return s.String()
}

func (s *GetServiceGroupSchedulingPreviewResponse) SetHeaders(v map[string]*string) *GetServiceGroupSchedulingPreviewResponse {
	s.Headers = v
	return s
}

func (s *GetServiceGroupSchedulingPreviewResponse) SetStatusCode(v int32) *GetServiceGroupSchedulingPreviewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceGroupSchedulingPreviewResponse) SetBody(v *GetServiceGroupSchedulingPreviewResponseBody) *GetServiceGroupSchedulingPreviewResponse {
	s.Body = v
	return s
}

type GetServiceGroupSpecialPersonSchedulingRequest struct {
	ClientToken    *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	ServiceGroupId *int64  `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
	UserId         *int64  `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetServiceGroupSpecialPersonSchedulingRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceGroupSpecialPersonSchedulingRequest) GoString() string {
	return s.String()
}

func (s *GetServiceGroupSpecialPersonSchedulingRequest) SetClientToken(v string) *GetServiceGroupSpecialPersonSchedulingRequest {
	s.ClientToken = &v
	return s
}

func (s *GetServiceGroupSpecialPersonSchedulingRequest) SetServiceGroupId(v int64) *GetServiceGroupSpecialPersonSchedulingRequest {
	s.ServiceGroupId = &v
	return s
}

func (s *GetServiceGroupSpecialPersonSchedulingRequest) SetUserId(v int64) *GetServiceGroupSpecialPersonSchedulingRequest {
	s.UserId = &v
	return s
}

type GetServiceGroupSpecialPersonSchedulingResponseBody struct {
	Data []*GetServiceGroupSpecialPersonSchedulingResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetServiceGroupSpecialPersonSchedulingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceGroupSpecialPersonSchedulingResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceGroupSpecialPersonSchedulingResponseBody) SetData(v []*GetServiceGroupSpecialPersonSchedulingResponseBodyData) *GetServiceGroupSpecialPersonSchedulingResponseBody {
	s.Data = v
	return s
}

func (s *GetServiceGroupSpecialPersonSchedulingResponseBody) SetRequestId(v string) *GetServiceGroupSpecialPersonSchedulingResponseBody {
	s.RequestId = &v
	return s
}

type GetServiceGroupSpecialPersonSchedulingResponseBodyData struct {
	SchedulingDate      *string `json:"schedulingDate,omitempty" xml:"schedulingDate,omitempty"`
	SchedulingEndTime   *string `json:"schedulingEndTime,omitempty" xml:"schedulingEndTime,omitempty"`
	SchedulingStartTime *string `json:"schedulingStartTime,omitempty" xml:"schedulingStartTime,omitempty"`
	SchedulingUserId    *int64  `json:"schedulingUserId,omitempty" xml:"schedulingUserId,omitempty"`
	ServiceGroupId      *int64  `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
	ServiceGroupName    *string `json:"serviceGroupName,omitempty" xml:"serviceGroupName,omitempty"`
}

func (s GetServiceGroupSpecialPersonSchedulingResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetServiceGroupSpecialPersonSchedulingResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetServiceGroupSpecialPersonSchedulingResponseBodyData) SetSchedulingDate(v string) *GetServiceGroupSpecialPersonSchedulingResponseBodyData {
	s.SchedulingDate = &v
	return s
}

func (s *GetServiceGroupSpecialPersonSchedulingResponseBodyData) SetSchedulingEndTime(v string) *GetServiceGroupSpecialPersonSchedulingResponseBodyData {
	s.SchedulingEndTime = &v
	return s
}

func (s *GetServiceGroupSpecialPersonSchedulingResponseBodyData) SetSchedulingStartTime(v string) *GetServiceGroupSpecialPersonSchedulingResponseBodyData {
	s.SchedulingStartTime = &v
	return s
}

func (s *GetServiceGroupSpecialPersonSchedulingResponseBodyData) SetSchedulingUserId(v int64) *GetServiceGroupSpecialPersonSchedulingResponseBodyData {
	s.SchedulingUserId = &v
	return s
}

func (s *GetServiceGroupSpecialPersonSchedulingResponseBodyData) SetServiceGroupId(v int64) *GetServiceGroupSpecialPersonSchedulingResponseBodyData {
	s.ServiceGroupId = &v
	return s
}

func (s *GetServiceGroupSpecialPersonSchedulingResponseBodyData) SetServiceGroupName(v string) *GetServiceGroupSpecialPersonSchedulingResponseBodyData {
	s.ServiceGroupName = &v
	return s
}

type GetServiceGroupSpecialPersonSchedulingResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetServiceGroupSpecialPersonSchedulingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetServiceGroupSpecialPersonSchedulingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceGroupSpecialPersonSchedulingResponse) GoString() string {
	return s.String()
}

func (s *GetServiceGroupSpecialPersonSchedulingResponse) SetHeaders(v map[string]*string) *GetServiceGroupSpecialPersonSchedulingResponse {
	s.Headers = v
	return s
}

func (s *GetServiceGroupSpecialPersonSchedulingResponse) SetStatusCode(v int32) *GetServiceGroupSpecialPersonSchedulingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceGroupSpecialPersonSchedulingResponse) SetBody(v *GetServiceGroupSpecialPersonSchedulingResponseBody) *GetServiceGroupSpecialPersonSchedulingResponse {
	s.Body = v
	return s
}

type GetSimilarIncidentStatisticsRequest struct {
	ClientToken      *string   `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	CreateTime       *string   `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Events           []*string `json:"events,omitempty" xml:"events,omitempty" type:"Repeated"`
	IncidentId       *int64    `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
	IncidentTitle    *string   `json:"incidentTitle,omitempty" xml:"incidentTitle,omitempty"`
	RelatedServiceId *int64    `json:"relatedServiceId,omitempty" xml:"relatedServiceId,omitempty"`
}

func (s GetSimilarIncidentStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSimilarIncidentStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetSimilarIncidentStatisticsRequest) SetClientToken(v string) *GetSimilarIncidentStatisticsRequest {
	s.ClientToken = &v
	return s
}

func (s *GetSimilarIncidentStatisticsRequest) SetCreateTime(v string) *GetSimilarIncidentStatisticsRequest {
	s.CreateTime = &v
	return s
}

func (s *GetSimilarIncidentStatisticsRequest) SetEvents(v []*string) *GetSimilarIncidentStatisticsRequest {
	s.Events = v
	return s
}

func (s *GetSimilarIncidentStatisticsRequest) SetIncidentId(v int64) *GetSimilarIncidentStatisticsRequest {
	s.IncidentId = &v
	return s
}

func (s *GetSimilarIncidentStatisticsRequest) SetIncidentTitle(v string) *GetSimilarIncidentStatisticsRequest {
	s.IncidentTitle = &v
	return s
}

func (s *GetSimilarIncidentStatisticsRequest) SetRelatedServiceId(v int64) *GetSimilarIncidentStatisticsRequest {
	s.RelatedServiceId = &v
	return s
}

type GetSimilarIncidentStatisticsResponseBody struct {
	// data
	Data *GetSimilarIncidentStatisticsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetSimilarIncidentStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSimilarIncidentStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetSimilarIncidentStatisticsResponseBody) SetData(v *GetSimilarIncidentStatisticsResponseBodyData) *GetSimilarIncidentStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBody) SetRequestId(v string) *GetSimilarIncidentStatisticsResponseBody {
	s.RequestId = &v
	return s
}

type GetSimilarIncidentStatisticsResponseBodyData struct {
	CountInSevenDays      *int64                                                               `json:"countInSevenDays,omitempty" xml:"countInSevenDays,omitempty"`
	CountInSixMonths      *int64                                                               `json:"countInSixMonths,omitempty" xml:"countInSixMonths,omitempty"`
	DailySimilarIncidents []*GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidents `json:"dailySimilarIncidents,omitempty" xml:"dailySimilarIncidents,omitempty" type:"Repeated"`
	// id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// topFiveIncidents
	TopFiveIncidents []*GetSimilarIncidentStatisticsResponseBodyDataTopFiveIncidents `json:"topFiveIncidents,omitempty" xml:"topFiveIncidents,omitempty" type:"Repeated"`
}

func (s GetSimilarIncidentStatisticsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSimilarIncidentStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSimilarIncidentStatisticsResponseBodyData) SetCountInSevenDays(v int64) *GetSimilarIncidentStatisticsResponseBodyData {
	s.CountInSevenDays = &v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBodyData) SetCountInSixMonths(v int64) *GetSimilarIncidentStatisticsResponseBodyData {
	s.CountInSixMonths = &v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBodyData) SetDailySimilarIncidents(v []*GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidents) *GetSimilarIncidentStatisticsResponseBodyData {
	s.DailySimilarIncidents = v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBodyData) SetRequestId(v string) *GetSimilarIncidentStatisticsResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBodyData) SetTopFiveIncidents(v []*GetSimilarIncidentStatisticsResponseBodyDataTopFiveIncidents) *GetSimilarIncidentStatisticsResponseBodyData {
	s.TopFiveIncidents = v
	return s
}

type GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidents struct {
	Commitment       *int64                                                                               `json:"commitment,omitempty" xml:"commitment,omitempty"`
	Date             *string                                                                              `json:"date,omitempty" xml:"date,omitempty"`
	Day              *int64                                                                               `json:"day,omitempty" xml:"day,omitempty"`
	Month            *int64                                                                               `json:"month,omitempty" xml:"month,omitempty"`
	SimilarIncidents []*GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidentsSimilarIncidents `json:"similarIncidents,omitempty" xml:"similarIncidents,omitempty" type:"Repeated"`
	Week             *string                                                                              `json:"week,omitempty" xml:"week,omitempty"`
}

func (s GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidents) String() string {
	return tea.Prettify(s)
}

func (s GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidents) GoString() string {
	return s.String()
}

func (s *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidents) SetCommitment(v int64) *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidents {
	s.Commitment = &v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidents) SetDate(v string) *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidents {
	s.Date = &v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidents) SetDay(v int64) *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidents {
	s.Day = &v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidents) SetMonth(v int64) *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidents {
	s.Month = &v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidents) SetSimilarIncidents(v []*GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidentsSimilarIncidents) *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidents {
	s.SimilarIncidents = v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidents) SetWeek(v string) *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidents {
	s.Week = &v
	return s
}

type GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidentsSimilarIncidents struct {
	AssignUserId              *int64  `json:"assignUserId,omitempty" xml:"assignUserId,omitempty"`
	AssignUserName            *string `json:"assignUserName,omitempty" xml:"assignUserName,omitempty"`
	CreateTime                *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DurationTime              *int64  `json:"durationTime,omitempty" xml:"durationTime,omitempty"`
	FinishReason              *int64  `json:"finishReason,omitempty" xml:"finishReason,omitempty"`
	FinishReasonDescription   *string `json:"finishReasonDescription,omitempty" xml:"finishReasonDescription,omitempty"`
	FinishSolutionDescription *string `json:"finishSolutionDescription,omitempty" xml:"finishSolutionDescription,omitempty"`
	IncidentFinishSolution    *int64  `json:"incidentFinishSolution,omitempty" xml:"incidentFinishSolution,omitempty"`
	IncidentId                *int64  `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
	IncidentNumber            *string `json:"incidentNumber,omitempty" xml:"incidentNumber,omitempty"`
	IncidentTitle             *string `json:"incidentTitle,omitempty" xml:"incidentTitle,omitempty"`
	RelatedRouteRuleId        *int64  `json:"relatedRouteRuleId,omitempty" xml:"relatedRouteRuleId,omitempty"`
	RelatedRouteRuleName      *string `json:"relatedRouteRuleName,omitempty" xml:"relatedRouteRuleName,omitempty"`
	SimilarScore              *string `json:"similarScore,omitempty" xml:"similarScore,omitempty"`
}

func (s GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidentsSimilarIncidents) String() string {
	return tea.Prettify(s)
}

func (s GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidentsSimilarIncidents) GoString() string {
	return s.String()
}

func (s *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidentsSimilarIncidents) SetAssignUserId(v int64) *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidentsSimilarIncidents {
	s.AssignUserId = &v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidentsSimilarIncidents) SetAssignUserName(v string) *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidentsSimilarIncidents {
	s.AssignUserName = &v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidentsSimilarIncidents) SetCreateTime(v string) *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidentsSimilarIncidents {
	s.CreateTime = &v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidentsSimilarIncidents) SetDurationTime(v int64) *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidentsSimilarIncidents {
	s.DurationTime = &v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidentsSimilarIncidents) SetFinishReason(v int64) *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidentsSimilarIncidents {
	s.FinishReason = &v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidentsSimilarIncidents) SetFinishReasonDescription(v string) *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidentsSimilarIncidents {
	s.FinishReasonDescription = &v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidentsSimilarIncidents) SetFinishSolutionDescription(v string) *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidentsSimilarIncidents {
	s.FinishSolutionDescription = &v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidentsSimilarIncidents) SetIncidentFinishSolution(v int64) *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidentsSimilarIncidents {
	s.IncidentFinishSolution = &v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidentsSimilarIncidents) SetIncidentId(v int64) *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidentsSimilarIncidents {
	s.IncidentId = &v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidentsSimilarIncidents) SetIncidentNumber(v string) *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidentsSimilarIncidents {
	s.IncidentNumber = &v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidentsSimilarIncidents) SetIncidentTitle(v string) *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidentsSimilarIncidents {
	s.IncidentTitle = &v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidentsSimilarIncidents) SetRelatedRouteRuleId(v int64) *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidentsSimilarIncidents {
	s.RelatedRouteRuleId = &v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidentsSimilarIncidents) SetRelatedRouteRuleName(v string) *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidentsSimilarIncidents {
	s.RelatedRouteRuleName = &v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidentsSimilarIncidents) SetSimilarScore(v string) *GetSimilarIncidentStatisticsResponseBodyDataDailySimilarIncidentsSimilarIncidents {
	s.SimilarScore = &v
	return s
}

type GetSimilarIncidentStatisticsResponseBodyDataTopFiveIncidents struct {
	AssignUserId              *string `json:"assignUserId,omitempty" xml:"assignUserId,omitempty"`
	AssignUserName            *string `json:"assignUserName,omitempty" xml:"assignUserName,omitempty"`
	CreateTime                *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DurationTime              *int64  `json:"durationTime,omitempty" xml:"durationTime,omitempty"`
	FinishReason              *int64  `json:"finishReason,omitempty" xml:"finishReason,omitempty"`
	FinishReasonDescription   *string `json:"finishReasonDescription,omitempty" xml:"finishReasonDescription,omitempty"`
	FinishSolutionDescription *string `json:"finishSolutionDescription,omitempty" xml:"finishSolutionDescription,omitempty"`
	IncidentFinishSolution    *int64  `json:"incidentFinishSolution,omitempty" xml:"incidentFinishSolution,omitempty"`
	IncidentId                *int64  `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
	IncidentNumber            *string `json:"incidentNumber,omitempty" xml:"incidentNumber,omitempty"`
	IncidentTitle             *string `json:"incidentTitle,omitempty" xml:"incidentTitle,omitempty"`
	RelatedRouteRuleId        *int64  `json:"relatedRouteRuleId,omitempty" xml:"relatedRouteRuleId,omitempty"`
	RelatedRouteRuleName      *string `json:"relatedRouteRuleName,omitempty" xml:"relatedRouteRuleName,omitempty"`
	SimilarScore              *string `json:"similarScore,omitempty" xml:"similarScore,omitempty"`
}

func (s GetSimilarIncidentStatisticsResponseBodyDataTopFiveIncidents) String() string {
	return tea.Prettify(s)
}

func (s GetSimilarIncidentStatisticsResponseBodyDataTopFiveIncidents) GoString() string {
	return s.String()
}

func (s *GetSimilarIncidentStatisticsResponseBodyDataTopFiveIncidents) SetAssignUserId(v string) *GetSimilarIncidentStatisticsResponseBodyDataTopFiveIncidents {
	s.AssignUserId = &v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBodyDataTopFiveIncidents) SetAssignUserName(v string) *GetSimilarIncidentStatisticsResponseBodyDataTopFiveIncidents {
	s.AssignUserName = &v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBodyDataTopFiveIncidents) SetCreateTime(v string) *GetSimilarIncidentStatisticsResponseBodyDataTopFiveIncidents {
	s.CreateTime = &v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBodyDataTopFiveIncidents) SetDurationTime(v int64) *GetSimilarIncidentStatisticsResponseBodyDataTopFiveIncidents {
	s.DurationTime = &v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBodyDataTopFiveIncidents) SetFinishReason(v int64) *GetSimilarIncidentStatisticsResponseBodyDataTopFiveIncidents {
	s.FinishReason = &v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBodyDataTopFiveIncidents) SetFinishReasonDescription(v string) *GetSimilarIncidentStatisticsResponseBodyDataTopFiveIncidents {
	s.FinishReasonDescription = &v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBodyDataTopFiveIncidents) SetFinishSolutionDescription(v string) *GetSimilarIncidentStatisticsResponseBodyDataTopFiveIncidents {
	s.FinishSolutionDescription = &v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBodyDataTopFiveIncidents) SetIncidentFinishSolution(v int64) *GetSimilarIncidentStatisticsResponseBodyDataTopFiveIncidents {
	s.IncidentFinishSolution = &v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBodyDataTopFiveIncidents) SetIncidentId(v int64) *GetSimilarIncidentStatisticsResponseBodyDataTopFiveIncidents {
	s.IncidentId = &v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBodyDataTopFiveIncidents) SetIncidentNumber(v string) *GetSimilarIncidentStatisticsResponseBodyDataTopFiveIncidents {
	s.IncidentNumber = &v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBodyDataTopFiveIncidents) SetIncidentTitle(v string) *GetSimilarIncidentStatisticsResponseBodyDataTopFiveIncidents {
	s.IncidentTitle = &v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBodyDataTopFiveIncidents) SetRelatedRouteRuleId(v int64) *GetSimilarIncidentStatisticsResponseBodyDataTopFiveIncidents {
	s.RelatedRouteRuleId = &v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBodyDataTopFiveIncidents) SetRelatedRouteRuleName(v string) *GetSimilarIncidentStatisticsResponseBodyDataTopFiveIncidents {
	s.RelatedRouteRuleName = &v
	return s
}

func (s *GetSimilarIncidentStatisticsResponseBodyDataTopFiveIncidents) SetSimilarScore(v string) *GetSimilarIncidentStatisticsResponseBodyDataTopFiveIncidents {
	s.SimilarScore = &v
	return s
}

type GetSimilarIncidentStatisticsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSimilarIncidentStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSimilarIncidentStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSimilarIncidentStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetSimilarIncidentStatisticsResponse) SetHeaders(v map[string]*string) *GetSimilarIncidentStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetSimilarIncidentStatisticsResponse) SetStatusCode(v int32) *GetSimilarIncidentStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSimilarIncidentStatisticsResponse) SetBody(v *GetSimilarIncidentStatisticsResponseBody) *GetSimilarIncidentStatisticsResponse {
	s.Body = v
	return s
}

type GetSubscriptionRequest struct {
	NotFilterScopeObjectDeleted *bool  `json:"notFilterScopeObjectDeleted,omitempty" xml:"notFilterScopeObjectDeleted,omitempty"`
	SubscriptionId              *int64 `json:"subscriptionId,omitempty" xml:"subscriptionId,omitempty"`
}

func (s GetSubscriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *GetSubscriptionRequest) SetNotFilterScopeObjectDeleted(v bool) *GetSubscriptionRequest {
	s.NotFilterScopeObjectDeleted = &v
	return s
}

func (s *GetSubscriptionRequest) SetSubscriptionId(v int64) *GetSubscriptionRequest {
	s.SubscriptionId = &v
	return s
}

type GetSubscriptionResponseBody struct {
	// Object
	Data *GetSubscriptionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetSubscriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubscriptionResponseBody) SetData(v *GetSubscriptionResponseBodyData) *GetSubscriptionResponseBody {
	s.Data = v
	return s
}

func (s *GetSubscriptionResponseBody) SetRequestId(v string) *GetSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

type GetSubscriptionResponseBodyData struct {
	EndTime            *string                                              `json:"endTime,omitempty" xml:"endTime,omitempty"`
	ExpiredType        *string                                              `json:"expiredType,omitempty" xml:"expiredType,omitempty"`
	NotifyObjectList   []*GetSubscriptionResponseBodyDataNotifyObjectList   `json:"notifyObjectList,omitempty" xml:"notifyObjectList,omitempty" type:"Repeated"`
	NotifyObjectType   *string                                              `json:"notifyObjectType,omitempty" xml:"notifyObjectType,omitempty"`
	NotifyStrategyList []*GetSubscriptionResponseBodyDataNotifyStrategyList `json:"notifyStrategyList,omitempty" xml:"notifyStrategyList,omitempty" type:"Repeated"`
	Period             *string                                              `json:"period,omitempty" xml:"period,omitempty"`
	Scope              *string                                              `json:"scope,omitempty" xml:"scope,omitempty"`
	// Array
	ScopeObjectList   []*GetSubscriptionResponseBodyDataScopeObjectList `json:"scopeObjectList,omitempty" xml:"scopeObjectList,omitempty" type:"Repeated"`
	StartTime         *string                                           `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status            *string                                           `json:"status,omitempty" xml:"status,omitempty"`
	SubscriptionId    *int64                                            `json:"subscriptionId,omitempty" xml:"subscriptionId,omitempty"`
	SubscriptionTitle *string                                           `json:"subscriptionTitle,omitempty" xml:"subscriptionTitle,omitempty"`
}

func (s GetSubscriptionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSubscriptionResponseBodyData) SetEndTime(v string) *GetSubscriptionResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetSubscriptionResponseBodyData) SetExpiredType(v string) *GetSubscriptionResponseBodyData {
	s.ExpiredType = &v
	return s
}

func (s *GetSubscriptionResponseBodyData) SetNotifyObjectList(v []*GetSubscriptionResponseBodyDataNotifyObjectList) *GetSubscriptionResponseBodyData {
	s.NotifyObjectList = v
	return s
}

func (s *GetSubscriptionResponseBodyData) SetNotifyObjectType(v string) *GetSubscriptionResponseBodyData {
	s.NotifyObjectType = &v
	return s
}

func (s *GetSubscriptionResponseBodyData) SetNotifyStrategyList(v []*GetSubscriptionResponseBodyDataNotifyStrategyList) *GetSubscriptionResponseBodyData {
	s.NotifyStrategyList = v
	return s
}

func (s *GetSubscriptionResponseBodyData) SetPeriod(v string) *GetSubscriptionResponseBodyData {
	s.Period = &v
	return s
}

func (s *GetSubscriptionResponseBodyData) SetScope(v string) *GetSubscriptionResponseBodyData {
	s.Scope = &v
	return s
}

func (s *GetSubscriptionResponseBodyData) SetScopeObjectList(v []*GetSubscriptionResponseBodyDataScopeObjectList) *GetSubscriptionResponseBodyData {
	s.ScopeObjectList = v
	return s
}

func (s *GetSubscriptionResponseBodyData) SetStartTime(v string) *GetSubscriptionResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetSubscriptionResponseBodyData) SetStatus(v string) *GetSubscriptionResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetSubscriptionResponseBodyData) SetSubscriptionId(v int64) *GetSubscriptionResponseBodyData {
	s.SubscriptionId = &v
	return s
}

func (s *GetSubscriptionResponseBodyData) SetSubscriptionTitle(v string) *GetSubscriptionResponseBodyData {
	s.SubscriptionTitle = &v
	return s
}

type GetSubscriptionResponseBodyDataNotifyObjectList struct {
	Id               *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name             *string `json:"name,omitempty" xml:"name,omitempty"`
	NotifyObjectId   *int64  `json:"notifyObjectId,omitempty" xml:"notifyObjectId,omitempty"`
	NotifyObjectType *int64  `json:"notifyObjectType,omitempty" xml:"notifyObjectType,omitempty"`
}

func (s GetSubscriptionResponseBodyDataNotifyObjectList) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionResponseBodyDataNotifyObjectList) GoString() string {
	return s.String()
}

func (s *GetSubscriptionResponseBodyDataNotifyObjectList) SetId(v int64) *GetSubscriptionResponseBodyDataNotifyObjectList {
	s.Id = &v
	return s
}

func (s *GetSubscriptionResponseBodyDataNotifyObjectList) SetName(v string) *GetSubscriptionResponseBodyDataNotifyObjectList {
	s.Name = &v
	return s
}

func (s *GetSubscriptionResponseBodyDataNotifyObjectList) SetNotifyObjectId(v int64) *GetSubscriptionResponseBodyDataNotifyObjectList {
	s.NotifyObjectId = &v
	return s
}

func (s *GetSubscriptionResponseBodyDataNotifyObjectList) SetNotifyObjectType(v int64) *GetSubscriptionResponseBodyDataNotifyObjectList {
	s.NotifyObjectType = &v
	return s
}

type GetSubscriptionResponseBodyDataNotifyStrategyList struct {
	InstanceType *int64                                                         `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	Strategies   []*GetSubscriptionResponseBodyDataNotifyStrategyListStrategies `json:"strategies,omitempty" xml:"strategies,omitempty" type:"Repeated"`
}

func (s GetSubscriptionResponseBodyDataNotifyStrategyList) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionResponseBodyDataNotifyStrategyList) GoString() string {
	return s.String()
}

func (s *GetSubscriptionResponseBodyDataNotifyStrategyList) SetInstanceType(v int64) *GetSubscriptionResponseBodyDataNotifyStrategyList {
	s.InstanceType = &v
	return s
}

func (s *GetSubscriptionResponseBodyDataNotifyStrategyList) SetStrategies(v []*GetSubscriptionResponseBodyDataNotifyStrategyListStrategies) *GetSubscriptionResponseBodyDataNotifyStrategyList {
	s.Strategies = v
	return s
}

type GetSubscriptionResponseBodyDataNotifyStrategyListStrategies struct {
	Channels      *string                                                                   `json:"channels,omitempty" xml:"channels,omitempty"`
	Conditions    []*GetSubscriptionResponseBodyDataNotifyStrategyListStrategiesConditions  `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	Id            *int64                                                                    `json:"id,omitempty" xml:"id,omitempty"`
	PeriodChannel *GetSubscriptionResponseBodyDataNotifyStrategyListStrategiesPeriodChannel `json:"periodChannel,omitempty" xml:"periodChannel,omitempty" type:"Struct"`
}

func (s GetSubscriptionResponseBodyDataNotifyStrategyListStrategies) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionResponseBodyDataNotifyStrategyListStrategies) GoString() string {
	return s.String()
}

func (s *GetSubscriptionResponseBodyDataNotifyStrategyListStrategies) SetChannels(v string) *GetSubscriptionResponseBodyDataNotifyStrategyListStrategies {
	s.Channels = &v
	return s
}

func (s *GetSubscriptionResponseBodyDataNotifyStrategyListStrategies) SetConditions(v []*GetSubscriptionResponseBodyDataNotifyStrategyListStrategiesConditions) *GetSubscriptionResponseBodyDataNotifyStrategyListStrategies {
	s.Conditions = v
	return s
}

func (s *GetSubscriptionResponseBodyDataNotifyStrategyListStrategies) SetId(v int64) *GetSubscriptionResponseBodyDataNotifyStrategyListStrategies {
	s.Id = &v
	return s
}

func (s *GetSubscriptionResponseBodyDataNotifyStrategyListStrategies) SetPeriodChannel(v *GetSubscriptionResponseBodyDataNotifyStrategyListStrategiesPeriodChannel) *GetSubscriptionResponseBodyDataNotifyStrategyListStrategies {
	s.PeriodChannel = v
	return s
}

type GetSubscriptionResponseBodyDataNotifyStrategyListStrategiesConditions struct {
	Action            *string `json:"action,omitempty" xml:"action,omitempty"`
	Effection         *string `json:"effection,omitempty" xml:"effection,omitempty"`
	Level             *string `json:"level,omitempty" xml:"level,omitempty"`
	ProblemNotifyType *string `json:"problemNotifyType,omitempty" xml:"problemNotifyType,omitempty"`
}

func (s GetSubscriptionResponseBodyDataNotifyStrategyListStrategiesConditions) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionResponseBodyDataNotifyStrategyListStrategiesConditions) GoString() string {
	return s.String()
}

func (s *GetSubscriptionResponseBodyDataNotifyStrategyListStrategiesConditions) SetAction(v string) *GetSubscriptionResponseBodyDataNotifyStrategyListStrategiesConditions {
	s.Action = &v
	return s
}

func (s *GetSubscriptionResponseBodyDataNotifyStrategyListStrategiesConditions) SetEffection(v string) *GetSubscriptionResponseBodyDataNotifyStrategyListStrategiesConditions {
	s.Effection = &v
	return s
}

func (s *GetSubscriptionResponseBodyDataNotifyStrategyListStrategiesConditions) SetLevel(v string) *GetSubscriptionResponseBodyDataNotifyStrategyListStrategiesConditions {
	s.Level = &v
	return s
}

func (s *GetSubscriptionResponseBodyDataNotifyStrategyListStrategiesConditions) SetProblemNotifyType(v string) *GetSubscriptionResponseBodyDataNotifyStrategyListStrategiesConditions {
	s.ProblemNotifyType = &v
	return s
}

type GetSubscriptionResponseBodyDataNotifyStrategyListStrategiesPeriodChannel struct {
	NonWorkday *string `json:"nonWorkday,omitempty" xml:"nonWorkday,omitempty"`
	Workday    *string `json:"workday,omitempty" xml:"workday,omitempty"`
}

func (s GetSubscriptionResponseBodyDataNotifyStrategyListStrategiesPeriodChannel) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionResponseBodyDataNotifyStrategyListStrategiesPeriodChannel) GoString() string {
	return s.String()
}

func (s *GetSubscriptionResponseBodyDataNotifyStrategyListStrategiesPeriodChannel) SetNonWorkday(v string) *GetSubscriptionResponseBodyDataNotifyStrategyListStrategiesPeriodChannel {
	s.NonWorkday = &v
	return s
}

func (s *GetSubscriptionResponseBodyDataNotifyStrategyListStrategiesPeriodChannel) SetWorkday(v string) *GetSubscriptionResponseBodyDataNotifyStrategyListStrategiesPeriodChannel {
	s.Workday = &v
	return s
}

type GetSubscriptionResponseBodyDataScopeObjectList struct {
	Id            *int64  `json:"id,omitempty" xml:"id,omitempty"`
	IsValid       *int64  `json:"isValid,omitempty" xml:"isValid,omitempty"`
	Scope         *string `json:"scope,omitempty" xml:"scope,omitempty"`
	ScopeObject   *string `json:"scopeObject,omitempty" xml:"scopeObject,omitempty"`
	ScopeObjectId *int64  `json:"scopeObjectId,omitempty" xml:"scopeObjectId,omitempty"`
}

func (s GetSubscriptionResponseBodyDataScopeObjectList) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionResponseBodyDataScopeObjectList) GoString() string {
	return s.String()
}

func (s *GetSubscriptionResponseBodyDataScopeObjectList) SetId(v int64) *GetSubscriptionResponseBodyDataScopeObjectList {
	s.Id = &v
	return s
}

func (s *GetSubscriptionResponseBodyDataScopeObjectList) SetIsValid(v int64) *GetSubscriptionResponseBodyDataScopeObjectList {
	s.IsValid = &v
	return s
}

func (s *GetSubscriptionResponseBodyDataScopeObjectList) SetScope(v string) *GetSubscriptionResponseBodyDataScopeObjectList {
	s.Scope = &v
	return s
}

func (s *GetSubscriptionResponseBodyDataScopeObjectList) SetScopeObject(v string) *GetSubscriptionResponseBodyDataScopeObjectList {
	s.ScopeObject = &v
	return s
}

func (s *GetSubscriptionResponseBodyDataScopeObjectList) SetScopeObjectId(v int64) *GetSubscriptionResponseBodyDataScopeObjectList {
	s.ScopeObjectId = &v
	return s
}

type GetSubscriptionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSubscriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *GetSubscriptionResponse) SetHeaders(v map[string]*string) *GetSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *GetSubscriptionResponse) SetStatusCode(v int32) *GetSubscriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSubscriptionResponse) SetBody(v *GetSubscriptionResponseBody) *GetSubscriptionResponse {
	s.Body = v
	return s
}

type GetTenantApplicationRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s GetTenantApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTenantApplicationRequest) GoString() string {
	return s.String()
}

func (s *GetTenantApplicationRequest) SetClientToken(v string) *GetTenantApplicationRequest {
	s.ClientToken = &v
	return s
}

type GetTenantApplicationResponseBody struct {
	// data
	Data *GetTenantApplicationResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetTenantApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTenantApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetTenantApplicationResponseBody) SetData(v *GetTenantApplicationResponseBodyData) *GetTenantApplicationResponseBody {
	s.Data = v
	return s
}

func (s *GetTenantApplicationResponseBody) SetRequestId(v string) *GetTenantApplicationResponseBody {
	s.RequestId = &v
	return s
}

type GetTenantApplicationResponseBodyData struct {
	BizId          *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	Channel        *string `json:"channel,omitempty" xml:"channel,omitempty"`
	CorporationId  *string `json:"corporationId,omitempty" xml:"corporationId,omitempty"`
	OriginalCorpId *string `json:"originalCorpId,omitempty" xml:"originalCorpId,omitempty"`
	Progress       *string `json:"progress,omitempty" xml:"progress,omitempty"`
}

func (s GetTenantApplicationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTenantApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTenantApplicationResponseBodyData) SetBizId(v string) *GetTenantApplicationResponseBodyData {
	s.BizId = &v
	return s
}

func (s *GetTenantApplicationResponseBodyData) SetChannel(v string) *GetTenantApplicationResponseBodyData {
	s.Channel = &v
	return s
}

func (s *GetTenantApplicationResponseBodyData) SetCorporationId(v string) *GetTenantApplicationResponseBodyData {
	s.CorporationId = &v
	return s
}

func (s *GetTenantApplicationResponseBodyData) SetOriginalCorpId(v string) *GetTenantApplicationResponseBodyData {
	s.OriginalCorpId = &v
	return s
}

func (s *GetTenantApplicationResponseBodyData) SetProgress(v string) *GetTenantApplicationResponseBodyData {
	s.Progress = &v
	return s
}

type GetTenantApplicationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTenantApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTenantApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTenantApplicationResponse) GoString() string {
	return s.String()
}

func (s *GetTenantApplicationResponse) SetHeaders(v map[string]*string) *GetTenantApplicationResponse {
	s.Headers = v
	return s
}

func (s *GetTenantApplicationResponse) SetStatusCode(v int32) *GetTenantApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTenantApplicationResponse) SetBody(v *GetTenantApplicationResponseBody) *GetTenantApplicationResponse {
	s.Body = v
	return s
}

type GetTenantStatusRequest struct {
	TenantRamId *int64 `json:"tenantRamId,omitempty" xml:"tenantRamId,omitempty"`
}

func (s GetTenantStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTenantStatusRequest) GoString() string {
	return s.String()
}

func (s *GetTenantStatusRequest) SetTenantRamId(v int64) *GetTenantStatusRequest {
	s.TenantRamId = &v
	return s
}

type GetTenantStatusResponseBody struct {
	Data      *GetTenantStatusResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	RequestId *string                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetTenantStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTenantStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetTenantStatusResponseBody) SetData(v *GetTenantStatusResponseBodyData) *GetTenantStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetTenantStatusResponseBody) SetRequestId(v string) *GetTenantStatusResponseBody {
	s.RequestId = &v
	return s
}

type GetTenantStatusResponseBodyData struct {
	TenantStatus *int32 `json:"tenantStatus,omitempty" xml:"tenantStatus,omitempty"`
}

func (s GetTenantStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTenantStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTenantStatusResponseBodyData) SetTenantStatus(v int32) *GetTenantStatusResponseBodyData {
	s.TenantStatus = &v
	return s
}

type GetTenantStatusResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTenantStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTenantStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTenantStatusResponse) GoString() string {
	return s.String()
}

func (s *GetTenantStatusResponse) SetHeaders(v map[string]*string) *GetTenantStatusResponse {
	s.Headers = v
	return s
}

func (s *GetTenantStatusResponse) SetStatusCode(v int32) *GetTenantStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTenantStatusResponse) SetBody(v *GetTenantStatusResponseBody) *GetTenantStatusResponse {
	s.Body = v
	return s
}

type GetUserRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	UserId      *int64  `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetUserRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) SetClientToken(v string) *GetUserRequest {
	s.ClientToken = &v
	return s
}

func (s *GetUserRequest) SetUserId(v int64) *GetUserRequest {
	s.UserId = &v
	return s
}

type GetUserResponseBody struct {
	Data *GetUserResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) SetData(v *GetUserResponseBodyData) *GetUserResponseBody {
	s.Data = v
	return s
}

func (s *GetUserResponseBody) SetRequestId(v string) *GetUserResponseBody {
	s.RequestId = &v
	return s
}

type GetUserResponseBodyData struct {
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	CreateTime  *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// email
	Email          *string `json:"email,omitempty" xml:"email,omitempty"`
	IsEditableUser *bool   `json:"isEditableUser,omitempty" xml:"isEditableUser,omitempty"`
	IsRelated      *string `json:"isRelated,omitempty" xml:"isRelated,omitempty"`
	Phone          *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// ramId
	RamId         *string                                 `json:"ramId,omitempty" xml:"ramId,omitempty"`
	ServiceGroups []*GetUserResponseBodyDataServiceGroups `json:"serviceGroups,omitempty" xml:"serviceGroups,omitempty" type:"Repeated"`
	UserId        *int64                                  `json:"userId,omitempty" xml:"userId,omitempty"`
	Username      *string                                 `json:"username,omitempty" xml:"username,omitempty"`
}

func (s GetUserResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyData) SetAccountType(v string) *GetUserResponseBodyData {
	s.AccountType = &v
	return s
}

func (s *GetUserResponseBodyData) SetCreateTime(v string) *GetUserResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetUserResponseBodyData) SetEmail(v string) *GetUserResponseBodyData {
	s.Email = &v
	return s
}

func (s *GetUserResponseBodyData) SetIsEditableUser(v bool) *GetUserResponseBodyData {
	s.IsEditableUser = &v
	return s
}

func (s *GetUserResponseBodyData) SetIsRelated(v string) *GetUserResponseBodyData {
	s.IsRelated = &v
	return s
}

func (s *GetUserResponseBodyData) SetPhone(v string) *GetUserResponseBodyData {
	s.Phone = &v
	return s
}

func (s *GetUserResponseBodyData) SetRamId(v string) *GetUserResponseBodyData {
	s.RamId = &v
	return s
}

func (s *GetUserResponseBodyData) SetServiceGroups(v []*GetUserResponseBodyDataServiceGroups) *GetUserResponseBodyData {
	s.ServiceGroups = v
	return s
}

func (s *GetUserResponseBodyData) SetUserId(v int64) *GetUserResponseBodyData {
	s.UserId = &v
	return s
}

func (s *GetUserResponseBodyData) SetUsername(v string) *GetUserResponseBodyData {
	s.Username = &v
	return s
}

type GetUserResponseBodyDataServiceGroups struct {
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	ServiceGroupId *int64  `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
}

func (s GetUserResponseBodyDataServiceGroups) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyDataServiceGroups) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyDataServiceGroups) SetName(v string) *GetUserResponseBodyDataServiceGroups {
	s.Name = &v
	return s
}

func (s *GetUserResponseBodyDataServiceGroups) SetServiceGroupId(v int64) *GetUserResponseBodyDataServiceGroups {
	s.ServiceGroupId = &v
	return s
}

type GetUserResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponse) GoString() string {
	return s.String()
}

func (s *GetUserResponse) SetHeaders(v map[string]*string) *GetUserResponse {
	s.Headers = v
	return s
}

func (s *GetUserResponse) SetStatusCode(v int32) *GetUserResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserResponse) SetBody(v *GetUserResponseBody) *GetUserResponse {
	s.Body = v
	return s
}

type GetUserGuideStatusRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s GetUserGuideStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserGuideStatusRequest) GoString() string {
	return s.String()
}

func (s *GetUserGuideStatusRequest) SetClientToken(v string) *GetUserGuideStatusRequest {
	s.ClientToken = &v
	return s
}

type GetUserGuideStatusResponseBody struct {
	// map
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetUserGuideStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserGuideStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserGuideStatusResponseBody) SetData(v map[string]interface{}) *GetUserGuideStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetUserGuideStatusResponseBody) SetRequestId(v string) *GetUserGuideStatusResponseBody {
	s.RequestId = &v
	return s
}

type GetUserGuideStatusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserGuideStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserGuideStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserGuideStatusResponse) GoString() string {
	return s.String()
}

func (s *GetUserGuideStatusResponse) SetHeaders(v map[string]*string) *GetUserGuideStatusResponse {
	s.Headers = v
	return s
}

func (s *GetUserGuideStatusResponse) SetStatusCode(v int32) *GetUserGuideStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserGuideStatusResponse) SetBody(v *GetUserGuideStatusResponseBody) *GetUserGuideStatusResponse {
	s.Body = v
	return s
}

type ListAlertsRequest struct {
	AlertLevel      *string `json:"alertLevel,omitempty" xml:"alertLevel,omitempty"`
	AlertName       *string `json:"alertName,omitempty" xml:"alertName,omitempty"`
	AlertSourceName *string `json:"alertSourceName,omitempty" xml:"alertSourceName,omitempty"`
	// 2020-09-10 21:00:00
	EndTime          *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	MonitorSourceId  *string `json:"monitorSourceId,omitempty" xml:"monitorSourceId,omitempty"`
	PageNumber       *int64  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize         *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	RelatedServiceId *int64  `json:"relatedServiceId,omitempty" xml:"relatedServiceId,omitempty"`
	RuleName         *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// 2020-09-10 13:00:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListAlertsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAlertsRequest) GoString() string {
	return s.String()
}

func (s *ListAlertsRequest) SetAlertLevel(v string) *ListAlertsRequest {
	s.AlertLevel = &v
	return s
}

func (s *ListAlertsRequest) SetAlertName(v string) *ListAlertsRequest {
	s.AlertName = &v
	return s
}

func (s *ListAlertsRequest) SetAlertSourceName(v string) *ListAlertsRequest {
	s.AlertSourceName = &v
	return s
}

func (s *ListAlertsRequest) SetEndTime(v string) *ListAlertsRequest {
	s.EndTime = &v
	return s
}

func (s *ListAlertsRequest) SetMonitorSourceId(v string) *ListAlertsRequest {
	s.MonitorSourceId = &v
	return s
}

func (s *ListAlertsRequest) SetPageNumber(v int64) *ListAlertsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAlertsRequest) SetPageSize(v int64) *ListAlertsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAlertsRequest) SetRelatedServiceId(v int64) *ListAlertsRequest {
	s.RelatedServiceId = &v
	return s
}

func (s *ListAlertsRequest) SetRuleName(v string) *ListAlertsRequest {
	s.RuleName = &v
	return s
}

func (s *ListAlertsRequest) SetStartTime(v string) *ListAlertsRequest {
	s.StartTime = &v
	return s
}

type ListAlertsResponseBody struct {
	Data       []*ListAlertsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	PageNumber *int32                        `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32                        `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	RequestId  *string                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TotalCount *int64                        `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListAlertsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAlertsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlertsResponseBody) SetData(v []*ListAlertsResponseBodyData) *ListAlertsResponseBody {
	s.Data = v
	return s
}

func (s *ListAlertsResponseBody) SetPageNumber(v int32) *ListAlertsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAlertsResponseBody) SetPageSize(v int32) *ListAlertsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAlertsResponseBody) SetRequestId(v string) *ListAlertsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlertsResponseBody) SetTotalCount(v int64) *ListAlertsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAlertsResponseBodyData struct {
	AlertId              *int64  `json:"alertId,omitempty" xml:"alertId,omitempty"`
	AlertLevel           *string `json:"alertLevel,omitempty" xml:"alertLevel,omitempty"`
	AlertNumber          *string `json:"alertNumber,omitempty" xml:"alertNumber,omitempty"`
	AlertSourceName      *string `json:"alertSourceName,omitempty" xml:"alertSourceName,omitempty"`
	CreateTime           *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	FirstEventTime       *string `json:"firstEventTime,omitempty" xml:"firstEventTime,omitempty"`
	MonitorSourceName    *string `json:"monitorSourceName,omitempty" xml:"monitorSourceName,omitempty"`
	RelServiceDeleteType *int32  `json:"relServiceDeleteType,omitempty" xml:"relServiceDeleteType,omitempty"`
	RelatedServiceName   *string `json:"relatedServiceName,omitempty" xml:"relatedServiceName,omitempty"`
	RouteRuleDeleteType  *int32  `json:"routeRuleDeleteType,omitempty" xml:"routeRuleDeleteType,omitempty"`
	RouteRuleId          *int64  `json:"routeRuleId,omitempty" xml:"routeRuleId,omitempty"`
	RouteRuleName        *string `json:"routeRuleName,omitempty" xml:"routeRuleName,omitempty"`
	SourceEventCount     *int64  `json:"sourceEventCount,omitempty" xml:"sourceEventCount,omitempty"`
	Title                *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ListAlertsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAlertsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAlertsResponseBodyData) SetAlertId(v int64) *ListAlertsResponseBodyData {
	s.AlertId = &v
	return s
}

func (s *ListAlertsResponseBodyData) SetAlertLevel(v string) *ListAlertsResponseBodyData {
	s.AlertLevel = &v
	return s
}

func (s *ListAlertsResponseBodyData) SetAlertNumber(v string) *ListAlertsResponseBodyData {
	s.AlertNumber = &v
	return s
}

func (s *ListAlertsResponseBodyData) SetAlertSourceName(v string) *ListAlertsResponseBodyData {
	s.AlertSourceName = &v
	return s
}

func (s *ListAlertsResponseBodyData) SetCreateTime(v string) *ListAlertsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListAlertsResponseBodyData) SetFirstEventTime(v string) *ListAlertsResponseBodyData {
	s.FirstEventTime = &v
	return s
}

func (s *ListAlertsResponseBodyData) SetMonitorSourceName(v string) *ListAlertsResponseBodyData {
	s.MonitorSourceName = &v
	return s
}

func (s *ListAlertsResponseBodyData) SetRelServiceDeleteType(v int32) *ListAlertsResponseBodyData {
	s.RelServiceDeleteType = &v
	return s
}

func (s *ListAlertsResponseBodyData) SetRelatedServiceName(v string) *ListAlertsResponseBodyData {
	s.RelatedServiceName = &v
	return s
}

func (s *ListAlertsResponseBodyData) SetRouteRuleDeleteType(v int32) *ListAlertsResponseBodyData {
	s.RouteRuleDeleteType = &v
	return s
}

func (s *ListAlertsResponseBodyData) SetRouteRuleId(v int64) *ListAlertsResponseBodyData {
	s.RouteRuleId = &v
	return s
}

func (s *ListAlertsResponseBodyData) SetRouteRuleName(v string) *ListAlertsResponseBodyData {
	s.RouteRuleName = &v
	return s
}

func (s *ListAlertsResponseBodyData) SetSourceEventCount(v int64) *ListAlertsResponseBodyData {
	s.SourceEventCount = &v
	return s
}

func (s *ListAlertsResponseBodyData) SetTitle(v string) *ListAlertsResponseBodyData {
	s.Title = &v
	return s
}

type ListAlertsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAlertsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAlertsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAlertsResponse) GoString() string {
	return s.String()
}

func (s *ListAlertsResponse) SetHeaders(v map[string]*string) *ListAlertsResponse {
	s.Headers = v
	return s
}

func (s *ListAlertsResponse) SetStatusCode(v int32) *ListAlertsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlertsResponse) SetBody(v *ListAlertsResponseBody) *ListAlertsResponse {
	s.Body = v
	return s
}

type ListByMonitorSourceIdRequest struct {
	MonitorSourceId *string `json:"monitorSourceId,omitempty" xml:"monitorSourceId,omitempty"`
}

func (s ListByMonitorSourceIdRequest) String() string {
	return tea.Prettify(s)
}

func (s ListByMonitorSourceIdRequest) GoString() string {
	return s.String()
}

func (s *ListByMonitorSourceIdRequest) SetMonitorSourceId(v string) *ListByMonitorSourceIdRequest {
	s.MonitorSourceId = &v
	return s
}

type ListByMonitorSourceIdResponseBody struct {
	Data      []*ListByMonitorSourceIdResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListByMonitorSourceIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListByMonitorSourceIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListByMonitorSourceIdResponseBody) SetData(v []*ListByMonitorSourceIdResponseBodyData) *ListByMonitorSourceIdResponseBody {
	s.Data = v
	return s
}

func (s *ListByMonitorSourceIdResponseBody) SetRequestId(v string) *ListByMonitorSourceIdResponseBody {
	s.RequestId = &v
	return s
}

type ListByMonitorSourceIdResponseBodyData struct {
	Id       *int64  `json:"id,omitempty" xml:"id,omitempty"`
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
}

func (s ListByMonitorSourceIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListByMonitorSourceIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListByMonitorSourceIdResponseBodyData) SetId(v int64) *ListByMonitorSourceIdResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListByMonitorSourceIdResponseBodyData) SetRuleName(v string) *ListByMonitorSourceIdResponseBodyData {
	s.RuleName = &v
	return s
}

type ListByMonitorSourceIdResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListByMonitorSourceIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListByMonitorSourceIdResponse) String() string {
	return tea.Prettify(s)
}

func (s ListByMonitorSourceIdResponse) GoString() string {
	return s.String()
}

func (s *ListByMonitorSourceIdResponse) SetHeaders(v map[string]*string) *ListByMonitorSourceIdResponse {
	s.Headers = v
	return s
}

func (s *ListByMonitorSourceIdResponse) SetStatusCode(v int32) *ListByMonitorSourceIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListByMonitorSourceIdResponse) SetBody(v *ListByMonitorSourceIdResponseBody) *ListByMonitorSourceIdResponse {
	s.Body = v
	return s
}

type ListChartDataForServiceGroupRequest struct {
	// clientToken
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	EndTime     *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	StartTime   *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListChartDataForServiceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListChartDataForServiceGroupRequest) GoString() string {
	return s.String()
}

func (s *ListChartDataForServiceGroupRequest) SetClientToken(v string) *ListChartDataForServiceGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *ListChartDataForServiceGroupRequest) SetEndTime(v string) *ListChartDataForServiceGroupRequest {
	s.EndTime = &v
	return s
}

func (s *ListChartDataForServiceGroupRequest) SetStartTime(v string) *ListChartDataForServiceGroupRequest {
	s.StartTime = &v
	return s
}

type ListChartDataForServiceGroupResponseBody struct {
	// data
	Data []*ListChartDataForServiceGroupResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListChartDataForServiceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListChartDataForServiceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListChartDataForServiceGroupResponseBody) SetData(v []*ListChartDataForServiceGroupResponseBodyData) *ListChartDataForServiceGroupResponseBody {
	s.Data = v
	return s
}

func (s *ListChartDataForServiceGroupResponseBody) SetRequestId(v string) *ListChartDataForServiceGroupResponseBody {
	s.RequestId = &v
	return s
}

type ListChartDataForServiceGroupResponseBodyData struct {
	EffectionLevel                        map[string]interface{} `json:"effectionLevel,omitempty" xml:"effectionLevel,omitempty"`
	EscalationIncidentCount               *int64                 `json:"escalationIncidentCount,omitempty" xml:"escalationIncidentCount,omitempty"`
	IncidentCount                         *int64                 `json:"incidentCount,omitempty" xml:"incidentCount,omitempty"`
	MeanTimeToAcknowledge                 *int64                 `json:"meanTimeToAcknowledge,omitempty" xml:"meanTimeToAcknowledge,omitempty"`
	MeanTimeToRepair                      *int64                 `json:"meanTimeToRepair,omitempty" xml:"meanTimeToRepair,omitempty"`
	Time                                  *string                `json:"time,omitempty" xml:"time,omitempty"`
	TotalMeanTimeToAcknowledge            *int64                 `json:"totalMeanTimeToAcknowledge,omitempty" xml:"totalMeanTimeToAcknowledge,omitempty"`
	TotalMeanTimeToRepair                 *int64                 `json:"totalMeanTimeToRepair,omitempty" xml:"totalMeanTimeToRepair,omitempty"`
	UnAcknowledgedEscalationIncidentCount *int64                 `json:"unAcknowledgedEscalationIncidentCount,omitempty" xml:"unAcknowledgedEscalationIncidentCount,omitempty"`
	UnFinishEscalationIncidentCount       *int64                 `json:"unFinishEscalationIncidentCount,omitempty" xml:"unFinishEscalationIncidentCount,omitempty"`
}

func (s ListChartDataForServiceGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListChartDataForServiceGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListChartDataForServiceGroupResponseBodyData) SetEffectionLevel(v map[string]interface{}) *ListChartDataForServiceGroupResponseBodyData {
	s.EffectionLevel = v
	return s
}

func (s *ListChartDataForServiceGroupResponseBodyData) SetEscalationIncidentCount(v int64) *ListChartDataForServiceGroupResponseBodyData {
	s.EscalationIncidentCount = &v
	return s
}

func (s *ListChartDataForServiceGroupResponseBodyData) SetIncidentCount(v int64) *ListChartDataForServiceGroupResponseBodyData {
	s.IncidentCount = &v
	return s
}

func (s *ListChartDataForServiceGroupResponseBodyData) SetMeanTimeToAcknowledge(v int64) *ListChartDataForServiceGroupResponseBodyData {
	s.MeanTimeToAcknowledge = &v
	return s
}

func (s *ListChartDataForServiceGroupResponseBodyData) SetMeanTimeToRepair(v int64) *ListChartDataForServiceGroupResponseBodyData {
	s.MeanTimeToRepair = &v
	return s
}

func (s *ListChartDataForServiceGroupResponseBodyData) SetTime(v string) *ListChartDataForServiceGroupResponseBodyData {
	s.Time = &v
	return s
}

func (s *ListChartDataForServiceGroupResponseBodyData) SetTotalMeanTimeToAcknowledge(v int64) *ListChartDataForServiceGroupResponseBodyData {
	s.TotalMeanTimeToAcknowledge = &v
	return s
}

func (s *ListChartDataForServiceGroupResponseBodyData) SetTotalMeanTimeToRepair(v int64) *ListChartDataForServiceGroupResponseBodyData {
	s.TotalMeanTimeToRepair = &v
	return s
}

func (s *ListChartDataForServiceGroupResponseBodyData) SetUnAcknowledgedEscalationIncidentCount(v int64) *ListChartDataForServiceGroupResponseBodyData {
	s.UnAcknowledgedEscalationIncidentCount = &v
	return s
}

func (s *ListChartDataForServiceGroupResponseBodyData) SetUnFinishEscalationIncidentCount(v int64) *ListChartDataForServiceGroupResponseBodyData {
	s.UnFinishEscalationIncidentCount = &v
	return s
}

type ListChartDataForServiceGroupResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListChartDataForServiceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListChartDataForServiceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListChartDataForServiceGroupResponse) GoString() string {
	return s.String()
}

func (s *ListChartDataForServiceGroupResponse) SetHeaders(v map[string]*string) *ListChartDataForServiceGroupResponse {
	s.Headers = v
	return s
}

func (s *ListChartDataForServiceGroupResponse) SetStatusCode(v int32) *ListChartDataForServiceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChartDataForServiceGroupResponse) SetBody(v *ListChartDataForServiceGroupResponseBody) *ListChartDataForServiceGroupResponse {
	s.Body = v
	return s
}

type ListChartDataForUserRequest struct {
	// clientToken
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	EndTime     *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	StartTime   *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListChartDataForUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ListChartDataForUserRequest) GoString() string {
	return s.String()
}

func (s *ListChartDataForUserRequest) SetClientToken(v string) *ListChartDataForUserRequest {
	s.ClientToken = &v
	return s
}

func (s *ListChartDataForUserRequest) SetEndTime(v string) *ListChartDataForUserRequest {
	s.EndTime = &v
	return s
}

func (s *ListChartDataForUserRequest) SetStartTime(v string) *ListChartDataForUserRequest {
	s.StartTime = &v
	return s
}

type ListChartDataForUserResponseBody struct {
	// data
	Data []*ListChartDataForUserResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListChartDataForUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListChartDataForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListChartDataForUserResponseBody) SetData(v []*ListChartDataForUserResponseBodyData) *ListChartDataForUserResponseBody {
	s.Data = v
	return s
}

func (s *ListChartDataForUserResponseBody) SetRequestId(v string) *ListChartDataForUserResponseBody {
	s.RequestId = &v
	return s
}

type ListChartDataForUserResponseBodyData struct {
	EffectionLevel                        map[string]interface{} `json:"effectionLevel,omitempty" xml:"effectionLevel,omitempty"`
	EscalationIncidentCount               *int64                 `json:"escalationIncidentCount,omitempty" xml:"escalationIncidentCount,omitempty"`
	IncidentCount                         *int64                 `json:"incidentCount,omitempty" xml:"incidentCount,omitempty"`
	MeanTimeToAcknowledge                 *int64                 `json:"meanTimeToAcknowledge,omitempty" xml:"meanTimeToAcknowledge,omitempty"`
	MeanTimeToRepair                      *int64                 `json:"meanTimeToRepair,omitempty" xml:"meanTimeToRepair,omitempty"`
	Time                                  *string                `json:"time,omitempty" xml:"time,omitempty"`
	TotalMeanTimeToAcknowledge            *int64                 `json:"totalMeanTimeToAcknowledge,omitempty" xml:"totalMeanTimeToAcknowledge,omitempty"`
	TotalMeanTimeToRepair                 *int64                 `json:"totalMeanTimeToRepair,omitempty" xml:"totalMeanTimeToRepair,omitempty"`
	UnAcknowledgedEscalationIncidentCount *int64                 `json:"unAcknowledgedEscalationIncidentCount,omitempty" xml:"unAcknowledgedEscalationIncidentCount,omitempty"`
	UnFinishEscalationIncidentCount       *int64                 `json:"unFinishEscalationIncidentCount,omitempty" xml:"unFinishEscalationIncidentCount,omitempty"`
}

func (s ListChartDataForUserResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListChartDataForUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListChartDataForUserResponseBodyData) SetEffectionLevel(v map[string]interface{}) *ListChartDataForUserResponseBodyData {
	s.EffectionLevel = v
	return s
}

func (s *ListChartDataForUserResponseBodyData) SetEscalationIncidentCount(v int64) *ListChartDataForUserResponseBodyData {
	s.EscalationIncidentCount = &v
	return s
}

func (s *ListChartDataForUserResponseBodyData) SetIncidentCount(v int64) *ListChartDataForUserResponseBodyData {
	s.IncidentCount = &v
	return s
}

func (s *ListChartDataForUserResponseBodyData) SetMeanTimeToAcknowledge(v int64) *ListChartDataForUserResponseBodyData {
	s.MeanTimeToAcknowledge = &v
	return s
}

func (s *ListChartDataForUserResponseBodyData) SetMeanTimeToRepair(v int64) *ListChartDataForUserResponseBodyData {
	s.MeanTimeToRepair = &v
	return s
}

func (s *ListChartDataForUserResponseBodyData) SetTime(v string) *ListChartDataForUserResponseBodyData {
	s.Time = &v
	return s
}

func (s *ListChartDataForUserResponseBodyData) SetTotalMeanTimeToAcknowledge(v int64) *ListChartDataForUserResponseBodyData {
	s.TotalMeanTimeToAcknowledge = &v
	return s
}

func (s *ListChartDataForUserResponseBodyData) SetTotalMeanTimeToRepair(v int64) *ListChartDataForUserResponseBodyData {
	s.TotalMeanTimeToRepair = &v
	return s
}

func (s *ListChartDataForUserResponseBodyData) SetUnAcknowledgedEscalationIncidentCount(v int64) *ListChartDataForUserResponseBodyData {
	s.UnAcknowledgedEscalationIncidentCount = &v
	return s
}

func (s *ListChartDataForUserResponseBodyData) SetUnFinishEscalationIncidentCount(v int64) *ListChartDataForUserResponseBodyData {
	s.UnFinishEscalationIncidentCount = &v
	return s
}

type ListChartDataForUserResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListChartDataForUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListChartDataForUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ListChartDataForUserResponse) GoString() string {
	return s.String()
}

func (s *ListChartDataForUserResponse) SetHeaders(v map[string]*string) *ListChartDataForUserResponse {
	s.Headers = v
	return s
}

func (s *ListChartDataForUserResponse) SetStatusCode(v int32) *ListChartDataForUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChartDataForUserResponse) SetBody(v *ListChartDataForUserResponseBody) *ListChartDataForUserResponse {
	s.Body = v
	return s
}

type ListConfigsRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s ListConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListConfigsRequest) SetClientToken(v string) *ListConfigsRequest {
	s.ClientToken = &v
	return s
}

type ListConfigsResponseBody struct {
	// data
	Data map[string][]*DataValue `json:"data,omitempty" xml:"data,omitempty"`
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConfigsResponseBody) SetData(v map[string][]*DataValue) *ListConfigsResponseBody {
	s.Data = v
	return s
}

func (s *ListConfigsResponseBody) SetRequestId(v string) *ListConfigsResponseBody {
	s.RequestId = &v
	return s
}

type ListConfigsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListConfigsResponse) SetHeaders(v map[string]*string) *ListConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListConfigsResponse) SetStatusCode(v int32) *ListConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConfigsResponse) SetBody(v *ListConfigsResponseBody) *ListConfigsResponse {
	s.Body = v
	return s
}

type ListDataReportForServiceGroupRequest struct {
	EndTime          *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	ServiceGroupName *string `json:"serviceGroupName,omitempty" xml:"serviceGroupName,omitempty"`
	StartTime        *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListDataReportForServiceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataReportForServiceGroupRequest) GoString() string {
	return s.String()
}

func (s *ListDataReportForServiceGroupRequest) SetEndTime(v string) *ListDataReportForServiceGroupRequest {
	s.EndTime = &v
	return s
}

func (s *ListDataReportForServiceGroupRequest) SetServiceGroupName(v string) *ListDataReportForServiceGroupRequest {
	s.ServiceGroupName = &v
	return s
}

func (s *ListDataReportForServiceGroupRequest) SetStartTime(v string) *ListDataReportForServiceGroupRequest {
	s.StartTime = &v
	return s
}

type ListDataReportForServiceGroupResponseBody struct {
	Data       []*ListDataReportForServiceGroupResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	PageNumber *int64                                           `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSIze   *int64                                           `json:"pageSIze,omitempty" xml:"pageSIze,omitempty"`
	// Id of the request
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TotalCount *int64  `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListDataReportForServiceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataReportForServiceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataReportForServiceGroupResponseBody) SetData(v []*ListDataReportForServiceGroupResponseBodyData) *ListDataReportForServiceGroupResponseBody {
	s.Data = v
	return s
}

func (s *ListDataReportForServiceGroupResponseBody) SetPageNumber(v int64) *ListDataReportForServiceGroupResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDataReportForServiceGroupResponseBody) SetPageSIze(v int64) *ListDataReportForServiceGroupResponseBody {
	s.PageSIze = &v
	return s
}

func (s *ListDataReportForServiceGroupResponseBody) SetRequestId(v string) *ListDataReportForServiceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataReportForServiceGroupResponseBody) SetTotalCount(v int64) *ListDataReportForServiceGroupResponseBody {
	s.TotalCount = &v
	return s
}

type ListDataReportForServiceGroupResponseBodyData struct {
	EscalationIncidentCount *int64  `json:"escalationIncidentCount,omitempty" xml:"escalationIncidentCount,omitempty"`
	FinishIncidentCount     *int64  `json:"finishIncidentCount,omitempty" xml:"finishIncidentCount,omitempty"`
	FinishProportion        *string `json:"finishProportion,omitempty" xml:"finishProportion,omitempty"`
	IncidentCount           *int64  `json:"incidentCount,omitempty" xml:"incidentCount,omitempty"`
	// MRRA
	MeanTimeToAcknowledge *int64 `json:"meanTimeToAcknowledge,omitempty" xml:"meanTimeToAcknowledge,omitempty"`
	// MTTR
	MeanTimeToRepair                      *int64  `json:"meanTimeToRepair,omitempty" xml:"meanTimeToRepair,omitempty"`
	ServiceGroupId                        *int64  `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
	ServiceGroupName                      *string `json:"serviceGroupName,omitempty" xml:"serviceGroupName,omitempty"`
	UnAcknowledgedEscalationIncidentCount *int64  `json:"unAcknowledgedEscalationIncidentCount,omitempty" xml:"unAcknowledgedEscalationIncidentCount,omitempty"`
	UnFinishEscalationIncidentCount       *int64  `json:"unFinishEscalationIncidentCount,omitempty" xml:"unFinishEscalationIncidentCount,omitempty"`
}

func (s ListDataReportForServiceGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDataReportForServiceGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataReportForServiceGroupResponseBodyData) SetEscalationIncidentCount(v int64) *ListDataReportForServiceGroupResponseBodyData {
	s.EscalationIncidentCount = &v
	return s
}

func (s *ListDataReportForServiceGroupResponseBodyData) SetFinishIncidentCount(v int64) *ListDataReportForServiceGroupResponseBodyData {
	s.FinishIncidentCount = &v
	return s
}

func (s *ListDataReportForServiceGroupResponseBodyData) SetFinishProportion(v string) *ListDataReportForServiceGroupResponseBodyData {
	s.FinishProportion = &v
	return s
}

func (s *ListDataReportForServiceGroupResponseBodyData) SetIncidentCount(v int64) *ListDataReportForServiceGroupResponseBodyData {
	s.IncidentCount = &v
	return s
}

func (s *ListDataReportForServiceGroupResponseBodyData) SetMeanTimeToAcknowledge(v int64) *ListDataReportForServiceGroupResponseBodyData {
	s.MeanTimeToAcknowledge = &v
	return s
}

func (s *ListDataReportForServiceGroupResponseBodyData) SetMeanTimeToRepair(v int64) *ListDataReportForServiceGroupResponseBodyData {
	s.MeanTimeToRepair = &v
	return s
}

func (s *ListDataReportForServiceGroupResponseBodyData) SetServiceGroupId(v int64) *ListDataReportForServiceGroupResponseBodyData {
	s.ServiceGroupId = &v
	return s
}

func (s *ListDataReportForServiceGroupResponseBodyData) SetServiceGroupName(v string) *ListDataReportForServiceGroupResponseBodyData {
	s.ServiceGroupName = &v
	return s
}

func (s *ListDataReportForServiceGroupResponseBodyData) SetUnAcknowledgedEscalationIncidentCount(v int64) *ListDataReportForServiceGroupResponseBodyData {
	s.UnAcknowledgedEscalationIncidentCount = &v
	return s
}

func (s *ListDataReportForServiceGroupResponseBodyData) SetUnFinishEscalationIncidentCount(v int64) *ListDataReportForServiceGroupResponseBodyData {
	s.UnFinishEscalationIncidentCount = &v
	return s
}

type ListDataReportForServiceGroupResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDataReportForServiceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDataReportForServiceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataReportForServiceGroupResponse) GoString() string {
	return s.String()
}

func (s *ListDataReportForServiceGroupResponse) SetHeaders(v map[string]*string) *ListDataReportForServiceGroupResponse {
	s.Headers = v
	return s
}

func (s *ListDataReportForServiceGroupResponse) SetStatusCode(v int32) *ListDataReportForServiceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataReportForServiceGroupResponse) SetBody(v *ListDataReportForServiceGroupResponseBody) *ListDataReportForServiceGroupResponse {
	s.Body = v
	return s
}

type ListDataReportForUserRequest struct {
	EndTime    *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	PageNumber *int64  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	StartTime  *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListDataReportForUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataReportForUserRequest) GoString() string {
	return s.String()
}

func (s *ListDataReportForUserRequest) SetEndTime(v string) *ListDataReportForUserRequest {
	s.EndTime = &v
	return s
}

func (s *ListDataReportForUserRequest) SetPageNumber(v int64) *ListDataReportForUserRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataReportForUserRequest) SetPageSize(v int64) *ListDataReportForUserRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataReportForUserRequest) SetStartTime(v string) *ListDataReportForUserRequest {
	s.StartTime = &v
	return s
}

type ListDataReportForUserResponseBody struct {
	Data []*ListDataReportForUserResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Id of the request
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TotalCount *int64  `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListDataReportForUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataReportForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataReportForUserResponseBody) SetData(v []*ListDataReportForUserResponseBodyData) *ListDataReportForUserResponseBody {
	s.Data = v
	return s
}

func (s *ListDataReportForUserResponseBody) SetRequestId(v string) *ListDataReportForUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataReportForUserResponseBody) SetTotalCount(v int64) *ListDataReportForUserResponseBody {
	s.TotalCount = &v
	return s
}

type ListDataReportForUserResponseBodyData struct {
	DistributionIncidentCount *int64  `json:"distributionIncidentCount,omitempty" xml:"distributionIncidentCount,omitempty"`
	EscalationIncidentCount   *int64  `json:"escalationIncidentCount,omitempty" xml:"escalationIncidentCount,omitempty"`
	FinishIncidentNumber      *int64  `json:"finishIncidentNumber,omitempty" xml:"finishIncidentNumber,omitempty"`
	FinishProportion          *string `json:"finishProportion,omitempty" xml:"finishProportion,omitempty"`
	// MRRA
	MeanTimeToAcknowledge *string `json:"meanTimeToAcknowledge,omitempty" xml:"meanTimeToAcknowledge,omitempty"`
	// MTTA
	MeanTimeToRepair                      *string `json:"meanTimeToRepair,omitempty" xml:"meanTimeToRepair,omitempty"`
	UnAcknowledgedEscalationIncidentCount *int64  `json:"unAcknowledgedEscalationIncidentCount,omitempty" xml:"unAcknowledgedEscalationIncidentCount,omitempty"`
	UnDistributionIncidentCount           *int64  `json:"unDistributionIncidentCount,omitempty" xml:"unDistributionIncidentCount,omitempty"`
	UnFinishEscalationIncidentCount       *int64  `json:"unFinishEscalationIncidentCount,omitempty" xml:"unFinishEscalationIncidentCount,omitempty"`
	UserId                                *int64  `json:"userId,omitempty" xml:"userId,omitempty"`
	UserName                              *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s ListDataReportForUserResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDataReportForUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataReportForUserResponseBodyData) SetDistributionIncidentCount(v int64) *ListDataReportForUserResponseBodyData {
	s.DistributionIncidentCount = &v
	return s
}

func (s *ListDataReportForUserResponseBodyData) SetEscalationIncidentCount(v int64) *ListDataReportForUserResponseBodyData {
	s.EscalationIncidentCount = &v
	return s
}

func (s *ListDataReportForUserResponseBodyData) SetFinishIncidentNumber(v int64) *ListDataReportForUserResponseBodyData {
	s.FinishIncidentNumber = &v
	return s
}

func (s *ListDataReportForUserResponseBodyData) SetFinishProportion(v string) *ListDataReportForUserResponseBodyData {
	s.FinishProportion = &v
	return s
}

func (s *ListDataReportForUserResponseBodyData) SetMeanTimeToAcknowledge(v string) *ListDataReportForUserResponseBodyData {
	s.MeanTimeToAcknowledge = &v
	return s
}

func (s *ListDataReportForUserResponseBodyData) SetMeanTimeToRepair(v string) *ListDataReportForUserResponseBodyData {
	s.MeanTimeToRepair = &v
	return s
}

func (s *ListDataReportForUserResponseBodyData) SetUnAcknowledgedEscalationIncidentCount(v int64) *ListDataReportForUserResponseBodyData {
	s.UnAcknowledgedEscalationIncidentCount = &v
	return s
}

func (s *ListDataReportForUserResponseBodyData) SetUnDistributionIncidentCount(v int64) *ListDataReportForUserResponseBodyData {
	s.UnDistributionIncidentCount = &v
	return s
}

func (s *ListDataReportForUserResponseBodyData) SetUnFinishEscalationIncidentCount(v int64) *ListDataReportForUserResponseBodyData {
	s.UnFinishEscalationIncidentCount = &v
	return s
}

func (s *ListDataReportForUserResponseBodyData) SetUserId(v int64) *ListDataReportForUserResponseBodyData {
	s.UserId = &v
	return s
}

func (s *ListDataReportForUserResponseBodyData) SetUserName(v string) *ListDataReportForUserResponseBodyData {
	s.UserName = &v
	return s
}

type ListDataReportForUserResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDataReportForUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDataReportForUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataReportForUserResponse) GoString() string {
	return s.String()
}

func (s *ListDataReportForUserResponse) SetHeaders(v map[string]*string) *ListDataReportForUserResponse {
	s.Headers = v
	return s
}

func (s *ListDataReportForUserResponse) SetStatusCode(v int32) *ListDataReportForUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataReportForUserResponse) SetBody(v *ListDataReportForUserResponseBody) *ListDataReportForUserResponse {
	s.Body = v
	return s
}

type ListDictionariesRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s ListDictionariesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDictionariesRequest) GoString() string {
	return s.String()
}

func (s *ListDictionariesRequest) SetClientToken(v string) *ListDictionariesRequest {
	s.ClientToken = &v
	return s
}

type ListDictionariesResponseBody struct {
	// data
	Data map[string][]*DataValue `json:"data,omitempty" xml:"data,omitempty"`
	// id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListDictionariesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDictionariesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDictionariesResponseBody) SetData(v map[string][]*DataValue) *ListDictionariesResponseBody {
	s.Data = v
	return s
}

func (s *ListDictionariesResponseBody) SetRequestId(v string) *ListDictionariesResponseBody {
	s.RequestId = &v
	return s
}

type ListDictionariesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDictionariesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDictionariesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDictionariesResponse) GoString() string {
	return s.String()
}

func (s *ListDictionariesResponse) SetHeaders(v map[string]*string) *ListDictionariesResponse {
	s.Headers = v
	return s
}

func (s *ListDictionariesResponse) SetStatusCode(v int32) *ListDictionariesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDictionariesResponse) SetBody(v *ListDictionariesResponseBody) *ListDictionariesResponse {
	s.Body = v
	return s
}

type ListEscalationPlanServicesRequest struct {
	// clientToken
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s ListEscalationPlanServicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEscalationPlanServicesRequest) GoString() string {
	return s.String()
}

func (s *ListEscalationPlanServicesRequest) SetClientToken(v string) *ListEscalationPlanServicesRequest {
	s.ClientToken = &v
	return s
}

type ListEscalationPlanServicesResponseBody struct {
	// data
	Data []*ListEscalationPlanServicesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListEscalationPlanServicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEscalationPlanServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListEscalationPlanServicesResponseBody) SetData(v []*ListEscalationPlanServicesResponseBodyData) *ListEscalationPlanServicesResponseBody {
	s.Data = v
	return s
}

func (s *ListEscalationPlanServicesResponseBody) SetRequestId(v string) *ListEscalationPlanServicesResponseBody {
	s.RequestId = &v
	return s
}

type ListEscalationPlanServicesResponseBodyData struct {
	Scope         *string `json:"scope,omitempty" xml:"scope,omitempty"`
	ScopeObjectId *int64  `json:"scopeObjectId,omitempty" xml:"scopeObjectId,omitempty"`
}

func (s ListEscalationPlanServicesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListEscalationPlanServicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEscalationPlanServicesResponseBodyData) SetScope(v string) *ListEscalationPlanServicesResponseBodyData {
	s.Scope = &v
	return s
}

func (s *ListEscalationPlanServicesResponseBodyData) SetScopeObjectId(v int64) *ListEscalationPlanServicesResponseBodyData {
	s.ScopeObjectId = &v
	return s
}

type ListEscalationPlanServicesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListEscalationPlanServicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEscalationPlanServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEscalationPlanServicesResponse) GoString() string {
	return s.String()
}

func (s *ListEscalationPlanServicesResponse) SetHeaders(v map[string]*string) *ListEscalationPlanServicesResponse {
	s.Headers = v
	return s
}

func (s *ListEscalationPlanServicesResponse) SetStatusCode(v int32) *ListEscalationPlanServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEscalationPlanServicesResponse) SetBody(v *ListEscalationPlanServicesResponseBody) *ListEscalationPlanServicesResponse {
	s.Body = v
	return s
}

type ListEscalationPlansRequest struct {
	ClientToken        *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	EscalationPlanName *string `json:"escalationPlanName,omitempty" xml:"escalationPlanName,omitempty"`
	PageNumber         *int64  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize           *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ServiceName        *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
}

func (s ListEscalationPlansRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEscalationPlansRequest) GoString() string {
	return s.String()
}

func (s *ListEscalationPlansRequest) SetClientToken(v string) *ListEscalationPlansRequest {
	s.ClientToken = &v
	return s
}

func (s *ListEscalationPlansRequest) SetEscalationPlanName(v string) *ListEscalationPlansRequest {
	s.EscalationPlanName = &v
	return s
}

func (s *ListEscalationPlansRequest) SetPageNumber(v int64) *ListEscalationPlansRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEscalationPlansRequest) SetPageSize(v int64) *ListEscalationPlansRequest {
	s.PageSize = &v
	return s
}

func (s *ListEscalationPlansRequest) SetServiceName(v string) *ListEscalationPlansRequest {
	s.ServiceName = &v
	return s
}

type ListEscalationPlansResponseBody struct {
	// data
	Data       []*ListEscalationPlansResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	PageNumber *int64                                 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64                                 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Id of the request
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TotalCount *int64  `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListEscalationPlansResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEscalationPlansResponseBody) GoString() string {
	return s.String()
}

func (s *ListEscalationPlansResponseBody) SetData(v []*ListEscalationPlansResponseBodyData) *ListEscalationPlansResponseBody {
	s.Data = v
	return s
}

func (s *ListEscalationPlansResponseBody) SetPageNumber(v int64) *ListEscalationPlansResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListEscalationPlansResponseBody) SetPageSize(v int64) *ListEscalationPlansResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListEscalationPlansResponseBody) SetRequestId(v string) *ListEscalationPlansResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEscalationPlansResponseBody) SetTotalCount(v int64) *ListEscalationPlansResponseBody {
	s.TotalCount = &v
	return s
}

type ListEscalationPlansResponseBodyData struct {
	EscalationPlanId           *int64                                                           `json:"escalationPlanId,omitempty" xml:"escalationPlanId,omitempty"`
	EscalationPlanName         *string                                                          `json:"escalationPlanName,omitempty" xml:"escalationPlanName,omitempty"`
	EscalationPlanScopeObjects []*ListEscalationPlansResponseBodyDataEscalationPlanScopeObjects `json:"escalationPlanScopeObjects,omitempty" xml:"escalationPlanScopeObjects,omitempty" type:"Repeated"`
	ModifyTime                 *string                                                          `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	Status                     *string                                                          `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListEscalationPlansResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListEscalationPlansResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEscalationPlansResponseBodyData) SetEscalationPlanId(v int64) *ListEscalationPlansResponseBodyData {
	s.EscalationPlanId = &v
	return s
}

func (s *ListEscalationPlansResponseBodyData) SetEscalationPlanName(v string) *ListEscalationPlansResponseBodyData {
	s.EscalationPlanName = &v
	return s
}

func (s *ListEscalationPlansResponseBodyData) SetEscalationPlanScopeObjects(v []*ListEscalationPlansResponseBodyDataEscalationPlanScopeObjects) *ListEscalationPlansResponseBodyData {
	s.EscalationPlanScopeObjects = v
	return s
}

func (s *ListEscalationPlansResponseBodyData) SetModifyTime(v string) *ListEscalationPlansResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *ListEscalationPlansResponseBodyData) SetStatus(v string) *ListEscalationPlansResponseBodyData {
	s.Status = &v
	return s
}

type ListEscalationPlansResponseBodyDataEscalationPlanScopeObjects struct {
	Scope                  *string `json:"scope,omitempty" xml:"scope,omitempty"`
	ScopeObjectDeletedType *int32  `json:"scopeObjectDeletedType,omitempty" xml:"scopeObjectDeletedType,omitempty"`
	ScopeObjectId          *int64  `json:"scopeObjectId,omitempty" xml:"scopeObjectId,omitempty"`
	ScopeObjectName        *string `json:"scopeObjectName,omitempty" xml:"scopeObjectName,omitempty"`
}

func (s ListEscalationPlansResponseBodyDataEscalationPlanScopeObjects) String() string {
	return tea.Prettify(s)
}

func (s ListEscalationPlansResponseBodyDataEscalationPlanScopeObjects) GoString() string {
	return s.String()
}

func (s *ListEscalationPlansResponseBodyDataEscalationPlanScopeObjects) SetScope(v string) *ListEscalationPlansResponseBodyDataEscalationPlanScopeObjects {
	s.Scope = &v
	return s
}

func (s *ListEscalationPlansResponseBodyDataEscalationPlanScopeObjects) SetScopeObjectDeletedType(v int32) *ListEscalationPlansResponseBodyDataEscalationPlanScopeObjects {
	s.ScopeObjectDeletedType = &v
	return s
}

func (s *ListEscalationPlansResponseBodyDataEscalationPlanScopeObjects) SetScopeObjectId(v int64) *ListEscalationPlansResponseBodyDataEscalationPlanScopeObjects {
	s.ScopeObjectId = &v
	return s
}

func (s *ListEscalationPlansResponseBodyDataEscalationPlanScopeObjects) SetScopeObjectName(v string) *ListEscalationPlansResponseBodyDataEscalationPlanScopeObjects {
	s.ScopeObjectName = &v
	return s
}

type ListEscalationPlansResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListEscalationPlansResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEscalationPlansResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEscalationPlansResponse) GoString() string {
	return s.String()
}

func (s *ListEscalationPlansResponse) SetHeaders(v map[string]*string) *ListEscalationPlansResponse {
	s.Headers = v
	return s
}

func (s *ListEscalationPlansResponse) SetStatusCode(v int32) *ListEscalationPlansResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEscalationPlansResponse) SetBody(v *ListEscalationPlansResponseBody) *ListEscalationPlansResponse {
	s.Body = v
	return s
}

type ListEscalationPlansByNoticeObjectRequest struct {
	NoticeObjectId   *int64 `json:"noticeObjectId,omitempty" xml:"noticeObjectId,omitempty"`
	NoticeObjectType *int64 `json:"noticeObjectType,omitempty" xml:"noticeObjectType,omitempty"`
}

func (s ListEscalationPlansByNoticeObjectRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEscalationPlansByNoticeObjectRequest) GoString() string {
	return s.String()
}

func (s *ListEscalationPlansByNoticeObjectRequest) SetNoticeObjectId(v int64) *ListEscalationPlansByNoticeObjectRequest {
	s.NoticeObjectId = &v
	return s
}

func (s *ListEscalationPlansByNoticeObjectRequest) SetNoticeObjectType(v int64) *ListEscalationPlansByNoticeObjectRequest {
	s.NoticeObjectType = &v
	return s
}

type ListEscalationPlansByNoticeObjectResponseBody struct {
	Data      []*ListEscalationPlansByNoticeObjectResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	RequestId *string                                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListEscalationPlansByNoticeObjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEscalationPlansByNoticeObjectResponseBody) GoString() string {
	return s.String()
}

func (s *ListEscalationPlansByNoticeObjectResponseBody) SetData(v []*ListEscalationPlansByNoticeObjectResponseBodyData) *ListEscalationPlansByNoticeObjectResponseBody {
	s.Data = v
	return s
}

func (s *ListEscalationPlansByNoticeObjectResponseBody) SetRequestId(v string) *ListEscalationPlansByNoticeObjectResponseBody {
	s.RequestId = &v
	return s
}

type ListEscalationPlansByNoticeObjectResponseBodyData struct {
	EscalationPlanId           *int64                                                                         `json:"escalationPlanId,omitempty" xml:"escalationPlanId,omitempty"`
	EscalationPlanName         *string                                                                        `json:"escalationPlanName,omitempty" xml:"escalationPlanName,omitempty"`
	EscalationPlanScopeObjects []*ListEscalationPlansByNoticeObjectResponseBodyDataEscalationPlanScopeObjects `json:"escalationPlanScopeObjects,omitempty" xml:"escalationPlanScopeObjects,omitempty" type:"Repeated"`
	ModifyTime                 *string                                                                        `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	Status                     *string                                                                        `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListEscalationPlansByNoticeObjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListEscalationPlansByNoticeObjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEscalationPlansByNoticeObjectResponseBodyData) SetEscalationPlanId(v int64) *ListEscalationPlansByNoticeObjectResponseBodyData {
	s.EscalationPlanId = &v
	return s
}

func (s *ListEscalationPlansByNoticeObjectResponseBodyData) SetEscalationPlanName(v string) *ListEscalationPlansByNoticeObjectResponseBodyData {
	s.EscalationPlanName = &v
	return s
}

func (s *ListEscalationPlansByNoticeObjectResponseBodyData) SetEscalationPlanScopeObjects(v []*ListEscalationPlansByNoticeObjectResponseBodyDataEscalationPlanScopeObjects) *ListEscalationPlansByNoticeObjectResponseBodyData {
	s.EscalationPlanScopeObjects = v
	return s
}

func (s *ListEscalationPlansByNoticeObjectResponseBodyData) SetModifyTime(v string) *ListEscalationPlansByNoticeObjectResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *ListEscalationPlansByNoticeObjectResponseBodyData) SetStatus(v string) *ListEscalationPlansByNoticeObjectResponseBodyData {
	s.Status = &v
	return s
}

type ListEscalationPlansByNoticeObjectResponseBodyDataEscalationPlanScopeObjects struct {
	Scope                  *string `json:"scope,omitempty" xml:"scope,omitempty"`
	ScopeObjectDeletedType *int32  `json:"scopeObjectDeletedType,omitempty" xml:"scopeObjectDeletedType,omitempty"`
	ScopeObjectId          *int64  `json:"scopeObjectId,omitempty" xml:"scopeObjectId,omitempty"`
	ScopeObjectName        *string `json:"scopeObjectName,omitempty" xml:"scopeObjectName,omitempty"`
}

func (s ListEscalationPlansByNoticeObjectResponseBodyDataEscalationPlanScopeObjects) String() string {
	return tea.Prettify(s)
}

func (s ListEscalationPlansByNoticeObjectResponseBodyDataEscalationPlanScopeObjects) GoString() string {
	return s.String()
}

func (s *ListEscalationPlansByNoticeObjectResponseBodyDataEscalationPlanScopeObjects) SetScope(v string) *ListEscalationPlansByNoticeObjectResponseBodyDataEscalationPlanScopeObjects {
	s.Scope = &v
	return s
}

func (s *ListEscalationPlansByNoticeObjectResponseBodyDataEscalationPlanScopeObjects) SetScopeObjectDeletedType(v int32) *ListEscalationPlansByNoticeObjectResponseBodyDataEscalationPlanScopeObjects {
	s.ScopeObjectDeletedType = &v
	return s
}

func (s *ListEscalationPlansByNoticeObjectResponseBodyDataEscalationPlanScopeObjects) SetScopeObjectId(v int64) *ListEscalationPlansByNoticeObjectResponseBodyDataEscalationPlanScopeObjects {
	s.ScopeObjectId = &v
	return s
}

func (s *ListEscalationPlansByNoticeObjectResponseBodyDataEscalationPlanScopeObjects) SetScopeObjectName(v string) *ListEscalationPlansByNoticeObjectResponseBodyDataEscalationPlanScopeObjects {
	s.ScopeObjectName = &v
	return s
}

type ListEscalationPlansByNoticeObjectResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListEscalationPlansByNoticeObjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEscalationPlansByNoticeObjectResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEscalationPlansByNoticeObjectResponse) GoString() string {
	return s.String()
}

func (s *ListEscalationPlansByNoticeObjectResponse) SetHeaders(v map[string]*string) *ListEscalationPlansByNoticeObjectResponse {
	s.Headers = v
	return s
}

func (s *ListEscalationPlansByNoticeObjectResponse) SetStatusCode(v int32) *ListEscalationPlansByNoticeObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEscalationPlansByNoticeObjectResponse) SetBody(v *ListEscalationPlansByNoticeObjectResponseBody) *ListEscalationPlansByNoticeObjectResponse {
	s.Body = v
	return s
}

type ListIncidentDetailEscalationPlansRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	IncidentId  *int64  `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
}

func (s ListIncidentDetailEscalationPlansRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIncidentDetailEscalationPlansRequest) GoString() string {
	return s.String()
}

func (s *ListIncidentDetailEscalationPlansRequest) SetClientToken(v string) *ListIncidentDetailEscalationPlansRequest {
	s.ClientToken = &v
	return s
}

func (s *ListIncidentDetailEscalationPlansRequest) SetIncidentId(v int64) *ListIncidentDetailEscalationPlansRequest {
	s.IncidentId = &v
	return s
}

type ListIncidentDetailEscalationPlansResponseBody struct {
	// data
	Data *ListIncidentDetailEscalationPlansResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListIncidentDetailEscalationPlansResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIncidentDetailEscalationPlansResponseBody) GoString() string {
	return s.String()
}

func (s *ListIncidentDetailEscalationPlansResponseBody) SetData(v *ListIncidentDetailEscalationPlansResponseBodyData) *ListIncidentDetailEscalationPlansResponseBody {
	s.Data = v
	return s
}

func (s *ListIncidentDetailEscalationPlansResponseBody) SetRequestId(v string) *ListIncidentDetailEscalationPlansResponseBody {
	s.RequestId = &v
	return s
}

type ListIncidentDetailEscalationPlansResponseBodyData struct {
	ConvergenceEscalationPlan   []*ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlan   `json:"convergenceEscalationPlan,omitempty" xml:"convergenceEscalationPlan,omitempty" type:"Repeated"`
	EscalationPlanId            *int64                                                                          `json:"escalationPlanId,omitempty" xml:"escalationPlanId,omitempty"`
	EscalationPlanName          *string                                                                         `json:"escalationPlanName,omitempty" xml:"escalationPlanName,omitempty"`
	NuAcknowledgeEscalationPlan []*ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlan `json:"nuAcknowledgeEscalationPlan,omitempty" xml:"nuAcknowledgeEscalationPlan,omitempty" type:"Repeated"`
	UnFinishEscalationPlan      []*ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlan      `json:"unFinishEscalationPlan,omitempty" xml:"unFinishEscalationPlan,omitempty" type:"Repeated"`
}

func (s ListIncidentDetailEscalationPlansResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListIncidentDetailEscalationPlansResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIncidentDetailEscalationPlansResponseBodyData) SetConvergenceEscalationPlan(v []*ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlan) *ListIncidentDetailEscalationPlansResponseBodyData {
	s.ConvergenceEscalationPlan = v
	return s
}

func (s *ListIncidentDetailEscalationPlansResponseBodyData) SetEscalationPlanId(v int64) *ListIncidentDetailEscalationPlansResponseBodyData {
	s.EscalationPlanId = &v
	return s
}

func (s *ListIncidentDetailEscalationPlansResponseBodyData) SetEscalationPlanName(v string) *ListIncidentDetailEscalationPlansResponseBodyData {
	s.EscalationPlanName = &v
	return s
}

func (s *ListIncidentDetailEscalationPlansResponseBodyData) SetNuAcknowledgeEscalationPlan(v []*ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlan) *ListIncidentDetailEscalationPlansResponseBodyData {
	s.NuAcknowledgeEscalationPlan = v
	return s
}

func (s *ListIncidentDetailEscalationPlansResponseBodyData) SetUnFinishEscalationPlan(v []*ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlan) *ListIncidentDetailEscalationPlansResponseBodyData {
	s.UnFinishEscalationPlan = v
	return s
}

type ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlan struct {
	EscalationPlanType *string                                                                                       `json:"escalationPlanType,omitempty" xml:"escalationPlanType,omitempty"`
	NoticeChannels     []*string                                                                                     `json:"noticeChannels,omitempty" xml:"noticeChannels,omitempty" type:"Repeated"`
	NoticeObjectList   []*ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlanNoticeObjectList `json:"noticeObjectList,omitempty" xml:"noticeObjectList,omitempty" type:"Repeated"`
	NoticeTime         *int64                                                                                        `json:"noticeTime,omitempty" xml:"noticeTime,omitempty"`
	ServiceGroupList   []*ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlanServiceGroupList `json:"serviceGroupList,omitempty" xml:"serviceGroupList,omitempty" type:"Repeated"`
	StartTime          *int64                                                                                        `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status             *string                                                                                       `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlan) String() string {
	return tea.Prettify(s)
}

func (s ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlan) GoString() string {
	return s.String()
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlan) SetEscalationPlanType(v string) *ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlan {
	s.EscalationPlanType = &v
	return s
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlan) SetNoticeChannels(v []*string) *ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlan {
	s.NoticeChannels = v
	return s
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlan) SetNoticeObjectList(v []*ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlanNoticeObjectList) *ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlan {
	s.NoticeObjectList = v
	return s
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlan) SetNoticeTime(v int64) *ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlan {
	s.NoticeTime = &v
	return s
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlan) SetServiceGroupList(v []*ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlanServiceGroupList) *ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlan {
	s.ServiceGroupList = v
	return s
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlan) SetStartTime(v int64) *ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlan {
	s.StartTime = &v
	return s
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlan) SetStatus(v string) *ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlan {
	s.Status = &v
	return s
}

type ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlanNoticeObjectList struct {
	NoticeObjectId    *int64  `json:"noticeObjectId,omitempty" xml:"noticeObjectId,omitempty"`
	NoticeObjectName  *string `json:"noticeObjectName,omitempty" xml:"noticeObjectName,omitempty"`
	NoticeObjectPhone *string `json:"noticeObjectPhone,omitempty" xml:"noticeObjectPhone,omitempty"`
}

func (s ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlanNoticeObjectList) String() string {
	return tea.Prettify(s)
}

func (s ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlanNoticeObjectList) GoString() string {
	return s.String()
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlanNoticeObjectList) SetNoticeObjectId(v int64) *ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlanNoticeObjectList {
	s.NoticeObjectId = &v
	return s
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlanNoticeObjectList) SetNoticeObjectName(v string) *ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlanNoticeObjectList {
	s.NoticeObjectName = &v
	return s
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlanNoticeObjectList) SetNoticeObjectPhone(v string) *ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlanNoticeObjectList {
	s.NoticeObjectPhone = &v
	return s
}

type ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlanServiceGroupList struct {
	Id   *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlanServiceGroupList) String() string {
	return tea.Prettify(s)
}

func (s ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlanServiceGroupList) GoString() string {
	return s.String()
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlanServiceGroupList) SetId(v int64) *ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlanServiceGroupList {
	s.Id = &v
	return s
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlanServiceGroupList) SetName(v string) *ListIncidentDetailEscalationPlansResponseBodyDataConvergenceEscalationPlanServiceGroupList {
	s.Name = &v
	return s
}

type ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlan struct {
	EscalationPlanType *string                                                                                         `json:"escalationPlanType,omitempty" xml:"escalationPlanType,omitempty"`
	NoticeChannels     []*string                                                                                       `json:"noticeChannels,omitempty" xml:"noticeChannels,omitempty" type:"Repeated"`
	NoticeObjectList   []*ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlanNoticeObjectList `json:"noticeObjectList,omitempty" xml:"noticeObjectList,omitempty" type:"Repeated"`
	NoticeTime         *int64                                                                                          `json:"noticeTime,omitempty" xml:"noticeTime,omitempty"`
	ServiceGroupList   []*ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlanServiceGroupList `json:"serviceGroupList,omitempty" xml:"serviceGroupList,omitempty" type:"Repeated"`
	StartTime          *int64                                                                                          `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status             *string                                                                                         `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlan) String() string {
	return tea.Prettify(s)
}

func (s ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlan) GoString() string {
	return s.String()
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlan) SetEscalationPlanType(v string) *ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlan {
	s.EscalationPlanType = &v
	return s
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlan) SetNoticeChannels(v []*string) *ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlan {
	s.NoticeChannels = v
	return s
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlan) SetNoticeObjectList(v []*ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlanNoticeObjectList) *ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlan {
	s.NoticeObjectList = v
	return s
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlan) SetNoticeTime(v int64) *ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlan {
	s.NoticeTime = &v
	return s
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlan) SetServiceGroupList(v []*ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlanServiceGroupList) *ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlan {
	s.ServiceGroupList = v
	return s
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlan) SetStartTime(v int64) *ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlan {
	s.StartTime = &v
	return s
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlan) SetStatus(v string) *ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlan {
	s.Status = &v
	return s
}

type ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlanNoticeObjectList struct {
	NoticeObjectId    *int64  `json:"noticeObjectId,omitempty" xml:"noticeObjectId,omitempty"`
	NoticeObjectName  *string `json:"noticeObjectName,omitempty" xml:"noticeObjectName,omitempty"`
	NoticeObjectPhone *string `json:"noticeObjectPhone,omitempty" xml:"noticeObjectPhone,omitempty"`
}

func (s ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlanNoticeObjectList) String() string {
	return tea.Prettify(s)
}

func (s ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlanNoticeObjectList) GoString() string {
	return s.String()
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlanNoticeObjectList) SetNoticeObjectId(v int64) *ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlanNoticeObjectList {
	s.NoticeObjectId = &v
	return s
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlanNoticeObjectList) SetNoticeObjectName(v string) *ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlanNoticeObjectList {
	s.NoticeObjectName = &v
	return s
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlanNoticeObjectList) SetNoticeObjectPhone(v string) *ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlanNoticeObjectList {
	s.NoticeObjectPhone = &v
	return s
}

type ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlanServiceGroupList struct {
	Id   *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlanServiceGroupList) String() string {
	return tea.Prettify(s)
}

func (s ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlanServiceGroupList) GoString() string {
	return s.String()
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlanServiceGroupList) SetId(v int64) *ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlanServiceGroupList {
	s.Id = &v
	return s
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlanServiceGroupList) SetName(v string) *ListIncidentDetailEscalationPlansResponseBodyDataNuAcknowledgeEscalationPlanServiceGroupList {
	s.Name = &v
	return s
}

type ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlan struct {
	EscalationPlanType *string                                                                                    `json:"escalationPlanType,omitempty" xml:"escalationPlanType,omitempty"`
	NoticeChannels     []*string                                                                                  `json:"noticeChannels,omitempty" xml:"noticeChannels,omitempty" type:"Repeated"`
	NoticeObjectList   []*ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlanNoticeObjectList `json:"noticeObjectList,omitempty" xml:"noticeObjectList,omitempty" type:"Repeated"`
	NoticeTime         *int32                                                                                     `json:"noticeTime,omitempty" xml:"noticeTime,omitempty"`
	ServiceGroupList   []*ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlanServiceGroupList `json:"serviceGroupList,omitempty" xml:"serviceGroupList,omitempty" type:"Repeated"`
	StartTime          *int64                                                                                     `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status             *string                                                                                    `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlan) String() string {
	return tea.Prettify(s)
}

func (s ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlan) GoString() string {
	return s.String()
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlan) SetEscalationPlanType(v string) *ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlan {
	s.EscalationPlanType = &v
	return s
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlan) SetNoticeChannels(v []*string) *ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlan {
	s.NoticeChannels = v
	return s
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlan) SetNoticeObjectList(v []*ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlanNoticeObjectList) *ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlan {
	s.NoticeObjectList = v
	return s
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlan) SetNoticeTime(v int32) *ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlan {
	s.NoticeTime = &v
	return s
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlan) SetServiceGroupList(v []*ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlanServiceGroupList) *ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlan {
	s.ServiceGroupList = v
	return s
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlan) SetStartTime(v int64) *ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlan {
	s.StartTime = &v
	return s
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlan) SetStatus(v string) *ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlan {
	s.Status = &v
	return s
}

type ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlanNoticeObjectList struct {
	NoticeObjectId    *int64  `json:"noticeObjectId,omitempty" xml:"noticeObjectId,omitempty"`
	NoticeObjectName  *string `json:"noticeObjectName,omitempty" xml:"noticeObjectName,omitempty"`
	NoticeObjectPhone *string `json:"noticeObjectPhone,omitempty" xml:"noticeObjectPhone,omitempty"`
}

func (s ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlanNoticeObjectList) String() string {
	return tea.Prettify(s)
}

func (s ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlanNoticeObjectList) GoString() string {
	return s.String()
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlanNoticeObjectList) SetNoticeObjectId(v int64) *ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlanNoticeObjectList {
	s.NoticeObjectId = &v
	return s
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlanNoticeObjectList) SetNoticeObjectName(v string) *ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlanNoticeObjectList {
	s.NoticeObjectName = &v
	return s
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlanNoticeObjectList) SetNoticeObjectPhone(v string) *ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlanNoticeObjectList {
	s.NoticeObjectPhone = &v
	return s
}

type ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlanServiceGroupList struct {
	Id   *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlanServiceGroupList) String() string {
	return tea.Prettify(s)
}

func (s ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlanServiceGroupList) GoString() string {
	return s.String()
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlanServiceGroupList) SetId(v int64) *ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlanServiceGroupList {
	s.Id = &v
	return s
}

func (s *ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlanServiceGroupList) SetName(v string) *ListIncidentDetailEscalationPlansResponseBodyDataUnFinishEscalationPlanServiceGroupList {
	s.Name = &v
	return s
}

type ListIncidentDetailEscalationPlansResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListIncidentDetailEscalationPlansResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListIncidentDetailEscalationPlansResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIncidentDetailEscalationPlansResponse) GoString() string {
	return s.String()
}

func (s *ListIncidentDetailEscalationPlansResponse) SetHeaders(v map[string]*string) *ListIncidentDetailEscalationPlansResponse {
	s.Headers = v
	return s
}

func (s *ListIncidentDetailEscalationPlansResponse) SetStatusCode(v int32) *ListIncidentDetailEscalationPlansResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIncidentDetailEscalationPlansResponse) SetBody(v *ListIncidentDetailEscalationPlansResponseBody) *ListIncidentDetailEscalationPlansResponse {
	s.Body = v
	return s
}

type ListIncidentDetailTimelinesRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	IdSort      *string `json:"idSort,omitempty" xml:"idSort,omitempty"`
	IncidentId  *int64  `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
	PageNumber  *int64  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize    *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListIncidentDetailTimelinesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIncidentDetailTimelinesRequest) GoString() string {
	return s.String()
}

func (s *ListIncidentDetailTimelinesRequest) SetClientToken(v string) *ListIncidentDetailTimelinesRequest {
	s.ClientToken = &v
	return s
}

func (s *ListIncidentDetailTimelinesRequest) SetIdSort(v string) *ListIncidentDetailTimelinesRequest {
	s.IdSort = &v
	return s
}

func (s *ListIncidentDetailTimelinesRequest) SetIncidentId(v int64) *ListIncidentDetailTimelinesRequest {
	s.IncidentId = &v
	return s
}

func (s *ListIncidentDetailTimelinesRequest) SetPageNumber(v int64) *ListIncidentDetailTimelinesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListIncidentDetailTimelinesRequest) SetPageSize(v int64) *ListIncidentDetailTimelinesRequest {
	s.PageSize = &v
	return s
}

type ListIncidentDetailTimelinesResponseBody struct {
	// data
	Data       []*ListIncidentDetailTimelinesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	PageNumber *int32                                         `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32                                         `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Id of the request
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TotalCount *int32  `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListIncidentDetailTimelinesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIncidentDetailTimelinesResponseBody) GoString() string {
	return s.String()
}

func (s *ListIncidentDetailTimelinesResponseBody) SetData(v []*ListIncidentDetailTimelinesResponseBodyData) *ListIncidentDetailTimelinesResponseBody {
	s.Data = v
	return s
}

func (s *ListIncidentDetailTimelinesResponseBody) SetPageNumber(v int32) *ListIncidentDetailTimelinesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListIncidentDetailTimelinesResponseBody) SetPageSize(v int32) *ListIncidentDetailTimelinesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListIncidentDetailTimelinesResponseBody) SetRequestId(v string) *ListIncidentDetailTimelinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIncidentDetailTimelinesResponseBody) SetTotalCount(v int32) *ListIncidentDetailTimelinesResponseBody {
	s.TotalCount = &v
	return s
}

type ListIncidentDetailTimelinesResponseBodyData struct {
	Action                 *string `json:"action,omitempty" xml:"action,omitempty"`
	CreateTime             *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description            *string `json:"description,omitempty" xml:"description,omitempty"`
	IncidentId             *int64  `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
	RelRouteRuleDeleteType *int32  `json:"relRouteRuleDeleteType,omitempty" xml:"relRouteRuleDeleteType,omitempty"`
	RelatedServiceName     *string `json:"relatedServiceName,omitempty" xml:"relatedServiceName,omitempty"`
	Remark                 *string `json:"remark,omitempty" xml:"remark,omitempty"`
	SnapshotData           *string `json:"snapshotData,omitempty" xml:"snapshotData,omitempty"`
	Title                  *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ListIncidentDetailTimelinesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListIncidentDetailTimelinesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIncidentDetailTimelinesResponseBodyData) SetAction(v string) *ListIncidentDetailTimelinesResponseBodyData {
	s.Action = &v
	return s
}

func (s *ListIncidentDetailTimelinesResponseBodyData) SetCreateTime(v string) *ListIncidentDetailTimelinesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListIncidentDetailTimelinesResponseBodyData) SetDescription(v string) *ListIncidentDetailTimelinesResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListIncidentDetailTimelinesResponseBodyData) SetIncidentId(v int64) *ListIncidentDetailTimelinesResponseBodyData {
	s.IncidentId = &v
	return s
}

func (s *ListIncidentDetailTimelinesResponseBodyData) SetRelRouteRuleDeleteType(v int32) *ListIncidentDetailTimelinesResponseBodyData {
	s.RelRouteRuleDeleteType = &v
	return s
}

func (s *ListIncidentDetailTimelinesResponseBodyData) SetRelatedServiceName(v string) *ListIncidentDetailTimelinesResponseBodyData {
	s.RelatedServiceName = &v
	return s
}

func (s *ListIncidentDetailTimelinesResponseBodyData) SetRemark(v string) *ListIncidentDetailTimelinesResponseBodyData {
	s.Remark = &v
	return s
}

func (s *ListIncidentDetailTimelinesResponseBodyData) SetSnapshotData(v string) *ListIncidentDetailTimelinesResponseBodyData {
	s.SnapshotData = &v
	return s
}

func (s *ListIncidentDetailTimelinesResponseBodyData) SetTitle(v string) *ListIncidentDetailTimelinesResponseBodyData {
	s.Title = &v
	return s
}

type ListIncidentDetailTimelinesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListIncidentDetailTimelinesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListIncidentDetailTimelinesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIncidentDetailTimelinesResponse) GoString() string {
	return s.String()
}

func (s *ListIncidentDetailTimelinesResponse) SetHeaders(v map[string]*string) *ListIncidentDetailTimelinesResponse {
	s.Headers = v
	return s
}

func (s *ListIncidentDetailTimelinesResponse) SetStatusCode(v int32) *ListIncidentDetailTimelinesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIncidentDetailTimelinesResponse) SetBody(v *ListIncidentDetailTimelinesResponseBody) *ListIncidentDetailTimelinesResponse {
	s.Body = v
	return s
}

type ListIncidentSubtotalsRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	IncidentId  *int64  `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
}

func (s ListIncidentSubtotalsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIncidentSubtotalsRequest) GoString() string {
	return s.String()
}

func (s *ListIncidentSubtotalsRequest) SetClientToken(v string) *ListIncidentSubtotalsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListIncidentSubtotalsRequest) SetIncidentId(v int64) *ListIncidentSubtotalsRequest {
	s.IncidentId = &v
	return s
}

type ListIncidentSubtotalsResponseBody struct {
	// data
	Data []*ListIncidentSubtotalsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListIncidentSubtotalsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIncidentSubtotalsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIncidentSubtotalsResponseBody) SetData(v []*ListIncidentSubtotalsResponseBodyData) *ListIncidentSubtotalsResponseBody {
	s.Data = v
	return s
}

func (s *ListIncidentSubtotalsResponseBody) SetRequestId(v string) *ListIncidentSubtotalsResponseBody {
	s.RequestId = &v
	return s
}

type ListIncidentSubtotalsResponseBodyData struct {
	CreateTime      *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreateUserId    *int64  `json:"createUserId,omitempty" xml:"createUserId,omitempty"`
	CreateUserName  *string `json:"createUserName,omitempty" xml:"createUserName,omitempty"`
	CreateUserPhone *string `json:"createUserPhone,omitempty" xml:"createUserPhone,omitempty"`
	Description     *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s ListIncidentSubtotalsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListIncidentSubtotalsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIncidentSubtotalsResponseBodyData) SetCreateTime(v string) *ListIncidentSubtotalsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListIncidentSubtotalsResponseBodyData) SetCreateUserId(v int64) *ListIncidentSubtotalsResponseBodyData {
	s.CreateUserId = &v
	return s
}

func (s *ListIncidentSubtotalsResponseBodyData) SetCreateUserName(v string) *ListIncidentSubtotalsResponseBodyData {
	s.CreateUserName = &v
	return s
}

func (s *ListIncidentSubtotalsResponseBodyData) SetCreateUserPhone(v string) *ListIncidentSubtotalsResponseBodyData {
	s.CreateUserPhone = &v
	return s
}

func (s *ListIncidentSubtotalsResponseBodyData) SetDescription(v string) *ListIncidentSubtotalsResponseBodyData {
	s.Description = &v
	return s
}

type ListIncidentSubtotalsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListIncidentSubtotalsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListIncidentSubtotalsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIncidentSubtotalsResponse) GoString() string {
	return s.String()
}

func (s *ListIncidentSubtotalsResponse) SetHeaders(v map[string]*string) *ListIncidentSubtotalsResponse {
	s.Headers = v
	return s
}

func (s *ListIncidentSubtotalsResponse) SetStatusCode(v int32) *ListIncidentSubtotalsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIncidentSubtotalsResponse) SetBody(v *ListIncidentSubtotalsResponseBody) *ListIncidentSubtotalsResponse {
	s.Body = v
	return s
}

type ListIncidentTimelinesRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	PageNumber  *int64  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize    *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListIncidentTimelinesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIncidentTimelinesRequest) GoString() string {
	return s.String()
}

func (s *ListIncidentTimelinesRequest) SetClientToken(v string) *ListIncidentTimelinesRequest {
	s.ClientToken = &v
	return s
}

func (s *ListIncidentTimelinesRequest) SetPageNumber(v int64) *ListIncidentTimelinesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListIncidentTimelinesRequest) SetPageSize(v int64) *ListIncidentTimelinesRequest {
	s.PageSize = &v
	return s
}

type ListIncidentTimelinesResponseBody struct {
	// data
	Data       []*ListIncidentTimelinesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	PageNumber *int32                                   `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32                                   `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// requestId
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TotalCount *int32  `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListIncidentTimelinesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIncidentTimelinesResponseBody) GoString() string {
	return s.String()
}

func (s *ListIncidentTimelinesResponseBody) SetData(v []*ListIncidentTimelinesResponseBodyData) *ListIncidentTimelinesResponseBody {
	s.Data = v
	return s
}

func (s *ListIncidentTimelinesResponseBody) SetPageNumber(v int32) *ListIncidentTimelinesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListIncidentTimelinesResponseBody) SetPageSize(v int32) *ListIncidentTimelinesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListIncidentTimelinesResponseBody) SetRequestId(v string) *ListIncidentTimelinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIncidentTimelinesResponseBody) SetTotalCount(v int32) *ListIncidentTimelinesResponseBody {
	s.TotalCount = &v
	return s
}

type ListIncidentTimelinesResponseBodyData struct {
	Action                 *string `json:"action,omitempty" xml:"action,omitempty"`
	CreateTime             *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description            *int64  `json:"description,omitempty" xml:"description,omitempty"`
	IncidentId             *int64  `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
	IncidentNumber         *string `json:"incidentNumber,omitempty" xml:"incidentNumber,omitempty"`
	IncidentTitle          *string `json:"incidentTitle,omitempty" xml:"incidentTitle,omitempty"`
	RelRouteRuleDeleteType *int32  `json:"relRouteRuleDeleteType,omitempty" xml:"relRouteRuleDeleteType,omitempty"`
	RelatedServiceName     *string `json:"relatedServiceName,omitempty" xml:"relatedServiceName,omitempty"`
	Remark                 *string `json:"remark,omitempty" xml:"remark,omitempty"`
	SnapshotData           *string `json:"snapshotData,omitempty" xml:"snapshotData,omitempty"`
	Title                  *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ListIncidentTimelinesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListIncidentTimelinesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIncidentTimelinesResponseBodyData) SetAction(v string) *ListIncidentTimelinesResponseBodyData {
	s.Action = &v
	return s
}

func (s *ListIncidentTimelinesResponseBodyData) SetCreateTime(v string) *ListIncidentTimelinesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListIncidentTimelinesResponseBodyData) SetDescription(v int64) *ListIncidentTimelinesResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListIncidentTimelinesResponseBodyData) SetIncidentId(v int64) *ListIncidentTimelinesResponseBodyData {
	s.IncidentId = &v
	return s
}

func (s *ListIncidentTimelinesResponseBodyData) SetIncidentNumber(v string) *ListIncidentTimelinesResponseBodyData {
	s.IncidentNumber = &v
	return s
}

func (s *ListIncidentTimelinesResponseBodyData) SetIncidentTitle(v string) *ListIncidentTimelinesResponseBodyData {
	s.IncidentTitle = &v
	return s
}

func (s *ListIncidentTimelinesResponseBodyData) SetRelRouteRuleDeleteType(v int32) *ListIncidentTimelinesResponseBodyData {
	s.RelRouteRuleDeleteType = &v
	return s
}

func (s *ListIncidentTimelinesResponseBodyData) SetRelatedServiceName(v string) *ListIncidentTimelinesResponseBodyData {
	s.RelatedServiceName = &v
	return s
}

func (s *ListIncidentTimelinesResponseBodyData) SetRemark(v string) *ListIncidentTimelinesResponseBodyData {
	s.Remark = &v
	return s
}

func (s *ListIncidentTimelinesResponseBodyData) SetSnapshotData(v string) *ListIncidentTimelinesResponseBodyData {
	s.SnapshotData = &v
	return s
}

func (s *ListIncidentTimelinesResponseBodyData) SetTitle(v string) *ListIncidentTimelinesResponseBodyData {
	s.Title = &v
	return s
}

type ListIncidentTimelinesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListIncidentTimelinesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListIncidentTimelinesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIncidentTimelinesResponse) GoString() string {
	return s.String()
}

func (s *ListIncidentTimelinesResponse) SetHeaders(v map[string]*string) *ListIncidentTimelinesResponse {
	s.Headers = v
	return s
}

func (s *ListIncidentTimelinesResponse) SetStatusCode(v int32) *ListIncidentTimelinesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIncidentTimelinesResponse) SetBody(v *ListIncidentTimelinesResponseBody) *ListIncidentTimelinesResponse {
	s.Body = v
	return s
}

type ListIncidentsRequest struct {
	ClientToken       *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	CreateEndTime     *string `json:"createEndTime,omitempty" xml:"createEndTime,omitempty"`
	CreateStartTime   *string `json:"createStartTime,omitempty" xml:"createStartTime,omitempty"`
	Effect            *string `json:"effect,omitempty" xml:"effect,omitempty"`
	IncidentLevel     *string `json:"incidentLevel,omitempty" xml:"incidentLevel,omitempty"`
	IncidentStatus    *string `json:"incidentStatus,omitempty" xml:"incidentStatus,omitempty"`
	Me                *int32  `json:"me,omitempty" xml:"me,omitempty"`
	PageNumber        *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize          *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	RelationServiceId *int64  `json:"relationServiceId,omitempty" xml:"relationServiceId,omitempty"`
	RuleName          *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
}

func (s ListIncidentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIncidentsRequest) GoString() string {
	return s.String()
}

func (s *ListIncidentsRequest) SetClientToken(v string) *ListIncidentsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListIncidentsRequest) SetCreateEndTime(v string) *ListIncidentsRequest {
	s.CreateEndTime = &v
	return s
}

func (s *ListIncidentsRequest) SetCreateStartTime(v string) *ListIncidentsRequest {
	s.CreateStartTime = &v
	return s
}

func (s *ListIncidentsRequest) SetEffect(v string) *ListIncidentsRequest {
	s.Effect = &v
	return s
}

func (s *ListIncidentsRequest) SetIncidentLevel(v string) *ListIncidentsRequest {
	s.IncidentLevel = &v
	return s
}

func (s *ListIncidentsRequest) SetIncidentStatus(v string) *ListIncidentsRequest {
	s.IncidentStatus = &v
	return s
}

func (s *ListIncidentsRequest) SetMe(v int32) *ListIncidentsRequest {
	s.Me = &v
	return s
}

func (s *ListIncidentsRequest) SetPageNumber(v int32) *ListIncidentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListIncidentsRequest) SetPageSize(v int32) *ListIncidentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListIncidentsRequest) SetRelationServiceId(v int64) *ListIncidentsRequest {
	s.RelationServiceId = &v
	return s
}

func (s *ListIncidentsRequest) SetRuleName(v string) *ListIncidentsRequest {
	s.RuleName = &v
	return s
}

type ListIncidentsResponseBody struct {
	// data
	Data       []*ListIncidentsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	PageNumber *int32                           `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32                           `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// requestId
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TotalCount *int32  `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListIncidentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIncidentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIncidentsResponseBody) SetData(v []*ListIncidentsResponseBodyData) *ListIncidentsResponseBody {
	s.Data = v
	return s
}

func (s *ListIncidentsResponseBody) SetPageNumber(v int32) *ListIncidentsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListIncidentsResponseBody) SetPageSize(v int32) *ListIncidentsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListIncidentsResponseBody) SetRequestId(v string) *ListIncidentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIncidentsResponseBody) SetTotalCount(v int32) *ListIncidentsResponseBody {
	s.TotalCount = &v
	return s
}

type ListIncidentsResponseBodyData struct {
	AssignToWhoIsValid *int64 `json:"assignToWhoIsValid,omitempty" xml:"assignToWhoIsValid,omitempty"`
	// 代表创建时间的资源属性字段
	AssignUserId *int64 `json:"assignUserId,omitempty" xml:"assignUserId,omitempty"`
	// 代表资源一级ID的资源属性字段
	AssignUserName  *string `json:"assignUserName,omitempty" xml:"assignUserName,omitempty"`
	AssignUserPhone *string `json:"assignUserPhone,omitempty" xml:"assignUserPhone,omitempty"`
	// 事件级别
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 时间指派人ID
	Effect *string `json:"effect,omitempty" xml:"effect,omitempty"`
	// 修改时间
	IncidentId *int64 `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
	// 影响程度
	IncidentLevel  *string `json:"incidentLevel,omitempty" xml:"incidentLevel,omitempty"`
	IncidentNumber *string `json:"incidentNumber,omitempty" xml:"incidentNumber,omitempty"`
	// 关联流转规则ID
	IncidentStatus *string `json:"incidentStatus,omitempty" xml:"incidentStatus,omitempty"`
	// 事件内容
	IncidentTitle          *string `json:"incidentTitle,omitempty" xml:"incidentTitle,omitempty"`
	IsManual               *bool   `json:"isManual,omitempty" xml:"isManual,omitempty"`
	RelRouteRuleDeleteType *int32  `json:"relRouteRuleDeleteType,omitempty" xml:"relRouteRuleDeleteType,omitempty"`
	RelServiceDeleteType   *int32  `json:"relServiceDeleteType,omitempty" xml:"relServiceDeleteType,omitempty"`
	// 事件状态
	RelatedServiceId   *int64  `json:"relatedServiceId,omitempty" xml:"relatedServiceId,omitempty"`
	RelatedServiceName *string `json:"relatedServiceName,omitempty" xml:"relatedServiceName,omitempty"`
	// 关联的服务ID
	RouteRuleId   *int64  `json:"routeRuleId,omitempty" xml:"routeRuleId,omitempty"`
	RouteRuleName *string `json:"routeRuleName,omitempty" xml:"routeRuleName,omitempty"`
}

func (s ListIncidentsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListIncidentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIncidentsResponseBodyData) SetAssignToWhoIsValid(v int64) *ListIncidentsResponseBodyData {
	s.AssignToWhoIsValid = &v
	return s
}

func (s *ListIncidentsResponseBodyData) SetAssignUserId(v int64) *ListIncidentsResponseBodyData {
	s.AssignUserId = &v
	return s
}

func (s *ListIncidentsResponseBodyData) SetAssignUserName(v string) *ListIncidentsResponseBodyData {
	s.AssignUserName = &v
	return s
}

func (s *ListIncidentsResponseBodyData) SetAssignUserPhone(v string) *ListIncidentsResponseBodyData {
	s.AssignUserPhone = &v
	return s
}

func (s *ListIncidentsResponseBodyData) SetCreateTime(v string) *ListIncidentsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListIncidentsResponseBodyData) SetEffect(v string) *ListIncidentsResponseBodyData {
	s.Effect = &v
	return s
}

func (s *ListIncidentsResponseBodyData) SetIncidentId(v int64) *ListIncidentsResponseBodyData {
	s.IncidentId = &v
	return s
}

func (s *ListIncidentsResponseBodyData) SetIncidentLevel(v string) *ListIncidentsResponseBodyData {
	s.IncidentLevel = &v
	return s
}

func (s *ListIncidentsResponseBodyData) SetIncidentNumber(v string) *ListIncidentsResponseBodyData {
	s.IncidentNumber = &v
	return s
}

func (s *ListIncidentsResponseBodyData) SetIncidentStatus(v string) *ListIncidentsResponseBodyData {
	s.IncidentStatus = &v
	return s
}

func (s *ListIncidentsResponseBodyData) SetIncidentTitle(v string) *ListIncidentsResponseBodyData {
	s.IncidentTitle = &v
	return s
}

func (s *ListIncidentsResponseBodyData) SetIsManual(v bool) *ListIncidentsResponseBodyData {
	s.IsManual = &v
	return s
}

func (s *ListIncidentsResponseBodyData) SetRelRouteRuleDeleteType(v int32) *ListIncidentsResponseBodyData {
	s.RelRouteRuleDeleteType = &v
	return s
}

func (s *ListIncidentsResponseBodyData) SetRelServiceDeleteType(v int32) *ListIncidentsResponseBodyData {
	s.RelServiceDeleteType = &v
	return s
}

func (s *ListIncidentsResponseBodyData) SetRelatedServiceId(v int64) *ListIncidentsResponseBodyData {
	s.RelatedServiceId = &v
	return s
}

func (s *ListIncidentsResponseBodyData) SetRelatedServiceName(v string) *ListIncidentsResponseBodyData {
	s.RelatedServiceName = &v
	return s
}

func (s *ListIncidentsResponseBodyData) SetRouteRuleId(v int64) *ListIncidentsResponseBodyData {
	s.RouteRuleId = &v
	return s
}

func (s *ListIncidentsResponseBodyData) SetRouteRuleName(v string) *ListIncidentsResponseBodyData {
	s.RouteRuleName = &v
	return s
}

type ListIncidentsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListIncidentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListIncidentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIncidentsResponse) GoString() string {
	return s.String()
}

func (s *ListIncidentsResponse) SetHeaders(v map[string]*string) *ListIncidentsResponse {
	s.Headers = v
	return s
}

func (s *ListIncidentsResponse) SetStatusCode(v int32) *ListIncidentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIncidentsResponse) SetBody(v *ListIncidentsResponseBody) *ListIncidentsResponse {
	s.Body = v
	return s
}

type ListIntegrationConfigTimelinesRequest struct {
	ClientToken         *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	IntegrationConfigId *int64  `json:"integrationConfigId,omitempty" xml:"integrationConfigId,omitempty"`
	PageNumber          *int64  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize            *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListIntegrationConfigTimelinesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIntegrationConfigTimelinesRequest) GoString() string {
	return s.String()
}

func (s *ListIntegrationConfigTimelinesRequest) SetClientToken(v string) *ListIntegrationConfigTimelinesRequest {
	s.ClientToken = &v
	return s
}

func (s *ListIntegrationConfigTimelinesRequest) SetIntegrationConfigId(v int64) *ListIntegrationConfigTimelinesRequest {
	s.IntegrationConfigId = &v
	return s
}

func (s *ListIntegrationConfigTimelinesRequest) SetPageNumber(v int64) *ListIntegrationConfigTimelinesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListIntegrationConfigTimelinesRequest) SetPageSize(v int64) *ListIntegrationConfigTimelinesRequest {
	s.PageSize = &v
	return s
}

type ListIntegrationConfigTimelinesResponseBody struct {
	// data
	Data []*ListIntegrationConfigTimelinesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// pageNumber
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// pageSize
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// totalCount
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListIntegrationConfigTimelinesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIntegrationConfigTimelinesResponseBody) GoString() string {
	return s.String()
}

func (s *ListIntegrationConfigTimelinesResponseBody) SetData(v []*ListIntegrationConfigTimelinesResponseBodyData) *ListIntegrationConfigTimelinesResponseBody {
	s.Data = v
	return s
}

func (s *ListIntegrationConfigTimelinesResponseBody) SetPageNumber(v int64) *ListIntegrationConfigTimelinesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListIntegrationConfigTimelinesResponseBody) SetPageSize(v int64) *ListIntegrationConfigTimelinesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListIntegrationConfigTimelinesResponseBody) SetRequestId(v string) *ListIntegrationConfigTimelinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIntegrationConfigTimelinesResponseBody) SetTotalCount(v int64) *ListIntegrationConfigTimelinesResponseBody {
	s.TotalCount = &v
	return s
}

type ListIntegrationConfigTimelinesResponseBodyData struct {
	CreateTime  *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Title       *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ListIntegrationConfigTimelinesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListIntegrationConfigTimelinesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIntegrationConfigTimelinesResponseBodyData) SetCreateTime(v string) *ListIntegrationConfigTimelinesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListIntegrationConfigTimelinesResponseBodyData) SetDescription(v string) *ListIntegrationConfigTimelinesResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListIntegrationConfigTimelinesResponseBodyData) SetTitle(v string) *ListIntegrationConfigTimelinesResponseBodyData {
	s.Title = &v
	return s
}

type ListIntegrationConfigTimelinesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListIntegrationConfigTimelinesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListIntegrationConfigTimelinesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIntegrationConfigTimelinesResponse) GoString() string {
	return s.String()
}

func (s *ListIntegrationConfigTimelinesResponse) SetHeaders(v map[string]*string) *ListIntegrationConfigTimelinesResponse {
	s.Headers = v
	return s
}

func (s *ListIntegrationConfigTimelinesResponse) SetStatusCode(v int32) *ListIntegrationConfigTimelinesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIntegrationConfigTimelinesResponse) SetBody(v *ListIntegrationConfigTimelinesResponseBody) *ListIntegrationConfigTimelinesResponse {
	s.Body = v
	return s
}

type ListIntegrationConfigsRequest struct {
	ClientToken       *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	MonitorSourceName *string `json:"monitorSourceName,omitempty" xml:"monitorSourceName,omitempty"`
}

func (s ListIntegrationConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIntegrationConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListIntegrationConfigsRequest) SetClientToken(v string) *ListIntegrationConfigsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListIntegrationConfigsRequest) SetMonitorSourceName(v string) *ListIntegrationConfigsRequest {
	s.MonitorSourceName = &v
	return s
}

type ListIntegrationConfigsResponseBody struct {
	// data
	Data []*ListIntegrationConfigsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListIntegrationConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIntegrationConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIntegrationConfigsResponseBody) SetData(v []*ListIntegrationConfigsResponseBodyData) *ListIntegrationConfigsResponseBody {
	s.Data = v
	return s
}

func (s *ListIntegrationConfigsResponseBody) SetRequestId(v string) *ListIntegrationConfigsResponseBody {
	s.RequestId = &v
	return s
}

type ListIntegrationConfigsResponseBodyData struct {
	IntegrationConfigId    *int64  `json:"integrationConfigId,omitempty" xml:"integrationConfigId,omitempty"`
	IsReceivedEvent        *bool   `json:"isReceivedEvent,omitempty" xml:"isReceivedEvent,omitempty"`
	MonitorSourceId        *int64  `json:"monitorSourceId,omitempty" xml:"monitorSourceId,omitempty"`
	MonitorSourceName      *string `json:"monitorSourceName,omitempty" xml:"monitorSourceName,omitempty"`
	MonitorSourceShortName *string `json:"monitorSourceShortName,omitempty" xml:"monitorSourceShortName,omitempty"`
	Status                 *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListIntegrationConfigsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListIntegrationConfigsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIntegrationConfigsResponseBodyData) SetIntegrationConfigId(v int64) *ListIntegrationConfigsResponseBodyData {
	s.IntegrationConfigId = &v
	return s
}

func (s *ListIntegrationConfigsResponseBodyData) SetIsReceivedEvent(v bool) *ListIntegrationConfigsResponseBodyData {
	s.IsReceivedEvent = &v
	return s
}

func (s *ListIntegrationConfigsResponseBodyData) SetMonitorSourceId(v int64) *ListIntegrationConfigsResponseBodyData {
	s.MonitorSourceId = &v
	return s
}

func (s *ListIntegrationConfigsResponseBodyData) SetMonitorSourceName(v string) *ListIntegrationConfigsResponseBodyData {
	s.MonitorSourceName = &v
	return s
}

func (s *ListIntegrationConfigsResponseBodyData) SetMonitorSourceShortName(v string) *ListIntegrationConfigsResponseBodyData {
	s.MonitorSourceShortName = &v
	return s
}

func (s *ListIntegrationConfigsResponseBodyData) SetStatus(v string) *ListIntegrationConfigsResponseBodyData {
	s.Status = &v
	return s
}

type ListIntegrationConfigsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListIntegrationConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListIntegrationConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIntegrationConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListIntegrationConfigsResponse) SetHeaders(v map[string]*string) *ListIntegrationConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListIntegrationConfigsResponse) SetStatusCode(v int32) *ListIntegrationConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIntegrationConfigsResponse) SetBody(v *ListIntegrationConfigsResponseBody) *ListIntegrationConfigsResponse {
	s.Body = v
	return s
}

type ListMonitorSourcesRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s ListMonitorSourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMonitorSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListMonitorSourcesRequest) SetClientToken(v string) *ListMonitorSourcesRequest {
	s.ClientToken = &v
	return s
}

type ListMonitorSourcesResponseBody struct {
	// data
	Data []*ListMonitorSourcesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListMonitorSourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMonitorSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMonitorSourcesResponseBody) SetData(v []*ListMonitorSourcesResponseBodyData) *ListMonitorSourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListMonitorSourcesResponseBody) SetRequestId(v string) *ListMonitorSourcesResponseBody {
	s.RequestId = &v
	return s
}

type ListMonitorSourcesResponseBodyData struct {
	FieldKeys         []*string `json:"fieldKeys,omitempty" xml:"fieldKeys,omitempty" type:"Repeated"`
	MonitorSourceId   *int64    `json:"monitorSourceId,omitempty" xml:"monitorSourceId,omitempty"`
	MonitorSourceName *string   `json:"monitorSourceName,omitempty" xml:"monitorSourceName,omitempty"`
}

func (s ListMonitorSourcesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListMonitorSourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMonitorSourcesResponseBodyData) SetFieldKeys(v []*string) *ListMonitorSourcesResponseBodyData {
	s.FieldKeys = v
	return s
}

func (s *ListMonitorSourcesResponseBodyData) SetMonitorSourceId(v int64) *ListMonitorSourcesResponseBodyData {
	s.MonitorSourceId = &v
	return s
}

func (s *ListMonitorSourcesResponseBodyData) SetMonitorSourceName(v string) *ListMonitorSourcesResponseBodyData {
	s.MonitorSourceName = &v
	return s
}

type ListMonitorSourcesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListMonitorSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMonitorSourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMonitorSourcesResponse) GoString() string {
	return s.String()
}

func (s *ListMonitorSourcesResponse) SetHeaders(v map[string]*string) *ListMonitorSourcesResponse {
	s.Headers = v
	return s
}

func (s *ListMonitorSourcesResponse) SetStatusCode(v int32) *ListMonitorSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMonitorSourcesResponse) SetBody(v *ListMonitorSourcesResponseBody) *ListMonitorSourcesResponse {
	s.Body = v
	return s
}

type ListProblemDetailOperationsRequest struct {
	ClientToken    *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	CreateTimeSort *string `json:"createTimeSort,omitempty" xml:"createTimeSort,omitempty"`
	PageNumber     *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize       *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ProblemId      *int64  `json:"problemId,omitempty" xml:"problemId,omitempty"`
}

func (s ListProblemDetailOperationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProblemDetailOperationsRequest) GoString() string {
	return s.String()
}

func (s *ListProblemDetailOperationsRequest) SetClientToken(v string) *ListProblemDetailOperationsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListProblemDetailOperationsRequest) SetCreateTimeSort(v string) *ListProblemDetailOperationsRequest {
	s.CreateTimeSort = &v
	return s
}

func (s *ListProblemDetailOperationsRequest) SetPageNumber(v int32) *ListProblemDetailOperationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProblemDetailOperationsRequest) SetPageSize(v int32) *ListProblemDetailOperationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListProblemDetailOperationsRequest) SetProblemId(v int64) *ListProblemDetailOperationsRequest {
	s.ProblemId = &v
	return s
}

type ListProblemDetailOperationsResponseBody struct {
	// data
	Data       []*ListProblemDetailOperationsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	PageNumber *int32                                         `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32                                         `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// requestId
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TotalCount *int32  `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListProblemDetailOperationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProblemDetailOperationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProblemDetailOperationsResponseBody) SetData(v []*ListProblemDetailOperationsResponseBodyData) *ListProblemDetailOperationsResponseBody {
	s.Data = v
	return s
}

func (s *ListProblemDetailOperationsResponseBody) SetPageNumber(v int32) *ListProblemDetailOperationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListProblemDetailOperationsResponseBody) SetPageSize(v int32) *ListProblemDetailOperationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListProblemDetailOperationsResponseBody) SetRequestId(v string) *ListProblemDetailOperationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProblemDetailOperationsResponseBody) SetTotalCount(v int32) *ListProblemDetailOperationsResponseBody {
	s.TotalCount = &v
	return s
}

type ListProblemDetailOperationsResponseBodyData struct {
	Action             *string `json:"action,omitempty" xml:"action,omitempty"`
	CreateTime         *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description        *string `json:"description,omitempty" xml:"description,omitempty"`
	RelatedServiceName *string `json:"relatedServiceName,omitempty" xml:"relatedServiceName,omitempty"`
	Remark             *string `json:"remark,omitempty" xml:"remark,omitempty"`
	SnapshotData       *string `json:"snapshotData,omitempty" xml:"snapshotData,omitempty"`
	Title              *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ListProblemDetailOperationsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListProblemDetailOperationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProblemDetailOperationsResponseBodyData) SetAction(v string) *ListProblemDetailOperationsResponseBodyData {
	s.Action = &v
	return s
}

func (s *ListProblemDetailOperationsResponseBodyData) SetCreateTime(v string) *ListProblemDetailOperationsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListProblemDetailOperationsResponseBodyData) SetDescription(v string) *ListProblemDetailOperationsResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListProblemDetailOperationsResponseBodyData) SetRelatedServiceName(v string) *ListProblemDetailOperationsResponseBodyData {
	s.RelatedServiceName = &v
	return s
}

func (s *ListProblemDetailOperationsResponseBodyData) SetRemark(v string) *ListProblemDetailOperationsResponseBodyData {
	s.Remark = &v
	return s
}

func (s *ListProblemDetailOperationsResponseBodyData) SetSnapshotData(v string) *ListProblemDetailOperationsResponseBodyData {
	s.SnapshotData = &v
	return s
}

func (s *ListProblemDetailOperationsResponseBodyData) SetTitle(v string) *ListProblemDetailOperationsResponseBodyData {
	s.Title = &v
	return s
}

type ListProblemDetailOperationsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListProblemDetailOperationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProblemDetailOperationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProblemDetailOperationsResponse) GoString() string {
	return s.String()
}

func (s *ListProblemDetailOperationsResponse) SetHeaders(v map[string]*string) *ListProblemDetailOperationsResponse {
	s.Headers = v
	return s
}

func (s *ListProblemDetailOperationsResponse) SetStatusCode(v int32) *ListProblemDetailOperationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProblemDetailOperationsResponse) SetBody(v *ListProblemDetailOperationsResponseBody) *ListProblemDetailOperationsResponse {
	s.Body = v
	return s
}

type ListProblemOperationsRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	PageNumber  *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize    *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListProblemOperationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProblemOperationsRequest) GoString() string {
	return s.String()
}

func (s *ListProblemOperationsRequest) SetClientToken(v string) *ListProblemOperationsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListProblemOperationsRequest) SetPageNumber(v int32) *ListProblemOperationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProblemOperationsRequest) SetPageSize(v int32) *ListProblemOperationsRequest {
	s.PageSize = &v
	return s
}

type ListProblemOperationsResponseBody struct {
	// data
	Data       []*ListProblemOperationsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	PageNumber *int32                                   `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32                                   `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Id of the request
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TotalCount *int32  `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListProblemOperationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProblemOperationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProblemOperationsResponseBody) SetData(v []*ListProblemOperationsResponseBodyData) *ListProblemOperationsResponseBody {
	s.Data = v
	return s
}

func (s *ListProblemOperationsResponseBody) SetPageNumber(v int32) *ListProblemOperationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListProblemOperationsResponseBody) SetPageSize(v int32) *ListProblemOperationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListProblemOperationsResponseBody) SetRequestId(v string) *ListProblemOperationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProblemOperationsResponseBody) SetTotalCount(v int32) *ListProblemOperationsResponseBody {
	s.TotalCount = &v
	return s
}

type ListProblemOperationsResponseBodyData struct {
	Action             *string `json:"action,omitempty" xml:"action,omitempty"`
	CreateTime         *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description        *string `json:"description,omitempty" xml:"description,omitempty"`
	ProblemId          *int64  `json:"problemId,omitempty" xml:"problemId,omitempty"`
	ProblemName        *string `json:"problemName,omitempty" xml:"problemName,omitempty"`
	ProblemNumber      *string `json:"problemNumber,omitempty" xml:"problemNumber,omitempty"`
	RelatedServiceName *string `json:"relatedServiceName,omitempty" xml:"relatedServiceName,omitempty"`
	SnapshotData       *string `json:"snapshotData,omitempty" xml:"snapshotData,omitempty"`
	Title              *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ListProblemOperationsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListProblemOperationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProblemOperationsResponseBodyData) SetAction(v string) *ListProblemOperationsResponseBodyData {
	s.Action = &v
	return s
}

func (s *ListProblemOperationsResponseBodyData) SetCreateTime(v string) *ListProblemOperationsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListProblemOperationsResponseBodyData) SetDescription(v string) *ListProblemOperationsResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListProblemOperationsResponseBodyData) SetProblemId(v int64) *ListProblemOperationsResponseBodyData {
	s.ProblemId = &v
	return s
}

func (s *ListProblemOperationsResponseBodyData) SetProblemName(v string) *ListProblemOperationsResponseBodyData {
	s.ProblemName = &v
	return s
}

func (s *ListProblemOperationsResponseBodyData) SetProblemNumber(v string) *ListProblemOperationsResponseBodyData {
	s.ProblemNumber = &v
	return s
}

func (s *ListProblemOperationsResponseBodyData) SetRelatedServiceName(v string) *ListProblemOperationsResponseBodyData {
	s.RelatedServiceName = &v
	return s
}

func (s *ListProblemOperationsResponseBodyData) SetSnapshotData(v string) *ListProblemOperationsResponseBodyData {
	s.SnapshotData = &v
	return s
}

func (s *ListProblemOperationsResponseBodyData) SetTitle(v string) *ListProblemOperationsResponseBodyData {
	s.Title = &v
	return s
}

type ListProblemOperationsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListProblemOperationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProblemOperationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProblemOperationsResponse) GoString() string {
	return s.String()
}

func (s *ListProblemOperationsResponse) SetHeaders(v map[string]*string) *ListProblemOperationsResponse {
	s.Headers = v
	return s
}

func (s *ListProblemOperationsResponse) SetStatusCode(v int32) *ListProblemOperationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProblemOperationsResponse) SetBody(v *ListProblemOperationsResponseBody) *ListProblemOperationsResponse {
	s.Body = v
	return s
}

type ListProblemSubtotalsRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	ProblemId   *int64  `json:"problemId,omitempty" xml:"problemId,omitempty"`
}

func (s ListProblemSubtotalsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProblemSubtotalsRequest) GoString() string {
	return s.String()
}

func (s *ListProblemSubtotalsRequest) SetClientToken(v string) *ListProblemSubtotalsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListProblemSubtotalsRequest) SetProblemId(v int64) *ListProblemSubtotalsRequest {
	s.ProblemId = &v
	return s
}

type ListProblemSubtotalsResponseBody struct {
	// data
	Data []*ListProblemSubtotalsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListProblemSubtotalsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProblemSubtotalsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProblemSubtotalsResponseBody) SetData(v []*ListProblemSubtotalsResponseBodyData) *ListProblemSubtotalsResponseBody {
	s.Data = v
	return s
}

func (s *ListProblemSubtotalsResponseBody) SetRequestId(v string) *ListProblemSubtotalsResponseBody {
	s.RequestId = &v
	return s
}

type ListProblemSubtotalsResponseBodyData struct {
	CreateRamName   *string `json:"createRamName,omitempty" xml:"createRamName,omitempty"`
	CreateTime      *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreateUserId    *int64  `json:"createUserId,omitempty" xml:"createUserId,omitempty"`
	CreateUserPhone *string `json:"createUserPhone,omitempty" xml:"createUserPhone,omitempty"`
	Description     *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s ListProblemSubtotalsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListProblemSubtotalsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProblemSubtotalsResponseBodyData) SetCreateRamName(v string) *ListProblemSubtotalsResponseBodyData {
	s.CreateRamName = &v
	return s
}

func (s *ListProblemSubtotalsResponseBodyData) SetCreateTime(v string) *ListProblemSubtotalsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListProblemSubtotalsResponseBodyData) SetCreateUserId(v int64) *ListProblemSubtotalsResponseBodyData {
	s.CreateUserId = &v
	return s
}

func (s *ListProblemSubtotalsResponseBodyData) SetCreateUserPhone(v string) *ListProblemSubtotalsResponseBodyData {
	s.CreateUserPhone = &v
	return s
}

func (s *ListProblemSubtotalsResponseBodyData) SetDescription(v string) *ListProblemSubtotalsResponseBodyData {
	s.Description = &v
	return s
}

type ListProblemSubtotalsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListProblemSubtotalsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProblemSubtotalsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProblemSubtotalsResponse) GoString() string {
	return s.String()
}

func (s *ListProblemSubtotalsResponse) SetHeaders(v map[string]*string) *ListProblemSubtotalsResponse {
	s.Headers = v
	return s
}

func (s *ListProblemSubtotalsResponse) SetStatusCode(v int32) *ListProblemSubtotalsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProblemSubtotalsResponse) SetBody(v *ListProblemSubtotalsResponseBody) *ListProblemSubtotalsResponse {
	s.Body = v
	return s
}

type ListProblemTimeLinesRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	ProblemId   *int64  `json:"problemId,omitempty" xml:"problemId,omitempty"`
}

func (s ListProblemTimeLinesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProblemTimeLinesRequest) GoString() string {
	return s.String()
}

func (s *ListProblemTimeLinesRequest) SetClientToken(v string) *ListProblemTimeLinesRequest {
	s.ClientToken = &v
	return s
}

func (s *ListProblemTimeLinesRequest) SetProblemId(v int64) *ListProblemTimeLinesRequest {
	s.ProblemId = &v
	return s
}

type ListProblemTimeLinesResponseBody struct {
	Data []*ListProblemTimeLinesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListProblemTimeLinesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProblemTimeLinesResponseBody) GoString() string {
	return s.String()
}

func (s *ListProblemTimeLinesResponseBody) SetData(v []*ListProblemTimeLinesResponseBodyData) *ListProblemTimeLinesResponseBody {
	s.Data = v
	return s
}

func (s *ListProblemTimeLinesResponseBody) SetRequestId(v string) *ListProblemTimeLinesResponseBody {
	s.RequestId = &v
	return s
}

type ListProblemTimeLinesResponseBodyData struct {
	Content           *string                                               `json:"content,omitempty" xml:"content,omitempty"`
	CreateTime        *string                                               `json:"createTime,omitempty" xml:"createTime,omitempty"`
	IsKey             *bool                                                 `json:"isKey,omitempty" xml:"isKey,omitempty"`
	KeyNode           *string                                               `json:"keyNode,omitempty" xml:"keyNode,omitempty"`
	ProblemTimelineId *int64                                                `json:"problemTimelineId,omitempty" xml:"problemTimelineId,omitempty"`
	Time              *string                                               `json:"time,omitempty" xml:"time,omitempty"`
	UpdateTime        *string                                               `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	UsersInContent    []*ListProblemTimeLinesResponseBodyDataUsersInContent `json:"usersInContent,omitempty" xml:"usersInContent,omitempty" type:"Repeated"`
}

func (s ListProblemTimeLinesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListProblemTimeLinesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProblemTimeLinesResponseBodyData) SetContent(v string) *ListProblemTimeLinesResponseBodyData {
	s.Content = &v
	return s
}

func (s *ListProblemTimeLinesResponseBodyData) SetCreateTime(v string) *ListProblemTimeLinesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListProblemTimeLinesResponseBodyData) SetIsKey(v bool) *ListProblemTimeLinesResponseBodyData {
	s.IsKey = &v
	return s
}

func (s *ListProblemTimeLinesResponseBodyData) SetKeyNode(v string) *ListProblemTimeLinesResponseBodyData {
	s.KeyNode = &v
	return s
}

func (s *ListProblemTimeLinesResponseBodyData) SetProblemTimelineId(v int64) *ListProblemTimeLinesResponseBodyData {
	s.ProblemTimelineId = &v
	return s
}

func (s *ListProblemTimeLinesResponseBodyData) SetTime(v string) *ListProblemTimeLinesResponseBodyData {
	s.Time = &v
	return s
}

func (s *ListProblemTimeLinesResponseBodyData) SetUpdateTime(v string) *ListProblemTimeLinesResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListProblemTimeLinesResponseBodyData) SetUsersInContent(v []*ListProblemTimeLinesResponseBodyDataUsersInContent) *ListProblemTimeLinesResponseBodyData {
	s.UsersInContent = v
	return s
}

type ListProblemTimeLinesResponseBodyDataUsersInContent struct {
	IsValid  *int64  `json:"isValid,omitempty" xml:"isValid,omitempty"`
	UserId   *int64  `json:"userId,omitempty" xml:"userId,omitempty"`
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ListProblemTimeLinesResponseBodyDataUsersInContent) String() string {
	return tea.Prettify(s)
}

func (s ListProblemTimeLinesResponseBodyDataUsersInContent) GoString() string {
	return s.String()
}

func (s *ListProblemTimeLinesResponseBodyDataUsersInContent) SetIsValid(v int64) *ListProblemTimeLinesResponseBodyDataUsersInContent {
	s.IsValid = &v
	return s
}

func (s *ListProblemTimeLinesResponseBodyDataUsersInContent) SetUserId(v int64) *ListProblemTimeLinesResponseBodyDataUsersInContent {
	s.UserId = &v
	return s
}

func (s *ListProblemTimeLinesResponseBodyDataUsersInContent) SetUsername(v string) *ListProblemTimeLinesResponseBodyDataUsersInContent {
	s.Username = &v
	return s
}

type ListProblemTimeLinesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListProblemTimeLinesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProblemTimeLinesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProblemTimeLinesResponse) GoString() string {
	return s.String()
}

func (s *ListProblemTimeLinesResponse) SetHeaders(v map[string]*string) *ListProblemTimeLinesResponse {
	s.Headers = v
	return s
}

func (s *ListProblemTimeLinesResponse) SetStatusCode(v int32) *ListProblemTimeLinesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProblemTimeLinesResponse) SetBody(v *ListProblemTimeLinesResponseBody) *ListProblemTimeLinesResponse {
	s.Body = v
	return s
}

type ListProblemsRequest struct {
	AffectServiceId    *int64  `json:"affectServiceId,omitempty" xml:"affectServiceId,omitempty"`
	ClientToken        *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	DiscoveryEndTime   *string `json:"discoveryEndTime,omitempty" xml:"discoveryEndTime,omitempty"`
	DiscoveryStartTime *string `json:"discoveryStartTime,omitempty" xml:"discoveryStartTime,omitempty"`
	MainHandlerId      *int64  `json:"mainHandlerId,omitempty" xml:"mainHandlerId,omitempty"`
	PageNumber         *int64  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize           *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ProblemLevel       *string `json:"problemLevel,omitempty" xml:"problemLevel,omitempty"`
	ProblemStatus      *string `json:"problemStatus,omitempty" xml:"problemStatus,omitempty"`
	QueryType          *string `json:"queryType,omitempty" xml:"queryType,omitempty"`
	RepeaterId         *int64  `json:"repeaterId,omitempty" xml:"repeaterId,omitempty"`
	RestoreEndTime     *string `json:"restoreEndTime,omitempty" xml:"restoreEndTime,omitempty"`
	RestoreStartTime   *string `json:"restoreStartTime,omitempty" xml:"restoreStartTime,omitempty"`
	ServiceGroupId     *int64  `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
}

func (s ListProblemsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProblemsRequest) GoString() string {
	return s.String()
}

func (s *ListProblemsRequest) SetAffectServiceId(v int64) *ListProblemsRequest {
	s.AffectServiceId = &v
	return s
}

func (s *ListProblemsRequest) SetClientToken(v string) *ListProblemsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListProblemsRequest) SetDiscoveryEndTime(v string) *ListProblemsRequest {
	s.DiscoveryEndTime = &v
	return s
}

func (s *ListProblemsRequest) SetDiscoveryStartTime(v string) *ListProblemsRequest {
	s.DiscoveryStartTime = &v
	return s
}

func (s *ListProblemsRequest) SetMainHandlerId(v int64) *ListProblemsRequest {
	s.MainHandlerId = &v
	return s
}

func (s *ListProblemsRequest) SetPageNumber(v int64) *ListProblemsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProblemsRequest) SetPageSize(v int64) *ListProblemsRequest {
	s.PageSize = &v
	return s
}

func (s *ListProblemsRequest) SetProblemLevel(v string) *ListProblemsRequest {
	s.ProblemLevel = &v
	return s
}

func (s *ListProblemsRequest) SetProblemStatus(v string) *ListProblemsRequest {
	s.ProblemStatus = &v
	return s
}

func (s *ListProblemsRequest) SetQueryType(v string) *ListProblemsRequest {
	s.QueryType = &v
	return s
}

func (s *ListProblemsRequest) SetRepeaterId(v int64) *ListProblemsRequest {
	s.RepeaterId = &v
	return s
}

func (s *ListProblemsRequest) SetRestoreEndTime(v string) *ListProblemsRequest {
	s.RestoreEndTime = &v
	return s
}

func (s *ListProblemsRequest) SetRestoreStartTime(v string) *ListProblemsRequest {
	s.RestoreStartTime = &v
	return s
}

func (s *ListProblemsRequest) SetServiceGroupId(v int64) *ListProblemsRequest {
	s.ServiceGroupId = &v
	return s
}

type ListProblemsResponseBody struct {
	Data       []*ListProblemsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	PageNumber *int64                          `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64                          `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Id of the request
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TotalCount *int64  `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListProblemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProblemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProblemsResponseBody) SetData(v []*ListProblemsResponseBodyData) *ListProblemsResponseBody {
	s.Data = v
	return s
}

func (s *ListProblemsResponseBody) SetPageNumber(v int64) *ListProblemsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListProblemsResponseBody) SetPageSize(v int64) *ListProblemsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListProblemsResponseBody) SetRequestId(v string) *ListProblemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProblemsResponseBody) SetTotalCount(v int64) *ListProblemsResponseBody {
	s.TotalCount = &v
	return s
}

type ListProblemsResponseBodyData struct {
	AffectServices     []*ListProblemsResponseBodyDataAffectServices `json:"affectServices,omitempty" xml:"affectServices,omitempty" type:"Repeated"`
	CancelTime         *string                                       `json:"cancelTime,omitempty" xml:"cancelTime,omitempty"`
	CreateTime         *string                                       `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DiscoverTime       *string                                       `json:"discoverTime,omitempty" xml:"discoverTime,omitempty"`
	FinishTime         *string                                       `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	IncidentId         *int64                                        `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
	IsManual           *bool                                         `json:"isManual,omitempty" xml:"isManual,omitempty"`
	IsUpgrade          *bool                                         `json:"isUpgrade,omitempty" xml:"isUpgrade,omitempty"`
	MainHandlerId      *int64                                        `json:"mainHandlerId,omitempty" xml:"mainHandlerId,omitempty"`
	MainHandlerIsValid *int64                                        `json:"mainHandlerIsValid,omitempty" xml:"mainHandlerIsValid,omitempty"`
	MainHandlerName    *string                                       `json:"mainHandlerName,omitempty" xml:"mainHandlerName,omitempty"`
	ProblemId          *int64                                        `json:"problemId,omitempty" xml:"problemId,omitempty"`
	ProblemLevel       *string                                       `json:"problemLevel,omitempty" xml:"problemLevel,omitempty"`
	ProblemName        *string                                       `json:"problemName,omitempty" xml:"problemName,omitempty"`
	ProblemNumber      *string                                       `json:"problemNumber,omitempty" xml:"problemNumber,omitempty"`
	ProblemStatus      *string                                       `json:"problemStatus,omitempty" xml:"problemStatus,omitempty"`
	RecoveryTime       *string                                       `json:"recoveryTime,omitempty" xml:"recoveryTime,omitempty"`
	RelatedServiceId   *string                                       `json:"relatedServiceId,omitempty" xml:"relatedServiceId,omitempty"`
	ReplayTime         *string                                       `json:"replayTime,omitempty" xml:"replayTime,omitempty"`
	ServiceDeletedType *int32                                        `json:"serviceDeletedType,omitempty" xml:"serviceDeletedType,omitempty"`
	ServiceName        *string                                       `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	UpdateTime         *string                                       `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListProblemsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListProblemsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProblemsResponseBodyData) SetAffectServices(v []*ListProblemsResponseBodyDataAffectServices) *ListProblemsResponseBodyData {
	s.AffectServices = v
	return s
}

func (s *ListProblemsResponseBodyData) SetCancelTime(v string) *ListProblemsResponseBodyData {
	s.CancelTime = &v
	return s
}

func (s *ListProblemsResponseBodyData) SetCreateTime(v string) *ListProblemsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListProblemsResponseBodyData) SetDiscoverTime(v string) *ListProblemsResponseBodyData {
	s.DiscoverTime = &v
	return s
}

func (s *ListProblemsResponseBodyData) SetFinishTime(v string) *ListProblemsResponseBodyData {
	s.FinishTime = &v
	return s
}

func (s *ListProblemsResponseBodyData) SetIncidentId(v int64) *ListProblemsResponseBodyData {
	s.IncidentId = &v
	return s
}

func (s *ListProblemsResponseBodyData) SetIsManual(v bool) *ListProblemsResponseBodyData {
	s.IsManual = &v
	return s
}

func (s *ListProblemsResponseBodyData) SetIsUpgrade(v bool) *ListProblemsResponseBodyData {
	s.IsUpgrade = &v
	return s
}

func (s *ListProblemsResponseBodyData) SetMainHandlerId(v int64) *ListProblemsResponseBodyData {
	s.MainHandlerId = &v
	return s
}

func (s *ListProblemsResponseBodyData) SetMainHandlerIsValid(v int64) *ListProblemsResponseBodyData {
	s.MainHandlerIsValid = &v
	return s
}

func (s *ListProblemsResponseBodyData) SetMainHandlerName(v string) *ListProblemsResponseBodyData {
	s.MainHandlerName = &v
	return s
}

func (s *ListProblemsResponseBodyData) SetProblemId(v int64) *ListProblemsResponseBodyData {
	s.ProblemId = &v
	return s
}

func (s *ListProblemsResponseBodyData) SetProblemLevel(v string) *ListProblemsResponseBodyData {
	s.ProblemLevel = &v
	return s
}

func (s *ListProblemsResponseBodyData) SetProblemName(v string) *ListProblemsResponseBodyData {
	s.ProblemName = &v
	return s
}

func (s *ListProblemsResponseBodyData) SetProblemNumber(v string) *ListProblemsResponseBodyData {
	s.ProblemNumber = &v
	return s
}

func (s *ListProblemsResponseBodyData) SetProblemStatus(v string) *ListProblemsResponseBodyData {
	s.ProblemStatus = &v
	return s
}

func (s *ListProblemsResponseBodyData) SetRecoveryTime(v string) *ListProblemsResponseBodyData {
	s.RecoveryTime = &v
	return s
}

func (s *ListProblemsResponseBodyData) SetRelatedServiceId(v string) *ListProblemsResponseBodyData {
	s.RelatedServiceId = &v
	return s
}

func (s *ListProblemsResponseBodyData) SetReplayTime(v string) *ListProblemsResponseBodyData {
	s.ReplayTime = &v
	return s
}

func (s *ListProblemsResponseBodyData) SetServiceDeletedType(v int32) *ListProblemsResponseBodyData {
	s.ServiceDeletedType = &v
	return s
}

func (s *ListProblemsResponseBodyData) SetServiceName(v string) *ListProblemsResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *ListProblemsResponseBodyData) SetUpdateTime(v string) *ListProblemsResponseBodyData {
	s.UpdateTime = &v
	return s
}

type ListProblemsResponseBodyDataAffectServices struct {
	ServiceDescription *string `json:"serviceDescription,omitempty" xml:"serviceDescription,omitempty"`
	ServiceId          *int64  `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	ServiceName        *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	UpdateTime         *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListProblemsResponseBodyDataAffectServices) String() string {
	return tea.Prettify(s)
}

func (s ListProblemsResponseBodyDataAffectServices) GoString() string {
	return s.String()
}

func (s *ListProblemsResponseBodyDataAffectServices) SetServiceDescription(v string) *ListProblemsResponseBodyDataAffectServices {
	s.ServiceDescription = &v
	return s
}

func (s *ListProblemsResponseBodyDataAffectServices) SetServiceId(v int64) *ListProblemsResponseBodyDataAffectServices {
	s.ServiceId = &v
	return s
}

func (s *ListProblemsResponseBodyDataAffectServices) SetServiceName(v string) *ListProblemsResponseBodyDataAffectServices {
	s.ServiceName = &v
	return s
}

func (s *ListProblemsResponseBodyDataAffectServices) SetUpdateTime(v string) *ListProblemsResponseBodyDataAffectServices {
	s.UpdateTime = &v
	return s
}

type ListProblemsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListProblemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProblemsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProblemsResponse) GoString() string {
	return s.String()
}

func (s *ListProblemsResponse) SetHeaders(v map[string]*string) *ListProblemsResponse {
	s.Headers = v
	return s
}

func (s *ListProblemsResponse) SetStatusCode(v int32) *ListProblemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProblemsResponse) SetBody(v *ListProblemsResponseBody) *ListProblemsResponse {
	s.Body = v
	return s
}

type ListRouteRulesRequest struct {
	ClientToken               *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	NotFilterRouteRuleDeleted *bool   `json:"notFilterRouteRuleDeleted,omitempty" xml:"notFilterRouteRuleDeleted,omitempty"`
	PageNumber                *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize                  *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	RouteType                 *int64  `json:"routeType,omitempty" xml:"routeType,omitempty"`
	RuleName                  []byte  `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	ServiceName               []byte  `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
}

func (s ListRouteRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRouteRulesRequest) GoString() string {
	return s.String()
}

func (s *ListRouteRulesRequest) SetClientToken(v string) *ListRouteRulesRequest {
	s.ClientToken = &v
	return s
}

func (s *ListRouteRulesRequest) SetNotFilterRouteRuleDeleted(v bool) *ListRouteRulesRequest {
	s.NotFilterRouteRuleDeleted = &v
	return s
}

func (s *ListRouteRulesRequest) SetPageNumber(v int32) *ListRouteRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRouteRulesRequest) SetPageSize(v int32) *ListRouteRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListRouteRulesRequest) SetRouteType(v int64) *ListRouteRulesRequest {
	s.RouteType = &v
	return s
}

func (s *ListRouteRulesRequest) SetRuleName(v []byte) *ListRouteRulesRequest {
	s.RuleName = v
	return s
}

func (s *ListRouteRulesRequest) SetServiceName(v []byte) *ListRouteRulesRequest {
	s.ServiceName = v
	return s
}

type ListRouteRulesResponseBody struct {
	Data       []*ListRouteRulesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	PageNumber *int64                            `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64                            `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	RequestId  *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TotalCount *int64                            `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListRouteRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRouteRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRouteRulesResponseBody) SetData(v []*ListRouteRulesResponseBodyData) *ListRouteRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListRouteRulesResponseBody) SetPageNumber(v int64) *ListRouteRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRouteRulesResponseBody) SetPageSize(v int64) *ListRouteRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRouteRulesResponseBody) SetRequestId(v string) *ListRouteRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRouteRulesResponseBody) SetTotalCount(v int64) *ListRouteRulesResponseBody {
	s.TotalCount = &v
	return s
}

type ListRouteRulesResponseBodyData struct {
	AssignObjectId       *int64  `json:"assignObjectId,omitempty" xml:"assignObjectId,omitempty"`
	AssignObjectType     *string `json:"assignObjectType,omitempty" xml:"assignObjectType,omitempty"`
	CreateTime           *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Effection            *string `json:"effection,omitempty" xml:"effection,omitempty"`
	EnableStatus         *string `json:"enableStatus,omitempty" xml:"enableStatus,omitempty"`
	IncidentLevel        *string `json:"incidentLevel,omitempty" xml:"incidentLevel,omitempty"`
	IsValid              *int32  `json:"isValid,omitempty" xml:"isValid,omitempty"`
	MatchCount           *int64  `json:"matchCount,omitempty" xml:"matchCount,omitempty"`
	MonitorSourceNames   *string `json:"monitorSourceNames,omitempty" xml:"monitorSourceNames,omitempty"`
	RelServiceDeleteType *int32  `json:"relServiceDeleteType,omitempty" xml:"relServiceDeleteType,omitempty"`
	RelatedServiceId     *int64  `json:"relatedServiceId,omitempty" xml:"relatedServiceId,omitempty"`
	RelatedServiceName   *string `json:"relatedServiceName,omitempty" xml:"relatedServiceName,omitempty"`
	RouteRuleId          *int64  `json:"routeRuleId,omitempty" xml:"routeRuleId,omitempty"`
	RouteType            *string `json:"routeType,omitempty" xml:"routeType,omitempty"`
	RuleName             *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	TenantRamId          *int64  `json:"tenantRamId,omitempty" xml:"tenantRamId,omitempty"`
	TimeWindow           *int64  `json:"timeWindow,omitempty" xml:"timeWindow,omitempty"`
	TimeWindowUnit       *int64  `json:"timeWindowUnit,omitempty" xml:"timeWindowUnit,omitempty"`
	UpdateTime           *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListRouteRulesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListRouteRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRouteRulesResponseBodyData) SetAssignObjectId(v int64) *ListRouteRulesResponseBodyData {
	s.AssignObjectId = &v
	return s
}

func (s *ListRouteRulesResponseBodyData) SetAssignObjectType(v string) *ListRouteRulesResponseBodyData {
	s.AssignObjectType = &v
	return s
}

func (s *ListRouteRulesResponseBodyData) SetCreateTime(v string) *ListRouteRulesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListRouteRulesResponseBodyData) SetEffection(v string) *ListRouteRulesResponseBodyData {
	s.Effection = &v
	return s
}

func (s *ListRouteRulesResponseBodyData) SetEnableStatus(v string) *ListRouteRulesResponseBodyData {
	s.EnableStatus = &v
	return s
}

func (s *ListRouteRulesResponseBodyData) SetIncidentLevel(v string) *ListRouteRulesResponseBodyData {
	s.IncidentLevel = &v
	return s
}

func (s *ListRouteRulesResponseBodyData) SetIsValid(v int32) *ListRouteRulesResponseBodyData {
	s.IsValid = &v
	return s
}

func (s *ListRouteRulesResponseBodyData) SetMatchCount(v int64) *ListRouteRulesResponseBodyData {
	s.MatchCount = &v
	return s
}

func (s *ListRouteRulesResponseBodyData) SetMonitorSourceNames(v string) *ListRouteRulesResponseBodyData {
	s.MonitorSourceNames = &v
	return s
}

func (s *ListRouteRulesResponseBodyData) SetRelServiceDeleteType(v int32) *ListRouteRulesResponseBodyData {
	s.RelServiceDeleteType = &v
	return s
}

func (s *ListRouteRulesResponseBodyData) SetRelatedServiceId(v int64) *ListRouteRulesResponseBodyData {
	s.RelatedServiceId = &v
	return s
}

func (s *ListRouteRulesResponseBodyData) SetRelatedServiceName(v string) *ListRouteRulesResponseBodyData {
	s.RelatedServiceName = &v
	return s
}

func (s *ListRouteRulesResponseBodyData) SetRouteRuleId(v int64) *ListRouteRulesResponseBodyData {
	s.RouteRuleId = &v
	return s
}

func (s *ListRouteRulesResponseBodyData) SetRouteType(v string) *ListRouteRulesResponseBodyData {
	s.RouteType = &v
	return s
}

func (s *ListRouteRulesResponseBodyData) SetRuleName(v string) *ListRouteRulesResponseBodyData {
	s.RuleName = &v
	return s
}

func (s *ListRouteRulesResponseBodyData) SetTenantRamId(v int64) *ListRouteRulesResponseBodyData {
	s.TenantRamId = &v
	return s
}

func (s *ListRouteRulesResponseBodyData) SetTimeWindow(v int64) *ListRouteRulesResponseBodyData {
	s.TimeWindow = &v
	return s
}

func (s *ListRouteRulesResponseBodyData) SetTimeWindowUnit(v int64) *ListRouteRulesResponseBodyData {
	s.TimeWindowUnit = &v
	return s
}

func (s *ListRouteRulesResponseBodyData) SetUpdateTime(v string) *ListRouteRulesResponseBodyData {
	s.UpdateTime = &v
	return s
}

type ListRouteRulesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRouteRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRouteRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRouteRulesResponse) GoString() string {
	return s.String()
}

func (s *ListRouteRulesResponse) SetHeaders(v map[string]*string) *ListRouteRulesResponse {
	s.Headers = v
	return s
}

func (s *ListRouteRulesResponse) SetStatusCode(v int32) *ListRouteRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRouteRulesResponse) SetBody(v *ListRouteRulesResponseBody) *ListRouteRulesResponse {
	s.Body = v
	return s
}

type ListRouteRulesByAssignWhoIdRequest struct {
	AssignWhoId   *int64 `json:"assignWhoId,omitempty" xml:"assignWhoId,omitempty"`
	AssignWhoType *int64 `json:"assignWhoType,omitempty" xml:"assignWhoType,omitempty"`
}

func (s ListRouteRulesByAssignWhoIdRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRouteRulesByAssignWhoIdRequest) GoString() string {
	return s.String()
}

func (s *ListRouteRulesByAssignWhoIdRequest) SetAssignWhoId(v int64) *ListRouteRulesByAssignWhoIdRequest {
	s.AssignWhoId = &v
	return s
}

func (s *ListRouteRulesByAssignWhoIdRequest) SetAssignWhoType(v int64) *ListRouteRulesByAssignWhoIdRequest {
	s.AssignWhoType = &v
	return s
}

type ListRouteRulesByAssignWhoIdResponseBody struct {
	Data      []*ListRouteRulesByAssignWhoIdResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	RequestId *string                                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListRouteRulesByAssignWhoIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRouteRulesByAssignWhoIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListRouteRulesByAssignWhoIdResponseBody) SetData(v []*ListRouteRulesByAssignWhoIdResponseBodyData) *ListRouteRulesByAssignWhoIdResponseBody {
	s.Data = v
	return s
}

func (s *ListRouteRulesByAssignWhoIdResponseBody) SetRequestId(v string) *ListRouteRulesByAssignWhoIdResponseBody {
	s.RequestId = &v
	return s
}

type ListRouteRulesByAssignWhoIdResponseBodyData struct {
	Id          *int64  `json:"id,omitempty" xml:"id,omitempty"`
	RuleName    *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	TenantRamId *int64  `json:"tenantRamId,omitempty" xml:"tenantRamId,omitempty"`
}

func (s ListRouteRulesByAssignWhoIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListRouteRulesByAssignWhoIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRouteRulesByAssignWhoIdResponseBodyData) SetId(v int64) *ListRouteRulesByAssignWhoIdResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListRouteRulesByAssignWhoIdResponseBodyData) SetRuleName(v string) *ListRouteRulesByAssignWhoIdResponseBodyData {
	s.RuleName = &v
	return s
}

func (s *ListRouteRulesByAssignWhoIdResponseBodyData) SetTenantRamId(v int64) *ListRouteRulesByAssignWhoIdResponseBodyData {
	s.TenantRamId = &v
	return s
}

type ListRouteRulesByAssignWhoIdResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRouteRulesByAssignWhoIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRouteRulesByAssignWhoIdResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRouteRulesByAssignWhoIdResponse) GoString() string {
	return s.String()
}

func (s *ListRouteRulesByAssignWhoIdResponse) SetHeaders(v map[string]*string) *ListRouteRulesByAssignWhoIdResponse {
	s.Headers = v
	return s
}

func (s *ListRouteRulesByAssignWhoIdResponse) SetStatusCode(v int32) *ListRouteRulesByAssignWhoIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRouteRulesByAssignWhoIdResponse) SetBody(v *ListRouteRulesByAssignWhoIdResponseBody) *ListRouteRulesByAssignWhoIdResponse {
	s.Body = v
	return s
}

type ListRouteRulesByServiceResponseBody struct {
	Data       []*ListRouteRulesByServiceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	RequestId  *string                                    `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TotalCount *int32                                     `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListRouteRulesByServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRouteRulesByServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ListRouteRulesByServiceResponseBody) SetData(v []*ListRouteRulesByServiceResponseBodyData) *ListRouteRulesByServiceResponseBody {
	s.Data = v
	return s
}

func (s *ListRouteRulesByServiceResponseBody) SetRequestId(v string) *ListRouteRulesByServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRouteRulesByServiceResponseBody) SetTotalCount(v int32) *ListRouteRulesByServiceResponseBody {
	s.TotalCount = &v
	return s
}

type ListRouteRulesByServiceResponseBodyData struct {
	Id       *int32  `json:"id,omitempty" xml:"id,omitempty"`
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
}

func (s ListRouteRulesByServiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListRouteRulesByServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRouteRulesByServiceResponseBodyData) SetId(v int32) *ListRouteRulesByServiceResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListRouteRulesByServiceResponseBodyData) SetRuleName(v string) *ListRouteRulesByServiceResponseBodyData {
	s.RuleName = &v
	return s
}

type ListRouteRulesByServiceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRouteRulesByServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRouteRulesByServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRouteRulesByServiceResponse) GoString() string {
	return s.String()
}

func (s *ListRouteRulesByServiceResponse) SetHeaders(v map[string]*string) *ListRouteRulesByServiceResponse {
	s.Headers = v
	return s
}

func (s *ListRouteRulesByServiceResponse) SetStatusCode(v int32) *ListRouteRulesByServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRouteRulesByServiceResponse) SetBody(v *ListRouteRulesByServiceResponseBody) *ListRouteRulesByServiceResponse {
	s.Body = v
	return s
}

type ListServiceGroupMonitorSourceTemplatesRequest struct {
	ClientToken    *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	RequestId      *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ServiceGroupId *int64  `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
}

func (s ListServiceGroupMonitorSourceTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceGroupMonitorSourceTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListServiceGroupMonitorSourceTemplatesRequest) SetClientToken(v string) *ListServiceGroupMonitorSourceTemplatesRequest {
	s.ClientToken = &v
	return s
}

func (s *ListServiceGroupMonitorSourceTemplatesRequest) SetRequestId(v string) *ListServiceGroupMonitorSourceTemplatesRequest {
	s.RequestId = &v
	return s
}

func (s *ListServiceGroupMonitorSourceTemplatesRequest) SetServiceGroupId(v int64) *ListServiceGroupMonitorSourceTemplatesRequest {
	s.ServiceGroupId = &v
	return s
}

type ListServiceGroupMonitorSourceTemplatesResponseBody struct {
	Data []*ListServiceGroupMonitorSourceTemplatesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListServiceGroupMonitorSourceTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceGroupMonitorSourceTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceGroupMonitorSourceTemplatesResponseBody) SetData(v []*ListServiceGroupMonitorSourceTemplatesResponseBodyData) *ListServiceGroupMonitorSourceTemplatesResponseBody {
	s.Data = v
	return s
}

func (s *ListServiceGroupMonitorSourceTemplatesResponseBody) SetRequestId(v string) *ListServiceGroupMonitorSourceTemplatesResponseBody {
	s.RequestId = &v
	return s
}

type ListServiceGroupMonitorSourceTemplatesResponseBodyData struct {
	Fields            []*string `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	MonitorSourceId   *int64    `json:"monitorSourceId,omitempty" xml:"monitorSourceId,omitempty"`
	MonitorSourceName *string   `json:"monitorSourceName,omitempty" xml:"monitorSourceName,omitempty"`
	TemplateContent   *string   `json:"templateContent,omitempty" xml:"templateContent,omitempty"`
	TemplateId        *int64    `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s ListServiceGroupMonitorSourceTemplatesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListServiceGroupMonitorSourceTemplatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListServiceGroupMonitorSourceTemplatesResponseBodyData) SetFields(v []*string) *ListServiceGroupMonitorSourceTemplatesResponseBodyData {
	s.Fields = v
	return s
}

func (s *ListServiceGroupMonitorSourceTemplatesResponseBodyData) SetMonitorSourceId(v int64) *ListServiceGroupMonitorSourceTemplatesResponseBodyData {
	s.MonitorSourceId = &v
	return s
}

func (s *ListServiceGroupMonitorSourceTemplatesResponseBodyData) SetMonitorSourceName(v string) *ListServiceGroupMonitorSourceTemplatesResponseBodyData {
	s.MonitorSourceName = &v
	return s
}

func (s *ListServiceGroupMonitorSourceTemplatesResponseBodyData) SetTemplateContent(v string) *ListServiceGroupMonitorSourceTemplatesResponseBodyData {
	s.TemplateContent = &v
	return s
}

func (s *ListServiceGroupMonitorSourceTemplatesResponseBodyData) SetTemplateId(v int64) *ListServiceGroupMonitorSourceTemplatesResponseBodyData {
	s.TemplateId = &v
	return s
}

type ListServiceGroupMonitorSourceTemplatesResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListServiceGroupMonitorSourceTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListServiceGroupMonitorSourceTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceGroupMonitorSourceTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListServiceGroupMonitorSourceTemplatesResponse) SetHeaders(v map[string]*string) *ListServiceGroupMonitorSourceTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListServiceGroupMonitorSourceTemplatesResponse) SetStatusCode(v int32) *ListServiceGroupMonitorSourceTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceGroupMonitorSourceTemplatesResponse) SetBody(v *ListServiceGroupMonitorSourceTemplatesResponseBody) *ListServiceGroupMonitorSourceTemplatesResponse {
	s.Body = v
	return s
}

type ListServiceGroupsRequest struct {
	ClientToken           *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	IsScheduled           *bool   `json:"isScheduled,omitempty" xml:"isScheduled,omitempty"`
	OrderByScheduleStatus *bool   `json:"orderByScheduleStatus,omitempty" xml:"orderByScheduleStatus,omitempty"`
	PageNumber            *int64  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize              *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	QueryName             *string `json:"queryName,omitempty" xml:"queryName,omitempty"`
	QueryType             *string `json:"queryType,omitempty" xml:"queryType,omitempty"`
	UserId                *int64  `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListServiceGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceGroupsRequest) SetClientToken(v string) *ListServiceGroupsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListServiceGroupsRequest) SetIsScheduled(v bool) *ListServiceGroupsRequest {
	s.IsScheduled = &v
	return s
}

func (s *ListServiceGroupsRequest) SetOrderByScheduleStatus(v bool) *ListServiceGroupsRequest {
	s.OrderByScheduleStatus = &v
	return s
}

func (s *ListServiceGroupsRequest) SetPageNumber(v int64) *ListServiceGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListServiceGroupsRequest) SetPageSize(v int64) *ListServiceGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListServiceGroupsRequest) SetQueryName(v string) *ListServiceGroupsRequest {
	s.QueryName = &v
	return s
}

func (s *ListServiceGroupsRequest) SetQueryType(v string) *ListServiceGroupsRequest {
	s.QueryType = &v
	return s
}

func (s *ListServiceGroupsRequest) SetUserId(v int64) *ListServiceGroupsRequest {
	s.UserId = &v
	return s
}

type ListServiceGroupsResponseBody struct {
	Data       []*ListServiceGroupsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	PageNumber *int64                               `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64                               `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Id of the request
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TotalCount *int64  `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListServiceGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceGroupsResponseBody) SetData(v []*ListServiceGroupsResponseBodyData) *ListServiceGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListServiceGroupsResponseBody) SetPageNumber(v int64) *ListServiceGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListServiceGroupsResponseBody) SetPageSize(v int64) *ListServiceGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListServiceGroupsResponseBody) SetRequestId(v string) *ListServiceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceGroupsResponseBody) SetTotalCount(v int64) *ListServiceGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListServiceGroupsResponseBodyData struct {
	EnableWebhook           *string                                   `json:"enableWebhook,omitempty" xml:"enableWebhook,omitempty"`
	IsScheduled             *bool                                     `json:"isScheduled,omitempty" xml:"isScheduled,omitempty"`
	ServiceGroupDescription *string                                   `json:"serviceGroupDescription,omitempty" xml:"serviceGroupDescription,omitempty"`
	ServiceGroupId          *int64                                    `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
	ServiceGroupName        *string                                   `json:"serviceGroupName,omitempty" xml:"serviceGroupName,omitempty"`
	UpdateTime              *string                                   `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	Users                   []*ListServiceGroupsResponseBodyDataUsers `json:"users,omitempty" xml:"users,omitempty" type:"Repeated"`
	WebhookLink             *string                                   `json:"webhookLink,omitempty" xml:"webhookLink,omitempty"`
	WebhookType             *string                                   `json:"webhookType,omitempty" xml:"webhookType,omitempty"`
}

func (s ListServiceGroupsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListServiceGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListServiceGroupsResponseBodyData) SetEnableWebhook(v string) *ListServiceGroupsResponseBodyData {
	s.EnableWebhook = &v
	return s
}

func (s *ListServiceGroupsResponseBodyData) SetIsScheduled(v bool) *ListServiceGroupsResponseBodyData {
	s.IsScheduled = &v
	return s
}

func (s *ListServiceGroupsResponseBodyData) SetServiceGroupDescription(v string) *ListServiceGroupsResponseBodyData {
	s.ServiceGroupDescription = &v
	return s
}

func (s *ListServiceGroupsResponseBodyData) SetServiceGroupId(v int64) *ListServiceGroupsResponseBodyData {
	s.ServiceGroupId = &v
	return s
}

func (s *ListServiceGroupsResponseBodyData) SetServiceGroupName(v string) *ListServiceGroupsResponseBodyData {
	s.ServiceGroupName = &v
	return s
}

func (s *ListServiceGroupsResponseBodyData) SetUpdateTime(v string) *ListServiceGroupsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListServiceGroupsResponseBodyData) SetUsers(v []*ListServiceGroupsResponseBodyDataUsers) *ListServiceGroupsResponseBodyData {
	s.Users = v
	return s
}

func (s *ListServiceGroupsResponseBodyData) SetWebhookLink(v string) *ListServiceGroupsResponseBodyData {
	s.WebhookLink = &v
	return s
}

func (s *ListServiceGroupsResponseBodyData) SetWebhookType(v string) *ListServiceGroupsResponseBodyData {
	s.WebhookType = &v
	return s
}

type ListServiceGroupsResponseBodyDataUsers struct {
	Email          *string `json:"email,omitempty" xml:"email,omitempty"`
	IsRelated      *int32  `json:"isRelated,omitempty" xml:"isRelated,omitempty"`
	Phone          *string `json:"phone,omitempty" xml:"phone,omitempty"`
	ServiceGroupId *int64  `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
	UserId         *int64  `json:"userId,omitempty" xml:"userId,omitempty"`
	UserName       *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s ListServiceGroupsResponseBodyDataUsers) String() string {
	return tea.Prettify(s)
}

func (s ListServiceGroupsResponseBodyDataUsers) GoString() string {
	return s.String()
}

func (s *ListServiceGroupsResponseBodyDataUsers) SetEmail(v string) *ListServiceGroupsResponseBodyDataUsers {
	s.Email = &v
	return s
}

func (s *ListServiceGroupsResponseBodyDataUsers) SetIsRelated(v int32) *ListServiceGroupsResponseBodyDataUsers {
	s.IsRelated = &v
	return s
}

func (s *ListServiceGroupsResponseBodyDataUsers) SetPhone(v string) *ListServiceGroupsResponseBodyDataUsers {
	s.Phone = &v
	return s
}

func (s *ListServiceGroupsResponseBodyDataUsers) SetServiceGroupId(v int64) *ListServiceGroupsResponseBodyDataUsers {
	s.ServiceGroupId = &v
	return s
}

func (s *ListServiceGroupsResponseBodyDataUsers) SetUserId(v int64) *ListServiceGroupsResponseBodyDataUsers {
	s.UserId = &v
	return s
}

func (s *ListServiceGroupsResponseBodyDataUsers) SetUserName(v string) *ListServiceGroupsResponseBodyDataUsers {
	s.UserName = &v
	return s
}

type ListServiceGroupsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListServiceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListServiceGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListServiceGroupsResponse) SetHeaders(v map[string]*string) *ListServiceGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListServiceGroupsResponse) SetStatusCode(v int32) *ListServiceGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceGroupsResponse) SetBody(v *ListServiceGroupsResponseBody) *ListServiceGroupsResponse {
	s.Body = v
	return s
}

type ListServiceGroupsByUserIdResponseBody struct {
	Data      *ListServiceGroupsByUserIdResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	RequestId *string                                    `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListServiceGroupsByUserIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceGroupsByUserIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceGroupsByUserIdResponseBody) SetData(v *ListServiceGroupsByUserIdResponseBodyData) *ListServiceGroupsByUserIdResponseBody {
	s.Data = v
	return s
}

func (s *ListServiceGroupsByUserIdResponseBody) SetRequestId(v string) *ListServiceGroupsByUserIdResponseBody {
	s.RequestId = &v
	return s
}

type ListServiceGroupsByUserIdResponseBodyData struct {
	IsScheduled      *bool   `json:"isScheduled,omitempty" xml:"isScheduled,omitempty"`
	ServiceGroupId   *int64  `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
	ServiceGroupName *string `json:"serviceGroupName,omitempty" xml:"serviceGroupName,omitempty"`
}

func (s ListServiceGroupsByUserIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListServiceGroupsByUserIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListServiceGroupsByUserIdResponseBodyData) SetIsScheduled(v bool) *ListServiceGroupsByUserIdResponseBodyData {
	s.IsScheduled = &v
	return s
}

func (s *ListServiceGroupsByUserIdResponseBodyData) SetServiceGroupId(v int64) *ListServiceGroupsByUserIdResponseBodyData {
	s.ServiceGroupId = &v
	return s
}

func (s *ListServiceGroupsByUserIdResponseBodyData) SetServiceGroupName(v string) *ListServiceGroupsByUserIdResponseBodyData {
	s.ServiceGroupName = &v
	return s
}

type ListServiceGroupsByUserIdResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListServiceGroupsByUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListServiceGroupsByUserIdResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceGroupsByUserIdResponse) GoString() string {
	return s.String()
}

func (s *ListServiceGroupsByUserIdResponse) SetHeaders(v map[string]*string) *ListServiceGroupsByUserIdResponse {
	s.Headers = v
	return s
}

func (s *ListServiceGroupsByUserIdResponse) SetStatusCode(v int32) *ListServiceGroupsByUserIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceGroupsByUserIdResponse) SetBody(v *ListServiceGroupsByUserIdResponseBody) *ListServiceGroupsByUserIdResponse {
	s.Body = v
	return s
}

type ListServicesRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	PageNumber  *int64  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize    *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
}

func (s ListServicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServicesRequest) GoString() string {
	return s.String()
}

func (s *ListServicesRequest) SetClientToken(v string) *ListServicesRequest {
	s.ClientToken = &v
	return s
}

func (s *ListServicesRequest) SetPageNumber(v int64) *ListServicesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListServicesRequest) SetPageSize(v int64) *ListServicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListServicesRequest) SetServiceName(v string) *ListServicesRequest {
	s.ServiceName = &v
	return s
}

type ListServicesResponseBody struct {
	Data       []*ListServicesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	PageNumber *int64                          `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64                          `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Id of the request
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TotalCount *int64  `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListServicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBody) SetData(v []*ListServicesResponseBodyData) *ListServicesResponseBody {
	s.Data = v
	return s
}

func (s *ListServicesResponseBody) SetPageNumber(v int64) *ListServicesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListServicesResponseBody) SetPageSize(v int64) *ListServicesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListServicesResponseBody) SetRequestId(v string) *ListServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServicesResponseBody) SetTotalCount(v int64) *ListServicesResponseBody {
	s.TotalCount = &v
	return s
}

type ListServicesResponseBodyData struct {
	IsValid            *int32  `json:"isValid,omitempty" xml:"isValid,omitempty"`
	ServiceDescription *string `json:"serviceDescription,omitempty" xml:"serviceDescription,omitempty"`
	ServiceId          *int64  `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	ServiceName        *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	UpdateTime         *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListServicesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBodyData) SetIsValid(v int32) *ListServicesResponseBodyData {
	s.IsValid = &v
	return s
}

func (s *ListServicesResponseBodyData) SetServiceDescription(v string) *ListServicesResponseBodyData {
	s.ServiceDescription = &v
	return s
}

func (s *ListServicesResponseBodyData) SetServiceId(v int64) *ListServicesResponseBodyData {
	s.ServiceId = &v
	return s
}

func (s *ListServicesResponseBodyData) SetServiceName(v string) *ListServicesResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *ListServicesResponseBodyData) SetUpdateTime(v string) *ListServicesResponseBodyData {
	s.UpdateTime = &v
	return s
}

type ListServicesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListServicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponse) GoString() string {
	return s.String()
}

func (s *ListServicesResponse) SetHeaders(v map[string]*string) *ListServicesResponse {
	s.Headers = v
	return s
}

func (s *ListServicesResponse) SetStatusCode(v int32) *ListServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServicesResponse) SetBody(v *ListServicesResponseBody) *ListServicesResponse {
	s.Body = v
	return s
}

type ListSourceEventsRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// 2020-09-18 13:00:00
	EndTime      *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	InstanceId   *int64  `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	PageNumber   *int64  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize     *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	StartRowKey  *string `json:"startRowKey,omitempty" xml:"startRowKey,omitempty"`
	// 2020-09-10 13:00:00
	StartTime  *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	StopRowKey *string `json:"stopRowKey,omitempty" xml:"stopRowKey,omitempty"`
}

func (s ListSourceEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSourceEventsRequest) GoString() string {
	return s.String()
}

func (s *ListSourceEventsRequest) SetClientToken(v string) *ListSourceEventsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListSourceEventsRequest) SetEndTime(v string) *ListSourceEventsRequest {
	s.EndTime = &v
	return s
}

func (s *ListSourceEventsRequest) SetInstanceId(v int64) *ListSourceEventsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListSourceEventsRequest) SetInstanceType(v string) *ListSourceEventsRequest {
	s.InstanceType = &v
	return s
}

func (s *ListSourceEventsRequest) SetPageNumber(v int64) *ListSourceEventsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSourceEventsRequest) SetPageSize(v int64) *ListSourceEventsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSourceEventsRequest) SetStartRowKey(v string) *ListSourceEventsRequest {
	s.StartRowKey = &v
	return s
}

func (s *ListSourceEventsRequest) SetStartTime(v string) *ListSourceEventsRequest {
	s.StartTime = &v
	return s
}

func (s *ListSourceEventsRequest) SetStopRowKey(v string) *ListSourceEventsRequest {
	s.StopRowKey = &v
	return s
}

type ListSourceEventsResponseBody struct {
	Data []*ListSourceEventsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// firstRowKey
	FirstRowKey *string `json:"firstRowKey,omitempty" xml:"firstRowKey,omitempty"`
	// lastRowKey
	LastRowKey *string `json:"lastRowKey,omitempty" xml:"lastRowKey,omitempty"`
	PageNumber *int64  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TotalCount *int64  `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListSourceEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSourceEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSourceEventsResponseBody) SetData(v []*ListSourceEventsResponseBodyData) *ListSourceEventsResponseBody {
	s.Data = v
	return s
}

func (s *ListSourceEventsResponseBody) SetFirstRowKey(v string) *ListSourceEventsResponseBody {
	s.FirstRowKey = &v
	return s
}

func (s *ListSourceEventsResponseBody) SetLastRowKey(v string) *ListSourceEventsResponseBody {
	s.LastRowKey = &v
	return s
}

func (s *ListSourceEventsResponseBody) SetPageNumber(v int64) *ListSourceEventsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSourceEventsResponseBody) SetPageSize(v int64) *ListSourceEventsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSourceEventsResponseBody) SetRequestId(v string) *ListSourceEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSourceEventsResponseBody) SetTotalCount(v int64) *ListSourceEventsResponseBody {
	s.TotalCount = &v
	return s
}

type ListSourceEventsResponseBodyData struct {
	EventJson         *string `json:"eventJson,omitempty" xml:"eventJson,omitempty"`
	EventTime         *string `json:"eventTime,omitempty" xml:"eventTime,omitempty"`
	InstanceId        *int64  `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	InstanceType      *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	MonitorSourceId   *int64  `json:"monitorSourceId,omitempty" xml:"monitorSourceId,omitempty"`
	MonitorSourceName *string `json:"monitorSourceName,omitempty" xml:"monitorSourceName,omitempty"`
	RouteRuleId       *int64  `json:"routeRuleId,omitempty" xml:"routeRuleId,omitempty"`
	TenantRamId       *int64  `json:"tenantRamId,omitempty" xml:"tenantRamId,omitempty"`
}

func (s ListSourceEventsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListSourceEventsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSourceEventsResponseBodyData) SetEventJson(v string) *ListSourceEventsResponseBodyData {
	s.EventJson = &v
	return s
}

func (s *ListSourceEventsResponseBodyData) SetEventTime(v string) *ListSourceEventsResponseBodyData {
	s.EventTime = &v
	return s
}

func (s *ListSourceEventsResponseBodyData) SetInstanceId(v int64) *ListSourceEventsResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListSourceEventsResponseBodyData) SetInstanceType(v string) *ListSourceEventsResponseBodyData {
	s.InstanceType = &v
	return s
}

func (s *ListSourceEventsResponseBodyData) SetMonitorSourceId(v int64) *ListSourceEventsResponseBodyData {
	s.MonitorSourceId = &v
	return s
}

func (s *ListSourceEventsResponseBodyData) SetMonitorSourceName(v string) *ListSourceEventsResponseBodyData {
	s.MonitorSourceName = &v
	return s
}

func (s *ListSourceEventsResponseBodyData) SetRouteRuleId(v int64) *ListSourceEventsResponseBodyData {
	s.RouteRuleId = &v
	return s
}

func (s *ListSourceEventsResponseBodyData) SetTenantRamId(v int64) *ListSourceEventsResponseBodyData {
	s.TenantRamId = &v
	return s
}

type ListSourceEventsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSourceEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSourceEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSourceEventsResponse) GoString() string {
	return s.String()
}

func (s *ListSourceEventsResponse) SetHeaders(v map[string]*string) *ListSourceEventsResponse {
	s.Headers = v
	return s
}

func (s *ListSourceEventsResponse) SetStatusCode(v int32) *ListSourceEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSourceEventsResponse) SetBody(v *ListSourceEventsResponseBody) *ListSourceEventsResponse {
	s.Body = v
	return s
}

type ListSourceEventsForMonitorSourceRequest struct {
	MonitorSourceId *int64 `json:"monitorSourceId,omitempty" xml:"monitorSourceId,omitempty"`
}

func (s ListSourceEventsForMonitorSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSourceEventsForMonitorSourceRequest) GoString() string {
	return s.String()
}

func (s *ListSourceEventsForMonitorSourceRequest) SetMonitorSourceId(v int64) *ListSourceEventsForMonitorSourceRequest {
	s.MonitorSourceId = &v
	return s
}

type ListSourceEventsForMonitorSourceResponseBody struct {
	Data []*ListSourceEventsForMonitorSourceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListSourceEventsForMonitorSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSourceEventsForMonitorSourceResponseBody) GoString() string {
	return s.String()
}

func (s *ListSourceEventsForMonitorSourceResponseBody) SetData(v []*ListSourceEventsForMonitorSourceResponseBodyData) *ListSourceEventsForMonitorSourceResponseBody {
	s.Data = v
	return s
}

func (s *ListSourceEventsForMonitorSourceResponseBody) SetRequestId(v string) *ListSourceEventsForMonitorSourceResponseBody {
	s.RequestId = &v
	return s
}

type ListSourceEventsForMonitorSourceResponseBodyData struct {
	EventJson         *string `json:"eventJson,omitempty" xml:"eventJson,omitempty"`
	EventTime         *string `json:"eventTime,omitempty" xml:"eventTime,omitempty"`
	MonitorSourceId   *bool   `json:"monitorSourceId,omitempty" xml:"monitorSourceId,omitempty"`
	MonitorSourceName *string `json:"monitorSourceName,omitempty" xml:"monitorSourceName,omitempty"`
}

func (s ListSourceEventsForMonitorSourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListSourceEventsForMonitorSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSourceEventsForMonitorSourceResponseBodyData) SetEventJson(v string) *ListSourceEventsForMonitorSourceResponseBodyData {
	s.EventJson = &v
	return s
}

func (s *ListSourceEventsForMonitorSourceResponseBodyData) SetEventTime(v string) *ListSourceEventsForMonitorSourceResponseBodyData {
	s.EventTime = &v
	return s
}

func (s *ListSourceEventsForMonitorSourceResponseBodyData) SetMonitorSourceId(v bool) *ListSourceEventsForMonitorSourceResponseBodyData {
	s.MonitorSourceId = &v
	return s
}

func (s *ListSourceEventsForMonitorSourceResponseBodyData) SetMonitorSourceName(v string) *ListSourceEventsForMonitorSourceResponseBodyData {
	s.MonitorSourceName = &v
	return s
}

type ListSourceEventsForMonitorSourceResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSourceEventsForMonitorSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSourceEventsForMonitorSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSourceEventsForMonitorSourceResponse) GoString() string {
	return s.String()
}

func (s *ListSourceEventsForMonitorSourceResponse) SetHeaders(v map[string]*string) *ListSourceEventsForMonitorSourceResponse {
	s.Headers = v
	return s
}

func (s *ListSourceEventsForMonitorSourceResponse) SetStatusCode(v int32) *ListSourceEventsForMonitorSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSourceEventsForMonitorSourceResponse) SetBody(v *ListSourceEventsForMonitorSourceResponseBody) *ListSourceEventsForMonitorSourceResponse {
	s.Body = v
	return s
}

type ListSubscriptionServiceGroupsRequest struct {
	ClientToken *string  `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	ServiceIds  []*int64 `json:"serviceIds,omitempty" xml:"serviceIds,omitempty" type:"Repeated"`
}

func (s ListSubscriptionServiceGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionServiceGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListSubscriptionServiceGroupsRequest) SetClientToken(v string) *ListSubscriptionServiceGroupsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListSubscriptionServiceGroupsRequest) SetServiceIds(v []*int64) *ListSubscriptionServiceGroupsRequest {
	s.ServiceIds = v
	return s
}

type ListSubscriptionServiceGroupsResponseBody struct {
	// data
	Data []*ListSubscriptionServiceGroupsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListSubscriptionServiceGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionServiceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubscriptionServiceGroupsResponseBody) SetData(v []*ListSubscriptionServiceGroupsResponseBodyData) *ListSubscriptionServiceGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListSubscriptionServiceGroupsResponseBody) SetRequestId(v string) *ListSubscriptionServiceGroupsResponseBody {
	s.RequestId = &v
	return s
}

type ListSubscriptionServiceGroupsResponseBodyData struct {
	ServiceGroupDescription *string `json:"serviceGroupDescription,omitempty" xml:"serviceGroupDescription,omitempty"`
	ServiceId               *int64  `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	ServiceName             *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
}

func (s ListSubscriptionServiceGroupsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionServiceGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSubscriptionServiceGroupsResponseBodyData) SetServiceGroupDescription(v string) *ListSubscriptionServiceGroupsResponseBodyData {
	s.ServiceGroupDescription = &v
	return s
}

func (s *ListSubscriptionServiceGroupsResponseBodyData) SetServiceId(v int64) *ListSubscriptionServiceGroupsResponseBodyData {
	s.ServiceId = &v
	return s
}

func (s *ListSubscriptionServiceGroupsResponseBodyData) SetServiceName(v string) *ListSubscriptionServiceGroupsResponseBodyData {
	s.ServiceName = &v
	return s
}

type ListSubscriptionServiceGroupsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSubscriptionServiceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSubscriptionServiceGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionServiceGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListSubscriptionServiceGroupsResponse) SetHeaders(v map[string]*string) *ListSubscriptionServiceGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListSubscriptionServiceGroupsResponse) SetStatusCode(v int32) *ListSubscriptionServiceGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSubscriptionServiceGroupsResponse) SetBody(v *ListSubscriptionServiceGroupsResponseBody) *ListSubscriptionServiceGroupsResponse {
	s.Body = v
	return s
}

type ListSubscriptionsRequest struct {
	ClientToken                 *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	NotFilterScopeObjectDeleted *bool   `json:"notFilterScopeObjectDeleted,omitempty" xml:"notFilterScopeObjectDeleted,omitempty"`
	NotifyObject                *string `json:"notifyObject,omitempty" xml:"notifyObject,omitempty"`
	NotifyObjectType            *string `json:"notifyObjectType,omitempty" xml:"notifyObjectType,omitempty"`
	PageNumber                  *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize                    *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Scope                       *string `json:"scope,omitempty" xml:"scope,omitempty"`
	ScopeObject                 *string `json:"scopeObject,omitempty" xml:"scopeObject,omitempty"`
	SubscriptionTitle           *string `json:"subscriptionTitle,omitempty" xml:"subscriptionTitle,omitempty"`
}

func (s ListSubscriptionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionsRequest) GoString() string {
	return s.String()
}

func (s *ListSubscriptionsRequest) SetClientToken(v string) *ListSubscriptionsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListSubscriptionsRequest) SetNotFilterScopeObjectDeleted(v bool) *ListSubscriptionsRequest {
	s.NotFilterScopeObjectDeleted = &v
	return s
}

func (s *ListSubscriptionsRequest) SetNotifyObject(v string) *ListSubscriptionsRequest {
	s.NotifyObject = &v
	return s
}

func (s *ListSubscriptionsRequest) SetNotifyObjectType(v string) *ListSubscriptionsRequest {
	s.NotifyObjectType = &v
	return s
}

func (s *ListSubscriptionsRequest) SetPageNumber(v int32) *ListSubscriptionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSubscriptionsRequest) SetPageSize(v int32) *ListSubscriptionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSubscriptionsRequest) SetScope(v string) *ListSubscriptionsRequest {
	s.Scope = &v
	return s
}

func (s *ListSubscriptionsRequest) SetScopeObject(v string) *ListSubscriptionsRequest {
	s.ScopeObject = &v
	return s
}

func (s *ListSubscriptionsRequest) SetSubscriptionTitle(v string) *ListSubscriptionsRequest {
	s.SubscriptionTitle = &v
	return s
}

type ListSubscriptionsResponseBody struct {
	// data
	Data       []*ListSubscriptionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	PageNumber *int64                               `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64                               `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// id of the request
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TotalCount *int64  `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListSubscriptionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubscriptionsResponseBody) SetData(v []*ListSubscriptionsResponseBodyData) *ListSubscriptionsResponseBody {
	s.Data = v
	return s
}

func (s *ListSubscriptionsResponseBody) SetPageNumber(v int64) *ListSubscriptionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSubscriptionsResponseBody) SetPageSize(v int64) *ListSubscriptionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSubscriptionsResponseBody) SetRequestId(v string) *ListSubscriptionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSubscriptionsResponseBody) SetTotalCount(v int64) *ListSubscriptionsResponseBody {
	s.TotalCount = &v
	return s
}

type ListSubscriptionsResponseBodyData struct {
	EndTime           *string                                              `json:"endTime,omitempty" xml:"endTime,omitempty"`
	ExpiredType       *string                                              `json:"expiredType,omitempty" xml:"expiredType,omitempty"`
	NotifyObjectList  []*ListSubscriptionsResponseBodyDataNotifyObjectList `json:"notifyObjectList,omitempty" xml:"notifyObjectList,omitempty" type:"Repeated"`
	NotifyObjectType  *int64                                               `json:"notifyObjectType,omitempty" xml:"notifyObjectType,omitempty"`
	Scope             *int64                                               `json:"scope,omitempty" xml:"scope,omitempty"`
	ScopeObjectList   []*ListSubscriptionsResponseBodyDataScopeObjectList  `json:"scopeObjectList,omitempty" xml:"scopeObjectList,omitempty" type:"Repeated"`
	StartTime         *string                                              `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status            *string                                              `json:"status,omitempty" xml:"status,omitempty"`
	SubscriptionId    *int64                                               `json:"subscriptionId,omitempty" xml:"subscriptionId,omitempty"`
	SubscriptionTitle *string                                              `json:"subscriptionTitle,omitempty" xml:"subscriptionTitle,omitempty"`
}

func (s ListSubscriptionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSubscriptionsResponseBodyData) SetEndTime(v string) *ListSubscriptionsResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListSubscriptionsResponseBodyData) SetExpiredType(v string) *ListSubscriptionsResponseBodyData {
	s.ExpiredType = &v
	return s
}

func (s *ListSubscriptionsResponseBodyData) SetNotifyObjectList(v []*ListSubscriptionsResponseBodyDataNotifyObjectList) *ListSubscriptionsResponseBodyData {
	s.NotifyObjectList = v
	return s
}

func (s *ListSubscriptionsResponseBodyData) SetNotifyObjectType(v int64) *ListSubscriptionsResponseBodyData {
	s.NotifyObjectType = &v
	return s
}

func (s *ListSubscriptionsResponseBodyData) SetScope(v int64) *ListSubscriptionsResponseBodyData {
	s.Scope = &v
	return s
}

func (s *ListSubscriptionsResponseBodyData) SetScopeObjectList(v []*ListSubscriptionsResponseBodyDataScopeObjectList) *ListSubscriptionsResponseBodyData {
	s.ScopeObjectList = v
	return s
}

func (s *ListSubscriptionsResponseBodyData) SetStartTime(v string) *ListSubscriptionsResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListSubscriptionsResponseBodyData) SetStatus(v string) *ListSubscriptionsResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListSubscriptionsResponseBodyData) SetSubscriptionId(v int64) *ListSubscriptionsResponseBodyData {
	s.SubscriptionId = &v
	return s
}

func (s *ListSubscriptionsResponseBodyData) SetSubscriptionTitle(v string) *ListSubscriptionsResponseBodyData {
	s.SubscriptionTitle = &v
	return s
}

type ListSubscriptionsResponseBodyDataNotifyObjectList struct {
	Id               *int64  `json:"id,omitempty" xml:"id,omitempty"`
	IsValid          *int64  `json:"isValid,omitempty" xml:"isValid,omitempty"`
	Name             *string `json:"name,omitempty" xml:"name,omitempty"`
	NotifyObjectId   *int64  `json:"notifyObjectId,omitempty" xml:"notifyObjectId,omitempty"`
	NotifyObjectType *int64  `json:"notifyObjectType,omitempty" xml:"notifyObjectType,omitempty"`
}

func (s ListSubscriptionsResponseBodyDataNotifyObjectList) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionsResponseBodyDataNotifyObjectList) GoString() string {
	return s.String()
}

func (s *ListSubscriptionsResponseBodyDataNotifyObjectList) SetId(v int64) *ListSubscriptionsResponseBodyDataNotifyObjectList {
	s.Id = &v
	return s
}

func (s *ListSubscriptionsResponseBodyDataNotifyObjectList) SetIsValid(v int64) *ListSubscriptionsResponseBodyDataNotifyObjectList {
	s.IsValid = &v
	return s
}

func (s *ListSubscriptionsResponseBodyDataNotifyObjectList) SetName(v string) *ListSubscriptionsResponseBodyDataNotifyObjectList {
	s.Name = &v
	return s
}

func (s *ListSubscriptionsResponseBodyDataNotifyObjectList) SetNotifyObjectId(v int64) *ListSubscriptionsResponseBodyDataNotifyObjectList {
	s.NotifyObjectId = &v
	return s
}

func (s *ListSubscriptionsResponseBodyDataNotifyObjectList) SetNotifyObjectType(v int64) *ListSubscriptionsResponseBodyDataNotifyObjectList {
	s.NotifyObjectType = &v
	return s
}

type ListSubscriptionsResponseBodyDataScopeObjectList struct {
	Id            *int64  `json:"id,omitempty" xml:"id,omitempty"`
	IsValid       *int64  `json:"isValid,omitempty" xml:"isValid,omitempty"`
	Scope         *int64  `json:"scope,omitempty" xml:"scope,omitempty"`
	ScopeObject   *string `json:"scopeObject,omitempty" xml:"scopeObject,omitempty"`
	ScopeObjectId *int64  `json:"scopeObjectId,omitempty" xml:"scopeObjectId,omitempty"`
}

func (s ListSubscriptionsResponseBodyDataScopeObjectList) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionsResponseBodyDataScopeObjectList) GoString() string {
	return s.String()
}

func (s *ListSubscriptionsResponseBodyDataScopeObjectList) SetId(v int64) *ListSubscriptionsResponseBodyDataScopeObjectList {
	s.Id = &v
	return s
}

func (s *ListSubscriptionsResponseBodyDataScopeObjectList) SetIsValid(v int64) *ListSubscriptionsResponseBodyDataScopeObjectList {
	s.IsValid = &v
	return s
}

func (s *ListSubscriptionsResponseBodyDataScopeObjectList) SetScope(v int64) *ListSubscriptionsResponseBodyDataScopeObjectList {
	s.Scope = &v
	return s
}

func (s *ListSubscriptionsResponseBodyDataScopeObjectList) SetScopeObject(v string) *ListSubscriptionsResponseBodyDataScopeObjectList {
	s.ScopeObject = &v
	return s
}

func (s *ListSubscriptionsResponseBodyDataScopeObjectList) SetScopeObjectId(v int64) *ListSubscriptionsResponseBodyDataScopeObjectList {
	s.ScopeObjectId = &v
	return s
}

type ListSubscriptionsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSubscriptionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSubscriptionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionsResponse) GoString() string {
	return s.String()
}

func (s *ListSubscriptionsResponse) SetHeaders(v map[string]*string) *ListSubscriptionsResponse {
	s.Headers = v
	return s
}

func (s *ListSubscriptionsResponse) SetStatusCode(v int32) *ListSubscriptionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSubscriptionsResponse) SetBody(v *ListSubscriptionsResponseBody) *ListSubscriptionsResponse {
	s.Body = v
	return s
}

type ListTrendForSourceEventRequest struct {
	EndTime      *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	InstanceId   *int64  `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	StartTime    *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	TimeUnit     *int64  `json:"timeUnit,omitempty" xml:"timeUnit,omitempty"`
}

func (s ListTrendForSourceEventRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTrendForSourceEventRequest) GoString() string {
	return s.String()
}

func (s *ListTrendForSourceEventRequest) SetEndTime(v string) *ListTrendForSourceEventRequest {
	s.EndTime = &v
	return s
}

func (s *ListTrendForSourceEventRequest) SetInstanceId(v int64) *ListTrendForSourceEventRequest {
	s.InstanceId = &v
	return s
}

func (s *ListTrendForSourceEventRequest) SetInstanceType(v string) *ListTrendForSourceEventRequest {
	s.InstanceType = &v
	return s
}

func (s *ListTrendForSourceEventRequest) SetRequestId(v string) *ListTrendForSourceEventRequest {
	s.RequestId = &v
	return s
}

func (s *ListTrendForSourceEventRequest) SetStartTime(v string) *ListTrendForSourceEventRequest {
	s.StartTime = &v
	return s
}

func (s *ListTrendForSourceEventRequest) SetTimeUnit(v int64) *ListTrendForSourceEventRequest {
	s.TimeUnit = &v
	return s
}

type ListTrendForSourceEventResponseBody struct {
	Data []*ListTrendForSourceEventResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListTrendForSourceEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTrendForSourceEventResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrendForSourceEventResponseBody) SetData(v []*ListTrendForSourceEventResponseBodyData) *ListTrendForSourceEventResponseBody {
	s.Data = v
	return s
}

func (s *ListTrendForSourceEventResponseBody) SetRequestId(v string) *ListTrendForSourceEventResponseBody {
	s.RequestId = &v
	return s
}

type ListTrendForSourceEventResponseBodyData struct {
	ConvergenceRate     *string                `json:"convergenceRate,omitempty" xml:"convergenceRate,omitempty"`
	MaxSustainTime      *int64                 `json:"maxSustainTime,omitempty" xml:"maxSustainTime,omitempty"`
	SkipDay             *bool                  `json:"skipDay,omitempty" xml:"skipDay,omitempty"`
	SourceEventsStatMap map[string]interface{} `json:"sourceEventsStatMap,omitempty" xml:"sourceEventsStatMap,omitempty"`
	Unit                *string                `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s ListTrendForSourceEventResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTrendForSourceEventResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTrendForSourceEventResponseBodyData) SetConvergenceRate(v string) *ListTrendForSourceEventResponseBodyData {
	s.ConvergenceRate = &v
	return s
}

func (s *ListTrendForSourceEventResponseBodyData) SetMaxSustainTime(v int64) *ListTrendForSourceEventResponseBodyData {
	s.MaxSustainTime = &v
	return s
}

func (s *ListTrendForSourceEventResponseBodyData) SetSkipDay(v bool) *ListTrendForSourceEventResponseBodyData {
	s.SkipDay = &v
	return s
}

func (s *ListTrendForSourceEventResponseBodyData) SetSourceEventsStatMap(v map[string]interface{}) *ListTrendForSourceEventResponseBodyData {
	s.SourceEventsStatMap = v
	return s
}

func (s *ListTrendForSourceEventResponseBodyData) SetUnit(v string) *ListTrendForSourceEventResponseBodyData {
	s.Unit = &v
	return s
}

type ListTrendForSourceEventResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTrendForSourceEventResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTrendForSourceEventResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTrendForSourceEventResponse) GoString() string {
	return s.String()
}

func (s *ListTrendForSourceEventResponse) SetHeaders(v map[string]*string) *ListTrendForSourceEventResponse {
	s.Headers = v
	return s
}

func (s *ListTrendForSourceEventResponse) SetStatusCode(v int32) *ListTrendForSourceEventResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTrendForSourceEventResponse) SetBody(v *ListTrendForSourceEventResponseBody) *ListTrendForSourceEventResponse {
	s.Body = v
	return s
}

type ListUserSerivceGroupsRequest struct {
	// clientToken
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	UserId      *int64  `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListUserSerivceGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserSerivceGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListUserSerivceGroupsRequest) SetClientToken(v string) *ListUserSerivceGroupsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListUserSerivceGroupsRequest) SetUserId(v int64) *ListUserSerivceGroupsRequest {
	s.UserId = &v
	return s
}

type ListUserSerivceGroupsResponseBody struct {
	// data
	Data *ListUserSerivceGroupsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListUserSerivceGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserSerivceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserSerivceGroupsResponseBody) SetData(v *ListUserSerivceGroupsResponseBodyData) *ListUserSerivceGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListUserSerivceGroupsResponseBody) SetRequestId(v string) *ListUserSerivceGroupsResponseBody {
	s.RequestId = &v
	return s
}

type ListUserSerivceGroupsResponseBodyData struct {
	Email         *string                                               `json:"email,omitempty" xml:"email,omitempty"`
	Phone         *string                                               `json:"phone,omitempty" xml:"phone,omitempty"`
	RamId         *int64                                                `json:"ramId,omitempty" xml:"ramId,omitempty"`
	ServiceGroups []*ListUserSerivceGroupsResponseBodyDataServiceGroups `json:"serviceGroups,omitempty" xml:"serviceGroups,omitempty" type:"Repeated"`
	UserId        *int64                                                `json:"userId,omitempty" xml:"userId,omitempty"`
	Username      *string                                               `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ListUserSerivceGroupsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListUserSerivceGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUserSerivceGroupsResponseBodyData) SetEmail(v string) *ListUserSerivceGroupsResponseBodyData {
	s.Email = &v
	return s
}

func (s *ListUserSerivceGroupsResponseBodyData) SetPhone(v string) *ListUserSerivceGroupsResponseBodyData {
	s.Phone = &v
	return s
}

func (s *ListUserSerivceGroupsResponseBodyData) SetRamId(v int64) *ListUserSerivceGroupsResponseBodyData {
	s.RamId = &v
	return s
}

func (s *ListUserSerivceGroupsResponseBodyData) SetServiceGroups(v []*ListUserSerivceGroupsResponseBodyDataServiceGroups) *ListUserSerivceGroupsResponseBodyData {
	s.ServiceGroups = v
	return s
}

func (s *ListUserSerivceGroupsResponseBodyData) SetUserId(v int64) *ListUserSerivceGroupsResponseBodyData {
	s.UserId = &v
	return s
}

func (s *ListUserSerivceGroupsResponseBodyData) SetUsername(v string) *ListUserSerivceGroupsResponseBodyData {
	s.Username = &v
	return s
}

type ListUserSerivceGroupsResponseBodyDataServiceGroups struct {
	ServiceGroupDescription *string `json:"serviceGroupDescription,omitempty" xml:"serviceGroupDescription,omitempty"`
	ServiceGroupId          *int64  `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
	ServiceGroupName        *string `json:"serviceGroupName,omitempty" xml:"serviceGroupName,omitempty"`
}

func (s ListUserSerivceGroupsResponseBodyDataServiceGroups) String() string {
	return tea.Prettify(s)
}

func (s ListUserSerivceGroupsResponseBodyDataServiceGroups) GoString() string {
	return s.String()
}

func (s *ListUserSerivceGroupsResponseBodyDataServiceGroups) SetServiceGroupDescription(v string) *ListUserSerivceGroupsResponseBodyDataServiceGroups {
	s.ServiceGroupDescription = &v
	return s
}

func (s *ListUserSerivceGroupsResponseBodyDataServiceGroups) SetServiceGroupId(v int64) *ListUserSerivceGroupsResponseBodyDataServiceGroups {
	s.ServiceGroupId = &v
	return s
}

func (s *ListUserSerivceGroupsResponseBodyDataServiceGroups) SetServiceGroupName(v string) *ListUserSerivceGroupsResponseBodyDataServiceGroups {
	s.ServiceGroupName = &v
	return s
}

type ListUserSerivceGroupsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUserSerivceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserSerivceGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserSerivceGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListUserSerivceGroupsResponse) SetHeaders(v map[string]*string) *ListUserSerivceGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListUserSerivceGroupsResponse) SetStatusCode(v int32) *ListUserSerivceGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserSerivceGroupsResponse) SetBody(v *ListUserSerivceGroupsResponseBody) *ListUserSerivceGroupsResponse {
	s.Body = v
	return s
}

type ListUsersRequest struct {
	// clientToken
	ClientToken    *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	PageNumber     *int64  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize       *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Phone          *string `json:"phone,omitempty" xml:"phone,omitempty"`
	RamId          *string `json:"ramId,omitempty" xml:"ramId,omitempty"`
	Scene          *int64  `json:"scene,omitempty" xml:"scene,omitempty"`
	SynergyChannel *string `json:"synergyChannel,omitempty" xml:"synergyChannel,omitempty"`
	Username       *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ListUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) SetClientToken(v string) *ListUsersRequest {
	s.ClientToken = &v
	return s
}

func (s *ListUsersRequest) SetPageNumber(v int64) *ListUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUsersRequest) SetPageSize(v int64) *ListUsersRequest {
	s.PageSize = &v
	return s
}

func (s *ListUsersRequest) SetPhone(v string) *ListUsersRequest {
	s.Phone = &v
	return s
}

func (s *ListUsersRequest) SetRamId(v string) *ListUsersRequest {
	s.RamId = &v
	return s
}

func (s *ListUsersRequest) SetScene(v int64) *ListUsersRequest {
	s.Scene = &v
	return s
}

func (s *ListUsersRequest) SetSynergyChannel(v string) *ListUsersRequest {
	s.SynergyChannel = &v
	return s
}

func (s *ListUsersRequest) SetUsername(v string) *ListUsersRequest {
	s.Username = &v
	return s
}

type ListUsersResponseBody struct {
	// data
	Data       []*ListUsersResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	PageNumber *int64                       `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64                       `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// id of the request
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TotalCount *int64  `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) SetData(v []*ListUsersResponseBodyData) *ListUsersResponseBody {
	s.Data = v
	return s
}

func (s *ListUsersResponseBody) SetPageNumber(v int64) *ListUsersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListUsersResponseBody) SetPageSize(v int64) *ListUsersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetTotalCount(v int64) *ListUsersResponseBody {
	s.TotalCount = &v
	return s
}

type ListUsersResponseBodyData struct {
	AccountType    *int64  `json:"accountType,omitempty" xml:"accountType,omitempty"`
	AppAccount     *string `json:"appAccount,omitempty" xml:"appAccount,omitempty"`
	Email          *string `json:"email,omitempty" xml:"email,omitempty"`
	IsEditableUser *int64  `json:"isEditableUser,omitempty" xml:"isEditableUser,omitempty"`
	IsOperation    *int32  `json:"isOperation,omitempty" xml:"isOperation,omitempty"`
	IsRam          *int32  `json:"isRam,omitempty" xml:"isRam,omitempty"`
	IsRelated      *string `json:"isRelated,omitempty" xml:"isRelated,omitempty"`
	Phone          *string `json:"phone,omitempty" xml:"phone,omitempty"`
	RamId          *int64  `json:"ramId,omitempty" xml:"ramId,omitempty"`
	SynergyChannel *string `json:"synergyChannel,omitempty" xml:"synergyChannel,omitempty"`
	UserId         *int64  `json:"userId,omitempty" xml:"userId,omitempty"`
	Username       *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ListUsersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyData) SetAccountType(v int64) *ListUsersResponseBodyData {
	s.AccountType = &v
	return s
}

func (s *ListUsersResponseBodyData) SetAppAccount(v string) *ListUsersResponseBodyData {
	s.AppAccount = &v
	return s
}

func (s *ListUsersResponseBodyData) SetEmail(v string) *ListUsersResponseBodyData {
	s.Email = &v
	return s
}

func (s *ListUsersResponseBodyData) SetIsEditableUser(v int64) *ListUsersResponseBodyData {
	s.IsEditableUser = &v
	return s
}

func (s *ListUsersResponseBodyData) SetIsOperation(v int32) *ListUsersResponseBodyData {
	s.IsOperation = &v
	return s
}

func (s *ListUsersResponseBodyData) SetIsRam(v int32) *ListUsersResponseBodyData {
	s.IsRam = &v
	return s
}

func (s *ListUsersResponseBodyData) SetIsRelated(v string) *ListUsersResponseBodyData {
	s.IsRelated = &v
	return s
}

func (s *ListUsersResponseBodyData) SetPhone(v string) *ListUsersResponseBodyData {
	s.Phone = &v
	return s
}

func (s *ListUsersResponseBodyData) SetRamId(v int64) *ListUsersResponseBodyData {
	s.RamId = &v
	return s
}

func (s *ListUsersResponseBodyData) SetSynergyChannel(v string) *ListUsersResponseBodyData {
	s.SynergyChannel = &v
	return s
}

func (s *ListUsersResponseBodyData) SetUserId(v int64) *ListUsersResponseBodyData {
	s.UserId = &v
	return s
}

func (s *ListUsersResponseBodyData) SetUsername(v string) *ListUsersResponseBodyData {
	s.Username = &v
	return s
}

type ListUsersResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponse) GoString() string {
	return s.String()
}

func (s *ListUsersResponse) SetHeaders(v map[string]*string) *ListUsersResponse {
	s.Headers = v
	return s
}

func (s *ListUsersResponse) SetStatusCode(v int32) *ListUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsersResponse) SetBody(v *ListUsersResponseBody) *ListUsersResponse {
	s.Body = v
	return s
}

type RecoverProblemRequest struct {
	ProblemId         *int64  `json:"problemId,omitempty" xml:"problemId,omitempty"`
	ProblemNotifyType *string `json:"problemNotifyType,omitempty" xml:"problemNotifyType,omitempty"`
	RecoveryTime      *string `json:"recoveryTime,omitempty" xml:"recoveryTime,omitempty"`
}

func (s RecoverProblemRequest) String() string {
	return tea.Prettify(s)
}

func (s RecoverProblemRequest) GoString() string {
	return s.String()
}

func (s *RecoverProblemRequest) SetProblemId(v int64) *RecoverProblemRequest {
	s.ProblemId = &v
	return s
}

func (s *RecoverProblemRequest) SetProblemNotifyType(v string) *RecoverProblemRequest {
	s.ProblemNotifyType = &v
	return s
}

func (s *RecoverProblemRequest) SetRecoveryTime(v string) *RecoverProblemRequest {
	s.RecoveryTime = &v
	return s
}

type RecoverProblemResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RecoverProblemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecoverProblemResponseBody) GoString() string {
	return s.String()
}

func (s *RecoverProblemResponseBody) SetRequestId(v string) *RecoverProblemResponseBody {
	s.RequestId = &v
	return s
}

type RecoverProblemResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecoverProblemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecoverProblemResponse) String() string {
	return tea.Prettify(s)
}

func (s RecoverProblemResponse) GoString() string {
	return s.String()
}

func (s *RecoverProblemResponse) SetHeaders(v map[string]*string) *RecoverProblemResponse {
	s.Headers = v
	return s
}

func (s *RecoverProblemResponse) SetStatusCode(v int32) *RecoverProblemResponse {
	s.StatusCode = &v
	return s
}

func (s *RecoverProblemResponse) SetBody(v *RecoverProblemResponseBody) *RecoverProblemResponse {
	s.Body = v
	return s
}

type RefreshIntegrationConfigKeyRequest struct {
	ClientToken         *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	IntegrationConfigId *int64  `json:"integrationConfigId,omitempty" xml:"integrationConfigId,omitempty"`
}

func (s RefreshIntegrationConfigKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshIntegrationConfigKeyRequest) GoString() string {
	return s.String()
}

func (s *RefreshIntegrationConfigKeyRequest) SetClientToken(v string) *RefreshIntegrationConfigKeyRequest {
	s.ClientToken = &v
	return s
}

func (s *RefreshIntegrationConfigKeyRequest) SetIntegrationConfigId(v int64) *RefreshIntegrationConfigKeyRequest {
	s.IntegrationConfigId = &v
	return s
}

type RefreshIntegrationConfigKeyResponseBody struct {
	// data
	Data *RefreshIntegrationConfigKeyResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RefreshIntegrationConfigKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshIntegrationConfigKeyResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshIntegrationConfigKeyResponseBody) SetData(v *RefreshIntegrationConfigKeyResponseBodyData) *RefreshIntegrationConfigKeyResponseBody {
	s.Data = v
	return s
}

func (s *RefreshIntegrationConfigKeyResponseBody) SetRequestId(v string) *RefreshIntegrationConfigKeyResponseBody {
	s.RequestId = &v
	return s
}

type RefreshIntegrationConfigKeyResponseBodyData struct {
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s RefreshIntegrationConfigKeyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RefreshIntegrationConfigKeyResponseBodyData) GoString() string {
	return s.String()
}

func (s *RefreshIntegrationConfigKeyResponseBodyData) SetKey(v string) *RefreshIntegrationConfigKeyResponseBodyData {
	s.Key = &v
	return s
}

type RefreshIntegrationConfigKeyResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RefreshIntegrationConfigKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RefreshIntegrationConfigKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshIntegrationConfigKeyResponse) GoString() string {
	return s.String()
}

func (s *RefreshIntegrationConfigKeyResponse) SetHeaders(v map[string]*string) *RefreshIntegrationConfigKeyResponse {
	s.Headers = v
	return s
}

func (s *RefreshIntegrationConfigKeyResponse) SetStatusCode(v int32) *RefreshIntegrationConfigKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshIntegrationConfigKeyResponse) SetBody(v *RefreshIntegrationConfigKeyResponseBody) *RefreshIntegrationConfigKeyResponse {
	s.Body = v
	return s
}

type RemoveIntegrationConfigRequest struct {
	ClientToken         *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	IntegrationConfigId *int64  `json:"integrationConfigId,omitempty" xml:"integrationConfigId,omitempty"`
}

func (s RemoveIntegrationConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveIntegrationConfigRequest) GoString() string {
	return s.String()
}

func (s *RemoveIntegrationConfigRequest) SetClientToken(v string) *RemoveIntegrationConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveIntegrationConfigRequest) SetIntegrationConfigId(v int64) *RemoveIntegrationConfigRequest {
	s.IntegrationConfigId = &v
	return s
}

type RemoveIntegrationConfigResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RemoveIntegrationConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveIntegrationConfigResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveIntegrationConfigResponseBody) SetRequestId(v string) *RemoveIntegrationConfigResponseBody {
	s.RequestId = &v
	return s
}

type RemoveIntegrationConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveIntegrationConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveIntegrationConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveIntegrationConfigResponse) GoString() string {
	return s.String()
}

func (s *RemoveIntegrationConfigResponse) SetHeaders(v map[string]*string) *RemoveIntegrationConfigResponse {
	s.Headers = v
	return s
}

func (s *RemoveIntegrationConfigResponse) SetStatusCode(v int32) *RemoveIntegrationConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveIntegrationConfigResponse) SetBody(v *RemoveIntegrationConfigResponseBody) *RemoveIntegrationConfigResponse {
	s.Body = v
	return s
}

type RemoveProblemServiceGroupRequest struct {
	ProblemId       *int64   `json:"problemId,omitempty" xml:"problemId,omitempty"`
	ServiceGroupIds []*int64 `json:"serviceGroupIds,omitempty" xml:"serviceGroupIds,omitempty" type:"Repeated"`
}

func (s RemoveProblemServiceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveProblemServiceGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveProblemServiceGroupRequest) SetProblemId(v int64) *RemoveProblemServiceGroupRequest {
	s.ProblemId = &v
	return s
}

func (s *RemoveProblemServiceGroupRequest) SetServiceGroupIds(v []*int64) *RemoveProblemServiceGroupRequest {
	s.ServiceGroupIds = v
	return s
}

type RemoveProblemServiceGroupResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RemoveProblemServiceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveProblemServiceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveProblemServiceGroupResponseBody) SetRequestId(v string) *RemoveProblemServiceGroupResponseBody {
	s.RequestId = &v
	return s
}

type RemoveProblemServiceGroupResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveProblemServiceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveProblemServiceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveProblemServiceGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveProblemServiceGroupResponse) SetHeaders(v map[string]*string) *RemoveProblemServiceGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveProblemServiceGroupResponse) SetStatusCode(v int32) *RemoveProblemServiceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveProblemServiceGroupResponse) SetBody(v *RemoveProblemServiceGroupResponseBody) *RemoveProblemServiceGroupResponse {
	s.Body = v
	return s
}

type ReplayProblemRequest struct {
	ClientToken      *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	ProblemId        *int64  `json:"problemId,omitempty" xml:"problemId,omitempty"`
	ReplayDutyUserId *int64  `json:"replayDutyUserId,omitempty" xml:"replayDutyUserId,omitempty"`
}

func (s ReplayProblemRequest) String() string {
	return tea.Prettify(s)
}

func (s ReplayProblemRequest) GoString() string {
	return s.String()
}

func (s *ReplayProblemRequest) SetClientToken(v string) *ReplayProblemRequest {
	s.ClientToken = &v
	return s
}

func (s *ReplayProblemRequest) SetProblemId(v int64) *ReplayProblemRequest {
	s.ProblemId = &v
	return s
}

func (s *ReplayProblemRequest) SetReplayDutyUserId(v int64) *ReplayProblemRequest {
	s.ReplayDutyUserId = &v
	return s
}

type ReplayProblemResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ReplayProblemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReplayProblemResponseBody) GoString() string {
	return s.String()
}

func (s *ReplayProblemResponseBody) SetRequestId(v string) *ReplayProblemResponseBody {
	s.RequestId = &v
	return s
}

type ReplayProblemResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReplayProblemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReplayProblemResponse) String() string {
	return tea.Prettify(s)
}

func (s ReplayProblemResponse) GoString() string {
	return s.String()
}

func (s *ReplayProblemResponse) SetHeaders(v map[string]*string) *ReplayProblemResponse {
	s.Headers = v
	return s
}

func (s *ReplayProblemResponse) SetStatusCode(v int32) *ReplayProblemResponse {
	s.StatusCode = &v
	return s
}

func (s *ReplayProblemResponse) SetBody(v *ReplayProblemResponseBody) *ReplayProblemResponse {
	s.Body = v
	return s
}

type RespondIncidentRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// 影响程度
	IncidentIds []*int64 `json:"incidentIds,omitempty" xml:"incidentIds,omitempty" type:"Repeated"`
}

func (s RespondIncidentRequest) String() string {
	return tea.Prettify(s)
}

func (s RespondIncidentRequest) GoString() string {
	return s.String()
}

func (s *RespondIncidentRequest) SetClientToken(v string) *RespondIncidentRequest {
	s.ClientToken = &v
	return s
}

func (s *RespondIncidentRequest) SetIncidentIds(v []*int64) *RespondIncidentRequest {
	s.IncidentIds = v
	return s
}

type RespondIncidentResponseBody struct {
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RespondIncidentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RespondIncidentResponseBody) GoString() string {
	return s.String()
}

func (s *RespondIncidentResponseBody) SetRequestId(v string) *RespondIncidentResponseBody {
	s.RequestId = &v
	return s
}

type RespondIncidentResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RespondIncidentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RespondIncidentResponse) String() string {
	return tea.Prettify(s)
}

func (s RespondIncidentResponse) GoString() string {
	return s.String()
}

func (s *RespondIncidentResponse) SetHeaders(v map[string]*string) *RespondIncidentResponse {
	s.Headers = v
	return s
}

func (s *RespondIncidentResponse) SetStatusCode(v int32) *RespondIncidentResponse {
	s.StatusCode = &v
	return s
}

func (s *RespondIncidentResponse) SetBody(v *RespondIncidentResponseBody) *RespondIncidentResponse {
	s.Body = v
	return s
}

type RevokeProblemRecoveryRequest struct {
	ClientToken       *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	ProblemId         *int64  `json:"problemId,omitempty" xml:"problemId,omitempty"`
	ProblemNotifyType *string `json:"problemNotifyType,omitempty" xml:"problemNotifyType,omitempty"`
}

func (s RevokeProblemRecoveryRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeProblemRecoveryRequest) GoString() string {
	return s.String()
}

func (s *RevokeProblemRecoveryRequest) SetClientToken(v string) *RevokeProblemRecoveryRequest {
	s.ClientToken = &v
	return s
}

func (s *RevokeProblemRecoveryRequest) SetProblemId(v int64) *RevokeProblemRecoveryRequest {
	s.ProblemId = &v
	return s
}

func (s *RevokeProblemRecoveryRequest) SetProblemNotifyType(v string) *RevokeProblemRecoveryRequest {
	s.ProblemNotifyType = &v
	return s
}

type RevokeProblemRecoveryResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RevokeProblemRecoveryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevokeProblemRecoveryResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeProblemRecoveryResponseBody) SetRequestId(v string) *RevokeProblemRecoveryResponseBody {
	s.RequestId = &v
	return s
}

type RevokeProblemRecoveryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RevokeProblemRecoveryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RevokeProblemRecoveryResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokeProblemRecoveryResponse) GoString() string {
	return s.String()
}

func (s *RevokeProblemRecoveryResponse) SetHeaders(v map[string]*string) *RevokeProblemRecoveryResponse {
	s.Headers = v
	return s
}

func (s *RevokeProblemRecoveryResponse) SetStatusCode(v int32) *RevokeProblemRecoveryResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeProblemRecoveryResponse) SetBody(v *RevokeProblemRecoveryResponseBody) *RevokeProblemRecoveryResponse {
	s.Body = v
	return s
}

type UnbindUserResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UnbindUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindUserResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindUserResponseBody) SetRequestId(v string) *UnbindUserResponseBody {
	s.RequestId = &v
	return s
}

type UnbindUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnbindUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindUserResponse) GoString() string {
	return s.String()
}

func (s *UnbindUserResponse) SetHeaders(v map[string]*string) *UnbindUserResponse {
	s.Headers = v
	return s
}

func (s *UnbindUserResponse) SetStatusCode(v int32) *UnbindUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindUserResponse) SetBody(v *UnbindUserResponseBody) *UnbindUserResponse {
	s.Body = v
	return s
}

type UpdateEscalationPlanRequest struct {
	// clientToken
	ClientToken                *string                                                  `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	EscalationPlanDescription  *string                                                  `json:"escalationPlanDescription,omitempty" xml:"escalationPlanDescription,omitempty"`
	EscalationPlanId           *int64                                                   `json:"escalationPlanId,omitempty" xml:"escalationPlanId,omitempty"`
	EscalationPlanName         *string                                                  `json:"escalationPlanName,omitempty" xml:"escalationPlanName,omitempty"`
	EscalationPlanRules        []*UpdateEscalationPlanRequestEscalationPlanRules        `json:"escalationPlanRules,omitempty" xml:"escalationPlanRules,omitempty" type:"Repeated"`
	EscalationPlanScopeObjects []*UpdateEscalationPlanRequestEscalationPlanScopeObjects `json:"escalationPlanScopeObjects,omitempty" xml:"escalationPlanScopeObjects,omitempty" type:"Repeated"`
}

func (s UpdateEscalationPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEscalationPlanRequest) GoString() string {
	return s.String()
}

func (s *UpdateEscalationPlanRequest) SetClientToken(v string) *UpdateEscalationPlanRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateEscalationPlanRequest) SetEscalationPlanDescription(v string) *UpdateEscalationPlanRequest {
	s.EscalationPlanDescription = &v
	return s
}

func (s *UpdateEscalationPlanRequest) SetEscalationPlanId(v int64) *UpdateEscalationPlanRequest {
	s.EscalationPlanId = &v
	return s
}

func (s *UpdateEscalationPlanRequest) SetEscalationPlanName(v string) *UpdateEscalationPlanRequest {
	s.EscalationPlanName = &v
	return s
}

func (s *UpdateEscalationPlanRequest) SetEscalationPlanRules(v []*UpdateEscalationPlanRequestEscalationPlanRules) *UpdateEscalationPlanRequest {
	s.EscalationPlanRules = v
	return s
}

func (s *UpdateEscalationPlanRequest) SetEscalationPlanScopeObjects(v []*UpdateEscalationPlanRequestEscalationPlanScopeObjects) *UpdateEscalationPlanRequest {
	s.EscalationPlanScopeObjects = v
	return s
}

type UpdateEscalationPlanRequestEscalationPlanRules struct {
	EscalationPlanConditions []*UpdateEscalationPlanRequestEscalationPlanRulesEscalationPlanConditions `json:"escalationPlanConditions,omitempty" xml:"escalationPlanConditions,omitempty" type:"Repeated"`
	EscalationPlanStrategies []*UpdateEscalationPlanRequestEscalationPlanRulesEscalationPlanStrategies `json:"escalationPlanStrategies,omitempty" xml:"escalationPlanStrategies,omitempty" type:"Repeated"`
	EscalationPlanType       *string                                                                   `json:"escalationPlanType,omitempty" xml:"escalationPlanType,omitempty"`
	Id                       *int64                                                                    `json:"id,omitempty" xml:"id,omitempty"`
}

func (s UpdateEscalationPlanRequestEscalationPlanRules) String() string {
	return tea.Prettify(s)
}

func (s UpdateEscalationPlanRequestEscalationPlanRules) GoString() string {
	return s.String()
}

func (s *UpdateEscalationPlanRequestEscalationPlanRules) SetEscalationPlanConditions(v []*UpdateEscalationPlanRequestEscalationPlanRulesEscalationPlanConditions) *UpdateEscalationPlanRequestEscalationPlanRules {
	s.EscalationPlanConditions = v
	return s
}

func (s *UpdateEscalationPlanRequestEscalationPlanRules) SetEscalationPlanStrategies(v []*UpdateEscalationPlanRequestEscalationPlanRulesEscalationPlanStrategies) *UpdateEscalationPlanRequestEscalationPlanRules {
	s.EscalationPlanStrategies = v
	return s
}

func (s *UpdateEscalationPlanRequestEscalationPlanRules) SetEscalationPlanType(v string) *UpdateEscalationPlanRequestEscalationPlanRules {
	s.EscalationPlanType = &v
	return s
}

func (s *UpdateEscalationPlanRequestEscalationPlanRules) SetId(v int64) *UpdateEscalationPlanRequestEscalationPlanRules {
	s.Id = &v
	return s
}

type UpdateEscalationPlanRequestEscalationPlanRulesEscalationPlanConditions struct {
	// LOW HIGH
	Effection *string `json:"effection,omitempty" xml:"effection,omitempty"`
	// P1 P2 P3 P4
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
}

func (s UpdateEscalationPlanRequestEscalationPlanRulesEscalationPlanConditions) String() string {
	return tea.Prettify(s)
}

func (s UpdateEscalationPlanRequestEscalationPlanRulesEscalationPlanConditions) GoString() string {
	return s.String()
}

func (s *UpdateEscalationPlanRequestEscalationPlanRulesEscalationPlanConditions) SetEffection(v string) *UpdateEscalationPlanRequestEscalationPlanRulesEscalationPlanConditions {
	s.Effection = &v
	return s
}

func (s *UpdateEscalationPlanRequestEscalationPlanRulesEscalationPlanConditions) SetLevel(v string) *UpdateEscalationPlanRequestEscalationPlanRulesEscalationPlanConditions {
	s.Level = &v
	return s
}

type UpdateEscalationPlanRequestEscalationPlanRulesEscalationPlanStrategies struct {
	EnableWebhook   *bool     `json:"enableWebhook,omitempty" xml:"enableWebhook,omitempty"`
	NoticeChannels  []*string `json:"noticeChannels,omitempty" xml:"noticeChannels,omitempty" type:"Repeated"`
	NoticeObjects   []*int64  `json:"noticeObjects,omitempty" xml:"noticeObjects,omitempty" type:"Repeated"`
	NoticeTime      *int64    `json:"noticeTime,omitempty" xml:"noticeTime,omitempty"`
	ServiceGroupIds []*int64  `json:"serviceGroupIds,omitempty" xml:"serviceGroupIds,omitempty" type:"Repeated"`
}

func (s UpdateEscalationPlanRequestEscalationPlanRulesEscalationPlanStrategies) String() string {
	return tea.Prettify(s)
}

func (s UpdateEscalationPlanRequestEscalationPlanRulesEscalationPlanStrategies) GoString() string {
	return s.String()
}

func (s *UpdateEscalationPlanRequestEscalationPlanRulesEscalationPlanStrategies) SetEnableWebhook(v bool) *UpdateEscalationPlanRequestEscalationPlanRulesEscalationPlanStrategies {
	s.EnableWebhook = &v
	return s
}

func (s *UpdateEscalationPlanRequestEscalationPlanRulesEscalationPlanStrategies) SetNoticeChannels(v []*string) *UpdateEscalationPlanRequestEscalationPlanRulesEscalationPlanStrategies {
	s.NoticeChannels = v
	return s
}

func (s *UpdateEscalationPlanRequestEscalationPlanRulesEscalationPlanStrategies) SetNoticeObjects(v []*int64) *UpdateEscalationPlanRequestEscalationPlanRulesEscalationPlanStrategies {
	s.NoticeObjects = v
	return s
}

func (s *UpdateEscalationPlanRequestEscalationPlanRulesEscalationPlanStrategies) SetNoticeTime(v int64) *UpdateEscalationPlanRequestEscalationPlanRulesEscalationPlanStrategies {
	s.NoticeTime = &v
	return s
}

func (s *UpdateEscalationPlanRequestEscalationPlanRulesEscalationPlanStrategies) SetServiceGroupIds(v []*int64) *UpdateEscalationPlanRequestEscalationPlanRulesEscalationPlanStrategies {
	s.ServiceGroupIds = v
	return s
}

type UpdateEscalationPlanRequestEscalationPlanScopeObjects struct {
	Id            *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Scope         *string `json:"scope,omitempty" xml:"scope,omitempty"`
	ScopeObjectId *int64  `json:"scopeObjectId,omitempty" xml:"scopeObjectId,omitempty"`
}

func (s UpdateEscalationPlanRequestEscalationPlanScopeObjects) String() string {
	return tea.Prettify(s)
}

func (s UpdateEscalationPlanRequestEscalationPlanScopeObjects) GoString() string {
	return s.String()
}

func (s *UpdateEscalationPlanRequestEscalationPlanScopeObjects) SetId(v int64) *UpdateEscalationPlanRequestEscalationPlanScopeObjects {
	s.Id = &v
	return s
}

func (s *UpdateEscalationPlanRequestEscalationPlanScopeObjects) SetScope(v string) *UpdateEscalationPlanRequestEscalationPlanScopeObjects {
	s.Scope = &v
	return s
}

func (s *UpdateEscalationPlanRequestEscalationPlanScopeObjects) SetScopeObjectId(v int64) *UpdateEscalationPlanRequestEscalationPlanScopeObjects {
	s.ScopeObjectId = &v
	return s
}

type UpdateEscalationPlanResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateEscalationPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEscalationPlanResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEscalationPlanResponseBody) SetRequestId(v string) *UpdateEscalationPlanResponseBody {
	s.RequestId = &v
	return s
}

type UpdateEscalationPlanResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateEscalationPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateEscalationPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEscalationPlanResponse) GoString() string {
	return s.String()
}

func (s *UpdateEscalationPlanResponse) SetHeaders(v map[string]*string) *UpdateEscalationPlanResponse {
	s.Headers = v
	return s
}

func (s *UpdateEscalationPlanResponse) SetStatusCode(v int32) *UpdateEscalationPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEscalationPlanResponse) SetBody(v *UpdateEscalationPlanResponseBody) *UpdateEscalationPlanResponse {
	s.Body = v
	return s
}

type UpdateIncidentRequest struct {
	ClientToken   *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	Effect        *string `json:"effect,omitempty" xml:"effect,omitempty"`
	IncidentId    *int64  `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
	IncidentLevel *string `json:"incidentLevel,omitempty" xml:"incidentLevel,omitempty"`
	IncidentTitle *string `json:"incidentTitle,omitempty" xml:"incidentTitle,omitempty"`
}

func (s UpdateIncidentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIncidentRequest) GoString() string {
	return s.String()
}

func (s *UpdateIncidentRequest) SetClientToken(v string) *UpdateIncidentRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateIncidentRequest) SetEffect(v string) *UpdateIncidentRequest {
	s.Effect = &v
	return s
}

func (s *UpdateIncidentRequest) SetIncidentId(v int64) *UpdateIncidentRequest {
	s.IncidentId = &v
	return s
}

func (s *UpdateIncidentRequest) SetIncidentLevel(v string) *UpdateIncidentRequest {
	s.IncidentLevel = &v
	return s
}

func (s *UpdateIncidentRequest) SetIncidentTitle(v string) *UpdateIncidentRequest {
	s.IncidentTitle = &v
	return s
}

type UpdateIncidentResponseBody struct {
	// data
	Data *UpdateIncidentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateIncidentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateIncidentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIncidentResponseBody) SetData(v *UpdateIncidentResponseBodyData) *UpdateIncidentResponseBody {
	s.Data = v
	return s
}

func (s *UpdateIncidentResponseBody) SetRequestId(v string) *UpdateIncidentResponseBody {
	s.RequestId = &v
	return s
}

type UpdateIncidentResponseBodyData struct {
	IncidentId *int64 `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
}

func (s UpdateIncidentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateIncidentResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateIncidentResponseBodyData) SetIncidentId(v int64) *UpdateIncidentResponseBodyData {
	s.IncidentId = &v
	return s
}

type UpdateIncidentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateIncidentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateIncidentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateIncidentResponse) GoString() string {
	return s.String()
}

func (s *UpdateIncidentResponse) SetHeaders(v map[string]*string) *UpdateIncidentResponse {
	s.Headers = v
	return s
}

func (s *UpdateIncidentResponse) SetStatusCode(v int32) *UpdateIncidentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIncidentResponse) SetBody(v *UpdateIncidentResponseBody) *UpdateIncidentResponse {
	s.Body = v
	return s
}

type UpdateIntegrationConfigRequest struct {
	// accessKey
	AccessKey           *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	ClientToken         *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	IntegrationConfigId *int64  `json:"integrationConfigId,omitempty" xml:"integrationConfigId,omitempty"`
}

func (s UpdateIntegrationConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIntegrationConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateIntegrationConfigRequest) SetAccessKey(v string) *UpdateIntegrationConfigRequest {
	s.AccessKey = &v
	return s
}

func (s *UpdateIntegrationConfigRequest) SetClientToken(v string) *UpdateIntegrationConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateIntegrationConfigRequest) SetIntegrationConfigId(v int64) *UpdateIntegrationConfigRequest {
	s.IntegrationConfigId = &v
	return s
}

type UpdateIntegrationConfigResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateIntegrationConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateIntegrationConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIntegrationConfigResponseBody) SetRequestId(v string) *UpdateIntegrationConfigResponseBody {
	s.RequestId = &v
	return s
}

type UpdateIntegrationConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateIntegrationConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateIntegrationConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateIntegrationConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateIntegrationConfigResponse) SetHeaders(v map[string]*string) *UpdateIntegrationConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateIntegrationConfigResponse) SetStatusCode(v int32) *UpdateIntegrationConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIntegrationConfigResponse) SetBody(v *UpdateIntegrationConfigResponseBody) *UpdateIntegrationConfigResponse {
	s.Body = v
	return s
}

type UpdateProblemRequest struct {
	Feedback                  *string  `json:"feedback,omitempty" xml:"feedback,omitempty"`
	Level                     *string  `json:"level,omitempty" xml:"level,omitempty"`
	MainHandlerId             *int64   `json:"mainHandlerId,omitempty" xml:"mainHandlerId,omitempty"`
	PreliminaryReason         *string  `json:"preliminaryReason,omitempty" xml:"preliminaryReason,omitempty"`
	ProblemId                 *int64   `json:"problemId,omitempty" xml:"problemId,omitempty"`
	ProblemName               *string  `json:"problemName,omitempty" xml:"problemName,omitempty"`
	ProgressSummary           *string  `json:"progressSummary,omitempty" xml:"progressSummary,omitempty"`
	ProgressSummaryRichTextId *int64   `json:"progressSummaryRichTextId,omitempty" xml:"progressSummaryRichTextId,omitempty"`
	RelatedServiceId          *int64   `json:"relatedServiceId,omitempty" xml:"relatedServiceId,omitempty"`
	ServiceGroupIds           []*int64 `json:"serviceGroupIds,omitempty" xml:"serviceGroupIds,omitempty" type:"Repeated"`
}

func (s UpdateProblemRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProblemRequest) GoString() string {
	return s.String()
}

func (s *UpdateProblemRequest) SetFeedback(v string) *UpdateProblemRequest {
	s.Feedback = &v
	return s
}

func (s *UpdateProblemRequest) SetLevel(v string) *UpdateProblemRequest {
	s.Level = &v
	return s
}

func (s *UpdateProblemRequest) SetMainHandlerId(v int64) *UpdateProblemRequest {
	s.MainHandlerId = &v
	return s
}

func (s *UpdateProblemRequest) SetPreliminaryReason(v string) *UpdateProblemRequest {
	s.PreliminaryReason = &v
	return s
}

func (s *UpdateProblemRequest) SetProblemId(v int64) *UpdateProblemRequest {
	s.ProblemId = &v
	return s
}

func (s *UpdateProblemRequest) SetProblemName(v string) *UpdateProblemRequest {
	s.ProblemName = &v
	return s
}

func (s *UpdateProblemRequest) SetProgressSummary(v string) *UpdateProblemRequest {
	s.ProgressSummary = &v
	return s
}

func (s *UpdateProblemRequest) SetProgressSummaryRichTextId(v int64) *UpdateProblemRequest {
	s.ProgressSummaryRichTextId = &v
	return s
}

func (s *UpdateProblemRequest) SetRelatedServiceId(v int64) *UpdateProblemRequest {
	s.RelatedServiceId = &v
	return s
}

func (s *UpdateProblemRequest) SetServiceGroupIds(v []*int64) *UpdateProblemRequest {
	s.ServiceGroupIds = v
	return s
}

type UpdateProblemResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateProblemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProblemResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProblemResponseBody) SetRequestId(v string) *UpdateProblemResponseBody {
	s.RequestId = &v
	return s
}

type UpdateProblemResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateProblemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateProblemResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProblemResponse) GoString() string {
	return s.String()
}

func (s *UpdateProblemResponse) SetHeaders(v map[string]*string) *UpdateProblemResponse {
	s.Headers = v
	return s
}

func (s *UpdateProblemResponse) SetStatusCode(v int32) *UpdateProblemResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProblemResponse) SetBody(v *UpdateProblemResponseBody) *UpdateProblemResponse {
	s.Body = v
	return s
}

type UpdateProblemEffectionServiceRequest struct {
	// clientToken
	ClientToken        *string   `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	Description        *string   `json:"description,omitempty" xml:"description,omitempty"`
	EffectionServiceId *int64    `json:"effectionServiceId,omitempty" xml:"effectionServiceId,omitempty"`
	Level              *string   `json:"level,omitempty" xml:"level,omitempty"`
	PicUrl             []*string `json:"picUrl,omitempty" xml:"picUrl,omitempty" type:"Repeated"`
	ProblemId          *int64    `json:"problemId,omitempty" xml:"problemId,omitempty"`
	ServiceId          *int64    `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	Status             *string   `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateProblemEffectionServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProblemEffectionServiceRequest) GoString() string {
	return s.String()
}

func (s *UpdateProblemEffectionServiceRequest) SetClientToken(v string) *UpdateProblemEffectionServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateProblemEffectionServiceRequest) SetDescription(v string) *UpdateProblemEffectionServiceRequest {
	s.Description = &v
	return s
}

func (s *UpdateProblemEffectionServiceRequest) SetEffectionServiceId(v int64) *UpdateProblemEffectionServiceRequest {
	s.EffectionServiceId = &v
	return s
}

func (s *UpdateProblemEffectionServiceRequest) SetLevel(v string) *UpdateProblemEffectionServiceRequest {
	s.Level = &v
	return s
}

func (s *UpdateProblemEffectionServiceRequest) SetPicUrl(v []*string) *UpdateProblemEffectionServiceRequest {
	s.PicUrl = v
	return s
}

func (s *UpdateProblemEffectionServiceRequest) SetProblemId(v int64) *UpdateProblemEffectionServiceRequest {
	s.ProblemId = &v
	return s
}

func (s *UpdateProblemEffectionServiceRequest) SetServiceId(v int64) *UpdateProblemEffectionServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *UpdateProblemEffectionServiceRequest) SetStatus(v string) *UpdateProblemEffectionServiceRequest {
	s.Status = &v
	return s
}

type UpdateProblemEffectionServiceResponseBody struct {
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateProblemEffectionServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProblemEffectionServiceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProblemEffectionServiceResponseBody) SetRequestId(v string) *UpdateProblemEffectionServiceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateProblemEffectionServiceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateProblemEffectionServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateProblemEffectionServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProblemEffectionServiceResponse) GoString() string {
	return s.String()
}

func (s *UpdateProblemEffectionServiceResponse) SetHeaders(v map[string]*string) *UpdateProblemEffectionServiceResponse {
	s.Headers = v
	return s
}

func (s *UpdateProblemEffectionServiceResponse) SetStatusCode(v int32) *UpdateProblemEffectionServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProblemEffectionServiceResponse) SetBody(v *UpdateProblemEffectionServiceResponseBody) *UpdateProblemEffectionServiceResponse {
	s.Body = v
	return s
}

type UpdateProblemImprovementRequest struct {
	ClientToken         *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	CustomProblemReason *string `json:"customProblemReason,omitempty" xml:"customProblemReason,omitempty"`
	DiscoverSource      *int64  `json:"discoverSource,omitempty" xml:"discoverSource,omitempty"`
	DutyDepartmentId    *int64  `json:"dutyDepartmentId,omitempty" xml:"dutyDepartmentId,omitempty"`
	DutyDepartmentName  *string `json:"dutyDepartmentName,omitempty" xml:"dutyDepartmentName,omitempty"`
	DutyUserId          *int64  `json:"dutyUserId,omitempty" xml:"dutyUserId,omitempty"`
	InjectionMode       *string `json:"injectionMode,omitempty" xml:"injectionMode,omitempty"`
	MonitorSourceName   *string `json:"monitorSourceName,omitempty" xml:"monitorSourceName,omitempty"`
	ProblemId           *int64  `json:"problemId,omitempty" xml:"problemId,omitempty"`
	ProblemReason       *string `json:"problemReason,omitempty" xml:"problemReason,omitempty"`
	RecentActivity      *string `json:"recentActivity,omitempty" xml:"recentActivity,omitempty"`
	RecoveryMode        *string `json:"recoveryMode,omitempty" xml:"recoveryMode,omitempty"`
	RelationChanges     *string `json:"relationChanges,omitempty" xml:"relationChanges,omitempty"`
	Remark              *string `json:"remark,omitempty" xml:"remark,omitempty"`
	ReplayDutyUserId    *int64  `json:"replayDutyUserId,omitempty" xml:"replayDutyUserId,omitempty"`
	UserReport          *int64  `json:"userReport,omitempty" xml:"userReport,omitempty"`
}

func (s UpdateProblemImprovementRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProblemImprovementRequest) GoString() string {
	return s.String()
}

func (s *UpdateProblemImprovementRequest) SetClientToken(v string) *UpdateProblemImprovementRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateProblemImprovementRequest) SetCustomProblemReason(v string) *UpdateProblemImprovementRequest {
	s.CustomProblemReason = &v
	return s
}

func (s *UpdateProblemImprovementRequest) SetDiscoverSource(v int64) *UpdateProblemImprovementRequest {
	s.DiscoverSource = &v
	return s
}

func (s *UpdateProblemImprovementRequest) SetDutyDepartmentId(v int64) *UpdateProblemImprovementRequest {
	s.DutyDepartmentId = &v
	return s
}

func (s *UpdateProblemImprovementRequest) SetDutyDepartmentName(v string) *UpdateProblemImprovementRequest {
	s.DutyDepartmentName = &v
	return s
}

func (s *UpdateProblemImprovementRequest) SetDutyUserId(v int64) *UpdateProblemImprovementRequest {
	s.DutyUserId = &v
	return s
}

func (s *UpdateProblemImprovementRequest) SetInjectionMode(v string) *UpdateProblemImprovementRequest {
	s.InjectionMode = &v
	return s
}

func (s *UpdateProblemImprovementRequest) SetMonitorSourceName(v string) *UpdateProblemImprovementRequest {
	s.MonitorSourceName = &v
	return s
}

func (s *UpdateProblemImprovementRequest) SetProblemId(v int64) *UpdateProblemImprovementRequest {
	s.ProblemId = &v
	return s
}

func (s *UpdateProblemImprovementRequest) SetProblemReason(v string) *UpdateProblemImprovementRequest {
	s.ProblemReason = &v
	return s
}

func (s *UpdateProblemImprovementRequest) SetRecentActivity(v string) *UpdateProblemImprovementRequest {
	s.RecentActivity = &v
	return s
}

func (s *UpdateProblemImprovementRequest) SetRecoveryMode(v string) *UpdateProblemImprovementRequest {
	s.RecoveryMode = &v
	return s
}

func (s *UpdateProblemImprovementRequest) SetRelationChanges(v string) *UpdateProblemImprovementRequest {
	s.RelationChanges = &v
	return s
}

func (s *UpdateProblemImprovementRequest) SetRemark(v string) *UpdateProblemImprovementRequest {
	s.Remark = &v
	return s
}

func (s *UpdateProblemImprovementRequest) SetReplayDutyUserId(v int64) *UpdateProblemImprovementRequest {
	s.ReplayDutyUserId = &v
	return s
}

func (s *UpdateProblemImprovementRequest) SetUserReport(v int64) *UpdateProblemImprovementRequest {
	s.UserReport = &v
	return s
}

type UpdateProblemImprovementResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateProblemImprovementResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProblemImprovementResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProblemImprovementResponseBody) SetRequestId(v string) *UpdateProblemImprovementResponseBody {
	s.RequestId = &v
	return s
}

type UpdateProblemImprovementResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateProblemImprovementResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateProblemImprovementResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProblemImprovementResponse) GoString() string {
	return s.String()
}

func (s *UpdateProblemImprovementResponse) SetHeaders(v map[string]*string) *UpdateProblemImprovementResponse {
	s.Headers = v
	return s
}

func (s *UpdateProblemImprovementResponse) SetStatusCode(v int32) *UpdateProblemImprovementResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProblemImprovementResponse) SetBody(v *UpdateProblemImprovementResponseBody) *UpdateProblemImprovementResponse {
	s.Body = v
	return s
}

type UpdateProblemMeasureRequest struct {
	CheckStandard  *string `json:"checkStandard,omitempty" xml:"checkStandard,omitempty"`
	CheckUserId    *int64  `json:"checkUserId,omitempty" xml:"checkUserId,omitempty"`
	ClientToken    *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	Content        *string `json:"content,omitempty" xml:"content,omitempty"`
	DirectorId     *int64  `json:"directorId,omitempty" xml:"directorId,omitempty"`
	MeasureId      *int64  `json:"measureId,omitempty" xml:"measureId,omitempty"`
	PlanFinishTime *string `json:"planFinishTime,omitempty" xml:"planFinishTime,omitempty"`
	ProblemId      *int64  `json:"problemId,omitempty" xml:"problemId,omitempty"`
	StalkerId      *int64  `json:"stalkerId,omitempty" xml:"stalkerId,omitempty"`
	Status         *string `json:"status,omitempty" xml:"status,omitempty"`
	Type           *int32  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateProblemMeasureRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProblemMeasureRequest) GoString() string {
	return s.String()
}

func (s *UpdateProblemMeasureRequest) SetCheckStandard(v string) *UpdateProblemMeasureRequest {
	s.CheckStandard = &v
	return s
}

func (s *UpdateProblemMeasureRequest) SetCheckUserId(v int64) *UpdateProblemMeasureRequest {
	s.CheckUserId = &v
	return s
}

func (s *UpdateProblemMeasureRequest) SetClientToken(v string) *UpdateProblemMeasureRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateProblemMeasureRequest) SetContent(v string) *UpdateProblemMeasureRequest {
	s.Content = &v
	return s
}

func (s *UpdateProblemMeasureRequest) SetDirectorId(v int64) *UpdateProblemMeasureRequest {
	s.DirectorId = &v
	return s
}

func (s *UpdateProblemMeasureRequest) SetMeasureId(v int64) *UpdateProblemMeasureRequest {
	s.MeasureId = &v
	return s
}

func (s *UpdateProblemMeasureRequest) SetPlanFinishTime(v string) *UpdateProblemMeasureRequest {
	s.PlanFinishTime = &v
	return s
}

func (s *UpdateProblemMeasureRequest) SetProblemId(v int64) *UpdateProblemMeasureRequest {
	s.ProblemId = &v
	return s
}

func (s *UpdateProblemMeasureRequest) SetStalkerId(v int64) *UpdateProblemMeasureRequest {
	s.StalkerId = &v
	return s
}

func (s *UpdateProblemMeasureRequest) SetStatus(v string) *UpdateProblemMeasureRequest {
	s.Status = &v
	return s
}

func (s *UpdateProblemMeasureRequest) SetType(v int32) *UpdateProblemMeasureRequest {
	s.Type = &v
	return s
}

type UpdateProblemMeasureResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateProblemMeasureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProblemMeasureResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProblemMeasureResponseBody) SetRequestId(v string) *UpdateProblemMeasureResponseBody {
	s.RequestId = &v
	return s
}

type UpdateProblemMeasureResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateProblemMeasureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateProblemMeasureResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProblemMeasureResponse) GoString() string {
	return s.String()
}

func (s *UpdateProblemMeasureResponse) SetHeaders(v map[string]*string) *UpdateProblemMeasureResponse {
	s.Headers = v
	return s
}

func (s *UpdateProblemMeasureResponse) SetStatusCode(v int32) *UpdateProblemMeasureResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProblemMeasureResponse) SetBody(v *UpdateProblemMeasureResponseBody) *UpdateProblemMeasureResponse {
	s.Body = v
	return s
}

type UpdateProblemNoticeRequest struct {
	ClientToken       *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	ProblemId         *int64  `json:"problemId,omitempty" xml:"problemId,omitempty"`
	ProblemNotifyType *string `json:"problemNotifyType,omitempty" xml:"problemNotifyType,omitempty"`
}

func (s UpdateProblemNoticeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProblemNoticeRequest) GoString() string {
	return s.String()
}

func (s *UpdateProblemNoticeRequest) SetClientToken(v string) *UpdateProblemNoticeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateProblemNoticeRequest) SetProblemId(v int64) *UpdateProblemNoticeRequest {
	s.ProblemId = &v
	return s
}

func (s *UpdateProblemNoticeRequest) SetProblemNotifyType(v string) *UpdateProblemNoticeRequest {
	s.ProblemNotifyType = &v
	return s
}

type UpdateProblemNoticeResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateProblemNoticeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProblemNoticeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProblemNoticeResponseBody) SetRequestId(v string) *UpdateProblemNoticeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateProblemNoticeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateProblemNoticeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateProblemNoticeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProblemNoticeResponse) GoString() string {
	return s.String()
}

func (s *UpdateProblemNoticeResponse) SetHeaders(v map[string]*string) *UpdateProblemNoticeResponse {
	s.Headers = v
	return s
}

func (s *UpdateProblemNoticeResponse) SetStatusCode(v int32) *UpdateProblemNoticeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProblemNoticeResponse) SetBody(v *UpdateProblemNoticeResponseBody) *UpdateProblemNoticeResponse {
	s.Body = v
	return s
}

type UpdateProblemTimelineRequest struct {
	ClientToken       *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	Content           *string `json:"content,omitempty" xml:"content,omitempty"`
	KeyNode           *string `json:"keyNode,omitempty" xml:"keyNode,omitempty"`
	ProblemId         *int64  `json:"problemId,omitempty" xml:"problemId,omitempty"`
	ProblemTimelineId *int64  `json:"problemTimelineId,omitempty" xml:"problemTimelineId,omitempty"`
	Time              *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s UpdateProblemTimelineRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProblemTimelineRequest) GoString() string {
	return s.String()
}

func (s *UpdateProblemTimelineRequest) SetClientToken(v string) *UpdateProblemTimelineRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateProblemTimelineRequest) SetContent(v string) *UpdateProblemTimelineRequest {
	s.Content = &v
	return s
}

func (s *UpdateProblemTimelineRequest) SetKeyNode(v string) *UpdateProblemTimelineRequest {
	s.KeyNode = &v
	return s
}

func (s *UpdateProblemTimelineRequest) SetProblemId(v int64) *UpdateProblemTimelineRequest {
	s.ProblemId = &v
	return s
}

func (s *UpdateProblemTimelineRequest) SetProblemTimelineId(v int64) *UpdateProblemTimelineRequest {
	s.ProblemTimelineId = &v
	return s
}

func (s *UpdateProblemTimelineRequest) SetTime(v string) *UpdateProblemTimelineRequest {
	s.Time = &v
	return s
}

type UpdateProblemTimelineResponseBody struct {
	// id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateProblemTimelineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProblemTimelineResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProblemTimelineResponseBody) SetRequestId(v string) *UpdateProblemTimelineResponseBody {
	s.RequestId = &v
	return s
}

type UpdateProblemTimelineResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateProblemTimelineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateProblemTimelineResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProblemTimelineResponse) GoString() string {
	return s.String()
}

func (s *UpdateProblemTimelineResponse) SetHeaders(v map[string]*string) *UpdateProblemTimelineResponse {
	s.Headers = v
	return s
}

func (s *UpdateProblemTimelineResponse) SetStatusCode(v int32) *UpdateProblemTimelineResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProblemTimelineResponse) SetBody(v *UpdateProblemTimelineResponseBody) *UpdateProblemTimelineResponse {
	s.Body = v
	return s
}

type UpdateRichTextRequest struct {
	InstanceId   *int64  `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	RichText     *string `json:"richText,omitempty" xml:"richText,omitempty"`
	RichTextId   *int64  `json:"richTextId,omitempty" xml:"richTextId,omitempty"`
}

func (s UpdateRichTextRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRichTextRequest) GoString() string {
	return s.String()
}

func (s *UpdateRichTextRequest) SetInstanceId(v int64) *UpdateRichTextRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateRichTextRequest) SetInstanceType(v string) *UpdateRichTextRequest {
	s.InstanceType = &v
	return s
}

func (s *UpdateRichTextRequest) SetRichText(v string) *UpdateRichTextRequest {
	s.RichText = &v
	return s
}

func (s *UpdateRichTextRequest) SetRichTextId(v int64) *UpdateRichTextRequest {
	s.RichTextId = &v
	return s
}

type UpdateRichTextResponseBody struct {
	// data
	Data *UpdateRichTextResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateRichTextResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRichTextResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRichTextResponseBody) SetData(v *UpdateRichTextResponseBodyData) *UpdateRichTextResponseBody {
	s.Data = v
	return s
}

func (s *UpdateRichTextResponseBody) SetRequestId(v string) *UpdateRichTextResponseBody {
	s.RequestId = &v
	return s
}

type UpdateRichTextResponseBodyData struct {
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s UpdateRichTextResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateRichTextResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateRichTextResponseBodyData) SetId(v int64) *UpdateRichTextResponseBodyData {
	s.Id = &v
	return s
}

type UpdateRichTextResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateRichTextResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRichTextResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRichTextResponse) GoString() string {
	return s.String()
}

func (s *UpdateRichTextResponse) SetHeaders(v map[string]*string) *UpdateRichTextResponse {
	s.Headers = v
	return s
}

func (s *UpdateRichTextResponse) SetStatusCode(v int32) *UpdateRichTextResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRichTextResponse) SetBody(v *UpdateRichTextResponseBody) *UpdateRichTextResponse {
	s.Body = v
	return s
}

type UpdateRouteRuleRequest struct {
	AssignObjectId           *int64                                   `json:"assignObjectId,omitempty" xml:"assignObjectId,omitempty"`
	AssignObjectType         *string                                  `json:"assignObjectType,omitempty" xml:"assignObjectType,omitempty"`
	ChildRuleRelation        *string                                  `json:"childRuleRelation,omitempty" xml:"childRuleRelation,omitempty"`
	ClientToken              *string                                  `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	CoverageProblemLevels    []*string                                `json:"coverageProblemLevels,omitempty" xml:"coverageProblemLevels,omitempty" type:"Repeated"`
	Effection                *string                                  `json:"effection,omitempty" xml:"effection,omitempty"`
	IncidentLevel            *string                                  `json:"incidentLevel,omitempty" xml:"incidentLevel,omitempty"`
	MatchCount               *int64                                   `json:"matchCount,omitempty" xml:"matchCount,omitempty"`
	NotifyChannels           []*string                                `json:"notifyChannels,omitempty" xml:"notifyChannels,omitempty" type:"Repeated"`
	ProblemEffectionServices []*int64                                 `json:"problemEffectionServices,omitempty" xml:"problemEffectionServices,omitempty" type:"Repeated"`
	ProblemLevelGroup        map[string]*ProblemLevelGroupValue       `json:"problemLevelGroup,omitempty" xml:"problemLevelGroup,omitempty"`
	RelatedServiceId         *int64                                   `json:"relatedServiceId,omitempty" xml:"relatedServiceId,omitempty"`
	RouteChildRules          []*UpdateRouteRuleRequestRouteChildRules `json:"routeChildRules,omitempty" xml:"routeChildRules,omitempty" type:"Repeated"`
	RouteRuleId              *int64                                   `json:"routeRuleId,omitempty" xml:"routeRuleId,omitempty"`
	RouteType                *string                                  `json:"routeType,omitempty" xml:"routeType,omitempty"`
	RuleName                 *string                                  `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	TimeWindow               *int32                                   `json:"timeWindow,omitempty" xml:"timeWindow,omitempty"`
	TimeWindowUnit           *string                                  `json:"timeWindowUnit,omitempty" xml:"timeWindowUnit,omitempty"`
}

func (s UpdateRouteRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRouteRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateRouteRuleRequest) SetAssignObjectId(v int64) *UpdateRouteRuleRequest {
	s.AssignObjectId = &v
	return s
}

func (s *UpdateRouteRuleRequest) SetAssignObjectType(v string) *UpdateRouteRuleRequest {
	s.AssignObjectType = &v
	return s
}

func (s *UpdateRouteRuleRequest) SetChildRuleRelation(v string) *UpdateRouteRuleRequest {
	s.ChildRuleRelation = &v
	return s
}

func (s *UpdateRouteRuleRequest) SetClientToken(v string) *UpdateRouteRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateRouteRuleRequest) SetCoverageProblemLevels(v []*string) *UpdateRouteRuleRequest {
	s.CoverageProblemLevels = v
	return s
}

func (s *UpdateRouteRuleRequest) SetEffection(v string) *UpdateRouteRuleRequest {
	s.Effection = &v
	return s
}

func (s *UpdateRouteRuleRequest) SetIncidentLevel(v string) *UpdateRouteRuleRequest {
	s.IncidentLevel = &v
	return s
}

func (s *UpdateRouteRuleRequest) SetMatchCount(v int64) *UpdateRouteRuleRequest {
	s.MatchCount = &v
	return s
}

func (s *UpdateRouteRuleRequest) SetNotifyChannels(v []*string) *UpdateRouteRuleRequest {
	s.NotifyChannels = v
	return s
}

func (s *UpdateRouteRuleRequest) SetProblemEffectionServices(v []*int64) *UpdateRouteRuleRequest {
	s.ProblemEffectionServices = v
	return s
}

func (s *UpdateRouteRuleRequest) SetProblemLevelGroup(v map[string]*ProblemLevelGroupValue) *UpdateRouteRuleRequest {
	s.ProblemLevelGroup = v
	return s
}

func (s *UpdateRouteRuleRequest) SetRelatedServiceId(v int64) *UpdateRouteRuleRequest {
	s.RelatedServiceId = &v
	return s
}

func (s *UpdateRouteRuleRequest) SetRouteChildRules(v []*UpdateRouteRuleRequestRouteChildRules) *UpdateRouteRuleRequest {
	s.RouteChildRules = v
	return s
}

func (s *UpdateRouteRuleRequest) SetRouteRuleId(v int64) *UpdateRouteRuleRequest {
	s.RouteRuleId = &v
	return s
}

func (s *UpdateRouteRuleRequest) SetRouteType(v string) *UpdateRouteRuleRequest {
	s.RouteType = &v
	return s
}

func (s *UpdateRouteRuleRequest) SetRuleName(v string) *UpdateRouteRuleRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateRouteRuleRequest) SetTimeWindow(v int32) *UpdateRouteRuleRequest {
	s.TimeWindow = &v
	return s
}

func (s *UpdateRouteRuleRequest) SetTimeWindowUnit(v string) *UpdateRouteRuleRequest {
	s.TimeWindowUnit = &v
	return s
}

type UpdateRouteRuleRequestRouteChildRules struct {
	ChildConditionRelation *int64                                             `json:"childConditionRelation,omitempty" xml:"childConditionRelation,omitempty"`
	ChildRouteRuleId       *int64                                             `json:"childRouteRuleId,omitempty" xml:"childRouteRuleId,omitempty"`
	Conditions             []*UpdateRouteRuleRequestRouteChildRulesConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	IsValidChildRule       *bool                                              `json:"isValidChildRule,omitempty" xml:"isValidChildRule,omitempty"`
	MonitorSourceId        *int64                                             `json:"monitorSourceId,omitempty" xml:"monitorSourceId,omitempty"`
	ProblemLevel           *string                                            `json:"problemLevel,omitempty" xml:"problemLevel,omitempty"`
}

func (s UpdateRouteRuleRequestRouteChildRules) String() string {
	return tea.Prettify(s)
}

func (s UpdateRouteRuleRequestRouteChildRules) GoString() string {
	return s.String()
}

func (s *UpdateRouteRuleRequestRouteChildRules) SetChildConditionRelation(v int64) *UpdateRouteRuleRequestRouteChildRules {
	s.ChildConditionRelation = &v
	return s
}

func (s *UpdateRouteRuleRequestRouteChildRules) SetChildRouteRuleId(v int64) *UpdateRouteRuleRequestRouteChildRules {
	s.ChildRouteRuleId = &v
	return s
}

func (s *UpdateRouteRuleRequestRouteChildRules) SetConditions(v []*UpdateRouteRuleRequestRouteChildRulesConditions) *UpdateRouteRuleRequestRouteChildRules {
	s.Conditions = v
	return s
}

func (s *UpdateRouteRuleRequestRouteChildRules) SetIsValidChildRule(v bool) *UpdateRouteRuleRequestRouteChildRules {
	s.IsValidChildRule = &v
	return s
}

func (s *UpdateRouteRuleRequestRouteChildRules) SetMonitorSourceId(v int64) *UpdateRouteRuleRequestRouteChildRules {
	s.MonitorSourceId = &v
	return s
}

func (s *UpdateRouteRuleRequestRouteChildRules) SetProblemLevel(v string) *UpdateRouteRuleRequestRouteChildRules {
	s.ProblemLevel = &v
	return s
}

type UpdateRouteRuleRequestRouteChildRulesConditions struct {
	Key             *string `json:"key,omitempty" xml:"key,omitempty"`
	OperationSymbol *string `json:"operationSymbol,omitempty" xml:"operationSymbol,omitempty"`
	Value           *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateRouteRuleRequestRouteChildRulesConditions) String() string {
	return tea.Prettify(s)
}

func (s UpdateRouteRuleRequestRouteChildRulesConditions) GoString() string {
	return s.String()
}

func (s *UpdateRouteRuleRequestRouteChildRulesConditions) SetKey(v string) *UpdateRouteRuleRequestRouteChildRulesConditions {
	s.Key = &v
	return s
}

func (s *UpdateRouteRuleRequestRouteChildRulesConditions) SetOperationSymbol(v string) *UpdateRouteRuleRequestRouteChildRulesConditions {
	s.OperationSymbol = &v
	return s
}

func (s *UpdateRouteRuleRequestRouteChildRulesConditions) SetValue(v string) *UpdateRouteRuleRequestRouteChildRulesConditions {
	s.Value = &v
	return s
}

type UpdateRouteRuleResponseBody struct {
	Data      *int64  `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateRouteRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRouteRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRouteRuleResponseBody) SetData(v int64) *UpdateRouteRuleResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateRouteRuleResponseBody) SetRequestId(v string) *UpdateRouteRuleResponseBody {
	s.RequestId = &v
	return s
}

type UpdateRouteRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateRouteRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRouteRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRouteRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateRouteRuleResponse) SetHeaders(v map[string]*string) *UpdateRouteRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateRouteRuleResponse) SetStatusCode(v int32) *UpdateRouteRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRouteRuleResponse) SetBody(v *UpdateRouteRuleResponseBody) *UpdateRouteRuleResponse {
	s.Body = v
	return s
}

type UpdateServiceRequest struct {
	ClientToken        *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	ServiceDescription *string `json:"serviceDescription,omitempty" xml:"serviceDescription,omitempty"`
	ServiceId          *int64  `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	ServiceName        *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
}

func (s UpdateServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequest) SetClientToken(v string) *UpdateServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServiceRequest) SetServiceDescription(v string) *UpdateServiceRequest {
	s.ServiceDescription = &v
	return s
}

func (s *UpdateServiceRequest) SetServiceId(v int64) *UpdateServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *UpdateServiceRequest) SetServiceName(v string) *UpdateServiceRequest {
	s.ServiceName = &v
	return s
}

type UpdateServiceResponseBody struct {
	Data *int64 `json:"data,omitempty" xml:"data,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceResponseBody) SetData(v int64) *UpdateServiceResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateServiceResponseBody) SetRequestId(v string) *UpdateServiceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceResponse) SetHeaders(v map[string]*string) *UpdateServiceResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceResponse) SetStatusCode(v int32) *UpdateServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceResponse) SetBody(v *UpdateServiceResponseBody) *UpdateServiceResponse {
	s.Body = v
	return s
}

type UpdateServiceGroupRequest struct {
	ClientToken             *string                                            `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	EnableWebhook           *string                                            `json:"enableWebhook,omitempty" xml:"enableWebhook,omitempty"`
	MonitorSourceTemplates  []*UpdateServiceGroupRequestMonitorSourceTemplates `json:"monitorSourceTemplates,omitempty" xml:"monitorSourceTemplates,omitempty" type:"Repeated"`
	ServiceGroupDescription *string                                            `json:"serviceGroupDescription,omitempty" xml:"serviceGroupDescription,omitempty"`
	ServiceGroupId          *int64                                             `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
	ServiceGroupName        *string                                            `json:"serviceGroupName,omitempty" xml:"serviceGroupName,omitempty"`
	UserIds                 []*int64                                           `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
	WebhookLink             *string                                            `json:"webhookLink,omitempty" xml:"webhookLink,omitempty"`
	WebhookType             *string                                            `json:"webhookType,omitempty" xml:"webhookType,omitempty"`
}

func (s UpdateServiceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceGroupRequest) SetClientToken(v string) *UpdateServiceGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServiceGroupRequest) SetEnableWebhook(v string) *UpdateServiceGroupRequest {
	s.EnableWebhook = &v
	return s
}

func (s *UpdateServiceGroupRequest) SetMonitorSourceTemplates(v []*UpdateServiceGroupRequestMonitorSourceTemplates) *UpdateServiceGroupRequest {
	s.MonitorSourceTemplates = v
	return s
}

func (s *UpdateServiceGroupRequest) SetServiceGroupDescription(v string) *UpdateServiceGroupRequest {
	s.ServiceGroupDescription = &v
	return s
}

func (s *UpdateServiceGroupRequest) SetServiceGroupId(v int64) *UpdateServiceGroupRequest {
	s.ServiceGroupId = &v
	return s
}

func (s *UpdateServiceGroupRequest) SetServiceGroupName(v string) *UpdateServiceGroupRequest {
	s.ServiceGroupName = &v
	return s
}

func (s *UpdateServiceGroupRequest) SetUserIds(v []*int64) *UpdateServiceGroupRequest {
	s.UserIds = v
	return s
}

func (s *UpdateServiceGroupRequest) SetWebhookLink(v string) *UpdateServiceGroupRequest {
	s.WebhookLink = &v
	return s
}

func (s *UpdateServiceGroupRequest) SetWebhookType(v string) *UpdateServiceGroupRequest {
	s.WebhookType = &v
	return s
}

type UpdateServiceGroupRequestMonitorSourceTemplates struct {
	MonitorSourceId   *int64  `json:"monitorSourceId,omitempty" xml:"monitorSourceId,omitempty"`
	MonitorSourceName *string `json:"monitorSourceName,omitempty" xml:"monitorSourceName,omitempty"`
	TemplateContent   *string `json:"templateContent,omitempty" xml:"templateContent,omitempty"`
	TemplateId        *int64  `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s UpdateServiceGroupRequestMonitorSourceTemplates) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceGroupRequestMonitorSourceTemplates) GoString() string {
	return s.String()
}

func (s *UpdateServiceGroupRequestMonitorSourceTemplates) SetMonitorSourceId(v int64) *UpdateServiceGroupRequestMonitorSourceTemplates {
	s.MonitorSourceId = &v
	return s
}

func (s *UpdateServiceGroupRequestMonitorSourceTemplates) SetMonitorSourceName(v string) *UpdateServiceGroupRequestMonitorSourceTemplates {
	s.MonitorSourceName = &v
	return s
}

func (s *UpdateServiceGroupRequestMonitorSourceTemplates) SetTemplateContent(v string) *UpdateServiceGroupRequestMonitorSourceTemplates {
	s.TemplateContent = &v
	return s
}

func (s *UpdateServiceGroupRequestMonitorSourceTemplates) SetTemplateId(v int64) *UpdateServiceGroupRequestMonitorSourceTemplates {
	s.TemplateId = &v
	return s
}

type UpdateServiceGroupResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateServiceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceGroupResponseBody) SetRequestId(v string) *UpdateServiceGroupResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServiceGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateServiceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateServiceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceGroupResponse) SetHeaders(v map[string]*string) *UpdateServiceGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceGroupResponse) SetStatusCode(v int32) *UpdateServiceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceGroupResponse) SetBody(v *UpdateServiceGroupResponseBody) *UpdateServiceGroupResponse {
	s.Body = v
	return s
}

type UpdateServiceGroupSchedulingRequest struct {
	ClientToken    *string                                            `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	FastScheduling *UpdateServiceGroupSchedulingRequestFastScheduling `json:"fastScheduling,omitempty" xml:"fastScheduling,omitempty" type:"Struct"`
	FineScheduling *UpdateServiceGroupSchedulingRequestFineScheduling `json:"fineScheduling,omitempty" xml:"fineScheduling,omitempty" type:"Struct"`
	SchedulingWay  *string                                            `json:"schedulingWay,omitempty" xml:"schedulingWay,omitempty"`
	ServiceGroupId *int64                                             `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
}

func (s UpdateServiceGroupSchedulingRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceGroupSchedulingRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceGroupSchedulingRequest) SetClientToken(v string) *UpdateServiceGroupSchedulingRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServiceGroupSchedulingRequest) SetFastScheduling(v *UpdateServiceGroupSchedulingRequestFastScheduling) *UpdateServiceGroupSchedulingRequest {
	s.FastScheduling = v
	return s
}

func (s *UpdateServiceGroupSchedulingRequest) SetFineScheduling(v *UpdateServiceGroupSchedulingRequestFineScheduling) *UpdateServiceGroupSchedulingRequest {
	s.FineScheduling = v
	return s
}

func (s *UpdateServiceGroupSchedulingRequest) SetSchedulingWay(v string) *UpdateServiceGroupSchedulingRequest {
	s.SchedulingWay = &v
	return s
}

func (s *UpdateServiceGroupSchedulingRequest) SetServiceGroupId(v int64) *UpdateServiceGroupSchedulingRequest {
	s.ServiceGroupId = &v
	return s
}

type UpdateServiceGroupSchedulingRequestFastScheduling struct {
	DutyPlan           *string                                                             `json:"dutyPlan,omitempty" xml:"dutyPlan,omitempty"`
	Id                 *int64                                                              `json:"id,omitempty" xml:"id,omitempty"`
	SchedulingUsers    []*UpdateServiceGroupSchedulingRequestFastSchedulingSchedulingUsers `json:"schedulingUsers,omitempty" xml:"schedulingUsers,omitempty" type:"Repeated"`
	SingleDuration     *int32                                                              `json:"singleDuration,omitempty" xml:"singleDuration,omitempty"`
	SingleDurationUnit *string                                                             `json:"singleDurationUnit,omitempty" xml:"singleDurationUnit,omitempty"`
}

func (s UpdateServiceGroupSchedulingRequestFastScheduling) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceGroupSchedulingRequestFastScheduling) GoString() string {
	return s.String()
}

func (s *UpdateServiceGroupSchedulingRequestFastScheduling) SetDutyPlan(v string) *UpdateServiceGroupSchedulingRequestFastScheduling {
	s.DutyPlan = &v
	return s
}

func (s *UpdateServiceGroupSchedulingRequestFastScheduling) SetId(v int64) *UpdateServiceGroupSchedulingRequestFastScheduling {
	s.Id = &v
	return s
}

func (s *UpdateServiceGroupSchedulingRequestFastScheduling) SetSchedulingUsers(v []*UpdateServiceGroupSchedulingRequestFastSchedulingSchedulingUsers) *UpdateServiceGroupSchedulingRequestFastScheduling {
	s.SchedulingUsers = v
	return s
}

func (s *UpdateServiceGroupSchedulingRequestFastScheduling) SetSingleDuration(v int32) *UpdateServiceGroupSchedulingRequestFastScheduling {
	s.SingleDuration = &v
	return s
}

func (s *UpdateServiceGroupSchedulingRequestFastScheduling) SetSingleDurationUnit(v string) *UpdateServiceGroupSchedulingRequestFastScheduling {
	s.SingleDurationUnit = &v
	return s
}

type UpdateServiceGroupSchedulingRequestFastSchedulingSchedulingUsers struct {
	SchedulingOrder  *int32 `json:"schedulingOrder,omitempty" xml:"schedulingOrder,omitempty"`
	SchedulingUserId *int64 `json:"schedulingUserId,omitempty" xml:"schedulingUserId,omitempty"`
}

func (s UpdateServiceGroupSchedulingRequestFastSchedulingSchedulingUsers) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceGroupSchedulingRequestFastSchedulingSchedulingUsers) GoString() string {
	return s.String()
}

func (s *UpdateServiceGroupSchedulingRequestFastSchedulingSchedulingUsers) SetSchedulingOrder(v int32) *UpdateServiceGroupSchedulingRequestFastSchedulingSchedulingUsers {
	s.SchedulingOrder = &v
	return s
}

func (s *UpdateServiceGroupSchedulingRequestFastSchedulingSchedulingUsers) SetSchedulingUserId(v int64) *UpdateServiceGroupSchedulingRequestFastSchedulingSchedulingUsers {
	s.SchedulingUserId = &v
	return s
}

type UpdateServiceGroupSchedulingRequestFineScheduling struct {
	Id                           *int64                                                                           `json:"id,omitempty" xml:"id,omitempty"`
	Period                       *int32                                                                           `json:"period,omitempty" xml:"period,omitempty"`
	PeriodUnit                   *string                                                                          `json:"periodUnit,omitempty" xml:"periodUnit,omitempty"`
	SchedulingFineShifts         []*UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts         `json:"schedulingFineShifts,omitempty" xml:"schedulingFineShifts,omitempty" type:"Repeated"`
	SchedulingTemplateFineShifts []*UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingTemplateFineShifts `json:"schedulingTemplateFineShifts,omitempty" xml:"schedulingTemplateFineShifts,omitempty" type:"Repeated"`
	ShiftType                    *string                                                                          `json:"shiftType,omitempty" xml:"shiftType,omitempty"`
}

func (s UpdateServiceGroupSchedulingRequestFineScheduling) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceGroupSchedulingRequestFineScheduling) GoString() string {
	return s.String()
}

func (s *UpdateServiceGroupSchedulingRequestFineScheduling) SetId(v int64) *UpdateServiceGroupSchedulingRequestFineScheduling {
	s.Id = &v
	return s
}

func (s *UpdateServiceGroupSchedulingRequestFineScheduling) SetPeriod(v int32) *UpdateServiceGroupSchedulingRequestFineScheduling {
	s.Period = &v
	return s
}

func (s *UpdateServiceGroupSchedulingRequestFineScheduling) SetPeriodUnit(v string) *UpdateServiceGroupSchedulingRequestFineScheduling {
	s.PeriodUnit = &v
	return s
}

func (s *UpdateServiceGroupSchedulingRequestFineScheduling) SetSchedulingFineShifts(v []*UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts) *UpdateServiceGroupSchedulingRequestFineScheduling {
	s.SchedulingFineShifts = v
	return s
}

func (s *UpdateServiceGroupSchedulingRequestFineScheduling) SetSchedulingTemplateFineShifts(v []*UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingTemplateFineShifts) *UpdateServiceGroupSchedulingRequestFineScheduling {
	s.SchedulingTemplateFineShifts = v
	return s
}

func (s *UpdateServiceGroupSchedulingRequestFineScheduling) SetShiftType(v string) *UpdateServiceGroupSchedulingRequestFineScheduling {
	s.ShiftType = &v
	return s
}

type UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts struct {
	CycleOrder          *int32  `json:"cycleOrder,omitempty" xml:"cycleOrder,omitempty"`
	SchedulingEndTime   *string `json:"schedulingEndTime,omitempty" xml:"schedulingEndTime,omitempty"`
	SchedulingOrder     *int64  `json:"schedulingOrder,omitempty" xml:"schedulingOrder,omitempty"`
	SchedulingStartTime *string `json:"schedulingStartTime,omitempty" xml:"schedulingStartTime,omitempty"`
	SchedulingUserId    *int64  `json:"schedulingUserId,omitempty" xml:"schedulingUserId,omitempty"`
	ShiftName           *string `json:"shiftName,omitempty" xml:"shiftName,omitempty"`
	SkipOneDay          *bool   `json:"skipOneDay,omitempty" xml:"skipOneDay,omitempty"`
}

func (s UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts) GoString() string {
	return s.String()
}

func (s *UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts) SetCycleOrder(v int32) *UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts {
	s.CycleOrder = &v
	return s
}

func (s *UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts) SetSchedulingEndTime(v string) *UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts {
	s.SchedulingEndTime = &v
	return s
}

func (s *UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts) SetSchedulingOrder(v int64) *UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts {
	s.SchedulingOrder = &v
	return s
}

func (s *UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts) SetSchedulingStartTime(v string) *UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts {
	s.SchedulingStartTime = &v
	return s
}

func (s *UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts) SetSchedulingUserId(v int64) *UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts {
	s.SchedulingUserId = &v
	return s
}

func (s *UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts) SetShiftName(v string) *UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts {
	s.ShiftName = &v
	return s
}

func (s *UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts) SetSkipOneDay(v bool) *UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingFineShifts {
	s.SkipOneDay = &v
	return s
}

type UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingTemplateFineShifts struct {
	SchedulingEndTime   *string `json:"schedulingEndTime,omitempty" xml:"schedulingEndTime,omitempty"`
	SchedulingOrder     *int32  `json:"schedulingOrder,omitempty" xml:"schedulingOrder,omitempty"`
	SchedulingStartTime *string `json:"schedulingStartTime,omitempty" xml:"schedulingStartTime,omitempty"`
	SchedulingUserId    *int64  `json:"schedulingUserId,omitempty" xml:"schedulingUserId,omitempty"`
	ShiftName           *string `json:"shiftName,omitempty" xml:"shiftName,omitempty"`
	SkipOneDay          *bool   `json:"skipOneDay,omitempty" xml:"skipOneDay,omitempty"`
}

func (s UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingTemplateFineShifts) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingTemplateFineShifts) GoString() string {
	return s.String()
}

func (s *UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingTemplateFineShifts) SetSchedulingEndTime(v string) *UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingTemplateFineShifts {
	s.SchedulingEndTime = &v
	return s
}

func (s *UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingTemplateFineShifts) SetSchedulingOrder(v int32) *UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingTemplateFineShifts {
	s.SchedulingOrder = &v
	return s
}

func (s *UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingTemplateFineShifts) SetSchedulingStartTime(v string) *UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingTemplateFineShifts {
	s.SchedulingStartTime = &v
	return s
}

func (s *UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingTemplateFineShifts) SetSchedulingUserId(v int64) *UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingTemplateFineShifts {
	s.SchedulingUserId = &v
	return s
}

func (s *UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingTemplateFineShifts) SetShiftName(v string) *UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingTemplateFineShifts {
	s.ShiftName = &v
	return s
}

func (s *UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingTemplateFineShifts) SetSkipOneDay(v bool) *UpdateServiceGroupSchedulingRequestFineSchedulingSchedulingTemplateFineShifts {
	s.SkipOneDay = &v
	return s
}

type UpdateServiceGroupSchedulingResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateServiceGroupSchedulingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceGroupSchedulingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceGroupSchedulingResponseBody) SetRequestId(v string) *UpdateServiceGroupSchedulingResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServiceGroupSchedulingResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateServiceGroupSchedulingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateServiceGroupSchedulingResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceGroupSchedulingResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceGroupSchedulingResponse) SetHeaders(v map[string]*string) *UpdateServiceGroupSchedulingResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceGroupSchedulingResponse) SetStatusCode(v int32) *UpdateServiceGroupSchedulingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceGroupSchedulingResponse) SetBody(v *UpdateServiceGroupSchedulingResponseBody) *UpdateServiceGroupSchedulingResponse {
	s.Body = v
	return s
}

type UpdateServiceGroupSpecialDaySchedulingRequest struct {
	ClientToken           *string                                                               `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	SchedulingDate        *string                                                               `json:"schedulingDate,omitempty" xml:"schedulingDate,omitempty"`
	SchedulingSpecialDays []*UpdateServiceGroupSpecialDaySchedulingRequestSchedulingSpecialDays `json:"schedulingSpecialDays,omitempty" xml:"schedulingSpecialDays,omitempty" type:"Repeated"`
	ServiceGroupId        *int64                                                                `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
}

func (s UpdateServiceGroupSpecialDaySchedulingRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceGroupSpecialDaySchedulingRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceGroupSpecialDaySchedulingRequest) SetClientToken(v string) *UpdateServiceGroupSpecialDaySchedulingRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServiceGroupSpecialDaySchedulingRequest) SetSchedulingDate(v string) *UpdateServiceGroupSpecialDaySchedulingRequest {
	s.SchedulingDate = &v
	return s
}

func (s *UpdateServiceGroupSpecialDaySchedulingRequest) SetSchedulingSpecialDays(v []*UpdateServiceGroupSpecialDaySchedulingRequestSchedulingSpecialDays) *UpdateServiceGroupSpecialDaySchedulingRequest {
	s.SchedulingSpecialDays = v
	return s
}

func (s *UpdateServiceGroupSpecialDaySchedulingRequest) SetServiceGroupId(v int64) *UpdateServiceGroupSpecialDaySchedulingRequest {
	s.ServiceGroupId = &v
	return s
}

type UpdateServiceGroupSpecialDaySchedulingRequestSchedulingSpecialDays struct {
	SchedulingEndTime   *string `json:"schedulingEndTime,omitempty" xml:"schedulingEndTime,omitempty"`
	SchedulingOrder     *int32  `json:"schedulingOrder,omitempty" xml:"schedulingOrder,omitempty"`
	SchedulingStartTime *string `json:"schedulingStartTime,omitempty" xml:"schedulingStartTime,omitempty"`
	SchedulingUserId    *int64  `json:"schedulingUserId,omitempty" xml:"schedulingUserId,omitempty"`
}

func (s UpdateServiceGroupSpecialDaySchedulingRequestSchedulingSpecialDays) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceGroupSpecialDaySchedulingRequestSchedulingSpecialDays) GoString() string {
	return s.String()
}

func (s *UpdateServiceGroupSpecialDaySchedulingRequestSchedulingSpecialDays) SetSchedulingEndTime(v string) *UpdateServiceGroupSpecialDaySchedulingRequestSchedulingSpecialDays {
	s.SchedulingEndTime = &v
	return s
}

func (s *UpdateServiceGroupSpecialDaySchedulingRequestSchedulingSpecialDays) SetSchedulingOrder(v int32) *UpdateServiceGroupSpecialDaySchedulingRequestSchedulingSpecialDays {
	s.SchedulingOrder = &v
	return s
}

func (s *UpdateServiceGroupSpecialDaySchedulingRequestSchedulingSpecialDays) SetSchedulingStartTime(v string) *UpdateServiceGroupSpecialDaySchedulingRequestSchedulingSpecialDays {
	s.SchedulingStartTime = &v
	return s
}

func (s *UpdateServiceGroupSpecialDaySchedulingRequestSchedulingSpecialDays) SetSchedulingUserId(v int64) *UpdateServiceGroupSpecialDaySchedulingRequestSchedulingSpecialDays {
	s.SchedulingUserId = &v
	return s
}

type UpdateServiceGroupSpecialDaySchedulingResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateServiceGroupSpecialDaySchedulingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceGroupSpecialDaySchedulingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceGroupSpecialDaySchedulingResponseBody) SetRequestId(v string) *UpdateServiceGroupSpecialDaySchedulingResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServiceGroupSpecialDaySchedulingResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateServiceGroupSpecialDaySchedulingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateServiceGroupSpecialDaySchedulingResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceGroupSpecialDaySchedulingResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceGroupSpecialDaySchedulingResponse) SetHeaders(v map[string]*string) *UpdateServiceGroupSpecialDaySchedulingResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceGroupSpecialDaySchedulingResponse) SetStatusCode(v int32) *UpdateServiceGroupSpecialDaySchedulingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceGroupSpecialDaySchedulingResponse) SetBody(v *UpdateServiceGroupSpecialDaySchedulingResponseBody) *UpdateServiceGroupSpecialDaySchedulingResponse {
	s.Body = v
	return s
}

type UpdateSubscriptionRequest struct {
	EndTime            *string                                        `json:"endTime,omitempty" xml:"endTime,omitempty"`
	ExpiredType        *string                                        `json:"expiredType,omitempty" xml:"expiredType,omitempty"`
	NotifyObjectList   []*UpdateSubscriptionRequestNotifyObjectList   `json:"notifyObjectList,omitempty" xml:"notifyObjectList,omitempty" type:"Repeated"`
	NotifyObjectType   *string                                        `json:"notifyObjectType,omitempty" xml:"notifyObjectType,omitempty"`
	NotifyStrategyList []*UpdateSubscriptionRequestNotifyStrategyList `json:"notifyStrategyList,omitempty" xml:"notifyStrategyList,omitempty" type:"Repeated"`
	Period             *string                                        `json:"period,omitempty" xml:"period,omitempty"`
	Scope              *string                                        `json:"scope,omitempty" xml:"scope,omitempty"`
	ScopeObjectList    []*UpdateSubscriptionRequestScopeObjectList    `json:"scopeObjectList,omitempty" xml:"scopeObjectList,omitempty" type:"Repeated"`
	StartTime          *string                                        `json:"startTime,omitempty" xml:"startTime,omitempty"`
	SubscriptionId     *int64                                         `json:"subscriptionId,omitempty" xml:"subscriptionId,omitempty"`
	SubscriptionTitle  *string                                        `json:"subscriptionTitle,omitempty" xml:"subscriptionTitle,omitempty"`
}

func (s UpdateSubscriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateSubscriptionRequest) SetEndTime(v string) *UpdateSubscriptionRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetExpiredType(v string) *UpdateSubscriptionRequest {
	s.ExpiredType = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetNotifyObjectList(v []*UpdateSubscriptionRequestNotifyObjectList) *UpdateSubscriptionRequest {
	s.NotifyObjectList = v
	return s
}

func (s *UpdateSubscriptionRequest) SetNotifyObjectType(v string) *UpdateSubscriptionRequest {
	s.NotifyObjectType = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetNotifyStrategyList(v []*UpdateSubscriptionRequestNotifyStrategyList) *UpdateSubscriptionRequest {
	s.NotifyStrategyList = v
	return s
}

func (s *UpdateSubscriptionRequest) SetPeriod(v string) *UpdateSubscriptionRequest {
	s.Period = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetScope(v string) *UpdateSubscriptionRequest {
	s.Scope = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetScopeObjectList(v []*UpdateSubscriptionRequestScopeObjectList) *UpdateSubscriptionRequest {
	s.ScopeObjectList = v
	return s
}

func (s *UpdateSubscriptionRequest) SetStartTime(v string) *UpdateSubscriptionRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetSubscriptionId(v int64) *UpdateSubscriptionRequest {
	s.SubscriptionId = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetSubscriptionTitle(v string) *UpdateSubscriptionRequest {
	s.SubscriptionTitle = &v
	return s
}

type UpdateSubscriptionRequestNotifyObjectList struct {
	Id             *int64 `json:"id,omitempty" xml:"id,omitempty"`
	NotifyObjectId *int64 `json:"notifyObjectId,omitempty" xml:"notifyObjectId,omitempty"`
}

func (s UpdateSubscriptionRequestNotifyObjectList) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubscriptionRequestNotifyObjectList) GoString() string {
	return s.String()
}

func (s *UpdateSubscriptionRequestNotifyObjectList) SetId(v int64) *UpdateSubscriptionRequestNotifyObjectList {
	s.Id = &v
	return s
}

func (s *UpdateSubscriptionRequestNotifyObjectList) SetNotifyObjectId(v int64) *UpdateSubscriptionRequestNotifyObjectList {
	s.NotifyObjectId = &v
	return s
}

type UpdateSubscriptionRequestNotifyStrategyList struct {
	InstanceType *int64                                                   `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	Strategies   []*UpdateSubscriptionRequestNotifyStrategyListStrategies `json:"strategies,omitempty" xml:"strategies,omitempty" type:"Repeated"`
}

func (s UpdateSubscriptionRequestNotifyStrategyList) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubscriptionRequestNotifyStrategyList) GoString() string {
	return s.String()
}

func (s *UpdateSubscriptionRequestNotifyStrategyList) SetInstanceType(v int64) *UpdateSubscriptionRequestNotifyStrategyList {
	s.InstanceType = &v
	return s
}

func (s *UpdateSubscriptionRequestNotifyStrategyList) SetStrategies(v []*UpdateSubscriptionRequestNotifyStrategyListStrategies) *UpdateSubscriptionRequestNotifyStrategyList {
	s.Strategies = v
	return s
}

type UpdateSubscriptionRequestNotifyStrategyListStrategies struct {
	Channels      *string                                                             `json:"channels,omitempty" xml:"channels,omitempty"`
	Conditions    []*UpdateSubscriptionRequestNotifyStrategyListStrategiesConditions  `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	Id            *string                                                             `json:"id,omitempty" xml:"id,omitempty"`
	PeriodChannel *UpdateSubscriptionRequestNotifyStrategyListStrategiesPeriodChannel `json:"periodChannel,omitempty" xml:"periodChannel,omitempty" type:"Struct"`
}

func (s UpdateSubscriptionRequestNotifyStrategyListStrategies) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubscriptionRequestNotifyStrategyListStrategies) GoString() string {
	return s.String()
}

func (s *UpdateSubscriptionRequestNotifyStrategyListStrategies) SetChannels(v string) *UpdateSubscriptionRequestNotifyStrategyListStrategies {
	s.Channels = &v
	return s
}

func (s *UpdateSubscriptionRequestNotifyStrategyListStrategies) SetConditions(v []*UpdateSubscriptionRequestNotifyStrategyListStrategiesConditions) *UpdateSubscriptionRequestNotifyStrategyListStrategies {
	s.Conditions = v
	return s
}

func (s *UpdateSubscriptionRequestNotifyStrategyListStrategies) SetId(v string) *UpdateSubscriptionRequestNotifyStrategyListStrategies {
	s.Id = &v
	return s
}

func (s *UpdateSubscriptionRequestNotifyStrategyListStrategies) SetPeriodChannel(v *UpdateSubscriptionRequestNotifyStrategyListStrategiesPeriodChannel) *UpdateSubscriptionRequestNotifyStrategyListStrategies {
	s.PeriodChannel = v
	return s
}

type UpdateSubscriptionRequestNotifyStrategyListStrategiesConditions struct {
	Action            *string `json:"action,omitempty" xml:"action,omitempty"`
	Effection         *string `json:"effection,omitempty" xml:"effection,omitempty"`
	Level             *string `json:"level,omitempty" xml:"level,omitempty"`
	ProblemNotifyType *string `json:"problemNotifyType,omitempty" xml:"problemNotifyType,omitempty"`
}

func (s UpdateSubscriptionRequestNotifyStrategyListStrategiesConditions) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubscriptionRequestNotifyStrategyListStrategiesConditions) GoString() string {
	return s.String()
}

func (s *UpdateSubscriptionRequestNotifyStrategyListStrategiesConditions) SetAction(v string) *UpdateSubscriptionRequestNotifyStrategyListStrategiesConditions {
	s.Action = &v
	return s
}

func (s *UpdateSubscriptionRequestNotifyStrategyListStrategiesConditions) SetEffection(v string) *UpdateSubscriptionRequestNotifyStrategyListStrategiesConditions {
	s.Effection = &v
	return s
}

func (s *UpdateSubscriptionRequestNotifyStrategyListStrategiesConditions) SetLevel(v string) *UpdateSubscriptionRequestNotifyStrategyListStrategiesConditions {
	s.Level = &v
	return s
}

func (s *UpdateSubscriptionRequestNotifyStrategyListStrategiesConditions) SetProblemNotifyType(v string) *UpdateSubscriptionRequestNotifyStrategyListStrategiesConditions {
	s.ProblemNotifyType = &v
	return s
}

type UpdateSubscriptionRequestNotifyStrategyListStrategiesPeriodChannel struct {
	NonWorkday *string `json:"nonWorkday,omitempty" xml:"nonWorkday,omitempty"`
	Workday    *string `json:"workday,omitempty" xml:"workday,omitempty"`
}

func (s UpdateSubscriptionRequestNotifyStrategyListStrategiesPeriodChannel) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubscriptionRequestNotifyStrategyListStrategiesPeriodChannel) GoString() string {
	return s.String()
}

func (s *UpdateSubscriptionRequestNotifyStrategyListStrategiesPeriodChannel) SetNonWorkday(v string) *UpdateSubscriptionRequestNotifyStrategyListStrategiesPeriodChannel {
	s.NonWorkday = &v
	return s
}

func (s *UpdateSubscriptionRequestNotifyStrategyListStrategiesPeriodChannel) SetWorkday(v string) *UpdateSubscriptionRequestNotifyStrategyListStrategiesPeriodChannel {
	s.Workday = &v
	return s
}

type UpdateSubscriptionRequestScopeObjectList struct {
	Id            *int64 `json:"id,omitempty" xml:"id,omitempty"`
	ScopeObjectId *int64 `json:"scopeObjectId,omitempty" xml:"scopeObjectId,omitempty"`
}

func (s UpdateSubscriptionRequestScopeObjectList) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubscriptionRequestScopeObjectList) GoString() string {
	return s.String()
}

func (s *UpdateSubscriptionRequestScopeObjectList) SetId(v int64) *UpdateSubscriptionRequestScopeObjectList {
	s.Id = &v
	return s
}

func (s *UpdateSubscriptionRequestScopeObjectList) SetScopeObjectId(v int64) *UpdateSubscriptionRequestScopeObjectList {
	s.ScopeObjectId = &v
	return s
}

type UpdateSubscriptionResponseBody struct {
	// id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateSubscriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSubscriptionResponseBody) SetRequestId(v string) *UpdateSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSubscriptionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSubscriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateSubscriptionResponse) SetHeaders(v map[string]*string) *UpdateSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateSubscriptionResponse) SetStatusCode(v int32) *UpdateSubscriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSubscriptionResponse) SetBody(v *UpdateSubscriptionResponseBody) *UpdateSubscriptionResponse {
	s.Body = v
	return s
}

type UpdateUserRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	Email       *string `json:"email,omitempty" xml:"email,omitempty"`
	Phone       *string `json:"phone,omitempty" xml:"phone,omitempty"`
	RamId       *int64  `json:"ramId,omitempty" xml:"ramId,omitempty"`
	UserId      *int64  `json:"userId,omitempty" xml:"userId,omitempty"`
	Username    *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s UpdateUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserRequest) SetClientToken(v string) *UpdateUserRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateUserRequest) SetEmail(v string) *UpdateUserRequest {
	s.Email = &v
	return s
}

func (s *UpdateUserRequest) SetPhone(v string) *UpdateUserRequest {
	s.Phone = &v
	return s
}

func (s *UpdateUserRequest) SetRamId(v int64) *UpdateUserRequest {
	s.RamId = &v
	return s
}

func (s *UpdateUserRequest) SetUserId(v int64) *UpdateUserRequest {
	s.UserId = &v
	return s
}

func (s *UpdateUserRequest) SetUsername(v string) *UpdateUserRequest {
	s.Username = &v
	return s
}

type UpdateUserResponseBody struct {
	// id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserResponseBody) SetRequestId(v string) *UpdateUserResponseBody {
	s.RequestId = &v
	return s
}

type UpdateUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserResponse) SetHeaders(v map[string]*string) *UpdateUserResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserResponse) SetStatusCode(v int32) *UpdateUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserResponse) SetBody(v *UpdateUserResponseBody) *UpdateUserResponse {
	s.Body = v
	return s
}

type UpdateUserGuideStatusRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	GuideAction *string `json:"guideAction,omitempty" xml:"guideAction,omitempty"`
}

func (s UpdateUserGuideStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserGuideStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserGuideStatusRequest) SetClientToken(v string) *UpdateUserGuideStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateUserGuideStatusRequest) SetGuideAction(v string) *UpdateUserGuideStatusRequest {
	s.GuideAction = &v
	return s
}

type UpdateUserGuideStatusResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateUserGuideStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserGuideStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserGuideStatusResponseBody) SetRequestId(v string) *UpdateUserGuideStatusResponseBody {
	s.RequestId = &v
	return s
}

type UpdateUserGuideStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateUserGuideStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateUserGuideStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserGuideStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserGuideStatusResponse) SetHeaders(v map[string]*string) *UpdateUserGuideStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserGuideStatusResponse) SetStatusCode(v int32) *UpdateUserGuideStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserGuideStatusResponse) SetBody(v *UpdateUserGuideStatusResponseBody) *UpdateUserGuideStatusResponse {
	s.Body = v
	return s
}

type VerifyRouteRuleRequest struct {
	RouteRuleId      *int64                                    `json:"routeRuleId,omitempty" xml:"routeRuleId,omitempty"`
	TestSourceEvents []*VerifyRouteRuleRequestTestSourceEvents `json:"testSourceEvents,omitempty" xml:"testSourceEvents,omitempty" type:"Repeated"`
}

func (s VerifyRouteRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyRouteRuleRequest) GoString() string {
	return s.String()
}

func (s *VerifyRouteRuleRequest) SetRouteRuleId(v int64) *VerifyRouteRuleRequest {
	s.RouteRuleId = &v
	return s
}

func (s *VerifyRouteRuleRequest) SetTestSourceEvents(v []*VerifyRouteRuleRequestTestSourceEvents) *VerifyRouteRuleRequest {
	s.TestSourceEvents = v
	return s
}

type VerifyRouteRuleRequestTestSourceEvents struct {
	EventJson         *string `json:"eventJson,omitempty" xml:"eventJson,omitempty"`
	EventTime         *string `json:"eventTime,omitempty" xml:"eventTime,omitempty"`
	MonitorSourceId   *int64  `json:"monitorSourceId,omitempty" xml:"monitorSourceId,omitempty"`
	MonitorSourceName *string `json:"monitorSourceName,omitempty" xml:"monitorSourceName,omitempty"`
}

func (s VerifyRouteRuleRequestTestSourceEvents) String() string {
	return tea.Prettify(s)
}

func (s VerifyRouteRuleRequestTestSourceEvents) GoString() string {
	return s.String()
}

func (s *VerifyRouteRuleRequestTestSourceEvents) SetEventJson(v string) *VerifyRouteRuleRequestTestSourceEvents {
	s.EventJson = &v
	return s
}

func (s *VerifyRouteRuleRequestTestSourceEvents) SetEventTime(v string) *VerifyRouteRuleRequestTestSourceEvents {
	s.EventTime = &v
	return s
}

func (s *VerifyRouteRuleRequestTestSourceEvents) SetMonitorSourceId(v int64) *VerifyRouteRuleRequestTestSourceEvents {
	s.MonitorSourceId = &v
	return s
}

func (s *VerifyRouteRuleRequestTestSourceEvents) SetMonitorSourceName(v string) *VerifyRouteRuleRequestTestSourceEvents {
	s.MonitorSourceName = &v
	return s
}

type VerifyRouteRuleResponseBody struct {
	Data *VerifyRouteRuleResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s VerifyRouteRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyRouteRuleResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyRouteRuleResponseBody) SetData(v *VerifyRouteRuleResponseBodyData) *VerifyRouteRuleResponseBody {
	s.Data = v
	return s
}

func (s *VerifyRouteRuleResponseBody) SetRequestId(v string) *VerifyRouteRuleResponseBody {
	s.RequestId = &v
	return s
}

type VerifyRouteRuleResponseBodyData struct {
	EscalationPlans         []*VerifyRouteRuleResponseBodyDataEscalationPlans         `json:"escalationPlans,omitempty" xml:"escalationPlans,omitempty" type:"Repeated"`
	IsValidRule             *bool                                                     `json:"isValidRule,omitempty" xml:"isValidRule,omitempty"`
	MonitorSourceIds        []*int64                                                  `json:"monitorSourceIds,omitempty" xml:"monitorSourceIds,omitempty" type:"Repeated"`
	NotifySubscriptionNames []*VerifyRouteRuleResponseBodyDataNotifySubscriptionNames `json:"notifySubscriptionNames,omitempty" xml:"notifySubscriptionNames,omitempty" type:"Repeated"`
	RouteRuleFailReason     []*string                                                 `json:"routeRuleFailReason,omitempty" xml:"routeRuleFailReason,omitempty" type:"Repeated"`
	RouteType               *string                                                   `json:"routeType,omitempty" xml:"routeType,omitempty"`
}

func (s VerifyRouteRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s VerifyRouteRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *VerifyRouteRuleResponseBodyData) SetEscalationPlans(v []*VerifyRouteRuleResponseBodyDataEscalationPlans) *VerifyRouteRuleResponseBodyData {
	s.EscalationPlans = v
	return s
}

func (s *VerifyRouteRuleResponseBodyData) SetIsValidRule(v bool) *VerifyRouteRuleResponseBodyData {
	s.IsValidRule = &v
	return s
}

func (s *VerifyRouteRuleResponseBodyData) SetMonitorSourceIds(v []*int64) *VerifyRouteRuleResponseBodyData {
	s.MonitorSourceIds = v
	return s
}

func (s *VerifyRouteRuleResponseBodyData) SetNotifySubscriptionNames(v []*VerifyRouteRuleResponseBodyDataNotifySubscriptionNames) *VerifyRouteRuleResponseBodyData {
	s.NotifySubscriptionNames = v
	return s
}

func (s *VerifyRouteRuleResponseBodyData) SetRouteRuleFailReason(v []*string) *VerifyRouteRuleResponseBodyData {
	s.RouteRuleFailReason = v
	return s
}

func (s *VerifyRouteRuleResponseBodyData) SetRouteType(v string) *VerifyRouteRuleResponseBodyData {
	s.RouteType = &v
	return s
}

type VerifyRouteRuleResponseBodyDataEscalationPlans struct {
	EscalationPlanId   *int64  `json:"escalationPlanId,omitempty" xml:"escalationPlanId,omitempty"`
	EscalationPlanName *string `json:"escalationPlanName,omitempty" xml:"escalationPlanName,omitempty"`
}

func (s VerifyRouteRuleResponseBodyDataEscalationPlans) String() string {
	return tea.Prettify(s)
}

func (s VerifyRouteRuleResponseBodyDataEscalationPlans) GoString() string {
	return s.String()
}

func (s *VerifyRouteRuleResponseBodyDataEscalationPlans) SetEscalationPlanId(v int64) *VerifyRouteRuleResponseBodyDataEscalationPlans {
	s.EscalationPlanId = &v
	return s
}

func (s *VerifyRouteRuleResponseBodyDataEscalationPlans) SetEscalationPlanName(v string) *VerifyRouteRuleResponseBodyDataEscalationPlans {
	s.EscalationPlanName = &v
	return s
}

type VerifyRouteRuleResponseBodyDataNotifySubscriptionNames struct {
	SubscriptionId *int64  `json:"subscriptionId,omitempty" xml:"subscriptionId,omitempty"`
	Title          *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s VerifyRouteRuleResponseBodyDataNotifySubscriptionNames) String() string {
	return tea.Prettify(s)
}

func (s VerifyRouteRuleResponseBodyDataNotifySubscriptionNames) GoString() string {
	return s.String()
}

func (s *VerifyRouteRuleResponseBodyDataNotifySubscriptionNames) SetSubscriptionId(v int64) *VerifyRouteRuleResponseBodyDataNotifySubscriptionNames {
	s.SubscriptionId = &v
	return s
}

func (s *VerifyRouteRuleResponseBodyDataNotifySubscriptionNames) SetTitle(v string) *VerifyRouteRuleResponseBodyDataNotifySubscriptionNames {
	s.Title = &v
	return s
}

type VerifyRouteRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *VerifyRouteRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifyRouteRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyRouteRuleResponse) GoString() string {
	return s.String()
}

func (s *VerifyRouteRuleResponse) SetHeaders(v map[string]*string) *VerifyRouteRuleResponse {
	s.Headers = v
	return s
}

func (s *VerifyRouteRuleResponse) SetStatusCode(v int32) *VerifyRouteRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyRouteRuleResponse) SetBody(v *VerifyRouteRuleResponseBody) *VerifyRouteRuleResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("gemp"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddProblemServiceGroupWithOptions(request *AddProblemServiceGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddProblemServiceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProblemId)) {
		body["problemId"] = request.ProblemId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceGroupIds)) {
		body["serviceGroupIds"] = request.ServiceGroupIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddProblemServiceGroup"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/problem/addServiceGroup"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddProblemServiceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddProblemServiceGroup(request *AddProblemServiceGroupRequest) (_result *AddProblemServiceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddProblemServiceGroupResponse{}
	_body, _err := client.AddProblemServiceGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BillingStatisticsWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *BillingStatisticsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("BillingStatistics"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/charging/details"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BillingStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BillingStatistics() (_result *BillingStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BillingStatisticsResponse{}
	_body, _err := client.BillingStatisticsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelProblemWithOptions(request *CancelProblemRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CancelProblemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CancelReason)) {
		body["cancelReason"] = request.CancelReason
	}

	if !tea.BoolValue(util.IsUnset(request.CancelReasonDescription)) {
		body["cancelReasonDescription"] = request.CancelReasonDescription
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemId)) {
		body["problemId"] = request.ProblemId
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemNotifyType)) {
		body["problemNotifyType"] = request.ProblemNotifyType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelProblem"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/problem/cancel"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelProblemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelProblem(request *CancelProblemRequest) (_result *CancelProblemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelProblemResponse{}
	_body, _err := client.CancelProblemWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckWebhookWithOptions(request *CheckWebhookRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CheckWebhookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Webhook)) {
		body["webhook"] = request.Webhook
	}

	if !tea.BoolValue(util.IsUnset(request.WebhookType)) {
		body["webhookType"] = request.WebhookType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckWebhook"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/services/check/webhook"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckWebhookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckWebhook(request *CheckWebhookRequest) (_result *CheckWebhookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckWebhookResponse{}
	_body, _err := client.CheckWebhookWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfirmIntegrationConfigWithOptions(request *ConfirmIntegrationConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ConfirmIntegrationConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.IntegrationConfigId)) {
		body["integrationConfigId"] = request.IntegrationConfigId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfirmIntegrationConfig"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/integrationConfig/confirm"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfirmIntegrationConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfirmIntegrationConfig(request *ConfirmIntegrationConfigRequest) (_result *ConfirmIntegrationConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ConfirmIntegrationConfigResponse{}
	_body, _err := client.ConfirmIntegrationConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEscalationPlanWithOptions(request *CreateEscalationPlanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateEscalationPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EscalationPlanDescription)) {
		body["escalationPlanDescription"] = request.EscalationPlanDescription
	}

	if !tea.BoolValue(util.IsUnset(request.EscalationPlanName)) {
		body["escalationPlanName"] = request.EscalationPlanName
	}

	if !tea.BoolValue(util.IsUnset(request.EscalationPlanRules)) {
		body["escalationPlanRules"] = request.EscalationPlanRules
	}

	if !tea.BoolValue(util.IsUnset(request.EscalationPlanScopeObjects)) {
		body["escalationPlanScopeObjects"] = request.EscalationPlanScopeObjects
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEscalationPlan"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/escalationPlan/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateEscalationPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEscalationPlan(request *CreateEscalationPlanRequest) (_result *CreateEscalationPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateEscalationPlanResponse{}
	_body, _err := client.CreateEscalationPlanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateIncidentWithOptions(request *CreateIncidentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateIncidentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssignUserId)) {
		body["assignUserId"] = request.AssignUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Channels)) {
		body["channels"] = request.Channels
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Effect)) {
		body["effect"] = request.Effect
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentDescription)) {
		body["incidentDescription"] = request.IncidentDescription
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentLevel)) {
		body["incidentLevel"] = request.IncidentLevel
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentTitle)) {
		body["incidentTitle"] = request.IncidentTitle
	}

	if !tea.BoolValue(util.IsUnset(request.RelatedServiceId)) {
		body["relatedServiceId"] = request.RelatedServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceGroupId)) {
		body["serviceGroupId"] = request.ServiceGroupId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIncident"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/incident/manualSave"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateIncidentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateIncident(request *CreateIncidentRequest) (_result *CreateIncidentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateIncidentResponse{}
	_body, _err := client.CreateIncidentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateIncidentSubtotalWithOptions(request *CreateIncidentSubtotalRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateIncidentSubtotalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentId)) {
		body["incidentId"] = request.IncidentId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIncidentSubtotal"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/incident/save/subtotal"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateIncidentSubtotalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateIncidentSubtotal(request *CreateIncidentSubtotalRequest) (_result *CreateIncidentSubtotalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateIncidentSubtotalResponse{}
	_body, _err := client.CreateIncidentSubtotalWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateIntegrationConfigWithOptions(request *CreateIntegrationConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateIntegrationConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.MonitorSourceId)) {
		body["monitorSourceId"] = request.MonitorSourceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIntegrationConfig"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/integrationConfig/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateIntegrationConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateIntegrationConfig(request *CreateIntegrationConfigRequest) (_result *CreateIntegrationConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateIntegrationConfigResponse{}
	_body, _err := client.CreateIntegrationConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateProblemWithOptions(request *CreateProblemRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateProblemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AffectServiceIds)) {
		body["affectServiceIds"] = request.AffectServiceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DiscoverTime)) {
		body["discoverTime"] = request.DiscoverTime
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentId)) {
		body["incidentId"] = request.IncidentId
	}

	if !tea.BoolValue(util.IsUnset(request.MainHandlerId)) {
		body["mainHandlerId"] = request.MainHandlerId
	}

	if !tea.BoolValue(util.IsUnset(request.PreliminaryReason)) {
		body["preliminaryReason"] = request.PreliminaryReason
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemLevel)) {
		body["problemLevel"] = request.ProblemLevel
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemName)) {
		body["problemName"] = request.ProblemName
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemNotifyType)) {
		body["problemNotifyType"] = request.ProblemNotifyType
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemStatus)) {
		body["problemStatus"] = request.ProblemStatus
	}

	if !tea.BoolValue(util.IsUnset(request.ProgressSummary)) {
		body["progressSummary"] = request.ProgressSummary
	}

	if !tea.BoolValue(util.IsUnset(request.ProgressSummaryRichTextId)) {
		body["progressSummaryRichTextId"] = request.ProgressSummaryRichTextId
	}

	if !tea.BoolValue(util.IsUnset(request.RecoveryTime)) {
		body["recoveryTime"] = request.RecoveryTime
	}

	if !tea.BoolValue(util.IsUnset(request.RelatedServiceId)) {
		body["relatedServiceId"] = request.RelatedServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceGroupIds)) {
		body["serviceGroupIds"] = request.ServiceGroupIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProblem"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/problem/upgrade"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProblemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProblem(request *CreateProblemRequest) (_result *CreateProblemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateProblemResponse{}
	_body, _err := client.CreateProblemWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateProblemEffectionServiceWithOptions(request *CreateProblemEffectionServiceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateProblemEffectionServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Level)) {
		body["level"] = request.Level
	}

	if !tea.BoolValue(util.IsUnset(request.PictureUrl)) {
		body["pictureUrl"] = request.PictureUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemId)) {
		body["problemId"] = request.ProblemId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		body["serviceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProblemEffectionService"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/problem/process/effectionService/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProblemEffectionServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProblemEffectionService(request *CreateProblemEffectionServiceRequest) (_result *CreateProblemEffectionServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateProblemEffectionServiceResponse{}
	_body, _err := client.CreateProblemEffectionServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateProblemMeasureWithOptions(request *CreateProblemMeasureRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateProblemMeasureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CheckStandard)) {
		body["checkStandard"] = request.CheckStandard
	}

	if !tea.BoolValue(util.IsUnset(request.CheckUserId)) {
		body["checkUserId"] = request.CheckUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.DirectorId)) {
		body["directorId"] = request.DirectorId
	}

	if !tea.BoolValue(util.IsUnset(request.PlanFinishTime)) {
		body["planFinishTime"] = request.PlanFinishTime
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemId)) {
		body["problemId"] = request.ProblemId
	}

	if !tea.BoolValue(util.IsUnset(request.StalkerId)) {
		body["stalkerId"] = request.StalkerId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProblemMeasure"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/problem/improvement/measure/save"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProblemMeasureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProblemMeasure(request *CreateProblemMeasureRequest) (_result *CreateProblemMeasureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateProblemMeasureResponse{}
	_body, _err := client.CreateProblemMeasureWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateProblemSubtotalWithOptions(request *CreateProblemSubtotalRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateProblemSubtotalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemId)) {
		body["problemId"] = request.ProblemId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProblemSubtotal"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/problem/save/subtotal"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProblemSubtotalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProblemSubtotal(request *CreateProblemSubtotalRequest) (_result *CreateProblemSubtotalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateProblemSubtotalResponse{}
	_body, _err := client.CreateProblemSubtotalWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateProblemTimelineWithOptions(request *CreateProblemTimelineRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateProblemTimelineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.KeyNode)) {
		body["keyNode"] = request.KeyNode
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemId)) {
		body["problemId"] = request.ProblemId
	}

	if !tea.BoolValue(util.IsUnset(request.Time)) {
		body["time"] = request.Time
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProblemTimeline"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/problem/process/timeline/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProblemTimelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProblemTimeline(request *CreateProblemTimelineRequest) (_result *CreateProblemTimelineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateProblemTimelineResponse{}
	_body, _err := client.CreateProblemTimelineWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateProblemTimelinesWithOptions(request *CreateProblemTimelinesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateProblemTimelinesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemId)) {
		body["problemId"] = request.ProblemId
	}

	if !tea.BoolValue(util.IsUnset(request.TimelineNodes)) {
		body["timelineNodes"] = request.TimelineNodes
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProblemTimelines"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/problem/process/timeline/batchCreate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProblemTimelinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProblemTimelines(request *CreateProblemTimelinesRequest) (_result *CreateProblemTimelinesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateProblemTimelinesResponse{}
	_body, _err := client.CreateProblemTimelinesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRichTextWithOptions(request *CreateRichTextRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateRichTextResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		body["instanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.RichText)) {
		body["richText"] = request.RichText
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRichText"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/rich/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRichTextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRichText(request *CreateRichTextRequest) (_result *CreateRichTextResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateRichTextResponse{}
	_body, _err := client.CreateRichTextWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRouteRuleWithOptions(request *CreateRouteRuleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateRouteRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssignObjectId)) {
		body["assignObjectId"] = request.AssignObjectId
	}

	if !tea.BoolValue(util.IsUnset(request.AssignObjectType)) {
		body["assignObjectType"] = request.AssignObjectType
	}

	if !tea.BoolValue(util.IsUnset(request.ChildRuleRelation)) {
		body["childRuleRelation"] = request.ChildRuleRelation
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CoverageProblemLevels)) {
		body["coverageProblemLevels"] = request.CoverageProblemLevels
	}

	if !tea.BoolValue(util.IsUnset(request.Effection)) {
		body["effection"] = request.Effection
	}

	if !tea.BoolValue(util.IsUnset(request.EnableStatus)) {
		body["enableStatus"] = request.EnableStatus
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentLevel)) {
		body["incidentLevel"] = request.IncidentLevel
	}

	if !tea.BoolValue(util.IsUnset(request.MatchCount)) {
		body["matchCount"] = request.MatchCount
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyChannels)) {
		body["notifyChannels"] = request.NotifyChannels
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemEffectionServices)) {
		body["problemEffectionServices"] = request.ProblemEffectionServices
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemLevelGroup)) {
		body["problemLevelGroup"] = request.ProblemLevelGroup
	}

	if !tea.BoolValue(util.IsUnset(request.RelatedServiceId)) {
		body["relatedServiceId"] = request.RelatedServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.RouteChildRules)) {
		body["routeChildRules"] = request.RouteChildRules
	}

	if !tea.BoolValue(util.IsUnset(request.RouteType)) {
		body["routeType"] = request.RouteType
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		body["ruleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.TimeWindow)) {
		body["timeWindow"] = request.TimeWindow
	}

	if !tea.BoolValue(util.IsUnset(request.TimeWindowUnit)) {
		body["timeWindowUnit"] = request.TimeWindowUnit
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRouteRule"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/routeRule/save"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRouteRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRouteRule(request *CreateRouteRuleRequest) (_result *CreateRouteRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateRouteRuleResponse{}
	_body, _err := client.CreateRouteRuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateServiceWithOptions(request *CreateServiceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceDescription)) {
		body["serviceDescription"] = request.ServiceDescription
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		body["serviceName"] = request.ServiceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateService"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/services/save"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateService(request *CreateServiceRequest) (_result *CreateServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateServiceResponse{}
	_body, _err := client.CreateServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateServiceGroupWithOptions(request *CreateServiceGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateServiceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EnableWebhook)) {
		body["enableWebhook"] = request.EnableWebhook
	}

	if !tea.BoolValue(util.IsUnset(request.MonitorSourceTemplates)) {
		body["monitorSourceTemplates"] = request.MonitorSourceTemplates
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceGroupDescription)) {
		body["serviceGroupDescription"] = request.ServiceGroupDescription
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceGroupName)) {
		body["serviceGroupName"] = request.ServiceGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["userIds"] = request.UserIds
	}

	if !tea.BoolValue(util.IsUnset(request.WebhookLink)) {
		body["webhookLink"] = request.WebhookLink
	}

	if !tea.BoolValue(util.IsUnset(request.WebhookType)) {
		body["webhookType"] = request.WebhookType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServiceGroup"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/services/group/insert"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateServiceGroup(request *CreateServiceGroupRequest) (_result *CreateServiceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateServiceGroupResponse{}
	_body, _err := client.CreateServiceGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateServiceGroupSchedulingWithOptions(request *CreateServiceGroupSchedulingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateServiceGroupSchedulingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.FastScheduling)) {
		body["fastScheduling"] = request.FastScheduling
	}

	if !tea.BoolValue(util.IsUnset(request.FineScheduling)) {
		body["fineScheduling"] = request.FineScheduling
	}

	if !tea.BoolValue(util.IsUnset(request.SchedulingWay)) {
		body["schedulingWay"] = request.SchedulingWay
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceGroupId)) {
		body["serviceGroupId"] = request.ServiceGroupId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServiceGroupScheduling"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/services/group/scheduling/save"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceGroupSchedulingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateServiceGroupScheduling(request *CreateServiceGroupSchedulingRequest) (_result *CreateServiceGroupSchedulingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateServiceGroupSchedulingResponse{}
	_body, _err := client.CreateServiceGroupSchedulingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSubscriptionWithOptions(request *CreateSubscriptionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateSubscriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExpiredType)) {
		body["expiredType"] = request.ExpiredType
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyObjectList)) {
		body["notifyObjectList"] = request.NotifyObjectList
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyObjectType)) {
		body["notifyObjectType"] = request.NotifyObjectType
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyStrategyList)) {
		body["notifyStrategyList"] = request.NotifyStrategyList
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		body["period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		body["scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeObjectList)) {
		body["scopeObjectList"] = request.ScopeObjectList
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriptionTitle)) {
		body["subscriptionTitle"] = request.SubscriptionTitle
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSubscription"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/notify/subscription/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSubscriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSubscription(request *CreateSubscriptionRequest) (_result *CreateSubscriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSubscriptionResponse{}
	_body, _err := client.CreateSubscriptionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTenantApplicationWithOptions(request *CreateTenantApplicationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateTenantApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		body["channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTenantApplication"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/mobileApp/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTenantApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTenantApplication(request *CreateTenantApplicationRequest) (_result *CreateTenantApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTenantApplicationResponse{}
	_body, _err := client.CreateTenantApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUserWithOptions(request *CreateUserRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		body["email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		body["phone"] = request.Phone
	}

	if !tea.BoolValue(util.IsUnset(request.RamId)) {
		body["ramId"] = request.RamId
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		body["username"] = request.Username
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUser"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/user/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUser(request *CreateUserRequest) (_result *CreateUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateUserResponse{}
	_body, _err := client.CreateUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEscalationPlanWithOptions(request *DeleteEscalationPlanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteEscalationPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EscalationPlanId)) {
		body["escalationPlanId"] = request.EscalationPlanId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEscalationPlan"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/escalationPlan/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteEscalationPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEscalationPlan(request *DeleteEscalationPlanRequest) (_result *DeleteEscalationPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteEscalationPlanResponse{}
	_body, _err := client.DeleteEscalationPlanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteIncidentWithOptions(request *DeleteIncidentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteIncidentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentId)) {
		body["incidentId"] = request.IncidentId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIncident"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/incident/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIncidentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteIncident(request *DeleteIncidentRequest) (_result *DeleteIncidentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteIncidentResponse{}
	_body, _err := client.DeleteIncidentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteIntegrationConfigWithOptions(request *DeleteIntegrationConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteIntegrationConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.IntegrationConfigId)) {
		body["integrationConfigId"] = request.IntegrationConfigId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIntegrationConfig"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/integrationConfig/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIntegrationConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteIntegrationConfig(request *DeleteIntegrationConfigRequest) (_result *DeleteIntegrationConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteIntegrationConfigResponse{}
	_body, _err := client.DeleteIntegrationConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProblemWithOptions(request *DeleteProblemRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteProblemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemId)) {
		body["problemId"] = request.ProblemId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProblem"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/problem/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProblemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProblem(request *DeleteProblemRequest) (_result *DeleteProblemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteProblemResponse{}
	_body, _err := client.DeleteProblemWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProblemEffectionServiceWithOptions(request *DeleteProblemEffectionServiceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteProblemEffectionServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EffectionServiceId)) {
		body["effectionServiceId"] = request.EffectionServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemId)) {
		body["problemId"] = request.ProblemId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProblemEffectionService"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/problem/process/effectionService/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProblemEffectionServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProblemEffectionService(request *DeleteProblemEffectionServiceRequest) (_result *DeleteProblemEffectionServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteProblemEffectionServiceResponse{}
	_body, _err := client.DeleteProblemEffectionServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProblemMeasureWithOptions(request *DeleteProblemMeasureRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteProblemMeasureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.MeasureId)) {
		body["measureId"] = request.MeasureId
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemId)) {
		body["problemId"] = request.ProblemId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProblemMeasure"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/problem/improvement/measure/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProblemMeasureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProblemMeasure(request *DeleteProblemMeasureRequest) (_result *DeleteProblemMeasureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteProblemMeasureResponse{}
	_body, _err := client.DeleteProblemMeasureWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProblemTimelineWithOptions(request *DeleteProblemTimelineRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteProblemTimelineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemId)) {
		body["problemId"] = request.ProblemId
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemTimelineId)) {
		body["problemTimelineId"] = request.ProblemTimelineId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProblemTimeline"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/problem/process/timeline/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProblemTimelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProblemTimeline(request *DeleteProblemTimelineRequest) (_result *DeleteProblemTimelineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteProblemTimelineResponse{}
	_body, _err := client.DeleteProblemTimelineWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRouteRuleWithOptions(request *DeleteRouteRuleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteRouteRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RouteRuleId)) {
		body["routeRuleId"] = request.RouteRuleId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRouteRule"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/routeRule/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRouteRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRouteRule(request *DeleteRouteRuleRequest) (_result *DeleteRouteRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteRouteRuleResponse{}
	_body, _err := client.DeleteRouteRuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteServiceWithOptions(request *DeleteServiceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		body["serviceId"] = request.ServiceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteService"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/services/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteService(request *DeleteServiceRequest) (_result *DeleteServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteServiceResponse{}
	_body, _err := client.DeleteServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteServiceGroupWithOptions(request *DeleteServiceGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteServiceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceGroupId)) {
		body["serviceGroupId"] = request.ServiceGroupId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteServiceGroup"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/services/group/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServiceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteServiceGroup(request *DeleteServiceGroupRequest) (_result *DeleteServiceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteServiceGroupResponse{}
	_body, _err := client.DeleteServiceGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteServiceGroupSchedulingWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteServiceGroupSchedulingResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteServiceGroupScheduling"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/services/group/scheduling/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServiceGroupSchedulingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteServiceGroupScheduling() (_result *DeleteServiceGroupSchedulingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteServiceGroupSchedulingResponse{}
	_body, _err := client.DeleteServiceGroupSchedulingWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteServiceGroupUserWithOptions(request *DeleteServiceGroupUserRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteServiceGroupUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.NewUserId)) {
		body["newUserId"] = request.NewUserId
	}

	if !tea.BoolValue(util.IsUnset(request.OldUserId)) {
		body["oldUserId"] = request.OldUserId
	}

	if !tea.BoolValue(util.IsUnset(request.RemoveUser)) {
		body["removeUser"] = request.RemoveUser
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceGroupId)) {
		body["serviceGroupId"] = request.ServiceGroupId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteServiceGroupUser"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/services/group/deleteServiceGroupUser"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServiceGroupUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteServiceGroupUser(request *DeleteServiceGroupUserRequest) (_result *DeleteServiceGroupUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteServiceGroupUserResponse{}
	_body, _err := client.DeleteServiceGroupUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSubscriptionWithOptions(request *DeleteSubscriptionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteSubscriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubscriptionId)) {
		body["subscriptionId"] = request.SubscriptionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSubscription"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/notify/subscription/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSubscriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSubscription(request *DeleteSubscriptionRequest) (_result *DeleteSubscriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSubscriptionResponse{}
	_body, _err := client.DeleteSubscriptionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUserWithOptions(request *DeleteUserRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUser"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/user/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUser(request *DeleteUserRequest) (_result *DeleteUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteUserResponse{}
	_body, _err := client.DeleteUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeliverIncidentWithOptions(request *DeliverIncidentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeliverIncidentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssignUserId)) {
		body["assignUserId"] = request.AssignUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentId)) {
		body["incidentId"] = request.IncidentId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeliverIncident"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/incident/deliver"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeliverIncidentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeliverIncident(request *DeliverIncidentRequest) (_result *DeliverIncidentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeliverIncidentResponse{}
	_body, _err := client.DeliverIncidentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableEscalationPlanWithOptions(request *DisableEscalationPlanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DisableEscalationPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EscalationPlanId)) {
		body["escalationPlanId"] = request.EscalationPlanId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableEscalationPlan"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/escalationPlan/disable"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableEscalationPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableEscalationPlan(request *DisableEscalationPlanRequest) (_result *DisableEscalationPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DisableEscalationPlanResponse{}
	_body, _err := client.DisableEscalationPlanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableIntegrationConfigWithOptions(request *DisableIntegrationConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DisableIntegrationConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.IntegrationConfigId)) {
		body["integrationConfigId"] = request.IntegrationConfigId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableIntegrationConfig"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/integrationConfig/disable"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableIntegrationConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableIntegrationConfig(request *DisableIntegrationConfigRequest) (_result *DisableIntegrationConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DisableIntegrationConfigResponse{}
	_body, _err := client.DisableIntegrationConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableRouteRuleWithOptions(request *DisableRouteRuleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DisableRouteRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RouteRuleId)) {
		body["routeRuleId"] = request.RouteRuleId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableRouteRule"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/routeRule/disable"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableRouteRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableRouteRule(request *DisableRouteRuleRequest) (_result *DisableRouteRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DisableRouteRuleResponse{}
	_body, _err := client.DisableRouteRuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableServiceGroupWebhookWithOptions(request *DisableServiceGroupWebhookRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DisableServiceGroupWebhookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceGroupId)) {
		body["serviceGroupId"] = request.ServiceGroupId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableServiceGroupWebhook"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/services/group/disableWebhook"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableServiceGroupWebhookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableServiceGroupWebhook(request *DisableServiceGroupWebhookRequest) (_result *DisableServiceGroupWebhookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DisableServiceGroupWebhookResponse{}
	_body, _err := client.DisableServiceGroupWebhookWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableSubscriptionWithOptions(request *DisableSubscriptionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DisableSubscriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubscriptionId)) {
		body["subscriptionId"] = request.SubscriptionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableSubscription"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/notify/subscription/doDisable"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableSubscriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableSubscription(request *DisableSubscriptionRequest) (_result *DisableSubscriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DisableSubscriptionResponse{}
	_body, _err := client.DisableSubscriptionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableEscalationPlanWithOptions(request *EnableEscalationPlanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *EnableEscalationPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EscalationPlanId)) {
		body["escalationPlanId"] = request.EscalationPlanId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableEscalationPlan"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/escalationPlan/enable"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableEscalationPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableEscalationPlan(request *EnableEscalationPlanRequest) (_result *EnableEscalationPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EnableEscalationPlanResponse{}
	_body, _err := client.EnableEscalationPlanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableIntegrationConfigWithOptions(request *EnableIntegrationConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *EnableIntegrationConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.IntegrationConfigId)) {
		body["integrationConfigId"] = request.IntegrationConfigId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableIntegrationConfig"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/integrationConfig/enable"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableIntegrationConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableIntegrationConfig(request *EnableIntegrationConfigRequest) (_result *EnableIntegrationConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EnableIntegrationConfigResponse{}
	_body, _err := client.EnableIntegrationConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableRouteRuleWithOptions(request *EnableRouteRuleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *EnableRouteRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RouteRuleId)) {
		body["routeRuleId"] = request.RouteRuleId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableRouteRule"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/routeRule/enable"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableRouteRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableRouteRule(request *EnableRouteRuleRequest) (_result *EnableRouteRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EnableRouteRuleResponse{}
	_body, _err := client.EnableRouteRuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableServiceGroupWebhookWithOptions(request *EnableServiceGroupWebhookRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *EnableServiceGroupWebhookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceGroupId)) {
		body["serviceGroupId"] = request.ServiceGroupId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableServiceGroupWebhook"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/services/group/enableWebhook"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableServiceGroupWebhookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableServiceGroupWebhook(request *EnableServiceGroupWebhookRequest) (_result *EnableServiceGroupWebhookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EnableServiceGroupWebhookResponse{}
	_body, _err := client.EnableServiceGroupWebhookWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableSubscriptionWithOptions(request *EnableSubscriptionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *EnableSubscriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubscriptionId)) {
		body["subscriptionId"] = request.SubscriptionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableSubscription"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/notify/subscription/enable"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableSubscriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableSubscription(request *EnableSubscriptionRequest) (_result *EnableSubscriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EnableSubscriptionResponse{}
	_body, _err := client.EnableSubscriptionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FinishIncidentWithOptions(request *FinishIncidentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *FinishIncidentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentFinishReason)) {
		body["incidentFinishReason"] = request.IncidentFinishReason
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentFinishReasonDescription)) {
		body["incidentFinishReasonDescription"] = request.IncidentFinishReasonDescription
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentFinishSolution)) {
		body["incidentFinishSolution"] = request.IncidentFinishSolution
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentFinishSolutionDescription)) {
		body["incidentFinishSolutionDescription"] = request.IncidentFinishSolutionDescription
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentIds)) {
		body["incidentIds"] = request.IncidentIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FinishIncident"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/incident/finish"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &FinishIncidentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FinishIncident(request *FinishIncidentRequest) (_result *FinishIncidentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &FinishIncidentResponse{}
	_body, _err := client.FinishIncidentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FinishProblemWithOptions(request *FinishProblemRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *FinishProblemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemId)) {
		body["problemId"] = request.ProblemId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FinishProblem"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/problem/finish"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &FinishProblemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FinishProblem(request *FinishProblemRequest) (_result *FinishProblemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &FinishProblemResponse{}
	_body, _err := client.FinishProblemWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GeneratePictureLinkWithOptions(request *GeneratePictureLinkRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GeneratePictureLinkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keys)) {
		body["keys"] = request.Keys
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemId)) {
		body["problemId"] = request.ProblemId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GeneratePictureLink"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/rich/oss/getPictureLink"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GeneratePictureLinkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GeneratePictureLink(request *GeneratePictureLinkRequest) (_result *GeneratePictureLinkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GeneratePictureLinkResponse{}
	_body, _err := client.GeneratePictureLinkWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GeneratePictureUploadSignWithOptions(request *GeneratePictureUploadSignRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GeneratePictureUploadSignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Files)) {
		body["files"] = request.Files
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		body["instanceType"] = request.InstanceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GeneratePictureUploadSign"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/rich/oss/generatePostPolicy"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GeneratePictureUploadSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GeneratePictureUploadSign(request *GeneratePictureUploadSignRequest) (_result *GeneratePictureUploadSignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GeneratePictureUploadSignResponse{}
	_body, _err := client.GeneratePictureUploadSignWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateProblemPictureLinkWithOptions(request *GenerateProblemPictureLinkRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GenerateProblemPictureLinkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keys)) {
		body["keys"] = request.Keys
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemId)) {
		body["problemId"] = request.ProblemId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateProblemPictureLink"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/problem/process/oss/getPresignedLink"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateProblemPictureLinkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateProblemPictureLink(request *GenerateProblemPictureLinkRequest) (_result *GenerateProblemPictureLinkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenerateProblemPictureLinkResponse{}
	_body, _err := client.GenerateProblemPictureLinkWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateProblemPictureUploadSignWithOptions(request *GenerateProblemPictureUploadSignRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GenerateProblemPictureUploadSignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		body["fileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileSize)) {
		body["fileSize"] = request.FileSize
	}

	if !tea.BoolValue(util.IsUnset(request.FileType)) {
		body["fileType"] = request.FileType
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemId)) {
		body["problemId"] = request.ProblemId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateProblemPictureUploadSign"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/problem/process/oss/generatePostPolicy"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateProblemPictureUploadSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateProblemPictureUploadSign(request *GenerateProblemPictureUploadSignRequest) (_result *GenerateProblemPictureUploadSignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenerateProblemPictureUploadSignResponse{}
	_body, _err := client.GenerateProblemPictureUploadSignWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEscalationPlanWithOptions(request *GetEscalationPlanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetEscalationPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EscalationPlanId)) {
		body["escalationPlanId"] = request.EscalationPlanId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEscalationPlan"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/escalationPlan/detail"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEscalationPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEscalationPlan(request *GetEscalationPlanRequest) (_result *GetEscalationPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEscalationPlanResponse{}
	_body, _err := client.GetEscalationPlanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEventWithOptions(request *GetEventRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MonitorSourceId)) {
		body["monitorSourceId"] = request.MonitorSourceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEvent"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/events/getLastTimeEvent"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEvent(request *GetEventRequest) (_result *GetEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEventResponse{}
	_body, _err := client.GetEventWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHomePageGuidanceWithOptions(request *GetHomePageGuidanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetHomePageGuidanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHomePageGuidance"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/guidance/detail"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHomePageGuidanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHomePageGuidance(request *GetHomePageGuidanceRequest) (_result *GetHomePageGuidanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetHomePageGuidanceResponse{}
	_body, _err := client.GetHomePageGuidanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetIncidentWithOptions(request *GetIncidentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetIncidentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentId)) {
		body["incidentId"] = request.IncidentId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetIncident"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/incident/detail"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIncidentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetIncident(request *GetIncidentRequest) (_result *GetIncidentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIncidentResponse{}
	_body, _err := client.GetIncidentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetIncidentStatisticsWithOptions(request *GetIncidentStatisticsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetIncidentStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetIncidentStatistics"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/incident/count"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIncidentStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetIncidentStatistics(request *GetIncidentStatisticsRequest) (_result *GetIncidentStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIncidentStatisticsResponse{}
	_body, _err := client.GetIncidentStatisticsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetIncidentSubtotalCountWithOptions(request *GetIncidentSubtotalCountRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetIncidentSubtotalCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentIds)) {
		body["incidentIds"] = request.IncidentIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetIncidentSubtotalCount"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/incident/subtotal/count"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIncidentSubtotalCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetIncidentSubtotalCount(request *GetIncidentSubtotalCountRequest) (_result *GetIncidentSubtotalCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIncidentSubtotalCountResponse{}
	_body, _err := client.GetIncidentSubtotalCountWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetIntegrationConfigWithOptions(request *GetIntegrationConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetIntegrationConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.IntegrationConfigId)) {
		body["integrationConfigId"] = request.IntegrationConfigId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetIntegrationConfig"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/integrationConfig/detail"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIntegrationConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetIntegrationConfig(request *GetIntegrationConfigRequest) (_result *GetIntegrationConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIntegrationConfigResponse{}
	_body, _err := client.GetIntegrationConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProblemWithOptions(request *GetProblemRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProblemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemId)) {
		body["problemId"] = request.ProblemId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProblem"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/problem/detail"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProblemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProblem(request *GetProblemRequest) (_result *GetProblemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProblemResponse{}
	_body, _err := client.GetProblemWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProblemEffectionServiceWithOptions(request *GetProblemEffectionServiceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProblemEffectionServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EffectionServiceId)) {
		body["effectionServiceId"] = request.EffectionServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemId)) {
		body["problemId"] = request.ProblemId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProblemEffectionService"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/problem/process/effectionService/detail"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProblemEffectionServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProblemEffectionService(request *GetProblemEffectionServiceRequest) (_result *GetProblemEffectionServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProblemEffectionServiceResponse{}
	_body, _err := client.GetProblemEffectionServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProblemImprovementWithOptions(request *GetProblemImprovementRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProblemImprovementResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemId)) {
		body["problemId"] = request.ProblemId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProblemImprovement"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/problem/improvement/detail"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProblemImprovementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProblemImprovement(request *GetProblemImprovementRequest) (_result *GetProblemImprovementResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProblemImprovementResponse{}
	_body, _err := client.GetProblemImprovementWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProblemPreviewWithOptions(request *GetProblemPreviewRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProblemPreviewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EffectServiceIds)) {
		body["effectServiceIds"] = request.EffectServiceIds
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentId)) {
		body["incidentId"] = request.IncidentId
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemId)) {
		body["problemId"] = request.ProblemId
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemLevel)) {
		body["problemLevel"] = request.ProblemLevel
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemNotifyType)) {
		body["problemNotifyType"] = request.ProblemNotifyType
	}

	if !tea.BoolValue(util.IsUnset(request.RelatedServiceId)) {
		body["relatedServiceId"] = request.RelatedServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceGroupIds)) {
		body["serviceGroupIds"] = request.ServiceGroupIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProblemPreview"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/problem/preview"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProblemPreviewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProblemPreview(request *GetProblemPreviewRequest) (_result *GetProblemPreviewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProblemPreviewResponse{}
	_body, _err := client.GetProblemPreviewWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetResourceStatisticsWithOptions(request *GetResourceStatisticsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetResourceStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResourceStatistics"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/config/resource/count"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResourceStatistics(request *GetResourceStatisticsRequest) (_result *GetResourceStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetResourceStatisticsResponse{}
	_body, _err := client.GetResourceStatisticsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRichTextWithOptions(request *GetRichTextRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetRichTextResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		body["instanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.RichTextId)) {
		body["richTextId"] = request.RichTextId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRichText"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/rich/detail"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRichTextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRichText(request *GetRichTextRequest) (_result *GetRichTextResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRichTextResponse{}
	_body, _err := client.GetRichTextWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRouteRuleWithOptions(request *GetRouteRuleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetRouteRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RouteRuleId)) {
		body["routeRuleId"] = request.RouteRuleId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRouteRule"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/routeRule/detail"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRouteRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRouteRule(request *GetRouteRuleRequest) (_result *GetRouteRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRouteRuleResponse{}
	_body, _err := client.GetRouteRuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetServiceWithOptions(request *GetServiceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		body["serviceId"] = request.ServiceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetService"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/services/detail"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetService(request *GetServiceRequest) (_result *GetServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetServiceResponse{}
	_body, _err := client.GetServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetServiceGroupWithOptions(request *GetServiceGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetServiceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceGroupId)) {
		body["serviceGroupId"] = request.ServiceGroupId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceGroup"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/services/group/detail"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetServiceGroup(request *GetServiceGroupRequest) (_result *GetServiceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetServiceGroupResponse{}
	_body, _err := client.GetServiceGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetServiceGroupPersonSchedulingWithOptions(request *GetServiceGroupPersonSchedulingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetServiceGroupPersonSchedulingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceGroupId)) {
		body["serviceGroupId"] = request.ServiceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceGroupPersonScheduling"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/services/group/scheduling/user/getScheduling"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceGroupPersonSchedulingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetServiceGroupPersonScheduling(request *GetServiceGroupPersonSchedulingRequest) (_result *GetServiceGroupPersonSchedulingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetServiceGroupPersonSchedulingResponse{}
	_body, _err := client.GetServiceGroupPersonSchedulingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetServiceGroupSchedulingWithOptions(request *GetServiceGroupSchedulingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetServiceGroupSchedulingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceGroupId)) {
		body["serviceGroupId"] = request.ServiceGroupId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceGroupScheduling"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/services/group/scheduling/detail"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceGroupSchedulingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetServiceGroupScheduling(request *GetServiceGroupSchedulingRequest) (_result *GetServiceGroupSchedulingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetServiceGroupSchedulingResponse{}
	_body, _err := client.GetServiceGroupSchedulingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetServiceGroupSchedulingPreviewWithOptions(request *GetServiceGroupSchedulingPreviewRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetServiceGroupSchedulingPreviewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.FastScheduling)) {
		body["fastScheduling"] = request.FastScheduling
	}

	if !tea.BoolValue(util.IsUnset(request.FineScheduling)) {
		body["fineScheduling"] = request.FineScheduling
	}

	if !tea.BoolValue(util.IsUnset(request.SchedulingWay)) {
		body["schedulingWay"] = request.SchedulingWay
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceGroupId)) {
		body["serviceGroupId"] = request.ServiceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceGroupSchedulingPreview"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/services/group/scheduling/preview"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceGroupSchedulingPreviewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetServiceGroupSchedulingPreview(request *GetServiceGroupSchedulingPreviewRequest) (_result *GetServiceGroupSchedulingPreviewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetServiceGroupSchedulingPreviewResponse{}
	_body, _err := client.GetServiceGroupSchedulingPreviewWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetServiceGroupSpecialPersonSchedulingWithOptions(request *GetServiceGroupSpecialPersonSchedulingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetServiceGroupSpecialPersonSchedulingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceGroupId)) {
		body["serviceGroupId"] = request.ServiceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceGroupSpecialPersonScheduling"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/services/group/scheduling/getUserScheduling"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceGroupSpecialPersonSchedulingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetServiceGroupSpecialPersonScheduling(request *GetServiceGroupSpecialPersonSchedulingRequest) (_result *GetServiceGroupSpecialPersonSchedulingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetServiceGroupSpecialPersonSchedulingResponse{}
	_body, _err := client.GetServiceGroupSpecialPersonSchedulingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSimilarIncidentStatisticsWithOptions(request *GetSimilarIncidentStatisticsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSimilarIncidentStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTime)) {
		body["createTime"] = request.CreateTime
	}

	if !tea.BoolValue(util.IsUnset(request.Events)) {
		body["events"] = request.Events
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentId)) {
		body["incidentId"] = request.IncidentId
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentTitle)) {
		body["incidentTitle"] = request.IncidentTitle
	}

	if !tea.BoolValue(util.IsUnset(request.RelatedServiceId)) {
		body["relatedServiceId"] = request.RelatedServiceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSimilarIncidentStatistics"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/incident/similarIncident/statistics"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSimilarIncidentStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSimilarIncidentStatistics(request *GetSimilarIncidentStatisticsRequest) (_result *GetSimilarIncidentStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSimilarIncidentStatisticsResponse{}
	_body, _err := client.GetSimilarIncidentStatisticsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSubscriptionWithOptions(request *GetSubscriptionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSubscriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NotFilterScopeObjectDeleted)) {
		body["notFilterScopeObjectDeleted"] = request.NotFilterScopeObjectDeleted
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriptionId)) {
		body["subscriptionId"] = request.SubscriptionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSubscription"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/notify/subscription/detail"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSubscriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSubscription(request *GetSubscriptionRequest) (_result *GetSubscriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSubscriptionResponse{}
	_body, _err := client.GetSubscriptionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTenantApplicationWithOptions(request *GetTenantApplicationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTenantApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTenantApplication"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/mobileApp/detail"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTenantApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTenantApplication(request *GetTenantApplicationRequest) (_result *GetTenantApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTenantApplicationResponse{}
	_body, _err := client.GetTenantApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTenantStatusWithOptions(request *GetTenantStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTenantStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TenantRamId)) {
		body["tenantRamId"] = request.TenantRamId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTenantStatus"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/tenant/getTenantStatus"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTenantStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTenantStatus(request *GetTenantStatusRequest) (_result *GetTenantStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTenantStatusResponse{}
	_body, _err := client.GetTenantStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserWithOptions(request *GetUserRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUser"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/user/getUser"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUser(request *GetUserRequest) (_result *GetUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserGuideStatusWithOptions(request *GetUserGuideStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetUserGuideStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserGuideStatus"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/user/guide/status"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserGuideStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserGuideStatus(request *GetUserGuideStatusRequest) (_result *GetUserGuideStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUserGuideStatusResponse{}
	_body, _err := client.GetUserGuideStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAlertsWithOptions(request *ListAlertsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAlertsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertLevel)) {
		body["alertLevel"] = request.AlertLevel
	}

	if !tea.BoolValue(util.IsUnset(request.AlertName)) {
		body["alertName"] = request.AlertName
	}

	if !tea.BoolValue(util.IsUnset(request.AlertSourceName)) {
		body["alertSourceName"] = request.AlertSourceName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MonitorSourceId)) {
		body["monitorSourceId"] = request.MonitorSourceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RelatedServiceId)) {
		body["relatedServiceId"] = request.RelatedServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		body["ruleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAlerts"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/alerts/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAlertsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAlerts(request *ListAlertsRequest) (_result *ListAlertsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAlertsResponse{}
	_body, _err := client.ListAlertsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListByMonitorSourceIdWithOptions(request *ListByMonitorSourceIdRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListByMonitorSourceIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MonitorSourceId)) {
		body["monitorSourceId"] = request.MonitorSourceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListByMonitorSourceId"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/routeRule/listByMonitorSourceId"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListByMonitorSourceIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListByMonitorSourceId(request *ListByMonitorSourceIdRequest) (_result *ListByMonitorSourceIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListByMonitorSourceIdResponse{}
	_body, _err := client.ListByMonitorSourceIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListChartDataForServiceGroupWithOptions(request *ListChartDataForServiceGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListChartDataForServiceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListChartDataForServiceGroup"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/statistics/chartDataForServiceGroup/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListChartDataForServiceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListChartDataForServiceGroup(request *ListChartDataForServiceGroupRequest) (_result *ListChartDataForServiceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListChartDataForServiceGroupResponse{}
	_body, _err := client.ListChartDataForServiceGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListChartDataForUserWithOptions(request *ListChartDataForUserRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListChartDataForUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListChartDataForUser"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/statistics/chartDataForUser/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListChartDataForUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListChartDataForUser(request *ListChartDataForUserRequest) (_result *ListChartDataForUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListChartDataForUserResponse{}
	_body, _err := client.ListChartDataForUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListConfigsWithOptions(request *ListConfigsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListConfigs"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/config/all"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListConfigs(request *ListConfigsRequest) (_result *ListConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListConfigsResponse{}
	_body, _err := client.ListConfigsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDataReportForServiceGroupWithOptions(request *ListDataReportForServiceGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDataReportForServiceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceGroupName)) {
		body["serviceGroupName"] = request.ServiceGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDataReportForServiceGroup"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/statistics/listDataReportForServiceGroup"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDataReportForServiceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDataReportForServiceGroup(request *ListDataReportForServiceGroupRequest) (_result *ListDataReportForServiceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDataReportForServiceGroupResponse{}
	_body, _err := client.ListDataReportForServiceGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDataReportForUserWithOptions(request *ListDataReportForUserRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDataReportForUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDataReportForUser"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/statistics/listDataReportForUser"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDataReportForUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDataReportForUser(request *ListDataReportForUserRequest) (_result *ListDataReportForUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDataReportForUserResponse{}
	_body, _err := client.ListDataReportForUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDictionariesWithOptions(request *ListDictionariesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDictionariesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDictionaries"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dict/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDictionariesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDictionaries(request *ListDictionariesRequest) (_result *ListDictionariesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDictionariesResponse{}
	_body, _err := client.ListDictionariesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEscalationPlanServicesWithOptions(request *ListEscalationPlanServicesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEscalationPlanServicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEscalationPlanServices"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/escalationPlan/services"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEscalationPlanServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEscalationPlanServices(request *ListEscalationPlanServicesRequest) (_result *ListEscalationPlanServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEscalationPlanServicesResponse{}
	_body, _err := client.ListEscalationPlanServicesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEscalationPlansWithOptions(request *ListEscalationPlansRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEscalationPlansResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EscalationPlanName)) {
		body["escalationPlanName"] = request.EscalationPlanName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		body["serviceName"] = request.ServiceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEscalationPlans"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/escalationPlan/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEscalationPlansResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEscalationPlans(request *ListEscalationPlansRequest) (_result *ListEscalationPlansResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEscalationPlansResponse{}
	_body, _err := client.ListEscalationPlansWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEscalationPlansByNoticeObjectWithOptions(request *ListEscalationPlansByNoticeObjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEscalationPlansByNoticeObjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NoticeObjectId)) {
		body["noticeObjectId"] = request.NoticeObjectId
	}

	if !tea.BoolValue(util.IsUnset(request.NoticeObjectType)) {
		body["noticeObjectType"] = request.NoticeObjectType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEscalationPlansByNoticeObject"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/escalationPlan/listByNoticeObject"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEscalationPlansByNoticeObjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEscalationPlansByNoticeObject(request *ListEscalationPlansByNoticeObjectRequest) (_result *ListEscalationPlansByNoticeObjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEscalationPlansByNoticeObjectResponse{}
	_body, _err := client.ListEscalationPlansByNoticeObjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListIncidentDetailEscalationPlansWithOptions(request *ListIncidentDetailEscalationPlansRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListIncidentDetailEscalationPlansResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentId)) {
		body["incidentId"] = request.IncidentId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIncidentDetailEscalationPlans"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/incident/detail/escalation"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIncidentDetailEscalationPlansResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListIncidentDetailEscalationPlans(request *ListIncidentDetailEscalationPlansRequest) (_result *ListIncidentDetailEscalationPlansResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIncidentDetailEscalationPlansResponse{}
	_body, _err := client.ListIncidentDetailEscalationPlansWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListIncidentDetailTimelinesWithOptions(request *ListIncidentDetailTimelinesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListIncidentDetailTimelinesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.IdSort)) {
		body["idSort"] = request.IdSort
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentId)) {
		body["incidentId"] = request.IncidentId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIncidentDetailTimelines"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/incident/detail/timeline"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIncidentDetailTimelinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListIncidentDetailTimelines(request *ListIncidentDetailTimelinesRequest) (_result *ListIncidentDetailTimelinesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIncidentDetailTimelinesResponse{}
	_body, _err := client.ListIncidentDetailTimelinesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListIncidentSubtotalsWithOptions(request *ListIncidentSubtotalsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListIncidentSubtotalsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentId)) {
		body["incidentId"] = request.IncidentId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIncidentSubtotals"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/incident/list/subtotal"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIncidentSubtotalsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListIncidentSubtotals(request *ListIncidentSubtotalsRequest) (_result *ListIncidentSubtotalsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIncidentSubtotalsResponse{}
	_body, _err := client.ListIncidentSubtotalsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListIncidentTimelinesWithOptions(request *ListIncidentTimelinesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListIncidentTimelinesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIncidentTimelines"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/incident/timeline"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIncidentTimelinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListIncidentTimelines(request *ListIncidentTimelinesRequest) (_result *ListIncidentTimelinesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIncidentTimelinesResponse{}
	_body, _err := client.ListIncidentTimelinesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListIncidentsWithOptions(request *ListIncidentsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListIncidentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CreateEndTime)) {
		body["createEndTime"] = request.CreateEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.CreateStartTime)) {
		body["createStartTime"] = request.CreateStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Effect)) {
		body["effect"] = request.Effect
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentLevel)) {
		body["incidentLevel"] = request.IncidentLevel
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentStatus)) {
		body["incidentStatus"] = request.IncidentStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Me)) {
		body["me"] = request.Me
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RelationServiceId)) {
		body["relationServiceId"] = request.RelationServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		body["ruleName"] = request.RuleName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIncidents"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/incident/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIncidentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListIncidents(request *ListIncidentsRequest) (_result *ListIncidentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIncidentsResponse{}
	_body, _err := client.ListIncidentsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListIntegrationConfigTimelinesWithOptions(request *ListIntegrationConfigTimelinesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListIntegrationConfigTimelinesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.IntegrationConfigId)) {
		body["integrationConfigId"] = request.IntegrationConfigId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIntegrationConfigTimelines"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/integrationConfig/timeline"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIntegrationConfigTimelinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListIntegrationConfigTimelines(request *ListIntegrationConfigTimelinesRequest) (_result *ListIntegrationConfigTimelinesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIntegrationConfigTimelinesResponse{}
	_body, _err := client.ListIntegrationConfigTimelinesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListIntegrationConfigsWithOptions(request *ListIntegrationConfigsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListIntegrationConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.MonitorSourceName)) {
		body["monitorSourceName"] = request.MonitorSourceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIntegrationConfigs"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/integrationConfig/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIntegrationConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListIntegrationConfigs(request *ListIntegrationConfigsRequest) (_result *ListIntegrationConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIntegrationConfigsResponse{}
	_body, _err := client.ListIntegrationConfigsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMonitorSourcesWithOptions(request *ListMonitorSourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListMonitorSourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMonitorSources"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/monitorSource/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMonitorSourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMonitorSources(request *ListMonitorSourcesRequest) (_result *ListMonitorSourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMonitorSourcesResponse{}
	_body, _err := client.ListMonitorSourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProblemDetailOperationsWithOptions(request *ListProblemDetailOperationsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProblemDetailOperationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTimeSort)) {
		body["createTimeSort"] = request.CreateTimeSort
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemId)) {
		body["problemId"] = request.ProblemId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProblemDetailOperations"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/problem/detail/operations"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProblemDetailOperationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProblemDetailOperations(request *ListProblemDetailOperationsRequest) (_result *ListProblemDetailOperationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProblemDetailOperationsResponse{}
	_body, _err := client.ListProblemDetailOperationsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProblemOperationsWithOptions(request *ListProblemOperationsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProblemOperationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProblemOperations"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/problem/operations"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProblemOperationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProblemOperations(request *ListProblemOperationsRequest) (_result *ListProblemOperationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProblemOperationsResponse{}
	_body, _err := client.ListProblemOperationsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProblemSubtotalsWithOptions(request *ListProblemSubtotalsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProblemSubtotalsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemId)) {
		body["problemId"] = request.ProblemId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProblemSubtotals"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/problem/list/subtotal"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProblemSubtotalsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProblemSubtotals(request *ListProblemSubtotalsRequest) (_result *ListProblemSubtotalsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProblemSubtotalsResponse{}
	_body, _err := client.ListProblemSubtotalsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProblemTimeLinesWithOptions(request *ListProblemTimeLinesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProblemTimeLinesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemId)) {
		body["problemId"] = request.ProblemId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProblemTimeLines"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/problem/detail/timeLines"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProblemTimeLinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProblemTimeLines(request *ListProblemTimeLinesRequest) (_result *ListProblemTimeLinesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProblemTimeLinesResponse{}
	_body, _err := client.ListProblemTimeLinesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProblemsWithOptions(request *ListProblemsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProblemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AffectServiceId)) {
		body["affectServiceId"] = request.AffectServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DiscoveryEndTime)) {
		body["discoveryEndTime"] = request.DiscoveryEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.DiscoveryStartTime)) {
		body["discoveryStartTime"] = request.DiscoveryStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.MainHandlerId)) {
		body["mainHandlerId"] = request.MainHandlerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemLevel)) {
		body["problemLevel"] = request.ProblemLevel
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemStatus)) {
		body["problemStatus"] = request.ProblemStatus
	}

	if !tea.BoolValue(util.IsUnset(request.QueryType)) {
		body["queryType"] = request.QueryType
	}

	if !tea.BoolValue(util.IsUnset(request.RepeaterId)) {
		body["repeaterId"] = request.RepeaterId
	}

	if !tea.BoolValue(util.IsUnset(request.RestoreEndTime)) {
		body["restoreEndTime"] = request.RestoreEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RestoreStartTime)) {
		body["restoreStartTime"] = request.RestoreStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceGroupId)) {
		body["serviceGroupId"] = request.ServiceGroupId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProblems"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/problem/listProblems"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProblemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProblems(request *ListProblemsRequest) (_result *ListProblemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProblemsResponse{}
	_body, _err := client.ListProblemsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRouteRulesWithOptions(request *ListRouteRulesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRouteRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.NotFilterRouteRuleDeleted)) {
		body["notFilterRouteRuleDeleted"] = request.NotFilterRouteRuleDeleted
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RouteType)) {
		body["routeType"] = request.RouteType
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		body["ruleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		body["serviceName"] = request.ServiceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRouteRules"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/routeRule/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRouteRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRouteRules(request *ListRouteRulesRequest) (_result *ListRouteRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRouteRulesResponse{}
	_body, _err := client.ListRouteRulesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRouteRulesByAssignWhoIdWithOptions(request *ListRouteRulesByAssignWhoIdRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRouteRulesByAssignWhoIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssignWhoId)) {
		body["assignWhoId"] = request.AssignWhoId
	}

	if !tea.BoolValue(util.IsUnset(request.AssignWhoType)) {
		body["assignWhoType"] = request.AssignWhoType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRouteRulesByAssignWhoId"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/routeRule/listByAssignWhoId"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRouteRulesByAssignWhoIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRouteRulesByAssignWhoId(request *ListRouteRulesByAssignWhoIdRequest) (_result *ListRouteRulesByAssignWhoIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRouteRulesByAssignWhoIdResponse{}
	_body, _err := client.ListRouteRulesByAssignWhoIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRouteRulesByServiceWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRouteRulesByServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListRouteRulesByService"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/routeRule/listByService"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRouteRulesByServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRouteRulesByService() (_result *ListRouteRulesByServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRouteRulesByServiceResponse{}
	_body, _err := client.ListRouteRulesByServiceWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListServiceGroupMonitorSourceTemplatesWithOptions(request *ListServiceGroupMonitorSourceTemplatesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListServiceGroupMonitorSourceTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["requestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceGroupId)) {
		body["serviceGroupId"] = request.ServiceGroupId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceGroupMonitorSourceTemplates"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/services/group/listServiceGroupMonitorSourceTemplates"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceGroupMonitorSourceTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServiceGroupMonitorSourceTemplates(request *ListServiceGroupMonitorSourceTemplatesRequest) (_result *ListServiceGroupMonitorSourceTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListServiceGroupMonitorSourceTemplatesResponse{}
	_body, _err := client.ListServiceGroupMonitorSourceTemplatesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListServiceGroupsWithOptions(request *ListServiceGroupsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListServiceGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.IsScheduled)) {
		body["isScheduled"] = request.IsScheduled
	}

	if !tea.BoolValue(util.IsUnset(request.OrderByScheduleStatus)) {
		body["orderByScheduleStatus"] = request.OrderByScheduleStatus
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryName)) {
		body["queryName"] = request.QueryName
	}

	if !tea.BoolValue(util.IsUnset(request.QueryType)) {
		body["queryType"] = request.QueryType
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceGroups"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/services/group/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServiceGroups(request *ListServiceGroupsRequest) (_result *ListServiceGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListServiceGroupsResponse{}
	_body, _err := client.ListServiceGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListServiceGroupsByUserIdWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListServiceGroupsByUserIdResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceGroupsByUserId"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/services/group/listByUserId"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceGroupsByUserIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServiceGroupsByUserId() (_result *ListServiceGroupsByUserIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListServiceGroupsByUserIdResponse{}
	_body, _err := client.ListServiceGroupsByUserIdWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListServicesWithOptions(request *ListServicesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListServicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		body["serviceName"] = request.ServiceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServices"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/services/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServices(request *ListServicesRequest) (_result *ListServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListServicesResponse{}
	_body, _err := client.ListServicesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSourceEventsWithOptions(request *ListSourceEventsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSourceEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		body["instanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartRowKey)) {
		body["startRowKey"] = request.StartRowKey
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.StopRowKey)) {
		body["stopRowKey"] = request.StopRowKey
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSourceEvents"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/events/listOriginalEvent"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSourceEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSourceEvents(request *ListSourceEventsRequest) (_result *ListSourceEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSourceEventsResponse{}
	_body, _err := client.ListSourceEventsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSourceEventsForMonitorSourceWithOptions(request *ListSourceEventsForMonitorSourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSourceEventsForMonitorSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MonitorSourceId)) {
		body["monitorSourceId"] = request.MonitorSourceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSourceEventsForMonitorSource"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/events/queryLastestEvents"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSourceEventsForMonitorSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSourceEventsForMonitorSource(request *ListSourceEventsForMonitorSourceRequest) (_result *ListSourceEventsForMonitorSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSourceEventsForMonitorSourceResponse{}
	_body, _err := client.ListSourceEventsForMonitorSourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSubscriptionServiceGroupsWithOptions(request *ListSubscriptionServiceGroupsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSubscriptionServiceGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceIds)) {
		body["serviceIds"] = request.ServiceIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSubscriptionServiceGroups"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/problem/serviceGroup/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSubscriptionServiceGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSubscriptionServiceGroups(request *ListSubscriptionServiceGroupsRequest) (_result *ListSubscriptionServiceGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSubscriptionServiceGroupsResponse{}
	_body, _err := client.ListSubscriptionServiceGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSubscriptionsWithOptions(request *ListSubscriptionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSubscriptionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.NotFilterScopeObjectDeleted)) {
		body["notFilterScopeObjectDeleted"] = request.NotFilterScopeObjectDeleted
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyObject)) {
		body["notifyObject"] = request.NotifyObject
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyObjectType)) {
		body["notifyObjectType"] = request.NotifyObjectType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		body["scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeObject)) {
		body["scopeObject"] = request.ScopeObject
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriptionTitle)) {
		body["subscriptionTitle"] = request.SubscriptionTitle
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSubscriptions"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/notify/subscription/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSubscriptionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSubscriptions(request *ListSubscriptionsRequest) (_result *ListSubscriptionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSubscriptionsResponse{}
	_body, _err := client.ListSubscriptionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTrendForSourceEventWithOptions(request *ListTrendForSourceEventRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTrendForSourceEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		body["instanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["requestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TimeUnit)) {
		body["timeUnit"] = request.TimeUnit
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTrendForSourceEvent"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/events/querySourceEventTrend"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTrendForSourceEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTrendForSourceEvent(request *ListTrendForSourceEventRequest) (_result *ListTrendForSourceEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTrendForSourceEventResponse{}
	_body, _err := client.ListTrendForSourceEventWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserSerivceGroupsWithOptions(request *ListUserSerivceGroupsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListUserSerivceGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserSerivceGroups"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/user/preview/detail"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserSerivceGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserSerivceGroups(request *ListUserSerivceGroupsRequest) (_result *ListUserSerivceGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListUserSerivceGroupsResponse{}
	_body, _err := client.ListUserSerivceGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUsersWithOptions(request *ListUsersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		body["phone"] = request.Phone
	}

	if !tea.BoolValue(util.IsUnset(request.RamId)) {
		body["ramId"] = request.RamId
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		body["scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.SynergyChannel)) {
		body["synergyChannel"] = request.SynergyChannel
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		body["username"] = request.Username
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUsers"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/user/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUsers(request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecoverProblemWithOptions(request *RecoverProblemRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RecoverProblemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProblemId)) {
		body["problemId"] = request.ProblemId
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemNotifyType)) {
		body["problemNotifyType"] = request.ProblemNotifyType
	}

	if !tea.BoolValue(util.IsUnset(request.RecoveryTime)) {
		body["recoveryTime"] = request.RecoveryTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecoverProblem"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/problem/recovery"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RecoverProblemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecoverProblem(request *RecoverProblemRequest) (_result *RecoverProblemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RecoverProblemResponse{}
	_body, _err := client.RecoverProblemWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefreshIntegrationConfigKeyWithOptions(request *RefreshIntegrationConfigKeyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RefreshIntegrationConfigKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.IntegrationConfigId)) {
		body["integrationConfigId"] = request.IntegrationConfigId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RefreshIntegrationConfigKey"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/integrationConfig/refreshKey"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RefreshIntegrationConfigKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefreshIntegrationConfigKey(request *RefreshIntegrationConfigKeyRequest) (_result *RefreshIntegrationConfigKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RefreshIntegrationConfigKeyResponse{}
	_body, _err := client.RefreshIntegrationConfigKeyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveIntegrationConfigWithOptions(request *RemoveIntegrationConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveIntegrationConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.IntegrationConfigId)) {
		body["integrationConfigId"] = request.IntegrationConfigId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveIntegrationConfig"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/integrationConfig/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveIntegrationConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveIntegrationConfig(request *RemoveIntegrationConfigRequest) (_result *RemoveIntegrationConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveIntegrationConfigResponse{}
	_body, _err := client.RemoveIntegrationConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveProblemServiceGroupWithOptions(request *RemoveProblemServiceGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveProblemServiceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProblemId)) {
		body["problemId"] = request.ProblemId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceGroupIds)) {
		body["serviceGroupIds"] = request.ServiceGroupIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveProblemServiceGroup"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/problem/removeServiceGroup"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveProblemServiceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveProblemServiceGroup(request *RemoveProblemServiceGroupRequest) (_result *RemoveProblemServiceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveProblemServiceGroupResponse{}
	_body, _err := client.RemoveProblemServiceGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReplayProblemWithOptions(request *ReplayProblemRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ReplayProblemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemId)) {
		body["problemId"] = request.ProblemId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplayDutyUserId)) {
		body["replayDutyUserId"] = request.ReplayDutyUserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReplayProblem"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/problem/replay"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ReplayProblemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReplayProblem(request *ReplayProblemRequest) (_result *ReplayProblemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReplayProblemResponse{}
	_body, _err := client.ReplayProblemWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RespondIncidentWithOptions(request *RespondIncidentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RespondIncidentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentIds)) {
		body["incidentIds"] = request.IncidentIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RespondIncident"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/incident/response"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RespondIncidentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RespondIncident(request *RespondIncidentRequest) (_result *RespondIncidentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RespondIncidentResponse{}
	_body, _err := client.RespondIncidentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RevokeProblemRecoveryWithOptions(request *RevokeProblemRecoveryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RevokeProblemRecoveryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemId)) {
		body["problemId"] = request.ProblemId
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemNotifyType)) {
		body["problemNotifyType"] = request.ProblemNotifyType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RevokeProblemRecovery"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/problem/revoke"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RevokeProblemRecoveryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RevokeProblemRecovery(request *RevokeProblemRecoveryRequest) (_result *RevokeProblemRecoveryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RevokeProblemRecoveryResponse{}
	_body, _err := client.RevokeProblemRecoveryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindUserWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *UnbindUserResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindUser"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/user/unbind"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindUser() (_result *UnbindUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UnbindUserResponse{}
	_body, _err := client.UnbindUserWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateEscalationPlanWithOptions(request *UpdateEscalationPlanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateEscalationPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EscalationPlanDescription)) {
		body["escalationPlanDescription"] = request.EscalationPlanDescription
	}

	if !tea.BoolValue(util.IsUnset(request.EscalationPlanId)) {
		body["escalationPlanId"] = request.EscalationPlanId
	}

	if !tea.BoolValue(util.IsUnset(request.EscalationPlanName)) {
		body["escalationPlanName"] = request.EscalationPlanName
	}

	if !tea.BoolValue(util.IsUnset(request.EscalationPlanRules)) {
		body["escalationPlanRules"] = request.EscalationPlanRules
	}

	if !tea.BoolValue(util.IsUnset(request.EscalationPlanScopeObjects)) {
		body["escalationPlanScopeObjects"] = request.EscalationPlanScopeObjects
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateEscalationPlan"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/escalationPlan/update"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateEscalationPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateEscalationPlan(request *UpdateEscalationPlanRequest) (_result *UpdateEscalationPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateEscalationPlanResponse{}
	_body, _err := client.UpdateEscalationPlanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateIncidentWithOptions(request *UpdateIncidentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateIncidentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Effect)) {
		body["effect"] = request.Effect
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentId)) {
		body["incidentId"] = request.IncidentId
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentLevel)) {
		body["incidentLevel"] = request.IncidentLevel
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentTitle)) {
		body["incidentTitle"] = request.IncidentTitle
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateIncident"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/incident/update"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateIncidentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateIncident(request *UpdateIncidentRequest) (_result *UpdateIncidentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateIncidentResponse{}
	_body, _err := client.UpdateIncidentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateIntegrationConfigWithOptions(request *UpdateIntegrationConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateIntegrationConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		body["accessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.IntegrationConfigId)) {
		body["integrationConfigId"] = request.IntegrationConfigId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateIntegrationConfig"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/integrationConfig/update"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateIntegrationConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateIntegrationConfig(request *UpdateIntegrationConfigRequest) (_result *UpdateIntegrationConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateIntegrationConfigResponse{}
	_body, _err := client.UpdateIntegrationConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProblemWithOptions(request *UpdateProblemRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateProblemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Feedback)) {
		body["feedback"] = request.Feedback
	}

	if !tea.BoolValue(util.IsUnset(request.Level)) {
		body["level"] = request.Level
	}

	if !tea.BoolValue(util.IsUnset(request.MainHandlerId)) {
		body["mainHandlerId"] = request.MainHandlerId
	}

	if !tea.BoolValue(util.IsUnset(request.PreliminaryReason)) {
		body["preliminaryReason"] = request.PreliminaryReason
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemId)) {
		body["problemId"] = request.ProblemId
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemName)) {
		body["problemName"] = request.ProblemName
	}

	if !tea.BoolValue(util.IsUnset(request.ProgressSummary)) {
		body["progressSummary"] = request.ProgressSummary
	}

	if !tea.BoolValue(util.IsUnset(request.ProgressSummaryRichTextId)) {
		body["progressSummaryRichTextId"] = request.ProgressSummaryRichTextId
	}

	if !tea.BoolValue(util.IsUnset(request.RelatedServiceId)) {
		body["relatedServiceId"] = request.RelatedServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceGroupIds)) {
		body["serviceGroupIds"] = request.ServiceGroupIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProblem"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/problem/update"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateProblemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProblem(request *UpdateProblemRequest) (_result *UpdateProblemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateProblemResponse{}
	_body, _err := client.UpdateProblemWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProblemEffectionServiceWithOptions(request *UpdateProblemEffectionServiceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateProblemEffectionServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EffectionServiceId)) {
		body["effectionServiceId"] = request.EffectionServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.Level)) {
		body["level"] = request.Level
	}

	if !tea.BoolValue(util.IsUnset(request.PicUrl)) {
		body["picUrl"] = request.PicUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemId)) {
		body["problemId"] = request.ProblemId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		body["serviceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProblemEffectionService"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/problem/process/effectionService/update"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateProblemEffectionServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProblemEffectionService(request *UpdateProblemEffectionServiceRequest) (_result *UpdateProblemEffectionServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateProblemEffectionServiceResponse{}
	_body, _err := client.UpdateProblemEffectionServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProblemImprovementWithOptions(request *UpdateProblemImprovementRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateProblemImprovementResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CustomProblemReason)) {
		body["customProblemReason"] = request.CustomProblemReason
	}

	if !tea.BoolValue(util.IsUnset(request.DiscoverSource)) {
		body["discoverSource"] = request.DiscoverSource
	}

	if !tea.BoolValue(util.IsUnset(request.DutyDepartmentId)) {
		body["dutyDepartmentId"] = request.DutyDepartmentId
	}

	if !tea.BoolValue(util.IsUnset(request.DutyDepartmentName)) {
		body["dutyDepartmentName"] = request.DutyDepartmentName
	}

	if !tea.BoolValue(util.IsUnset(request.DutyUserId)) {
		body["dutyUserId"] = request.DutyUserId
	}

	if !tea.BoolValue(util.IsUnset(request.InjectionMode)) {
		body["injectionMode"] = request.InjectionMode
	}

	if !tea.BoolValue(util.IsUnset(request.MonitorSourceName)) {
		body["monitorSourceName"] = request.MonitorSourceName
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemId)) {
		body["problemId"] = request.ProblemId
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemReason)) {
		body["problemReason"] = request.ProblemReason
	}

	if !tea.BoolValue(util.IsUnset(request.RecentActivity)) {
		body["recentActivity"] = request.RecentActivity
	}

	if !tea.BoolValue(util.IsUnset(request.RecoveryMode)) {
		body["recoveryMode"] = request.RecoveryMode
	}

	if !tea.BoolValue(util.IsUnset(request.RelationChanges)) {
		body["relationChanges"] = request.RelationChanges
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.ReplayDutyUserId)) {
		body["replayDutyUserId"] = request.ReplayDutyUserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserReport)) {
		body["userReport"] = request.UserReport
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProblemImprovement"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/problem/improvement/update"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateProblemImprovementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProblemImprovement(request *UpdateProblemImprovementRequest) (_result *UpdateProblemImprovementResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateProblemImprovementResponse{}
	_body, _err := client.UpdateProblemImprovementWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProblemMeasureWithOptions(request *UpdateProblemMeasureRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateProblemMeasureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CheckStandard)) {
		body["checkStandard"] = request.CheckStandard
	}

	if !tea.BoolValue(util.IsUnset(request.CheckUserId)) {
		body["checkUserId"] = request.CheckUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.DirectorId)) {
		body["directorId"] = request.DirectorId
	}

	if !tea.BoolValue(util.IsUnset(request.MeasureId)) {
		body["measureId"] = request.MeasureId
	}

	if !tea.BoolValue(util.IsUnset(request.PlanFinishTime)) {
		body["planFinishTime"] = request.PlanFinishTime
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemId)) {
		body["problemId"] = request.ProblemId
	}

	if !tea.BoolValue(util.IsUnset(request.StalkerId)) {
		body["stalkerId"] = request.StalkerId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProblemMeasure"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/problem/improvement/measure/update"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateProblemMeasureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProblemMeasure(request *UpdateProblemMeasureRequest) (_result *UpdateProblemMeasureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateProblemMeasureResponse{}
	_body, _err := client.UpdateProblemMeasureWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProblemNoticeWithOptions(request *UpdateProblemNoticeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateProblemNoticeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemId)) {
		body["problemId"] = request.ProblemId
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemNotifyType)) {
		body["problemNotifyType"] = request.ProblemNotifyType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProblemNotice"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/problem/notify"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateProblemNoticeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProblemNotice(request *UpdateProblemNoticeRequest) (_result *UpdateProblemNoticeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateProblemNoticeResponse{}
	_body, _err := client.UpdateProblemNoticeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProblemTimelineWithOptions(request *UpdateProblemTimelineRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateProblemTimelineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.KeyNode)) {
		body["keyNode"] = request.KeyNode
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemId)) {
		body["problemId"] = request.ProblemId
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemTimelineId)) {
		body["problemTimelineId"] = request.ProblemTimelineId
	}

	if !tea.BoolValue(util.IsUnset(request.Time)) {
		body["time"] = request.Time
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProblemTimeline"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/problem/process/timeline/update"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateProblemTimelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProblemTimeline(request *UpdateProblemTimelineRequest) (_result *UpdateProblemTimelineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateProblemTimelineResponse{}
	_body, _err := client.UpdateProblemTimelineWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRichTextWithOptions(request *UpdateRichTextRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateRichTextResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		body["instanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.RichText)) {
		body["richText"] = request.RichText
	}

	if !tea.BoolValue(util.IsUnset(request.RichTextId)) {
		body["richTextId"] = request.RichTextId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRichText"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/rich/update"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRichTextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRichText(request *UpdateRichTextRequest) (_result *UpdateRichTextResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateRichTextResponse{}
	_body, _err := client.UpdateRichTextWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRouteRuleWithOptions(request *UpdateRouteRuleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateRouteRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssignObjectId)) {
		body["assignObjectId"] = request.AssignObjectId
	}

	if !tea.BoolValue(util.IsUnset(request.AssignObjectType)) {
		body["assignObjectType"] = request.AssignObjectType
	}

	if !tea.BoolValue(util.IsUnset(request.ChildRuleRelation)) {
		body["childRuleRelation"] = request.ChildRuleRelation
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CoverageProblemLevels)) {
		body["coverageProblemLevels"] = request.CoverageProblemLevels
	}

	if !tea.BoolValue(util.IsUnset(request.Effection)) {
		body["effection"] = request.Effection
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentLevel)) {
		body["incidentLevel"] = request.IncidentLevel
	}

	if !tea.BoolValue(util.IsUnset(request.MatchCount)) {
		body["matchCount"] = request.MatchCount
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyChannels)) {
		body["notifyChannels"] = request.NotifyChannels
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemEffectionServices)) {
		body["problemEffectionServices"] = request.ProblemEffectionServices
	}

	if !tea.BoolValue(util.IsUnset(request.ProblemLevelGroup)) {
		body["problemLevelGroup"] = request.ProblemLevelGroup
	}

	if !tea.BoolValue(util.IsUnset(request.RelatedServiceId)) {
		body["relatedServiceId"] = request.RelatedServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.RouteChildRules)) {
		body["routeChildRules"] = request.RouteChildRules
	}

	if !tea.BoolValue(util.IsUnset(request.RouteRuleId)) {
		body["routeRuleId"] = request.RouteRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.RouteType)) {
		body["routeType"] = request.RouteType
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		body["ruleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.TimeWindow)) {
		body["timeWindow"] = request.TimeWindow
	}

	if !tea.BoolValue(util.IsUnset(request.TimeWindowUnit)) {
		body["timeWindowUnit"] = request.TimeWindowUnit
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRouteRule"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/routeRule/edit"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRouteRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRouteRule(request *UpdateRouteRuleRequest) (_result *UpdateRouteRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateRouteRuleResponse{}
	_body, _err := client.UpdateRouteRuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateServiceWithOptions(request *UpdateServiceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceDescription)) {
		body["serviceDescription"] = request.ServiceDescription
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		body["serviceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		body["serviceName"] = request.ServiceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateService"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/services/update"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateService(request *UpdateServiceRequest) (_result *UpdateServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateServiceResponse{}
	_body, _err := client.UpdateServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateServiceGroupWithOptions(request *UpdateServiceGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateServiceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EnableWebhook)) {
		body["enableWebhook"] = request.EnableWebhook
	}

	if !tea.BoolValue(util.IsUnset(request.MonitorSourceTemplates)) {
		body["monitorSourceTemplates"] = request.MonitorSourceTemplates
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceGroupDescription)) {
		body["serviceGroupDescription"] = request.ServiceGroupDescription
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceGroupId)) {
		body["serviceGroupId"] = request.ServiceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceGroupName)) {
		body["serviceGroupName"] = request.ServiceGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["userIds"] = request.UserIds
	}

	if !tea.BoolValue(util.IsUnset(request.WebhookLink)) {
		body["webhookLink"] = request.WebhookLink
	}

	if !tea.BoolValue(util.IsUnset(request.WebhookType)) {
		body["webhookType"] = request.WebhookType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServiceGroup"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/services/group/modify"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateServiceGroup(request *UpdateServiceGroupRequest) (_result *UpdateServiceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateServiceGroupResponse{}
	_body, _err := client.UpdateServiceGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateServiceGroupSchedulingWithOptions(request *UpdateServiceGroupSchedulingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateServiceGroupSchedulingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.FastScheduling)) {
		body["fastScheduling"] = request.FastScheduling
	}

	if !tea.BoolValue(util.IsUnset(request.FineScheduling)) {
		body["fineScheduling"] = request.FineScheduling
	}

	if !tea.BoolValue(util.IsUnset(request.SchedulingWay)) {
		body["schedulingWay"] = request.SchedulingWay
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceGroupId)) {
		body["serviceGroupId"] = request.ServiceGroupId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServiceGroupScheduling"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/services/group/scheduling/update"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceGroupSchedulingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateServiceGroupScheduling(request *UpdateServiceGroupSchedulingRequest) (_result *UpdateServiceGroupSchedulingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateServiceGroupSchedulingResponse{}
	_body, _err := client.UpdateServiceGroupSchedulingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateServiceGroupSpecialDaySchedulingWithOptions(request *UpdateServiceGroupSpecialDaySchedulingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateServiceGroupSpecialDaySchedulingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.SchedulingDate)) {
		body["schedulingDate"] = request.SchedulingDate
	}

	if !tea.BoolValue(util.IsUnset(request.SchedulingSpecialDays)) {
		body["schedulingSpecialDays"] = request.SchedulingSpecialDays
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceGroupId)) {
		body["serviceGroupId"] = request.ServiceGroupId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServiceGroupSpecialDayScheduling"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/services/group/scheduling/updateSpecialDayScheduling"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceGroupSpecialDaySchedulingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateServiceGroupSpecialDayScheduling(request *UpdateServiceGroupSpecialDaySchedulingRequest) (_result *UpdateServiceGroupSpecialDaySchedulingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateServiceGroupSpecialDaySchedulingResponse{}
	_body, _err := client.UpdateServiceGroupSpecialDaySchedulingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSubscriptionWithOptions(request *UpdateSubscriptionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateSubscriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExpiredType)) {
		body["expiredType"] = request.ExpiredType
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyObjectList)) {
		body["notifyObjectList"] = request.NotifyObjectList
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyObjectType)) {
		body["notifyObjectType"] = request.NotifyObjectType
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyStrategyList)) {
		body["notifyStrategyList"] = request.NotifyStrategyList
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		body["period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		body["scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeObjectList)) {
		body["scopeObjectList"] = request.ScopeObjectList
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriptionId)) {
		body["subscriptionId"] = request.SubscriptionId
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriptionTitle)) {
		body["subscriptionTitle"] = request.SubscriptionTitle
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSubscription"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/notify/subscription/update"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSubscriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSubscription(request *UpdateSubscriptionRequest) (_result *UpdateSubscriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateSubscriptionResponse{}
	_body, _err := client.UpdateSubscriptionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateUserWithOptions(request *UpdateUserRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		body["email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		body["phone"] = request.Phone
	}

	if !tea.BoolValue(util.IsUnset(request.RamId)) {
		body["ramId"] = request.RamId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		body["username"] = request.Username
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUser"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/user/update"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateUser(request *UpdateUserRequest) (_result *UpdateUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateUserResponse{}
	_body, _err := client.UpdateUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateUserGuideStatusWithOptions(request *UpdateUserGuideStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateUserGuideStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.GuideAction)) {
		body["guideAction"] = request.GuideAction
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUserGuideStatus"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/user/update/guide/status"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserGuideStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateUserGuideStatus(request *UpdateUserGuideStatusRequest) (_result *UpdateUserGuideStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateUserGuideStatusResponse{}
	_body, _err := client.UpdateUserGuideStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyRouteRuleWithOptions(request *VerifyRouteRuleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *VerifyRouteRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RouteRuleId)) {
		body["routeRuleId"] = request.RouteRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.TestSourceEvents)) {
		body["testSourceEvents"] = request.TestSourceEvents
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyRouteRule"),
		Version:     tea.String("2021-04-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/routeRule/verify"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyRouteRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyRouteRule(request *VerifyRouteRuleRequest) (_result *VerifyRouteRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &VerifyRouteRuleResponse{}
	_body, _err := client.VerifyRouteRuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
