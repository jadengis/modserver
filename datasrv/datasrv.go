package main

import (
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/jadengis/modserver/duid"
	"github.com/jadengis/modserver/models"
	pb "github.com/jadengis/modserver/protoapi"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	host     string = "localhost"
	user     string = "postgres"
	password string = "postgres"
	dbname   string = "test"
)

var db *gorm.DB

func main() {
	models.SetIdGenerator(duid.NewGenerator(0))
	dbConfig := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		host, user, dbname, password)

	var err error
	db, err = gorm.Open("postgres", dbConfig)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if db.DB().Ping() != nil {
		panic(err)
	}
	log.Println("connected to db")

	// Migrate schema
	db.AutoMigrate(&models.User{})
	log.Println("migration successful")

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9092))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("preparing to listen and serve")
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterUserDAOServer(grpcServer, &userDAOServer{})
	grpcServer.Serve(lis)
}

type userDAOServer struct{}

func (u *userDAOServer) SaveUser(ctx context.Context, user *pb.User) (*pb.SaveResponse, error) {
	newUser := models.NewUser(user.GetName(), user.GetEmail(), int(user.GetAge()))
	db.Create(newUser)

	resultUser, err := NewUserProto(newUser)
	if err != nil {
		return nil, err
	}

	response := &pb.SaveResponse{resultUser}
	return response, nil
}

func (u *userDAOServer) GetUser(ctx context.Context, userId *pb.Id) (*pb.UserResponse, error) {
	var user models.User
	db.First(&user, uint(userId.GetValue()))

	resultUser, err := NewUserProto(&user)
	if err != nil {
		return nil, err
	}

	response := &pb.UserResponse{resultUser}
	return response, nil
}

func (u *userDAOServer) UpdateUser(ctx context.Context, user *pb.User) (*pb.UpdateResponse, error) {
	updateUser, err := NewUserModel(user)
	if err != nil {
		return nil, err
	}
	db.Update(&updateUser)
	resultUser, err := NewUserProto(updateUser)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateResponse{resultUser}, nil
}

func (u *userDAOServer) DeleteUser(ctx context.Context, userId *pb.Id) (*pb.DeleteResponse, error) {
	user := &models.User{}
	user.ID = uint(userId.GetValue())
	db.Delete(user)
	return &pb.DeleteResponse{}, nil
}

func NewUserProto(user *models.User) (*pb.User, error) {
	resultUser := &pb.User{}
	var err error

	resultUser.Id = &pb.Id{uint64(user.ID)}
	resultUser.Age = int64(user.Age)
	resultUser.CreatedAt, err = ptypes.TimestampProto(user.CreatedAt)
	if err != nil {
		return nil, err
	}
	resultUser.UpdatedAt, err = ptypes.TimestampProto(user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if user.DeletedAt == nil {
		resultUser.DeletedAt = nil
	} else {
		resultUser.DeletedAt, err = ptypes.TimestampProto(*user.DeletedAt)
		if err != nil {
			return nil, err
		}
	}
	resultUser.Name = user.Name
	resultUser.Email = user.Email
	return resultUser, nil
}

func NewUserModel(user *pb.User) (*models.User, error) {
	resultUser := &models.User{}
	var err error

	resultUser.ID = uint(user.GetId().GetValue())
	resultUser.Age = int(user.GetAge())
	resultUser.CreatedAt, err = ptypes.Timestamp(user.GetCreatedAt())
	if err != nil {
		return nil, err
	}
	resultUser.UpdatedAt, err = ptypes.Timestamp(user.GetUpdatedAt())
	if err != nil {
		return nil, err
	}
	deletedAt, err := ptypes.Timestamp(user.GetDeletedAt())
	if err != nil {
		return nil, err
	}
	resultUser.DeletedAt = &deletedAt
	resultUser.Name = user.Name
	resultUser.Email = user.Email
	return resultUser, nil
}
