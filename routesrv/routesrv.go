package main

import (
	"encoding/json"
	"github.com/celrenheit/lion"
	"github.com/celrenheit/lion/middleware"
	"github.com/jadengis/modserver/models"
	pb "github.com/jadengis/modserver/protoapi"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	"net/http"
	"strconv"
)

var (
	serverAddr    string = "localhost:9092"
	userDAOClient pb.UserDAOClient
)

// Lion modules
type Users struct{}

func (Users) Base() string {
	return "/users"
}

func (Users) Post(w http.ResponseWriter, r *http.Request) {
	user, err := decodeUser(r)
	if err != nil {
		http.Error(w, "Malformed input json:"+err.Error(), http.StatusUnprocessableEntity)
	}
	protoUser, err := models.NewUserProto(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}
	response, err := userDAOClient.SaveUser(context.Background(), protoUser)
	if err != nil {
		http.Error(w, "error saving user:"+err.Error(), http.StatusUnprocessableEntity)
	}
	user, err = models.NewUserModel(response.GetUser())
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
	return
}

func (Users) Routes(r *lion.Router) {
	// Defining a resource for getting, editing and deleting a single product
	r.Resource("/:id", OneUser{})
}

type OneUser struct{}

func (OneUser) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(lion.Param(r, "id"), 0, 0)
	if err != nil {
		http.Error(w, "bad path param:"+err.Error(), http.StatusUnprocessableEntity)
	}
	response, err := userDAOClient.GetUser(context.Background(), &pb.Id{id})
	if err != nil {
		http.Error(w, "error getting user:"+err.Error(), http.StatusUnprocessableEntity)
	}
	user, err := models.NewUserModel(response.GetUser())
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
	return
}

func (OneUser) Put(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(lion.Param(r, "id"), 0, 0)
	if err != nil {
		http.Error(w, "bad path param:"+err.Error(), http.StatusUnprocessableEntity)
	}
	user, err := decodeUser(r)
	if err != nil {
		http.Error(w, "Malformed input json:"+err.Error(), http.StatusUnprocessableEntity)
	}
	user.ID = uint(id)
	protoUser, err := models.NewUserProto(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}
	response, err := userDAOClient.UpdateUser(context.Background(), protoUser)
	if err != nil {
		http.Error(w, "error updating user:"+err.Error(), http.StatusUnprocessableEntity)
	}
	user, err = models.NewUserModel(response.GetUser())
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
	return
}

func (OneUser) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(lion.Param(r, "id"), 0, 0)
	if err != nil {
		http.Error(w, "bad path param:"+err.Error(), http.StatusUnprocessableEntity)
	}
	_, err = userDAOClient.DeleteUser(context.Background(), &pb.Id{id})
	if err != nil {
		http.Error(w, "error deleting user:"+err.Error(), http.StatusUnprocessableEntity)
	}
	w.WriteHeader(204)
}

func decodeUser(r *http.Request) (*models.User, error) {
	var user models.User
	dec := json.NewDecoder(r.Body)
	for {
		if err := dec.Decode(&user); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		//log.Printf("%s\n", t.Test)
	}
	return &user, nil
}

func main() {
	conn, err := dialConnection()
	if err != nil {
		log.Fatalf("could not connect to rpc server: error = %s", err.Error())
	}
	userDAOClient = pb.NewUserDAOClient(conn)

	l := lion.New(middleware.Classic())
	api := l.Group("/api")
	api.Module(Users{})
	l.Run(":8080")
}

func dialConnection() (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(serverAddr, opts...)
	return conn, err
}
