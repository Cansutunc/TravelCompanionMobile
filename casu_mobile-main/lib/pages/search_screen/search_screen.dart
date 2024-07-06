
import 'package:casu_mobile/pages/search_screen/widgets/search_list_widget.dart';
import 'package:flutter/cupertino.dart';

import '../../repositories/movie_repository.dart';

class SearchScreen extends StatelessWidget {
  const SearchScreen(
      {super.key,
        required this.movieRepository});
  final MovieRepository movieRepository;

  @override
  Widget build(BuildContext context) {
    return SearchListWidget(
      movieRepository: movieRepository,
    );
  }
}