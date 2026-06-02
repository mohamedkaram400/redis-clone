package main

import	"github.com/mohamedkaram400/redis-clone/resp"


var Handlers = map[string]func([]resp.Value) resp.Value{
	"PING": ping,
}


func ping(args []resp.Value) resp.Value {
	if len(args) == 0 {
		return resp.Value{Typ: "string", Str: "PONG"}
	}
	return resp.Value{Typ: "string", Str: args[0].Bulk}
}