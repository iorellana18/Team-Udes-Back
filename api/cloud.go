package api

import (
	"github.com/NaySoftware/go-fcm"
)

const (
	serverKey = "AAAAwPFQ8Qw:APA91bF_LFLBtOIxVpZA8naarJBS8e5tolbbed2UrN9NsO52bB5FuJFSXLKGCcsLz7sogwRf7vXhXI6oeIkeInkh-4YrVQclhNl7LBDKFg3EacGT3HtTcYmzmV2slcAsjRC0HWx1FYQI"
)

func SendNotification(appToken string, title string, notification string) error {

	data := map[string]string{
		"type":  "Voluntarios",
		"id":    "1",
		"title": title,
		"body":  notification,
	}

	ids := []string{
		appToken,
	}

	c := fcm.NewFcmClient(serverKey)
	c.NewFcmRegIdsMsg(ids, data)

	status, err := c.Send()

	if err == nil {
		status.PrintResults()
	}

	return err
}
