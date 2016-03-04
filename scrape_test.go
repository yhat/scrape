package scrape

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

const testHTML = `
<html>
  <body>
    <div class="bigbird">
      <div class="container">
        <div class="bigbird">
          Hi, I'm Big Bird
        </div>
      </div>
    </div>
  </body>
</html>
`

func TestFindAllNestedReturnsNestedMatchingNodes(t *testing.T) {
	node, _ := html.Parse(strings.NewReader(testHTML))
	allResults := FindAllNested(node, ByClass("bigbird"))

	if len(allResults) != 2 {
		t.Error("Expected 2 nodes returned but only found", len(allResults))
	}
}

func TestFindAllDoesNotReturnNestedMatchingNodes(t *testing.T) {
	node, _ := html.Parse(strings.NewReader(testHTML))
	allResults := FindAll(node, ByClass("bigbird"))

	if len(allResults) != 1 {
		t.Error("Expected 1 node returned but found", len(allResults))
	}
}
