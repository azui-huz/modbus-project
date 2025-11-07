# Modbus Server Project (Go)

A modular **Modbus TCP server** with REST API and web SCADA support, built in Go.  
This project allows you to configure a Modbus server via YAML, force or release registers, read registers, and generate server architecture for automatic client creation.

---

## Features

- **Modbus Server**  
  - TCP Modbus server with configurable holding registers and coils  
  - Thread-safe read/write operations  
  - Force/release registers  

- **REST API**  
  - Force or release registers  
  - Read a single register or all registers  
  - Return server architecture (for automatic client generation)  

- **Cyclic Readers**  
  - Periodically read specified registers  

- **Clients**  
  - Connect to external Modbus clients automatically based on YAML description  

- **Web SCADA**  
  - Frontend in JS for real-time visualization  

- **Dockerized**  
  - Run server and API in a container easily  

---

## Project Structure

```
modbus-project/
├── cmd/
│   └── modbus-server/       # main.go: entry point for the Modbus server
├── internal/
│   ├── api/                 # REST API handlers
│   │   └── handlers.go
│   ├── clientmgr/           # Modbus client manager from server description
│   │   └── clientmgr.go
│   ├── config/              # YAML configuration loader
│   │   └── config.go
│   ├── cycreader/           # Cyclic readers implementation
│   │   └── cycreader.go
│   ├── modbussrv/           # Modbus server implementation
│   │   ├── server.go
│   │   └── types.go         # Not implemented yet
│   └── web/                 # Web SCADA frontend
│       ├── app.js           # Main JS frontend logic
│       └── index.html       # Web SCADA UI
├── config.yaml              # Project configuration
├── docker-compose.yml       # Docker Compose for container orchestration
├── Dockerfile               # Docker image for the server
├── Makefile                 # Build and run commands
├── README.md                # Project documentation
└── go.mod / go.sum          # Go module dependencies
```

---

## Configuration (\`config.yaml\`)

```yaml
server:
  host: "0.0.0.0"          
  port: 5020               
  unit_id: 1               
  holding_registers:
    size: 100              
  coils:
    size: 200              

api:
  port: 8080               

cyclic_readers:
  - name: meter
    interval_ms: 1000
    registers:
      - {address: 0, length: 10}

clients:
  - name: client
    host: 192.168.1.20
    port: 502
    unit_id: 1
```

### Notes

- `server.host` / `server.port`: Modbus server address and port  
- `unit_id`: Modbus unit ID  
- `holding_registers.size`: number of holding registers  
- `coils.size`: number of coils  
- `api.port`: REST API port  
- `cyclic_readers`: list of registers to read periodically  
- `clients`: external Modbus clients to connect to  

---

## Getting Started

### 1. Install Go

Go 1.25 or higher is required:

```bash
go version
```

### 2. Clone the repository

```bash
git clone https://github.com/your-username/modbus-project.git
cd modbus-project
```

### 3. Build and run locally

```bash
go mod tidy
go build -o modbus-server ./cmd/modbus-server
./modbus-server -config config.yaml
```

### 4. Run with \`go run\`

```bash
go run ./cmd/modbus-server -config config.yaml
```

---

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET    | `/api/read/all` | Read all holding registers |
| GET    | `/api/read/holding/{addr}` | Read a specific holding register |
| POST   | `/api/force` | Force a holding register (`{"type":"holding","addr":5,"value":123}`) |
| POST   | `/api/release` | Release a forced register (`{"type":"holding","addr":5}`) |
| GET    | `/api/architecture` | Get server architecture |

Examples:

```bash
curl http://localhost:8080/api/read/holding/5
```

```bash
curl http://localhost:8080/api/architecture
```

```bash
curl -X POST -H "Content-Type: application/json" \
     -d '{"type":"holding","addr":5,"value":123}' \
     http://localhost:8080/api/force
```

```bash
curl -X POST -H "Content-Type: application/json" \
     -d '{"type":"holding","addr":5}' \
     http://localhost:8080/api/release
```

---

## Makefile Commands

| Command | Description |
|---------|-------------|
| `make build` | Compile the Go binary |
| `make run` | Run the server locally |
| `make clean` | Remove the binary |
| `make docker-build` | Build Docker image |
| `make docker-run` | Run server in Docker |

---

## Using Docker

### Build Docker Image

```bash
make docker-build
```

or manually:

```bash
docker build -t modbus-server .
```

### Run Container

```bash
make docker-run
```

or manually:

```bash
docker run -p 5020:5020 -p 8080:8080 \
  -v $(pwd)/config.yaml:/app/config.yaml \
  modbus-server
```

### Run docker-compose

```bash
docker compose up -d
```

---

## License

MIT License © azui-huz
