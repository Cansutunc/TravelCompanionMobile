import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../common/global.dart';
import '../../../repositories/movie_repository.dart';
import '../../../widgets/cast_widget_loader.dart';
import '../bloc/route_bloc/route_cubit.dart';
import 'bottom_sheet.dart';

class RouteListView extends StatelessWidget {
  const RouteListView({super.key, required this.movieRepository});

  final MovieRepository movieRepository;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: BlocProvider(
        create: (_) => RouteCubit(
          repository: context.read<MovieRepository>(),
        ),
        child: RouteView(),
      ),
    );
  }
}

class RouteView extends StatelessWidget {
  const RouteView({super.key});

  void _addModal(BuildContext context2, String userId, String username) {
    showModalBottomSheet(
      context: context2,
      builder: (context) {
        return SingleChildScrollView(
          child: Container(
            color: Colors.transparent,
            padding: EdgeInsets.only(
              bottom: MediaQuery.of(context).viewInsets.bottom,
            ),
            child: CustomRouteModelBottomSheet(
                context2: context2, userId: userId, username: username),
          ),
        );
      },
    );
  }

  @override
  Widget build(BuildContext context) {
    String userId = Global.storageService.getUserId();
    String userName = Global.storageService.getUsername();
    bool isChanged = false;

    return DefaultTabController(
      length: 2,
      child: Scaffold(
        appBar: AppBar(
          title: const Text("Routes"),
          bottom: const TabBar(
              tabs: <Widget>[Tab(text: "Routes"), Tab(text: "History")]),
        ),
        body: TabBarView(
          children: [
            BlocBuilder<RouteCubit, RouteState>(
              buildWhen: (previous, current) {
                if (isChanged) {
                  isChanged = false;
                  return true;
                }
                if (current.routes.length != previous.routes.length) {
                  return true;
                }
                return false;
              },
              builder: (context, state) {
                switch (state.status) {
                  case ListStatus.failure:
                    return const Center(
                        child: Text(
                      'Oops something went wrong!',
                      style: TextStyle(color: Colors.white),
                    ));
                  case ListStatus.success:
                    return Scaffold(
                      floatingActionButton: FloatingActionButton(
                        child: Icon(Icons.add),
                        onPressed: () {
                          _addModal(context, userId, userName);
                        },
                      ),
                      body: ListView.builder(
                        itemCount: state.routes.length,
                        itemBuilder: (context, index) {
                          return Card(
                            child: ListTile(
                                leading: CircleAvatar(
                                  radius: 15,
                                  child: Text((index + 1).toString()),
                                ),
                                title: Text(state.routes[index].title),
                                subtitle: Text(state.routes[index].description),
                                trailing: Row(
                                  mainAxisSize: MainAxisSize.min,
                                  mainAxisAlignment:
                                      MainAxisAlignment.spaceEvenly,
                                  crossAxisAlignment: CrossAxisAlignment.center,
                                  children: [
                                    if (state.routes[index].createdUsername ==
                                        userName) ...[
                                      Text("You"),
                                      SizedBox(width: 10),
                                      IconButton(
                                        onPressed: () {
                                          context
                                              .read<RouteCubit>()
                                              .removeRoutes(
                                                  state.routes[index]);
                                        },
                                        icon: Icon(Icons.remove),
                                      ),
                                    ] else ...[
                                      Text(state.routes[index].createdUsername),
                                      SizedBox(width: 10),
                                      if (state.routes[index].allPersons
                                          .contains(userName)) ...[
                                        IconButton(
                                          onPressed: () {
                                            isChanged = true;
                                            context
                                                .read<RouteCubit>()
                                                .deleteRoute(
                                                    state.routes[index],
                                                    userName);
                                            context
                                                .read<RouteCubit>()
                                                .getRoute();
                                          },
                                          icon:
                                              const Icon(Icons.remove_circle_outline),
                                        ),
                                      ] else ...[
                                        if (state.routes[index].maxPerson >
                                            state.routes[index].allPersons
                                                .length) ...[
                                          IconButton(
                                            onPressed: () {
                                              isChanged = true;
                                              context
                                                  .read<RouteCubit>()
                                                  .acceptRoute(
                                                      state.routes[index],
                                                      userName);
                                              context
                                                  .read<RouteCubit>()
                                                  .getRoute();
                                            },
                                            icon: const Icon(Icons.plus_one_outlined),
                                          ),
                                        ] else ...[
                                          const Column(
                                            mainAxisAlignment:
                                                MainAxisAlignment.center,
                                            children: [
                                              Icon(Icons.add_box_outlined),
                                              Text("Full")
                                            ],
                                          )
                                        ]
                                      ]
                                    ]
                                  ],
                                )),
                          );
                        },
                      ),
                    );
                  default:
                    return Scaffold(
                      floatingActionButton: FloatingActionButton(
                        child: Icon(Icons.add),
                        onPressed: () {
                          _addModal(context, userId, userName);
                        },
                      ),
                      body: buildCastslistLoaderWidget(context),
                    );
                }
              },
            ),
            BlocBuilder<RouteCubit, RouteState>(
              buildWhen: (previous, current) {
                if (isChanged) {
                  isChanged = false;
                  return true;
                }
                if (current.routes.length != previous.routes.length) {
                  return true;
                }
                return false;
              },
              builder: (context, state) {
                switch (state.status) {
                  case ListStatus.failure:
                    return const Center(
                        child: Text(
                          'Oops something went wrong!',
                          style: TextStyle(color: Colors.white),
                        ));
                  case ListStatus.success:
                    return Scaffold(
                      body: ListView.builder(
                        itemCount: state.acceptedRoutes.length,
                        itemBuilder: (context, index) {
                          return Card(
                            child: ListTile(
                                leading: CircleAvatar(
                                  radius: 15,
                                  child: Text((index + 1).toString()),
                                ),
                                title: Text(state.acceptedRoutes[index].title),
                                subtitle: Text(state.acceptedRoutes[index].description),
                              trailing: IconButton(
                                onPressed: () {
                                  isChanged = true;
                                  context
                                      .read<RouteCubit>()
                                      .deleteRoute(
                                      state.acceptedRoutes[index],
                                      userName);
                                  context
                                      .read<RouteCubit>()
                                      .getRoute();
                                },
                                icon:
                                const Icon(Icons.remove_circle_outline),
                              ),
                            )
                          );
                        },
                      ),
                    );
                  default:
                    return Scaffold(
                      floatingActionButton: FloatingActionButton(

                        onPressed: () {
                          _addModal(context, userId, userName);
                        },
                        child: Icon(Icons.add),
                      ),
                      body: buildCastslistLoaderWidget(context),
                    );
                }
              },
            ),
          ],
        ),
      ),
    );
  }
}
