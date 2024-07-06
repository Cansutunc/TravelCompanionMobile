

import 'package:casu_mobile/pages/route_screen/widget/route_list_widget.dart';
import 'package:flutter/cupertino.dart';

import '../../repositories/movie_repository.dart';

class RouteScreen extends StatelessWidget {
  const RouteScreen(
      {super.key,
        required this.movieRepository});
  final MovieRepository movieRepository;

  @override
  Widget build(BuildContext context) {
    return RouteListWidget(
      movieRepository: movieRepository,
    );
  }
}