# interop-core


curl --location --request GET 'http://localhost:8080/graphql' \
--header 'Content-Type: application/json' \
--data-raw '{"query":"query Users {\r\n\t\t\tusers {\r\n\t\t\t\tuserID\r\n\t\t\t\tusername\r\n\t\t\t\tnim\r\n\t\t\t}\r\n\t\t}","variables":{}}'

curl --location --request GET 'http://localhost:8080/graphql' \
--header 'Content-Type: application/json' \
--data-raw '{"query":"query User($userID: ID!) {\r\n\t\t\tuser(userID: $userID) {\r\n\t\t\t\tuserID\r\n\t\t\t\tusername\r\n\t\t\t}\r\n\t\t}","variables":{"userID":"u-001"}}'
