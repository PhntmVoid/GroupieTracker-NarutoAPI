# 🌌 Naruto Data Explorer — GroupieTracker  

Bienvenue sur **Naruto Data Explorer**, un projet web inspiré de la célèbre série anime/manga **Naruto**. Développée en **Go (Golang)**, cette application permet aux fans d'**explorer, rechercher et découvrir** l'univers de Naruto grâce à une API dédiée.  

Toutes les données sont récupérées dynamiquement depuis l'API [Dattebayo](https://api-dattebayo.vercel.app/).  

---

## 📜 Présentation du Projet  

En tant que grand fan de Naruto depuis mon enfance, j'ai voulu allier ma passion pour la série et le développement web afin de créer une **plateforme interactive et complète** permettant de découvrir tout l’univers de Naruto : personnages, clans, Kekkei-Genkai, et bijū (démons à queues).  

---

## 🚀 Ce que le projet propose  

### 🎭 Informations sur les Personnages  
Découvrez une large collection de personnages de Naruto, avec leurs **pouvoirs, histoires et particularités uniques**. Que vous soyez fan des ninjas emblématiques comme **Naruto, Sasuke, Sakura**, ou curieux des personnages secondaires.

### 🏯 Clans et Kekkei-Genkai  
Explorez le monde des **clans et Kekkei-Genkai**, pour voir les lignées puissantes dans l’histoire de Naruto.  

### 🦊 Démons à Queues (Bijū)  
Apprenez-en plus sur les **démons à queues**, créatures mythiques qui influencent profondément l’équilibre des forces dans le monde ninja.  

---

## 🌐 Endpoints API utilisés  

| Fonctionnalité          | URL de l'API                                                  |
|-------------------------|---------------------------------------------------------------|
| Personnages             | https://dattebayo-api.onrender.com/characters                 |
| Clans                   | https://dattebayo-api.onrender.com/clans                      |
| Kekkei-Genkai           | https://dattebayo-api.onrender.com/kekkei-genkai              |
| Démons à Queues         | https://dattebayo-api.onrender.com/tailed-beasts              |

---

## ⚙️ Fonctionnalités principales  

- ✅ **Interface intuitive et simple d'utilisation**.  
- ✅ **Pagination** : navigation fluide à travers de nombreuses données.  
- ✅ **Recherche** : trouvez facilement des personnages, clans ou Kekkei-Genkai.  
- ✅ **Système de filtres** : filtrez les personnages selon plusieurs critères.  
- 🚧 **Système de favoris (à venir)** : ajoutez vos personnages et clans préférés.  

---

## 🗺️ Routes de l'application  

| Route                 | Description                                    |
|-----------------------|------------------------------------------------|
| `/`                   | Page d'accueil et de présentation.             |
| `/charactersPage`     | Liste complète des personnages.                |
| `/clansPage`          | Découverte des clans célèbres.                 |
| `/kgPage`             | L'ensemble des Kekkei-Genkai.                  |
| `/tailedBeastsPage`   | Profils détaillés des démons à queues (Bijū).  |

---

## 📦 Guide complet d'installation et de lancement  


### 🔽 Étapes détaillées d'installation  

#### 1. **Cloner le dépôt GitHub**  

Commencez par télécharger une copie locale du projet :  

```bash
git clone https://github.com/PhntmVoid/GroupieTracker.git 
```
Et ecris ```go run . ``` dans le terminal (vous devez vérifier que vous êtes dans le fichier où se trouve le projet).

