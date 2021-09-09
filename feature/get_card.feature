Feature: Get card
    
    In order to know card information
    I need to get them/it first

    Scenario: Card is empty
        Given the card owned by user 1 is empty
        When I get all cards owned by user 1
        Then response status code must be 200
        And response must match json
            """
            {
                "cards": []
            }
            """
    
    Scenario: Many cards exists
        Given there are cards owned by user 1 with body
            | {"walletId": "oWx0b8DZ1a", "limitDaily": 1000000, "limitMonthly": 20000000} |
            | {"walletId": "oWx0b8DZ1a", "limitDaily": 2000000, "limitMonthly": 30000000} |
        When I get all cards owned by user 1
        Then response status code must be 200
        And number of cards retrieved must be 2

    Scenario: Get single card when card doesnt exist
        Given the card owned by user 1 is empty
        When I get single card with id 1 owned by user 1
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
    
    Scenario: Get single card when card exists
        Given there are cards owned by user 1 with body
            | {"walletId": "oWx0b8DZ1a", "limitDaily": 1000000, "limitMonthly": 20000000} |
        When I get single card with index 0 owned by user 1
        Then response status code must be 200
        And response must be single card
