package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"simple-webcrawler/lib/logger"
)

// Fetch 根据一个Url从网页拉取整个html内容并送到解析器
func Fetch(url string) ([]byte, error) {
	// 创建一个新的请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}
	req.Header.Set("Cookie", "FSSBBIl1UgzbN7NO=5HaX9_xEbT04XlDFr3YkQkpx0zkH4Oqmn.EtFRj74entf4cN8rHoCHYdtTGgXbwjIll7fQVj.1QnjwNlpOxQ.wA;sid=4b52de1c-661f-47ea-8f33-5d42fffaf035;Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1690870648;ec=u2sJDAAf-1690870648568-38ece5744f763-1455864551;token=1800238612.1690871380447.52b8b987b5cc3803e5ec0f5d9db76bc8;refreshToken=1800238612.1690957780447.384bdc635d7079c3853a5fd4415ac207;_pc_myzhenai_showdialog_=1; _pc_myzhenai_memberid_=%22%2C1800238612%22;recommendId=ZaMainpageTa-0001-online_xgb-0.0.1-za_user_profile-0.0.1;_exid=%2Fp1xJU7Jd0XWiShsA9ILFkNcK66PlxwWX5Nj6cD5SWaJfBLmiyVpvZo1vLJZI3lQoT0H2EonWixMl9U9X1NWVA%3D%3D;_efmdata=TC%2BGD5u7b67QIdrcJTeKkTCVCra0t%2Fd6PpfY6vS3pPGNcJCj4V%2Btwmbhc%2FyCgKzxbIyV4mqlFGtIHtDkPvA9ZciSqPYttjTvMWdOxU2RS%2F4%3D;FSSBBIl1UgzbN7NP=5RWP.ICsEj6VqqqDx.e6ukGtCIG.ANXo6nwk7T6HaNfVoqOYjOu2Iwf5_nK6iE8XsWkUp7sqH6w8yH9dI549J..b5e3k.HZRsYGjYkchvs5rAAsGUknEzQS8IpC.rqPNiixUlR63v7RPFYiwJTHxo637vtYcPbzQlLXQPOvg7IbC4CXfV_Iw877ncNWC8457VH6_0XX5r4bBLK3419MCIi7YYKoSJOWAqAp2sL6kGUa.2lFHuoZLUodN9SN.5AY5vG;Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1690872696;lrt=1690872697224")
	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		logger.Error("Error: status code: ", resp.StatusCode)
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	} else {
		bodyReader := bufio.NewReader(resp.Body)
		e := determineEncoding(bodyReader)
		utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
		return ioutil.ReadAll(utf8Reader)
	}
}

// determineEncoding 根据html头解析指定的字符集
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
