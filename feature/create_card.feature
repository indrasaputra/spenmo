Feature: Create new user card
    
    In order to wallet
    I need to create a card first

    Scenario: Invalid json request body
        Given the card owned by user 1 is empty
        When I create card owned by user 1 with body
            | string |
        Then response status code must be 400
        And response must match json
            """
            {
                "code": 3,
                "message": "invalid character 's' looking for beginning of value",
                "details": []
            }
            """

    Scenario: Wallet ID is not hashid
        Given the card owned by user 1 is empty
        When I create card owned by user 1 with body
            | {"walletId": "abc", "limitDaily": 1000000, "limitMonthly": 20000000} |
            | {"walletId": "def", "limitDaily": 1000000, "limitMonthly": 20000000} |
        Then response status code must be 400
        And response must match json
            """
            {
                "code": 3,
                "message": "",
                "details": [
                    {
                        "@type": "type.googleapis.com/google.rpc.BadRequest",
                        "fieldViolations": [
                            {
                            "field": "walletId",
                            "description": "wallet is invalid"
                            }
                        ]
                    },
                    {
                        "@type": "type.googleapis.com/proto.indrasaputra.spenmo.v1.SpenmoCardError",
                        "errorCode": "INVALID_WALLET"
                    }
                ]
            }
            """

    Scenario: Invalid json value for limit
        Given the card owned by user 1 is empty
        When I create card owned by user 1 with body
            | {"walletId": "oWx0b8DZ1a", "limitDaily": 0, "limitMonthly": 20000000} |
            | {"walletId": "oWx0b8DZ1a", "limitDaily": 1000000, "limitMonthly": 0} |
        Then response status code must be 400
        And response must match json
            """
            {
                "code": 3,
                "message": "",
                "details": [
                    {
                        "@type": "type.googleapis.com/google.rpc.BadRequest",
                        "fieldViolations": [
                            {
                            "field": "limit*",
                            "description": "limit is invalid. it must be greater than 0"
                            }
                        ]
                    },
                    {
                        "@type": "type.googleapis.com/proto.indrasaputra.spenmo.v1.SpenmoCardError",
                        "errorCode": "INVALID_LIMIT"
                    }
                ]
            }
            """

    Scenario: Valid json request body
        Given the card owned by user 1 is empty
        When I create card owned by user 1 with body
            | {"walletId": "oWx0b8DZ1a", "limitDaily": 1000000, "limitMonthly": 20000000} |
        Then response status code must be 200
        And response must match json
            """
            {}
            """
