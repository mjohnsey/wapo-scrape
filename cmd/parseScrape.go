// Copyright © 2019 Michael Johnsey <mjohnsey@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	wapo "github.com/mjohnsey/wapo-scrape/pkg"
	"github.com/spf13/cobra"
)

// parseScrapeCmd represents the parseScrape command
var parseScrapeCmd = &cobra.Command{
	Use:   "parseScrape [FILE]",
	Short: "Parse a scrape json file",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Requires to pass in the json file location")
		}
		fileLocation := args[0]
		if _, err := os.Stat(fileLocation); !os.IsNotExist(err) {
			return err
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fileLocation := args[0]
		scrape, err := wapo.WashingtonPostScrape{}.FromJsonFile(fileLocation)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(scrape.Stats())
	},
}

func init() {
	rootCmd.AddCommand(parseScrapeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// parseScrapeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// parseScrapeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
