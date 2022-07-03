// 这个示例程序展示如何写基础单元测试
package listing01

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

// TestDownload 确认http包的Get函数可以下载内容
func TestDownload(t *testing.T) {
	// url := "http://www.goinggo.net/feeds/posts/default?alt=rss"
	url := "https://www.baidu.com/"
	statusCode := 200

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"", url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call.", ballotX, err)
			}
			t.Log("\t\tShould be able to make the call.", checkMark)

			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t\tShould receive a \"%d\" status. %v", statusCode, checkMark)
			} else {
				t.Errorf("\t\tShould receive a \"%d\" status. %v %v", statusCode, ballotX, resp.StatusCode)
			}
		}
	}
}

/*
$ go test -v
=== RUN   TestDownload
    listing01_test.go:18: Given the need to test downloading content.
    listing01_test.go:20:       When checking "https://www.baidu.com/" for status code "200"
    listing01_test.go:26:               Should be able to make the call. ✓
    listing01_test.go:31:               Should receive a "200" status. ✓
--- PASS: TestDownload (0.24s)
PASS
ok      github.com/heqingbao/goinaction/chapter9/listing01      0.818s
*/
