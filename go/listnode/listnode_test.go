package listnode

import (
	"fmt"
	"testing"
)

func TestListNode(t *testing.T) {
	node := getListNode()
	fmt.Print(node.toString())
}
