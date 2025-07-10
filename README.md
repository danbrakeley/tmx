# tmx

## Overview

Go package for parsing [Tiled](https://www.mapeditor.org/) [TMX map files](https://doc.mapeditor.org/en/stable/reference/tmx-map-format/).

This is a fork of [Noofbiz/tmx](https://github.com/Noofbiz/tmx), which is licensed under the [MIT license](LICENSE).

This fork is also licensed under the [MIT License](LICENSE).

## Goals of this fork

- [x] Favor streaming (e.g. change `xml.Unmarshal()` to `xml.NewDecoder().Decode()`).
- [x] Replace TMXURL with something safe to use in a concurrent environment.
- [x] Add option to skip loading any external files
- [ ] Clean up `tmx:"ref"` tag <-- is this needed? Just look for LoadRefs on everything?
  - is there a performance impact?
- [ ] Are nested templates allowed? Does this code handle them properly?
- [ ] Build PR to submit back to Noofbiz/tmx

## Example Usage

```go
m, err := tmx.ParseFile("test1.tmx")
if err != nil {
  fmt.Println(err)
  return
}
// Do stuff with your tmx map...
```
