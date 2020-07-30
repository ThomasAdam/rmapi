package main

import (
	"os"

	_ "github.com/juruen/rmapi/annotations/license"
	"github.com/juruen/rmapi/api"
	"github.com/juruen/rmapi/log"
	"github.com/juruen/rmapi/shell"
)

const AUTH_RETRIES = 3

func run_shell(ctx *api.ApiCtx) {
	err := shell.RunShell(ctx)

	if err != nil {
		log.Error.Println("Error: ", err)
		os.Exit(1)
	}
}

func main() {
	log.InitLog()

	var ctx *api.ApiCtx
	var err error
	for i := 0; i < AUTH_RETRIES; i++ {
		ctx, err = api.CreateApiCtx(api.AuthHttpCtx())

		if err != nil {
			log.Trace.Println(err)
		}
	}

	if ctx == nil {
		log.Error.Fatal("failed to build documents tree, last error: ", err)
	}

	run_shell(ctx)
}
