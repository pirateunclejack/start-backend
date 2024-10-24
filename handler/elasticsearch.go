package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"start-backend/els"
	"start-backend/model"

	"github.com/gin-gonic/gin"
)

func ElasticsearchPut(c *gin.Context)  {
    var product model.Product
    err := c.BindJSON(&product)
    if err != nil {
        c.JSON(
            http.StatusBadRequest, 
            gin.H{
                "error": "failed to convert request body to product",
                "error_detail": err.Error(),
            },
    	)
    }
    // data, err := json.Marshal(product)
    // if err != nil {
    //     c.JSON(
    //         http.StatusInternalServerError, 
    //         gin.H{
    //             "error": "failed to marshal product",
    //             "error_detail": err.Error(),
    //         },
    // 	)
    // }
    els_client, err := els.NewElasticsearchClient()
    if err != nil {
        c.JSON(
            http.StatusInternalServerError, 
            gin.H{
                "error": "failed to create Elasticsearch client",
                "error_detail": err.Error(),
            },
        )   
    }
    res, err := els_client.Index(els.Elasticsearch_index).
    Id(product.ID).
    Request(product).
    Do(context.TODO())
    if err != nil {
        c.JSON(
            http.StatusInternalServerError, 
            gin.H{
                "error": "failed to create index product to elasticsearch",
                "error_detail": err.Error(),
            },
        )
    }
    c.JSON(
        http.StatusOK, 
        gin.H{
            "product": res.Result.String(),
        },
    )
}

func ElasticsearchGet(c *gin.Context)  {
    id := c.Params.ByName("id")
    if id == "" {
        c.JSON(
            http.StatusBadRequest, 
            gin.H{
                "error": "no product id provided",
            },
    	)
    }

    client, err := els.NewElasticsearchClient()
    if err != nil {
        c.JSON(
            http.StatusInternalServerError, 
            gin.H{
                "error": "failed to create Elasticsearch client",
                "error_detail": err.Error(),
            },
        )   
    }

    res, err := client.Get(els.Elasticsearch_index, id).
        Do(context.TODO())
    if err != nil {
        c.JSON(
            http.StatusInternalServerError, 
            gin.H{
                "error": "failed to query product",
                "error_detail": err.Error(),
            },
        )   
    }

    var product model.Product
    log.Println("get product result", res)
    err = json.Unmarshal(res.Source_, &product)

    if err != nil {
        c.JSON(
            http.StatusInternalServerError, 
            gin.H{
                "error": "failed to unmarshal product",
                "error_detail": err.Error(),
            },
        )
        return
    }

    c.JSON(
        http.StatusOK, 
        gin.H{
            "result": product,
        },
    )
}



func ElasticsearchDelete(c *gin.Context)  {
    id := c.Params.ByName("id")
    if id == "" {
        c.JSON(
            http.StatusBadRequest, 
            gin.H{
                "error": "no product id provided",
            },
        )
    }

    client, err := els.NewElasticsearchClient()
    if err != nil {
        c.JSON(
            http.StatusInternalServerError, 
            gin.H{
                "error": "failed to create Elasticsearch client",
                "error_detail": err.Error(),
            },
        )   
    }

    res, err := client.Delete(els.Elasticsearch_index, id).
        Do(context.TODO())
    if err != nil {
        c.JSON(
            http.StatusInternalServerError,
            gin.H{
                "error": "failed to delete product",
                "error_detail": err.Error(),
            },
        )
    }

    c.JSON(
        http.StatusOK, 
        gin.H{
            "result": res.Result.String(),
        },
    )
}
