# Bitcoin Pulse

![DALL·E 2024-11-03 11 01 04 - A cartoon-style logo for a project called 'Bitcoin Pulse', inspired by the simplicity of Go's gopher logo  The logo features a cute, simplified Bitcoi](https://github.com/user-attachments/assets/3f8ac386-b3e9-474e-bc75-a1aaeabf9672)

Bitcoin Pulse is a real-time cryptocurrency tracker that fetches live Bitcoin prices using the Coinbase Websocket Feed
and displays the data on an interactive graph. The application is built using a microservices architecture,
leveraging Spring Boot, Kafka, and React to provide a scalable and efficient solution for real-time tracking.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Tech Stack

![Java](https://img.shields.io/badge/Java-ED8B00?style=for-the-badge&logo=openjdk&logoColor=white)
![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Typescript](https://img.shields.io/badge/TypeScript-007ACC?style=for-the-badge&logo=typescript&logoColor=white)
![React](https://img.shields.io/badge/react-%2320232a.svg?style=for-the-badge&logo=react&logoColor=%2361DAFB)
![Postgres](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)

## Features

* Real-time Bitcoin price tracking.
* Kafka-based messaging for publishing and consuming price updates.
* Interactive graphical visualization of Bitcoin prices.

## Architecture

![Local Image](docs/architecture_bitcoin_pulse.jpeg)

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing
purposes.

### Installation

A `docker-compose.yaml` is available to run Kafka and Zookeeper images, use the following commands if you see the need:

Docker run images

```bash
docker compose up -d
```

Docker stop images

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
