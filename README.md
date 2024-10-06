# Bitcoin Pulse


Bitcoin Pulse is a real-time cryptocurrency tracker that fetches live Bitcoin prices using the Coinbase API 
and displays the data on an interactive graph. The application is built using a microservices architecture,
leveraging Spring Boot, Kafka, and React to provide a scalable and efficient solution for real-time tracking.

[![Java](https://img.shields.io/badge/java-17-blue)](https://www.oracle.com/java/technologies/javase-jdk17-downloads.html)
[![Spring Boot](https://img.shields.io/badge/spring--boot-3.3.3-brightgreen)](https://spring.io/projects/spring-boot)
[![Spring Kafka](https://img.shields.io/badge/spring--kafka-3.2.4-brightgreen)](https://spring.io/projects/spring-kafka)
[![Maven](https://img.shields.io/badge/maven-3.8.5-orange)](https://maven.apache.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Features

* Real-time Bitcoin price tracking.
* Kafka-based messaging for publishing and consuming price updates.
* Interactive graphical visualization of Bitcoin prices.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing
purposes.

### Installation

Build the project:

```bash
mvn clean install
```

Run the project:

```bash
mvn spring-boot:run
```

Docker build image

```bash
docker build -t bitcoin-pulse .
```

Docker run image

```bash
docker compose up -d
```

Docker stop image

```bash
docker compose down -v
```

## Open Source files

For more information about the project and how to contribute, please refer to the following files:

- [CONTRIBUTING](CONTRIBUTING.md)
- [CODE_OF_CONDUCT](CODE_OF_CONDUCT.md)
- [LICENSE](LICENSE)

## License

This project is lisenced under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

Thanks to all contributors who decide to participate in this project.