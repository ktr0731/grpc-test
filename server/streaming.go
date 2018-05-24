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
			log.Println("end of client streaming")
			return stm.SendAndClose(&api.SimpleResponse{
				Message: fmt.Sprintf(`%s, you greet %d times.`, name, t),
			})
		}
		if err != nil {
			return err
		}
		name = req.GetName()
		t++
		log.Println("client streaming request received")
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

func (s *ExampleService) BidiStreaming(stm api.Example_BidiStreamingServer) error {
	for {
		req, err := stm.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		if err := s.ServerStreaming(req, stm); err != nil {
			return err
		}
	}
}
