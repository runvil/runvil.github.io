# Runvil

The landing page for the Runvil ecosystem, served at
[https://runvil.github.io/](https://runvil.github.io/) via GitHub Pages.

The site is built with the [Runvil framework](https://github.com/runvil/framework)
static site generator (`web/ssg`): pages are composed from named components
with scoped styles, wrapped by a layout, and themed with the `web` theming
system. Links to the documentation site at
[https://runvil.github.io/docs/](https://runvil.github.io/docs/) and to the
ecosystem repositories under [github.com/runvil](https://github.com/runvil).

## Build

The build writes the generated site into the repository root, which GitHub
Pages serves directly.

```sh
../runvil build
# created index.html
# created 404.html
# created assets/style.css
```

## Serve locally

```sh
../runvil build
python3 -m http.server
# open http://localhost:8000
```

## Test

```sh
go test ./...
```

## License

MIT — see [LICENSE](./LICENSE).