# Go Stock Service

![Go](https://img.shields.io/badge/Go-1.23.2-blue.svg)
![Kafka](https://img.shields.io/badge/KafkaGo-0.4.47-blue.svg)
![License](https://img.shields.io/badge/license-MIT-green.svg)

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Architecture](#architecture)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Service](#running-the-service)

## Introduction

The **Go Stock Service** is a high-performance, scalable service built with Go that consumes stock price data from a Kafka broker. It processes real-time stock information, making it ideal for applications requiring up-to-date financial data. Whether you're building a trading platform, analytics dashboard, or any financial application, this service provides a robust foundation for handling stock data efficiently.

## Features

- **Real-Time Data Consumption:** Connects to Kafka to consume live stock price updates.
- **Scalable Architecture:** Designed to handle high-throughput data streams.
- **Configurable:** Easily configurable through environment variables.
- **Logging & Monitoring:** Integrated logging for easy debugging and monitoring.
- **Graceful Shutdown:** Ensures all processes terminate smoothly without data loss.
- **Error Handling:** Robust error handling mechanisms to maintain service reliability.

## Architecture

![Architecture Diagram](https://via.placeholder.com/800x400?text=Architecture+Diagram)

1. **Kafka Broker:** Acts as the central hub for stock price data.
2. **Go Stock Service:** Consumes data from Kafka, processes it, and performs necessary actions.
3. **Storage/Database (Optional):** Stores processed data for further analysis or retrieval.
4. **Monitoring Tools:** Keeps track of service performance and health.

## Prerequisites

Before setting up the Go Stock Service, ensure you have the following installed:

- **Go:** [Download and Install Go](https://golang.org/dl/)
- **Kafka:** [Download and Install Apache Kafka](https://kafka.apache.org/quickstart)

## Installation

1. **Clone the Repository:**

    ```bash
    git clone https://github.com/ZiyadBouazara/stock-service-go.git
    cd stock-service-go
    ```

2. **Install Dependencies:**

   The project uses Go modules. Ensure you're inside the project directory.

    ```bash
    go mod download
    ```

## Configuration

The service is configured using environment variables defined in a `.env` file. Create a `.env` file in the root directory of the project with the following content:

```env
# Kafka Configuration
KAFKA_BROKER_URL=localhost:9092
KAFKA_TOPIC=bitcoin-price-topic
KAFKA_GROUP_ID=stockservice-go-consumer
```
## Running the service

1. **Run Kafka:**

    Two options: Either you run your own instances of Kafka and Zookeeper, OR you can use the `docker-compose.yaml` at the root fo the bitcoin pulse project that will run the images for you.


2. **Build the service:**

    ```bash
   go build -o stockservice
    ```
    ```bash
   ./stockservice
    ```
