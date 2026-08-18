// The Runvil ecosystem landing page, built with the web/ssg static site
// generator: components with scoped styles, a layout with a content slot,
// and the ui theming system. Build writes the site into the repository
// root, which GitHub Pages serves directly.
package main

import (
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/runvil/framework/ui"
	"github.com/runvil/framework/web/ssg"
)

const favicon = "data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 32 32'%3E%3Crect width='32' height='32' rx='7' fill='%2300c853'/%3E%3Cpath d='M9 22 L9 10 L13 10 L13 18 L19 18 L19 10 L23 10 L23 22 Z' fill='%23fff'/%3E%3C/svg%3E"

// Repo is one ecosystem repository shown in the repositories grid.
type Repo struct {
	Name string
	URL  string
	Desc template.HTML
}

// SiteData is the value passed to every component and layout.
type SiteData struct {
	Title       string
	Description string
	Tagline     string
	Theme       *ui.Theme
	Features    []template.HTML
	Repos       []Repo
	Snippet     string
}

func main() {
	out := "."
	if len(os.Args) > 1 {
		out = os.Args[1]
	}
	created, err := buildSite().Build(out)
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range created {
		fmt.Println("created " + p)
	}
}

func buildSite() *ssg.Site {
	data := siteData()
	return ssg.New().
		Component(ssg.Component{
			Name: "topbar",
			Body: `<header class="top">
  <a class="brand" href="/">Run<span>vil</span></a>
  <nav>
    <a href="/docs/">Documentation</a>
    <a href="https://github.com/runvil">GitHub</a>
    {{if .Theme}}{{.Theme.Button}}{{end}}
  </nav>
</header>`,
			Style: `.top {
  display: flex; align-items: center; justify-content: space-between;
  padding: 1.1rem 1.5rem; max-width: 64rem; margin: 0 auto;
}
.top .brand { font-weight: 800; font-size: 1.15rem; letter-spacing: .02em; color: var(--ink); text-decoration: none; }
.top .brand span { color: var(--accent); }
.top nav a { text-decoration: none; font-weight: 600; margin-left: 1.25rem; color: var(--ink); }
.top nav a:hover { color: var(--accent); }`,
		}).
		Component(ssg.Component{
			Name: "hero",
			Body: `<section class="hero">
  <h1>A meta-framework, <span>written in Go</span>.</h1>
  <p class="tagline">{{.Tagline}}</p>
  <p class="cta">
    <a class="primary" href="/docs/">Read the documentation</a>
    <a class="secondary" href="https://github.com/runvil">Browse the code</a>
  </p>
</section>`,
			Style: `.hero { text-align: center; padding: 3rem 0 2rem; }
.hero h1 { font-size: clamp(2.2rem, 6vw, 3.4rem); line-height: 1.1; margin: 0 0 .6rem; }
.hero h1 span { color: var(--accent); }
.hero .tagline { font-size: 1.2rem; color: var(--muted); max-width: 44rem; margin: 0 auto 1.8rem; }`,
		}).
		Component(ssg.Component{
			Name: "features",
			Body: `<section>
  <h2>Why Runvil</h2>
  <ul class="features">
    {{range .Features}}<li>{{.}}</li>{{end}}
  </ul>
</section>`,
		}).
		Component(ssg.Component{
			Name: "repos",
			Body: `<section>
  <h2>Repositories</h2>
  <div class="grid">
    {{range .Repos}}
    <a class="card" href="{{.URL}}">
      <h3>{{.Name}}</h3>
      <p>{{.Desc}}</p>
    </a>
    {{end}}
  </div>
</section>`,
			Style: `.grid { display: grid; grid-template-columns: repeat(auto-fit, minmax(15rem, 1fr)); gap: 1rem; }
.card { background: var(--card); border: 1px solid var(--neutral); border-radius: var(--radius); padding: 1.25rem; text-decoration: none; color: inherit; transition: border-color .15s ease, transform .15s ease; }
.card:hover { border-color: var(--primary); transform: translateY(-2px); }
.card h3 { margin: 0 0 .35rem; font-size: 1.05rem; }
.card p { margin: 0; color: var(--muted); font-size: .95rem; }`,
		}).
		Component(ssg.Component{
			Name: "start",
			Body: `<section>
  <h2>Getting started</h2>
  <p>Scaffold a new Runvil project in seconds with the developer tool.</p>
  <pre>{{.Snippet}}</pre>
</section>`,
		}).
		Component(ssg.Component{
			Name: "home",
			Body: `{{component "hero" .}}
{{component "features" .}}
{{component "repos" .}}
{{component "start" .}}`,
		}).
		Component(ssg.Component{
			Name: "notfound",
			Body: `<section class="nf">
  <h1>404 — page not found</h1>
  <p>The page you are looking for does not exist.</p>
  <p class="cta"><a class="primary" href="/">Back home</a></p>
</section>`,
			Style: `.nf { text-align: center; padding: 3rem 0; }
.nf h1 { font-size: 2rem; }`,
		}).
		Component(ssg.Component{
			Name: "footer",
			Body: `<footer>
  <p>&copy; <span id="year">2026</span> Runvil Contributors &middot; MIT License</p>
</footer>
<script>
  document.getElementById('year').textContent = new Date().getFullYear();
</script>`,
			Style: `footer { max-width: 64rem; margin: 0 auto; padding: 1.5rem; color: var(--muted); font-size: .9rem; border-top: 1px solid var(--neutral); }`,
		}).
		Layout(ssg.Layout{
			Name: "site",
			Body: `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta name="color-scheme" content="light dark">
<meta name="description" content="{{.Data.Description}}">
<title>{{.Title}}</title>
<link rel="icon" href="` + favicon + `">
<link rel="stylesheet" href="/assets/style.css">
<link rel="stylesheet" href="/assets/theme.css">
{{if .Data.Theme}}{{.Data.Theme.Script}}{{end}}
</head>
<body>
{{component "topbar" .Data}}
<main>
{{.Content}}
</main>
{{component "footer" .Data}}
</body>
</html>`,
			Style: `:root {
  --radius: 12px;
  --card: var(--base-2);
  --bg: var(--base-1);
  --ink: var(--base-1-content);
  --muted: var(--base-2-content);
  --accent: var(--primary);
  --accent-soft: var(--ghost);
  --rule: var(--neutral);
}
* { box-sizing: border-box; }
body { margin: 0; font: 17px/1.6 ui-sans-serif, system-ui, -apple-system, "Segoe UI", sans-serif; color: var(--ink); background: var(--bg); transition: background-color .2s ease, color .2s ease; }
a { color: var(--accent); }
main { max-width: 64rem; margin: 0 auto; padding: 2rem 1.5rem 4rem; }
section { margin-top: 3.5rem; }
h2 { font-size: 1.5rem; margin: 0 0 1rem; }
pre { background: var(--ghost); border: 1px solid var(--neutral); border-radius: var(--radius); padding: 1rem 1.25rem; overflow-x: auto; font: .95rem/1.5 ui-monospace, "SF Mono", Menlo, Consolas, monospace; color: var(--base-1-content); }
ul.features { padding-left: 1.25rem; max-width: 44rem; }
ul.features li { margin: .45rem 0; }
.cta { display: flex; gap: .8rem; justify-content: center; flex-wrap: wrap; }
.cta a { display: inline-block; padding: .7rem 1.5rem; border-radius: 999px; text-decoration: none; font-weight: 700; }
.cta a.primary { background: var(--primary); color: var(--primary-content); }
.cta a.secondary { border: 1px solid var(--neutral); color: var(--ink); }`,
		}).
		Asset("assets/theme.css", ui.ThemeModeVarsCSS+"\n"+ui.ThemeToggleCSS).
		Page(ssg.Page{Path: "/", Title: data.Title, Layout: "site", Root: "home", Data: data}).
		Page(ssg.Page{Path: "/404.html", Title: "404 — " + data.Title, Layout: "site", Root: "notfound", Data: data})
}

func siteData() SiteData {
	return SiteData{
		Title:       "Runvil — A Meta-Framework Written in Go",
		Description: "Runvil composes modules sourced across multiple ecosystems into one cohesive, high-performance foundation for building applications in Go.",
		Tagline:     "Runvil composes modules sourced across multiple ecosystems and repositories into one cohesive, high-performance foundation for building web services, CLI tools, background workers, and desktop applications.",
		Theme:       &ui.Theme{},
		Features: []template.HTML{
			template.HTML(`<strong>Meta-framework by design</strong> — integrates what other ecosystems do well instead of re-implementing it.`),
			template.HTML(`<strong>Stdlib-first</strong> — argument parsing, logging, and configuration come from the Go standard library.`),
			template.HTML(`<strong>Safe by design</strong> — Go's memory safety, no <code>unsafe</code>, no manual memory management.`),
			template.HTML(`<strong>Spec-driven</strong> — every component ships with a formal specification describing requirements and conventions.`),
			template.HTML(`<strong>Dogfooded</strong> — Runvil's own tooling and websites are built with Runvil.`),
		},
		Repos: []Repo{
			{Name: "framework", URL: "https://github.com/runvil/framework", Desc: template.HTML(`The meta-framework: integrated <code>cli</code> and <code>web</code> packages.`)},
			{Name: "libs", URL: "https://github.com/runvil/libs", Desc: template.HTML(`Shared, reusable libraries: <code>core</code> and <code>term</code>.`)},
			{Name: "runvil", URL: "https://github.com/runvil/runvil", Desc: template.HTML(`The developer tool for scaffolding, testing, and project information.`)},
			{Name: "mdbind", URL: "https://github.com/runvil/mdbind", Desc: template.HTML(`The site builder: Markdown folders into book-shaped static websites.`)},
			{Name: "docs", URL: "https://github.com/runvil/docs", Desc: template.HTML(`This ecosystem's website — <a href="/docs/">read it online</a>.`)},
		},
		Snippet: "go install github.com/runvil/runvil/cmd/runvil@v0.1.0\nrunvil new hello --module example.com/hello\ncd hello && runvil test",
	}
}
