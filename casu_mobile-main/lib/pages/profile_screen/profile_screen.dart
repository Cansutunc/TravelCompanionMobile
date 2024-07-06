import 'package:casu_mobile/pages/hobbie_screen/hobbie_screen.dart';
import 'package:casu_mobile/repositories/movie_repository.dart';
import 'package:flutter/material.dart';

import '../../common/constants.dart';
import '../../common/global.dart';

class ProfileScreen extends StatefulWidget {
  const ProfileScreen({super.key, required this.profileName,required this.movieRepository});

  final String profileName;
  final MovieRepository movieRepository;

  @override
  State<ProfileScreen> createState() => _ProfileScreenState();
}

class _ProfileScreenState extends State<ProfileScreen> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          title: const Text(
            "Profile",
            style: TextStyle(color: Colors.black),
          ),
        ),
        body: SafeArea(
          child: ListView(
            padding: const EdgeInsets.only(left: 50, top: 100, right: 50),
            children: [
              CircleAvatar(
                  radius: 50.0,
                  child: ClipOval(
                    child: Image.asset(
                      fit: BoxFit.cover,
                      "assets/img/cast_placeholder.png",
                    ),
                  )),
              Padding(
                padding: const EdgeInsets.only(left: 100),
                child: Text(widget.profileName),
              ),
              ElevatedButton(
                  onPressed: () {
                    Navigator.push(
                      context,
                      MaterialPageRoute(
                        builder: (context) => HobbieScreen(movieRepository: widget.movieRepository,),
                      ),
                    );
                  },
                  child: Text("Hobbies")),
              ElevatedButton(
                  onPressed: () {
                    Global.storageService
                        .remove(AppConstants.STORAGE_USER_TOKEN_KEY);
                    Navigator.of(context)
                        .pushNamedAndRemoveUntil("/", (route) => false);
                  },
                  child: Text("Çıkış Yap"))
            ],
          ),
        ));
  }
}
