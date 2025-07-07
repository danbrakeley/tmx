# tmx

## Overview

Go package for parsing [Tiled](https://www.mapeditor.org/) [TMX map files](https://doc.mapeditor.org/en/stable/reference/tmx-map-format/).

This is a fork of [Noofbiz/tmx](https://github.com/Noofbiz/tmx), which is licensed under the [MIT license](LICENSE).

This fork is also licensed under the [MIT License](LICENSE).

## Goals of this fork

- [x] Favor streaming (e.g. change `xml.Unmarshal()` to `xml.NewDecoder().Decode()`).
- [ ] Replace the global gTMXPath (née TMXURL) with something safe to use in a concurrent environment.

## Example Usage

```go
m, err := tmx.ParseFile("test1.tmx")
if err != nil {
  fmt.Println(err)
  return
}
// Do stuff with your tmx map...
```
