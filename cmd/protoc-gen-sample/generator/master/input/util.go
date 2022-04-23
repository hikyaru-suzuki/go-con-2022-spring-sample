package input

var AdminAccessorSet = NewAccessorSet(
	AccessorType_All,
	AccessorType_OnlyAdmin,
	AccessorType_AdminAndServer,
	AccessorType_AdminAndClient,
)

var ServerAccessorSet = NewAccessorSet(
	AccessorType_All,
	AccessorType_OnlyServer,
	AccessorType_AdminAndServer,
	AccessorType_ServerAndClient,
)

var ClientAccessorSet = NewAccessorSet(
	AccessorType_All,
	AccessorType_OnlyClient,
	AccessorType_AdminAndClient,
	AccessorType_ServerAndClient,
)
