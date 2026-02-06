package main

import (
	"context"
	"log"

	sdk "github.com/apito-io/go-apito-plugin-sdk"
)

func main() {
	log.Printf("🎯 [hc-push-notification-plugin] Starting Push Notification plugin...")
	plugin := sdk.Init("hc-push-notification-plugin", "1.0.0", "apito-plugin-key")
	statusType := sdk.NewObjectType("PushStatus", "Push notification plugin status").
		AddStringField("status", "Plugin status", false).
		AddStringField("version", "Plugin version", false).
		Build()
	plugin.RegisterQuery("getPushStatus",
		sdk.ComplexObjectField("Get push notification plugin status", statusType),
		func(ctx context.Context, rawArgs map[string]interface{}) (interface{}, error) {
			return map[string]interface{}{"status": "ready", "version": "1.0.0"}, nil
		})
	log.Printf("🚀 [hc-push-notification-plugin] Plugin ready")
	plugin.Serve()
}
