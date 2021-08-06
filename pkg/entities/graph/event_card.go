package graph

type DBEventCard struct {
	IDGraph

	DataUI
	DataEventCard
}

type DataEvent struct {
	Platform string

	OwnerType string
	OwnerId   string

	ResourceType string
	ResourceId   string

	ContextType string
	ContextId   string

	InitiatorType string
	InitiatorId   string

	StaticType string
	StaticId   string
}

type DataEventCard struct {
	HttpVote uint64

	//Placement
	TargetId  uint64
	SlidePort string

	DataEvent
}
