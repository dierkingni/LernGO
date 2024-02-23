# I am learning GO in this repo

1) Erstellen einer Calculator Cloud Function API

Erstelle eine kleine REST-API, welche zwei Zahlen addieren bzw. subtrahieren kann (jeweils ein Endpunkt).
Die zwei Zahlen sollen als Path-Parameter übergeben werden. Das Ergebnis soll sowohl als Response zurückgegeben als auch in einem Cloud Storage Bucket als JSON gespeichert werden.
Zudem sollen die API-Endpunkte mittels einer OpenAPI Spec beschrieben werden und diese über einen zusätzlichen API-Endpunkt aufrufbar sein.

Teck-Stack
- Sprache: Go
- Verwende Cloud Functions als Platform
- Terraform für die notwendige Cloud Infrastruktur
- Automatisches Deployment mittels GitHub Actions