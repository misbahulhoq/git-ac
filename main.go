/*
Copyright © 2026 Md Mezbah Uddin extraordinarymisbah@gmail.com
*/

package main

import (
	"github.com/joho/godotenv"
	"github.com/misbahulhoq/gcm/cmd"
)

func main() {
	godotenv.Load()
	cmd.Execute()
}
