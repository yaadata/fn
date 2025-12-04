package forgedomain

import . "github.com/yaadata/optionsgo"

type Notification struct {
	Context    NotificationContext
	Id         string
	Kind       Option[Value[NotificationKind]]
	Repository Option[Repository]
	Trigger    Value[NotificationTrigger]
	Status     Value[NotificationStatus]
	Viewed     bool
}
