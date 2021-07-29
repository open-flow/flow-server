package utils

//
//import (
//	"context"
//	"encoding/json"
//	"github.com/nats-io/nats.go"
//	"gitlab.com/yautoflow/interfaces/dtos"
//	"log"
//	"reflect"
//)
//
//func ConnectFactory(nc *nats.Conn) func(listen string, fn interface{}) {
//	return func(listen string, fn interface{}) {
//		reqType := reflect.TypeOf(fn).In(1).Elem()
//		fnValue := reflect.ValueOf(fn)
//
//		_, err := nc.Subscribe(listen, func(msg *nats.Msg) {
//			respond := func(val interface{}, err error) {
//				response := dtos.NatsResponse{}
//
//				if err == nil {
//					bytes, err := json.Marshal(val)
//					if err != nil {
//						log.Println("couldn't marshal response json", err)
//						response.Error = err.Error()
//					} else {
//						response.Message = bytes
//					}
//				} else {
//					log.Println("function returned error", err)
//					response.Error = err.Error()
//				}
//
//				bytes, err := json.Marshal(response)
//				if err != nil {
//					log.Println("couldn't marshal response", err)
//					return
//				}
//				err = msg.Respond(bytes)
//				if err != nil {
//					log.Println("couldn't respond", err)
//					return
//				}
//			}
//
//			reqValue := reflect.New(reqType)
//			err := json.Unmarshal(msg.Data, reqValue.Interface())
//			if err != nil {
//				log.Println("couldn't unmarshall message")
//				respond(nil, err)
//				return
//			}
//			args := []reflect.Value{reflect.ValueOf(context.Background()), reqValue}
//			resArr := fnValue.Call(args)
//			resVal := resArr[0]
//			errVal := resArr[1]
//
//			respond(resVal.Interface(), errVal.Interface().(error))
//		})
//
//		if err != nil {
//			panic(err)
//		}
//	}
//}
