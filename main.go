package main

import (
	"emojify-ipv6/emojidb"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cobra"
	"os"
)

var gitTag string

type EmojiDataBase struct {
	Emoji          string   `json:"emoji"`
	Description    string   `json:"description"`
	Category       string   `json:"category"`
	Aliases        []string `json:"aliases"`
	Tags           []string `json:"tags"`
	UnicodeVersion string   `json:"unicode_version"`
	IOSVersion     string   `json:"ios_version"`
}

var emojiData []EmojiDataBase
var inputData string

var rootCmd = &cobra.Command{
	Use:   "emojify-ipv6",
	Short: "Convert an array of emojis to an IPv6 address and vice versa",
	Long: `The IPv6 address is expressed as UTF-8 encoded data of emojis.
This provides higher readability compared to the existing IPv6 notation.
 
  😀😀😀😀 <-> f09f:9880:f09f:9880:f09f:9880:f09f:9880
	`,
	Example: `
	# Convert emojis to IPv6. Emoji-cli("https://github.com/b4b4r07/emoji-cli") is strongly required to run this example.
  echo 😀😀😀😀 | emojify-ipv6 

  # Reverse IPv6 to emojis 
  emojify-ipv6 -r "f09f:9880:f09f:9880:f09f:9880:f09f:9880"
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			inputData = args[0]
		}
		return runCommand()
	},
}

var flags struct {
	reverse string
	version bool
	list    bool
}

var flagsName = struct {
	reverse, reverseShort string
	version, versionShort string
	list, listShort       string
}{
	"reverse", "r",
	"version", "v",
	"list", "l",
}

func main() {
	rootCmd.Flags().StringVarP(
		&flags.reverse,
		flagsName.reverse,
		flagsName.reverseShort,
		"", "IPv6 address")
	rootCmd.PersistentFlags().BoolVarP(
		&flags.version,
		flagsName.version,
		flagsName.versionShort,
		false, "version number")
	rootCmd.PersistentFlags().BoolVarP(
		&flags.list,
		flagsName.list,
		flagsName.listShort,
		false, "show the list of supported emojis")

	bytejson, _ := emojidb.Asset("emojidb/emoji.json")

	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	if err := json.Unmarshal(bytejson, &emojiData); err != nil {
		panic(err)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
