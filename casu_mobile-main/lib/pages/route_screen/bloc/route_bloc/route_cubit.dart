
import 'package:equatable/equatable.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../common/global.dart';
import '../../../../model/route.dart';
import '../../../../repositories/movie_repository.dart';


part 'route_state.dart';


class RouteCubit extends Cubit<RouteState> {
  RouteCubit({required this.repository}) : super(RouteState()){
    getRoute();
  }
  final MovieRepository repository;
  String username = Global.storageService.getUsername();

  Future<void> getRoute() async {
    try {
      final res = await repository.getRoutes();
      final res2 = await repository.getAcceptedRoutes(username);
      emit(RouteState(status:ListStatus.success,routes: res.routes,acceptedRoutes: res2.routes));
    } on Exception {
    }
  }

  Future<void> addRoute(Routes route) async {
    try {
      bool response = await repository.addRoute(route);
      final state = this.state;
      if (response) {
        emit(RouteState(status:ListStatus.success,routes: List.from(state.routes)..add(route)));
      }
    } on Exception {
      // emit(const HobbieState.failure());
    }
  }

  Future<void> removeRoutes(Routes route) async {
    try {
      await repository.removeRoute(route.city_id);
      final state = this.state;
      emit(RouteState(status:ListStatus.success,routes: List.from(state.routes)..remove(route)));
    } on Exception {
      // emit(const HobbieState.failure());
    }
  }

  Future<void> acceptRoute(Routes route,String username) async {
    try {
      await repository.acceptRoute(route.city_id,username);
      final state = this.state;
      // emit(RouteState(status:ListStatus.success,routes: List.from(state.routes)..remove(route)));
    } on Exception {
      // emit(const HobbieState.failure());
    }
  }

  Future<void> deleteRoute(Routes route,String username) async {
    try {
      await repository.deleteRoute(route.city_id,username);
      final state = this.state;
      // emit(RouteState(status:ListStatus.success,routes: List.from(state.routes)..remove(route)));
    } on Exception {
      // emit(const HobbieState.failure());
    }
  }


}


