package cmd

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/blcksec/stripe-cli/pkg/plugins"
)

func createPluginCmd() *pluginTemplateCmd {
	plugin := plugins.Plugin{
		Shortname:        "test",
		Shortdesc:        "test your stuff",
		Binary:           "stripe-cli-test",
		MagicCookieValue: "magic",
		Releases: []plugins.Release{{
			Arch:    "amd64",
			OS:      "darwin",
			Version: "0.0.1",
			Sum:     "c53a98c3fa63563227eb8b5601acedb5e0e70fed2e1d52a5918a17ac755f17f7",
		}},
	}

	pluginCmd := newPluginTemplateCmd(&Config, &plugin)

	return pluginCmd
}

// TestFlagsArePassedAsArgs ensures that the plugin is passing all args and flags as expected.
// This is a complex dance between the CLI itself and the plugin, so the flags come from
// two different sources as a result. This test is here to catch any non-obvious regressions
func TestFlagsArePassedAsArgs(t *testing.T) {
	pluginCmd := createPluginCmd()
	rootCmd.AddCommand(pluginCmd.cmd)

	Execute(context.Background())

	// temp override for the os.Args so that the pluginCmd can use them
	oldArgs := os.Args
	os.Args = []string{"stripe", "test", "testarg", "--log-level=info"}
	defer func() { os.Args = oldArgs }()

	rootCmd.SetArgs([]string{"test", "testarg", "--log-level=info"})
	executeCommandC(rootCmd, "test", "testarg", "--log-level=info")

	require.Equal(t, 2, len(pluginCmd.ParsedArgs))
	require.Equal(t, "testarg --log-level=info", strings.Join(pluginCmd.ParsedArgs, " "))
}
