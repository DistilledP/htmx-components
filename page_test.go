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

	expectedScriptTags := []string{
		`<script crossorigin="anonymous" integrity="sha384-QFjmbokDn2DjBjq+fM+8LUIVrAgqcNW2s0PjAxHETgRn9l4fvX31ZxDxvwQnyMOX" src="https://unpkg.com/htmx.org@1.9.9"></script>`,
		`<script src="https://unpkg.com/htmx.org/dist/ext/ws.js"></script>`,
		`<script src="/static/js/reset-on-success.htmx.ext.js"></script>`,
		`<script src="/static/js/ws-events.htmx.ext.js"></script>`,
	}

	tags := []types.ScriptTag{
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

	component := Page("Test", tags)

	err := component.Render(context.Background(), testWriter)
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	if !testWriter.ContainsString(expectedString) {
		t.Errorf("failed to find: '%s' in '%s'", expectedString, testWriter.ToString())
	}

	foundScripts, missingScripts := testWriter.ContainsAllString(expectedScriptTags)
	if !foundScripts {
		t.Errorf("missing script tag(s): \n%s", strings.Join(missingScripts, "\n"))
	}
}
