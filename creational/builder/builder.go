package main

type NotificationBuilder struct {
	Title    string
	Subtitle string
	Message  string
	Image    string
	Priority int
}

func newNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

func (nb *NotificationBuilder) SetTitle(title string) {
	nb.Title = title
}

func (nb *NotificationBuilder) SetSubTitle(subtitle string) {
	nb.Subtitle = subtitle
}

func (nb *NotificationBuilder) SetMessage(message string) {
	nb.Message = message
}

func (nb *NotificationBuilder) SetImage(img string) {
	nb.Image = img
}

func (nb *NotificationBuilder) SetPriority(prt int) {
	nb.Priority = prt
}

func (nb *NotificationBuilder) Build() (*Notification, error) {
	return &Notification{
		title:    nb.Title,
		subtitle: nb.Subtitle,
		message:  nb.Message,
		image:    nb.Image,
		priority: nb.Priority,
	}, nil
	return nil, nil
}
