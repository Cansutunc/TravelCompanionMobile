

import 'package:casu_mobile/pages/search_screen/widgets/search_list_view.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../repositories/movie_repository.dart';

class SearchListWidget extends StatefulWidget {
  const SearchListWidget(
      {super.key,
        required this.movieRepository,
        });

  final MovieRepository movieRepository;

  @override
  State<SearchListWidget> createState() => _MovieDetailWidgetState();
}

class _MovieDetailWidgetState extends State<SearchListWidget> {
  @override
  Widget build(BuildContext context) {
    return RepositoryProvider.value(
      value: widget.movieRepository,
      child: SearchListView(
        movieRepository: widget.movieRepository,
      ),
    );
  }
}