## Assumptions

1. The scopes of the API are:
    
    a. Create user's card
    
    b. Update user's card information
    
    c. Get single user's card
    
    d. Get all user's card
    
    e. Delete a single user's card

2. To call the API, authentication is required. For simplicity, authentication is using HTTP Authorization header with user's ID (integer) as value.

    ```
    $ curl http://localhost:8081/v1/users/cards --header 'Authorization: 1'
    ```

3. User's card depends on user and user's wallet. Since this service focuses only on user's card, we need to seed the data for user and user's wallet.
    The data are automatically inserted during the very first database migration process.
    The data and query can be seen in [20210908163105_seed_for_users_and_wallets.up.sql](../db/migrations/20210908163105_seed_for_users_and_wallets.up.sql).
    There will be 5 users with ID 1 to 5. Each user has 5 wallets.

4. Rate limit is applied per user ID.

5. Any unique identifier (ID) is exposed to user/public using [hashids](https://hashids.org/). Hashids is choosen because it only obfuscates the output while we still can use integer internally. Integer is choosen over UUID due to its simplicity, easy to remember, and easy to find, such as in a log.

6. The only place that the ID doesn't use hashids is in authorization. The reason is authorization is already encoded. For example, we can use OAuth2. In OAuth2, there is a choice to use a random string as token. We can exchange the token with user ID internally.

    But, in this service, for simplicity and due to its scope, the authorization is made simple as explained in #2.

7. Card deletion is soft delete.