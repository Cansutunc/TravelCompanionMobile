
import 'package:casu_mobile/pages/main_screen/main_screen.dart';
import 'package:casu_mobile/pages/register_screen/register_screen.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

import '../../pages/login_screen/login_screen.dart';
import '../../repositories/movie_repository.dart';
import '../global.dart';

MovieRepository movieRepository = MovieRepository();

class AppPages {



  static MaterialPageRoute generateRouteSettings(RouteSettings settings) {

    if(settings.name=="/"){
      bool isLogged = Global.storageService.getIsLoggedIn();
      if (isLogged){
        String profileName = Global.storageService.getUsername();
        return MaterialPageRoute(
            builder: (_) => MainScreen(profileName: profileName,movieRepository: movieRepository ), settings: settings);
      }
    }
    if (settings.name == "/register") {
      return MaterialPageRoute(
          builder: (_) => const RegisterScreen(), settings: settings);
    }
    else if(settings.name  =="/main"){
      String profileName = Global.storageService.getUsername();
      return MaterialPageRoute(

          builder: (_) => MainScreen(profileName:profileName,movieRepository: movieRepository ), settings: settings);
    }

    return MaterialPageRoute(
        builder: (_) => const LoginScreen(), settings: settings);
  }
}
