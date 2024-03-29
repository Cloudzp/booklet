package main

import (
	"protocolBuffers/example"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

func main(){
	test := &example.Test{
		Label:proto.String("Hello"),
		Type: proto.Int32(17),
		Reps: []int64{1,2,3},
	}

	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("Marshal", err)
	}

	newTest := example.Test{}
	err = proto.Unmarshal(data, &newTest)
	if err != nil {
		log.Fatal("Unmarshal:",err)
	}
	// Now test and newTest contain the same data.
	if test.GetLabel() != newTest.GetLabel() {
		log.Fatalf("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
	}else {
		fmt.Println(newTest.GetLabel(),newTest.GetType(),newTest.GetReps())
	}
}
