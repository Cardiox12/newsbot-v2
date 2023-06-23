# Project spec

The bot must be designed to : 
- be self-hosted
- be highly configurable :
    - configure subscribing schedule time
    - configure subscribers infos (API KEY, SECRETS, ...)
- connect existing subscribers :
    - discord
    - slack
    - ...
- be modular
- let the community write their own subscriber
- schedule collection of data sources
    ```yaml
    hackernews: every.day.at(10).hour().timezone("America/Los_Angeles")
    twitter: every.day.at(18).hour().timezone("Europe/Paris")
    ```
- define subscriber failure policy
    - define number of retry
    - define retry timeout

# Data

A collection is a set of "Article".
An article is an object with the following fields :
```json
{
    "title"     : "",
    "url"       : "",
    "source"    : "",
}
```


## Database spec
- Configure retention time [TTL](https://www.mongodb.com/docs/manual/tutorial/expire-data/)

## Data spec
- Create a collection
- Get a collection
- Get latest collection
- Add article to collection

# Domains
**Article** : Un article est un texte qui relate un événement, présente des faits ou expose un point de vue.

**Source** : Personne ou organisme à l’origine d’une information.

**Légende** : Texte court qui accompagne une photo, un dessin, une infographie.

**Information chaude** : Information qui traite de ce qui vient ou est en train de se produire.

**Information froide** : Information dont la parution peut être différée.

**Information sensible** : Une information sensible est une information ou une connaissance qui, si elle est révélée au 
public, aurait une influence sur les entités qu’elle concerne.

**Éditeur** : Professionnel qui publie et diffuse des ouvrages imprimés.

**Publier** : Action de mettre un journal à la disposition du public.

**News Cycle** :
The period of time that elapses before one news story or set of stories is replaced by another. Historically, news 
cycles have often been based on the publication deadlines of daily newspapers, typically lasting 24 hours, but for 
much of the public has continually decreased in duration with the emergence of radio, television, the Internet, and 
social media, such that cycles of definitive length have largely been supplanted by rolling news.

**News Aggregator** : In computing, a news aggregator, also termed a feed aggregator, feed reader, news reader, RSS 
reader, or simply an aggregator is client software or a web application that aggregates syndicated web content such as 
online newspapers, blogs, podcasts, and video blogs (vlogs) in one location for easy viewing. The updates distributed
may include journal tables of contents, podcasts, videos, and news items.
