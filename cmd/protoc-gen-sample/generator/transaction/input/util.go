package input

var ServerMessageAccessorSet = NewMessageAccessorSet(
	MessageAccessorType_OnlyServer,
	MessageAccessorType_ServerAndClient,
	MessageAccessorType_ServerAndClientWithCommonResponse,
)

var ClientMessageAccessorSet = NewMessageAccessorSet(
	MessageAccessorType_OnlyClient,
	MessageAccessorType_OnlyClientWithCommonResponse,
	MessageAccessorType_ServerAndClient,
	MessageAccessorType_ServerAndClientWithCommonResponse,
)

var ClientMessageCommonResponseAccessorSet = NewMessageAccessorSet(
	MessageAccessorType_OnlyClientWithCommonResponse,
	MessageAccessorType_ServerAndClientWithCommonResponse,
)

var ServerFieldAccessorSet = NewFieldAccessorSet(
	FieldAccessorType_All,
	FieldAccessorType_OnlyServer,
)

var ClientFieldAccessorSet = NewFieldAccessorSet(
	FieldAccessorType_All,
	FieldAccessorType_OnlyClient,
)
