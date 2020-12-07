import 'package:flutter/material.dart';
import 'package:graphql/client.dart';

class GQLClient {
  GQLClient({@required graphQLClient}) : _graphQLClient = graphQLClient;

  factory GQLClient.create() {
    final httpLink = HttpLink(uri: 'http://localhost:8080/query');
    final link = Link.from([httpLink]);
    return GQLClient(
      graphQLClient: GraphQLClient(cache: InMemoryCache(), link: link),
    );
  }

  final GraphQLClient _graphQLClient;

  GraphQLClient get client => _graphQLClient;
}
