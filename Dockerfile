# Utiliser l'image officielle Go en version alpine
FROM golang:1.25-alpine

# Installer git et autres dépendances nécessaires
RUN apk add --no-cache git bash

# Définir le répertoire de travail
WORKDIR /app

# Copier les fichiers de module et télécharger les dépendances
COPY go.mod go.sum ./
RUN go mod download

# Copier le reste du projet
COPY . .

# Compiler le binaire
RUN go build -o modbus-server ./cmd/modbus-server

# Exposer les ports Modbus et API
EXPOSE 5020 8080

# Définir la commande par défaut pour lancer le serveur
CMD ["./modbus-server", "-config", "config.yaml"]
