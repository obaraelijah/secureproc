package jobctl

import (
	"context"

	"github.com/obaraelijah/secureproc/pkg/client/jobmanager"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:     "stop",
	Short:   "Stop a job",
	Long:    "Stop a job managed by the JobManger.  If the job is not running, this has no effect.",
	Example: "jobctl stop 8de11b74-5cd9-4769-b40d-53de13faf77f",
	RunE:    stop,
}

func init() {
	rootCmd.AddCommand(stopCmd)
}

func stop(cmd *cobra.Command, jobIDs []string) error {

	c, err := jobmanager.NewClient(argUserID, argServerHostPort)
	if err != nil {
		return err
	}
	defer c.Close()

	for _, jobID := range jobIDs {
		err = func() error {
			ctx, cancel := context.WithTimeout(cmd.Context(), shortOperationTimeout)
			defer cancel()

			return c.Stop(ctx, jobID)
		}()

		if err != nil {
			return err
		}
	}

	return nil
}
