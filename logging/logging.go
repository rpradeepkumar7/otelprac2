package logging

import (
	"context"
	"log"
	"os"
	"time"
	"encoding/json"
)

type OTelLog struct {
	Timestamp           string `json:"Timestamp"`
	ObservedTimestamp   string `json:"ObservedTimestamp"`
	TraceId             string `json:"TraceId"`
	SpanId              string `json:"SpanId"`
	SeverityText        string `json:"SeverityText"`
	SeverityNumber      string `json:"SeverityNumber"`
	Body                string `json:"Body"`
	Resource            struct {
		ServiceName string `json:"service.name"`
		HostName    string `json:"host.name"`
		HostIP      string `json:"host.ip"`
		HostMAC     string `json:"host.mac"`
		OSType      string `json:"os.type"`
		OSVersion   string `json:"os.version"`
	} `json:"Resource"`
	InstrumentationScope struct {
		Name    string `json:"Name"`
		Version string `json:"Version"`
	} `json:"InstrumentationScope"`
	Attributes struct {
		HTTPMethod     string `json:"http.method"`
		HTTPStatusCode string `json:"http.status_code"`
		HTTPURL        string `json:"http.url"`
		DBOperation    string `json:"db.operation"`
		FirewallStatus string `json:"firewall.status"`
		NetworkLatency string `json:"network.latency"`
		CPUUsage       string `json:"cpu.usage"`
		MemoryUsage    string `json:"memory.usage"`
	} `json:"Attributes"`
	EventData struct {
		EventName string `json:"event.name"`
		EventType string `json:"event.type"`
	} `json:"EventData"`
	Exception struct {
		Message    string `json:"exception.message"`
		Type       string `json:"exception.type"`
		StackTrace string `json:"exception.stacktrace"`
	} `json:"Exception"`
	Duration string `json:"Duration"`
	Status   string `json:"Status"`
	LogLevel string `json:"LogLevel"`
}

func InitLogger(ctx context.Context) {
	// Example log creation
	logEntry := OTelLog{
		Timestamp:         time.Now().Format(time.RFC3339),
		ObservedTimestamp: time.Now().Format(time.RFC3339),
		TraceId:           "abcd1234",
		SpanId:            "efgh5678",
		SeverityText:      "ERROR",
		SeverityNumber:    "17",
		Body:              "An error occurred while processing the request.",
		Duration:          "500ms",
		Status:            "Error",
		LogLevel:          "Critical",
	}

	logEntry.Resource.ServiceName = "web-backend"
	logEntry.Resource.HostName = "web-server-1"
	logEntry.Resource.HostIP = "192.168.1.1"
	logEntry.Resource.HostMAC = "00:1A:2B:3C:4D:5E"
	logEntry.Resource.OSType = "Linux"
	logEntry.Resource.OSVersion = "Ubuntu 20.04"

	logEntry.InstrumentationScope.Name = "example-logger"
	logEntry.InstrumentationScope.Version = "1.0.0"

	logEntry.Attributes.HTTPMethod = "GET"
	logEntry.Attributes.HTTPStatusCode = "500"
	logEntry.Attributes.HTTPURL = "http://example.com"
	logEntry.Attributes.DBOperation = "SELECT"
	logEntry.Attributes.FirewallStatus = "enabled"
	logEntry.Attributes.NetworkLatency = "100ms"
	logEntry.Attributes.CPUUsage = "75%"
	logEntry.Attributes.MemoryUsage = "60%"

	logEntry.EventData.EventName = "example-event"
	logEntry.EventData.EventType = "example-type"

	logEntry.Exception.Message = "Null Pointer Exception"
	logEntry.Exception.Type = "NullPointerException"
	logEntry.Exception.StackTrace = "stack trace details here"

	// Print log as JSON
	jsonData, err := json.MarshalIndent(logEntry, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling log entry: %v", err)
	}

	log.SetOutput(os.Stdout)
	log.Println(string(jsonData))
}

func LogEvent(ctx context.Context) {
	InitLogger(ctx)
}
