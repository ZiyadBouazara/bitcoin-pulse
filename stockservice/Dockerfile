FROM maven:3-amazoncorretto-21

WORKDIR /app

COPY pom.xml .
RUN mvn --batch-mode exec:java || exit 0

COPY . .
RUN mvn --batch-mode compiler:compile

CMD ["mvn", "--batch-mode", "--offline", "exec:java"]