Feature: Delete card
    
    In order to remove card
    I need to delete it first

    Scenario: Non-exists card can't be deleted
        Given the card owned by user 1 is empty
        When I delete card with id 1 owned by user 1
        Then response status code must be 404
        And response must match json
            """
            {
                "code": 5,
                "message": "",
                "details": [
                    {
                        "@type": "type.googleapis.com/proto.indrasaputra.spenmo.v1.SpenmoCardError",
                        "errorCode": "NOT_FOUND"
                    }
                ]
            }
            """

    Scenario: Newly created card can be deleted
        Given there are cards owned by user 1 with body
            | {"walletId": "oWx0b8DZ1a", "limitDaily": 1000000, "limitMonthly": 20000000} |
        When I delete card with index 0 owned by user 1
        Then response status code must be 200
        And response must match json
            """
            {}
            """