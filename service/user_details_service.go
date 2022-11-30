package service

import (
	"context"
	"encoding/json"
	"es/model"
	"log"

	elastic "github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
)

var dbName = "users_go"
var elasticClient1 *elastic.Client

func GetESClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL("https://192.168.1.6:9200"),
		elastic.SetBasicAuth("elastic", "fLiLoRYSTSEF=HLmgzoR"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
	)
	elasticClient1 = client
	info, code, err := elasticClient1.Ping("https://192.168.1.6:9200").Do(context.Background())
	if err != nil {
		panic(err)
	}
	log.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
	return elasticClient1, err
}

func Insert(request model.UserDetails) (model.UserDetails, error) {
	log.Println("Creating connection........")
	_, err := GetESClient()
	if err != nil {
		log.Println("Error initializing : ", err)
		panic("Client fail ")
	}
	log.Println("Connected........")
	newUser := model.UserDetails{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		FullName:  request.FirstName + request.LastName,
		Role:      request.Role,
		EmailId:   request.EmailId,
		Status:    "A",
		UserName:  request.EmailId,
	}

	dataJSON, err := json.Marshal(newUser)
	if err != nil {
		log.Println(err)
	}
	js := string(dataJSON)
	ind, err := elasticClient1.Index().
		Index(dbName).
		BodyJson(js).
		Do(context.Background())
	if err != nil {
		log.Println("Unable to insert :", err)
	}
	flushESDB(dbName)
	log.Println("Insertion Successful:", ind.Id)
	newUser.Id = ind.Id
	return newUser, nil
}

func flushESDB(indexname string) error {
	_, err := elasticClient1.Flush().Index(indexname).Do(context.TODO())
	return err
}

func FetchById(idStr, name string) ([]model.UserDetails, error) {
	var response []model.UserDetails
	log.Println("Creating connection........")
	_, err := GetESClient()
	if err != nil {
		log.Println("Error initializing : ", err)
		return response, errors.New("Client fail ")
	}
	log.Println("Connected........")

	searchSource := elastic.NewSearchSource()
	if idStr != "" {
		searchSource.Query(elastic.NewMatchQuery("_id", idStr))
	}
	if name != "" {
		searchSource.Query(elastic.NewMatchQuery("firstName", name))
	}
	queryStr, err1 := searchSource.Source()
	_, err2 := json.Marshal(queryStr)

	if err1 != nil || err2 != nil {
		log.Println("[esclient][GetResponse]err during query marshal=", err1, err2)
		return response, errors.New("err during query marshal")
	}
	searchResult, err := elasticClient1.Search().Index(dbName).SearchSource(searchSource).Do(context.Background())
	if err != nil {
		log.Println("[ProductsES][GetPIds]Error=", err)
		return response, err
	}
	log.Println("Search:", searchResult.Hits)
	for _, hit := range searchResult.Hits.Hits {
		var user model.UserDetails
		err := json.Unmarshal(hit.Source, &user)
		if err != nil {
			log.Println("[Getting Students][Unmarshal] Err=", err)
		}
		user.Id = hit.Id
		log.Println("User:", user)
		response = append(response, user)
	}
	log.Println("Response:", response)
	if len(response) == 0 {
		return response, errors.New("no records found")
	}
	return response, err
}
