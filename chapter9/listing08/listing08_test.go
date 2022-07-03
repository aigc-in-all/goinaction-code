// 这个示例程序展示如何写一个基本的表组测试
package listing08

import (
	"net/http"

	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestDownload(t *testing.T) {
	var urls = []struct {
		url        string
		statusCode int
	}{
		{
			// "http://www.goinggo.net/feeds/posts/default?alt=rss",
			"https://www.baidu.com/",
			http.StatusOK,
		},
		{
			// "http://rss.cnn.com/rss/cnn_topstbadurl.rss",
			"https://www.baidu.com/hello",
			http.StatusNotFound,
		},
	}

	t.Log("Given the need to test downloading different content.")

	for _, u := range urls {
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"", u.url, u.statusCode)
		{
			resp, err := http.Get(u.url)
			if err != nil {
				t.Fatal("\t\tShould be able to Get the url.",
					ballotX, err)
			}
			t.Log("\t\tShould be able to Get the url.",
				checkMark)

			defer resp.Body.Close()

			if resp.StatusCode == u.statusCode {
				t.Logf("\t\tShould have a \"%d\" status. %v", u.statusCode, checkMark)
			} else {
				t.Errorf("\t\tShould have a \"%d\" status. %v %v", u.statusCode, ballotX, resp.StatusCode)
			}
		}
	}
}

/*
$ go test -v
=== RUN   TestDownload
    listing08_test.go:30: Given the need to test downloading different content.
    listing08_test.go:33:       When checking "https://www.baidu.com/" for status code "200"
    listing08_test.go:40:               Should be able to Get the url. ✓
    listing08_test.go:46:               Should have a "200" status. ✓
    listing08_test.go:33:       When checking "https://www.baidu.com/hello" for status code "404"
    listing08_test.go:40:               Should be able to Get the url. ✓
    listing08_test.go:46:               Should have a "404" status. ✓
--- PASS: TestDownload (0.30s)
PASS
ok      github.com/heqingbao/goinaction/chapter9/listing08      0.590s
*/
