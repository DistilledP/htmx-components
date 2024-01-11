package component

import (
	"context"
	"testing"

	"github.com/DistilledP/htmx-components/test"
)

func TestPage(t *testing.T) {
	expectedString := "<title>Test</title>"

	testWriter := test.NewTestWriter()
	component := Page("Test")

	err := component.Render(context.Background(), testWriter)
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	if !testWriter.ContainsString(expectedString) {
		t.Errorf("failed to find: '%s' in '%s'", expectedString, testWriter.ToString())
	}
}
