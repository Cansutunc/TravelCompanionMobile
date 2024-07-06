part of 'hobbie_cubit.dart';

enum ListStatus { loading, success, failure }


class HobbieState extends Equatable {
  final ListStatus status;
  final List<Hobbie> hobbies;

  HobbieState({
    this.status = ListStatus.loading,
    this.hobbies = const <Hobbie>[],
  });



  @override
  // TODO: implement props
  List<Object?> get props => [status,hobbies];
}


class HobbieStateInitial extends HobbieState{}