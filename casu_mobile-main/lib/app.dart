import 'package:casu_mobile/common/routes/pages.dart';
import 'package:flutter/material.dart';
import 'package:casu_mobile/common/theme/theme_controller.dart';

import 'package:flutter_screenutil/flutter_screenutil.dart';

import 'common/theme/style/custom_theme.dart';

/// The Widget that configures your application.
class App extends StatelessWidget {
  const App(
      {super.key,
      required this.themeController,
      });

  final ThemeController themeController;

  @override
  Widget build(BuildContext context) {
    return ScreenUtilInit(
      builder: (context, child) => MaterialApp(
        debugShowCheckedModeBanner: false,
        restorationScopeId: 'app',
        theme: ThemeData.light(useMaterial3: true),
        darkTheme:ThemeData.dark(useMaterial3: true),
        themeMode: ThemeMode.light,
        onGenerateRoute: AppPages.generateRouteSettings,
      ),
    );
  }
}
