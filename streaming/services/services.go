package services

import (
	"context"
	"io"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	pb "github.com/hqd888/grpc-examples/streaming/services/pb"
)

type UserService struct{}

// 普通调用
func (u *UserService) Login(ctx context.Context, request *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	var phones = []*pb.UserLoginResponse_PhoneNumber{
		&pb.UserLoginResponse_PhoneNumber{Number: "1111111", Type: pb.UserLoginResponse_PhoneType(1)},
		&pb.UserLoginResponse_PhoneNumber{Number: "2222222", Type: pb.UserLoginResponse_PhoneType(2)},
	}
	var lastUpdateDate = &timestamp.Timestamp{
		Seconds: time.Now().Unix(),
		Nanos:   int32(999999999),
	}
	resp := &pb.UserLoginResponse{
		UserId:         1000,
		UserName:       request.GetUserName(),
		UserNickname:   "hqd888",
		UserCountry:    "china",
		UserGender:     1,
		UserCredits:    20000,
		IsAdmin:        true,
		AuthCode:       []byte("xxxxxx"),
		UserBalance:    20000.222,
		UserHobby:      0,
		Phones:         phones,
		LoginStatus:    pb.LoginStatus_success,
		LastUpdateDate: lastUpdateDate,
	}
	return resp, nil
}

// 请求使用单向流
func (u *UserService) LoginClientStreaming(us pb.UserService_LoginClientStreamingServer) error {
	var phones = []*pb.UserLoginResponse_PhoneNumber{
		&pb.UserLoginResponse_PhoneNumber{Number: "1111111", Type: pb.UserLoginResponse_PhoneType(1)},
		&pb.UserLoginResponse_PhoneNumber{Number: "2222222", Type: pb.UserLoginResponse_PhoneType(2)},
	}
	var lastUpdateDate = &timestamp.Timestamp{
		Seconds: time.Now().Unix(),
		Nanos:   int32(999999999),
	}
	for {
		req, err := us.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		resp := &pb.UserLoginResponse{
			UserId:         1001,
			UserName:       req.GetUserName(),
			UserNickname:   "hqd888",
			UserCountry:    "china",
			UserGender:     1,
			UserCredits:    30000,
			IsAdmin:        true,
			AuthCode:       []byte("abd"),
			UserBalance:    30000.333,
			UserHobby:      1,
			Phones:         phones,
			LoginStatus:    pb.LoginStatus_success,
			LastUpdateDate: lastUpdateDate,
		}
		return us.SendAndClose(resp)
	}
}

// 响应使用单向流
func (u *UserService) LoginServerStreaming(req *pb.UserLoginRequest, us pb.UserService_LoginServerStreamingServer) error {
	var phones = []*pb.UserLoginResponse_PhoneNumber{
		&pb.UserLoginResponse_PhoneNumber{Number: "1111111", Type: pb.UserLoginResponse_PhoneType(1)},
		&pb.UserLoginResponse_PhoneNumber{Number: "2222222", Type: pb.UserLoginResponse_PhoneType(2)},
	}
	var lastUpdateDate = &timestamp.Timestamp{
		Seconds: time.Now().Unix(),
		Nanos:   int32(999999999),
	}
	resp := &pb.UserLoginResponse{
		UserId:         1002,
		UserName:       req.GetUserName(),
		UserNickname:   "hqd888",
		UserCountry:    "china",
		UserGender:     1,
		UserCredits:    40000,
		IsAdmin:        true,
		AuthCode:       []byte("abd"),
		UserBalance:    40000.555,
		UserHobby:      2,
		Phones:         phones,
		LoginStatus:    pb.LoginStatus_success,
		LastUpdateDate: lastUpdateDate,
	}
	return us.Send(resp)
}

// 使用双向流
func (u *UserService) LoginStreaming(stream pb.UserService_LoginStreamingServer) error {
	var phones = []*pb.UserLoginResponse_PhoneNumber{
		&pb.UserLoginResponse_PhoneNumber{Number: "1111111", Type: pb.UserLoginResponse_PhoneType(1)},
		&pb.UserLoginResponse_PhoneNumber{Number: "2222222", Type: pb.UserLoginResponse_PhoneType(2)},
	}
	var lastUpdateDate = &timestamp.Timestamp{
		Seconds: time.Now().Unix(),
		Nanos:   int32(999999999),
	}
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		resp := &pb.UserLoginResponse{
			UserId:         1003,
			UserName:       req.GetUserName(),
			UserNickname:   "hqd888",
			UserCountry:    "china",
			UserGender:     1,
			UserCredits:    50000,
			IsAdmin:        true,
			AuthCode:       []byte("abd"),
			UserBalance:    5555555.555,
			UserHobby:      2,
			Phones:         phones,
			LoginStatus:    pb.LoginStatus_success,
			LastUpdateDate: lastUpdateDate,
		}

		err = stream.Send(resp)
		if err != nil {
			return err
		}
		return nil
	}
}
