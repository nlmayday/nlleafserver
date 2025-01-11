package internal

import (
	"context"
	"fmt"
	"io"
	"net"
	"server/base"
	"server/global"
	"server/model"
	"server/protocol"
	protojson "server/protocol/json"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	pb "server/protocol/msgproto"
)

type GameServe struct {
	GameId uint64
	Addr   string
	Name   string

	Conn     *grpc.ClientConn
	Stream   pb.Serviceagent_ChannelClient
	SgClient pb.ServiceagentClient
	Status   int // 状态
}

func (g *GameServe) LoadGame(game *model.GameServers) error {
	g.GameId = game.ID
	g.Addr = game.Addr
	g.Name = game.Name
	g.Status = protocol.GAME_STATUS_NONE
	g.Connect()
	return nil
}

func (g *GameServe) Connect() error {
	if g.Status == protocol.GAME_STATUS_LOADING {
		return nil
	}
	if g.Conn != nil {
		global.G_LOG.Warn("重连，断开老的连接", zap.String("addr", g.Addr))
		g.Close()
		// return nil
	}

	g.Status = protocol.GAME_STATUS_LOADING

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 拨号选项
	opts := []grpc.DialOption{
		grpc.WithInsecure(), // 禁用 TLS
		grpc.WithBlock(),    // 阻塞直到连接成功
		grpc.WithContextDialer(func(ctx context.Context, addr string) (net.Conn, error) {
			return (&net.Dialer{}).DialContext(ctx, "tcp", addr)
		}),
	}

	// 建立连接
	conn, err := grpc.DialContext(ctx, g.Addr, opts...)
	if err != nil {
		global.G_LOG.Error("连接失败", zap.Error(err), zap.String("addr", g.Addr))
		return err
	}

	c := pb.NewServiceagentClient(conn)
	// Channel
	stream, err := c.Channel(context.Background())
	if err != nil {
		global.G_LOG.Error("连接失败", zap.Error(err), zap.String("addr", g.Addr))
		return err
	}

	g.SgClient = c
	g.Stream = stream
	g.Conn = conn
	global.G_LOG.Info("连接成功", zap.String("addr", g.Addr))
	g.Status = protocol.GAME_STATUS_RUNNING
	return nil
}

func (g *GameServe) Reconnect() error {
	if g.Status != protocol.GAME_STATUS_CLOSE {
		return nil
	}

	if err := g.Connect(); err != nil {
		return err
	}

	// g.Run()
	return nil
}

func (g *GameServe) Close() error {
	if g.Conn != nil {
		g.Conn.Close()
		g.Conn = nil
		g.Stream = nil
		g.SgClient = nil
	}
	return nil
}

func (g *GameServe) Send(data []byte) error {
	if g.Status == protocol.GAME_STATUS_RUNNING && g.Conn != nil {
		if err := g.Stream.Send(&pb.CommonMessage{Body: "Hello from client"}); err != nil {
			// log.Fatalf("Failed to send message: %v", err)
			global.G_LOG.Error("发送消息失败", zap.Error(err))
		}
	}
	return fmt.Errorf("连接已断开,当前状态:%d", g.Status)
}

func (g *GameServe) Recv() ([]byte, error) {
	if g.Status == protocol.GAME_STATUS_RUNNING && g.Conn != nil {
		in, err := g.Stream.Recv()
		if err == io.EOF {
			global.G_LOG.Info("连接已断开", zap.String("addr", g.Addr))
			g.Status = protocol.GAME_STATUS_CLOSE
			return nil, fmt.Errorf("连接已断开,当前状态:%d", g.Status)
		}
		if err != nil {
			// log.Fatalf("Failed to receive message: %v", err)
			global.G_LOG.Error("接收消息失败", zap.Error(err))
			g.Status = protocol.GAME_STATUS_CLOSE
			return nil, fmt.Errorf("连接已断开,当前状态:%d", g.Status)
		}
		// log.Printf("Got message: %v", in)
		global.G_LOG.Info("收到消息", zap.String("data", in.Body))
		return []byte(in.Body), nil
	}
	return nil, fmt.Errorf("连接已断开,当前状态:%d", g.Status)
}

// 心跳
func (g *GameServe) Heartbeat() error {
	if g.Status == protocol.GAME_STATUS_RUNNING && g.Conn != nil {
		hb := &pb.HeartbeatMessage{}
		hb.Times = uint64(time.Now().Unix())
		ret, err := g.SgClient.Heartbeat(context.Background(), hb)
		if err != nil {
			global.G_LOG.Error("heart", zap.Error(err))
			g.Status = protocol.GAME_STATUS_CLOSE
			return err
		}
		global.G_LOG.Debug("heart bk", zap.Any("times", ret.Times))
	}
	return nil
}

func (g *GameServe) Run() {
	// 心跳 3秒
	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()

	// 重连时间 10秒
	ticker2 := time.NewTicker(time.Second * 10)
	defer ticker2.Stop()

	go g.Read()
	for {
		// 3秒中发送一个心跳
		select {
		case <-ticker.C:
			g.Heartbeat()
		case <-ticker2.C:
			g.Reconnect()
		}
	}
}

func (g *GameServe) Destroy() {
	g.Close()
}

// 读取消息
func (g *GameServe) Read() {
	for {

		if g.Status != protocol.GAME_STATUS_RUNNING {
			return
		}
		data, err := g.Recv()
		if err != nil {
			global.G_LOG.Error("读取消息失败", zap.Error(err))
			g.Status = protocol.GAME_STATUS_CLOSE
			continue
		}
		global.G_LOG.Info("收到消息", zap.ByteString("data", data))
	}
}

// 加入游戏
func (g *GameServe) JoinGameReq(userInfo *base.UserInfo, m *protojson.JoinGameReq) {
	global.G_LOG.Info("加入游戏", zap.Uint64("gameId", m.GameId))

	if g.Status == protocol.GAME_STATUS_RUNNING && g.Conn != nil {
		hb := &pb.JoinMessage{}
		hb.Uid = userInfo.Uid
		ret, err := g.SgClient.Join(context.Background(), hb)
		if err != nil {
			global.G_LOG.Error("join", zap.Error(err))
		}
		global.G_LOG.Debug("join bk", zap.Any("times", ret.Uid))
	}
}

// 离开游戏
func (g *GameServe) LeaveGameReq(userInfo *base.UserInfo, m *protojson.LeaveGameReq) {
	global.G_LOG.Info("离开游戏", zap.Uint64("gameId", m.GameId))

	if g.Status == protocol.GAME_STATUS_RUNNING && g.Conn != nil {
		hb := &pb.LeaveMessage{}
		hb.Uid = userInfo.Uid
		ret, err := g.SgClient.Leave(context.Background(), hb)
		if err != nil {
			global.G_LOG.Error("leave", zap.Error(err))
		}
		global.G_LOG.Debug("leave bk", zap.Any("times", ret.Uid))
	}
}

// 进入游戏
