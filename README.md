# GOstructure

This repository proposes to be an scaffolding repo to get some architectural patterns applied on go

## Architectures

To exemplify these patterns, I'll solve a single `BDD` scenario aplying each of below patterns:

- DDD
- MVC (developing)

### Future:

- Clean architecture
- Transaction script


## Problem

The BDD is:

```
Given I want to reserve a book in a book store
When I reserve the book
And there is just one available book
Then the book has to be unavailable
And an email must be send to inform that the user has five days to return it
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
