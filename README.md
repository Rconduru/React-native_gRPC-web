# Poc Arch GRPC

## Config Environment

Para a geração do código de forma automática baseado nos arquivos **.proto** serão necessários a instalação e configuração das variáveis de ambiente dos seguinte projetos:

- Buf (https://github.com/bufbuild/buf)
- Protoc (https://github.com/protocolbuffers/protobuf/releases/tag/v21.9)

## Envoy

Envoy será necessário para fazer o proxy entre o gRPC-web com gRPC no server. Por que? Porque ainda é necessário transformar a conexão HTTP/1.1 para o HTTP/2. Outros poderiam ser escolhidos para fazer o proxy e/ou o proxy reverso, tal qual o NGINX, entretanto, o Envoy já possui suporte integrado ao gRPC-web e amplamente suportada pela comunidade de Remote Procedure Call.

## Commands

A pasta gen desse repositório deve estar no **.gitignore**

```
go mod init

go mod tidy

npm install --save-dev protoc-gen-grpc-web
npm install --save-dev @protobuf-ts/plugin

yarn add grpc-web google-protobuf

buf generate

```

## gRpc-ui

```
grpcui -plaintext 0.0.0.0:50051
```

## Docker

Buildar a imagem

```
docker build -t grpcweb/envoy .
```

Subir o container
