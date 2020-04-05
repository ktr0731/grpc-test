package server

import (
	"fmt"
	"io"
	"math/rand"
	"strings"
	"time"

	"github.com/ktr0731/grpc-test/api"
	"google.golang.org/grpc/metadata"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (s *ExampleService) ClientStreaming(stm api.Example_ClientStreamingServer) error {
	var t int
	var names []string
	stm.SendHeader(metadata.New(map[string]string{"header_key1": "header_val1", "header_key2": "header_val2"}))
	stm.SetTrailer(metadata.New(map[string]string{"trailer_key1": "trailer_val1", "trailer_key2": "trailer_val2"}))
	for {
		req, err := stm.Recv()
		if err == io.EOF {
			s.Logger.Println("end of client streaming")
			return stm.SendAndClose(&api.SimpleResponse{
				Message: fmt.Sprintf(`you sent requests %d times (%s).`, t, strings.Join(names, ", ")),
			})
		}
		if err != nil {
			return err
		}
		names = append(names, req.GetName())
		t++
		s.Logger.Println("client streaming request received")
	}
	return nil
}

func (s *ExampleService) ServerStreaming(req *api.SimpleRequest, stm api.Example_ServerStreamingServer) error {
	defer s.Logger.Println("end of server streaming")
	if err := stm.SendHeader(metadata.New(map[string]string{"ss_header_key1": "header_val1", "ss_header_key2": "header_val2"})); err != nil {
		return err
	}
	stm.SetTrailer(metadata.New(map[string]string{"ss_trailer_key1": "trailer_val1", "ss_trailer_key2": "trailer_val2"}))
	if err := s.serverStreaming(req, stm); err != nil {
		return err
	}
	return nil
}

func (s *ExampleService) serverStreaming(req *api.SimpleRequest, stm api.Example_ServerStreamingServer) error {
	n := 3
	s.Logger.Printf("send %d times\n", n)
	for i := 0; i < n; i++ {
		s.Logger.Printf("send %d\n", i+1)
		err := stm.Send(&api.SimpleResponse{
			Message: fmt.Sprintf(`hello %s, I greet %d times.`, req.GetName(), i+1),
		})
		if err != nil {
			return err
		}
		time.Sleep(50 * time.Millisecond)
	}
	return nil
}

func (s *ExampleService) BidiStreaming(stm api.Example_BidiStreamingServer) error {
	if err := stm.SendHeader(metadata.New(map[string]string{"header_key1": "header_val1", "header_key2": "header_val2"})); err != nil {
		return err
	}
	stm.SetTrailer(metadata.New(map[string]string{"trailer_key1": "trailer_val1", "trailer_key2": "trailer_val2"}))
	for {
		req, err := stm.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		if err := s.serverStreaming(req, stm); err != nil {
			return err
		}
	}
}
