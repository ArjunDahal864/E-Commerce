package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	recommendations "recommendations/proto"
)

type Recommendations struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Recommendations) Call(ctx context.Context, req *recommendations.Request, rsp *recommendations.Response) error {
	log.Info("Received Recommendations.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Recommendations) Stream(ctx context.Context, req *recommendations.StreamingRequest, stream recommendations.Recommendations_StreamStream) error {
	log.Infof("Received Recommendations.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&recommendations.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Recommendations) PingPong(ctx context.Context, stream recommendations.Recommendations_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&recommendations.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
