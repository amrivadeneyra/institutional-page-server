package httpresponses

const (
	LoginError = iota + 1
	DeletionNotAllowed
	LogOperationNotAllowed
	UpdateNotAllowed
	OutOfPlan
	InvalidEmail
	ServiceStatusChangeDenied
	ServiceCustomFieldChangeDenied
	ServiceSystemDuplicated
)
