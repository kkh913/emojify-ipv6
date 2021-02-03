# emojify-ipv6
Convert emojis to an IPv6 address and vice versa

## Warning ⚠️

This emoji IPv6 address system should **never** be used. This is simply my personal project to develop Golang's development skills. See [Limitations] below for a detailed explanation.


## Installation

``` bash
$ make && make install 
```
The binary `emojify-ipv6` will be located in `$GOPATH/bin`.


## Dependencies

- [emoji-cli]: Emoji completion on the command line 
- [emoji.json]: emoji data from gemoji (official?)
- [go-bindata]: inject json into executable
- [json-iterator]: High-Performance JSON Tool
- [cobra]: Modern CLI

## Example 

`emojify-ipv6` uses 4 emoticons to represent IPv6 addresses.
``` bash
emojify-ipv6 😀😀😀😀
# or
echo 😀😀😀😀 | emojify-ipv6
```
output:
> f09f:9880:f09f:9880:f09f:9880:f09f:9880

And vice versa with flag `-r, --reverse`
``` bash
emojify-ipv6 -r f09f:9880:f09f:92af:f09f:918d:f09f:9890
# or 
echo f09f:9880:f09f:92af:f09f:918d:f09f:9890 | emojify-ipv6 -r - 
```
output:
> 😀💯👍😐

## Demo 

For [emoji-cli] to work as shown in the demo, you need to add the follwing to `.zshrc`: 
```
[ -f ~/emoji-cli/emoji-cli.zsh ] && source ~/emoji-cli/emoji-cli.zsh
export EMOJI_CLI_FILTER="fzf --height 40%"
export EMOJI_CLI_USE_EMOJI=1
```

## Emoji 

An [emoji] is a character in the form of a picture, which is different from [emoticon] which is a combination of characters, and even has a different etymology. (Japanese ↔ English) The emojis have spread greatly due to smartphones and SNS, and are actively occupying part of the Unicode map. As of this writing, the total number of emojis in v13.1 is [3,521][1].

## Motivation 

[IPv6] is 16 bytes (128 bits) in size and can be easily represented by 4 emojis with a UTF-8 encoding size of 4 bytes. I am not the first pioneer to express IP addresses in a combination of emojis. [Their address system][2], which was the only one found in my research, can represent IPv4 as well as IPv6. It maps 256 emojis to 1 byte, so it can represent any IP address. However, there is arguable over the reasons for mapping numbers and emojis. In contrast, my method uses 2 byte UTF-8 encoded emojis. Thus, we can check the IPv6 address without a additional mapping table. Although it is very limited to 1,053 which is the number of emojis that can be used in this address system, it offers 286.26 times larger address space than the IPv4 address space.

## Limitations 

### Address system ignoring RFC 

This address system cannot represent any of the IPv6 address formats - Global Unicast Address, Link-local Address, and Unique Local Addresses - even if it excludes the addresses, such as loopback address, Solicited-node multicast address, and multicast address.

### Sparse address space 

It is 2<sup>32</sup>, which is the maximum number of endpoints that can be represented in 2 bytes - equivalent to the IPv4 address space. However, this address system inefficiently uses only 1,053 endpoints in this large address space. It is also inefficient attempt to emojify a 64-bit interface ID of a Unique Local Address. 

### Configure network subnets with only 4 prefixes? Well...

There are only 4 network prefixes in this address system: `/32`,`/64`,`/96`,`/128`. Suppose this address system is available for private use. Since address resources are much richer compared to IPv4, network operators will be happy to assign addresses. However, this forces you to allocate `/96` subnets with 1,053 addresses for teams of less than 10 people, although the dev teams might be very welcome. 

## Credits 

Especially thanks to [@b4b4r07][b4b4r07]. Thanks to his work, [emoji-cli] and [zplug], I am able to reduce the time it takes to achieve the goal of this project. 

## License 

MIT

[emoji]: https://en.wikipedia.org/wiki/Emoji
[emoticon]: https://en.wikipedia.org/wiki/Emoticon
[Limitations]: https://github.com/kkh913/emojify-ipv6#limitations 
[emoji-cli]: https://github.com/b4b4r07/emoji-cli
[emoji.json]: https://github.com/github/gemoji/blob/master/db/emoji.json
[go-bindata]: https://github.com/go-bindata/go-bindata
[json-iterator]: https://github.com/json-iterator/go
[cobra]: https://github.com/spf13/cobra
[b4b4r07]: https://github.com/b4b4r07
[zplug]: https://github.com/zplug/zplug
[1]: https://www.unicode.org/emoji/charts/emoji-counts.html
[IPv6]: https://en.wikipedia.org/wiki/IPv6_address
[2]: https://www.6connect.com/resources/how-to-view-ip-addresses-as-emojis/

