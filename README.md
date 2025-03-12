# Microservices

Création d'une architecture microservices avec Kafka, Keycloak, Postgres, et plusieurs microservices.

Plusieurs microservices sont créés :

- Guild
- Member
- Role

## Keycloak

Keycloak est un serveur d'identité et d'authentification open-source, il est utilisé pour l'authentification et l'autorisation des microservices.

Chaque micro-service a son propre client Keycloak, avec ses propres accès. L'avantage est que les microservices ne doivent pas se soucier de l'authentification et de l'autorisation, seulement de l'utilisation des API. De surcroît il sera d'avantage plus facile de rajouter de la traçabilité à l'application.

## Kafka

Kafka est un système de messagerie distribuée, il est utilisé pour la communication asynchrone entre les microservices. Nous l'utiliserons afin d'éviter les problèmes de dépendances entre les microservices ainsi que la mise en place d'un système de transaction distribuée.

## Postgres

Création des bases de données pour chaque microservice.

```bash
docker exec -it postgres psql -U postgres
CREATE DATABASE member_service;
CREATE DATABASE guild_service;
```
