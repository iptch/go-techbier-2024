# go-techbier-2024

This is the repository for ipt's **TechBier Go**, an event to get an
introduction to Golang's many features.

This event revolves around a small TUI app, based on the amazing
[PokéAPI](https://pokeapi.co/), built using the
[bubbletea framework](https://github.com/charmbracelet/bubbletea). Our app
introduces the following Go concepts:

- structs and implementing interfaces
- use of standard library, with `net/http` and `encoding/json`
- importing a new package
- goroutines and channels
- and Golang's new generics!

Run the application with `go run .`:

![](/assets/demo.gif)

## Branches/exercises

We have prepared branches as exercises for our participants, each with their
respective solutions:

- tasks/1[-solution]
- tasks/2[-solution]
- tasks/3[-solution]
- tasks/bonus[-solution]

## Playing along

If you want to go through the TechBier, we recommend starting We prepared some
markdown [slides](/slides/theory.md) which we used together with
https://github.com/maaslalani/slides. You can either view it as plain markdown
or run it with:

```bash
slides slides/theory.md
```

You can execute most of the code snippets in the slides with `ctrl-E`.
