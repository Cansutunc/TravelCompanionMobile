part of 'user_location_cubit.dart';
enum ListStatus { loading, success, failure }

class UserLocationState extends Equatable {
  const UserLocationState._({
    this.status = ListStatus.loading,
    this.userLocations = const <UserLocation>[],
    this.hobbies = const <Hobbie>[],
  });

  const UserLocationState.loading() : this._();

  const UserLocationState.success(List<UserLocation> userLocations,List<Hobbie> hobbies)
      : this._(status: ListStatus.success,userLocations: userLocations,hobbies: hobbies);

  const UserLocationState.failure() : this._(status: ListStatus.failure);

  final ListStatus status;

  final List<UserLocation> userLocations;
  final List<Hobbie> hobbies;

  @override
  List<Object> get props => [status, userLocations,hobbies];
}