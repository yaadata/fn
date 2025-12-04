package forgedomain

import (
	"time"
)

type notificationGetOptions struct {
	state NotificationStatus
	since time.Time
}

type NotificationGetOptions func(*notificationGetOptions)

func WithNotificationState(state NotificationStatus) NotificationGetOptions {
	return func(o *notificationGetOptions) {
		o.state = state
	}
}

func WithNotificationSince(since time.Time) NotificationGetOptions {
	return func(o *notificationGetOptions) {
		o.since = since
	}
}
