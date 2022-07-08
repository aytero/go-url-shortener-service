package dataStorage

import (
	"testing"
	"urlShortener/pkg/config"
	"urlShortener/pkg/model"
)

// Test the function that fetches all articles
//func TestGetAllArticles(t *testing.T) {
//	alist := getAllArticles()
//
//	// Check that the length of the list of articles returned is the
//	// same as the length of the global variable holding the list
//	if len(alist) != len(articleList) {
//		t.Fail()
//	}
//
//	// Check that each member is identical
//	for i, v := range alist {
//		if v.Content != articleList[i].Content ||
//			v.ID != articleList[i].ID ||
//			v.Title != articleList[i].Title {
//
//			t.Fail()
//			break
//		}
//	}
//}

var validEntry = []model.DbEntry{
	{Id: 1, ShortUrl: "zDGkSmE0Qp", FullUrl: "https://lalalaal/sadasfasw"},
	{Id: 2, ShortUrl: "aaakSmE0Qp", FullUrl: "url"},
	{Id: 3, ShortUrl: "zfGksME0Qp", FullUrl: "https://example.com/sa309uadjuaio_s"},
}

func TestLocalStoragePost(t *testing.T) {
	local := NewLocalStorage(&config.Config{})
	fullUrl := "https://www.youtube.com/watch?v=dQw4w9WgXcQ"
	shortUrl := "zDGkSmE0Qp"
	_, err := local.PostUrl(fullUrl, shortUrl)
	if err != nil {
		t.Errorf("Local storage POST failed\n")
	}
}

func TestLocalStorageGet(t *testing.T) {
	local := NewLocalStorage(&config.Config{})

	for _, e := range validEntry {
		_, _ = local.PostUrl(e.ShortUrl, e.FullUrl)
	}
	for _, f := range validEntry {
		res, _ := local.GetUrl(f.ShortUrl)
		if f.FullUrl != res {
			t.Errorf("Local storage GET full_url failed\n")
		}
		res, _ = local.GetShortUrl(f.FullUrl)
		if f.ShortUrl != res {
			t.Errorf("Local storage GET short_url failed\n")
		}
	}
}
