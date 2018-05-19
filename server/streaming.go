package server

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"

	"github.com/ktr0731/grpc-test/api"
)

func (s *ExampleService) ClientStreaming(stm api.Example_ClientStreamingServer) error {
	var t int
	var name string
	for {
		req, err := stm.Recv()
		if err == io.EOF {
			return stm.SendAndClose(&api.SimpleResponse{
				Message: fmt.Sprintf(`%s, you greet %d times.`, name, t),
			})
		}
		if err != nil {
			return err
		}
		name = req.GetName()
		t++
	}
	return nil
}

func (s *ExampleService) ServerStreaming(req *api.SimpleRequest, stm api.Example_ServerStreamingServer) error {
	n := rand.Intn(10)
	for i := 0; i < n; i++ {
		log.Printf("send %d\n", i+1)
		err := stm.Send(&api.SimpleResponse{
			Message: fmt.Sprintf(`hello %s, I greet %d times.`, req.GetName(), i),
		})
		time.Sleep(1 * time.Second)
		if err != nil {
			return err
		}
	}
	return nil
}
