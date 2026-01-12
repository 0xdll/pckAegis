# pckAegis  

![Build](https://img.shields.io/badge/build-passing-brightgreen) ![Go](https://img.shields.io/badge/Go-1.25.5-00ADD8?logo=go) ![License](https://img.shields.io/badge/license-MIT-blue)

pckAegis est un outil CLI écrit en Go qui analyse vos fichiers `requirements.txt` et `package.json`, interroge l'API OSV.dev et retourne un code de sortie `1` dès qu'une vulnérabilité est détectée afin de bloquer immédiatement vos pipelines CI/CD.

## Features
- **Modularité** : parsing dédié pour `requirements.txt` (PyPI) et `package.json` (npm).
- **OSV API** : requêtes ciblées vers `https://api.osv.dev` pour identifier les CVE connues.
- **CI/CD Ready** : code de sortie non nul dès la première vulnérabilité pour stopper les pipelines.
- **Export JSON** : génération optionnelle de rapports au format JSON pour archivage et traçabilité.
- **Mode strict** : contrôle du comportement de sortie (bloquer ou continuer) via l'option `-strict`.
- **Conteneurisé** : image Docker minimale prête à être poussée dans vos registres.

## Installation rapide

### Option Go
1. Prérequis : Go >= 1.25.
2. Compiler le binaire :
   ```bash
   go build -o pckaegis ./cmd/scan
   ```

### Option Docker
1. Construire l'image :
   ```bash
   docker build -t pckaegis .
   ```
2. Exécuter en montant le répertoire courant :
   ```bash
   docker run --rm -v ${PWD}:/workspace pckaegis -file /workspace/requirements.txt -eco PyPI
   ```
   Adaptez `-file` et `-eco` selon l'écosystème (`PyPI` ou `npm`).

## Exemples d'utilisation

### Analyser un `requirements.txt`
```bash
./pckaegis -file ./requirements.txt -eco PyPI
```

### Analyser un `package.json`
```bash
./pckaegis -file ./package.json -eco npm
```

### Générer un rapport JSON
```bash
./pckaegis -file ./requirements.txt -eco PyPI -out resultat.json
```

### Mode non-bloquant (audit uniquement)
```bash
./pckaegis -file ./package.json -eco npm -strict=false
```

### Exécution directe sans compilation
```bash
go run ./cmd/scan -file ./requirements.txt -eco PyPI
```

## Notes
- Retourne `Exit Code 1` dès qu'une vulnérabilité est détectée pour faciliter le fail-fast en CI/CD.
- Pensez à exécuter l'analyse sur les deux écosystèmes si votre projet mélange Python et JavaScript.
- L'API OSV renvoyant un résumé minimal, enrichissez votre reporting en ajoutant une récupération de scores CVSS si nécessaire.

## Améliorations à venir 
Pour faire évoluer **pckAegis** vers une suite de sécurité plus robuste, les fonctionnalités suivantes sont prévues :
- [ ] **Rapports HTML** : Générer des audits de sécurité visuels pour les parties prenantes non techniques.
- [ ] **Webhooks** : Envoyer des alertes en temps réel aux équipes de développement lorsqu’une vulnérabilité est détectée.
- [ ] **Filtrage CVSS** : Implémenter un indicateur pour bloquer uniquement les scores supérieurs à 7,0.