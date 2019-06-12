package ahtml

import (
	"fmt"
	"testing"
)

func TestStripTags(t *testing.T) {
	src := `<p>Test paragraph.</p><!-- Comment -->  <a href="#fragment">Other text</a>`
	dst := `Test paragraph.  Other text`
	fmt.Println(StripTags(src) == dst)
	fmt.Println(StripTags(src))
}

func TestEntities(t *testing.T) {
	src := `A 'quote' "is" <b>bold</b>`
	dst := `A &#39;quote&#39; &#34;is&#34; &lt;b&gt;bold&lt;/b&gt;`
	fmt.Println(Entities(src) == dst)
	fmt.Println(Entities(src))

	fmt.Println(EntitiesDecode(dst) == src)
	fmt.Println(EntitiesDecode(dst))
}

func TestSpecialChars(t *testing.T) {
	src := `A 'quote' "is" <b>bold</b>`
	dst := `A &#39;quote&#39; &#34;is&#34; &lt;b&gt;bold&lt;/b&gt;`
	fmt.Println(SpecialChars(src) == dst)
	fmt.Println(SpecialChars(src))

	fmt.Println(SpecialCharsDecode(dst) == src)
	fmt.Println(SpecialCharsDecode(dst))
}
