package jobctl

import (
	"context"
	"fmt"
	"time"

	"github.com/obaraelijah/secureproc/pkg/client/jobmanager"
	"github.com/spf13/cobra"
)

const (
	// With 1000 jobs, list takes ~200ms (~10ms more than the "short")
	//      2000 jobs, list takes ~210ms (~10ms more than 1000 jobs)
	//      3000 jobs, list takes ~220ms (~10ms more than 1000 jobs)
	// Each job adds about 0.01 ms
	// Assume a maximum of 100k jobs, that'd be ~1 second
	listOperationTimeout = shortOperationTimeout + (1 * time.Second)
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List job",
	Long:    "List jobs managed by the JobManager",
	Example: "jobctl list",
	RunE:    list,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func list(cmd *cobra.Command, _ []string) error {
	ctx, cancel := context.WithTimeout(cmd.Context(), listOperationTimeout)
	defer cancel()

	c, err := jobmanager.NewClient(argUserID, argServerHostPort)
	if err != nil {
		return err
	}
	defer c.Close()

	jobList, err := c.List(ctx)
	if err != nil {
		return err
	}

	if len(jobList) == 0 {
		fmt.Println("There are no jobs")
		return nil
	}

	renderJobStatusList(jobList)

	return nil
}
