package node

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeSize(t *testing.T) {
	root := NewNode("", 0, nil)
	rootChild := NewNode("a", 1, root)
	root.AddChild(rootChild)

	assert.Equal(t, 1, root.ComputeValue())

	foo := NewNode("foo", 0, root)
	root.AddChild(foo)
	fooChild1 := NewNode("fooChild1", 1, foo)
	foo.AddChild(fooChild1)
	fooChild2 := NewNode("fooChild2", 2, foo)
	foo.AddChild(fooChild2)

	assert.Equal(t, 3, foo.ComputeValue())
	assert.Equal(t, 4, root.ComputeValue())
}

func TestRemoveChild(t *testing.T) {
	root := NewNode("", 0, nil)
	rootChild := NewNode("rootChild", 1, root)
	root.AddChild(rootChild)

	assert.Equal(t, 1, len(root.Children))

	root.RemoveChild(rootChild)

	assert.Equal(t, 0, len(root.Children))

	// Shouldn't panic
	root.RemoveChild(rootChild)
}

func TestFindChildById(t *testing.T) {
	root := NewNode("", 0, nil)
	rootChild := NewNode("a", 1, root)
	root.AddChild(rootChild)

	assert.Equal(t, &rootChild, root.FindChildById("a"))
	assert.Nil(t, root.FindChildById("b"))
}

func TestIsLastParent(t *testing.T) {
	root := NewNode("", 0, nil)
	rootChild := NewNode("a", 1, root)
	root.AddChild(rootChild)

	assert.False(t, root.IsLastParent())
	assert.False(t, rootChild.IsLastParent())

	parent := NewNode("parent", 0, root)
	root.AddChild(parent)
	child := NewNode("child", 1, parent)
	parent.AddChild(child)

	assert.True(t, parent.IsLastParent())
}
