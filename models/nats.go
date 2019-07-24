package models

type NatsSubscription struct {
	ClientId     string `json:"client_id"`
	Inbox        string `json:"inbox"`
	AckInbox     string `json:"ack_inbox"`
	QueueName    string `json:"queue_name"`
	IsDurable    bool   `json:"is_durable"`
	IsOffline    bool   `json:"is_offline"`
	MaxInflight  int    `json:"max_inflight"`
	AckWait      int    `json:"ack_wait"`
	LastSent     int    `json:"last_sent"`
	PendingCount int    `json:"pending_count"`
	IsStalled    bool   `json:"is_stalled"`
}

// IsHealthy returns true if the subscription is concidered healthy. This is
// comprised by the other facts that we know about the subscription.
func (sub NatsSubscription) IsHealthy() bool {
	return !sub.IsOffline && !sub.IsStalled && sub.PendingCount == 0
}

type NatsChannel struct {
	Name          string             `json:"name"`
	MessagesCount int                `json:"msgs"`
	BytesCount    int                `json:"bytes"`
	FirstSequence int                `json:"first_seq"`
	LastSequence  int                `json:"last_seq"`
	Subscriptions []NatsSubscription `json:"subscriptions"`
}

// Color returns a string with the color code for the status of a given channel.
// This is based on the IsHealthy status from the corresponding subscriptions.
func (ch NatsChannel) Color() string {
	color := "green"

	if len(ch.Subscriptions) == 0 {
		color = "orange"
	}

	for _, sub := range ch.Subscriptions {
		if !sub.IsHealthy() {
			color = "red"
		}
	}

	return color
}

type NatsChannels struct {
	ClusterId string        `json:"cluster_id"`
	ServerId  string        `json:"server_id"`
	Timestamp string        `json:"now"`
	Offset    int           `json:"offset"`
	Limit     int           `json:"limit"`
	Count     int           `json:"count"`
	Total     int           `json:"total"`
	Channels  []NatsChannel `json:"channels"`
}
