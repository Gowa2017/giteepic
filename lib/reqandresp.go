package lib

// {"content":{"name":"teetst.png","path":"pic/teetst.png","size":643,"sha":"8193ff25d1c1ec412e9fe61264fdb8b2df9562d2","type":"file","url":"https://gitee.com/api/v5/repos/gowa2017/pics/contents/pic/teetst.png","html_url":"https://gitee.com/gowa2017/pics/blob/master/pic/teetst.png","download_url":"https://gitee.com/gowa2017/pics/raw/master/pic/teetst.png","_links":{"self":"https://gitee.com/api/v5/repos/gowa2017/pics/contents/pic/teetst.png","html":"https://gitee.com/gowa2017/pics/blob/master/pic/teetst.png"}},"commit":{"sha":"851c8ca9d59852c31079648659c8fad069e54843","author":{"name":"Gowa2017","date":"2021-08-24T03:17:29+00:00","email":"gowa2017@aliyun.com"},"committer":{"name":"Gitee","date":"2021-08-24T03:17:29+00:00","email":"noreply@gitee.com"},"message":"upload","tree":{"sha":"9e31af13ca9ddf558ec47587cabbb0fe8e8abf98","url":"https://gitee.com/api/v5/repos/gowa2017/pics/git/trees/9e31af13ca9ddf558ec47587cabbb0fe8e8abf98"},"parents":[{"sha":"f6aae123b98131979a3462eff8e1ef3b2f28838a","url":"https://gitee.com/api/v5/repos/gowa2017/pics/commits/f6aae123b98131979a3462eff8e1ef3b2f28838a"}]}// Generated by https://quicktype.io

type Req struct {
	Access_token string `json:"access_token"`
	Content      string `json:"content"`
	Message      string `json:"message"`
	Branch       string `json:"branch"`
}

type Resp struct {
	Content Content `json:"content"`
	Commit  Commit  `json:"commit"`
}

type Commit struct {
	SHA       string `json:"sha"`
	Author    Author `json:"author"`
	Committer Author `json:"committer"`
	Message   string `json:"message"`
	Tree      Tree   `json:"tree"`
	Parents   []Tree `json:"parents"`
}

type Author struct {
	Name  string `json:"name"`
	Date  string `json:"date"`
	Email string `json:"email"`
}

type Tree struct {
	SHA string `json:"sha"`
	URL string `json:"url"`
}

type Content struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Size        int64  `json:"size"`
	SHA         string `json:"sha"`
	Type        string `json:"type"`
	URL         string `json:"url"`
	HTMLURL     string `json:"html_url"`
	DownloadURL string `json:"download_url"`
	Links       Links  `json:"_links"`
}

type Links struct {
	Self string `json:"self"`
	HTML string `json:"html"`
}
