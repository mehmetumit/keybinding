# keybinding

A golang wrapper for parsing gocui keybindings.

```go
// get a single keybinding
keybinding.Parse("ctrl+c")
keybinding.Parse("/")
keybinding.Parse("ü")
keybinding.Parse("ö")

// get a list of keybindings
keybinding.ParseAll("ctrl+A, ctrl+B, CTRL+ALT+C")
```
