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