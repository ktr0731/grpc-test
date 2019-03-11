package server

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/ktr0731/grpc-test/api"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (s *ExampleService) ClientStreaming(stm api.Example_ClientStreamingServer) error {
	var t int
	var names []string
	for {
		req, err := stm.Recv()
		if err == io.EOF {
			log.Println("end of client streaming")
			return stm.SendAndClose(&api.SimpleResponse{
				Message: fmt.Sprintf(`you sent requests %d times (%s).`, t, strings.Join(names, ", ")),
			})
		}
		if err != nil {
			return err
		}
		names = append(names, req.GetName())
		t++
		log.Println("client streaming request received")
	}
	return nil
}

func (s *ExampleService) ServerStreaming(req *api.SimpleRequest, stm api.Example_ServerStreamingServer) error {
	defer log.Println("end of server streaming")
	n := rand.Intn(5) + 1
	log.Printf("send %d times\n", n)
	for i := 0; i < n; i++ {
		log.Printf("send %d\n", i+1)
		err := stm.Send(&api.SimpleResponse{
			Message: fmt.Sprintf(`hello %s, I greet %d times.`, req.GetName(), i+1),
		})
		if err != nil {
			return err
		}
		time.Sleep(300 * time.Millisecond)
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
