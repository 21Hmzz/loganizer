# GoLog Analyzer

GoLog Analyzer est un outil CLI développé en Go permettant d'analyser plusieurs fichiers de logs de manière parallèle, robuste et automatisée. Il centralise les résultats et les exporte au format JSON.

---

## Fonctionnalités

- Analyse parallèle de plusieurs fichiers de logs avec goroutines
- Gestion d'erreurs personnalisées (fichier introuvable, erreur de parsing)
- Interface CLI avec Cobra (sous-commandes, flags)
- Chargement d’une configuration JSON listant les logs à analyser
- Export d’un rapport d’analyse au format JSON (avec gestion des dossiers)
- Nom de fichier de rapport horodaté (format AAMMJJ)
- Filtrage des résultats avec `--status OK/FAILED`
- Commande `add-log` pour enrichir le fichier de configuration

---

## Installation

```bash
git clone https://github.com/21Hmzz/loganizer.git
cd loganalyzer
go mod tidy
go build -o loganalyzer
```

## Commandes disponibles

- analyze – Analyser les logs
loganalyzer analyze -c config.json -o rapports/resultat.json --status OK
  --config, -c
  Chemin vers le fichier config.json
  --output, -o
  Chemin vers le rapport exporté (les dossiers sont créés automatiquement)
  --status
  Filtrer les résultats

- add-log – Ajouter un log au fichier de configuration
loganalyzer add-log --id app-3 --path /var/log/app3.log --type custom --file config.json
--id
Identifiant unique du log
--path
Chemin du fichier de log
--type
Type du log (informel)
--file
Fichier de configuration existant

## Fonctionnement interne
	•	main.go démarre l’application CLI via Cobra
	•	cmd/ contient les sous-commandes analyze et add-log
	•	internal/config charge la configuration JSON
	•	internal/analyzer simule l’analyse avec gestion des erreurs
	•	internal/reporter exporte le rapport JSON final

## Bonus techniques implémentés
Création automatique des dossiers d’export
via os.MkdirAll()

Préfixe AAMMJJ du fichier de sortie
via time.Now().Format("060102")

Filtrage par statut avec --status
utile pour n’exporter que les erreurs

Ajout dynamique avec add-log
ajoute une entrée à config.json

## Équipe de développement
Hamza BELLA

Valentin POVEDA AMARAL

Serhat KUS
