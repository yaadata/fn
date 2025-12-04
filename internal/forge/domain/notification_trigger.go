package forgedomain

type NotificationTrigger string

const (
	ForgeNotificationTriggerAssigned        NotificationTrigger = "assigned"
	ForgeNotificationTriggerReviewRequested NotificationTrigger = "review_requested"
	ForgeNotificationTriggerMentioned       NotificationTrigger = "mentioned"
	ForgeNotificationTriggerOther           NotificationTrigger = "other"
)
