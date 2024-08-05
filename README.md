# Fleet Management System

## Visão Geral

Este projeto é uma aplicação backend para gerenciar motoristas e veículos, e associar motoristas a veículos. A aplicação é construída usando Golang e PostgreSQL, e utiliza gRPC para comunicação.

## Funcionalidades

1. Operações CRUD para motoristas (Criar, Ler, Atualizar, Deletar).
2. Operações CRUD para veículos (Criar, Ler, Atualizar, Deletar).
3. Associação de motoristas com veículos.

## Pré-requisitos

- Docker
- Docker Compose
- Golang 1.22.5

## Configuração

1. Clone o repositório:
   
   git clone https://github.com/blpilla/fleet_management.git
   cd fleet_management
   

2. Crie o arquivo .env na raiz do projeto com as seguintes variáveis de ambiente:
   env
   DB_HOST=db
   DB_PORT=5432
   DB_USER=user
   DB_PASSWORD=password
   DB_NAME=fleet_management
   

3. Execute os containers Docker:
   
   docker compose up --build
   

4. Aplique as migrações para criar as tabelas no banco de dados:
   
   docker exec -it fleet_management-app-1 sh -c 'migrate -path /app/db/migrations -database "postgres://user:password@db:5432/fleet_management?sslmode=disable" up'
   

## Testando a aplicação

1. Verifique se o servidor está funcionando:
   
   docker logs fleet_management-app-1
   

2. Testar a criação de um motorista via gRPC:
   
   grpcurl -plaintext -d '{"first_name": "John", "last_name": "Doe", "license": "ABC123"}' localhost:50051 proto.DriverService/CreateDriver
   

3. Testar a obtenção de um motorista por ID via gRPC:
   
   grpcurl -plaintext -d '{"id": 1}' localhost:50051 proto.DriverService/GetDriverByID
   

4. Testar a criação de um veículo via gRPC:
   
   grpcurl -plaintext -d '{"make": "Toyota", "model": "Corolla"}' localhost:50051 proto.VehicleService/CreateVehicle
   

5. Testar a obtenção de um veículo por ID via gRPC:
   
   grpcurl -plaintext -d '{"id": 1}' localhost:50051 proto.VehicleService/GetVehicleByID
   

6. Testar a associação de um motorista a um veículo via gRPC:
   
   grpcurl -plaintext -d '{"driver_id": 1, "vehicle_id": 1}' localhost:50051 proto.AssociationService/AssociateDriverToVehicle
   

7. Listar todos os motoristas:
   
   grpcurl -plaintext -d '{}' localhost:50051 proto.DriverService/GetAllDrivers
   

8. Listar todos os veículos:
   
   grpcurl -plaintext -d '{}' localhost:50051 proto.VehicleService/GetAllVehicles
   

9. Listar todas as associações de motoristas com veículos:
   
   grpcurl -plaintext -d '{}' localhost:50051 proto.AssociationService/GetAllAssociations
