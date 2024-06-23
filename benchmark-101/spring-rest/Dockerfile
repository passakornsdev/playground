FROM openjdk:21-slim-bullseye

RUN mkdir -p /home/authentication/
WORKDIR /home/authentication/

COPY ./target/spring-rest-0.0.1-SNAPSHOT.jar .

EXPOSE 8080

CMD ["java","-jar", "spring-rest-0.0.1-SNAPSHOT.jar"]
