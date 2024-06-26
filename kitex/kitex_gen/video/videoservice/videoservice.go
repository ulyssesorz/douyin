// Code generated by Kitex v0.4.4. DO NOT EDIT.

package videoservice

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	video "github.com/ulyssesorz/douyin/kitex/kitex_gen/video"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return videoServiceServiceInfo
}

var videoServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "VideoService"
	handlerType := (*video.VideoService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Feed":          kitex.NewMethodInfo(feedHandler, newFeedArgs, newFeedResult, false),
		"PublishAction": kitex.NewMethodInfo(publishActionHandler, newPublishActionArgs, newPublishActionResult, false),
		"PublishList":   kitex.NewMethodInfo(publishListHandler, newPublishListArgs, newPublishListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "video",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func feedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.FeedRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.VideoService).Feed(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FeedArgs:
		success, err := handler.(video.VideoService).Feed(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FeedResult)
		realResult.Success = success
	}
	return nil
}
func newFeedArgs() interface{} {
	return &FeedArgs{}
}

func newFeedResult() interface{} {
	return &FeedResult{}
}

type FeedArgs struct {
	Req *video.FeedRequest
}

func (p *FeedArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(video.FeedRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FeedArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FeedArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FeedArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FeedArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FeedArgs) Unmarshal(in []byte) error {
	msg := new(video.FeedRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FeedArgs_Req_DEFAULT *video.FeedRequest

func (p *FeedArgs) GetReq() *video.FeedRequest {
	if !p.IsSetReq() {
		return FeedArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FeedArgs) IsSetReq() bool {
	return p.Req != nil
}

type FeedResult struct {
	Success *video.FeedResponse
}

var FeedResult_Success_DEFAULT *video.FeedResponse

func (p *FeedResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(video.FeedResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FeedResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FeedResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FeedResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FeedResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FeedResult) Unmarshal(in []byte) error {
	msg := new(video.FeedResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FeedResult) GetSuccess() *video.FeedResponse {
	if !p.IsSetSuccess() {
		return FeedResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FeedResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.FeedResponse)
}

func (p *FeedResult) IsSetSuccess() bool {
	return p.Success != nil
}

func publishActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.PublishActionRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.VideoService).PublishAction(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *PublishActionArgs:
		success, err := handler.(video.VideoService).PublishAction(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PublishActionResult)
		realResult.Success = success
	}
	return nil
}
func newPublishActionArgs() interface{} {
	return &PublishActionArgs{}
}

func newPublishActionResult() interface{} {
	return &PublishActionResult{}
}

type PublishActionArgs struct {
	Req *video.PublishActionRequest
}

func (p *PublishActionArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(video.PublishActionRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *PublishActionArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *PublishActionArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *PublishActionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PublishActionArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *PublishActionArgs) Unmarshal(in []byte) error {
	msg := new(video.PublishActionRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PublishActionArgs_Req_DEFAULT *video.PublishActionRequest

func (p *PublishActionArgs) GetReq() *video.PublishActionRequest {
	if !p.IsSetReq() {
		return PublishActionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PublishActionArgs) IsSetReq() bool {
	return p.Req != nil
}

type PublishActionResult struct {
	Success *video.PublishActionResponse
}

var PublishActionResult_Success_DEFAULT *video.PublishActionResponse

func (p *PublishActionResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(video.PublishActionResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *PublishActionResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *PublishActionResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *PublishActionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PublishActionResult")
	}
	return proto.Marshal(p.Success)
}

func (p *PublishActionResult) Unmarshal(in []byte) error {
	msg := new(video.PublishActionResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PublishActionResult) GetSuccess() *video.PublishActionResponse {
	if !p.IsSetSuccess() {
		return PublishActionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PublishActionResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.PublishActionResponse)
}

func (p *PublishActionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func publishListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.PublishListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.VideoService).PublishList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *PublishListArgs:
		success, err := handler.(video.VideoService).PublishList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PublishListResult)
		realResult.Success = success
	}
	return nil
}
func newPublishListArgs() interface{} {
	return &PublishListArgs{}
}

func newPublishListResult() interface{} {
	return &PublishListResult{}
}

type PublishListArgs struct {
	Req *video.PublishListRequest
}

func (p *PublishListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(video.PublishListRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *PublishListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *PublishListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *PublishListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PublishListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *PublishListArgs) Unmarshal(in []byte) error {
	msg := new(video.PublishListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PublishListArgs_Req_DEFAULT *video.PublishListRequest

func (p *PublishListArgs) GetReq() *video.PublishListRequest {
	if !p.IsSetReq() {
		return PublishListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PublishListArgs) IsSetReq() bool {
	return p.Req != nil
}

type PublishListResult struct {
	Success *video.PublishListResponse
}

var PublishListResult_Success_DEFAULT *video.PublishListResponse

func (p *PublishListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(video.PublishListResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *PublishListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *PublishListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *PublishListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PublishListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *PublishListResult) Unmarshal(in []byte) error {
	msg := new(video.PublishListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PublishListResult) GetSuccess() *video.PublishListResponse {
	if !p.IsSetSuccess() {
		return PublishListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PublishListResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.PublishListResponse)
}

func (p *PublishListResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Feed(ctx context.Context, Req *video.FeedRequest) (r *video.FeedResponse, err error) {
	var _args FeedArgs
	_args.Req = Req
	var _result FeedResult
	if err = p.c.Call(ctx, "Feed", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PublishAction(ctx context.Context, Req *video.PublishActionRequest) (r *video.PublishActionResponse, err error) {
	var _args PublishActionArgs
	_args.Req = Req
	var _result PublishActionResult
	if err = p.c.Call(ctx, "PublishAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PublishList(ctx context.Context, Req *video.PublishListRequest) (r *video.PublishListResponse, err error) {
	var _args PublishListArgs
	_args.Req = Req
	var _result PublishListResult
	if err = p.c.Call(ctx, "PublishList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
