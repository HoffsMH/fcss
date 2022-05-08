package fcss

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFindClassesInText(t *testing.T) {
	Convey("when the search term is in a file in the dir", t, func() {
		got := FindClassesInText(`
.test-class {
		.hello {
			&__hi {}
			.some-other-selector {
				some-rule: some-value;
				#some__selector { some-rule: some-value; }
				&__hi {
					someklsjdf: aslkjdf;
				}
			}
		}
}
		`)

		// doesn't get the rule or the value
		So(got, ShouldResemble, []string{"test-class", "hello", "some-other-selector", "some__selector", "__hi"})
	})
}

func TestFindClassesInFile(t *testing.T) {
	Convey("when the search term is in a file in the dir", t, func() {
		got := FindClassesInFile(`./test-dir/test.scss`)

		// doesn't get the rule or the value
		So(got, ShouldResemble, []string{"test-class", "hello", "some-other-selector", "some__selector", "__hi"})
	})
}
