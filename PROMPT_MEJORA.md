# Prompt para Mejorar el Codigo Base

Copia y pega el siguiente contenido completo en un asistente de IA (Claude, ChatGPT, etc.)
para obtener un ZIP con el proyecto corregido y listo para compilar.

---

```
Eres un asistente experto en análisis, corrección y generación de archivos de cualquier tipo:
código fuente, documentación, hojas de cálculo, documentos Word, configuraciones, entre otros.
Voy a enviarte una cadena de texto que contiene uno o más archivos. Cada archivo está delimitado por un marcador con el siguiente formato:
// === ARCHIVO: ruta/del/archivo.extension ===
o también puede aparecer como:
## === ARCHIVO: ruta/del/archivo.extension ===
Lo que sigue al marcador puede ser:

El contenido real del archivo (código, texto, YAML, etc.)
Una descripción en lenguaje natural de lo que debe contener el archivo


TU TAREA
PASO 1 — Detección y extracción
Identifica todos los archivos presentes en la cadena. Para cada archivo extrae:

Su ruta completa (ej: src/main/java/com/pragma/Service.java)
Su contenido o descripción

PASO 2 — Clasificación por tipo
Clasifica cada archivo en una de estas categorías:
A) Código fuente (Java, Python, TypeScript, JavaScript, Kotlin, etc.)
B) Configuración / documentación (YAML, properties, Markdown, JSON, txt, etc.)
C) Excel (.xlsx, .xls, .csv)
D) Word (.docx, .doc)
E) Otro tipo de archivo binario o especial
PASO 3 — Clasificación de errores en código fuente

Objetivo prioritario: que el proyecto compile. No corrijas flujo de negocio ni lógica funcional.

Antes de modificar cualquier archivo de código fuente, clasifica cada problema encontrado en una de estas dos categorías:
🔴 ERROR DE COMPILACIÓN — corregir siempre
Son errores que impiden que el proyecto arranque, sin valor pedagógico:

Import faltante o incorrecto
Clase, método o variable referenciada que no existe en ningún archivo del proyecto
Error de sintaxis
Anotación con atributos inválidos
Dependencia ausente en pom.xml, package.json, etc.
Archivo referenciado que no existe y debe ser creado con implementación mínima

→ CORREGIR estos errores.
🟡 PROBLEMA FUNCIONAL O DE CALIDAD — preservar siempre
Son problemas que no impiden compilar. Pueden ser intencionales para el aprendizaje:

Clave secreta hardcodeada ("secret", "password123")
API deprecada que funciona pero tiene reemplazo moderno
Lógica de negocio incorrecta o incompleta
Código redundante o de baja legibilidad
Falta de validaciones en flujo de negocio
Patrones de diseño incorrectos pero funcionales
Concurrencia no segura
Configuración funcional pero no óptima

→ PRESERVAR tal cual. No corregir, no mejorar, no comentar.
PASO 4 — Procesamiento según tipo de archivo
Tipo A — Código fuente
Aplica únicamente las correcciones clasificadas como 🔴 ERROR DE COMPILACIÓN.
No alteres ningún elemento clasificado como 🟡 PROBLEMA FUNCIONAL O DE CALIDAD.
Si falta un archivo referenciado, créalo con la implementación mínima necesaria para compilar.
Tipo B — Configuración / documentación
Extrae el contenido tal cual, sin modificaciones salvo errores evidentes de sintaxis
(ej: YAML mal indentado).
Tipo C — Excel (.xlsx)
Si viene con contenido real, genera el archivo respetando ese contenido.
Si viene con descripción en lenguaje natural, genera un archivo Excel funcional con:

Fila de encabezados en negrita con color de fondo distintivo
Columnas con ancho ajustado al contenido
Tipos de dato correctos por columna
Validaciones si la descripción lo indica
Hojas nombradas descriptivamente si hay más de una
Filas de ejemplo si no hay datos reales

Tipo D — Word (.docx)
Si viene con contenido real, genera el archivo respetando ese contenido.
Si viene con descripción en lenguaje natural, genera un documento Word funcional con:

Estilos de título (Título 1, Título 2) para jerarquía de secciones
Fuente legible (Calibri o equivalente), tamaño 11-12pt para cuerpo
Márgenes estándar
Tabla de contenido si tiene múltiples secciones
Tablas con encabezados en negrita si aplica

Tipo E — Otro
Genera el archivo con el contenido o estructura más apropiada según la descripción.
PASO 5 — Exportación en ZIP
Empaqueta todos los archivos en un único archivo ZIP descargable respetando exactamente
la estructura de rutas indicada por los marcadores.
El ZIP debe incluir:

Archivos de código con únicamente los errores de compilación corregidos
Archivos de configuración y documentación sin cambios
Archivos nuevos creados para resolver dependencias de compilación faltantes
Archivos Excel y Word generados desde descripción

IMPORTANTE: El ZIP debe estar listo para descargar al finalizar. No preguntes si el usuario
quiere generarlo. Simplemente genera el archivo y proporciona el enlace de descarga; No debes desplegar en el chat el resumen de lo que arreglaste al Zip, solo entregalo.

REGLAS IMPORTANTES

No omitas ningún archivo aunque no tenga errores ni modificaciones
Respeta los nombres y rutas exactas indicadas por los marcadores
Si un archivo no tiene marcador claro, infiere el nombre desde su contenido
Si la cadena contiene solo documentación o descripciones sin código, genera los archivos
correspondientes sin aplicar análisis de compilación
No agregues texto después del enlace de descarga del ZIP
No preguntes si el usuario quiere el ZIP: simplemente generalo siempre
Si detectas que falta un archivo de configuración necesario para compilar
(pom.xml, package.json, requirements.txt, build.gradle, etc.), créalo e inclúyelo
inferiendo su contenido desde los imports y frameworks detectados en el código
Nunca corrijas problemas 🟡 aunque parezcan obvios o fáciles de mejorar.
El participante que recibirá este proyecto los debe encontrar y resolver él mismo.


INPUT
Aquí está la cadena con los archivos:
// === ARCHIVO: src/cmd/main.go ===
package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/yourusername/payment-microservice/src/internal/handlers"
	"google.golang.org/grpc"
)

func main() {
	// Create a gRPC server
	grpcServer := grpc.NewServer()
	paymentHandler := &handlers.PaymentHandler{}
	pb.RegisterPaymentServiceServer(grpcServer, paymentHandler)

	// Create a gRPC gateway
	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := runtime.RegisterService(ctx, gwmux, "payment.PaymentService", paymentHandler)
	if err!= nil {
		log.Fatalf("failed to register service: %v", err)
	}

	// Serve the gRPC server and gateway
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: gwmux,
	}
	log.Println("Server started at :8080")
	if err := httpServer.ListenAndServe(); err!= nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// === ARCHIVO: src/internal/handlers/payment_handler.go ===
package handlers

import (
	"context"
	"log"

	pb "github.com/yourusername/payment-microservice/src/internal/handlers"
	"github.com/yourusername/payment-microservice/src/internal/services"
)

type PaymentHandler struct {}

func (h *PaymentHandler) ProcessPayment(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	log.Println("Processing payment")
	service := services.NewPaymentService()
	response, err := service.ProcessPayment(ctx, req)
	if err!= nil {
		return nil, err
	}
	return response, nil
}

// === ARCHIVO: src/internal/services/payment_service.go ===
package services

import (
	"context"
	"log"

	"github.com/yourusername/payment-microservice/src/internal/repository"
	"github.com/yourusername/payment-microservice/src/pkg/external"
	"github.com/yourusername/payment-microservice/src/pkg/idempotency"
	pb "github.com/yourusername/payment-microservice/src/internal/handlers"
)

type PaymentService struct {}

func NewPaymentService() *PaymentService {
	return &PaymentService{}
}

func (s *PaymentService) ProcessPayment(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	idKey := idempotency.GenerateIdempotencyKey(req)
	if idempotency.IsIdempotent(idKey) {
		log.Println("Payment is idempotent")
		return &pb.PaymentResponse{Status: "SUCCESS"}, nil
	}
		// Validate payment
	if err := validatePayment(req); err!= nil {
		return nil, err
	}
		// Check external service
	client := external.NewExternalServiceClient()
	if err := client.CheckFunds(ctx, req.AccountID); err!= nil {
		return nil, err
	}
		// Save payment
	repo := repository.NewPaymentRepository()
	if err := repo.SavePayment(ctx, req); err!= nil {
		return nil, err
	}
		// Mark as idempotent
	idempotency.MarkAsIdempotent(idKey)
	return &pb.PaymentResponse{Status: "SUCCESS"}, nil
}

func validatePayment(req *pb.PaymentRequest) error {
	// Implement validation logic
	return nil
}

// === ARCHIVO: src/internal/repository/payment_repository.go ===
package repository

import (
	"context"
	"log"

	pb "github.com/yourusername/payment-microservice/src/internal/handlers"
)

type PaymentRepository struct {}

func NewPaymentRepository() *PaymentRepository {
	return &PaymentRepository{}
}

func (r *PaymentRepository) SavePayment(ctx context.Context, req *pb.PaymentRequest) error {
	log.Println("Saving payment")
	// Implement save logic
	return nil
}

// === ARCHIVO: src/pkg/idempotency/idempotency_key.go ===
package idempotency

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateIdempotencyKey(req interface{}) string {
	// Implement idempotency key generation
	hash := sha256.New()
	hash.Write([]byte("some-data"))
	return hex.EncodeToString(hash.Sum(nil))
}

func IsIdempotent(key string) bool {
	// Implement idempotency check
	return false
}

func MarkAsIdempotent(key string) {
	// Implement idempotency mark
}

// === ARCHIVO: src/pkg/external/external_service_client.go ===
package external

import (
	"context"
	"log"
)

type ExternalServiceClient struct {}

func NewExternalServiceClient() *ExternalServiceClient {
	return &ExternalServiceClient{}
}

func (c *ExternalServiceClient) CheckFunds(ctx context.Context, accountID string) error {
	log.Println("Checking funds")
	// Implement fund check logic
	return nil
}

// === ARCHIVO: src/config/config.yaml ===
# Configuration for the microservice
port: 8080
external_service_url: http://external-service:8080

// === ARCHIVO: src/test/payment_service_test.go ===
package test

import (
	"testing"
	"context"

	"github.com/yourusername/payment-microservice/src/internal/services"
	pb "github.com/yourusername/payment-microservice/src/internal/handlers"
)

func TestProcessPayment(t *testing.T) {
	service := services.NewPaymentService()
	req := &pb.PaymentRequest{
		AccountID: "1234567890",
		Amount:     100.0,
	}
	_, err := service.ProcessPayment(context.Background(), req)
	if err!= nil {
		 t.Errorf("ProcessPayment failed: %v", err)
	}
}

// === ARCHIVO: docker/Dockerfile ===
# Dockerfile for the microservice
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum./
RUN go mod download

COPY../
RUN go build -o payment-microservice

FROM alpine:latest

COPY --from=builder /app/payment-microservice /app/

WORKDIR /app

EXPOSE 8080

CMD ["./payment-microservice"]

// === ARCHIVO: docker/docker-compose.yml ===
version: '3.8'

services:
  payment-microservice:
    build:.
    ports:
      - "8080:8080"
    depends_on:
      - external-service

  external-service:
    image: external-service:latest
    ports:
      - "8080:8080"

// === ARCHIVO: docker/.dockerignore ===
# Files to ignore in the Docker build
*.go
*.sum
*.mod

// === ARCHIVO: docker/healthcheck.sh ===
#!/bin/sh

# Health check script for Docker
curl -f http://localhost:8080/health || exit 1

```
