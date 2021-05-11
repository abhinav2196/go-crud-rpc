package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"net/http"
	"os"
	"rpc-tutorial/Item"
	"rpc-tutorial/rpc/basicCrudService"
)

func main() {
	client := basicCrudService.NewBasicCrudServiceProtobufClient("http://localhost:8085", &http.Client{})
	item := basicCrudService.Item{}
	item.Body = "UpdatedVal2"
	item.Title = "BITS"

	//addItemCall(client,item)
	updateItemCall(client,item)
	//deleteItemCall(client,item)
	getItemsCall(client)
}

func deleteItemCall(client basicCrudService.BasicCrudService, item basicCrudService.Item) {
	response, err := client.DeleteItem(context.Background(), &item)
	if err != nil {
		fmt.Printf("oh no: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Result after execution: %+v", response)
}

func updateItemCall(client basicCrudService.BasicCrudService, item basicCrudService.Item) {
	response, err := client.UpdateItem(context.Background(), &item)
	if err != nil {
		fmt.Printf("oh no: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Result after execution: %+v", response)
}

func addItemCall(client basicCrudService.BasicCrudService, item basicCrudService.Item) {
	response, err := client.AddItem(context.Background(), &item)
	if err != nil {
		fmt.Printf("oh no: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Result after execution: %+v", response)
}

func getItemsCall(client basicCrudService.BasicCrudService) {
	response, err := client.GetItems(context.Background(), &basicCrudService.NullVal{})
	if err != nil {
		fmt.Printf("oh no: %v", err)
		os.Exit(1)
	}
	var itemsFetched []Item.Item
	err = json.Unmarshal([]byte(response.Response), &itemsFetched)
	if err != nil {
		log.Panicln("Unmarshalling failed")
	}

	fmt.Printf("Result after execution: %+v", itemsFetched)
}

func toJson(pb proto.Message) []byte {
	//marshaler := jsonpb.Marshaler{}
	//out,err := marshaler.MarshalToString(pb)


	out,err := json.Marshal(pb)

	if err != nil {
		log.Fatalln("Can't convert to JSON",err)
		return nil
	}
	return out

}