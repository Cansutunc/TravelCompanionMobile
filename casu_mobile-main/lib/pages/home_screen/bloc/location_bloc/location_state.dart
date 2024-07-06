part of 'location_cubit.dart';
enum ListStatus { loading, success, failure }

class LocationState extends Equatable {
  const LocationState._({
    this.status = ListStatus.loading,
    this.location = Location.empty,

  });

  const LocationState.loading() : this._();

  const LocationState.success(Location location)
      : this._(status: ListStatus.success, location: location);

  const LocationState.failure() : this._(status: ListStatus.failure);

  final ListStatus status;
  final Location location;


  @override
  List<Object> get props => [status, location];
}