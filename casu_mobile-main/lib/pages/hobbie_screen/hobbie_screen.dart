import 'package:casu_mobile/pages/hobbie_screen/widget/hobbie_view.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../repositories/movie_repository.dart';
import 'bloc/hobbie_cubit.dart';


class HobbieScreen extends StatelessWidget {
  const HobbieScreen(
      {super.key,
        required this.movieRepository});
  final MovieRepository movieRepository;

  @override
  Widget build(BuildContext context) {
    return HobbieRepoProvider(
      movieRepository: movieRepository,
    );
  }
}

class HobbieRepoProvider extends StatefulWidget {
  const HobbieRepoProvider(
      {super.key,
        required this.movieRepository,});

  final MovieRepository movieRepository;

  @override
  State<HobbieRepoProvider> createState() => _HobbieRepoProviderWidgetState();
}

class _HobbieRepoProviderWidgetState extends State<HobbieRepoProvider> {
  @override
  Widget build(BuildContext context) {
    return RepositoryProvider.value(
      value: widget.movieRepository,
      child: HobbieBlocProvider(
        movieRepository: widget.movieRepository,
      ),
    );
  }
}

class HobbieBlocProvider extends StatelessWidget {
  const HobbieBlocProvider(
      {super.key,
        required this.movieRepository});
  final MovieRepository movieRepository;

  @override
  Widget build(BuildContext context) {
    return BlocProvider(
        create: (_) => HobbieCubit(
          repository: context.read<MovieRepository>(),
        )..fetchHobbies(),
        child: HobbieView(
          movieRepository: movieRepository,
        ),
      );

  }
}