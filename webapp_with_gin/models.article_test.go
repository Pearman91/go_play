package main

import (
	"testing"
)

func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()
	
	// check ze delka articleListu stejna jako z getAllArticles
	if len(alist) != len(articleList) {
		t.Fail()
	}

	// check ze info o clancich je ok
	for i, v := range alist {
		if v.Content != articleList[i].Content ||
			v.ID != articleList[i].ID ||
			v.Title != articleList[i].Title {
			
			t.Fail()
			break
		}
	}
}
