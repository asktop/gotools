package ahtml

import (
	"testing"
)

func TestStripTags(t *testing.T) {
	src := `<p>Test paragraph.</p><!-- Comment -->  <a href="#fragment">Other text</a>`
	dst := `Test paragraph.  Other text`

	t.Log(StripTags(src) == dst)
	t.Log(StripTags(src))
}

func TestEntities(t *testing.T) {
	src := `A 'quote' "is" <b>bold</b>`
	dst := `A &#39;quote&#39; &#34;is&#34; &lt;b&gt;bold&lt;/b&gt;`
	t.Log(Entities(src) == dst)
	t.Log(Entities(src))

	t.Log(EntitiesDecode(dst) == src)
	t.Log(EntitiesDecode(dst))
}

func TestSpecialChars(t *testing.T) {
	src := `A 'quote' "is" <b>bold</b>`
	dst := `A &#39;quote&#39; &#34;is&#34; &lt;b&gt;bold&lt;/b&gt;`
	t.Log(SpecialChars(src) == dst)
	t.Log(SpecialChars(src))

	t.Log(SpecialCharsDecode(dst) == src)
	t.Log(SpecialCharsDecode(dst))
}
