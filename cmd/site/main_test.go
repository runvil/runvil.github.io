package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestBuildLanding(t *testing.T) {
	out := t.TempDir()
	if _, err := buildSite().Build(out); err != nil {
		t.Fatal(err)
	}
	index, err := os.ReadFile(filepath.Join(out, "index.html"))
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"Runvil — A Meta-Framework Written in Go",
		`data-rv-component="topbar"`,
		`data-rv-component="repos"`,
		`data-rv-layout="site"`,
		`data-theme-toggle`,
		`runvilTheme`,
		`window.runvilTheme`,
	} {
		if !strings.Contains(string(index), want) {
			t.Errorf("index.html missing %q", want)
		}
	}
	if strings.Contains(string(index), `data-rv-component="hero" data-rv-component="home"`) {
		t.Error("duplicate scope attributes must be avoided")
	}
	css, err := os.ReadFile(filepath.Join(out, "assets", "style.css"))
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		`[data-rv-component="topbar"].top{`, // root element style
		`[data-rv-layout="site"] body{`,     // layout descendant style
		`[data-rv-component="repos"].card{`,
	} {
		if !strings.Contains(string(css), want) {
			t.Errorf("style.css missing %q", want)
		}
	}
	if _, err := os.Stat(filepath.Join(out, "404.html")); err != nil {
		t.Error("404.html missing")
	}
	if _, err := os.Stat(filepath.Join(out, "index.html")); err != nil {
		t.Error("index.html missing")
	}
}
