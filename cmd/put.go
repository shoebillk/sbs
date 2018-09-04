// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/shoebillk/sbs/blob"
	"github.com/spf13/cobra"
)

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put",
	Short: "put a file to sbs server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("put called")

		fmt.Printf("%#v\n", args)

		log.Println(cmd.Flags().GetString("host"))
		log.Println(cmd.Flags().GetInt("port"))

		f, err := os.Open(args[0])
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		reader := bufio.NewReader(f)

		chunkContentSize := 0x1000
		chunk := blob.Chunk{
			Id:      args[0],
			Content: make([]byte, chunkContentSize),
		}

		for {
			n, err := reader.Read(chunk.Content)

			if err == io.EOF {
				log.Printf("Done to read")
				break
			} else if err != nil {
				log.Fatal(err)
				break
			}

			chunk.Content = chunk.Content[:n]
			log.Printf("Read %d", n)

			// TODO send chunk to server
		}

	},
}

func init() {
	rootCmd.AddCommand(putCmd)

	putCmd.Flags().StringP("host", "s", DefaultHost, "Host string of server")
	putCmd.Flags().IntP("port", "p", DefaultPort, "Port number of server")
}
