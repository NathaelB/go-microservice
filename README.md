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
CREATE DATABASE role_service;
```

## A faire

### Création des Topics

- guild-events
- role-events
- member-events

### Développement des Interfaces et Modèles

- Définir les Modèles :
        Créer des structures de données en Go pour représenter les événements que que l'on va publier/consommer via Kafka.
- Interfaces :
        Définir les interfaces pour les producteurs et consommateurs Kafka dans les services Go.

### Implémentation des Transactions Distribuées

- Transactions Kafka :
        Configurer les producteurs Kafka pour utiliser les transactions afin de garantir l'atomicité des opérations.
        Vérifier que les consommateurs peuvent gérer les transactions.

### Intégration de Keycloak pour la Sécurité

- Configuration de Keycloak :
        Configurer Keycloak pour gérer les utilisateurs, rôles, et permissions nécessaires pour nos services.
- Intégration avec nos Services :
        Utiliser les bibliothèques Go pour intégrer Keycloak dans nos services afin de gérer l'authentification et l'autorisation.

### Implémentation d'Avro pour la Sérialisation

- Définir les Schémas Avro :
        Créer des fichiers .avsc pour définir les schémas de nos événements.
- Utilisation dans Go :
        Utiliser une bibliothèque Go pour Avro pour sérialiser et désérialiser les messages Kafka.
