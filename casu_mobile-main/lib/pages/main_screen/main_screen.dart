import 'package:casu_mobile/common/constants.dart';
import 'package:casu_mobile/pages/route_screen/route_screen.dart';
import 'package:casu_mobile/pages/search_screen/search_screen.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:casu_mobile/pages/main_screen/bottom_navbar_bloc.dart';
import 'package:casu_mobile/repositories/movie_repository.dart';
import '../../common/global.dart';
import '../home_screen/home_screen.dart';
import '../profile_screen/profile_screen.dart';

class MainScreen extends StatefulWidget {
  const MainScreen(
      {super.key,required this.profileName, required this.movieRepository});

  final MovieRepository movieRepository;
  final String profileName;

  @override
  State<MainScreen> createState() => _MainScreenState();
}

class _MainScreenState extends State<MainScreen> {
  final BottomNavBarBloc _bottomNavBarBloc = BottomNavBarBloc();
  late bool isDarkMode;

  @override
  void initState() {
    super.initState();
  }

  @override
  void dispose() {
    _bottomNavBarBloc.close();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return AnnotatedRegion<SystemUiOverlayStyle>(
      value: const SystemUiOverlayStyle(
        // For Android.
        // Use [light] for white status bar and [dark] for black status bar.
        statusBarIconBrightness: Brightness.light,
        // For iOS.
        // Use [dark] for white status bar and [light] for black status bar.
        statusBarBrightness: Brightness.dark,
      ),
      child: Scaffold(
          body: StreamBuilder<NavBarItem>(
            stream: _bottomNavBarBloc.itemStream,
            initialData: _bottomNavBarBloc.defaultItem,
            builder:
                (BuildContext context, AsyncSnapshot<NavBarItem> snapshot) {
              switch (snapshot.data) {
                case NavBarItem.home:
                  return HomeScreen(
                      movieRepository: widget.movieRepository);
                case NavBarItem.route:
                  return RouteScreen(movieRepository: widget.movieRepository);
                case NavBarItem.search:
                  return SearchScreen(movieRepository: widget.movieRepository);
                case NavBarItem.profile:
                  return ProfileScreen(profileName: widget.profileName,movieRepository: widget.movieRepository);
                default:
                  return Container();
              }
            },
          ),
          bottomNavigationBar: StreamBuilder(
            stream: _bottomNavBarBloc.itemStream,
            initialData: _bottomNavBarBloc.defaultItem,
            builder:
                (BuildContext context, AsyncSnapshot<NavBarItem> snapshot) {
              return Container(
                decoration: BoxDecoration(
                    border: Border(
                        top: BorderSide(
                            width: 0.5, color: Colors.black.withOpacity(0.4)))),
                child: BottomNavigationBar(
                  selectedItemColor: Colors.grey,
                  unselectedItemColor: Colors.black,
                  backgroundColor: Colors.white,
                  elevation: 0.9,
                  iconSize: 21,
                  unselectedFontSize: 10.0,
                  selectedFontSize: 12.0,
                  type: BottomNavigationBarType.fixed,
                  currentIndex: snapshot.data!.index,
                  onTap: _bottomNavBarBloc.pickItem,
                  items: [
                    BottomNavigationBarItem(
                      label: "Home",
                      icon: SizedBox(
                        child: SvgPicture.asset(
                          "assets/icons/home.svg",
                          colorFilter: const ColorFilter.mode(Colors.white, BlendMode.dstIn),
                          height: 25.0,
                          width: 25.0,
                        ),
                      ),
                      activeIcon: SizedBox(
                        child: SvgPicture.asset(
                          "assets/icons/home-active.svg",
                          colorFilter: const ColorFilter.mode(Colors.black, BlendMode.dstIn),
                          height: 25.0,
                          width: 25.0,
                        ),
                      ),
                    ),
                    BottomNavigationBarItem(
                      label: "Route",
                      icon: SvgPicture.asset(
                        "assets/icons/layers.svg",
                        colorFilter: const ColorFilter.mode(Colors.white, BlendMode.dstIn),
                        height: 25.0,
                        width: 25.0,
                      ),
                      activeIcon: SizedBox(
                        child: SvgPicture.asset(
                          "assets/icons/layers-active.svg",
                          colorFilter: const ColorFilter.mode(Colors.black, BlendMode.dstIn),
                          height: 25.0,
                          width: 25.0,
                        ),
                      ),
                    ),
                    BottomNavigationBarItem(
                      label: "Search",
                      icon: SvgPicture.asset(
                        "assets/icons/search.svg",
                        colorFilter: const ColorFilter.mode(Colors.white, BlendMode.dstIn),
                        height: 25.0,
                        width: 25.0,
                      ),
                      activeIcon: SizedBox(
                        child: SvgPicture.asset(
                          "assets/icons/search-active.svg",
                          colorFilter: const ColorFilter.mode(Colors.black, BlendMode.dstIn),
                          height: 25.0,
                          width: 25.0,
                        ),
                      ),
                    ),
                    BottomNavigationBarItem(
                      label: "Profile",
                      icon: SvgPicture.asset(
                        "assets/icons/profile.svg",
                        colorFilter: const ColorFilter.mode(Colors.white, BlendMode.dstIn),
                        height: 25.0,
                        width: 25.0,
                      ),
                      activeIcon: SizedBox(
                        child: SvgPicture.asset(
                          "assets/icons/profile-active.svg",
                          colorFilter: const ColorFilter.mode(Colors.black, BlendMode.dstIn),
                          height: 25.0,
                          width: 25.0,
                        ),
                      ),
                    ),
                  ],
                ),
              );
            },
          )),
    );
  }
}
