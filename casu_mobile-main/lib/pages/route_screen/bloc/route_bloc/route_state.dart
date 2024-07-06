part of 'route_cubit.dart';

enum ListStatus { loading, success, failure }


class RouteState extends Equatable {
  final ListStatus status;
  final List<Routes> routes;
  final List<Routes> acceptedRoutes;

  RouteState({
    this.status = ListStatus.loading,
    this.routes = const <Routes>[],
    this.acceptedRoutes = const <Routes>[],
  });



  @override
  // TODO: implement props
  List<Object?> get props => [status,routes,acceptedRoutes];
}


class RouteStateInitial extends RouteState{}