

[![N|Solid](https://www.jojonomic.com/wp-content/uploads/2020/05/logo-jojonomic-with-ISO@2x-1.png)](https://jojonomic.com)

# Coding Test Jojonomic
Create all endpoint bellow. Every request has Authorization Header
```
Authorization:  Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJKb2pvbm9taWMiLCJpYXQiOjE2MDY2OTYyOTYsImV4cCI6MTYzODIzMjI5NiwiYXVkIjoiam9qb25vbWljLmNvbSIsInN1YiI6Impvam9hcmllZiIsImNvbXBhbnlfaWQiOiIxMzAiLCJ1c2VyX2lkIjoiMTIwIn0.jdnqi7hHmeeQlOJ1o-nZBxynTah-QuDj-SjItbol8XE
```
Note: The JWT Token info is:
```
{
  "iss": "Jojonomic",
  "iat": 1606696296,
  "exp": 1638232296,
  "aud": "jojonomic.com",
  "sub": "jojoarief",
  "company_id": "130",
  "user_id": "120"
}
```
JWT Private Key:
```
qwertyuiopasdfghjklzxcvbnm123456
```
You must validate every request by their JWT token. If the token is valid, you can use that information for authorize the document service. You can createw your own JWT Token builder for the test.

## Endpoints
### Root List
Get list of a root of document. The responses will show all document and folder from the user's company. The company_id is provided by jwt token.
```
GET http://api-gateway.co.id/document-service
```
Response:
```
{
    "error": false,
    "data": [
         {
             "id": "82b07a6f-60cc-4403-8fd2-329ef0de0d3d",
             "name": "Folder Tech",
             "type": "folder",
             "is_public": true,
             "owner_id": 1231,
             "share": [123,232,4333],
             "timestamp": 16576232323
         },
         {
             "id": "82b07a6f-60cc-4403-8fd2-329ef0de045d",
             "name": "Folder hrd",
             "type": "folder",
             "is_public": true,
             "owner_id": 1231,
             "share": [123,232,4333],
             "timestamp": 16576232323
         },
        {
             "id": "82b07a6f-60cc-4403-8fd2-329ef0de045d",
             "name": "Document Job desc Tech",
             "type": "document",
             "owner_id": 1231,
             "share": [123,232,4333],
             "timestamp": 16576232323
         }
    ]
}
```
If is_public is true, all user from all company can see the file or folder. If is_public is false, only owner and user in share can see the documents.
