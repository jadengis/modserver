package tests

import (
	pb "github.com/jadengis/modserver/protoapi"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"testing"
)

var (
	serverAddr string = "localhost:9092"
)

func TestUserWorkflow(t *testing.T) {
	conn, err := dialConnection()
	if err != nil {
		t.Fatalf("could not connect to rpc server: error = %s", err.Error())
	}
	client := pb.NewUserDAOClient(conn)
	user := &pb.User{Name: "Billy", Email: "billy@billy.billy", Age: 400}
	response, err := client.SaveUser(context.Background(), user)
	if err != nil {
		t.FailNow()
	}
	defer client.DeleteUser(context.Background(), response.GetUser().GetId())

	getResponse, err := client.GetUser(context.Background(), response.GetUser().GetId())
	if err != nil {
		t.FailNow()
	}
	if response.User.Id.Value != getResponse.User.Id.Value {
		t.Errorf("The id value is effed: out: %d != in: %d", response.User.Id.Value, getResponse.User.Id.Value)
	}

}

func dialConnection() (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(serverAddr, opts...)
	return conn, err
}
