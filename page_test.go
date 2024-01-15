package component

import (
	"context"
	"strings"
	"testing"

	"github.com/DistilledP/htmx-components/test"
	"github.com/DistilledP/htmx-components/types"
)

func TestPage(t *testing.T) {
	expectedString := "<title>Test</title>"

	testWriter := test.NewTestWriter()

	expectedMetaTags := []string{
		`<meta charset="utf-8" />`,
		`<meta content="width=device-width, initial-scale=1" name="viewport" />`,
	}

	expectedLinkTags := []string{
		`<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css" rel="stylesheet" />`,
	}

	expectedScriptTags := []string{
		`<script crossorigin="anonymous" integrity="sha384-QFjmbokDn2DjBjq+fM+8LUIVrAgqcNW2s0PjAxHETgRn9l4fvX31ZxDxvwQnyMOX" src="https://unpkg.com/htmx.org@1.9.9"></script>`,
		`<script src="https://unpkg.com/htmx.org/dist/ext/ws.js"></script>`,
		`<script src="/static/js/reset-on-success.htmx.ext.js"></script>`,
		`<script src="/static/js/ws-events.htmx.ext.js"></script>`,
	}

	metaTags := []types.MetaTag{
		{
			Charset: "utf-8",
		},
		{
			Name:    "viewport",
			Content: "width=device-width, initial-scale=1",
		},
	}

	linkTags := []types.LinkTag{
		{
			HRef: "https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css",
			Rel:  "stylesheet",
		},
	}

	scriptTags := []types.ScriptTag{
		{
			Src:         "https://unpkg.com/htmx.org@1.9.9",
			Integrity:   "sha384-QFjmbokDn2DjBjq+fM+8LUIVrAgqcNW2s0PjAxHETgRn9l4fvX31ZxDxvwQnyMOX",
			CrossOrigin: "anonymous",
		},
		{
			Src: "https://unpkg.com/htmx.org/dist/ext/ws.js",
		},
		{
			Src: "/static/js/reset-on-success.htmx.ext.js",
		},
		{
			Src: "/static/js/ws-events.htmx.ext.js",
		},
	}

	component := Page("Test", metaTags, linkTags, scriptTags)

	err := component.Render(context.Background(), testWriter)
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	if !testWriter.ContainsString(expectedString) {
		t.Errorf("failed to find: '%s' in '%s'", expectedString, testWriter.ToString())
	}

	foundMetas, missingMeta := testWriter.ContainsAllString(expectedMetaTags)
	if !foundMetas {
		t.Errorf("missing meta tag(s): \n%s", strings.Join(missingMeta, "\n"))
	}

	foundLinks, missingLinks := testWriter.ContainsAllString(expectedLinkTags)
	if !foundLinks {
		t.Errorf("missing link tag(s): \n%s", strings.Join(missingLinks, "\n"))
	}

	foundScripts, missingScripts := testWriter.ContainsAllString(expectedScriptTags)
	if !foundScripts {
		t.Errorf("missing script tag(s): \n%s", strings.Join(missingScripts, "\n"))
	}
}
