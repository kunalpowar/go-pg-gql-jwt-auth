import 'package:app/gql/gql.dart';
import 'package:flutter/material.dart';

void main() {
  runApp(App(
    gqlClient: GQLClient.create(),
  ));
}

class App extends StatelessWidget {
  const App({Key key, this.gqlClient}) : super(key: key);

  final GQLClient gqlClient;

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter GraphQL JWT Demo',
      home: Scaffold(
        body: Center(
          child: Text('WIP'),
        ),
      ),
    );
  }
}
