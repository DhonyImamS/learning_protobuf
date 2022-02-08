package main

import (
    "fmt"
    "bytes"
    "os"
    "learn_protobuf/model"
    "github.com/golang/protobuf/jsonpb"
)

func main() {
  var user1 = &model.User{
    Id: "u001",
    Name: "Dhony Imam Saputra",
    Password: "password123",
    Gender: model.UserGender_MALE,
  }

  var userList = &model.UserList{
     List: []*model.User{
         user1,
     },
  }

  var garage1 = &model.Garage{
    Id: "g001",
    Name: "Auto2000 BSD",
    Coordinate: &model.GarageCoordinate {
      Latitude: 0,
      Longitude: 0,
    },
  }

 var garageList = &model.GarageList{
     List: []*model.Garage{
         garage1,
     },
 }
 
 var garageListByUser = &model.GarageListByUser{
    List: map[string]*model.GarageList{
      user1.Id: garageList,
    },
 }

 _ = garageListByUser
 _ = userList

 // =========== original
 fmt.Printf("# ==== Original\n       %#v \n", user1)

 // =========== as string
 fmt.Printf("# ==== As String\n       %v \n", user1.String())

 // =========== as json string
  var buf bytes.Buffer
  err1 := (&jsonpb.Marshaler{}).Marshal(&buf, garageList)
  if err1 != nil {
      fmt.Println(err1.Error())
      os.Exit(0)
  }
  jsonString := buf.String()
  fmt.Printf("# ==== As JSON String\n       %v \n", jsonString)

}
