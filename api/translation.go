package api

import (
	"google.golang.org/api/option"

	"cloud.google.com/go/translate"
	"golang.org/x/net/context"
	"golang.org/x/text/language"
)

var apiKeyTranslate = "AIzaSyDpsV4XaNs4l8vD2lVdPR9YVmElTR7-wuU"

func translateText(text string) (string, error) {
	ctx := context.Background()

	if client, err := translate.NewClient(ctx, option.WithAPIKey(apiKeyTranslate)); err != nil {
		return "", err
	} else {
		if translations, err := client.Translate(ctx,
			[]string{text}, language.Spanish,
			&translate.Options{
				Source: language.English,
				Format: translate.Text,
			}); err != nil {
			return "", err
		} else {
			return translations[0].Text, nil
		}
	}
}
