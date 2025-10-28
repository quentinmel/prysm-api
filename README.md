# 🏆 Prysm API

API REST pour la plateforme de prédictions Prysm. Cette API permet d'accéder aux données des compétitions, matchs et prédictions de manière simple et efficace.

## 📋 Table des matières

- [Fonctionnalités](#fonctionnalités)
- [Technologies](#technologies)
- [Installation](#installation)
- [Configuration](#configuration)
- [Utilisation](#utilisation)
- [Endpoints](#endpoints)
- [Documentation](#documentation)
- [Tests](#tests)
- [Déploiement](#déploiement)

## ✨ Fonctionnalités

- 🏆 Gestion des compétitions
- ⚽ Suivi des matchs en temps réel
- 📊 Accès aux prédictions des participants
- 🔍 Filtres avancés (statut, pays, date)
- 📚 Documentation interactive avec Swagger UI
- 🚀 API REST performante avec Go Fiber
- 💾 Base de données Supabase (PostgreSQL)
- 🔒 CORS configuré pour les applications web

## 🛠️ Technologies

- **Go** 1.21+
- **Fiber** v2 - Framework web ultra-rapide
- **Supabase** - Base de données PostgreSQL
- **Swagger UI** - Documentation interactive
- **godotenv** - Gestion des variables d'environnement

## 📦 Installation

### Prérequis

- Go 1.21 ou supérieur
- Git
- Un compte Supabase (gratuit)

### Cloner le projet

```bash
git clone https://github.com/quentinmel/prysm-api.git
cd prysm-api
```

### Installer les dépendances

```bash
go mod download
```

## ⚙️ Configuration

### 1. Créer le fichier `.env`

Créez un fichier `.env` à la racine du projet :

```bash
cp .env
```

### 2. Configurer les variables d'environnement

Éditez le fichier `.env` avec vos informations Supabase :

```env
# Supabase Configuration
SUPABASE_URL=https://votre-projet.supabase.co
SUPABASE_KEY=votre_cle_anon_publique

# Server Configuration
PORT=3000
```

### 3. Obtenir vos identifiants Supabase

1. Connectez-vous à [Supabase](https://supabase.com)
2. Créez un nouveau projet ou sélectionnez un projet existant
3. Allez dans **Settings** → **API**
4. Copiez :
   - **URL** : `https://xxxxx.supabase.co`
   - **anon/public key** : `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...`

⚠️ **Important** : Ne partagez jamais votre clé `service_role` publiquement !

### 4. Structure de la base de données

Créez une table `rooms` dans Supabase avec la structure suivante :

```sql
CREATE TABLE rooms (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  team_home TEXT NOT NULL,
  team_away TEXT NOT NULL,
  status TEXT NOT NULL,
  match_date TIMESTAMP WITH TIME ZONE NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  score_home INTEGER,
  score_away INTEGER
);
```

## 🚀 Utilisation

### Démarrer le serveur en développement

```bash
go run main.go
```

Le serveur démarre sur `http://localhost:3000`

### Compiler pour la production

```bash
# Linux/Mac
go build -o prysm-api main.go

# Windows
go build -o prysm-api.exe main.go
```

### Exécuter le binaire

```bash
./prysm-api
```

## 📡 Endpoints

### Base URL

```
http://localhost:3000/api/v1
```

### Health Check

```http
GET /api/v1/health
```

Vérifie l'état de l'API et les variables d'environnement.

**Réponse :**
```json
{
  "success": true,
  "timestamp": "2025-10-28T12:00:00Z",
  "env_check": {
    "supabase_url_exists": true,
    "supabase_url_value": "https://xxxxx.supabase.co",
    "supabase_key_exists": true,
    "supabase_key_length": 200
  }
}
```

### Compétitions

#### Liste des compétitions

```http
GET /api/v1/competitions?status=active&country=France&limit=10
```

**Paramètres :**
- `status` (optionnel) : Filtrer par statut
- `country` (optionnel) : Filtrer par pays
- `limit` (optionnel) : Nombre max de résultats (défaut: 50)

**Réponse :**
```json
{
  "success": true,
  "data": [
    {
      "id": "uuid-1",
      "name": "PSG vs OM",
      "type": "football",
      "country": "France",
      "status": "active",
      "teams": ["PSG", "OM"],
      "match_date": "2025-10-30T20:00:00Z",
      "created_at": "2025-10-20T10:00:00Z",
      "total_rooms": 5,
      "total_participants": 0
    }
  ],
  "meta": {
    "total": 1,
    "page": 1,
    "limit": 50
  }
}
```

### Matchs

#### Liste des matchs

```http
GET /api/v1/matches?status=open&limit=20
```

**Paramètres :**
- `status` (optionnel) : Filtrer par statut (open, closed, finished)
- `limit` (optionnel) : Nombre max de résultats (défaut: 50)

#### Détails d'un match

```http
GET /api/v1/matches/{id}
```

**Réponse :**
```json
{
  "success": true,
  "data": {
    "id": "uuid-1",
    "team_home": "PSG",
    "team_away": "OM",
    "status": "open",
    "match_date": "2025-10-30T20:00:00Z",
    "score_home": null,
    "score_away": null,
    "created_at": "2025-10-20T10:00:00Z"
  }
}
```

## 📚 Documentation

### Swagger UI

Accédez à la documentation interactive :

```
http://localhost:3000/swagger
```

Vous pouvez :
- ✅ Visualiser tous les endpoints
- ✅ Tester les requêtes directement
- ✅ Voir les paramètres et réponses
- ✅ Télécharger la spécification OpenAPI

### Spécification OpenAPI (JSON)

```
http://localhost:3000/api/v1/swagger
```

## 🧪 Tests

### Tester avec curl

```bash
# Health check
curl http://localhost:3000/api/v1/health | jq

# Competitions
curl http://localhost:3000/api/v1/competitions | jq

# Matches
curl http://localhost:3000/api/v1/matches | jq

# Match spécifique
curl http://localhost:3000/api/v1/matches/uuid-1 | jq
```

### Script de test automatisé

```bash
chmod +x test-api.sh
./test-api.sh
```

### Tests unitaires

```bash
go test ./... -v
```

## 🚢 Déploiement

### Docker

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o prysm-api main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/prysm-api .
COPY --from=builder /app/static ./static
EXPOSE 3000
CMD ["./prysm-api"]
```

Construire et lancer :

```bash
docker build -t prysm-api .
docker run -p 3000:3000 --env-file .env prysm-api
```

### Variables d'environnement en production

Configurez ces variables sur votre plateforme de déploiement :

```env
SUPABASE_URL=https://votre-projet.supabase.co
SUPABASE_KEY=votre_cle_anon_publique
PORT=3000
```

## 📂 Structure du projet

```
prysm-api/
├── handlers/           # Gestionnaires de routes
│   ├── handlers.go    # Endpoints principaux
│   └── swagger.go     # Spécification OpenAPI
├── models/            # Modèles de données
│   └── models.go      # Structures Room, Competition, Match
├── database/          # Configuration base de données
│   └── supabase.go    # Client Supabase
├── static/            # Fichiers statiques
│   └── swagger.html   # Interface Swagger UI
├── .env               # Variables d'environnement (non versionné)
├── .env.example       # Exemple de configuration
├── .gitignore         # Fichiers à ignorer
├── go.mod             # Dépendances Go
├── go.sum             # Checksums des dépendances
├── main.go            # Point d'entrée
├── README.md          # Documentation
└── test-api.sh        # Script de test
```

## 🤝 Contribution

Les contributions sont les bienvenues ! N'hésitez pas à :

1. Fork le projet
2. Créer une branche (`git checkout -b feature/AmazingFeature`)
3. Commit vos changements (`git commit -m 'Add AmazingFeature'`)
4. Push vers la branche (`git push origin feature/AmazingFeature`)
5. Ouvrir une Pull Request

## 📝 License

Ce projet est sous licence MIT.

## 👤 Auteur

**Quentin Mel**

- GitHub: [@quentinmel](https://github.com/quentinmel)

## 🙏 Remerciements

- [Fiber](https://gofiber.io/) - Framework web Go
- [Supabase](https://supabase.com/) - Backend as a Service
- [Swagger UI](https://swagger.io/tools/swagger-ui/) - Documentation interactive

---

⭐ Si ce projet vous aide, n'hésitez pas à lui donner une étoile !