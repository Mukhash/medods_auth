package main

import (
	"go.uber.org/zap"
)

func main() {
	//ctx := context.Background()
	//cfg := config.DefaultConfig()
	logger := zap.NewExample()

	defer logger.Sync()

}
