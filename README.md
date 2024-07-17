# weather-api

### Verzija GO-a

 - 1.22.5

Preprost REST API implementiran v go programskem jeziku, z Gin ogrodjem, ki vrne informacije o vremenu.

### Endpoint 

- GET /weather : vrne vreme, za vsa mesta, ki so definirana v arso-xmls.json
- GET /weather/{city} : vrne vreme, za mesto, ki je podano, ce obstaja v arso-xmls.json

### Dodajanje novih mest

- Aplikacija omogoča dodajanje novih mest, tudi ko je že v produkciji. Podatke o mestih in URL-jih, od koder se pridobivajo vremenski podatki, hranimo v JSON datoteki (arso-xmls.json) -> simulacija baze.

### Pohitritev vrnitve podatkov

- Aplikacija ima implementiran cache mehanizem, ki omogča hitrejšo vrnitev podatkov. Saj ob zahtevi za /weather se podatki shranijo v cache. Tako dobimo hitrejši odgovor, saj ne rabimo podatke pridobiti iz xml in jih razčlenit v pravilno obliko. 
- Cache se prazni na vsakih 20 min, da zagotovimo, da so podatki pravilni.

### Prevajanje (Build)
 
- make build

### Zagon lokalno

- make run <- build-a in zažene

### Zagon z Dokerjem

- docker build -t weather-api:1.0 .    
- docker run -p 8080:8080 weather-api:1.0

### Zagon testov

- make test