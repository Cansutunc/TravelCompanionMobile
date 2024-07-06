import 'package:casu_mobile/pages/home_screen/bloc/user_location_bloc/user_location_cubit.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:casu_mobile/repositories/movie_repository.dart';
import '../../../../common/global.dart';
import '../../../../widgets/cast_widget_loader.dart';

class UserLocationWidget extends StatefulWidget {
  const UserLocationWidget({super.key, required this.movieRepository});

  final MovieRepository movieRepository;

  @override
  State<UserLocationWidget> createState() => _UserLocationWidgetWidgetState();
}

class _UserLocationWidgetWidgetState extends State<UserLocationWidget> {
  @override
  Widget build(BuildContext context) {
    return RepositoryProvider.value(
      value: widget.movieRepository,
      child: const UserLocationBlocProvider(),
    );
  }
}

class UserLocationBlocProvider extends StatelessWidget {
  const UserLocationBlocProvider({super.key});

  @override
  Widget build(BuildContext context) {
    return BlocProvider(
      create: (_) => UserLocationCubit(
        repository: context.read<MovieRepository>(),
      ),
      child: UserLocationView(),
    );
  }
}

class UserLocationView extends StatelessWidget {
  UserLocationView({super.key});

  @override
  Widget build(BuildContext context) {
    return BlocBuilder<UserLocationCubit, UserLocationState>(
        buildWhen: (previous, current) {
      if (current.userLocations.isNotEmpty) {
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
          return SingleChildScrollView(
              child: ListView.builder(
                  shrinkWrap: true,
                  padding: const EdgeInsets.all(8),
                  itemCount: state.userLocations.length,
                  itemBuilder: (BuildContext context, int index) {
                    var li = state.hobbies
                        .map((e) => e.author == state.userLocations[index].user_name ? e.name: null)
                        .toList();
                    li.removeWhere((element) => element == null);

                    return Card(
                      child: ListTile(
                        leading: CircleAvatar(
                          radius: 15,
                          child: Text((index + 1).toString()),
                        ),
                        title: Text(state.userLocations[index].user_name),
                        subtitle: Text("Hobbies => ${li}"),
                        trailing: Column(
                          children: [
                            Text(state.userLocations[index].country),
                            Text(state.userLocations[index].province)
                          ],
                        ),
                      ),
                    );
                  }));
        default:
          return buildCastslistLoaderWidget(context);
      }
    });
  }
}
