package app

import (
	"regexp"
	"testing"
)

func _TestCommands(t *testing.T, regexp *regexp.Regexp, commands []string) {
	for _, cmd := range commands {
		if !regexp.MatchString(cmd) {
			t.Errorf("`%v` should match with `%v`, but not", cmd, regexp)
		}
	}
}

func TestBuildSetBotID(t *testing.T) {
	app := &App{}
	app.SetBotID("rabot-test")
	_TestCommands(t, app.Commands.List, []string{
		"<@rabot-test>  list   container  ",
		"<@rabot-test>  list   containers  ",
		"<@rabot-test>  ls  containers ",
		"<@rabot-test>  ls ",
	})
	_TestCommands(t, app.Commands.Start, []string{
		"<@rabot-test>  start  record    ALPHA-STATION for  12 min ",
		"<@rabot-test>  start  recording    ALPHA-STATION   1minute ",
	})
	_TestCommands(t, app.Commands.Remove, []string{
		"<@rabot-test>  remove  container  foo  ",
		"<@rabot-test>  rm  container  foo  ",
		"<@rabot-test>  rm   foo  ",
	})
	_TestCommands(t, app.Commands.Ping, []string{
		"<@rabot-test>  ping  ",
	})
}