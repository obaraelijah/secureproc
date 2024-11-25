package jobctl

import (
	"context"
	"errors"
	"os"
	"strconv"

	"github.com/obaraelijah/secureproc/pkg/client/jobmanager"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var queryCmd = &cobra.Command{
	Use:     "query",
	Short:   "Query job state",
	Long:    "Query the state of a job managed by JobManager",
	Example: "jobctl query ba90b623-3dae-4bdd-8b96-c1ea4a999c44",
	RunE:    query,
}

func init() {
	rootCmd.AddCommand(queryCmd)
}

func query(cmd *cobra.Command, jobIDs []string) error {
	if len(jobIDs) == 0 {
		return errors.New("no jobs specified")
	}

	c, err := jobmanager.NewClient(argUserID, argServerHostPort)
	if err != nil {
		return err
	}
	defer c.Close()
	var lastError error

	jobStatusList := make([]*jobmanager.JobStatus, 0, len(jobIDs))

	for _, jobID := range jobIDs {
		func() {
			ctx, cancel := context.WithTimeout(cmd.Context(), shortOperationTimeout)
			defer cancel()

			status, err := c.Query(ctx, jobID)
			if err != nil {
				lastError = err
				return
			}

			jobStatusList = append(jobStatusList, status)
		}()
	}

	renderJobStatusList(jobStatusList)

	return lastError
}

func renderJobStatusList(jobStatus []*jobmanager.JobStatus) {
	isAdmin := argUserID == "administrator"
	header := []string{"Owner", "Name", "ID", "Running", "Pid", "Exit Code", "Signal", "Error"}

	if !isAdmin {
		header = header[1:]
	}

	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader(header)

	for _, js := range jobStatus {
		runErr := ""
		if js.RunError != nil {
			runErr = js.RunError.Error()
		}

		sigStr := ""
		if js.SignalNum > 0 {
			sigStr = js.SignalNum.String()
		}

		exitCode := ""
		if js.ExitCode >= 0 {
			exitCode = strconv.FormatInt(int64(js.ExitCode), 10)
		}

		pid := ""
		if js.Pid > 0 {
			pid = strconv.FormatInt(int64(js.Pid), 10)
		}

		columns := make([]string, 0, 8)

		if isAdmin {
			columns = append(columns, js.Owner)
		}

		columns = append(columns, js.Name)
		columns = append(columns, js.ID)
		columns = append(columns, strconv.FormatBool(js.Running))
		columns = append(columns, pid)
		columns = append(columns, exitCode)
		columns = append(columns, sigStr)
		columns = append(columns, runErr)

		table.Append(columns)
	}

	table.Render()
}
