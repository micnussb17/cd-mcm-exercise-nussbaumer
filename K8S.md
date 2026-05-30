Health Checks

Readiness vs Liveness probe: what's the difference?

Die Readiness Probe entscheidet, ob ein Pod Traffic empfangen darf. Die Liveness Probe entscheidet, ob ein Pod neu gestartet werden muss. Readiness schuetzt Benutzer vor nicht betriebsbereiten Pods, waehrend Liveness die Verfuegbarkeit durch automatische Neustarts fehlerhafter Pods erhoeht.

What happens when each probe fails?

Wenn eine Readiness Probe fehlschlaegt:

- Der Pod wird als Not Ready markiert.
- Kubernetes entfernt den Pod aus den Service Endpoints.
- Neue Anfragen werden nicht mehr an diesen Pod weitergeleitet.
- Der Pod laeuft weiterhin und wird nicht neu gestartet.
- Sobald die Probe wieder erfolgreich ist, erhaelt der Pod wieder Traffic.

Wenn eine Liveness Probe fehlschlaegt:

- Kubernetes geht davon aus, dass die Anwendung fehlerhaft oder festgefahren ist.
- Der Container wird automatisch beendet.
- Kubernetes startet den Container bzw. Pod neu.
- Dadurch kann sich die Anwendung selbststaendig von Fehlern erholen.

Why different initialDelaySeconds values?

Die Readiness Probe startet meist frueher, damit Kubernetes schnell erkennen kann, wann ein Pod Anfragen verarbeiten darf. Die Liveness Probe startet spaeter, um der Anwendung genuegend Zeit zum Starten zu geben und unnoetige Neustarts zu vermeiden.

Resource Limits

What happens if memory/CPU limit is exceeded?

Wird das CPU-Limit Ueberschritten, wird der Container gedrosselt (throttling). Wird das Memory-Limit ueberschritten, wird der Container von Kubernetes beendet und neu gestartet (Out Of Memory Kill).

Why specify both requests and limits?

Requests reservieren die benoetigten Ressourcen fuer einen Pod und helfen Kubernetes bei der Planung. Limits begrenzen den maximalen Ressourcenverbrauch, damit ein Pod nicht alle verfuegbaren Ressourcen des Clusters verbraucht.

Screenshots befinden sich in der beigelegten PDF ex04_Nußbaumer.pdf