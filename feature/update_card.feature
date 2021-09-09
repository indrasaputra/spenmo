Feature: Enable card
    
    In order to update card
    I need to update it first

    Scenario: Non-exists card can't be enabled
        Given the card owned by user 1 is empty
        When I update card with id 1 owned by user 1 with body
            | {"limitDaily": 1000000, "limitMonthly": 20000000} |
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

    Scenario: Cant update card with invalid limit input
        Given there are cards owned by user 1 with body
            | {"walletId": "oWx0b8DZ1a", "limitDaily": 1000000, "limitMonthly": 20000000} |
        When I update card with index 0 owned by user 1 with body
            | {"limitDaily": 0, "limitMonthly": 20000000} |
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

    Scenario: Success update card
        Given there are cards owned by user 1 with body
            | {"walletId": "oWx0b8DZ1a", "limitDaily": 1000000, "limitMonthly": 20000000} |
        When I update card with index 0 owned by user 1 with body
            | {"limitDaily": 2000000, "limitMonthly": 50000000} |
        Then response status code must be 200
        And response must match json
            """
            {}
            """
