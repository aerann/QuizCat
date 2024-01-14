## QuizCat
AI powered studying assistant that generates flash cards automatically based on your text input.

Built using AngularJS + TypeScript for the user-interface, and Go for the REST API. Making use of Cohere's LLM service for text generation.

### Run locally
Client 
```
cd client
npm install
ng serve
```

Server
```
cd server
go run cmd/api/main.go
```

