const { ApolloServer } = require("@apollo/server");
const { startStandaloneServer } = require("@apollo/server/standalone");
const { ApolloGateway, IntrospectAndCompose } = require("@apollo/gateway");

// TODO: https://www.apollographql.com/docs/apollo-server/using-federation/apollo-gateway-setup

(async function () {
  // TODO: replace `IntrospectAndCompose` with something more stable for prod
  const gateway = new ApolloGateway({
    supergraphSdl: new IntrospectAndCompose({
      subgraphs: [
        { name: "feedback", url: "http://localhost:4001/query" },
        { name: "users", url: "http://localhost:4002/query" },
        // ...additional subgraphs...
      ],
    }),
  });

  const server = new ApolloServer({
    gateway,
    subscriptions: false,
  });

  const { url } = await startStandaloneServer(server);
  console.log(`🚀  Server ready at ${url}`);
})();
