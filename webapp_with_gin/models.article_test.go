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
	for i, v range alist {
		if v.Content != articleList[i].Content ||
			ID.Content != articleList[i].ID ||
			Title.Content != articleList[i].Title {
			
			t.Fail()
			break
		}
	}
}
