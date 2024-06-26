package statscmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/sagan/ptool/cmd"
	"github.com/sagan/ptool/config"
	"github.com/sagan/ptool/stats"
)

var command = &cobra.Command{
	Use:         "stats [clients]... [flags]",
	Annotations: map[string]string{"cobra-prompt-dynamic-suggestions": "stats"},
	Short:       "Show client brushing traffic statistics.",
	Long: `Show client brushing traffic statistics.
Only torrents added by ptool (of this machine) will be counted.
The traffic info of a torrent will ONLY be recorded when it's been DELETED from the client.
To use this command, enable the statistics feature by adding the "brushEnableStats = true"
line to ptool.toml config file.`,
	RunE: statscmd,
}

var (
	statsFilename = ""
)

func init() {
	command.Flags().StringVarP(&statsFilename, "stats-file", "", "",
		"Manually specify stats file ("+config.STATS_FILENAME+") path")
	cmd.RootCmd.AddCommand(command)
}

func statscmd(cmd *cobra.Command, args []string) error {
	clientnames := args
	if !config.Get().BrushEnableStats {
		return fmt.Errorf("statistics feature is NOT enabled currently. " +
			"To enable it, add the \"brushEnableStats = true\" line to the top of ptool.toml config file. " +
			"It will use the \"ptool_stats.txt\" (in the same dir of ptool.toml file) as the statistics data file")
	}
	if statsFilename == "" {
		statsFilename = filepath.Join(config.ConfigDir, config.STATS_FILENAME)
	}
	statDb, err := stats.NewDb(statsFilename)
	if err != nil {
		return fmt.Errorf("failed to create stats db: %w", err)
	}
	if len(clientnames) == 0 {
		statDb.ShowTrafficStats("")
		return nil
	}

	doneFlag := map[string]bool{}
	for i, clientname := range clientnames {
		if clientname == "_" || doneFlag[clientname] {
			continue
		}
		doneFlag[clientname] = true
		if i > 0 {
			fmt.Printf("\n")
		}
		statDb.ShowTrafficStats(clientname)
	}
	return nil
}
