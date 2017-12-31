# modserver
A simple modular server comprised of an HTTP routing node (which could house business logic) and a data node that acts as
gatekeeper to the underlying data store. HTTP routes are implemented using the lion framework for Go, object relation management is handled with gorm, and these to servers communicate using a grpc service layer.
