package basicCrudServiceServer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"rpc-tutorial/Item"
	"rpc-tutorial/database"
	pb "rpc-tutorial/rpc/basicCrudService"
)

// Server implements the Haberdasher service
type Server struct {}

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

func modelToJson(items []Item.Item) []byte {
	//marshaler := jsonpb.Marshaler{}
	//out,err := marshaler.MarshalToString(pb)


	out,err := json.Marshal(items)

	if err != nil {
		log.Fatalln("Can't convert to JSON",err)
		return nil
	}
	return out
}

//func toItemFromJson(in string){
//:= Item.Item{}
//	jsonpb.UnmarshalString(in)
//}

func (s *Server) AddItem(ctx context.Context, item *pb.Item) (responseMsg *pb.ResponseMsg, err error) {

	jsonReq := toJson(item)
	fmt.Println("Here's marshalled output string of item passed",string(jsonReq))

	itemToBeCreated := Item.Item{}
	err = json.Unmarshal(jsonReq, &itemToBeCreated)
	if err != nil {
		log.Panicln("Unmarshalling failed")
	}

	fmt.Println("Here's unmarshalled output string of item passed",itemToBeCreated)

	db := database.DBConn
	//var items []Item.Item
	db.Create(&itemToBeCreated)
	//db.Find(&items)

	return &pb.ResponseMsg{
		Response : "Success",
	}, nil
}

func convertPbMsgToModel(item *pb.Item) (Item.Item,error) {
	jsonReq := toJson(item)
	fmt.Println("Here's marshalled output string of item passed",string(jsonReq))

	itemToBeCreated := Item.Item{}
	err := json.Unmarshal(jsonReq, &itemToBeCreated)
	if err != nil {
		log.Panicln("Unmarshalling failed")
		return itemToBeCreated,err
	}
	return itemToBeCreated, nil
}

func convertModelToPbMsg(items []Item.Item) (*pb.Items,error) {
	jsonResult := modelToJson(items)
	fmt.Println("Here's marshalled output string of item passed to convert to message",string(jsonResult))

	var convertedMessage *pb.Items
	err := json.Unmarshal(jsonResult, &convertedMessage)
	// I guess the unmarshalling failed due to message having less fields than original struct i.e gorm.model fields
	if err != nil {
		log.Panicln("Unmarshalling failed")
		return convertedMessage,err
	}
	return convertedMessage, nil
}

func (s *Server) DeleteItem(ctx context.Context, item *pb.Item) (responseMsg *pb.ResponseMsg, err error) {

	itemTobeDeleted,err := convertPbMsgToModel(item)

	if err != nil {
		log.Fatalln("Conversion failed from proto message to model struct")
	}

	db := database.DBConn
	db.Where("Title = ?", itemTobeDeleted.Title).Find(&itemTobeDeleted)

	fmt.Println("Item to be deleted= ",itemTobeDeleted)

	//var items []Item.Item
	db.Delete(&itemTobeDeleted)
	//db.Find(&items)

	return &pb.ResponseMsg{
		Response : "Successfully deleted",
	}, nil
}

func (s *Server) UpdateItem(ctx context.Context, item *pb.Item) (responseMsg *pb.ResponseMsg, err error) {

	//changedItem,err := convertPbMsgToModel(item)
	var itemTobeUpdated Item.Item

	//if err != nil {
	//	log.Fatalln("Conversion failed from proto message to model struct")
	//}

	db := database.DBConn
	db.Where("Title = ?", item.Title).Find(&itemTobeUpdated)
	//itemTobeUpdated ,err = convertPbMsgToModel(item)
	//if err != nil {
	//	log.Fatalln("Conversion failed from proto message to model struct")
	//}
	itemTobeUpdated.Body = item.Body
	fmt.Println("Item to be updated= ", itemTobeUpdated)

	//var items []Item.Item
	db.Save(&itemTobeUpdated)
	//db.Find(&items)

	return &pb.ResponseMsg{
		Response : "Successfully Updated",
	}, nil
}

func (s *Server) GetItems(ctx context.Context, nullVal  *pb.NullVal) (*pb.ResponseMsg, error) {
	var items []Item.Item
	db := database.DBConn
	//var items []Item.Item
	db.Find(&items)

	jsonOut, err := json.Marshal(items)



	//response,err := convertModelToPbMsg(items)

	if err != nil {
		log.Fatalln("Model to Pb Msg conversion failed")
	}


	//fmt.Println("Json Output String",string(jsonOut))

	fmt.Println("Item fetched= ",items)

	return &pb.ResponseMsg{
		Response : string(jsonOut),
	}, nil
}


//
//func (s *Server) GetItems(ctx context.Context, null *pb.NullVal) (responseMsg *pb.Items, err error) {
//
//	db := database.DBConn
//	var items []pb.Item
//	db.Find(&items)
//
//	return &pb.Items{
//		Items : items,
//	}, nil
//}
