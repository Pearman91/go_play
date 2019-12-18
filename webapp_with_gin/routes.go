package main

func initializeRoutes() {
	//index route handler
	router.GET("/", showIndexPage)

	//article route handler
	router.GET("/article/view/:article_id", getArticle)
}
