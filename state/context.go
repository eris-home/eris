package state

import "fmt"

type Context struct {
	ActorID   string
	SessionID string
	Data      map[string]interface{}
}

func NewContext(actorID, sessionID string) *Context {
	return &Context{
		ActorID:   actorID,
		SessionID: sessionID,
		Data:      make(map[string]interface{}),
	}
}

func (ctx *Context) Set(key string, value interface{}) {
	ctx.Data[key] = value
	fmt.Printf("Context updated: %s = %v\n", key, value)
}

func (ctx *Context) Get(key string) (interface{}, bool) {
	value, exists := ctx.Data[key]
	return value, exists
}

func (ctx *Context) Delete(key string) {
	delete(ctx.Data, key)
	fmt.Printf("Context key deleted: %s\n", key)
}

func (ctx *Context) Reset() {
	ctx.Data = make(map[string]interface{})
	fmt.Printf("Context reset for Session: %s\n", ctx.SessionID)
}