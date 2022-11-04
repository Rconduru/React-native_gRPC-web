package alert

import (
	"log"

	alertPb "github.com/Rconduru/React-native_gRPC-web/server/gen"

	"golang.org/x/net/context"
)

type Server struct {
	alertPb.UnimplementedAlertServiceServer
}

func (s *Server) SendAlert(ctx context.Context, in *alertPb.SendAlertRequest) (*alertPb.SendAlertResponse, error) {
	log.Printf("Receive message body from client: %s", in.Notification)
	return &alertPb.SendAlertResponse{Id: "Simbora meu povo!"}, nil
}
