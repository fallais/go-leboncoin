# Go-LeBonCoin

![Coop](https://github.com/fallais/go-leboncoin/blob/master/gopher.png)

**go-leboncoin** is a tool that is able to search ads in **Le Bon Coin** and informs you of new ads. The development is in **progress**

## Filters

The idea is to create filters based on the URL that is built when you search on *Le Bon Coin*. The filter has a `name`, a `is_enabled` boolean, and the URL.

```yaml
filters:
  - name: moto
    is_enabled: true
    url: https://www.leboncoin.fr/recherche/?category=3&locations=d_31&moto_type=moto&price=1000-2500&cubic_capacity=500-600
```

## Notifications

TODO :

- Emails notification
- SMS notification
- IFTTT
- ...

## Contributing

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request

## Credits

Implemented by François Allais
