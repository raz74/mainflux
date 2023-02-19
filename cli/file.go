package cli

import (
	"encoding/json"
	mfxsdk "github.com/mainflux/mainflux/pkg/sdk/go"
	"github.com/spf13/cobra"
)

var cmdFile = []cobra.Command{
	{
		Use:   "upload <JSON_string> <file> <user_auth_token>",
		Short: "Upload file",
		Long:  `Upload new file`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 3 {
				logUsage(cmd.Use)
				return
			}
			var file mfxsdk.File
			if err := json.Unmarshal([]byte(args[0]), &file); err != nil {
				logError(err)
				return
			}
			name := args[1]
			token := args[2]

			id, err := sdk.UploadFile(file, name, token)
			if err != nil {
				logError(err)
				return
			}

			logCreated(id)
		},
	},
}

// NewFileCmd returns files command.
func NewFileCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "Files [upload]",
		Short: "Files management",
		Long:  `Files management: upload files"`,
	}

	for i := range cmdFile {
		cmd.AddCommand(&cmdFile[i])
	}

	return &cmd
}
