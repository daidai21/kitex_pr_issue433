namespace go api

struct Request {
    1: required list<string>  messages
}

struct Response {
    1: required string message
}

service Echo {
    Response echo(1: Request req)
}
