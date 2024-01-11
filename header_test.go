package component

import (
	"context"
	"testing"

	"github.com/DistilledP/htmx-components/test"
	"github.com/a-h/templ"
)

func TestHeaderFuncs(t *testing.T) {
	testCases := []struct {
		name     string
		text     string
		classes  []string
		expected string
		testFn   func(string, ...string) templ.Component
	}{
		{
			name:     "Header1 without classes",
			text:     "Test 0",
			classes:  []string{},
			expected: `<h1>Test 0</h1>`,
			testFn:   Header1,
		},
		{
			name:     "Header1 with classes",
			text:     "Test 1",
			classes:  []string{"class1", "class2"},
			expected: `<h1 class="class1 class2">Test 1</h1>`,
			testFn:   Header1,
		},
		{
			name:     "Header2 without classes",
			text:     "Test 2",
			classes:  []string{},
			expected: `<h2>Test 2</h2>`,
			testFn:   Header2,
		},
		{
			name:     "Header2 with classes",
			text:     "Test 3",
			classes:  []string{"class1", "class2"},
			expected: `<h2 class="class1 class2">Test 3</h2>`,
			testFn:   Header2,
		},
		{
			name:     "Header3 without classes",
			text:     "Test 4",
			classes:  []string{},
			expected: `<h3>Test 4</h3>`,
			testFn:   Header3,
		},
		{
			name:     "Header3 with classes",
			text:     "Test 5",
			classes:  []string{"class1", "class2"},
			expected: `<h3 class="class1 class2">Test 5</h3>`,
			testFn:   Header3,
		},
		{
			name:     "Header4 without classes",
			text:     "Test 6",
			classes:  []string{},
			expected: `<h4>Test 6</h4>`,
			testFn:   Header4,
		},
		{
			name:     "Header4 with classes",
			text:     "Test 7",
			classes:  []string{"class1", "class2"},
			expected: `<h4 class="class1 class2">Test 7</h4>`,
			testFn:   Header4,
		},
		{
			name:     "Header5 without classes",
			text:     "Test 8",
			classes:  []string{},
			expected: `<h5>Test 8</h5>`,
			testFn:   Header5,
		},
		{
			name:     "Header5 with classes",
			text:     "Test 9",
			classes:  []string{"class1", "class2"},
			expected: `<h5 class="class1 class2">Test 9</h5>`,
			testFn:   Header5,
		},
		{
			name:     "Header6 without classes",
			text:     "Test 10",
			classes:  []string{},
			expected: `<h6>Test 10</h6>`,
			testFn:   Header6,
		},
		{
			name:     "Header6 with classes",
			text:     "Test 11",
			classes:  []string{"class1", "class2"},
			expected: `<h6 class="class1 class2">Test 11</h6>`,
			testFn:   Header6,
		},
	}

	t.Parallel()

	for i, testCase := range testCases {
		idx := i
		tc := testCase

		testWriter := test.NewTestWriter()

		component := tc.testFn(tc.text, tc.classes...)

		err := component.Render(context.Background(), testWriter)
		if err != nil {
			t.Errorf("[%d]%s failed, render error: %s", idx, tc.name, err)
		}

		if !testWriter.ContainsString(tc.expected) {
			t.Errorf("[%d]%s failed, expected %s got %s", idx, tc.name, tc.expected, testWriter.ToString())
		}
	}
}
