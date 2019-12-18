package main

import "errors"

type article struct {
  ID      int    "json:id"
  Title   string "json:title"
  Content string "json:content"
}

var articleList = []article{
	article{ID: 1, Title: "Prvni clanek", Content: "Tma neni ve dne videt"},
	article{ID: 2, Title: "Druhy clanek", Content: "Svetlo neni v noci videt"},
}

func getAllArticles() []article {
  return articleList
}

func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}
