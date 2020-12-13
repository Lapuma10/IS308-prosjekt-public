# is308-prosjekt
This is the codebase for our automatic Shopify &amp; Fiken accounting software. 

## Setup Guide 
For at tjenestene skal kunne motta meldinger, så må en instans av RabbitMQ startes:
```
docker run --name rabbitmq -p 5672:5672 rabbitmq
```

Messaging service bruker en eksternal library, så vi anbefaler følgende nedlastinger:
```
go get github.com/streadway/amqp
```


## Tjenester
### web-tjeneste 
registrering av brukere med config info
### bruker-tjeneste
oppdaterer fiken med nye shopify brukere
### transaksjon-tjeneste
oppdaterer fiken med nye transaksjoner (vipps, stripe, shopify)
### logging-tjeneste
mottar eventuelle feilmeldinger, også historikk for handlinger 
### messaging-tjeneste
leser cronjob info, sender meldinger til meldingskøen til rabbitMQ
