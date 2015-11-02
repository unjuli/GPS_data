package main

import (
  "fmt"
  "log"
  "gopkg.in/olivere/elastic.v2"
  "reflect"
  "strconv"
  // "encoding/json"
)

type GPSdata struct {
    Lat   float32
    Lng   float32
}

func ping_elastic(ElasticClient *elastic.Client) {
  info, code, err := ElasticClient.Ping().Do()
  if err != nil {
      log.Fatal(err)
  }
  fmt.Printf("Elasticsearch returned with code %d and version  %s\n", code, info.Version.Number)
}

func get_gps_data(client *elastic.Client, index string, car_id int) *elastic.SearchResult{
  searchResult, err := client.Search().
      Index("test").              // search in index "date"
      Type(strconv.Itoa(car_id)). // type as car_id
      Sort("_timestamp", false).  // sort by "message" field, ascending
      From(0).Size(10).           // take documents 0-9
      Pretty(true).               // pretty print request and response JSON
      Do()                        // execute
  if err != nil {
      log.Fatal(err)
  }
  fmt.Printf("Query took %d milliseconds, total hits: %s\n", searchResult.TookInMillis, searchResult.Hits.Hits[0].Source)
  return searchResult
}

func main() {
  client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
  if err != nil {
      log.Fatal(err)
  }
  go ping_elastic(client)
  go searchResult := get_gps_data(client, "test", 2)
  var ttyp GPSdata
  for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
      t := item.(GPSdata)
      fmt.Println("Found ", t.Lat, t.Lng)
  }    
}