# Runvil

The landing page for the Runvil ecosystem, served at
[https://runvil.github.io/](https://runvil.github.io/) via GitHub Pages.

The site is built with the [Runvil framework](https://github.com/runvil/framework)
static site generator (`web/ssg`): pages are composed from named components
with scoped styles, wrapped by a layout, and themed with the `web` theming
system. Links to the documentation site at
[https://runvil.github.io/docs/](https://runvil.github.io/docs/) and to the
ecosystem repositories under [github.com/runvil](https://github.com/runvil).

## Repository layout

The repository follows a source/deploy branch split:

- `main` — the source of truth: `runvil.yaml` (SSG config), specs under
  `docs/specs/`, and this README. No generated output is committed here.
- `gh-pages` — the built static site (`site/` output), served directly by
  GitHub Pages.

## Build

The build writes the generated site into `site/`:

```sh
runvil build
# created site/index.html
# created site/404.html
# created site/assets/style.css
```

## Deploy to GitHub Pages

After building, publish the contents of `site/` to the `gh-pages` branch:

```sh
git checkout gh-pages
git rm -rq . ':(exclude).'
cp -r site/* .
git add -A && git commit -m "deploy: update landing page"
git push origin gh-pages
```

## Serve locally

```sh
runvil build
python3 -m http.server -d site
# open http://localhost:8000
```

## Test

```sh
go test ./...
```

## License

MIT — see [LICENSE](./LICENSE).