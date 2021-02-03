package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net"
	"os"
	"strings"
)

func runCommand() error {
	if flags.version {
		fmt.Println(gitTag)
		return nil
	} else if flags.list {
		for _, item := range emojiData {
			if len(item.Emoji) == 4 {
				fmt.Printf("%x%x:%x%x %s\n", item.Emoji[0], item.Emoji[1], item.Emoji[2], item.Emoji[3], item.Emoji)
			}
		}
		return nil
	} else if isInputFromPipe() {
		if flags.reverse != "" {
			return toEmoji(os.Stdin, os.Stdout)
		} else {
			return toIPv6(os.Stdin, os.Stdout)
		}
	} else if flags.reverse != "" {
		return toEmoji(strings.NewReader(flags.reverse), os.Stdout)
	} else if inputData != "" {
		return toIPv6(strings.NewReader(inputData), os.Stdout)
	}
	return errors.New("no valid operations.")
}

func isInputFromPipe() bool {
	fi, _ := os.Stdin.Stat()
	return fi.Mode()&os.ModeCharDevice == 0
}

func processConvertEmoji(s string) (string, error) {

	var rawData net.IP
	rawData = net.ParseIP(s)

	if rawData == nil {
		return "", errors.New("Invaild IPv6 address")
	}

	Contains := func(s string) bool {
		for _, item := range emojiData {
			if item.Emoji == s {
				return true
			}
		}
		return false
	}

	for i := 0; i < 4; i++ {
		emojiValue := rawData[(i * 4):(i*4 + 4)]
		if !Contains(string(emojiValue)) {
			return "", errors.New("There are no convertible emojis.")
		}
	}

	return string(rawData), nil
}

func processConvertIPv6(s string) (string, error) {

	if len(s) != 16 {
		return "", errors.New("Invalid length of emojis.")
	}

	var ipv6_addr net.IP

	ipv6_addr = []byte(s)

	return ipv6_addr.String(), nil
}

func toIPv6(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(bufio.NewReader(r))
	for scanner.Scan() {

		var emojis string
		var e error

		emojis, e = processConvertIPv6(scanner.Text())
		if e != nil {
			return e
		}

		_, e = fmt.Fprintln(
			w, emojis)
		if e != nil {
			return e
		}
	}
	return nil
}

func toEmoji(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(bufio.NewReader(r))
	for scanner.Scan() {

		var ipv6_addr string
		var e error

		ipv6_addr, e = processConvertEmoji(scanner.Text())
		if e != nil {
			return e
		}

		_, e = fmt.Fprintln(
			w, ipv6_addr)
		if e != nil {
			return e
		}
	}
	return nil
}

func fileExists(filepath string) bool {
	info, e := os.Stat(filepath)
	if os.IsNotExist(e) {
		return false
	}
	return !info.IsDir()
}
