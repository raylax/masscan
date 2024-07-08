package main

import (
	"github.com/raylax/masscan/pkg/runner"
	"github.com/zan8in/gologger"
)

func main() {
	opts, err := runner.NewOptions()
	if err != nil {
		gologger.Error().Msg(err.Error())
		return
	}

	runner, err := runner.NewRunner(opts)
	if err != nil {
		gologger.Error().Msg(err.Error())
		return
	}

	ts, err := runner.GetTargets()
	if err != nil {
		gologger.Error().Msg(err.Error())
		return
	}

	for t := range ts {
		gologger.Debug().Msg(t)
	}
}
