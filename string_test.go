package httpexpect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringFailed(t *testing.T) {
	chain := makeChain(newMockReporter(t))

	chain.fail("fail")

	value := &String{chain, ""}

	value.Empty()
	value.NotEmpty()
	value.Equal("")
	value.NotEqual("")
	value.EqualFold("")
	value.NotEqualFold("")
	value.Contains("")
	value.NotContains("")
	value.ContainsFold("")
	value.NotContainsFold("")
}

func TestStringEmpty(t *testing.T) {
	reporter := newMockReporter(t)

	value1 := NewString(reporter, "")

	value1.Empty()
	value1.chain.assertOK(t)
	value1.chain.reset()

	value1.NotEmpty()
	value1.chain.assertFailed(t)
	value1.chain.reset()

	value2 := NewString(reporter, "a")

	value2.Empty()
	value2.chain.assertFailed(t)
	value2.chain.reset()

	value2.NotEmpty()
	value2.chain.assertOK(t)
	value2.chain.reset()
}

func TestStringEqual(t *testing.T) {
	reporter := newMockReporter(t)

	value := NewString(reporter, "foo")

	assert.Equal(t, "foo", value.Raw())

	value.Equal("foo")
	value.chain.assertOK(t)
	value.chain.reset()

	value.Equal("FOO")
	value.chain.assertFailed(t)
	value.chain.reset()

	value.NotEqual("FOO")
	value.chain.assertOK(t)
	value.chain.reset()

	value.NotEqual("foo")
	value.chain.assertFailed(t)
	value.chain.reset()
}

func TestStringEqualFold(t *testing.T) {
	reporter := newMockReporter(t)

	value := NewString(reporter, "foo")

	value.EqualFold("foo")
	value.chain.assertOK(t)
	value.chain.reset()

	value.EqualFold("FOO")
	value.chain.assertOK(t)
	value.chain.reset()

	value.EqualFold("foo2")
	value.chain.assertFailed(t)
	value.chain.reset()

	value.NotEqualFold("foo")
	value.chain.assertFailed(t)
	value.chain.reset()

	value.NotEqualFold("FOO")
	value.chain.assertFailed(t)
	value.chain.reset()

	value.NotEqualFold("foo2")
	value.chain.assertOK(t)
	value.chain.reset()
}

func TestStringContains(t *testing.T) {
	reporter := newMockReporter(t)

	value := NewString(reporter, "11-foo-22")

	value.Contains("foo")
	value.chain.assertOK(t)
	value.chain.reset()

	value.Contains("FOO")
	value.chain.assertFailed(t)
	value.chain.reset()

	value.NotContains("FOO")
	value.chain.assertOK(t)
	value.chain.reset()

	value.NotContains("foo")
	value.chain.assertFailed(t)
	value.chain.reset()
}

func TestStringContainsFold(t *testing.T) {
	reporter := newMockReporter(t)

	value := NewString(reporter, "11-foo-22")

	value.ContainsFold("foo")
	value.chain.assertOK(t)
	value.chain.reset()

	value.ContainsFold("FOO")
	value.chain.assertOK(t)
	value.chain.reset()

	value.ContainsFold("foo3")
	value.chain.assertFailed(t)
	value.chain.reset()

	value.NotContainsFold("foo")
	value.chain.assertFailed(t)
	value.chain.reset()

	value.NotContainsFold("FOO")
	value.chain.assertFailed(t)
	value.chain.reset()

	value.NotContainsFold("foo3")
	value.chain.assertOK(t)
	value.chain.reset()
}

func TestStringLength(t *testing.T) {
	reporter := newMockReporter(t)

	value := NewString(reporter, "1234567")

	num := value.Length()
	value.chain.assertOK(t)
	num.chain.assertOK(t)
	assert.Equal(t, 7.0, num.Raw())
}

func TestStringMatchOne(t *testing.T) {
	reporter := newMockReporter(t)

	value := NewString(reporter, "http://example.com/users/john")

	m1 := value.Match(`http://(?P<host>.+)/users/(?P<user>.+)`)
	m1.chain.assertOK(t)

	assert.Equal(t,
		[]string{"http://example.com/users/john", "example.com", "john"},
		m1.submatches)

	m2 := value.Match(`http://(.+)/users/(.+)`)
	m2.chain.assertOK(t)

	assert.Equal(t,
		[]string{"http://example.com/users/john", "example.com", "john"},
		m2.submatches)
}

func TestStringMatchAll(t *testing.T) {
	reporter := newMockReporter(t)

	value := NewString(reporter,
		"http://example.com/users/john http://example.com/users/bob")

	m := value.MatchAll(`http://(\S+)/users/(\S+)`)

	assert.Equal(t, 2, len(m))

	m[0].chain.assertOK(t)
	m[1].chain.assertOK(t)

	assert.Equal(t,
		[]string{"http://example.com/users/john", "example.com", "john"},
		m[0].submatches)

	assert.Equal(t,
		[]string{"http://example.com/users/bob", "example.com", "bob"},
		m[1].submatches)
}

func TestStringMatchStatus(t *testing.T) {
	reporter := newMockReporter(t)

	value := NewString(reporter, "a")

	value.Match(`a`)
	value.chain.assertOK(t)
	value.chain.reset()

	value.MatchAll(`a`)
	value.chain.assertOK(t)
	value.chain.reset()

	value.NotMatch(`a`)
	value.chain.assertFailed(t)
	value.chain.reset()

	value.Match(`[^a]`)
	value.chain.assertFailed(t)
	value.chain.reset()

	value.MatchAll(`[^a]`)
	value.chain.assertFailed(t)
	value.chain.reset()

	value.NotMatch(`[^a]`)
	value.chain.assertOK(t)
	value.chain.reset()

	assert.Equal(t, []string{}, value.Match(`[^a]`).submatches)
	assert.Equal(t, []Match{}, value.MatchAll(`[^a]`))
}

func TestStringMatchInvalid(t *testing.T) {
	reporter := newMockReporter(t)

	value := NewString(reporter, "a")

	value.Match(`[`)
	value.chain.assertFailed(t)
	value.chain.reset()

	value.MatchAll(`[`)
	value.chain.assertFailed(t)
	value.chain.reset()

	value.NotMatch(`[`)
	value.chain.assertFailed(t)
	value.chain.reset()
}
