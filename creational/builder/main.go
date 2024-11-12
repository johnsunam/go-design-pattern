package main

import "fmt"

type Notification struct {
	title    string
	subtitle string
	message  string
	image    string
	priority int
}

func main() {
	nbldr := newNotificationBuilder()
	nbldr.SetTitle("New Notification")
	nbldr.SetSubTitle("Urgent Notification")
	nbldr.SetMessage("This is test notificaiton")
	nbldr.SetImage("notification.png")
	nbldr.SetPriority(4)
	nf, err := nbldr.Build()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\n notification %+v ", nf)
	}

}
