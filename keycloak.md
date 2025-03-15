# Keycloak

## Explanations

### 1. Configuration de Keycloak

#### a. Créer un Realm

- Un realm est un espace dans Keycloak où vous pouvez gérer les utilisateurs, rôles, et clients.
- Connectez-vous à l'interface d'administration de Keycloak (`http://localhost:8080` par défaut).
- Créez un nouveau realm pour votre application.

#### b. Ajouter des Clients

- Un client représente une application qui peut interagir avec Keycloak.
- Ajoutez un client pour chaque service Go que vous souhaitez sécuriser.
- Configurez les paramètres du client, notamment les URL de redirection et les protocoles (par exemple, OpenID Connect).

#### c. Configurer les Rôles et Permissions

- Définissez les rôles nécessaires pour vos utilisateurs (par exemple, admin, user).
- Assignez ces rôles aux utilisateurs ou groupes appropriés.

#### d. Configurer les Clients Confidentiels

- Pour les services backend, configurez les clients comme étant confidentiels pour permettre l'authentification basée sur les credentials du client.

### 2. Intégration avec vos Services Go

#### a. Utiliser une Bibliothèque Go pour Keycloak

- Vous pouvez utiliser une bibliothèque Go pour faciliter l'intégration avec Keycloak, comme `go-keycloak`.

```bash
go get github.com/Nerzal/gocloak/v12
```

#### b. Configurer l'Authentification

- Configurez votre service Go pour utiliser Keycloak pour l'authentification. Voici un exemple de base :

```go
package main

import (
    "context"
    "log"
    "net/http"

    "github.com/Nerzal/gocloak/v12"
    "github.com/gin-gonic/gin"
)

func main() {
    client := gocloak.NewClient("http://localhost:8080/auth")

    // Authenticate with Keycloak
    token, err := client.LoginAdmin(context.Background(), "admin", "password", "master")
    if err != nil {
        log.Fatalf("Failed to authenticate with Keycloak: %v", err)
    }

    // Use the token to access Keycloak resources
    log.Printf("Access Token: %s", token.AccessToken)

    // Example Gin route
    r := gin.Default()
    r.GET("/secure-endpoint", func(c *gin.Context) {
        // Protect this endpoint with Keycloak authentication
        c.JSON(http.StatusOK, gin.H{"message": "This is a secure endpoint"})
    })

    r.Run(":8081")
}
```

#### c. Configurer l'Autorisation

- Utilisez les rôles et permissions définis dans Keycloak pour contrôler l'accès aux différentes parties de votre application.
- Vous pouvez vérifier les rôles dans le token JWT reçu de Keycloak pour appliquer des règles d'autorisation.

### 3. Sécuriser les Communications

- Assurez-vous que toutes les communications entre vos services et Keycloak sont sécurisées en utilisant HTTPS.
- Configurez les certificats SSL/TLS pour Keycloak et vos services Go.

### 4. Tester l'Intégration

- Testez l'authentification et l'autorisation en accédant aux endpoints sécurisés de votre application.
- Assurez-vous que les utilisateurs ne peuvent accéder qu'aux ressources pour lesquelles ils ont les permissions appropriées.

### 5. Documentation et Maintenance

- Documentez bien la configuration et l'intégration pour faciliter la maintenance future.
- Surveillez les logs de Keycloak et de vos services pour détecter et résoudre rapidement les problèmes de sécurité
