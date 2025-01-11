Leaf server
===========
A game server based on [Leaf framework](https://github.com/nlmayday/nlleaf).

Licensing
---------

Leaf server is licensed under the Apache License, Version 2.0. See [LICENSE](https://github.com/nlmayday/nlleafserver/blob/master/LICENSE) for the full license text.

游戏平台
    http api
    ws/tcp 客户端
    grpc <-> 子游戏 

子游戏
    游戏逻辑
    grpc <-> 游戏平台

===========

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative msgproto/serviceagent.proto