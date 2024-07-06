import 'package:casu_mobile/pages/home_screen/bloc/location_bloc/location_cubit.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:casu_mobile/repositories/movie_repository.dart';

import '../../../../widgets/cast_widget_loader.dart';

class LocationWidget extends StatefulWidget {
  const LocationWidget({super.key, required this.movieRepository});

  final MovieRepository movieRepository;

  @override
  State<LocationWidget> createState() => _LocationWidgetWidgetState();
}

class _LocationWidgetWidgetState extends State<LocationWidget> {
  @override
  Widget build(BuildContext context) {
    return RepositoryProvider.value(
      value: widget.movieRepository,
      child: const LocationBlocProvider(),
    );
  }
}

class LocationBlocProvider extends StatelessWidget {
  const LocationBlocProvider({super.key});

  @override
  Widget build(BuildContext context) {
    return BlocProvider(
      create: (_) => LocationCubit(
        repository: context.read<MovieRepository>(),
      ),
      child: const LocationView(),
    );
  }
}

class LocationView extends StatelessWidget {
  const LocationView({super.key});

  @override
  Widget build(BuildContext context) {
    return BlocBuilder<LocationCubit, LocationState>(
        buildWhen: (previous, current) {
      if (current.location.country.isNotEmpty) {
        return true;
      }
      return false;
    }, builder: (context, state) {
      switch (state.status) {
        case ListStatus.failure:
          return const Center(
              child: Text(
            'Oops something went wrong!',
            style: TextStyle(color: Colors.white),
          ));
        case ListStatus.success:
          return Card(
            color: Colors.transparent,
            child: Column(children: [
              Text(state.location.country),
              Text(state.location.province),
              Text(state.location.town),
              Text(state.location.road),
              Text(state.location.suburb),
              Text(state.location.amenity),
            ],),
          );
        default:
          return buildCastslistLoaderWidget(context);
      }
    });
  }
}
