package articles

import (
	"log"
	"net/http"
	"sync"

	"example.com/test/core/db"
	"example.com/test/core/logger"
	"example.com/test/core/req"
	"example.com/test/core/res"
	"example.com/test/core/scrape"
	"github.com/gin-gonic/gin"
)

func create(c *gin.Context) {
	auth, _ := req.Auth(c)
	result, err := _getScrapedData(c)

	if err != nil {
		return
	}

	articlesModel := []Article{}

	for _, article := range result {
		articlesModel = append(articlesModel, Article{
			Title:         article.Title,
			Url:           article.Url,
			TitleSelector: article.Selector,
			UserID:        auth.ID,
		})
	}

	data, createError := db.CreateMany(articlesModel)

	if createError != nil {
		res.Error(c, res.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    createError.Error(),
		})
		return
	}

	res.Success(c, res.Response{
		Message:    "Articles created successfully",
		StatusCode: http.StatusCreated,
		Data:       data,
	})
}

func findAll(c *gin.Context) {
	auth, _ := req.Auth(c)

	data, err := db.Query[Article](db.Options{
		// Where: map[string]any{"user_id": auth.ID},
		// WhereNot: map[string]any{"user_id": auth.ID},
		WhereRaw: []any{"user_id = ?", auth.ID},
		Offset:   0,
		Limit:    10,
		OrderBy:  "id asc",
	})

	if err != nil {
		res.Error(c, res.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	res.Success(c, res.Response{
		StatusCode: http.StatusOK,
		Data:       data,
	})
}

func paginate(c *gin.Context) {
	auth, _ := req.Auth(c)
	page, pageSize := req.Pagination(c)

	data, totalRecords, err := db.Paginate[Article](page, pageSize, db.Options{
		WhereRaw: []any{"user_id = ?", auth.ID},
		OrderBy:  "id desc",
	})

	if err != nil {
		res.Error(c, res.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	resData := map[string]any{
		"totalRecords": totalRecords,
		"data":         data,
		"page":         1,
		"pageSize":     20,
		"totalPages":   (totalRecords / 20) + 1,
	}

	res.Success(c, res.Response{
		StatusCode: http.StatusOK,
		Data:       resData,
	})
}

func inspect(c *gin.Context) {
	result, err := _getScrapedData(c)

	if err != nil {
		return
	}

	res.Success(c, res.Response{
		StatusCode: http.StatusOK,
		Data:       result,
	})
}

func _getScrapedData(c *gin.Context) ([]scrape.ScrappingResponse, error) {
	body, err := req.Body[createArticleDto](c)

	if err != nil {
		res.Error(c, res.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return nil, err
	}

	scrapper := scrape.Scraper{
		Mutex: &sync.Mutex{},
		DataToScrap: scrape.DataToScrap{
			Urls:          body.Urls,
			TitleSelector: body.TitleSelector,
		},
	}

	result, errScrapping := scrape.ScrapeURLs(scrapper)

	if errScrapping != nil {
		res.Error(c, res.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    errScrapping.Error(),
		})
		return nil, errScrapping
	}

	return result, nil
}

func UpdateScheduler() {
	articles, err := db.Find[Article]()

	if err != nil {
		logger.Error(err)
		return
	}

	for _, article := range articles {
		go func(article Article) {
			scrapper := scrape.Scraper{
				Mutex: &sync.Mutex{},
				DataToScrap: scrape.DataToScrap{
					Urls:          []string{article.Url},
					TitleSelector: article.TitleSelector,
				},
			}

			result, err := scrape.ScrapeURLs(scrapper)

			if err != nil {
				logger.Error(err)
				return
			}

			article.Title = result[0].Title
			article.Url = result[0].Url

			_, errUpdate := db.Update(article)

			if errUpdate != nil {
				logger.Error(errUpdate)
				return
			}

			log.Print("Updating article: ", article)
		}(article)
	}
}
