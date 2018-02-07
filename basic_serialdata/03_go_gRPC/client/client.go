package mainb

import (
	"fmt"
	"log"

	proto "github.com/muhfaris/goFun/basic_serialdata/03_go_gRPC/server/kittens"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect to server", err)
	}

	client := proto.NewKittensClient(conn)
	response, err := client.Hello(context.Background(), &proto.Request{Name: "Faris"})

	if err != nil {
		log.Fatal("Error calling service:", err)
	}

	fmt.Println(response.Name)
}
