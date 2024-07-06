import 'package:casu_mobile/pages/home_screen/widgets/location_widgets/location_widget.dart';
import 'package:casu_mobile/pages/home_screen/widgets/user_location_widgets/user_location_widget.dart';
import 'package:flutter/material.dart';
import 'package:casu_mobile/repositories/movie_repository.dart';

class HomeScreen extends StatefulWidget {
  const HomeScreen({super.key, required this.movieRepository});

  final MovieRepository movieRepository;

  @override
  State<HomeScreen> createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: ListView(
        padding: EdgeInsets.zero,
        children: [
          // UpComingWidget(
          //     movieRepository: widget.movieRepository),
          const Padding(
            padding: EdgeInsets.only(left: 15.0,top: 40.0),
            child: Text("Your Location"),
          ),
          LocationWidget(movieRepository: widget.movieRepository),
          const Padding(
            padding: EdgeInsets.only(top: 10.0,left: 15.0),
            child: Text("People Around You"),
          ),
          UserLocationWidget(movieRepository: widget.movieRepository),



        ],
      ),
    );
  }
}
