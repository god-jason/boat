package db

import "github.com/god-jason/boat/boot"

func init() {
	boot.Register("database", &boot.Task{
		Startup:  Startup,
		Shutdown: Shutdown,
		Depends:  []string{"config"},
	})
}
