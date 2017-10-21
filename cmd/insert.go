// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/chentex/go-sif/service"
	"github.com/chentex/go-sif/service/io"
	"github.com/spf13/cobra"
)

// insertCmd represents the insert command
var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "gosif insert will insert a string into a file of your choice",
	Long: `gosif insert will insert a string into a file of your choice
Line count starts at 1.

You can specify a specific line number. For example:

$ gosif insert -f testfile -l 10 string to insert in the file

this command will, when executed correctly, insert the string in the file 'testfile' in line 10

$ gosif insert -f testfile string to insert in the file

this command will, when executed correctly, insert the string at the end of file 'testfile' in line a new line`,
	RunE: func(cmd *cobra.Command, args []string) error {
		insertService := service.NewInsertService()
		fm := io.NewFileManager()
		insertService.SetManager(fm)
		err := insertService.Insert(file, line, text)
		return err
	},
}

func init() {
	RootCmd.AddCommand(insertCmd)
}
