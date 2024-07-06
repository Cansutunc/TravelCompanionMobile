


import 'package:casu_mobile/pages/route_screen/widget/route_list_view.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../repositories/movie_repository.dart';
import '../../search_screen/widgets/search_list_view.dart';

class RouteListWidget extends StatefulWidget {
  const RouteListWidget(
      {super.key,
        required this.movieRepository,
      });

  final MovieRepository movieRepository;

  @override
  State<RouteListWidget> createState() => _RouteListWidgetState();
}

class _RouteListWidgetState extends State<RouteListWidget> {
  @override
  Widget build(BuildContext context) {
    return RepositoryProvider.value(
      value: widget.movieRepository,
      child: RouteListView(
        movieRepository: widget.movieRepository,
      ),
    );
  }
}